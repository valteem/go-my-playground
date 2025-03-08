package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type metrics struct {
	metricGauge    prometheus.Gauge
	metricCounters *prometheus.CounterVec
}

func newMetrics(reg prometheus.Registerer) *metrics {

	m := &metrics{
		metricGauge: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "custom_gauge",
			Help: "custom gauge",
		}),
		metricCounters: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "custom_counter",
				Help: "custom counter",
			},
			[]string{"custom_label"}),
	}

	reg.MustRegister(m.metricGauge)
	reg.MustRegister(m.metricCounters)

	return m
}

func run() {

	reg := prometheus.NewRegistry()

	m := newMetrics(reg)

	m.metricGauge.Set(42.0)
	m.metricCounters.With(prometheus.Labels{"custom_label": "custom_counter_label"}).Inc()

	// reg is both Registry and Gatherer
	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg}))

	http.ListenAndServe(":3001", nil)

}
