package isi40core

import (
	"fmt"
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

func NewProtocol(protocolName, protocolTitle, protocolVersion string) *Protocol {

	protocol := new(Protocol)

	LogFileHandle, err := os.OpenFile("./"+protocolName+".log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}

	protocol.Log = log.New(io.MultiWriter(os.Stdout, LogFileHandle), "["+protocolName+"] ", log.Ldate|log.Ltime|log.Lmsgprefix)

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

	p.InternalLogInfo("Protocol " + p.Name + " is running!!!!!")

	p.InternalLogInfo("AQUI VOU MANDAR O AVISO DE QUE O PROTOCOLO ESTÁ RODANDO PARA O MQTT INTERNO")

	mqttClient := mqtt.NewClient(mqtt.NewClientOptions().AddBroker("tcp://localhost:1883"))

	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		p.InternalLogError(fmt.Sprintf("MQTT Internal Broker is not running. Error %s", token.Error()))
		os.Exit(1)
	}

	mqttClient.Publish("isi40/protocol/"+p.Name+"/status", 0, false, "running")
}

func (p *Protocol) SendMessageMQTT(topic string, message string) {

	mqttClient := mqtt.NewClient(mqtt.NewClientOptions().AddBroker("tcp://localhost:1883"))
	mqttClient.Publish(topic, 0, false, message)

}
