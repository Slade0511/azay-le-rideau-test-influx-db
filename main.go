package main

import (
	"fmt"
	"github.com/influxdata/influxdb-client-go/v2"
	"time"
)

func main() {
	// Set a log level constant
	const debugLevel uint = 4

	/**
	  * Instantiate a client with a configuration object
	  * that contains your InfluxDB URL and token.
	**/

	clientOptions := influxdb2.DefaultOptions().
		SetBatchSize(20).
		SetLogLevel(debugLevel).
		SetPrecision(time.Second)

	client := influxdb2.NewClientWithOptions("https://us-east-1-1.aws.cloud2.influxdata.com",
		"JQcQA2A21cCh2S9Zf7PHd8eDiFg_6sp1g5_6cXNpNavlBmLpaN4rsgdV4TfcY90T8xoBskiGdbtdRWdty7o6VQ==",
		clientOptions)

	/**
	  * Create an asynchronous, non-blocking write client.
	  * Provide your InfluxDB org and bucket as arguments
	**/
	writeAPI := client.WriteAPI("a56f745d70b85346", "test")

	// Get the errors channel for the asynchronous write client.
	errorsCh := writeAPI.Errors()

	/** Create a point.
	  * Provide measurement, tags, and fields as arguments.
	**/ //  Air
	/*p := influxdb2.NewPointWithMeasurement("Air").
		AddField("humidity", 80).
		AddField("temperature", 22).
		SetTime(time.Now())*/
      //  Habitation
    p := influxdb2.NewPointWithMeasurement("Habitation").
		AddField("intrusion", 0).
		AddField("temperature", 25).
		AddField("luminosite", 2000).
		SetTime(time.Now())

	// Define a proc for handling errors.
	go func() {
		for err := range errorsCh {
			fmt.Printf("write error: %s\n", err.Error())
		}
	}()

	// Write the point asynchronously
	writeAPI.WritePoint(p)

	// Send pending writes from the buffer to the bucket.
	writeAPI.Flush()

	// Ensure background processes finish and release resources.
	client.Close()
}
