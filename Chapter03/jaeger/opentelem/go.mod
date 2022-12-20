module chapter.3/trace

go 1.17

require (
	go.opentelemetry.io/otel v1.2.0
	go.opentelemetry.io/otel/exporters/jaeger v1.2.0
	go.opentelemetry.io/otel/sdk v1.2.0
	go.opentelemetry.io/otel/trace v1.2.0
)

require golang.org/x/sys v0.0.0-20210423185535-09eb48e85fd7 // indirect
replace chapter.3/trace => ./
