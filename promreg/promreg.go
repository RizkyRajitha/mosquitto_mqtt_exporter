package promreg

import (
	"fmt"

	prometheus "github.com/prometheus/client_golang/prometheus"
)

var namespace = "mosquitto_mqtt_prometheus_exporter"

var mosquittoMessagesReceivedGauge= prometheus.NewGauge(prometheus.GaugeOpts{Name: "mosquitto_messages_received" , Help: "mosquitto broker messages received" , Namespace: namespace})
var mosquittoMessagesSentGauge= prometheus.NewGauge(prometheus.GaugeOpts{Name: "mosquitto_messages_sent" , Help: "mosquitto broker messages sent" , Namespace: namespace})
var mosquittoUptimeGauge = prometheus.NewGauge(prometheus.GaugeOpts{Name: "mosquitto_uptime" , Help: "mosquitto broker uptime" , Namespace: namespace})

func init()  {
	prometheus.MustRegister(mosquittoUptimeGauge)
	prometheus.MustRegister(mosquittoMessagesSentGauge)
	prometheus.MustRegister(mosquittoMessagesReceivedGauge)
}

func MessagesReceivedGauge(MessagesReceivedCount float64)  {
	fmt.Println("MessagesReceivedCount")
    mosquittoMessagesReceivedGauge.Set(MessagesReceivedCount)
}

func MessagesSentGauge(MessagesSentCount float64)  {
	fmt.Println("MessagesSentCount")
    mosquittoMessagesSentGauge.Set(MessagesSentCount)
}

func UptimeGauge(Uptime float64)  {
	fmt.Println("Uptime")
    mosquittoUptimeGauge.Set(Uptime)
}


// func PromRegister()  {
// 	fmt.Println("go go cargo")
// 	i = i+10
//     gauge.Set(float64(i))
// }

// func PromRegister()  {
// 	fmt.Println("go go cargo")
// 	i = i+10
//     gauge.Set(float64(i))
// }

// func PromRegister()  {
// 	fmt.Println("go go cargo")
// 	i = i+10
//     gauge.Set(float64(i))
// }