package main

import (
	"flag"
)

type Config struct {
	Distance      float64
	JWTToken      string
	ServerAddress string
}

func ParseFlags() Config {
	var config Config

	flag.Float64Var(&config.Distance, "distance", 0, "Total distance in km for the simulation")
	flag.StringVar(&config.JWTToken, "jwt", "", "JWT token for authentication")
	flag.StringVar(&config.ServerAddress, "server", "", "TCP server address (ip:port)")
	flag.Parse()

	return config
}
