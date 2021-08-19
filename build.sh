#!/bin/bash

version=0.1.0

platforms=("linux/386" "linux/amd64" "linux/arm64" )

for platform in "${platforms[@]}"
do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}

    BinaryName="mosquitto_mqtt_exporter-${version}-${GOOS}-${GOARCH}"
    echo $BinaryName

    env GOOS=$GOOS GOARCH=$GOARCH go build -ldflags="-X 'main.version=${version}'" -o $BinaryName

done


#arm build
armversions=("5" "6" "7")

for armversion in "${armversions[@]}"
do
    platform_split=(${platform//\// })
    GOOS=linux
    GOARCH=arm
    GOARM=$armversion

    BinaryName="mosquitto_mqtt_exporter-${version}-${GOOS}-${GOARCH}v${armversion}"
    echo $BinaryName

    env GOOS=$GOOS GOARCH=$GOARCH GOARM=$GOARM go build -ldflags="-X 'main.version=${version}'" -o $BinaryName

done