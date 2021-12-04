package main

import (
	m "chapter.three/prometheus/metric"

	"context"
	"github.com/gorilla/mux"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/global"
	"log"
	"math/rand"
	"net/http"
)

const serviceName = "samplemetrics"

func main() {
	sMetrics, err := m.InitMetrics(serviceName)
	if err != nil {
		log.Fatalf("Failed to setup metrics: %v\n", err)
	}
	defer func() {
		if err := sMetrics(context.Background()); err != nil {
			log.Printf("Failed to shutdown metrics: %v\n", err)
		}
	}()

	r := mux.NewRouter()

	//setup automated tracing for incoming http request
	r.Use(otelmux.Middleware("my-api-server"))

	//setup meter and use `samplemetrics` as servicename
	meter := global.Meter("samplemetrics")

	//measure total number of requests
	ctr, err := meter.NewInt64Counter("metric.totalrequest", metric.WithDescription("description"))

	//setup handler for rqeuest
	r.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Reporting metric metric.totalrequest")

		ctx := r.Context()

		//add request metric counter
		ctr.Add(ctx, 1)

		rw.WriteHeader(http.StatusOK)

		//send back something to user
		rw.Header().Add("Content-Type", "application/text")
		rw.Write([]byte("All good!\n"))
	}).Methods("GET")

	createGauge(context.Background(), meter)

	log.Println("Starting up sever on port 8000")

	log.Fatal(http.ListenAndServe(":8000", r))
}

//createGauge function to create gauge observer metric
func createGauge(ctx context.Context, meter metric.Meter) {
	_, _ = meter.NewInt64GaugeObserver("metric.random",
		func(ctx context.Context, result metric.Int64ObserverResult) {
			min := 100
			max := 300
			log.Println("Reporting metric metric.random")
			result.Observe(int64(rand.Intn(max-min) + min))
		},
		metric.WithDescription("Random numbers"),
	)
}
