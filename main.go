package isi40core

import "log"

type Protocol struct {
	Name    string
	Version string
}

func NewProtocol(protocolName string, protocolVersion string) *Protocol {
	protocol := new(Protocol)
	protocol.Name = protocolName
	protocol.Version = protocolVersion
	return protocol
}

func (p *Protocol) GetProtocolName() string {
	return p.Name
}

func (p *Protocol) GetProtocolVersion() string {
	return p.Version
}

func (p *Protocol) Init() {
	LogInfo("Initializing protocol " + p.Name)
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
