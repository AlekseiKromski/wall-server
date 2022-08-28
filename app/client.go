package app

import (
	"fmt"
	"github.com/gorilla/websocket"
	"wall-server/app/actions"
)

type Client struct {
	conn     *websocket.Conn
	security Security
}

func CreateNewClient(connection *websocket.Conn, config *Config) *Client {
	return &Client{conn: connection, security: Security{attemptsAllowed: config.AttemptsAllowed, attemptsCount: 0}}
}

func (c *Client) Handler(app *App) {
	go c.startReceiveChannel(app)
}

func (c *Client) startReceiveChannel(app *App) {
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			fmt.Printf("error in clinet: %v", err)
		}
		fmt.Println(string(message))
		c.security.doAttempt()
		action, err := actions.DecodeAction(message)
		if err != nil {
			fmt.Printf("action decoder error: %v", err)

			if c.security.attemptsCount >= c.security.attemptsAllowed {
				break
			}
		} else {
			action.Realisation.Do(app.WallApp, action.ActionData)
			c.security.cleanAttempts()
			c.conn.WriteMessage(1, []byte("OK"))
		}

	}
	defer c.conn.Close()
}
