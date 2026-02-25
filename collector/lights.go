package collector

import (
	"github.com/bastiweb/bambulab-exporter/config"
	"github.com/bastiweb/bambulab-exporter/metrics"
)

func receiveLightsReport(report []any) {
	for _, light := range report {
		if lightMap, ok := light.(map[string]any); ok {
			if node, exists := lightMap["node"]; exists && safeString(node) == "chamber_light" {
				if mode, exists := lightMap["mode"]; exists {
					if safeString(mode) == "on" {
						metrics.LightsReportChamberLightMode.WithLabelValues(config.PrinterName).Set(1)
					} else {
						metrics.LightsReportChamberLightMode.WithLabelValues(config.PrinterName).Set(0)
					}
				}
			}
		}
	}
}
