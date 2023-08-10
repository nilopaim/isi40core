package isi40core

import (
	"io"
	"log"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Protocol struct {
	Name    string
	Title   string
	Version string

	Log *log.Logger
}

func NewProtocol(protocolName string, protocolTitle, protocolVersion string) *Protocol {

	protocol := new(Protocol)

	protocol.Log = log.New(io.MultiWriter(os.Stdout), "["+protocolName+"] ", log.Ldate|log.Ltime|log.Lshortfile|log.Lmsgprefix)

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

	p.LogInfo("Protocol " + p.Name + " is running!!!!!")

	p.LogInfo("AQUI VOU MANDAR O AVISO DE QUE O PROTOCOLO ESTÁ RODANDO PARA O MQTT INTERNO")

	mqttClient := mqtt.NewClient(mqtt.NewClientOptions().AddBroker("tcp://localhost:1883"))
	mqttClient.Publish("isi40/protocol/"+p.Name+"/status", 0, false, "running")
}

func (p *Protocol) SendMessageMQTT(topic string, message string) {

	mqttClient := mqtt.NewClient(mqtt.NewClientOptions().AddBroker("tcp://localhost:1883"))
	mqttClient.Publish(topic, 0, false, message)

}

func (p *Protocol) LogInfo(message string) {

	p.Log.Printf("[%s] INFO - %v", p.Name, message)

}

func (p *Protocol) LogWarning(message string) {

	p.Log.Printf("[%s] WARN - %v", p.Name, message)

}

func (p *Protocol) LogError(message string) {

	p.Log.Printf("[%s] ERROR - %v", p.Name, message)

}
