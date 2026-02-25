package collector

import (
	"strconv"

	"github.com/bastiweb/bambulab-exporter/config"
	"github.com/bastiweb/bambulab-exporter/metrics"
	"github.com/rs/zerolog/log"
)

func receiveAmsReport(report map[string]any) {
	for key, value := range report {
		switch key {
		case "ams_exist_bits":
			metrics.AmsExistBits.WithLabelValues(config.PrinterName).Set(safeFloat64(value))
		case "tray_exist_bits":
			metrics.TrayExistBits.WithLabelValues(config.PrinterName).Set(safeFloat64(value))
		case "tray_is_bbl_bits":
			metrics.TrayIsBblBits.WithLabelValues(config.PrinterName).Set(safeFloat64(value))
		case "tray_tar":
			metrics.TrayTarget.WithLabelValues(config.PrinterName).Set(safeFloat64(value))
		case "tray_now":
			metrics.TrayCurrent.WithLabelValues(config.PrinterName).Set(safeFloat64(value))
		case "tray_pre":
			metrics.TrayPrevious.WithLabelValues(config.PrinterName).Set(safeFloat64(value))
		case "tray_read_done_bits":
			metrics.TrayReadDoneBits.WithLabelValues(config.PrinterName).Set(safeFloat64(value))
		case "tray_reading_bits":
			metrics.TrayReadingBits.WithLabelValues(config.PrinterName).Set(safeFloat64(value))
		case "version":
			metrics.AmsVersion.WithLabelValues(config.PrinterName).Set(safeFloat64(value))
		case "insert_flag":
			if safeBool(value) {
				metrics.AmsInsertFlag.WithLabelValues(config.PrinterName).Set(1)
			} else {
				metrics.AmsInsertFlag.WithLabelValues(config.PrinterName).Set(0)
			}
		case "power_on_flag":
			if safeBool(value) {
				metrics.AmsPowerOnFlag.WithLabelValues(config.PrinterName).Set(1)
			} else {
				metrics.AmsPowerOnFlag.WithLabelValues(config.PrinterName).Set(0)
			}
		case "ams":
			if amsArray, ok := value.([]any); ok {
				receiveAmsArray(amsArray)
			}
		default:
			if isNewUnknownField("ams", key) {
				log.Warn().Str("key", key).Interface("value", value).Msg("Unknown field in ams report")
			}
		}
	}
}

func receiveAmsArray(amsArray []any) {
	for amsIndex, amsValue := range amsArray {
		if amsMap, ok := amsValue.(map[string]any); ok {
			amsId := strconv.Itoa(amsIndex)

			// Process AMS-level metrics
			for key, value := range amsMap {
				switch key {
				case "id":
					// AMS ID is already known from array index
				case "humidity":
					metrics.AmsHumidity.WithLabelValues(config.PrinterName, amsId).Set(safeFloat64(value))
				case "humidity_raw":
					metrics.AmsHumidityRaw.WithLabelValues(config.PrinterName, amsId).Set(safeFloat64(value))
				case "temp":
					metrics.AmsTemp.WithLabelValues(config.PrinterName, amsId).Set(safeFloat64(value))
				case "dry_time":
					metrics.AmsDryTime.WithLabelValues(config.PrinterName, amsId).Set(safeFloat64(value))
				case "info":
					metrics.AmsInfo.WithLabelValues(config.PrinterName, amsId).Set(safeFloat64(value))
				case "tray":
					if trayArray, ok := value.([]any); ok {
						receiveAmsTrayArray(amsId, trayArray)
					}
				default:
					if isNewUnknownField("ams_"+amsId, key) {
						log.Warn().Str("ams", amsId).Str("key", key).Interface("value", value).Msg("Unknown field in AMS report")
					}
				}
			}
		}
	}
}

func receiveAmsTrayArray(amsId string, trayArray []any) {
	for trayIndex, trayValue := range trayArray {
		if trayMap, ok := trayValue.(map[string]any); ok {
			trayId := strconv.Itoa(trayIndex)

			for key, value := range trayMap {
				switch key {
				case "id":
					metrics.AmsTrayId.WithLabelValues(config.PrinterName, amsId, trayId).Set(safeFloat64(value))
				case "state":
					metrics.AmsTrayState.WithLabelValues(config.PrinterName, amsId, trayId).Set(safeFloat64(value))
				case "remain":
					metrics.AmsTrayRemain.WithLabelValues(config.PrinterName, amsId, trayId).Set(safeFloat64(value))
				case "k":
					metrics.AmsTrayK.WithLabelValues(config.PrinterName, amsId, trayId).Set(safeFloat64(value))
				case "n":
					metrics.AmsTrayN.WithLabelValues(config.PrinterName, amsId, trayId).Set(safeFloat64(value))
				case "cali_idx":
					metrics.AmsTrayCaliIdx.WithLabelValues(config.PrinterName, amsId, trayId).Set(safeFloat64(value))
				case "total_len":
					metrics.AmsTrayTotalLen.WithLabelValues(config.PrinterName, amsId, trayId).Set(safeFloat64(value))
				case "tray_diameter":
					metrics.AmsTrayDiameter.WithLabelValues(config.PrinterName, amsId, trayId).Set(safeFloat64(value))
				case "tray_temp":
					metrics.AmsTrayTemp.WithLabelValues(config.PrinterName, amsId, trayId).Set(safeFloat64(value))
				case "tray_time":
					metrics.AmsTrayTime.WithLabelValues(config.PrinterName, amsId, trayId).Set(safeFloat64(value))
				case "bed_temp_type":
					metrics.AmsTrayBedTempType.WithLabelValues(config.PrinterName, amsId, trayId).Set(safeFloat64(value))
				case "bed_temp":
					metrics.AmsTrayBedTemp.WithLabelValues(config.PrinterName, amsId, trayId).Set(safeFloat64(value))
				case "nozzle_temp_max":
					metrics.AmsTrayNozzleTempMax.WithLabelValues(config.PrinterName, amsId, trayId).Set(safeFloat64(value))
				case "nozzle_temp_min":
					metrics.AmsTrayNozzleTempMin.WithLabelValues(config.PrinterName, amsId, trayId).Set(safeFloat64(value))
				case "ctype":
					metrics.AmsTrayCtype.WithLabelValues(config.PrinterName, amsId, trayId).Set(safeFloat64(value))
				// Known fields that don't need metrics
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
					if isNewUnknownField("ams_"+amsId+"_tray_"+trayId, key) {
						log.Warn().Str("ams", amsId).Str("tray", trayId).Str("key", key).Interface("value", value).Msg("Unknown field in AMS tray report")
					}
				}
			}
		}
	}
}
