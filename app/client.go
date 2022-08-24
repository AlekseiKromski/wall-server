package app

import (
	"fmt"
	"net"
)

type Client struct {
	Info string
}

func (c *Client) Handler(connection net.Conn) {
	buffer := make([]byte, 1024)
	mLen, err := connection.Read(buffer)
	if err != nil {
		fmt.Printf("error during client handling <reading> :%s \n", err)
	}
	fmt.Println("Received: ", string(buffer[:mLen]))
	_, err = connection.Write([]byte("Thanks! Got your message:" + string(buffer[:mLen])))
	connection.Close()
}
