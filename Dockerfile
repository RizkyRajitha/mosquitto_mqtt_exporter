FROM golang:1.16-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY main.go ./
COPY collector/* ./collector/
COPY mqttcon/* ./mqttcon/

RUN go build 

FROM alpine:latest AS runner

ENV brokerAddress=""
ENV listenPort=9992
ENV username=""
ENV password=""

WORKDIR /app

COPY --from=builder /app/mosquitto_exporter ./

EXPOSE ${listenPort}

CMD [ "/bin/sh" , "-c" , "./mosquitto_exporter -brokerAddress=${brokerAddress} -listenPort=${listenPort} -username=${username} -password=${password}} " ]
