package metric

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Metric struct {
	RequestCounter   *prometheus.CounterVec
	RequestHistogram *prometheus.HistogramVec
}

func NewMetric() *Metric {
	return &Metric{
		RequestCounter: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "http_request_total",
				Help: "Total number of HTTP requests",
			},
			[]string{"method"},
		),
		RequestHistogram: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Name: "http_request_duration_seconds",
				Help: "Duration of HTTP requests.",
			},
			[]string{"method"},
		),
	}
}
