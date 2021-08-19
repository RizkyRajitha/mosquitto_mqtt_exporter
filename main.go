package main

import (
	"flag"
	"fmt"
	"log"

	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	collector "rizkyrajitha.github.io/mosquitto_exporter/collector"
	mqttcon "rizkyrajitha.github.io/mosquitto_exporter/mqttcon"
)

var version, brokerAddress, username, password, listenPort string

func main() {

	flag.Usage = func() {
		fmt.Printf("mosquitto_exporter v%s\nA Mosquitto Mqtt Broker metric exporter for Prometheus\n\nOptions\n", version)
		flag.PrintDefaults()
	}

	flag.StringVar(&brokerAddress, "brokerAddress", "", "mqtt broker ip / url and port. ex -127.0.0.1:1883 / broker.yourdomain.com:1883")
	flag.StringVar(&listenPort, "listenPort", "9992", "exporter listening address")
	flag.StringVar(&username, "username", "", "mqtt broker username")
	flag.StringVar(&password, "password", "", "mqtt password username")

	flag.Parse()

	if brokerAddress == "" {
		fmt.Println("Required brokerAddress ex - 127.0.0.1:1883 / broker.yourdomain.com:1883")
		flag.PrintDefaults()
		os.Exit(1)
	}

	fmt.Printf("mosquitto_exporter v%s\nA Mosquitto Mqtt Broker metric exporter for Prometheus\n\n", version)
	log.Println("brokerAddress : ", brokerAddress)

	if username != "" {
		log.Println("username : ", username)
	}

	http.HandleFunc("/", homePageHandler)

	registry := collector.Registry()
	handler := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})

	http.Handle("/metrics", handler)

	log.Println("mosquitto_exporter listening on port : ", listenPort)

	mqttcon.MqttCon(brokerAddress, username, password)

	err := http.ListenAndServe(":"+listenPort, nil)
	if err != nil {
		log.Fatal("Error listening on port : ", listenPort)
	}

}

func homePageHandler(w http.ResponseWriter, r *http.Request) {

	homePageHtml := `
		<html> 
		<head>	
		<title> mosquitto_exporter v` + version + `</title>
		</head>	
		<body>
		
		<h1>mosquitto_exporter v` + version + `</h1> 
		<a href="/metrics" >metrics</a> 
		
		</body>
		</html>`

	fmt.Fprint(w, homePageHtml)
}
