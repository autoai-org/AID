package daemon

import (
	"log"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	stdout "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	oteltrace "go.opentelemetry.io/otel/trace"
)

func initTracer() *sdktrace.TracerProvider {
	exporter, err := stdout.New(stdout.WithPrettyPrint())
	if err != nil {
		log.Fatal(err)
	}
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exporter),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	return tp
}

func getUser(c *gin.Context, id string) string {
	// Pass the built-in `context.Context` object from http.Request to OpenTelemetry APIs
	// where required. It is available from gin.Context.Request.Context()
	_, span := tracer.Start(c.Request.Context(), "getUser", oteltrace.WithAttributes(attribute.String("id", id)))
	defer span.End()
	if id == "123" {
		return "otelgin tester"
	}
	return "unknown"
}
