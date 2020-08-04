package persistence

import (
	"reflect"

	data1 "github.com/pip-services-infrastructure/pip-services-eventlog-go/data/version1"
	cdata "github.com/pip-services3-go/pip-services3-commons-go/data"
	mpersist "github.com/pip-services3-go/pip-services3-mongodb-go/persistence"
	"go.mongodb.org/mongo-driver/bson"
)

type EventLogMongoDbPersistence struct {
	mpersist.IdentifiableMongoDbPersistence
}

func NewEventLogMongoDbPersistence() *EventLogMongoDbPersistence {
	proto := reflect.TypeOf(&data1.SystemEventV1{})
	c := EventLogMongoDbPersistence{}
	c.IdentifiableMongoDbPersistence = *mpersist.NewIdentifiableMongoDbPersistence(proto, "event_log")
	return &c
}

func (c *EventLogMongoDbPersistence) composeFilter(filter *cdata.FilterParams) interface{} {
	if filter == nil {
		filter = cdata.NewEmptyFilterParams()
	}

	criteria := make([]bson.M, 0, 0)

	// let search = filter.getAsNullableString('search');
	// if (search != null) {
	//     let searchRegex = new RegExp(search, "i");
	//     let searchCriteria = [];
	//     searchCriteria.push({ source: { $regex: searchRegex } });
	//     searchCriteria.push({ type: { $regex: searchRegex } });
	//     searchCriteria.push({ message: { $regex: searchRegex } });
	//     criteria.push({ $or: searchCriteria });
	// }

	id := filter.GetAsString("id")
	if id != "" {
		criteria = append(criteria, bson.M{"_id": id})
	}

	correlationId := filter.GetAsString("correlation_id")
	if correlationId != "" {
		criteria = append(criteria, bson.M{"correlation_id": correlationId})
	}

	source := filter.GetAsString("source")
	if source != "" {
		criteria = append(criteria, bson.M{"source": source})
	}

	typ := filter.GetAsString("type")
	if typ != "" {
		criteria = append(criteria, bson.M{"type": typ})
	}

	minSeverity := filter.GetAsNullableInteger("min_severity")
	if minSeverity != nil {
		criteria = append(criteria, bson.M{"severity": bson.M{"$gte": *minSeverity}})
	}

	fromTime := filter.GetAsNullableDateTime("from_time")
	if fromTime != nil {
		criteria = append(criteria, bson.M{"time": bson.M{"$gte": fromTime}})
	}

	toTime := filter.GetAsNullableDateTime("to_time")
	if toTime != nil {
		criteria = append(criteria, bson.M{"time": bson.M{"$lt": toTime}})
	}

	if len(criteria) > 0 {
		return bson.D{{"$and", criteria}}
	}
	return bson.M{}
}

func (c *EventLogMongoDbPersistence) GetPageByFilter(correlationId string, filter *cdata.FilterParams, paging *cdata.PagingParams) (*data1.SystemEventV1DataPage, error) {
	p, err := c.IdentifiableMongoDbPersistence.GetPageByFilter(correlationId, c.composeFilter(filter), paging, nil, nil)
	// Todo: Here we shall receive a reference instead of object
	if /*p == nil ||*/ err != nil {
		return nil, err
	}

	// Convert to EventLogV1DataPage
	d := make([]*data1.SystemEventV1, len(p.Data))
	for i, v := range p.Data {
		d[i] = v.(*data1.SystemEventV1)
	}
	page := data1.NewSystemEventV1DataPage(p.Total, d)
	return page, nil
}

func (c *EventLogMongoDbPersistence) GetOneById(correlationId string, id string) (*data1.SystemEventV1, error) {
	value, err := c.IdentifiableMongoDbPersistence.GetOneById(correlationId, id)

	if value == nil || err != nil {
		return nil, err
	}

	// Convert to SystemEventV1
	result, _ := value.(*data1.SystemEventV1)
	return result, nil
}

func (c *EventLogMongoDbPersistence) Create(correlationId string, item *data1.SystemEventV1) (*data1.SystemEventV1, error) {
	value, err := c.IdentifiableMongoDbPersistence.Create(correlationId, item)

	if value == nil || err != nil {
		return nil, err
	}

	// Convert to SystemEventV1
	result, _ := value.(*data1.SystemEventV1)
	return result, nil
}

func (c *EventLogMongoDbPersistence) Update(correlationId string, item *data1.SystemEventV1) (*data1.SystemEventV1, error) {
	value, err := c.IdentifiableMongoDbPersistence.Update(correlationId, item)

	if value == nil || err != nil {
		return nil, err
	}

	// Convert to SystemEventV1
	result, _ := value.(*data1.SystemEventV1)
	return result, nil
}

func (c *EventLogMongoDbPersistence) DeleteById(correlationId string, id string) (*data1.SystemEventV1, error) {
	value, err := c.IdentifiableMongoDbPersistence.DeleteById(correlationId, id)

	if value == nil || err != nil {
		return nil, err
	}

	// Convert to SystemEventV1
	result, _ := value.(*data1.SystemEventV1)
	return result, nil
}
