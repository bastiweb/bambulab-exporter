package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var UploadProgress = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "upload_progress_percent",
	Help: "Upload progress as a percentage",
}, []string{"printer"}))

var UploadStatus = makeCollector(prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: metricPrefix + "upload_status",
	Help: "Upload status (0 = idle, 1 = uploading, 2 = completed, 3 = error)",
}, []string{"printer"}))
