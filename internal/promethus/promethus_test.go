package promethus

import (
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	. "github.com/smartystreets/goconvey/convey"
)

func TestPromethus(t *testing.T) {
	Convey("test counter", t, func() {
		counter := prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "ns",
			Subsystem: "subsystem",
			Name:      "repository_pushes",
			Help:      "help xxx",
		})

		err := prometheus.Register(counter)
		So(err, ShouldBeNil)
		counter.Inc()
	})

	Convey("test gauge", t, func() {
		gauge := prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "our_company",
			Subsystem: "blob_storage",
			Name:      "ops_queued",
			Help:      "Number of blob storage operations waiting to be processed.",
		})
		err := prometheus.Register(gauge)
		So(err, ShouldBeNil)
		gauge.Inc()
		gauge.Dec()
		gauge.Add(10)
		gauge.Sub(10)
	})
}
