package main

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/baggage"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.43.0"
	"go.opentelemetry.io/otel/trace"
)

// Module level, concurrency safe
var tracer trace.Tracer

// Send data to the otel collector
func newExporter(ctx context.Context) (sdktrace.SpanExporter, error) {
	// some exporter, how the telemetry is set out
	return stdouttrace.New(stdouttrace.WithPrettyPrint())
}

func newTracerProvider(ctx context.Context) (*sdktrace.TracerProvider, error) {
	// Set of attrs describing the entity that produces the telemetry (e.g. service, host, pod)
	// Set once per provider, describes who emitted the data
	// Applies to metrics and logs
	// Backends lean on service.name a lot
	r, err := resource.Merge(
		// Gives all defaults, but gets overridden with custom attrs
		// resource.Merge(a, b), when both resources set the same key, b's value wins
		resource.Default(),
		// semconv are standard attr names (service.name, http.request.method)
		// names change between versions: http.method was renamed to http.request.method, without version tag backend can't tell if they mean same thing
		resource.NewWithAttributes(semconv.SchemaURL, semconv.ServiceName("pug-api")),
	)

	if err != nil {
		return nil, err
	}

	exp, _ := newExporter(ctx)

	return sdktrace.NewTracerProvider(
		sdktrace.WithResource(r),
		sdktrace.WithSyncer(exp),
	), nil
}

func main() {
	ctx := context.Background()

	provider, _ := newTracerProvider(ctx)

	// Otel package is global registry, set things once at startup,
	// any code in process can read it back
	//
	// Each setter has a getter
	otel.SetTracerProvider(provider)
	// tp := otel.GetTracerProvider()
	otel.SetTextMapPropagator(
		// propagates traceparent and tracestate
		propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{},
			propagation.Baggage{},
		),
	)

	tmp := otel.GetTextMapPropagator()

	defer func() {
		_ = provider.Shutdown(ctx)
	}()

	tracer = provider.Tracer("tracer")

	tier, _ := baggage.NewMember("user.tier", "premium")
	tenant, _ := baggage.NewMember("tenant.id", "acme")
	bag, _ := baggage.New(tier, tenant)
	ctx = baggage.ContextWithBaggage(ctx, bag)

	spanCtx, span := tracer.Start(ctx, "new span")
	defer span.End()

	carrier := propagation.MapCarrier{}
	tmp.Inject(spanCtx, carrier)
	// ASCII string which carries info on the current span:
	// version (00), trace-id, parent-id (id of sender's curr span), trace flags (01 means caller is sampling, 00 means it isn't)
	fmt.Println("traceparent ", carrier["traceparent"])

	// comma-separated list of key=value entries
	// any software which takes palce in creating or forwarding trace contexts between processes
	// Systems add this, key says which system owns the entry, value is system's private metadata
	// dd=s:1;ot=th:8
	// The key is a namespace. dd is Datadog, ot is OpenTelemetry, etc...
	// The value is opaque to everyone except the owner, no other system parses it
	fmt.Println("tracestate: ", carrier["tracestate"])

	// User defined k/v pairs
	// All downstream things can see it
	fmt.Println("baggage: ", carrier["baggage"])
}
