package collector

import (
	"github.com/bastiweb/bambulab-exporter/config"
	"github.com/bastiweb/bambulab-exporter/metrics"
	"github.com/rs/zerolog/log"
)

func receiveIpcamReport(report map[string]any) {
	for key, value := range report {
		switch key {
		case "ipcam_dev":
			if safeString(value) == "1" {
				metrics.IpcamDev.WithLabelValues(config.PrinterName).Set(1)
			} else {
				metrics.IpcamDev.WithLabelValues(config.PrinterName).Set(0)
			}
		case "ipcam_record":
			if safeString(value) == "enable" {
				metrics.IpcamRecord.WithLabelValues(config.PrinterName).Set(1)
			} else {
				metrics.IpcamRecord.WithLabelValues(config.PrinterName).Set(0)
			}
		case "timelapse":
			if safeString(value) == "enable" {
				metrics.IpcamTimelapse.WithLabelValues(config.PrinterName).Set(1)
			} else {
				metrics.IpcamTimelapse.WithLabelValues(config.PrinterName).Set(0)
			}
		case "tutk_server":
			if safeString(value) == "enable" {
				metrics.IpcamTutkServer.WithLabelValues(config.PrinterName).Set(1)
			} else {
				metrics.IpcamTutkServer.WithLabelValues(config.PrinterName).Set(0)
			}
		case "mode_bits":
			metrics.IpcamModeBits.WithLabelValues(config.PrinterName).Set(safeFloat64(value))
		case "resolution":
		default:
			if isNewUnknownField("ipcam", key) {
				log.Warn().Str("key", key).Interface("value", value).Msg("Unknown field in ipcam report")
			}
		}
	}
}
