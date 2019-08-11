package main

import (
	"github.com/IGPla/taskaut/actions"
	"github.com/IGPla/taskaut/utils"
)

func main() {
	actionsFile := utils.ParseFlags()
	allActions := actions.LoadActions(actionsFile)

	finished := make(chan bool, len(allActions))
	for _, action := range allActions {
		go action.RunAction(finished)
	}
	for range allActions {
		<-finished
	}

}
