package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var NozzleTemperature = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "nozzle_temperature",
	Help: "Current temperature of the nozzle",
}, []string{"printer"}))

var NozzleTargetTemperature = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "nozzle_target_temperature",
	Help: "Target temperature of the nozzle",
}, []string{"printer"}))

var BedTemperature = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "bed_temperature",
	Help: "Current temperature of the bed",
}, []string{"printer"}))

var BedTargetTemperature = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "bed_target_temperature",
	Help: "Target temperature of the bed",
}, []string{"printer"}))

var ChamberTemperature = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "chamber_temperature",
	Help: "Current temperature of the chamber",
}, []string{"printer"}))

var HeatbreakFanSpeed = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "heatbreak_fan_speed",
	Help: "Speed of the heatbreak fan",
}, []string{"printer"}))

var CoolingFanSpeed = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "cooling_fan_speed",
	Help: "Speed of the cooling fan",
}, []string{"printer"}))

var BigFan1Speed = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "big_fan1_speed",
	Help: "Speed of the first big fan",
}, []string{"printer"}))

var BigFan2Speed = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "big_fan2_speed",
	Help: "Speed of the second big fan",
}, []string{"printer"}))

var PrintProgress = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "print_progress_percent",
	Help: "Print progress as a percentage",
}, []string{"printer"}))

var PrintRemainingTime = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "print_remaining_time_minutes",
	Help: "Remaining print time in minutes",
}, []string{"printer"}))

var PrintSpeedMagnitude = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "print_speed_magnitude",
	Help: "Print speed magnitude",
}, []string{"printer"}))

var PrintSpeedLevel = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "print_speed_level",
	Help: "Print speed level",
}, []string{"printer"}))

var PrintError = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "print_error",
	Help: "Print error code (0 = no error)",
}, []string{"printer"}))

var WifiSignal = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "wifi_signal_dbm",
	Help: "WiFi signal strength in dBm",
}, []string{"printer"}))

var LayerNumber = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "current_layer",
	Help: "Current layer number being printed",
}, []string{"printer"}))

var TotalLayers = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "total_layers",
	Help: "Total number of layers in the print",
}, []string{"printer"}))

var AmsStatus = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "ams_status",
	Help: "AMS (Automatic Material System) status",
}, []string{"printer"}))

var AmsRfidStatus = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "ams_rfid_status",
	Help: "AMS RFID status",
}, []string{"printer"}))

var HardwareSwitchState = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "hardware_switch_state",
	Help: "Hardware switch state",
}, []string{"printer"}))

var Stage = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "stage",
	Help: "Current stage",
}, []string{"printer"}))

var PrintStage = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "print_stage",
	Help: "Current print stage",
}, []string{"printer"}))

var PrintSubStage = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "print_sub_stage",
	Help: "Current print sub-stage",
}, []string{"printer"}))

var GcodeFilePreparePercent = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "gcode_file_prepare_percent",
	Help: "G-code file preparation progress percentage",
}, []string{"printer"}))

var QueueNumber = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "queue_number",
	Help: "Current queue number",
}, []string{"printer"}))

var QueueTotal = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "queue_total",
	Help: "Total items in queue",
}, []string{"printer"}))

var QueueEstimatedTime = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "queue_estimated_time_minutes",
	Help: "Estimated queue time in minutes",
}, []string{"printer"}))

var QueueStatus = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "queue_status",
	Help: "Queue status",
}, []string{"printer"}))

var McPrintLineNumber = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "mc_print_line_number",
	Help: "Current G-code line number being executed",
}, []string{"printer"}))

var NozzleDiameter = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "nozzle_diameter_mm",
	Help: "Nozzle diameter in millimeters",
}, []string{"printer"}))

var CalibrationK = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "calibration_k",
	Help: "Calibration K value",
}, []string{"printer"}))

var FanGear = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "fan_gear",
	Help: "Fan gear setting",
}, []string{"printer"}))

var HomeFlag = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "home_flag",
	Help: "Home flag status",
}, []string{"printer"}))

var Flag3 = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "flag3",
	Help: "Flag3 status",
}, []string{"printer"}))

var CalibrationVersion = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "calibration_version",
	Help: "Calibration version",
}, []string{"printer"}))

var SdcardPresent = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "sdcard_present",
	Help: "SD card presence status (1 = present, 0 = not present)",
}, []string{"printer"}))

var ForceUpgrade = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "force_upgrade",
	Help: "Force upgrade status (1 = enabled, 0 = disabled)",
}, []string{"printer"}))

var WifiSignalRaw = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "wifi_signal_raw",
	Help: "Raw WiFi signal value",
}, []string{"printer"}))

var GcodeState = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "gcode_state",
	Help: "G-code state (0 = idle, 1 = running, 2 = paused, 3 = completed, 4 = error)",
}, []string{"printer"}))

var PrintType = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "print_type",
	Help: "Print type (0 = local, 1 = cloud)",
}, []string{"printer"}))

var MessProductionState = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "mess_production_state",
	Help: "Mess production state (0 = inactive, 1 = active)",
}, []string{"printer"}))

var SequenceId = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "sequence_id",
	Help: "Sequence ID",
}, []string{"printer"}))

var Msg = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "msg",
	Help: "Message value",
}, []string{"printer"}))
