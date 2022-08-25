package app

import (
	"fmt"
	"net"
)

type Client struct {
	Info string
}

type SocketClient struct {
	socket    net.Conn
	broadcast chan []byte
}

func (c *Client) Handler(connection net.Conn) {
	socketClient := SocketClient{
		socket:    connection,
		broadcast: make(chan []byte),
	}
	//go startListeners(&socketClient)
	go startReceiveChannel(&socketClient)

}

func startReceiveChannel(client *SocketClient) {
	for {
		buffer := make([]byte, 1024)
		mLen, err := client.socket.Read(buffer)
		if err != nil {
			fmt.Printf("error during client handling <reading> :%s \n", err)
			client.socket.Close()
			return
		}
		fmt.Println("\nReceived: ", string(buffer[:mLen]))
		//action, err := actions.DecodeAction(buffer)
		//if err != nil {
		//	fmt.Println(err)
		//	continue
		//}
		//action.Realisation.Do()
		_, err = client.socket.Write([]byte("Thanks! Got your message:" + string(buffer[:mLen])))
		fmt.Println(string(buffer))
	}
}
