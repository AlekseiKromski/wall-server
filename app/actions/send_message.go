package actions

import (
	"fmt"
	wall_app "wall-server/wall-app"
)

type SendMessage struct {
	ActionType string
	appPayload *wall_app.WallList
}

func RegisterSendMessageHandler() *ActionHandler {
	sendMessage := SendMessage{ActionType: "sendMessage"}
	return &ActionHandler{ActionType: sendMessage.ActionType, Realisation: &sendMessage}
}

func (sm *SendMessage) Do(appPayload *wall_app.WallList, data interface{}) error {
	sm.appPayload = appPayload
	switch data.(type) {
	case string:
		sm.addMessageToWall(fmt.Sprintf("%s", data))
	default:
		return nil
	}
	return nil
}

func (sm *SendMessage) addMessageToWall(data string) {
	sm.appPayload.PutRecord(wall_app.CreateWallRecord(data))
}
