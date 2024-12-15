package main

import (
	"log/slog"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "google.com:80")
	if err != nil {
		slog.Error("failed to connect",
			"error", err)
	}

	defer conn.Close()

	slog.Info("successfully connected")
}
