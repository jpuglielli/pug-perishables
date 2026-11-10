package telemetry

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

// The global propagator defaults to a no-op, so Inject writes nothing until this is called
func SetupPropagators() {
	otel.SetTextMapPropagator(
		// TraceContext propagates traceparent and tracestate, Baggage propagates baggage
		propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{},
			propagation.Baggage{},
		),
	)
}
