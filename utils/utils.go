package utils

import (
	"flag"
)

func ParseFlags() string {
	actionsFile := flag.String("actions-file", "./actions.json", "Actions filepath. JSON format expected")
	flag.Parse()
	return *actionsFile
}
