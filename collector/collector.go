package collector

import (
	// "log"

	prometheus "github.com/prometheus/client_golang/prometheus"
)

var namespace = "mosquitto_mqtt"

var mqttRegistry = prometheus.NewRegistry()

//uptime
var mosquittoUptimeGauge = prometheus.NewGauge(prometheus.GaugeOpts{Name: "mosquitto_uptime", Help: "The Mosquitto broker uptime", Namespace: namespace})

//messages
var mosquittoMessagesSentGauge = prometheus.NewGauge(prometheus.GaugeOpts{Name: "mosquitto_messages_sent", Help: "The total number of messages of any type sent since the broker started.", Namespace: namespace})
var mosquittoMessagesReceivedGauge = prometheus.NewGauge(prometheus.GaugeOpts{Name: "mosquitto_messages_received", Help: "The total number of messages of any type received since the broker started.", Namespace: namespace})
var mosquittoMessagesInflightGauge = prometheus.NewGauge(prometheus.GaugeOpts{Name: "mosquitto_messages_inflight", Help: "The number of messages with QoS>0 that are awaiting acknowledgments.", Namespace: namespace})

//clients
var mosquittoClientsTotalGauge = prometheus.NewGauge(prometheus.GaugeOpts{Name: "mosquitto_clients_total", Help: "The total number of active and inactive clients currently connected and registered on the broker.", Namespace: namespace})
var mosquittoClientsMaximumGauge = prometheus.NewGauge(prometheus.GaugeOpts{Name: "mosquitto_clients_maximum", Help: "The maximum number of clients that have been connected to the broker at the same time.", Namespace: namespace})
var mosquittoClientsConnectedGauge = prometheus.NewGauge(prometheus.GaugeOpts{Name: "mosquitto_clients_active", Help: "The number of currently connected clients.", Namespace: namespace})
var mosquittoClientsDisconnectedGauge = prometheus.NewGauge(prometheus.GaugeOpts{Name: "mosquitto_clients_disconnected", Help: "The total number of persistent clients (with clean session disabled) that are registered at the broker but are currently disconnected.", Namespace: namespace})
var mosquittoClientsExpiredGauge = prometheus.NewGauge(prometheus.GaugeOpts{Name: "mosquitto_clients_expired", Help: "The number of disconnected persistent clients that have been expired and removed through the persistent_client_expiration option.", Namespace: namespace})

//bytes
var mosquittoBytesSentGauge = prometheus.NewGauge(prometheus.GaugeOpts{Name: "mosquitto_bytes_sent", Help: "The total number of bytes sent since the broker started.", Namespace: namespace})
var mosquittoBytesReceivedGauge = prometheus.NewGauge(prometheus.GaugeOpts{Name: "mosquitto_bytes_received", Help: "The total number of bytes received since the broker started.", Namespace: namespace})

//heap
var mosquittoHeapCurrentGauge = prometheus.NewGauge(prometheus.GaugeOpts{Name: "mosquitto_heap_current", Help: "The current size of the heap memory in use by mosquitto. Note that this topic may be unavailable depending on compile time options.", Namespace: namespace})
var mosquittoHeapMaximumGauge = prometheus.NewGauge(prometheus.GaugeOpts{Name: "mosquitto_heap_maximum", Help: "The largest amount of heap memory used by mosquitto. Note that this topic may be unavailable depending on compile time options.", Namespace: namespace})

func init() {

	mqttRegistry.MustRegister(mosquittoUptimeGauge)

	mqttRegistry.MustRegister(mosquittoMessagesSentGauge)
	mqttRegistry.MustRegister(mosquittoMessagesReceivedGauge)
	mqttRegistry.MustRegister(mosquittoMessagesInflightGauge)

	mqttRegistry.MustRegister(mosquittoClientsTotalGauge)
	mqttRegistry.MustRegister(mosquittoClientsMaximumGauge)
	mqttRegistry.MustRegister(mosquittoClientsConnectedGauge)
	mqttRegistry.MustRegister(mosquittoClientsExpiredGauge)
	mqttRegistry.MustRegister(mosquittoClientsDisconnectedGauge)

	mqttRegistry.MustRegister(mosquittoBytesSentGauge)
	mqttRegistry.MustRegister(mosquittoBytesReceivedGauge)

	mqttRegistry.MustRegister(mosquittoHeapCurrentGauge)
	mqttRegistry.MustRegister(mosquittoHeapMaximumGauge)
}

func GetMqttRegistry() *prometheus.Registry {
	return mqttRegistry
}

func UptimeGauge(Uptime float64) {
	// log.Println("Uptime")
	mosquittoUptimeGauge.Set(Uptime)
}

func MessagesReceivedGauge(MessagesReceivedCount float64) {
	// log.Println("MessagesReceivedCount")
	mosquittoMessagesReceivedGauge.Set(MessagesReceivedCount)
}

func MessagesSentGauge(MessagesSentCount float64) {
	// log.Println("MessagesSentCount")
	mosquittoMessagesSentGauge.Set(MessagesSentCount)
}

func MessagesInflightGauge(MessagesInflight float64) {
	mosquittoMessagesInflightGauge.Set(MessagesInflight)
}

func ClientsTotalGauge(TotalClients float64) {
	// log.Println("TotalClients")
	mosquittoClientsTotalGauge.Set(TotalClients)
}

func ClientsMaximumGauge(MaximumClients float64) {
	// log.Println("MaximumClients")
	mosquittoClientsMaximumGauge.Set(MaximumClients)
}

func ClientsConnectedGauge(ConnectedClients float64) {
	// log.Println("ConnectedClients")
	mosquittoClientsConnectedGauge.Set(ConnectedClients)
}

func ClientsDisconnectedGauge(DisconnectedClients float64) {
	// log.Println("ConnectedClients")
	mosquittoClientsDisconnectedGauge.Set(DisconnectedClients)
}

func ClientsExpiredGauge(ExpiredClients float64) {
	// log.Println("ExpiredClients")
	mosquittoClientsExpiredGauge.Set(ExpiredClients)
}

func BytesReceivedGauge(BytesReceived float64) {
	// log.Println("BytesReceived")
	mosquittoBytesReceivedGauge.Set(float64(BytesReceived))
}

func BytesSentGauge(BytesSent float64) {
	// log.Println("BytesSent")
	mosquittoBytesSentGauge.Set(float64(BytesSent))
}

func HeapCurrentGauge(HeapCurrent float64) {
	// log.Println("HeapCurrent")
	mosquittoHeapCurrentGauge.Set(float64(HeapCurrent))
}

func HeapMaximumGauge(HeapMaximum float64) {
	// log.Println("HeapMaximum")
	mosquittoHeapMaximumGauge.Set(float64(HeapMaximum))
}
