package data

import (
	cconv "github.com/pip-services3-go/pip-services3-commons-go/convert"
	cvalid "github.com/pip-services3-go/pip-services3-commons-go/validate"
)

type SystemEventV1Schema struct {
	cvalid.ObjectSchema
}

func NewSystemEventV1Schema() *SystemEventV1Schema {
	c := SystemEventV1Schema{}
	c.ObjectSchema = *cvalid.NewObjectSchema()

	c.WithOptionalProperty("id", cconv.String)
	c.WithOptionalProperty("time", cconv.DateTime)
	c.WithOptionalProperty("correlation_id", cconv.String)
	c.WithOptionalProperty("source", cconv.String)
	c.WithRequiredProperty("type", cconv.String)
	c.WithRequiredProperty("severity", cconv.Long)
	c.WithOptionalProperty("message", cconv.String)
	c.WithOptionalProperty("details", cconv.Map)
	return &c
}
