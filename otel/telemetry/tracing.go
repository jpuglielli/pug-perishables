package telemetry

import (
	"context"

	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

// Send data to the otel collector
func newExporter(ctx context.Context) (sdktrace.SpanExporter, error) {
	// some exporter, how the telemetry is set out
	return stdouttrace.New(stdouttrace.WithPrettyPrint())
}

func NewTracerProvider(ctx context.Context, serviceName string) (*sdktrace.TracerProvider, error) {
	r, err := newResource(serviceName)
	if err != nil {
		return nil, err
	}

	exp, err := newExporter(ctx)
	if err != nil {
		return nil, err
	}

	return sdktrace.NewTracerProvider(
		sdktrace.WithResource(r),
		sdktrace.WithSyncer(exp),
	), nil
}
