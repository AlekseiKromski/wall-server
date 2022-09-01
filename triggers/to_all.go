package triggers

import (
	"fmt"
	"wall-server/app"
	wall_app "wall-server/wall-app"
)

type ToAll struct {
	WallAppInstance *wall_app.WallList
	data            string
	client          *app.Client
	clients         []*app.Client
}

func (ta *ToAll) SetClients(clients []*app.Client) {
	ta.clients = clients
}

func (ta *ToAll) Do() {
	ta.WallAppInstance = wall_app.GetAppInstance()
	for _, client := range ta.clients {
		if client.ID == ta.client.ID {
			continue
		}
		records := ta.WallAppInstance.GetRecords()
		lastData, err := records[len(records)-1].ToJson()
		if err != nil {
			fmt.Errorf("error in toJson method: %v", err)
			return
		}
		fmt.Println(string(lastData))
		client.Conn.WriteMessage(1, lastData)
	}
}

func (ta *ToAll) SetClient(c *app.Client) {
	ta.client = c
}

func (ta *ToAll) SetData(data string) {
	ta.data = data
}
