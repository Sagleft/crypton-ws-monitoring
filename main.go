package main

import (
	"log"

	utopiago "github.com/Sagleft/utopialib-go"
)

var (
	wsPort int = 25000
)

func main() {
	// create client
	client := utopiago.UtopiaClient{
		Protocol: "http",
		Host:     "127.0.0.1",
		Token:    "your utopia api token",
		Port:     20000,
		WsPort:   wsPort,
	}

	// check connection
	if !client.CheckClientConnection() {
		log.Fatalln("failed to connect to Utopia client")
	}

	// enable ws
	err := client.SetWebSocketState(utopiago.SetWsStateTask{
		Enabled:       true,
		Port:          wsPort,
		EnableSSL:     false,
		Notifications: "contact",
	})
	if err != nil {
		log.Fatalln(err)
	}

	// subscribe to ws events
	err = client.WsSubscribe(utopiago.WsSubscribeTask{
		OnConnected: handleWsConnected,
		Callback:    handleWsEvent,
		ErrCallback: handleWsError,
		Port:        wsPort,
	})
	if err != nil {
		log.Fatalln(err)
	}
}

func handleWsConnected() {
	log.Println("ws connection establsihed")
}

func handleWsEvent(event utopiago.WsEvent) {
	log.Println(event)
}

func handleWsError(err error) {
	log.Println("ERROR: " + err.Error())
}
