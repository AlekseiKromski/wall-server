package actions

import (
	"fmt"
)

type SendMessage struct {
	ActionType string
}

func RegisterSendMessageHandler() *ActionHandler {
	sendMessage := SendMessage{ActionType: "sendMessage"}
	return &ActionHandler{ActionType: sendMessage.ActionType, Realisation: &sendMessage}
}

func (sm *SendMessage) Do() error {
	fmt.Println("Do something")
	return nil
}
