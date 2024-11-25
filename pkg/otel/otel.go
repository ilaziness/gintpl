package otel

import (
	"context"

	"gintpl/pkg/config"
	"gintpl/pkg/log"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	oteltrace "go.opentelemetry.io/otel/trace"
)

var (
	Tracer        oteltrace.Tracer
	tracerStarted = false
)

// InitTracer 初始化链路追踪
func InitTracer(serviceName string, cfg *config.Otel) {
	r, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(serviceName),
		),
	)
	if err != nil {
		panic(err)
	}
	exporter, err := otlptracehttp.New(context.Background(), otlptracehttp.WithEndpointURL(cfg.TraceExporterURL))
	if err != nil {
		panic(err)
	}
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(r),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	tracerStarted = true
	Tracer = otel.Tracer(serviceName)
}

func Shutdown() {
	if !tracerStarted {
		return
	}
	tracerStarted = false
	tp := otel.GetTracerProvider()
	if newTp, ok := tp.(*sdktrace.TracerProvider); ok {
		if err := newTp.Shutdown(context.Background()); err != nil {
			log.Logger.Warnf("Tracer shutdown err: %v", err)
		}
	}
	log.Logger.Infof("Tracer shutdown")
}
