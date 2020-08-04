package test_persistence

import (
	"testing"

	persist "github.com/pip-services-infrastructure/pip-services-eventlog-go/persistence"
	cconf "github.com/pip-services3-go/pip-services3-commons-go/config"
)

func TestEventLogMemoryPersistence(t *testing.T) {
	var persistence *persist.EventLogMemoryPersistence
	var fixture *EventLogPersistenceFixture

	persistence = persist.NewEventLogMemoryPersistence()
	persistence.Configure(cconf.NewEmptyConfigParams())
	fixture = NewEventLogPersistenceFixture(persistence)

	persistence.Open("")

	defer persistence.Close("")

	t.Run("EventLogMemoryPersistence:CRUD Operations", fixture.TestCrudOperations)
	// persistence.Clear("")
	// t.Run("EventLogMemoryPersistence:Get with Filters", fixture.TestGetWithFilters)
}
