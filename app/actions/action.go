package actions

import (
	"encoding/json"
	"fmt"
)

var actions []*ActionHandler

//From message
type Action struct {
	ActionType string `json:"actionType"`
	Data       string `json:"data"`
}

//For internal registration
type ActionHandler struct {
	ActionType  string
	Realisation ActionHandlerInterface
}

//For standart
type ActionHandlerInterface interface {
	Do() error
}

func DecodeAction(buffer []byte) (*ActionHandler, error) {
	action := Action{}
	if err := json.Unmarshal(buffer, &action); err != nil {
		return &ActionHandler{}, fmt.Errorf("can't unmarshal json to action: %v", err)
	}
	actionHandler, err := action.defineActionInRegistered()
	if err != nil {
		return nil, err
	}
	return actionHandler, nil
}

func RegisterActions() {
	actions = []*ActionHandler{
		RegisterSendMessageHandler(),
	}
}

func (a *Action) defineActionInRegistered() (*ActionHandler, error) {
	for _, action := range actions {
		if a.ActionType == action.ActionType {
			return action, nil
		}
	}
	return nil, fmt.Errorf("can't find")
}
