package build

import (
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
	cbuild "github.com/pip-services3-go/pip-services3-components-go/build"
	logic "github.com/pip-services-infrastructure/pip-services-eventlog-go/logic"
	persist "github.com/pip-services-infrastructure/pip-services-eventlog-go/persistence"
	services1 "github.com/pip-services-infrastructure/pip-services-eventlog-go/services/version1"
)

type EventLogServiceFactory struct {
	cbuild.Factory
}

func NewEventLogServiceFactory() *EventLogServiceFactory {
	c := EventLogServiceFactory{}
	c.Factory = *cbuild.NewFactory()

	memoryPersistenceDescriptor := cref.NewDescriptor("pip-services-eventlog", "persistence", "memory", "*", "1.0")
	filePersistenceDescriptor := cref.NewDescriptor("pip-services-eventlog", "persistence", "file", "*", "1.0")
	mongoDbPersistenceDescriptor := cref.NewDescriptor("pip-services-eventlog", "persistence", "mongodb", "*", "1.0")
	controllerDescriptor := cref.NewDescriptor("pip-services-eventlog", "controller", "default", "*", "1.0")
	httpServiceV1Descriptor := cref.NewDescriptor("pip-services-eventlog", "service", "http", "*", "1.0")

	c.RegisterType(memoryPersistenceDescriptor, persist.NewEventLogMemoryPersistence)
	c.RegisterType(filePersistenceDescriptor, persist.NewEventLogFilePersistence)
	c.RegisterType(mongoDbPersistenceDescriptor, persist.NewEventLogMongoDbPersistence)
	c.RegisterType(controllerDescriptor, logic.NewEventLogController)
	c.RegisterType(httpServiceV1Descriptor, services1.NewEventLogHttpServiceV1)
	return &c
}
