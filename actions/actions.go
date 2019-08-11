package actions

import (
	"encoding/json"
	"github.com/logrusorgru/aurora"
	"io/ioutil"
	"log"
	"os/exec"
)

type Command struct {
	Command string   `json:"command"`
	Params  []string `json"params"`
}

type Action struct {
	Commands []Command `json:"commands"`
	Dir      string    `json:"dir"`
}

func (action Action) RunAction(finished chan bool) {
	defer func() { finished <- true }()
	for _, command := range action.Commands {
		log.Printf("Executing command %v, with params %v in directory %v",
			aurora.Blue(command.Command), aurora.Blue(command.Params),
			aurora.Blue(action.Dir))
		cmd := exec.Command(command.Command, command.Params...)
		cmd.Dir = action.Dir
		logs, err := cmd.CombinedOutput()
		if err != nil {
			log.Printf("Error arised while running command: %v\n%v", aurora.Red(err), aurora.Red(string(logs)))
		} else {
			log.Printf("Executed command %v, with params %v in directory %v. Output:\n%v",
				aurora.Blue(command.Command), aurora.Blue(command.Params),
				aurora.Blue(action.Dir), aurora.Green(string(logs)))
		}
	}
}

/*
Description: load actions from file and return them
*/
func LoadActions(actionsFile string) []Action {
	actions := make([]Action, 0, 1)
	rawActions, err := ioutil.ReadFile(actionsFile)
	if err != nil {
		log.Fatalf("Could not read %v file: %v", actionsFile, aurora.Red(err))
	}
	err = json.Unmarshal(rawActions, &actions)
	if err != nil {
		log.Fatalf("Could not parse %v file: %v", actionsFile, aurora.Red(err))
	}
	return actions
}
