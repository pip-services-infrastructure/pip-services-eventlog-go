package persistence

import (
	cconf "github.com/pip-services3-go/pip-services3-commons-go/config"
	cpersist "github.com/pip-services3-go/pip-services3-data-go/persistence"
)

type EventLogFilePersistence struct {
	EventLogMemoryPersistence
	persister *cpersist.JsonFilePersister
}

func NewEventLogFilePersistence(path string) *EventLogFilePersistence {
	c := EventLogFilePersistence{}
	c.EventLogMemoryPersistence = *NewEventLogMemoryPersistence()
	c.persister = cpersist.NewJsonFilePersister(c.Prototype, path)
	c.Loader = c.persister
	c.Saver = c.persister
	return &c
}

func (c *EventLogFilePersistence) Configure(config *cconf.ConfigParams) {
	c.EventLogMemoryPersistence.Configure(config)
	c.persister.Configure(config)
}
