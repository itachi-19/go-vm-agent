[![Go](https://github.com/itachi-19/go-vm-agent/actions/workflows/go.v1.yml/badge.svg?branch=master)](https://github.com/itachi-19/go-vm-agent/actions/workflows/go.v1.yml)


# Go VM Agent

A Go based agent that can be run as both root and non-root on a node to act as an interface for automations over TCP protocol.

Currently enabled features:

## TCP Listener
A Go based TCP listener that can be used for sending and running arbitrary code over the TCP session.



### Usage
- `go run cmd/tcp-listener/main.go` on the server side or host which is hosting resources
- `telnet <server IP> 9999` on client side
- run commands like `ls`, `ps -ef` and get their real-time output in console

