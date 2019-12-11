package fetchData

import (
	"context"
	"fmt"
	"html"
	"net/http"
	"strings"
	"time"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/iterator"
)

func fetchData(w http.ResponseWriter, r *http.Request) {
	//_, ok := r.URL.Query()["key"]
	//if !ok {
	//    fmt.Fprint(w, "Url Param 'key' is missing")
	//    return
	//}
	// get project info
	//proj := os.Getenv("GOOGLE_CLOUD_PROJECT")
	//if proj == "" {
	//        fmt.Println("GOOGLE_CLOUD_PROJECT environment variable must be set.")
	//        os.Exit(1)
	//}
	// code for big query part.
	ctx := context.Background()
	bq_client, err := bigquery.NewClient(ctx, `micro-reef-210013`)
	if err != nil {
		fmt.Fprint(w, "There was some issue, cannot create bqclient")
	}
	beforeQuery := time.Now()
	bq_query := "SELECT 1 AS test FROM `micro-reef-210013.kb.c15_fkblty` LIMIT 10"
	query_job := bq_client.Query(bq_query)
	it, err := query_job.Read(ctx)
	queryTime := time.Now().Sub(beforeQuery)
	if err != nil {
		fmt.Fprint(w, "There was some issue in reading the job")
	}
	var result []string
	counter := 0
	for {
		var c []bigquery.Value
		err := it.Next(&c)
		if err == iterator.Done {
			break
		}
		result = append(result, fmt.Sprintf("%d", c[0]))
		counter++
		if err != nil {
			fmt.Fprint(w, "nothing was found")
		}
	}
	fmt.Fprint(w, html.EscapeString(fmt.Sprintf("The time taken was %f and records traversed were %d and the record is %s", queryTime.Seconds(), counter, strings.Join(result, ","))))
}

// BQ query to run
// SELECT 1 as this FROM `micro-reef-210013.kb.c15_fkblty`
// - runs through 0 byte of data.
