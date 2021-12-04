package metric

import (
	"context"
	"go.opentelemetry.io/otel/sdk/export/metric/aggregation"
	"net/http"

	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/metric/global"
	controller "go.opentelemetry.io/otel/sdk/metric/controller/basic"
	processor "go.opentelemetry.io/otel/sdk/metric/processor/basic"
	selector "go.opentelemetry.io/otel/sdk/metric/selector/simple"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

type ShutdownMetrics func(ctx context.Context) error

// InitMetrics use Prometheus exporter
func InitMetrics(service string) (ShutdownMetrics, error) {
	config := prometheus.Config{}
	c := controller.New(
		processor.NewFactory(
			selector.NewWithExactDistribution(),
			aggregation.CumulativeTemporalitySelector(),
			processor.WithMemory(true),
		),
		controller.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(service),
		)),
	)
	exporter, err := prometheus.New(config, c)
	if err != nil {
		return func(ctx context.Context) error { return nil }, err
	}

	global.SetMeterProvider(exporter.MeterProvider())

	srv := &http.Server{Addr: ":2112", Handler: exporter}
	go func() {
		_ = srv.ListenAndServe()
	}()

	return srv.Shutdown, nil
}
