package test_logic

import (
	"testing"
	"time"

	data1 "github.com/pip-services-infrastructure/pip-services-eventlog-go/data/version1"
	logic "github.com/pip-services-infrastructure/pip-services-eventlog-go/logic"
	persist "github.com/pip-services-infrastructure/pip-services-eventlog-go/persistence"
	cconf "github.com/pip-services3-go/pip-services3-commons-go/config"
	cdata "github.com/pip-services3-go/pip-services3-commons-go/data"
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
	"github.com/stretchr/testify/assert"
)

var Event1 *data1.SystemEventV1 = &data1.SystemEventV1{
	Id:       "1",
	Time:     time.Now(),
	Source:   "test",
	Type:     data1.Restart,
	Severity: data1.Important,
	Message:  "test restart #1",
}

var Event2 *data1.SystemEventV1 = &data1.SystemEventV1{
	Id:       "2",
	Time:     time.Now(),
	Source:   "test",
	Type:     data1.Failure,
	Severity: data1.Critical,
	Message:  "test error",
}

var persistence *persist.EventLogMemoryPersistence
var controller *logic.EventLogController

func TestEventLogController(t *testing.T) {
	persistence = persist.NewEventLogMemoryPersistence()
	persistence.Configure(cconf.NewEmptyConfigParams())

	controller = logic.NewEventLogController()
	controller.Configure(cconf.NewEmptyConfigParams())

	references := cref.NewReferencesFromTuples(
		cref.NewDescriptor("pip-services-eventlog", "persistence", "memory", "default", "1.0"), persistence,
		cref.NewDescriptor("pip-services-eventlog", "controller", "default", "default", "1.0"), controller,
	)

	controller.SetReferences(references)

	persistence.Open("")

	defer persistence.Close("")

	t.Run("EventLogController:CRUD Operations", CrudOperations)
}

func CrudOperations(t *testing.T) {
	// Create one event
	err := controller.LogEvent("", Event1)
	assert.Nil(t, err)

	// Create another event
	err1 := controller.LogEvent("", Event2)
	assert.Nil(t, err1)

	// Get all system events
	events, err2 := persistence.GetPageByFilter("",
		cdata.NewFilterParamsFromTuples("source", "test"), cdata.NewEmptyPagingParams())

	assert.Nil(t, err2)
	assert.NotNil(t, events)
	assert.Len(t, events.Data, 2)
}
