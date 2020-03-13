package main

import (
	"fmt"
	"time"

	kiteapi "github.com/VarunBatraIT/kiteapi"
)

const (
	apiKey    string = "my_api_key"
	apiSecret string = "my_api_secret"
)

var (
	ticker *kiteapi.Ticker
)

// Triggered when any error is raised
func onError(err error) {
	fmt.Println("Error: ", err)
}

// Triggered when websocket connection is closed
func onClose(code int, reason string) {
	fmt.Println("Close: ", code, reason)
}

// Triggered when connection is established and ready to send and accept data
func onConnect() {
	fmt.Println("Connected")
	err := ticker.Subscribe([]uint32{53718535})
	if err != nil {
		fmt.Println("err: ", err)
	}
}

// Triggered when tick is recevived
func onTick(tick kiteapi.Tick) {
	fmt.Println("Tick: ", tick)
}

// Triggered when reconnection is attempted which is enabled by default
func onReconnect(attempt int, delay time.Duration) {
	fmt.Printf("Reconnect attempt %d in %fs\n", attempt, delay.Seconds())
}

// Triggered when maximum number of reconnect attempt is made and the program is terminated
func onNoReconnect(attempt int) {
	fmt.Printf("Maximum no of reconnect attempt reached: %d", attempt)
}

// Triggered when order update is received
func onOrderUpdate(order kiteapi.Order) {
	fmt.Printf("Order: %s", order.OrderID)
}

func main() {
	// Create a new Kite connect instance
	kc := kiteapi.NewConnect(apiKey)

	// Login URL from which request token can be obtained
	fmt.Println(kc.GetLoginURL())

	// Obtained request token after Kite Connect login flow
	// simulated here by scanning from stdin
	var requestToken string
	fmt.Scanf("%s\n", &requestToken)

	// Get user details and access token
	data, err := kc.GenerateSession(requestToken, apiSecret)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	// Create new Kite ticker instance
	ticker = kiteapi.NewTicker(apiKey, data.AccessToken)

	// Assign callbacks
	ticker.OnError(onError)
	ticker.OnClose(onClose)
	ticker.OnConnect(onConnect)
	ticker.OnReconnect(onReconnect)
	ticker.OnNoReconnect(onNoReconnect)
	ticker.OnTick(onTick)
	ticker.OnOrderUpdate(onOrderUpdate)

	// Start the connection
	ticker.Serve()
}
