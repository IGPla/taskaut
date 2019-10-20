package actions

import (
	"encoding/json"
	"github.com/logrusorgru/aurora"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"strings"
)

/*
LoadActions : function that will load all actions and commands from a configuration file
*/
func LoadActions(actionsFile string) []Action {
	var err error

	actions := make([]Action, 0, 1)
	rawActions, err := ioutil.ReadFile(actionsFile)
	if err != nil {
		log.Fatalf("Could not read %v file: %v", actionsFile, aurora.Red(err))
	}
	actionsFileLower := strings.ToLower(actionsFile)
	if strings.HasSuffix(actionsFileLower, ".json") {
		err = json.Unmarshal(rawActions, &actions)
	} else if strings.HasSuffix(actionsFile, ".yaml") || strings.HasSuffix(actionsFile, ".yml") {
		err = yaml.Unmarshal(rawActions, &actions)
	} else {
		log.Fatalf("Actions file format not supported. Please provide a .yaml or .json file")
	}
	if err != nil {
		log.Fatalf("Could not parse %v file: %v", actionsFile, aurora.Red(err))
	}
	return actions
}
