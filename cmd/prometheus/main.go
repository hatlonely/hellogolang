package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func recordMetrics() {
	go func() {
		for {
			counter.Inc()
			gauges.Add(10)
			histogram.Observe(float64(rand.Int63n(100)))
			counterVec.With(map[string]string{"key1": "val1", "key2": "val2"}).Inc()
			counterVec.WithLabelValues("val3", "val4").Inc()
			summary.Observe(rand.Float64())
			time.Sleep(2 * time.Second)
		}
	}()
}

var (
	counter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "hatlonely_counter",
		Help: "help counter",
	})
	gauges = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "hatlonely_gauge",
		Help: "help gauge",
	})
	histogram = promauto.NewHistogram(prometheus.HistogramOpts{
		Name:    "hatlonely_histogram",
		Help:    "help histogram",
		Buckets: []float64{10, 20, 40, 60, 80, 100},
		ConstLabels: map[string]string{
			"hello": "world",
		},
	})

	counterVec = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "hatlonely_counter_vec",
		Help: "help counter vec",
	}, []string{"key1", "key2"})

	summary = promauto.NewSummary(prometheus.SummaryOpts{
		Name:       "hatlonely_summary",
		Help:       "help summary",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	})
)

func main() {
	recordMetrics()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
