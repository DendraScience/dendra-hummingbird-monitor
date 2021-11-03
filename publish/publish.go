package publish

import (
	"bytes"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/DendraScience/dendra_hummingbird_monitor/proc"
	log "github.com/sirupsen/logrus"
)

func init() {
	var seed int64
	hostName := proc.GetHostname()

	for _, x := range hostName {
		seed += int64(x)
	}
	rand.Seed(seed)
}

func Post(jsonStr string, URL string, auth string, hostname string) {
	n := rand.Intn(10)
	time.Sleep(time.Duration(n) * time.Second)
	for retries := 0; retries <= 5; retries++ {
		var jsonBuf = []byte(jsonStr)
		endpoint := fmt.Sprintf("%s?key=%s&hostname=%s", URL, auth, hostname)
		req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonBuf))
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		client.Timeout = time.Second * 15
		resp, err := client.Do(req)
		if err != nil {
			log.Errorf("Error posting data: %v", err)
			return
		}
		defer resp.Body.Close()
		log.Printf("Data posted to: %s", endpoint)

		if resp.StatusCode == http.StatusOK {
			return
		}
		n = rand.Intn(10)
		time.Sleep(time.Minute * time.Duration(retries*n))
	}

}
