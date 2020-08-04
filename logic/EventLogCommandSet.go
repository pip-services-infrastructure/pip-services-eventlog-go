package logic

import (
	"encoding/json"

	data1 "github.com/pip-services-infrastructure/pip-services-eventlog-go/data/version1"
	ccmd "github.com/pip-services3-go/pip-services3-commons-go/commands"
	cdata "github.com/pip-services3-go/pip-services3-commons-go/data"
	crun "github.com/pip-services3-go/pip-services3-commons-go/run"
	cvalid "github.com/pip-services3-go/pip-services3-commons-go/validate"
)

type EventLogCommandSet struct {
	ccmd.CommandSet
	controller IEventLogController
}

func NewEventLogCommandSet(controller IEventLogController) *EventLogCommandSet {
	c := EventLogCommandSet{}
	c.CommandSet = *ccmd.NewCommandSet()
	c.controller = controller
	c.AddCommand(c.makeGetEventsCommand())
	c.AddCommand(c.makeLogEventCommand())
	return &c
}

func (c *EventLogCommandSet) makeGetEventsCommand() ccmd.ICommand {
	return ccmd.NewCommand(
		"get_events",
		cvalid.NewObjectSchema().
			WithOptionalProperty("filter", cvalid.NewFilterParamsSchema()).
			WithOptionalProperty("paging", cvalid.NewPagingParamsSchema()),
		func(correlationId string, args *crun.Parameters) (result interface{}, err error) {
			filter := cdata.NewFilterParamsFromValue(args.Get("filter"))
			paging := cdata.NewPagingParamsFromValue(args.Get("paging"))
			result, err = c.controller.GetEvents(correlationId, filter, paging)
			return result, err
		})
}

func (c *EventLogCommandSet) makeLogEventCommand() ccmd.ICommand {
	return ccmd.NewCommand(
		"log_event",
		cvalid.NewObjectSchema().
			WithRequiredProperty("event", data1.NewSystemEventV1Schema()),
		func(correlationId string, args *crun.Parameters) (result interface{}, err error) {
			val, _ := json.Marshal(args.Get("event"))
			var event data1.SystemEventV1
			json.Unmarshal(val, &event)

			err = c.controller.LogEvent(correlationId, &event)
			return nil, err
		})
}
