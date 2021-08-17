package mqttcon

import (
	"fmt"
	"strconv"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"rizkyrajitha.github.io/mqttexporter/promreg"
)

var topic = "$SYS/broker/#"
var mqttClient mqtt.Client
var brokerAddress = ""

func MqttCon( brokerIp string , username string ,password string ){
	
	brokerAddress = brokerIp

	mqttClientOptions := mqtt.NewClientOptions().AddBroker(brokerIp)

	if(username !=""){
		mqttClientOptions.Username = username
	}

	if(password !=""){
		mqttClientOptions.Password = password
	}

	mqttClientOptions.OnConnect = connectionHandler
	mqttClientOptions.DefaultPublishHandler = messageHandler
	mqttClient = mqtt.NewClient(mqttClientOptions)
	
    if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {    
		panic(token.Error())
	}

    if token :=  mqttClient.Subscribe(topic , 0 , nil); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

  	fmt.Println("Subscribed to topic: ", topic)

}


var connectionHandler mqtt.OnConnectHandler = func (client mqtt.Client)  {
	fmt.Println("connected to ",brokerAddress)
}

var messageHandler mqtt.MessageHandler = func(c mqtt.Client, m mqtt.Message) {

	topic := m.Topic()
	msg := string(m.Payload())


	if(topic=="$SYS/broker/uptime"){
		fmt.Println(msg)
		uptime := (strings.Split(msg , " "))[0]
		data , _ :=  strconv.ParseFloat(uptime,64)
		fmt.Println(data)
		promreg.UptimeGauge(data)	
	}

	if(topic=="$SYS/broker/messages/received"){
		fmt.Println(msg)
		data , _ :=  strconv.ParseFloat(msg,64)
		promreg.MessagesReceivedGauge(data)
	}

	if(topic=="$SYS/broker/messages/sent"){
		fmt.Println(msg)
		data , _ :=  strconv.ParseFloat(msg,64)
		fmt.Println(data)
		promreg.MessagesSentGauge(data)
	}

}