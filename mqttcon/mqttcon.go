package mqttcon

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"rizkyrajitha.github.io/mosquitto_mqtt_exporter/collector"
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
		collector.UptimeGauge(data)
	}

	//version
	// The version of the broker. Static.

	// if topic == "$SYS/broker/version" {
	// 	uptime := (strings.Split(msg, " "))[0]
	// 	data, _ := strconv.ParseFloat(uptime, 64)
	// 	// log.Println(data)
	// 	collector.UptimeGauge(data)
	// }
	
	//messages

	if topic == "$SYS/broker/messages/received" {
		data, _ := strconv.ParseFloat(msg, 64)
		collector.MessagesReceivedGauge(data)
	}

	if topic == "$SYS/broker/messages/sent" {
		data, _ := strconv.ParseFloat(msg, 64)
		// log.Println(data)
		collector.MessagesSentGauge(data)
	}

	if topic == "$SYS/broker/messages/inflight" {
		data, _ := strconv.ParseFloat(msg, 64)
		// log.Println(data)
		collector.MessagesInflightGauge(data)
	}

	// clients

	if topic == "$SYS/broker/clients/total" {
		data, _ := strconv.ParseFloat(msg, 64)
		// log.Println(data)
		collector.ClientsTotalGauge(data)
	}

	if topic == "$SYS/broker/clients/maximum" {
		data, _ := strconv.ParseFloat(msg, 64)
		// log.Println(data)
		collector.ClientsMaximumGauge(data)
	}

	if topic == "$SYS/broker/clients/connected" {
		data, _ := strconv.ParseFloat(msg, 64)
		// log.Println(data)
		collector.ClientsConnectedGauge(data)
	}

	if topic == "$SYS/broker/clients/disconnected" {
		data, _ := strconv.ParseFloat(msg, 64)
		// log.Println(data)
		collector.ClientsDisconnectedGauge(data)
	}

	if topic == "$SYS/broker/clients/expired" {
		data, _ := strconv.ParseFloat(msg, 64)
		// log.Println(data)
		collector.ClientsExpiredGauge(data)
	}

	//bytes

	if topic == "$SYS/broker/bytes/received" {
		data, _ := strconv.ParseFloat(msg, 64)
		// log.Println(data)
		collector.BytesReceivedGauge(data)
	}

	if topic == "$SYS/broker/bytes/sent" {
		data, _ := strconv.ParseFloat(msg, 64)
		// log.Println(data)
		collector.BytesSentGauge(data)
	}

	//heap

	if topic == "$SYS/broker/heap/current" {
		data, _ := strconv.ParseFloat(msg, 64)
		// log.Println(data)
		collector.HeapCurrentGauge(data)
	}

	if topic == "$SYS/broker/heap/maximum" {
		data, _ := strconv.ParseFloat(msg, 64)
		// log.Println(data)
		collector.HeapMaximumGauge(data)
	}

}
