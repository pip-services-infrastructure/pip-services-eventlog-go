package main

import (
	"os"

	cont "github.com/pip-services-infrastructure/pip-services-eventlog-go/container"
)

func main() {
	proc := cont.NewEventLogProcess()
	proc.SetConfigPath("./config/config.yml")
	proc.Run(os.Args)
}
