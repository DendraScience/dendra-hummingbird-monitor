package cloud_function

import (
	"github.com/DendraScience/dendra_hummingbird_monitor/cloud_function/bq"
	"github.com/DendraScience/dendra_hummingbird_monitor/types"

	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
)

const (
	Package = "save_metrics"
)

var (
	count      int
	countex    sync.Mutex
	primary    string
	route      string
	auth       string
	v          bool
	version    bool
	localCount int
)

func init() {
	auth = os.Getenv("HUMMINGBIRD_KEY")
	flag.BoolVar(&version, "version", false, "Get detailed version string")
	flag.BoolVar(&v, "v", false, "Get detailed version string")
	flag.Parse()
	count = 0
}

func Ingest(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("key") != auth {
		w.WriteHeader(401)
		return
	}
	hostname := r.URL.Query().Get("hostname")
	if hostname == "" {
		w.WriteHeader(400)
		return
	}
	switch r.Method {
	case "GET":
		fallthrough
	case "HEAD":
		fallthrough
	case "DELETE":
		fallthrough
	case "PUT":
		fallthrough
	case "PATCH":
		fmt.Fprintf(w, "Only POST supported.")
		return
	case "POST":
		break
	default:
		// Should be impossible to reach, Methods are listened on explicitly
		log.Printf("<%d> [0] Unknown verb: %s", localCount, r.Method)
		return
	}
	fmt.Printf("Received metrics from host: %s\n", hostname)

	var buffer []byte
	var err error
	buffer, err = ioutil.ReadAll(r.Body)
	if err != nil {
		// TODO
		w.WriteHeader(400)
		log.Panic(err)
	}
	var data types.QuarterHourly
	json.Unmarshal(buffer, &data)
	bq.Insert(data)

	w.WriteHeader(200)
	localCount++
	return
}
