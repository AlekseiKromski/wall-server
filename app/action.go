package app

import "encoding/json"

type ActionHandlerInterface interface {
	SetData(data string)
	Do()
	TrigType() string
}

type ActionHandler struct {
	actionType string
	Data       string
	action     ActionHandlerInterface
}

type Action struct {
	ActionType string
	Data       string
}

type ActionsWorker struct {
	actions []*ActionHandler
}

func (aw *ActionsWorker) registerHandler(handler *ActionHandler) {
	aw.actions = append(aw.actions, handler)
}

func (aw *ActionsWorker) defineAction(message []byte) (*ActionHandler, error) {
	var action Action
	err := json.Unmarshal(message, &action)
	if err != nil {
		return nil, err
	}
	for _, actionHandler := range aw.actions {
		if actionHandler.actionType == action.ActionType {
			actionHandler.action.SetData(action.Data)
			return actionHandler, nil
		}
	}
	return nil, nil
}
