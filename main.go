package main

import (
	"github.com/IGPla/taskaut/actions"
	"github.com/IGPla/taskaut/utils"
	"sync"
)

func main() {
	actionsFile := utils.ParseFlags()
	allActions := actions.LoadActions(actionsFile)

	var wg sync.WaitGroup
	for _, action := range allActions {
		wg.Add(1)
		go action.RunAction(&wg)
	}
	wg.Wait()
}
