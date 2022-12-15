package proxy

import (
	"core/network/channels"
	"core/utils/collections"
	"fmt"
	"net"
	"time"
)

type Server struct {
	listeners collections.List[*net.TCPListener]
	Emitter   channels.Emitter
	running   bool
}

func (server *Server) Listen(port int) error {
	listener, err := net.Listen("tcp4", fmt.Sprintf("0.0.0.0:%d", port))

	if err != nil {
		return err
	}

	server.listeners = server.listeners.Add(listener.(*net.TCPListener))
	return nil
}

func (server *Server) Start() {
	server.running = true

	server.Emitter.Trigger("running")

	go func() {
		for server.running {
			for _, listener := range server.listeners {
				listener.SetDeadline(time.Now().Add(1 * time.Microsecond))

				conn, err := listener.Accept()

				if err != nil {
					server.Emitter.Trigger("error", err)
				}

				if conn != nil {
					client := &Client{
						Transport: conn,
					}

					server.Emitter.Trigger("newclient", client)

					client.poll()
				}
			}
		}

		server.Emitter.Trigger("exit")
	}()
}

func (server *Server) Stop() {
	server.running = false
}
