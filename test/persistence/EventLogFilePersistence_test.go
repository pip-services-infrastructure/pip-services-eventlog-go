package test_persistence

import (
	"testing"

	persist "github.com/pip-services-infrastructure/pip-services-eventlog-go/persistence"
	cconf "github.com/pip-services3-go/pip-services3-commons-go/config"
)

func TestEventLogFilePersistence(t *testing.T) {
	var persistence *persist.EventLogFilePersistence
	var fixture *EventLogPersistenceFixture

	persistence = persist.NewEventLogFilePersistence("../../temp/eventlog.test.json")
	persistence.Configure(cconf.NewEmptyConfigParams())
	fixture = NewEventLogPersistenceFixture(persistence)

	opnErr := persistence.Open("")
	if opnErr == nil {
		persistence.Clear("")
	}

	defer persistence.Close("")

	t.Run("EventLogFilePersistence:CRUD Operations", fixture.TestCrudOperations)
	// persistence.Clear("")
	// t.Run("EventLogFilePersistence:Get with Filters", fixture.TestGetWithFilters)
}
