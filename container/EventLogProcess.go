package container

import (
	build "github.com/pip-services-infrastructure/pip-services-eventlog-go/build"
	cproc "github.com/pip-services3-go/pip-services3-container-go/container"
	rbuild "github.com/pip-services3-go/pip-services3-rpc-go/build"
)

type EventLogProcess struct {
	cproc.ProcessContainer
}

func NewEventLogProcess() *EventLogProcess {
	c := EventLogProcess{}
	c.ProcessContainer = *cproc.NewProcessContainer("eventlog", "System event logging microservice")
	c.AddFactory(build.NewEventLogServiceFactory())
	c.AddFactory(rbuild.NewDefaultRpcFactory())
	return &c
}
