#!/bin/sh

CGO_ENABLED=0 GOOS=linux GOARCH=arm go build ./src/main.go
mv main httpSh_linux_arm
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./src/main.go
mv main httpSh_linux_amd64
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build ./src/main.go
mv main httpSh_mac_amd64
