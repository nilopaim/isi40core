package main

func main() {

	// Creates the protocol
	p := NewProtocol("customprotocol", "Protocolo do Nilo", "1.0.0")

	// nome := p.GetProtocolName()
	// titulo := p.GetProtocolTitle()
	// versao := p.GetProtocolVersion()

	// msg := fmt.Sprintf("Protocolo [%s] (%s), versão %s", titulo, nome, versao)
	// p.LogInfo(msg)

	// Initializes the protocol
	p.Init()

	// Sends a message to MQTT broker
	p.SendMessageMQTT(p.GetProtocolName(), "Nilo")

	type msg struct {
		Msg string `json:"msg"`
	}
	x := msg{Msg: "Nilo, mas como json"}
	p.SendMessageMQTT(p.GetProtocolName(), x)
}
