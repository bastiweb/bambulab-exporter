package collector

import (
	"github.com/bastiweb/bambulab-exporter/config"
	"github.com/bastiweb/bambulab-exporter/metrics"
	"github.com/rs/zerolog/log"
)

func receiveVtTrayReport(report map[string]any) {
	for key, value := range report {
		switch key {
		case "id":
			metrics.VtTrayId.WithLabelValues(config.PrinterName).Set(safeFloat64(value))
		case "remain":
			metrics.VtTrayRemain.WithLabelValues(config.PrinterName).Set(safeFloat64(value))
		case "k":
			metrics.VtTrayK.WithLabelValues(config.PrinterName).Set(safeFloat64(value))
		case "n":
			metrics.VtTrayN.WithLabelValues(config.PrinterName).Set(safeFloat64(value))
		case "cali_idx":
			metrics.VtTrayCaliIdx.WithLabelValues(config.PrinterName).Set(safeFloat64(value))
		case "tray_diameter":
			metrics.VtTrayDiameter.WithLabelValues(config.PrinterName).Set(safeFloat64(value))
		case "tray_temp":
			metrics.VtTrayTemp.WithLabelValues(config.PrinterName).Set(safeFloat64(value))
		case "tray_time":
			metrics.VtTrayTime.WithLabelValues(config.PrinterName).Set(safeFloat64(value))
		case "bed_temp_type":
			metrics.VtTrayBedTempType.WithLabelValues(config.PrinterName).Set(safeFloat64(value))
		case "bed_temp":
			metrics.VtTrayBedTemp.WithLabelValues(config.PrinterName).Set(safeFloat64(value))
		case "nozzle_temp_max":
			metrics.VtTrayNozzleTempMax.WithLabelValues(config.PrinterName).Set(safeFloat64(value))
		case "nozzle_temp_min":
			metrics.VtTrayNozzleTempMin.WithLabelValues(config.PrinterName).Set(safeFloat64(value))
		case "ctype":
			metrics.VtTrayCtype.WithLabelValues(config.PrinterName).Set(safeFloat64(value))
		case "tag_uid":
		case "tray_id_name":
		case "tray_info_idx":
		case "tray_type":
		case "tray_sub_brands":
		case "tray_color":
		case "tray_weight":
		case "tray_uuid":
		case "xcam_info":
		case "cols":
		default:
			if isNewUnknownField("vt_tray", key) {
				log.Warn().Str("key", key).Interface("value", value).Msg("Unknown field in vt_tray report")
			}
		}
	}
}
