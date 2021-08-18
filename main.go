package main

import (
	"flag"
	"log"

	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	mqttcon "rizkyrajitha.github.io/mqttexporter/mqttcon"
)


var brokerAddress , username , password string

func main() {
	
	flag.StringVar(&brokerAddress, "brokerAddress", "127.0.0.1:1883", "mqtt broker ip / url and port")
	flag.StringVar(&username, "username", "", "mqtt broker username")
	flag.StringVar(&password, "password", "", "mqtt password username")
    
	if len(os.Args) < 2 {
		log.Fatal("expected brokerAddress")
    }

	flag.Parse()

	log.Println("brokerAddress : ", brokerAddress)
	if(username != ""){
		log.Println("username : ", username)
	}

	mqttcon.MqttCon(brokerAddress ,username ,password)

	http.Handle("/metrics", promhttp.Handler())

	 err :=  http.ListenAndServe(":9999", nil)
	 if (err != nil){
		 log.Fatal("Error listening on port : ",9999)
	 }

}