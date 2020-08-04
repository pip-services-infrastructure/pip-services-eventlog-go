package data

import (
	cdata "github.com/pip-services3-go/pip-services3-commons-go/data"
	"time"
)

type SystemEventV1 struct {
	Id            string               `json:"id" bson:"_id"`
	Time          time.Time            `json:"time" bson:"time"`
	CorrelationId string               `json:"correlation_id" bson:"correlation_id"`
	Source        string               `json:"source" bson:"source"`
	Type          string               `json:"type" bson:"type"`
	Severity      int64                `json:"severity" bson:"severity"`
	Message       string               `json:"message" bson:"message"`
	Details       cdata.StringValueMap `json:"details" bson:"details"`
}
