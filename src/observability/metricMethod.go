package observability

import (
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel/metric"
)

func Counter(name string, desc string, meter metric.Meter) (metric.Int64Counter, error){
	var err error
	ApiCounter, err := meter.Int64Counter(
		name , 
		metric.WithDescription(desc),
		metric.WithUnit("{API}"),
	)
	if err != nil {
		log.Err(err).Msg(err.Error())
		return nil, err 
	}
	return ApiCounter, nil 
}

func Histogram(name string, desc string, meter metric.Meter) (metric.Int64Histogram, error){
	var err error
	ApiHistogram, err := meter.Int64Histogram(
		name , 
		metric.WithDescription(desc),
		metric.WithUnit("{API}"),
	)
	if err != nil {
		log.Err(err).Msg(err.Error())
		return nil, err 
	}
	return ApiHistogram, nil 
}