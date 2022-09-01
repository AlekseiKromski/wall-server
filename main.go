package main

import (
	"fmt"
	"wall-server/app"
)

var actionHandlers = []*app.ActionHandler{}

var triggerHandlers = []*app.TriggerHandler{}

func main() {

	_, err := app.Start(actionHandlers, triggerHandlers)
	if err != nil {
		fmt.Println(err)
	}
}
