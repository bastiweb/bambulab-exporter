package mqtt

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"time"

	"github.com/bastiweb/bambulab-exporter/config"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/rs/zerolog/log"
)

var errChan chan error = make(chan error, 1)

func Start(reportReceiver func(report map[string]any)) <-chan error {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("ssl://%s:%d", config.PrinterHost, 8883))
	opts.SetClientID("bambulab-exporter")
	opts.SetUsername("bblp")
	opts.SetPassword(config.PrinterAccessCode)
	opts.SetDefaultPublishHandler(onMessageDefault)
	opts.SetTLSConfig(&tls.Config{InsecureSkipVerify: true})
	opts.SetAutoReconnect(true)
	opts.SetConnectRetry(true)
	opts.SetConnectRetryInterval(1 * time.Second)
	opts.OnConnect = onConnect
	opts.OnConnectionLost = onConnectionLost

	client := mqtt.NewClient(opts)
	tok := client.Connect()
	tok.Wait()
	if err := tok.Error(); err != nil {
		errChan <- err
		return errChan
	}

	token := client.Subscribe(fmt.Sprintf("device/%s/report", config.PrinterSerial), 0, onMessage(reportReceiver))
	go func() {
		<-token.Done()
		if err := token.Error(); err != nil {
			errChan <- err
		}
	}()

	return errChan
}

func onConnect(client mqtt.Client) {
	log.Info().Msg("Connected to MQTT broker")
}

func onConnectionLost(client mqtt.Client, err error) {
	log.Error().Err(err).Msg("Connection lost")
	errChan <- err
}

func onMessageDefault(client mqtt.Client, msg mqtt.Message) {
	log.Debug().Str("payload", string(msg.Payload())).Msg("Received unexpected message")
}

func onMessage(reportReceiver func(report map[string]any)) func(client mqtt.Client, msg mqtt.Message) {
	return func(client mqtt.Client, msg mqtt.Message) {
		payload := msg.Payload()
		log.Trace().Str("payload", string(payload)).Msg("Received message")

		var report map[string]any
		err := json.Unmarshal(payload, &report)
		if err != nil {
			log.Error().Err(err).Msg("Failed to unmarshal message")
			return
		}

		reportReceiver(report)
	}
}
