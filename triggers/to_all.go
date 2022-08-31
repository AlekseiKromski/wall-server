package triggers

import (
	"fmt"
	wall_app "wall-server/wall-app"
)

type ToAll struct {
	WallAppInstance *wall_app.WallList
	data            string
}

func (ta *ToAll) Do() {
	ta.WallAppInstance = wall_app.GetAppInstance()
	for _, client := range ta.WallAppInstance.GetClients() {
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

func (ta *ToAll) SetData(data string) {
	ta.data = data
}