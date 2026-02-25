package collector

import (
	"github.com/bastiweb/bambulab-exporter/config"
	"github.com/bastiweb/bambulab-exporter/metrics"
	"github.com/rs/zerolog/log"
)

func receiveNetReport(report map[string]any) {
	for key, value := range report {
		switch key {
		case "conf":
			metrics.NetConf.WithLabelValues(config.PrinterName).Set(safeFloat64(value))
		case "info":
		default:
			if isNewUnknownField("net", key) {
				log.Warn().Str("key", key).Interface("value", value).Msg("Unknown field in net report")
			}
		}
	}
}
