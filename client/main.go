package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"os"
)

type Client struct {
	conn *websocket.Conn
}

var c *Client

func main() {
	fmt.Println("Start client")
	fmt.Println("Client can just connect and send default hello world message")

	var input string
	for {
		fmt.Println("\n-------------\n")
		fmt.Println("Please, select option:")
		fmt.Println("0. exit")
		fmt.Println("1. connect to host")
		fmt.Println("2. listen")
		fmt.Scanln(&input)

		switch input {
		case "0":
			return
		case "1":
			connect()
		case "2":
			if c != nil && c.conn != nil {
				listen()
				return
			}
			fmt.Println("Need connection")
			connect()
			listen()
		default:
			continue
		}
	}
}

type sm struct {
	ActionType string `json:"action_type"`
	Data       string `json:"data"`
}

func connect() {

	var connectIp string
	fmt.Print("Connect to [hostname:port] (d - default <localhost:3000>): ")
	fmt.Scan(&connectIp)
	if connectIp == "d" {
		connectIp = "localhost:3000"
	}
	conn, _, err := websocket.DefaultDialer.Dial(fmt.Sprintf("ws://%s/", connectIp), nil)
	c = &Client{
		conn: conn,
	}
	if err != nil {
		fmt.Println("Error connecting to Websocket Server:", err)
	}
	fmt.Printf("Connected: %s\n", conn.RemoteAddr())
}

func send(data string) {
	message := sm{ActionType: "send-message", Data: data}
	messageJson, err := json.Marshal(message)
	if err != nil {
		fmt.Println("can't send message")
	}
	c.conn.WriteMessage(1, []byte(messageJson))
	fmt.Println("[YOU]: ", data)
}

func listen() {
	scanner := bufio.NewScanner(os.Stdin)
	go func() {
		for scanner.Scan() {
			send(scanner.Text())
		}
	}()
	for {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			fmt.Println("Error in receive:", err)
			return
		}
		fmt.Printf("[CLIENT]: %s\n", string(msg))

	}
}
