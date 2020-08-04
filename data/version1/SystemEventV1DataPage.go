package data

type SystemEventV1DataPage struct {
	Total *int64           `json:"total" bson:"total"`
	Data  []*SystemEventV1 `json:"data" bson:"data"`
}

func NewEmptySystemEventV1DataPage() *SystemEventV1DataPage {
	return &SystemEventV1DataPage{}
}

func NewSystemEventV1DataPage(total *int64, data []*SystemEventV1) *SystemEventV1DataPage {
	return &SystemEventV1DataPage{Total: total, Data: data}
}
