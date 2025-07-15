package collector

import (
	"github.com/Scrin/bambulab-exporter/config"
	"github.com/Scrin/bambulab-exporter/metrics"
	"github.com/rs/zerolog/log"
)

func receivePrintReport(report map[string]any) {
	for key, value := range report {
		switch key {
		case "nozzle_temper":
			metrics.NozzleTemperature.WithLabelValues(config.PrinterName).Set(safeFloat64(value))
		case "nozzle_target_temper":
			metrics.NozzleTargetTemperature.WithLabelValues(config.PrinterName).Set(safeFloat64(value))
		case "bed_temper":
			metrics.BedTemperature.WithLabelValues(config.PrinterName).Set(safeFloat64(value))
		case "bed_target_temper":
			metrics.BedTargetTemperature.WithLabelValues(config.PrinterName).Set(safeFloat64(value))
		case "chamber_temper":
			metrics.ChamberTemperature.WithLabelValues(config.PrinterName).Set(safeFloat64(value))
		case "heatbreak_fan_speed":
			metrics.HeatbreakFanSpeed.WithLabelValues(config.PrinterName).Set(safeFloat64(value))
		case "cooling_fan_speed":
			metrics.CoolingFanSpeed.WithLabelValues(config.PrinterName).Set(safeFloat64(value))
		case "big_fan1_speed":
			metrics.BigFan1Speed.WithLabelValues(config.PrinterName).Set(safeFloat64(value))
		case "big_fan2_speed":
			metrics.BigFan2Speed.WithLabelValues(config.PrinterName).Set(safeFloat64(value))
		case "mc_percent":
			metrics.PrintProgress.WithLabelValues(config.PrinterName).Set(safeFloat64(value))
		case "mc_remaining_time":
			metrics.PrintRemainingTime.WithLabelValues(config.PrinterName).Set(safeFloat64(value))
		case "spd_mag":
			metrics.PrintSpeedMagnitude.WithLabelValues(config.PrinterName).Set(safeFloat64(value))
		case "spd_lvl":
			metrics.PrintSpeedLevel.WithLabelValues(config.PrinterName).Set(safeFloat64(value))
		case "print_error":
			metrics.PrintError.WithLabelValues(config.PrinterName).Set(safeFloat64(value))
		case "wifi_signal":
			metrics.WifiSignal.WithLabelValues(config.PrinterName).Set(safeFloat64(value))
		case "layer_num":
			metrics.LayerNumber.WithLabelValues(config.PrinterName).Set(safeFloat64(value))
		case "total_layer_num":
			metrics.TotalLayers.WithLabelValues(config.PrinterName).Set(safeFloat64(value))
		case "ams_status":
			metrics.AmsStatus.WithLabelValues(config.PrinterName).Set(safeFloat64(value))
		case "ams_rfid_status":
			metrics.AmsRfidStatus.WithLabelValues(config.PrinterName).Set(safeFloat64(value))
		case "hw_switch_state":
			metrics.HardwareSwitchState.WithLabelValues(config.PrinterName).Set(safeFloat64(value))
		case "mc_print_stage":
			metrics.PrintStage.WithLabelValues(config.PrinterName).Set(safeFloat64(value))
		case "mc_print_sub_stage":
			metrics.PrintSubStage.WithLabelValues(config.PrinterName).Set(safeFloat64(value))
		case "gcode_file_prepare_percent":
			metrics.GcodeFilePreparePercent.WithLabelValues(config.PrinterName).Set(safeFloat64(value))
		case "queue_number":
			metrics.QueueNumber.WithLabelValues(config.PrinterName).Set(safeFloat64(value))
		case "queue_total":
			metrics.QueueTotal.WithLabelValues(config.PrinterName).Set(safeFloat64(value))
		case "queue_est":
			metrics.QueueEstimatedTime.WithLabelValues(config.PrinterName).Set(safeFloat64(value))
		case "queue_sts":
			metrics.QueueStatus.WithLabelValues(config.PrinterName).Set(safeFloat64(value))
		case "mc_print_line_number":
			metrics.McPrintLineNumber.WithLabelValues(config.PrinterName).Set(safeFloat64(value))
		case "nozzle_diameter":
			metrics.NozzleDiameter.WithLabelValues(config.PrinterName).Set(safeFloat64(value))
		case "k":
			metrics.CalibrationK.WithLabelValues(config.PrinterName).Set(safeFloat64(value))
		case "fan_gear":
			metrics.FanGear.WithLabelValues(config.PrinterName).Set(safeFloat64(value))
		case "home_flag":
			metrics.HomeFlag.WithLabelValues(config.PrinterName).Set(safeFloat64(value))
		case "flag3":
			metrics.Flag3.WithLabelValues(config.PrinterName).Set(safeFloat64(value))
		case "cali_version":
			metrics.CalibrationVersion.WithLabelValues(config.PrinterName).Set(safeFloat64(value))
		case "sdcard":
			if safeBool(value) {
				metrics.SdcardPresent.WithLabelValues(config.PrinterName).Set(1)
			} else {
				metrics.SdcardPresent.WithLabelValues(config.PrinterName).Set(0)
			}
		case "force_upgrade":
			if safeBool(value) {
				metrics.ForceUpgrade.WithLabelValues(config.PrinterName).Set(1)
			} else {
				metrics.ForceUpgrade.WithLabelValues(config.PrinterName).Set(0)
			}

		case "gcode_state":
			// Convert string state to numeric
			stateStr := safeString(value)
			var stateValue float64
			switch stateStr {
			case "IDLE":
				stateValue = 0
			case "RUNNING":
				stateValue = 1
			case "PAUSED":
				stateValue = 2
			case "COMPLETED":
				stateValue = 3
			case "ERROR":
				stateValue = 4
			default:
				stateValue = 0
			}
			metrics.GcodeState.WithLabelValues(config.PrinterName).Set(stateValue)
		case "print_type":
			// Convert string type to numeric
			typeStr := safeString(value)
			var typeValue float64
			switch typeStr {
			case "local":
				typeValue = 0
			case "cloud":
				typeValue = 1
			default:
				typeValue = 0
			}
			metrics.PrintType.WithLabelValues(config.PrinterName).Set(typeValue)
		case "mess_production_state":
			// Convert string state to numeric
			stateStr := safeString(value)
			var stateValue float64
			switch stateStr {
			case "inactive":
				stateValue = 0
			case "active":
				stateValue = 1
			default:
				stateValue = 0
			}
			metrics.MessProductionState.WithLabelValues(config.PrinterName).Set(stateValue)
		case "sequence_id":
			metrics.SequenceId.WithLabelValues(config.PrinterName).Set(safeFloat64(value))
		case "msg":
			metrics.Msg.WithLabelValues(config.PrinterName).Set(safeFloat64(value))
		case "ipcam":
			if ipcamReport, ok := value.(map[string]any); ok {
				receiveIpcamReport(ipcamReport)
			}
		case "upload":
			if uploadReport, ok := value.(map[string]any); ok {
				receiveUploadReport(uploadReport)
			}
		case "net":
			if netReport, ok := value.(map[string]any); ok {
				receiveNetReport(netReport)
			}
		case "ams":
			if amsReport, ok := value.(map[string]any); ok {
				receiveAmsReport(amsReport)
			}
		case "vt_tray":
			if vtTrayReport, ok := value.(map[string]any); ok {
				receiveVtTrayReport(vtTrayReport)
			}
		case "lights_report":
			if lightsReport, ok := value.([]any); ok {
				receiveLightsReport(lightsReport)
			}
		case "upgrade_state":
			if upgradeStateReport, ok := value.(map[string]any); ok {
				receiveUpgradeStateReport(upgradeStateReport)
			}
		case "online":
			if onlineReport, ok := value.(map[string]any); ok {
				receiveOnlineReport(onlineReport)
			}
		case "command":
		case "lifecycle":
		case "project_id":
		case "profile_id":
		case "task_id":
		case "subtask_id":
		case "subtask_name":
		case "gcode_file":
		case "stg":
		case "stg_cur":
		case "s_obj":
		case "filam_bak":
		case "nozzle_type":
		case "hms":
		default:
			if isNewUnknownField("print", key) {
				log.Warn().Str("key", key).Interface("value", value).Msg("Unknown field in print report")
			}
		}
	}
}
