package actions

import (
	"github.com/logrusorgru/aurora"
	"log"
	"os/exec"
	"sync"
)

/*
Command : struct to hold a single command plus all parameters
*/
type Command struct {
	Command string   `json:"command"`
	Params  []string `json:"params"`
	Retries int      `json:"retries"`
}

/*
Action : struct to bundle commands around a directory
*/
type Action struct {
	Commands    []Command `json:"commands"`
	Dir         string    `json:"dir"`
	AbortOnFail bool      `json:"abort_on_fail"`
}

/*
runCommand : method on command that will run that command
*/
func (command Command) runCommand(dir string) error {
	log.Printf("Executing command %v, with params %v in directory %v",
		aurora.Blue(command.Command), aurora.Blue(command.Params),
		aurora.Blue(dir))
	cmd := exec.Command(command.Command, command.Params...)
	cmd.Dir = dir
	logs, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error arose while running command %v, with params %v in directory %v: %v\n%v",
			aurora.Blue(command.Command), aurora.Blue(command.Params),
			aurora.Blue(dir), aurora.Red(err), aurora.Red(string(logs)))
		return err
	}
	log.Printf("Executed command %v, with params %v in directory %v. Output:\n%v",
		aurora.Blue(command.Command), aurora.Blue(command.Params),
		aurora.Blue(dir), aurora.Green(string(logs)))
	return nil
}

/*
RunAction : method on action that will run all chained command included in Action, sequentially
*/
func (action Action) RunAction(wg *sync.WaitGroup) {
	defer wg.Done()
	for _, command := range action.Commands {
		retries := 0
		abort := false
		for retries <= command.Retries {
			err := command.runCommand(action.Dir)
			if err == nil {
				break
			} else if retries == command.Retries && action.AbortOnFail {
				abort = true
				break
			}
			retries++
		}
		if abort {
			log.Printf("Aborting rest of commands due to reach max retries on %v, with params %v in directory %v",
				aurora.Blue(command.Command), aurora.Blue(command.Params), aurora.Blue(action.Dir))
			break
		}
	}
}
