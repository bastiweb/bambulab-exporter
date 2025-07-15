package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var LightsReportChamberLightMode = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "lights_report_chamber_light_mode",
	Help: "Chamber light mode (0 = off, 1 = on)",
}, []string{"printer"}))
