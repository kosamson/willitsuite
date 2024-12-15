package main

import (
	"fmt"
	"log/slog"
	"net"
	"strconv"
	"time"
)

type protocol int

const (
	protocolIP protocol = iota
	protocolTCP
	protocolUDP
)

var protocolName = map[protocol]string{
	protocolIP:  "ip:1",
	protocolTCP: "tcp",
	protocolUDP: "udp",
}

const dialTimeout time.Duration = 3 * time.Second

func main() {
	if err := canConnect(protocolIP, "127.0.0.1", -1); err != nil {
		slog.Error("failed to ping localhost",
			"error", err)
	} else {
		slog.Info("successfully pinged localhost")
	}

	if err := canConnect(protocolTCP, "google.com", 80); err != nil {
		slog.Error("failed to connect to google on port 80",
			"error", err)
	} else {
		slog.Info("successfully connected to google on port 80")
	}

	if err := canConnect(protocolTCP, "google.com", 81); err != nil {
		slog.Error("failed to connect to google on port 81",
			"error", err)
	} else {
		slog.Info("successfully connected to google on port 81")
	}
}

// Asserts whether the host running this application
// can connect to the specified address (optionally on the specified port).
//
// For IP protocol, set port to -1.
func canConnect(protocol protocol, address string, port int) error {
	if port > 0 {
		address = address + ":" + strconv.Itoa(port)
	}

	conn, err := net.DialTimeout(protocolName[protocol], address, dialTimeout)
	if err != nil {
		return fmt.Errorf("failed to connect: %w", err)
	}

	defer conn.Close()

	return nil
}
