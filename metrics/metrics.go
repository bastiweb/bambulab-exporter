package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

const metricPrefix = "bambulab_"

func makeCollector[C prometheus.Collector](collector C) C {
	prometheus.MustRegister(collector)
	return collector
}
