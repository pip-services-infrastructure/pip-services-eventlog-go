package logic

import (
	data1 "github.com/pip-services-infrastructure/pip-services-eventlog-go/data/version1"
	cdata "github.com/pip-services3-go/pip-services3-commons-go/data"
)

type IEventLogController interface {
	GetEvents(correlationId string, filter *cdata.FilterParams,
		paging *cdata.PagingParams) (page *data1.SystemEventV1DataPage, err error)

	LogEvent(correlationId string, event *data1.SystemEventV1) error
}
