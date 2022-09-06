package main

import (
	"fmt"
	"wall-server/actions"
	"wall-server/app"
	"wall-server/triggers"
)

var actionHandlers = []*app.ActionHandler{
	{
		ActionType: "send-message",
		Action:     &actions.SendMessage{},
	},
}

var triggerHandlers = []*app.TriggerHandler{
	{
		TriggerType: "to-all",
		Action:      &triggers.ToAll{},
	},
}

func main() {

	_, err := app.Start(actionHandlers, triggerHandlers)
	if err != nil {
		fmt.Println(err)
	}
}
