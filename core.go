package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// Create a new protocol
func NewProtocol(protocolName, protocolTitle, protocolVersion string) *Protocol {

	protocol := new(Protocol)

	// Create nd configures log file
	LogFileHandle, err := os.OpenFile("./"+protocolName+".log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	protocol.Log = log.New(io.MultiWriter(os.Stdout, LogFileHandle), "["+protocolName+"] ", log.Ldate|log.Ltime|log.Lmsgprefix)

	// Set protocol name, title and version
	protocol.Name = protocolName
	protocol.Title = protocolTitle
	protocol.Version = protocolVersion

	return protocol
}

// Returns protocol name
func (p *Protocol) GetProtocolName() string {
	return p.Name
}

// Returns protocol title
func (p *Protocol) GetProtocolTitle() string {
	return p.Title
}

// Returns protocol version
func (p *Protocol) GetProtocolVersion() string {
	return p.Version
}

// Initializes protocol
func (p *Protocol) Init() {

	// Connects to MQTT broker
	p.MqttClient = mqtt.NewClient(mqtt.NewClientOptions().AddBroker("tcp://localhost:1883"))

	if token := p.MqttClient.Connect(); token.Wait() && token.Error() != nil {
		p.InternalLogError(fmt.Sprintf("MQTT: [%s]", token.Error()))
		p.InternalLogInfo("Protocol " + p.Name + " cannot be started due to MQTT connection error.")
		os.Exit(1)
	}

	// Publishes status message indicating that the protocol is running
	var data MessageStatusProtocol
	data.MessageType = "status"
	data.Data.Status = "running"
	data.Data.Name = p.Name
	payload, _ := json.Marshal(data)

	p.MqttClient.Publish("general", 0, false, payload)

	p.InternalLogInfo("Protocol " + p.Name + " is running!")
}
