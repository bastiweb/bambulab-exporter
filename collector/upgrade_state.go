package collector

import (
	"github.com/Scrin/bambulab-exporter/config"
	"github.com/Scrin/bambulab-exporter/metrics"
	"github.com/rs/zerolog/log"
)

func receiveUpgradeStateReport(report map[string]any) {
	for key, value := range report {
		switch key {
		case "sequence_id":
			metrics.UpgradeStateSequenceId.WithLabelValues(config.PrinterName).Set(safeFloat64(value))
		case "progress":
			metrics.UpgradeStateProgress.WithLabelValues(config.PrinterName).Set(safeFloat64(value))
		case "status":
			// Convert string status to numeric
			statusStr := safeString(value)
			var statusValue float64
			switch statusStr {
			case "IDLE":
				statusValue = 0
			case "UPGRADING":
				statusValue = 1
			case "COMPLETED":
				statusValue = 2
			case "ERROR":
				statusValue = 3
			default:
				statusValue = 0
			}
			metrics.UpgradeStateStatus.WithLabelValues(config.PrinterName).Set(statusValue)
		case "consistency_request":
			if safeBool(value) {
				metrics.UpgradeStateConsistencyRequest.WithLabelValues(config.PrinterName).Set(1)
			} else {
				metrics.UpgradeStateConsistencyRequest.WithLabelValues(config.PrinterName).Set(0)
			}
		case "dis_state":
			metrics.UpgradeStateDisState.WithLabelValues(config.PrinterName).Set(safeFloat64(value))
		case "err_code":
			metrics.UpgradeStateErrCode.WithLabelValues(config.PrinterName).Set(safeFloat64(value))
		case "force_upgrade":
			if safeBool(value) {
				metrics.UpgradeStateForceUpgrade.WithLabelValues(config.PrinterName).Set(1)
			} else {
				metrics.UpgradeStateForceUpgrade.WithLabelValues(config.PrinterName).Set(0)
			}
		case "module":
			metrics.UpgradeStateModule.WithLabelValues(config.PrinterName).Set(safeFloat64(value))
		case "new_version_state":
			metrics.UpgradeStateNewVersionState.WithLabelValues(config.PrinterName).Set(safeFloat64(value))
		case "cur_state_code":
			metrics.UpgradeStateCurStateCode.WithLabelValues(config.PrinterName).Set(safeFloat64(value))
		case "idx2":
			metrics.UpgradeStateIdx2.WithLabelValues(config.PrinterName).Set(safeFloat64(value))
		case "message":
		case "new_ver_list":
		default:
			if isNewUnknownField("upgrade_state", key) {
				log.Warn().Str("key", key).Interface("value", value).Msg("Unknown field in upgrade_state report")
			}
		}
	}
}
