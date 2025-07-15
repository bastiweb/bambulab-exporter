package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var AmsExistBits = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "ams_exist_bits",
	Help: "AMS existence bits",
}, []string{"printer"}))

var TrayExistBits = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "tray_exist_bits",
	Help: "Tray existence bits",
}, []string{"printer"}))

var TrayIsBblBits = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "tray_is_bbl_bits",
	Help: "Tray is BBL bits",
}, []string{"printer"}))

var TrayTarget = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "tray_target",
	Help: "Target tray number",
}, []string{"printer"}))

var TrayCurrent = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "tray_current",
	Help: "Current tray number",
}, []string{"printer"}))

var TrayPrevious = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "tray_previous",
	Help: "Previous tray number",
}, []string{"printer"}))

var TrayReadDoneBits = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "tray_read_done_bits",
	Help: "Tray read done bits",
}, []string{"printer"}))

var TrayReadingBits = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "tray_reading_bits",
	Help: "Tray reading bits",
}, []string{"printer"}))

var AmsVersion = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "ams_version",
	Help: "AMS version",
}, []string{"printer"}))

var AmsInsertFlag = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "ams_insert_flag",
	Help: "AMS insert flag (1 = inserted, 0 = not inserted)",
}, []string{"printer"}))

var AmsPowerOnFlag = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "ams_power_on_flag",
	Help: "AMS power on flag (1 = powered on, 0 = powered off)",
}, []string{"printer"}))

var AmsHumidity = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "ams_humidity",
	Help: "AMS humidity",
}, []string{"printer", "ams"}))

var AmsHumidityRaw = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "ams_humidity_raw",
	Help: "AMS raw humidity value",
}, []string{"printer", "ams"}))

var AmsTemp = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "ams_temp_celsius",
	Help: "AMS temperature in Celsius",
}, []string{"printer", "ams"}))

var AmsDryTime = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "ams_dry_time_hours",
	Help: "AMS dry time in hours",
}, []string{"printer", "ams"}))

var AmsInfo = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "ams_info",
	Help: "AMS info value",
}, []string{"printer", "ams"}))

var AmsTrayId = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "ams_tray_id",
	Help: "AMS tray ID",
}, []string{"printer", "ams", "tray"}))

var AmsTrayState = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "ams_tray_state",
	Help: "AMS tray state",
}, []string{"printer", "ams", "tray"}))

var AmsTrayRemain = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "ams_tray_remain",
	Help: "AMS tray remaining material",
}, []string{"printer", "ams", "tray"}))

var AmsTrayK = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "ams_tray_k",
	Help: "AMS tray K value",
}, []string{"printer", "ams", "tray"}))

var AmsTrayN = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "ams_tray_n",
	Help: "AMS tray N value",
}, []string{"printer", "ams", "tray"}))

var AmsTrayCaliIdx = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "ams_tray_cali_idx",
	Help: "AMS tray calibration index",
}, []string{"printer", "ams", "tray"}))

var AmsTrayTotalLen = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "ams_tray_total_len",
	Help: "AMS tray total length",
}, []string{"printer", "ams", "tray"}))

var AmsTrayDiameter = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "ams_tray_diameter_mm",
	Help: "AMS tray diameter in millimeters",
}, []string{"printer", "ams", "tray"}))

var AmsTrayTemp = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "ams_tray_temp_celsius",
	Help: "AMS tray temperature in Celsius",
}, []string{"printer", "ams", "tray"}))

var AmsTrayTime = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "ams_tray_time_hours",
	Help: "AMS tray time in hours",
}, []string{"printer", "ams", "tray"}))

var AmsTrayBedTempType = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "ams_tray_bed_temp_type",
	Help: "AMS tray bed temperature type",
}, []string{"printer", "ams", "tray"}))

var AmsTrayBedTemp = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "ams_tray_bed_temp_celsius",
	Help: "AMS tray bed temperature in Celsius",
}, []string{"printer", "ams", "tray"}))

var AmsTrayNozzleTempMax = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "ams_tray_nozzle_temp_max_celsius",
	Help: "AMS tray maximum nozzle temperature in Celsius",
}, []string{"printer", "ams", "tray"}))

var AmsTrayNozzleTempMin = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "ams_tray_nozzle_temp_min_celsius",
	Help: "AMS tray minimum nozzle temperature in Celsius",
}, []string{"printer", "ams", "tray"}))

var AmsTrayCtype = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "ams_tray_ctype",
	Help: "AMS tray ctype",
}, []string{"printer", "ams", "tray"}))
