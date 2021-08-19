package metrics

import (
	// "log"

	prometheus "github.com/prometheus/client_golang/prometheus"
)

var namespace = "mosquitto_mqtt"

var registry = prometheus.NewRegistry()

var mosquittoUptimeGauge = prometheus.NewGauge(prometheus.GaugeOpts{Name: "mosquitto_uptime", Help: "mosquitto broker uptime", Namespace: namespace})

//messages
var mosquittoMessagesSentGauge = prometheus.NewGauge(prometheus.GaugeOpts{Name: "mosquitto_messages_sent", Help: "mosquitto broker messages sent", Namespace: namespace})
var mosquittoMessagesReceivedGauge = prometheus.NewGauge(prometheus.GaugeOpts{Name: "mosquitto_messages_received", Help: "mosquitto broker messages received", Namespace: namespace})

//clients
var mosquittoClientsTotalGauge = prometheus.NewGauge(prometheus.GaugeOpts{Name: "mosquitto_clients_total", Help: "The total number of active and inactive clients currently connected and registered on the broker.", Namespace: namespace})
var mosquittoClientsMaximumGauge = prometheus.NewGauge(prometheus.GaugeOpts{Name: "mosquitto_clients_maximum", Help: "The maximum number of clients that have been connected to the broker at the same time.", Namespace: namespace})
var mosquittoClientsConnectedGauge = prometheus.NewGauge(prometheus.GaugeOpts{Name: "mosquitto_clients_active", Help: "The number of currently connected clients.", Namespace: namespace})

//bytes
var mosquittoBytesSentGauge = prometheus.NewGauge(prometheus.GaugeOpts{Name: "mosquitto_bytes_sent", Help: "The total number of bytes sent since the broker started.", Namespace: namespace})
var mosquittoBytesReceivedGauge = prometheus.NewGauge(prometheus.GaugeOpts{Name: "mosquitto_bytes_received", Help: "The total number of bytes received since the broker started.", Namespace: namespace})

//heap
var mosquittoHeapCurrentGauge = prometheus.NewGauge(prometheus.GaugeOpts{Name: "mosquitto_heap_current", Help: "The current size of the heap memory in use by mosquitto. Note that this topic may be unavailable depending on compile time options.", Namespace: namespace})
var mosquittoHeapMaximumGauge = prometheus.NewGauge(prometheus.GaugeOpts{Name: "mosquitto_heap_maximum", Help: "The largest amount of heap memory used by mosquitto. Note that this topic may be unavailable depending on compile time options.", Namespace: namespace})

func init() {

	registry.MustRegister(mosquittoUptimeGauge)

	registry.MustRegister(mosquittoMessagesSentGauge)
	registry.MustRegister(mosquittoMessagesReceivedGauge)

	registry.MustRegister(mosquittoClientsTotalGauge)
	registry.MustRegister(mosquittoClientsMaximumGauge)
	registry.MustRegister(mosquittoClientsConnectedGauge)

	registry.MustRegister(mosquittoBytesSentGauge)
	registry.MustRegister(mosquittoBytesReceivedGauge)

	registry.MustRegister(mosquittoHeapCurrentGauge)
	registry.MustRegister(mosquittoHeapMaximumGauge)
}

func Registry() *prometheus.Registry {
	return registry
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
