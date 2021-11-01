package dendra_hummingbird_monitor

import (
	"net/http"

	cloud_function "github.com/DendraScience/dendra_hummingbird_monitor/cloud_function"
)

func Ingest(w http.ResponseWriter, r *http.Request) {
	cloud_function.Ingest(w, r)
}
