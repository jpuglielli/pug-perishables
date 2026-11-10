package main

import (
	"context"
	"fmt"
	"log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/baggage"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"

	"github.com/jpuglielli/pug-perishables/otel/telemetry"
)

// Module level, concurrency safe
var tracer trace.Tracer

func main() {
	ctx := context.Background()

	provider, err := telemetry.NewTracerProvider(ctx, "pug-api")
	if err != nil {
		log.Fatal(err)
	}

	// Otel package is global registry, set things once at startup,
	// any code in process can read it back
	//
	// Each setter has a getter
	otel.SetTracerProvider(provider)
	// tp := otel.GetTracerProvider()
	telemetry.SetupPropagators()

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
