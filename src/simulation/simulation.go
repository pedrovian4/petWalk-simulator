package simulation

import (
	"encoding/json"
	"fmt"
	"net"
	"time"
)

type Config struct {
	Distance float64
	JWTToken string
}

type Payload struct {
	Location LocationData `json:"location"`
	JWT      string       `json:"jwt,omitempty"`
}

type LocationData struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func SimulateMovement(conn net.Conn, distance float64, jwt string) {
	var lat, long float64 = 0.013183040713049118, -51.07407797173432
	latIncrement := distance / 100000.0

	for i := 0; i < int(distance*1000); i++ {
		lat += latIncrement
		payload := Payload{
			Location: LocationData{
				Latitude:  lat,
				Longitude: long,
			},
			JWT: jwt,
		}
		sendData(conn, payload)
		time.Sleep(1 * time.Second)
	}
}

func sendData(conn net.Conn, payload Payload) {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = conn.Write(jsonData)
	if err != nil {
		fmt.Println("Error sending data:", err)
		return
	}
	fmt.Println("Data sent:", payload)
}
