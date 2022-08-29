package app

type TriggerHandlerInterface interface {
	Do()
}

type TriggerHandler struct {
	triggerType string
	data        string
	action      TriggerHandlerInterface
}

type TriggersWorker struct {
	triggers []*TriggerHandler
}

func (th *TriggersWorker) registerHandler(handler *TriggerHandler) {
	th.triggers = append(th.triggers, handler)
}

func (th *TriggersWorker) defineTrigger(triggerType string) (*TriggerHandler, error) {
	for _, triggerHandler := range th.triggers {
		if triggerHandler.triggerType == triggerType {
			return triggerHandler, nil
		}
	}
	return nil, nil
}
