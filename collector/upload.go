package collector

import (
	"github.com/Scrin/bambulab-exporter/config"
	"github.com/Scrin/bambulab-exporter/metrics"
	"github.com/rs/zerolog/log"
)

func receiveUploadReport(report map[string]any) {
	for key, value := range report {
		switch key {
		case "progress":
			metrics.UploadProgress.WithLabelValues(config.PrinterName).Set(safeFloat64(value))
		case "status":
			// Convert string status to numeric
			statusStr := safeString(value)
			var statusValue float64
			switch statusStr {
			case "idle":
				statusValue = 0
			case "uploading":
				statusValue = 1
			case "completed":
				statusValue = 2
			case "error":
				statusValue = 3
			default:
				statusValue = 0
			}
			metrics.UploadStatus.WithLabelValues(config.PrinterName).Set(statusValue)
		case "message":
		default:
			if isNewUnknownField("upload", key) {
				log.Warn().Str("key", key).Interface("value", value).Msg("Unknown field in upload report")
			}
		}
	}
}
