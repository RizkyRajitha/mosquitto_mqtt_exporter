package main

import (
	"flag"
	"fmt"
	"log"

	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	mqttcon "rizkyrajitha.github.io/mosquitto_exporter/mqttcon"
)

var version, brokerAddress, username, password, listenPort string

func main() {

	flag.Usage = func() {
		fmt.Printf("mosquitto_exporter %s\nA Mosquitto Mqtt Broker metric exporter for Prometheus\n\nOptions\n", version)
        flag.PrintDefaults()
	}

	flag.StringVar(&brokerAddress, "brokerAddress", "127.0.0.1:1883", "mqtt broker ip / url and port")
	flag.StringVar(&listenPort, "listenPort", "9992", "exporter listening address")
	flag.StringVar(&username, "username", "", "mqtt broker username")
	flag.StringVar(&password, "password", "", "mqtt password username")

	if len(os.Args) < 2 {
		log.Fatal("expected brokerAddress")
	}

	flag.Parse()

	log.Println("brokerAddress : ", brokerAddress)
	if username != "" {
		log.Println("username : ", username)
	}

	http.Handle("/metrics", promhttp.Handler())

	log.Println("mosquitto_exporter listening on port", listenPort)

	mqttcon.MqttCon(brokerAddress, username, password)

	err := http.ListenAndServe(":"+listenPort, nil)
	if err != nil {
		log.Fatal("Error listening on port : ", listenPort)
	}

}
