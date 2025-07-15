package api

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Start() <-chan error {
	http.Handle("/metrics", promhttp.Handler())

	errChan := make(chan error)
	go func() {
		errChan <- http.ListenAndServe(":8080", nil)
	}()
	return errChan
}
