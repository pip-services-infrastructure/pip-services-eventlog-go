package logic

import (
	data1 "github.com/pip-services-infrastructure/pip-services-eventlog-go/data/version1"
	persist "github.com/pip-services-infrastructure/pip-services-eventlog-go/persistence"
	ccomand "github.com/pip-services3-go/pip-services3-commons-go/commands"
	cconf "github.com/pip-services3-go/pip-services3-commons-go/config"
	cdata "github.com/pip-services3-go/pip-services3-commons-go/data"
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
)

type EventLogController struct {
	persistence persist.IEventLogPersistence
	commandSet  *EventLogCommandSet
}

func NewEventLogController() *EventLogController {
	c := EventLogController{}
	return &c
}

func (c *EventLogController) Configure(config *cconf.ConfigParams) {
	// Todo: Read configuration parameters here...
}

func (c *EventLogController) SetReferences(references cref.IReferences) {
	p, err := references.GetOneRequired(cref.NewDescriptor("pip-services-eventlog", "persistence", "*", "*", "1.0"))
	if p != nil && err == nil {
		c.persistence = p.(persist.IEventLogPersistence)
	}
}

func (c *EventLogController) GetCommandSet() *ccomand.CommandSet {
	if c.commandSet == nil {
		c.commandSet = NewEventLogCommandSet(c)
	}
	return &c.commandSet.CommandSet
}

func (c *EventLogController) GetEvents(correlationId string, filter *cdata.FilterParams,
	paging *cdata.PagingParams) (*data1.SystemEventV1DataPage, error) {
	return c.persistence.GetPageByFilter(correlationId, filter, paging)
}

func (c *EventLogController) LogEvent(correlationId string, event *data1.SystemEventV1) error {
	// event.severity = event.severity || EventLogSeverityV1.Informational;
	// event.time = event.time || new Date();
	_, err := c.persistence.Create(correlationId, event)
	return err
}
