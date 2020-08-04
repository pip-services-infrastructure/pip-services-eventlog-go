package persistence

import (
	data1 "github.com/pip-services-infrastructure/pip-services-eventlog-go/data/version1"
	cdata "github.com/pip-services3-go/pip-services3-commons-go/data"
)

type IEventLogPersistence interface {
	GetPageByFilter(correlationId string, filter *cdata.FilterParams,
		paging *cdata.PagingParams) (page *data1.SystemEventV1DataPage, err error)

	GetOneById(correlationId string, id string) (res *data1.SystemEventV1, err error)

	Create(correlationId string, item *data1.SystemEventV1) (res *data1.SystemEventV1, err error)

	Update(correlationId string, item *data1.SystemEventV1) (res *data1.SystemEventV1, err error)

	DeleteById(correlationId string, id string) (item *data1.SystemEventV1, err error)
}
