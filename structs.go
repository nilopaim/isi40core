package main

import (
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// Global struct of the protocol
type Protocol struct {
	Name    string
	Title   string
	Version string

	Log *log.Logger

	MqttClient mqtt.Client
}

// Status message to be sent to MQTT broker
type MessageStatusProtocol struct {
	MessageType string            `json:"messageType"`
	Data        MessageDataStatus `json:"data"`
}

type MessageDataStatus struct {
	Status string `json:"status"`
	Name   string `json:"name"`
}
