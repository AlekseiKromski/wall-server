package app

import (
	"fmt"
	"github.com/gorilla/websocket"
)

type Client struct {
	Conn     *websocket.Conn
	security Security
}

func CreateNewClient(connection *websocket.Conn, config *Config) *Client {
	return &Client{Conn: connection, security: Security{attemptsAllowed: config.AttemptsAllowed, attemptsCount: 0}}
}

func (c *Client) Handler(app *App) {
	go c.startReceiveChannel(app)
}

func (c *Client) startReceiveChannel(app *App) {
	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			fmt.Printf("error in clinet: %v", err)
		}
		fmt.Println(string(message))
		c.security.doAttempt()
		actionHandler, err := app.ActionsWorker.defineAction(message)
		if err != nil || actionHandler == nil {
			if err == nil {
				err = fmt.Errorf("CAN'T FIND ACTION HANDLER")
			}
			fmt.Printf("action decoder error: %v", err)

			if c.security.attemptsCount >= c.security.attemptsAllowed {
				break
			}
		} else {
			actionHandler.Action.Do()
			triggerHandler, err := app.TriggersWorker.defineTrigger(actionHandler.Action.TrigType())
			if err != nil || triggerHandler == nil {
				if err == nil {
					err = fmt.Errorf("CAN'T FIND TRIGGER HANDLER")
				}
				fmt.Printf("error in trigger handler: %v", err)
			}
			triggerHandler.Action.Do()

			c.security.cleanAttempts()
			c.Conn.WriteMessage(1, []byte("OK"))
		}

	}
	defer c.Conn.Close()
}
