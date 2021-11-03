package bq

import (
	"context"
	"fmt"
	"log"
	"os"

	. "github.com/DendraScience/dendra_hummingbird_monitor/types"

	"cloud.google.com/go/bigquery"
)

var (
	projectID          string
	hostDataSet        string
	hostTableName      string
	containerDataSet   string
	containerTableName string
)

func init() {
	projectID = os.Getenv("GOOGLE_CLOUD_PROJECT")
	containerDataSet = os.Getenv("BIGQUERY_DATASET_CONTAINERS")
	containerTableName = os.Getenv("BIGQUERY_TABLE_CONTAINERS")
	hostDataSet = os.Getenv("BIGQUERY_DATASET_HOST")
	hostTableName = os.Getenv("BIGQUERY_TABLE_HOST")
	if projectID == "" {
		fmt.Println("GOOGLE_CLOUD_PROJECT environment variable must be set.")
		os.Exit(1)
	}
	if hostDataSet == "" {
		fmt.Println("BIGQUERY_DATASET_HOST environment variable must be set.")
		os.Exit(1)
	}
	if hostTableName == "" {
		fmt.Println("BIGQUERY_TABLE_HOST environment variable must be set.")
		os.Exit(1)
	}
	if containerDataSet == "" {
		fmt.Println("BIGQUERY_DATASET_CONTAINER environment variable must be set.")
		os.Exit(1)
	}
	if containerTableName == "" {
		fmt.Println("BIGQUERY_TABLE_CONTAINER environment variable must be set.")
		os.Exit(1)
	}
}

func Insert(data QuarterHourly) {
	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		log.Panicf("bigquery.NewClient: %v", err)
	}
	defer client.Close()

	err = queryMetrics(ctx, client, data)
	if err != nil {
		log.Panicf("bigquery insertion fail: %v", err)
	}

	err = queryContainers(ctx, client, data)
	if err != nil {
		log.Panicf("bigquery insertion fail: %v", err)
	}
}

func queryMetrics(ctx context.Context, client *bigquery.Client, data QuarterHourly) error {
	qstring := `INSERT INTO ` + projectID + "." + hostDataSet + "." + hostTableName + `(hostname, version, timestamp, collection_time, disk_usage, disk_free, lan_bytes_down, lan_bytes_up, memory_total, memory_buffered, memory_free, memory_percent, memory_avail, processor_count, load_average, cpu_percent, num_packages, updates_available, wan_bytes_down, wan_bytes_up, uptime, wan_ip, lan_ip) VALUES`
	qstring += "(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
	query := client.Query(qstring)
	query.Parameters = []bigquery.QueryParameter{
		{Value: data.Hostname},
		{Value: data.Version},
		{Value: data.Timestamp},
		{Value: data.CollectionTime},
		{Value: data.DiskUsage},
		{Value: data.DiskFree},
		{Value: data.LANBytesDown},
		{Value: data.LANBytesUp},
		{Value: data.MemTotal},
		{Value: data.MemBuffered},
		{Value: data.MemFree},
		{Value: data.MemPercent},
		{Value: data.MemAvail},
		{Value: data.ProcessorCount},
		{Value: data.LoadAverage},
		{Value: data.CPU_Percent},
		{Value: data.NumPackages},
		{Value: data.UpdatesAvail},
		{Value: data.WANBytesDown},
		{Value: data.WANBytesUp},
		{Value: data.Uptime},
		{Value: data.WANIP},
		{Value: data.LANIP},
	}

	job, err := query.Run(ctx)
	if err != nil {
		log.Printf("Error creating metrics query job: %s", err.Error())
		log.Printf("Query: %s", qstring)
		return err
	}
	status, err := job.Wait(ctx)
	if err != nil {
		log.Printf("Error running metrics query: %s", err.Error())
		return err
	}
	err = status.Err()
	return err
}
func queryContainers(ctx context.Context, client *bigquery.Client, data QuarterHourly) error {
	qstring := `INSERT INTO ` + projectID + "." + containerDataSet + "." + containerTableName + `(hostname, version, id, timestamp, image, name, created, cpu_percent, memory_usage, memory_allowed, memory_percent, uptime) VALUES`
	qps := []bigquery.QueryParameter{}
	for i, container := range data.Containers {
		if i == len(data.Containers)-1 {
			qstring += "(?,?,?,?,?,?,?,?,?,?,?,?);"
		} else {
			qstring += "(?,?,?,?,?,?,?,?,?,?,?,?),"
		}
		qps = append(qps,
			[]bigquery.QueryParameter{{Value: data.Hostname},
				{Value: data.Version},
				{Value: container.ID},
				{Value: data.Timestamp},
				{Value: container.Image},
				{Value: container.Name},
				{Value: container.Created},
				{Value: container.CPU},
				{Value: container.MemUsage},
				{Value: container.MemAllowed},
				{Value: container.MemPercent},
				{Value: container.Uptime}}...)
	}
	query := client.Query(qstring)
	query.Parameters = qps
	job, err := query.Run(ctx)
	if err != nil {
		log.Printf("Error creating container query job: %s", err.Error())
		log.Printf("Query: %s", qstring)
		return err
	}
	stat, err := job.Wait(ctx)
	if err != nil {
		log.Printf("Error running container query: %s", err.Error())
		return err
	}
	if stat.Err() != nil {
		return stat.Err()
	}

	return nil
}
