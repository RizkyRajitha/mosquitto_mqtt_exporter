## Prometheus exporter for Mosquitto mqtt broker

[![Publish-GHRC](https://github.com/RizkyRajitha/mosquitto_mqtt_exporter/actions/workflows/publishghrc.yml/badge.svg)](https://github.com/RizkyRajitha/mosquitto_mqtt_exporter/actions/workflows/publishghrc.yml)
[![Go-Build](https://github.com/RizkyRajitha/mosquitto_mqtt_exporter/actions/workflows/build.yml/badge.svg)](https://github.com/RizkyRajitha/mosquitto_mqtt_exporter/actions/workflows/build.yml)
[![Release](https://github.com/RizkyRajitha/mosquitto_mqtt_exporter/actions/workflows/release.yml/badge.svg)](https://github.com/RizkyRajitha/mosquitto_mqtt_exporter/actions/workflows/release.yml)
![GitHub all releases](https://img.shields.io/github/downloads/rizkyrajitha/mosquitto_mqtt_exporter/total)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/rizkyrajitha/mosquitto_mqtt_exporter)

<!-- A Mosquitto Mqtt Broker metric exporter for Prometheus  -->

## Flags

  - `brokerAddress` (required) - mqtt broker ip / url and port (default "127.0.0.1:1883")
  - `listenPort`    (optional) - exporter listening address (default "9992")
  - `password`      (optional) - mqtt password username
  - `username`      (optional) - mqtt broker username

## Run exporter 

### Binary

```bash
./mosquitto_mqtt_exporter --brokerAddress=localhost:1883`
```

### Docker

#### From GHRC
```bash
docker run -d -p 9992:9992 -e brokerAddress=localhost:1883  ghcr.io/rizkyrajitha/mosquitto_mqtt_exporter:latest
```

#### From source
```bash

docker build . --tag=mosquitto_mqtt_exporter

docker run -p 9992:9992 -e brokerAddress=localhost:1883 mosquitto_mqtt_exporter

# with host networking https://stackoverflow.com/a/24326540
docker run -p 9992:9992 --add-host host.docker.internal:host-gateway -e brokerAddress=host.docker.internal:1883  mosquitto_mqtt_exporter

```

## Metrics

|Mosquitto Topic|Prometheus Metric|Description|
|----------------|-----------------|-----------|
$SYS/broker/uptime |mosquitto_uptime| mosquitto broker uptime
$SYS/broker/messages/sent|mosquitto_messages_sent|The total number of messages of any type sent since the broker started.
$SYS/broker/messages/received|mosquitto_messages_received|The total number of messages of any type received since the broker started.
$SYS/broker/messages/inflight|mosquitto_messages_inflight|The number of messages with QoS>0 that are awaiting acknowledgments.
$SYS/broker/clients/total|mosquitto_clients_total|The total number of active and inactive clients currently connected and registered on the broker.
$SYS/broker/clients/maximum|mosquitto_clients_maximum|The maximum number of clients that have been connected to the broker at the same time.
$SYS/broker/clients/active|mosquitto_clients_active|The number of currently connected clients.
$SYS/broker/bytes/sent|mosquitto_bytes_sent|The total number of bytes sent since the broker started.
$SYS/broker/bytes/received|mosquitto_bytes_received|The total number of bytes received since the broker started.
$SYS/broker/heap/current|mosquitto_heap_current|The current size of the heap memory in use by mosquitto. Note that this topic may be unavailable depending on compile time options.
$SYS/broker/heap/maximum|mosquitto_heap_maximum|The largest amount of heap memory used by mosquitto. Note that this topic may be unavailable depending on compile time options.


![grafana dashboard](grafana.png)

[mosquitto documentation ]( https://mosquitto.org/documentation/)

[mosquitto metrics man page ]( https://mosquitto.org/man/mosquitto-8.html)

## Metrics example

```
# HELP mosquitto_mqtt_mosquitto_bytes_received The total number of bytes received since the broker started.
# TYPE mosquitto_mqtt_mosquitto_bytes_received gauge
mosquitto_mqtt_mosquitto_bytes_received 2250
# HELP mosquitto_mqtt_mosquitto_bytes_sent The total number of bytes sent since the broker started.
# TYPE mosquitto_mqtt_mosquitto_bytes_sent gauge
mosquitto_mqtt_mosquitto_bytes_sent 42054
# HELP mosquitto_mqtt_mosquitto_clients_active The number of currently connected clients.
# TYPE mosquitto_mqtt_mosquitto_clients_active gauge
mosquitto_mqtt_mosquitto_clients_active 1
# HELP mosquitto_mqtt_mosquitto_clients_disconnected The total number of persistent clients (with clean session disabled) that are registered at the broker but are currently disconnected.
# TYPE mosquitto_mqtt_mosquitto_clients_disconnected gauge
mosquitto_mqtt_mosquitto_clients_disconnected 0
# HELP mosquitto_mqtt_mosquitto_clients_expired The number of disconnected persistent clients that have been expired and removed through the persistent_client_expiration option.
# TYPE mosquitto_mqtt_mosquitto_clients_expired gauge
mosquitto_mqtt_mosquitto_clients_expired 0
# HELP mosquitto_mqtt_mosquitto_clients_maximum The maximum number of clients that have been connected to the broker at the same time.
# TYPE mosquitto_mqtt_mosquitto_clients_maximum gauge
mosquitto_mqtt_mosquitto_clients_maximum 23
# HELP mosquitto_mqtt_mosquitto_clients_total The total number of active and inactive clients currently connected and registered on the broker.
# TYPE mosquitto_mqtt_mosquitto_clients_total gauge
mosquitto_mqtt_mosquitto_clients_total 1
# HELP mosquitto_mqtt_mosquitto_heap_current The current size of the heap memory in use by mosquitto. Note that this topic may be unavailable depending on compile time options.
# TYPE mosquitto_mqtt_mosquitto_heap_current gauge
mosquitto_mqtt_mosquitto_heap_current 54680
# HELP mosquitto_mqtt_mosquitto_heap_maximum The largest amount of heap memory used by mosquitto. Note that this topic may be unavailable depending on compile time options.
# TYPE mosquitto_mqtt_mosquitto_heap_maximum gauge
mosquitto_mqtt_mosquitto_heap_maximum 99040
# HELP mosquitto_mqtt_mosquitto_messages_inflight The number of messages with QoS>0 that are awaiting acknowledgments.
# TYPE mosquitto_mqtt_mosquitto_messages_inflight gauge
mosquitto_mqtt_mosquitto_messages_inflight 0
# HELP mosquitto_mqtt_mosquitto_messages_received The total number of messages of any type received since the broker started.
# TYPE mosquitto_mqtt_mosquitto_messages_received gauge
mosquitto_mqtt_mosquitto_messages_received 135
# HELP mosquitto_mqtt_mosquitto_messages_sent The total number of messages of any type sent since the broker started.
# TYPE mosquitto_mqtt_mosquitto_messages_sent gauge
mosquitto_mqtt_mosquitto_messages_sent 1632
# HELP mosquitto_mqtt_mosquitto_uptime The Mosquitto broker uptime
# TYPE mosquitto_mqtt_mosquitto_uptime gauge
mosquitto_mqtt_mosquitto_uptime 3378

```

## Build exporter 

```bash
./build.sh
```
Build binaries for following platforms

1. linux/386 
2. linux/amd64
3. linux/arm64
4. linux/armv7
5. linux/armv6
6. linux/armv5


<!-- zap logging
 -->
