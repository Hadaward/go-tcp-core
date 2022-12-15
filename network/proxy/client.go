package proxy

import (
	"core/network/channels"
	"io"
	"net"
	"time"
)

type Client struct {
	Transport net.Conn
	Emitter   channels.Emitter
	polling   bool
}

func (client *Client) poll() {
	client.polling = true

	go func() {
		for client.polling {
			data := make([]byte, 2048)

			client.Transport.SetReadDeadline(time.Now().Add(1 * time.Microsecond))
			size, err := client.Transport.Read(data)

			if size > 0 && err == nil {
				client.Emitter.Trigger("data", data[:size])
			} else if err != nil && (err == io.EOF || !err.(net.Error).Timeout()) {
				client.Emitter.Trigger("disconnected")
				client.polling = false
			}
		}
	}()
}
