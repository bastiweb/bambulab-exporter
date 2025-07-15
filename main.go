package main

import (
	"github.com/Scrin/bambulab-exporter/api"
	"github.com/Scrin/bambulab-exporter/collector"
	"github.com/Scrin/bambulab-exporter/config"
	"github.com/Scrin/bambulab-exporter/logging"
	"github.com/Scrin/bambulab-exporter/mqtt"
	"github.com/rs/zerolog/log"
)

func main() {
	logging.Setup()
	err := config.LoadEnv()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load config")
	}

	apiErrChan := api.Start()
	mqttErrChan := mqtt.Start(collector.ReceiveReport)

	log.Info().Msg("Bambulab exporter started")

	select {
	case err := <-apiErrChan:
		log.Fatal().Err(err).Msg("API error")
	case err := <-mqttErrChan:
		log.Fatal().Err(err).Msg("MQTT error")
	}
}
