package trace

import (
	"context"

	"go.opentelemetry.io/otel"

	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	sc "go.opentelemetry.io/otel/semconv/v1.4.0"
)

type ShutdownTracing func(ctx context.Context) error

func InitTracing(service string) (ShutdownTracing, error) {
	// Create the Jaeger exporter.
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint())
	if err != nil {
		return func(ctx context.Context) error { return nil }, err
	}

	// Create the TracerProvider.
	tp := trace.NewTracerProvider(
		trace.WithBatcher(exp),
		trace.WithResource(resource.NewWithAttributes(
			sc.SchemaURL,
			sc.ServiceNameKey.String(service),
		)),
	)

	otel.SetTracerProvider(tp)
	return tp.Shutdown, nil
}
