package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// Sends a message to a local MQTT topic
func (p *Protocol) SendMessageMQTT(topic string, message any) {

	fmt.Println(reflect.TypeOf(message))

	if reflect.TypeOf(message).String() != "string" {
		x2, _ := json.Marshal(message)
		message = string(x2)
	}

	// Checks if MQTT client is connected
	if p.MqttClient.IsConnected() {

		// Publishes message to MQTT topic
		token := p.MqttClient.Publish(topic, 0, false, message)

		token.Wait()

		if token.Error() != nil {
			p.InternalLogError(fmt.Sprintf("cannot publish to MQTT: %v", token.Error()))
			p.InternalLogInfo("Sending message to MQTT broker: " + message.(string) + " (" + topic + ")")
		}
	} else {
		p.InternalLogError("cannot publish to MQTT broker: not connected")
	}

}
