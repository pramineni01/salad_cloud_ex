# README

### Description
This source code provides a binary tcp client, which connects to a tcp server, waits till any message appears, receives a message, and prints it.
At a high level, this provides a tcp client, which connects to a remote server, receive binary data into a buffer, parses the message into a custom struct and simply prints it to the standard output stream.

### Prerequisites
Binary data source: A TCP server is required, which sends the message in certain binary format.

### Dependency variables
TCP Source url, and port - Using hardcoded values for this exercise.

### How to build
1. ```cd``` to sala_cloud_ex/cmd directory
2. ```CGO_ENABLED=0 GOARCH=amd64 GO111MODULE=off go build -o ../salad_cloud_processor ./...```
3. One shoud see sala_cloud_ex/salad_cloud_processor binary

### Yet to be done
1. Makefile - With targets for building code, building docker, tagging and push to registry (given a container registry)
2. Dockerfile

