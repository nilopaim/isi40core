package isi40core

import (
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Protocol struct {
	Name    string
	Title   string
	Version string
}

func NewProtocol(protocolName string, protocolTitle, protocolVersion string) *Protocol {
	protocol := new(Protocol)
	protocol.Name = protocolName
	protocol.Title = protocolTitle
	protocol.Version = protocolVersion
	return protocol
}

func (p *Protocol) GetProtocolName() string {
	return p.Name
}

func (p *Protocol) GetProtocolTitle() string {
	return p.Title
}

func (p *Protocol) GetProtocolVersion() string {
	return p.Version
}

func (p *Protocol) Init() {
	LogInfo("Protocol " + p.Name + " is running!!!!!")
	LogInfo("AQUI VOU MANDAR O AVISO DE QUE O PROTOCOLO ESTÁ RODANDO PARA O MQTT INTERNO")
	mqttClient := mqtt.NewClient(mqtt.NewClientOptions().AddBroker("tcp://localhost:1883"))
	mqttClient.Publish("isi40/protocol/"+p.Name+"/status", 0, false, "running")
}

func (p *Protocol) SendMessageMQTT(topic string, message string) {
	mqttClient := mqtt.NewClient(mqtt.NewClientOptions().AddBroker("tcp://localhost:1883"))
	mqttClient.Publish(topic, 0, false, message)

}

func LogInfo(message string) {
	log.Printf("INFO - %v", message)
}

func LogWarning(message string) {
	log.Printf("WARN - %v", message)
}

func LogError(message string) {
	log.Printf("ERROR - %v", message)
}
