package api

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Start() <-chan error {
	http.Handle("/metrics", promhttp.Handler())

	errChan := make(chan error)
	go func() {
		errChan <- http.ListenAndServe(":9043", nil)
	}()
	return errChan
}
