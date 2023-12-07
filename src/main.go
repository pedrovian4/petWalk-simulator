package main

import (
	"fmt"
	"net"
	"petWalkSimulator/src/network"
	"petWalkSimulator/src/simulation"
)

func main() {
	config := ParseFlags()

	if config.ServerAddress == "" {
		fmt.Println("Server address is required. Use the --server flag to specify it.")
		return
	}
	conn, err := network.ConnectToServer(config.ServerAddress)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("Error closing the connection:", err)
		}
	}(conn)
	simulation.SimulateMovement(conn, config.Distance, config.JWTToken)
}
