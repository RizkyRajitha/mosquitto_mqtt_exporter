package main

import (
	"flag"
	"fmt"

	// "log"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	mqttcon "rizkyrajitha.github.io/mqttexporter/mqttcon"
	// "time"
)


func main() {

    if len(os.Args) < 2 {
        fmt.Println("expected brokerAddress")
        os.Exit(1)
    }

	var brokerAddress , username , password string

	flag.StringVar(&brokerAddress, "brokerAddress", "127.0.0.1:1883", "mqtt broker ip / url and port")
	flag.StringVar(&username, "username", "", "mqtt broker username")
	flag.StringVar(&password, "password", "", "mqtt password username")

	flag.Parse()

	fmt.Println("brokerAddress:", brokerAddress)
	fmt.Println("username:", username)
	fmt.Println("password:", password)

	mqttcon.MqttCon(brokerAddress ,username ,password)

	http.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(":9999", nil)
	// log.Fatal(http.ListenAndServe(*addr, nil))
	// log.Fatal()

}


