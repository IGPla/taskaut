package utils

import (
	"flag"
)

/*
ParseFlags : parse function that will collect all flags from command line
*/
func ParseFlags() string {
	actionsFile := flag.String("actions-file", "./actions.json", "Actions filepath. JSON format expected")
	flag.Parse()
	return *actionsFile
}
