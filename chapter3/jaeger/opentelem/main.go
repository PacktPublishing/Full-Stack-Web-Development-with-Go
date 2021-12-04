package main

import (
	t "chapter.3/trace/trace"
	"context"
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"log"
	"sync"
	"time"
)

const serviceName = "tracing"

func main() {
	sTracing, err := t.InitTracing(serviceName)
	if err != nil {
		log.Fatalf("Failed to setup tracing: %v\n", err)
	}
	defer func() {
		if err := sTracing(context.Background()); err != nil {
			log.Printf("Failed to shutdown tracing: %v\n", err)
		}
	}()
	var span trace.Span
	ctx, span := otel.Tracer(serviceName).Start(context.Background(), "outside")
	defer span.End()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		_, s := otel.Tracer(serviceName).Start(ctx, "inside")
		defer s.End()
		time.Sleep(1 * time.Second)
		s.SetAttributes(attribute.String("sleep", "done"))
		s.SetAttributes(attribute.String("go func", "1"))
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		_, ss := otel.Tracer(serviceName).Start(ctx, "inside")
		defer ss.End()
		time.Sleep(2 * time.Second)
		ss.SetAttributes(attribute.String("sleep", "done"))
		ss.SetAttributes(attribute.String("go func", "2"))
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("\nDone!")
}
