package services

import (
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
	services "github.com/pip-services3-go/pip-services3-rpc-go/services"
)

type EventLogHttpServiceV1 struct {
	*services.CommandableHttpService
}

func NewEventLogHttpServiceV1() *EventLogHttpServiceV1 {
	c := EventLogHttpServiceV1{
		CommandableHttpService: services.NewCommandableHttpService("v1/eventlog"),
	}
	c.DependencyResolver.Put("controller", cref.NewDescriptor("pip-services-eventlog", "controller", "*", "*", "1.0"))
	return &c
}
