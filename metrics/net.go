package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var NetConf = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "net_conf",
	Help: "Network configuration",
}, []string{"printer"}))
