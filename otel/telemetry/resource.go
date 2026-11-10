package telemetry

import (
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.43.0"
)

// Set of attrs describing the entity that produces the telemetry (e.g. service, host, pod)
// Set once per provider, describes who emitted the data
// Applies to metrics and logs
// Backends lean on service.name a lot
func newResource(serviceName string) (*resource.Resource, error) {
	return resource.Merge(
		// Gives all defaults, but gets overridden with custom attrs
		// resource.Merge(a, b), when both resources set the same key, b's value wins
		resource.Default(),
		// semconv are standard attr names (service.name, http.request.method)
		// names change between versions: http.method was renamed to http.request.method, without version tag backend can't tell if they mean same thing
		resource.NewWithAttributes(semconv.SchemaURL, semconv.ServiceName(serviceName)),
	)
}
