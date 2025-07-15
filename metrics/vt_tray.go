package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var VtTrayId = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "vt_tray_id",
	Help: "VT tray ID",
}, []string{"printer"}))

var VtTrayRemain = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "vt_tray_remain",
	Help: "VT tray remaining material",
}, []string{"printer"}))

var VtTrayK = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "vt_tray_k",
	Help: "VT tray K value",
}, []string{"printer"}))

var VtTrayN = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "vt_tray_n",
	Help: "VT tray N value",
}, []string{"printer"}))

var VtTrayCaliIdx = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "vt_tray_cali_idx",
	Help: "VT tray calibration index",
}, []string{"printer"}))

var VtTrayDiameter = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "vt_tray_diameter_mm",
	Help: "VT tray diameter in millimeters",
}, []string{"printer"}))

var VtTrayTemp = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "vt_tray_temp_celsius",
	Help: "VT tray temperature in Celsius",
}, []string{"printer"}))

var VtTrayTime = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "vt_tray_time_hours",
	Help: "VT tray time in hours",
}, []string{"printer"}))

var VtTrayBedTempType = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "vt_tray_bed_temp_type",
	Help: "VT tray bed temperature type",
}, []string{"printer"}))

var VtTrayBedTemp = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "vt_tray_bed_temp_celsius",
	Help: "VT tray bed temperature in Celsius",
}, []string{"printer"}))

var VtTrayNozzleTempMax = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "vt_tray_nozzle_temp_max_celsius",
	Help: "VT tray maximum nozzle temperature in Celsius",
}, []string{"printer"}))

var VtTrayNozzleTempMin = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "vt_tray_nozzle_temp_min_celsius",
	Help: "VT tray minimum nozzle temperature in Celsius",
}, []string{"printer"}))

var VtTrayCtype = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "vt_tray_ctype",
	Help: "VT tray ctype",
}, []string{"printer"}))
