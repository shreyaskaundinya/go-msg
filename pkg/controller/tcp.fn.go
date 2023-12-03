package controller

import (
	"fmt"
	"net"

	"go.uber.org/zap"
)

type TCPServer struct {
	// TODO : switch to Gain Framework when support for windows arrives
	Listener *net.Listener
}

func newServer(hostname string, port int) *TCPServer {
	t := &TCPServer{}
	t.Init(hostname, port)
	return t
}

func (t *TCPServer) Init(hostname string, port int) {
	tcp, err := net.Listen(
		"tcp",
		fmt.Sprintf("%s:%d", hostname, port),
	)

	if err != nil {
		zap.L().Sugar().Panicln(err)
		return
	}

	zap.L().Sugar().Infof("[TCP] Listening on %s:%d", hostname, port)
	t.Listener = &tcp
}

func (t *TCPServer) Serve() {
	for {
		conn, err := (*t.Listener).Accept()
		if err != nil {
			zap.L().Sugar().Errorln(err)
			continue
		}

		go t.handleConnection(conn)
	}
}

func (t *TCPServer) handleConnection(conn net.Conn) {
	// Close the connection when we're done
	defer conn.Close()

	// Read incoming data
	buf := make([]byte, 2048)
	_, err := conn.Read(buf)

	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the incoming data
	fmt.Printf("[TCP] Received: %s", buf)
}
