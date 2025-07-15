package collector

import (
	"github.com/Scrin/bambulab-exporter/config"
	"github.com/Scrin/bambulab-exporter/metrics"
	"github.com/rs/zerolog/log"
)

func receiveOnlineReport(report map[string]any) {
	for key, value := range report {
		switch key {
		case "ahb":
			if safeBool(value) {
				metrics.OnlineAhb.WithLabelValues(config.PrinterName).Set(1)
			} else {
				metrics.OnlineAhb.WithLabelValues(config.PrinterName).Set(0)
			}
		case "rfid":
			if safeBool(value) {
				metrics.OnlineRfid.WithLabelValues(config.PrinterName).Set(1)
			} else {
				metrics.OnlineRfid.WithLabelValues(config.PrinterName).Set(0)
			}
		case "version":
			metrics.OnlineVersion.WithLabelValues(config.PrinterName).Set(safeFloat64(value))
		default:
			if isNewUnknownField("online", key) {
				log.Warn().Str("key", key).Interface("value", value).Msg("Unknown field in online report")
			}
		}
	}
}
