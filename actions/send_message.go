package actions

import (
	"wall-server/app"
	wall_app "wall-server/wall-app"
)

type SendMessage struct {
	Data            string
	WallAppInstance *wall_app.WallList
	client          *app.Client
}

func (sm *SendMessage) SetData(data string) {
	sm.Data = data
}

func (sm *SendMessage) Do() {
	sm.WallAppInstance = wall_app.GetAppInstance()
	sm.run()
}
func (sm *SendMessage) TrigType() string {
	return "to-all"
}

func (sm *SendMessage) SetClient(client *app.Client) {
	sm.client = client
}

func (sm *SendMessage) run() {
	record := sm.WallAppInstance.CreateWallRecord(sm.Data, sm.client.ID)
	sm.WallAppInstance.PutRecord(record)
}
