package observability

import (
	"context"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
)

// Initilize a meter provider and set it as the global meter provider
func InitMeter(ctx context.Context, endPoint string, serviceName string) (func(context.Context) error, error) {
	metricExporter, err := otlpmetricgrpc.New(ctx,
		otlpmetricgrpc.WithInsecure(),
		otlpmetricgrpc.WithEndpoint(endPoint),
	)
	if err != nil {
		return nil, err
	}

	resource, err := newResource(serviceName)
	if err != nil {
		return nil, err
	}

	meterProvider, err := newMeterProvider(ctx, metricExporter, endPoint, resource)
	if err != nil {
		return nil, err
	}

	otel.SetMeterProvider(meterProvider)
	return meterProvider.Shutdown, nil
}

func newMeterProvider(ctx context.Context, exp *otlpmetricgrpc.Exporter, endPoint string, res *resource.Resource) (*metric.MeterProvider, error) {
	meterProvider := metric.NewMeterProvider(
		metric.WithResource(res),
		metric.WithReader(metric.NewPeriodicReader(
			exp,
			metric.WithInterval(3*time.Second),
		)),
	)

	return meterProvider, nil
}
