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
        fmt.Println("expected brokerIp")
        os.Exit(1)
    }
	var brokerIp , username , password string

	flag.StringVar(&brokerIp, "brokerIp", "127.0.0.1:1883", "mqtt broker ip and port")
	flag.StringVar(&username, "username", "", "mqtt broker username")
	flag.StringVar(&password, "password", "", "mqtt password username")

	flag.Parse()

	fmt.Println("brokerIp:", brokerIp)
	fmt.Println("username:", username)
	fmt.Println("password:", password)

	mqttcon.MqttCon(brokerIp ,username ,password)

	http.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(":9999", nil)
	// log.Fatal(http.ListenAndServe(*addr, nil))
	// log.Fatal()

}


