package network

import (
	"fmt"
	"net"
)

func ConnectToServer(serverAddress string) (net.Conn, error) {
	conn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		return nil, fmt.Errorf("error connecting to server: %v", err)
	}
	return conn, nil
}
