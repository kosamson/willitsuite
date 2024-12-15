package main

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"connectrpc.com/connect"
	willitsuitev1 "github.com/kosamson/willitsuite/gen/proto/v1"
	"github.com/kosamson/willitsuite/gen/proto/v1/willitsuitev1connect"
)

type protocol int

const (
	protocolIP protocol = iota
	protocolTCP
	protocolUDP
)

var textInputProtocol = map[string]protocol{
	"ip":  protocolIP,
	"tcp": protocolTCP,
	"udp": protocolUDP,
}

var protoInputProtocol = map[willitsuitev1.Protocol]protocol{
	willitsuitev1.Protocol_PROTOCOL_IP:  protocolIP,
	willitsuitev1.Protocol_PROTOCOL_TCP: protocolTCP,
	willitsuitev1.Protocol_PROTOCOL_UDP: protocolUDP,
}

var protocolName = map[protocol]string{
	protocolIP:  "ip:1",
	protocolTCP: "tcp",
	protocolUDP: "udp",
}

const dialTimeout time.Duration = 3 * time.Second

type WillitSuiteServer struct{}

func (ws *WillitSuiteServer) Connect(
	ctx context.Context,
	req *connect.Request[willitsuitev1.ConnectRequest],
) (*connect.Response[willitsuitev1.ConnectResponse], error) {
	slog.InfoContext(ctx, "received request")

	var status string

	protocol, ok := protoInputProtocol[req.Msg.GetProtocol()]
	if !ok {
		return nil, connect.NewError(
			connect.CodeInvalidArgument,
			fmt.Errorf("invalid protocol: %s", req.Msg.GetProtocol()),
		)
	}

	if err := canConnect(protocol, req.Msg.GetAddress(), int(req.Msg.GetPort())); err != nil {
		status = fmt.Sprintf("failed to connect: %s", err)
	} else {
		status = "successfully connected"
	}

	slog.InfoContext(ctx, "sending response")

	return connect.NewResponse(&willitsuitev1.ConnectResponse{
		Status: status,
	}), nil
}

func main() {
	//	if err := canConnect(protocolIP, "127.0.0.1", -1); err != nil {
	//		slog.Error("failed to ping localhost",
	//			"error", err)
	//	} else {
	//		slog.Info("successfully pinged localhost")
	//	}
	//
	//	if err := canConnect(protocolTCP, "google.com", 80); err != nil {
	//		slog.Error("failed to connect to google on port 80",
	//			"error", err)
	//	} else {
	//		slog.Info("successfully connected to google on port 80")
	//	}
	//
	//	if err := canConnect(protocolTCP, "google.com", 81); err != nil {
	//		slog.Error("failed to connect to google on port 81",
	//			"error", err)
	//	} else {
	//		slog.Info("successfully connected to google on port 81")
	//	}

	listener, err := net.Listen(protocolName[protocolTCP], ":8080")
	if err != nil {
		panic(err)
	}

	srv := &WillitSuiteServer{}
	path, handler := willitsuitev1connect.NewWillitSuiteServiceHandler(srv)

	http.Handle(path, handler)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		slog.InfoContext(ctx, "received request")

		tmpl, err := template.New("site.html.go.tmpl").ParseFiles("site.html.go.tmpl")
		if err != nil {
			panic(err)
		}

		protocolValue := r.FormValue("protocol")

		if protocolValue == "" {
			protocolValue = "tcp"
		}

		addressValue := r.FormValue("address")
		portValue := r.FormValue("port")

		var status string

		if addressValue != "" || portValue != "" {
			protocol, ok := textInputProtocol[protocolValue]
			if !ok {
				status = "invalid protocol"
			}

			port, err := strconv.Atoi(portValue)
			if err != nil {
				status = "invalid port value"
			}

			if status == "" {
				if err = canConnect(protocol, addressValue, port); err != nil {
					status = fmt.Sprintf("failed to connect: %s", err)
				} else {
					status = "successfully connected"
				}
			}

			slog.InfoContext(ctx, "connection tested",
				"status", status)

		}

		connectData := struct {
			InputProtocol string
			InputAddress  string
			InputPort     string
			Status        string
		}{
			InputProtocol: protocolValue,
			InputAddress:  addressValue,
			InputPort:     portValue,
			Status:        status,
		}

		if err := tmpl.Execute(w, connectData); err != nil {
			panic(err)
		}

		slog.InfoContext(ctx, "executed html template")
	})

	slog.Info("listening on port 8080")

	if err := http.Serve(listener, nil); err != nil {
		panic(err)
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
