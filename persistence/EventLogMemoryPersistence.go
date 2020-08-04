package persistence

import (
	"reflect"
	"strings"

	data1 "github.com/pip-services-infrastructure/pip-services-eventlog-go/data/version1"
	cdata "github.com/pip-services3-go/pip-services3-commons-go/data"
	cpersist "github.com/pip-services3-go/pip-services3-data-go/persistence"
)

type EventLogMemoryPersistence struct {
	cpersist.IdentifiableMemoryPersistence
}

func NewEventLogMemoryPersistence() *EventLogMemoryPersistence {
	proto := reflect.TypeOf(&data1.SystemEventV1{})
	c := EventLogMemoryPersistence{}
	c.IdentifiableMemoryPersistence = *cpersist.NewIdentifiableMemoryPersistence(proto)
	return &c
}

func (c *EventLogMemoryPersistence) matchString(value string, search string) bool {
	if value == "" && search == "" {
		return true
	}
	if value == "" || search == "" {
		return false
	}
	return strings.Index(strings.ToLower(value), search) >= 0
}

func (c *EventLogMemoryPersistence) matchSearch(item *data1.SystemEventV1, search string) bool {
	search = strings.ToLower(search)
	if c.matchString(item.Source, search) {
		return true
	}
	if c.matchString(item.Message, search) {
		return true
	}
	if c.matchString(item.Type, search) {
		return true
	}
	return false
}

func (c *EventLogMemoryPersistence) composeFilter(filter *cdata.FilterParams) func(item interface{}) bool {
	if filter == nil {
		filter = cdata.NewEmptyFilterParams()
	}

	search := filter.GetAsString("search")
	id := filter.GetAsString("id")
	correlationId := filter.GetAsString("correlation_id")
	source := filter.GetAsString("source")
	typ := filter.GetAsString("type")
	minSeverity := filter.GetAsNullableLong("min_severity")
	fromTime := filter.GetAsNullableDateTime("from_time")
	toTime := filter.GetAsNullableDateTime("to_time")

	return func(data interface{}) bool {
		item, ok := data.(data1.SystemEventV1)
		if !ok {
			return false
		}

		if search != "" && !c.matchSearch(&item, search) {
			return false
		}
		if id != "" && id != item.Id {
			return false
		}
		if correlationId != "" && correlationId != item.CorrelationId {
			return false
		}
		if source != "" && source != item.Source {
			return false
		}
		if typ != "" && typ != item.Type {
			return false
		}
		if minSeverity != nil && item.Severity < *minSeverity {
			return false
		}
		if fromTime != nil && item.Time.Before(*fromTime) {
			return false
		}
		if toTime != nil && (item.Time.Equal(*toTime) || item.Time.After(*toTime)) {
			return false
		}
		return true
	}
}

func (c *EventLogMemoryPersistence) GetPageByFilter(correlationId string, filter *cdata.FilterParams,
	paging *cdata.PagingParams) (*data1.SystemEventV1DataPage, error) {
	p, err := c.IdentifiableMemoryPersistence.GetPageByFilter(correlationId, c.composeFilter(filter), paging, nil, nil)

	if p == nil || err != nil {
		return nil, err
	}

	// Convert to SystemEventV1DataPage
	d := make([]*data1.SystemEventV1, len(p.Data))
	for i, v := range p.Data {
		d[i] = v.(*data1.SystemEventV1)
	}

	page := data1.NewSystemEventV1DataPage(p.Total, d)
	return page, nil
}

func (c *EventLogMemoryPersistence) GetOneById(correlationId string, id string) (*data1.SystemEventV1, error) {
	value, err := c.IdentifiableMemoryPersistence.GetOneById(correlationId, id)

	if value == nil || err != nil {
		return nil, err
	}

	// Convert to SystemEventV1
	result, _ := value.(*data1.SystemEventV1)
	return result, nil
}

func (c *EventLogMemoryPersistence) Create(correlationId string, item *data1.SystemEventV1) (*data1.SystemEventV1, error) {
	value, err := c.IdentifiableMemoryPersistence.Create(correlationId, item)

	if value == nil || err != nil {
		return nil, err
	}

	// Convert to SystemEventV1
	result, _ := value.(*data1.SystemEventV1)
	return result, nil
}

func (c *EventLogMemoryPersistence) Update(correlationId string, item *data1.SystemEventV1) (*data1.SystemEventV1, error) {
	value, err := c.IdentifiableMemoryPersistence.Update(correlationId, item)

	if value == nil || err != nil {
		return nil, err
	}

	// Convert to SystemEventV1
	result, _ := value.(*data1.SystemEventV1)
	return result, nil
}

func (c *EventLogMemoryPersistence) DeleteById(correlationId string, id string) (*data1.SystemEventV1, error) {
	value, err := c.IdentifiableMemoryPersistence.DeleteById(correlationId, id)

	if value == nil || err != nil {
		return nil, err
	}

	// Convert to SystemEventV1
	result, _ := value.(*data1.SystemEventV1)
	return result, nil
}
