package isi40core

import "log"

type Protocol struct {
	ProtocolName string
}

func (p *Protocol) GetProtocolName() string {
	return p.ProtocolName
}

func (p *Protocol) SetProtocolName(protocolName string) {
	p.ProtocolName = protocolName
}

func (p *Protocol) Init() {
	LogInfo("Initializing protocol " + p.ProtocolName)
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
