package test_persistence

import (
	"testing"
	"time"

	data1 "github.com/pip-services-infrastructure/pip-services-eventlog-go/data/version1"
	persist "github.com/pip-services-infrastructure/pip-services-eventlog-go/persistence"
	cdata "github.com/pip-services3-go/pip-services3-commons-go/data"
	"github.com/stretchr/testify/assert"
)

type EventLogPersistenceFixture struct {
	Event1      *data1.SystemEventV1
	Event2      *data1.SystemEventV1
	persistence persist.IEventLogPersistence
}

func NewEventLogPersistenceFixture(persistence persist.IEventLogPersistence) *EventLogPersistenceFixture {
	c := EventLogPersistenceFixture{}
	c.Event1 = &data1.SystemEventV1{
		Id:       "1",
		Time:     time.Now(),
		Source:   "test",
		Type:     data1.Restart,
		Severity: data1.Important,
		Message:  "test restart #1",
	}
	c.Event2 = &data1.SystemEventV1{
		Id:       "2",
		Time:     time.Now(),
		Source:   "test",
		Type:     data1.Failure,
		Severity: data1.Critical,
		Message:  "test error",
	}
	c.persistence = persistence
	return &c
}

func (c *EventLogPersistenceFixture) TestCrudOperations(t *testing.T) {
	// Create one event
	event1, err := c.persistence.Create("", c.Event1)
	assert.Nil(t, err)
	assert.NotNil(t, event1)
	assert.Equal(t, c.Event1.Id, event1.Id)
	assert.Equal(t, c.Event1.Time, event1.Time)
	assert.Equal(t, c.Event1.Type, event1.Type)
	assert.Equal(t, c.Event1.Severity, event1.Severity)
	assert.Equal(t, c.Event1.Message, event1.Message)

	// Create another event
	event2, err1 := c.persistence.Create("", c.Event2)
	assert.Nil(t, err1)
	assert.NotNil(t, event2)
	assert.Equal(t, c.Event2.Id, event2.Id)
	assert.Equal(t, c.Event2.Time, event2.Time)
	assert.Equal(t, c.Event2.Type, event2.Type)
	assert.Equal(t, c.Event2.Severity, event2.Severity)
	assert.Equal(t, c.Event2.Message, event2.Message)

	// Get all system events
	events, err2 := c.persistence.GetPageByFilter("",
		cdata.NewFilterParamsFromTuples("source", "test"), cdata.NewEmptyPagingParams())

	assert.Nil(t, err2)
	assert.NotNil(t, events)
	assert.Len(t, events.Data, 2)
}
