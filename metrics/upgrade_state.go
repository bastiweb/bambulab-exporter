package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var UpgradeStateSequenceId = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "upgrade_state_sequence_id",
	Help: "Upgrade state sequence ID",
}, []string{"printer"}))

var UpgradeStateProgress = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "upgrade_state_progress_percent",
	Help: "Upgrade state progress as a percentage",
}, []string{"printer"}))

var UpgradeStateStatus = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "upgrade_state_status",
	Help: "Upgrade state status (0 = idle, 1 = upgrading, 2 = completed, 3 = error)",
}, []string{"printer"}))

var UpgradeStateConsistencyRequest = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "upgrade_state_consistency_request",
	Help: "Upgrade state consistency request (1 = requested, 0 = not requested)",
}, []string{"printer"}))

var UpgradeStateDisState = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "upgrade_state_dis_state",
	Help: "Upgrade state dis state",
}, []string{"printer"}))

var UpgradeStateErrCode = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "upgrade_state_err_code",
	Help: "Upgrade state error code",
}, []string{"printer"}))

var UpgradeStateForceUpgrade = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "upgrade_state_force_upgrade",
	Help: "Upgrade state force upgrade (1 = enabled, 0 = disabled)",
}, []string{"printer"}))

var UpgradeStateModule = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "upgrade_state_module",
	Help: "Upgrade state module",
}, []string{"printer"}))

var UpgradeStateNewVersionState = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "upgrade_state_new_version_state",
	Help: "Upgrade state new version state",
}, []string{"printer"}))

var UpgradeStateCurStateCode = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "upgrade_state_cur_state_code",
	Help: "Upgrade state current state code",
}, []string{"printer"}))

var UpgradeStateIdx2 = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "upgrade_state_idx2",
	Help: "Upgrade state idx2",
}, []string{"printer"}))
