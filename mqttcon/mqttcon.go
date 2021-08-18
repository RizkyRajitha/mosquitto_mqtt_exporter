package mqttcon

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"rizkyrajitha.github.io/mqttexporter/metrics"
)

var topic = "$SYS/broker/#"
var mqttClient mqtt.Client
var brokerAddress = ""

func MqttCon(brokerIp string, username string, password string) {

	brokerAddress = brokerIp

	mqttClientOptions := mqtt.NewClientOptions().AddBroker(brokerIp)

	if username != "" {
		mqttClientOptions.Username = username
	}

	if password != "" {
		mqttClientOptions.Password = password
	}

	mqttClientOptions.OnConnect = connectionHandler
	mqttClientOptions.DefaultPublishHandler = messageHandler
	mqttClient = mqtt.NewClient(mqttClientOptions)

	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal("Error connecting to mqtt broker\n", token.Error())
		// panic(token.Error())
	}

	if token := mqttClient.Subscribe(topic, 0, nil); token.Wait() && token.Error() != nil {
		logMsg := fmt.Sprintf("Error subscribing to topic : %s on  mqtt broker\n%s", topic, token.Error())
		log.Fatal(logMsg)
		// panic(token.Error())
	}

	log.Println("Subscribed to  : ", topic)

}

var connectionHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	log.Println("Connected to : ", brokerAddress)
}

var messageHandler mqtt.MessageHandler = func(client mqtt.Client, message mqtt.Message) {

	topic := message.Topic()
	msg := string(message.Payload())

	//uptime

	if topic == "$SYS/broker/uptime" {
		uptime := (strings.Split(msg, " "))[0]
		data, _ := strconv.ParseFloat(uptime, 64)
		// log.Println(data)
		metrics.UptimeGauge(data)
	}

	//messages

	if topic == "$SYS/broker/messages/received" {
		data, _ := strconv.ParseFloat(msg, 64)
		metrics.MessagesReceivedGauge(data)
	}

	if topic == "$SYS/broker/messages/sent" {
		data, _ := strconv.ParseFloat(msg, 64)
		// log.Println(data)
		metrics.MessagesSentGauge(data)
	}

	// clients

	if topic == "$SYS/broker/clients/total" {
		data, _ := strconv.ParseFloat(msg, 64)
		// log.Println(data)
		metrics.ClientsTotalGauge(data)
	}

	if topic == "$SYS/broker/clients/maximum" {
		data, _ := strconv.ParseFloat(msg, 64)
		// log.Println(data)
		metrics.ClientsMaximumGauge(data)
	}

	if topic == "$SYS/broker/clients/connected" {
		data, _ := strconv.ParseFloat(msg, 64)
		// log.Println(data)
		metrics.ClientsConnectedGauge(data)
	}

	//bytes

	if topic == "$SYS/broker/bytes/received" {
		data, _ := strconv.ParseFloat(msg, 64)
		// log.Println(data)
		metrics.BytesReceivedGauge(data)
	}

	if topic == "$SYS/broker/bytes/sent" {
		data, _ := strconv.ParseFloat(msg, 64)
		// log.Println(data)
		metrics.BytesSentGauge(data)
	}

	//heap

	if topic == "$SYS/broker/heap/current" {
		data, _ := strconv.ParseFloat(msg, 64)
		// log.Println(data)
		metrics.HeapCurrentGauge(data)
	}

	if topic == "$SYS/broker/heap/maximum" {
		data, _ := strconv.ParseFloat(msg, 64)
		// log.Println(data)
		metrics.HeapMaximumGauge(data)
	}

}
