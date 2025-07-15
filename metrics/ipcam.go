package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var IpcamDev = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "ipcam_dev",
	Help: "IP camera device status (1 = enabled, 0 = disabled)",
}, []string{"printer"}))

var IpcamRecord = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "ipcam_record",
	Help: "IP camera record status (1 = enabled, 0 = disabled)",
}, []string{"printer"}))

var IpcamTimelapse = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "ipcam_timelapse",
	Help: "IP camera timelapse status (1 = enabled, 0 = disabled)",
}, []string{"printer"}))

var IpcamTutkServer = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "ipcam_tutk_server",
	Help: "IP camera TUTK server status (1 = enabled, 0 = disabled)",
}, []string{"printer"}))

var IpcamModeBits = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "ipcam_mode_bits",
	Help: "IP camera mode bits",
}, []string{"printer"}))
