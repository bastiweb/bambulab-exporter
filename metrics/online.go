package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var OnlineAhb = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "online_ahb",
	Help: "Online AHB status (1 = online, 0 = offline)",
}, []string{"printer"}))

var OnlineRfid = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "online_rfid",
	Help: "Online RFID status (1 = online, 0 = offline)",
}, []string{"printer"}))

var OnlineVersion = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "online_version",
	Help: "Online version",
}, []string{"printer"}))
