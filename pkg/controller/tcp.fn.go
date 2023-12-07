package controller

import (
	"bufio"
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
	// zap.L().Sugar().Info("[TCP] Stopped listening")
}

func (t *TCPServer) handleConnection(conn net.Conn) {
	for {
		// Read incoming data
		s, err := bufio.NewReader(conn).ReadString('\n')

		if err != nil {
			if err.Error() != "EOF" {
				zap.L().Sugar().Errorln(err)
			}
			conn.Close()
			zap.L().Sugar().Info("[TCP] Closed connection")
			return
		}

		// Print the incoming data
		zap.L().Sugar().Infof("[TCP] Received: %s", s)

		if s == "EOF" {
			conn.Close()
			zap.L().Sugar().Info("[TCP] Closed connection")
			return
		} else {
			conn.Write([]byte("ACK\n"))
		}

	}
}
