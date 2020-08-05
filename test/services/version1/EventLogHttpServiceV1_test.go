package test_services

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	data1 "github.com/pip-services-infrastructure/pip-services-eventlog-go/data/version1"
	logic "github.com/pip-services-infrastructure/pip-services-eventlog-go/logic"
	persist "github.com/pip-services-infrastructure/pip-services-eventlog-go/persistence"
	services1 "github.com/pip-services-infrastructure/pip-services-eventlog-go/services/version1"
	cconf "github.com/pip-services3-go/pip-services3-commons-go/config"
	cdata "github.com/pip-services3-go/pip-services3-commons-go/data"
	cerr "github.com/pip-services3-go/pip-services3-commons-go/errors"
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
	"github.com/stretchr/testify/assert"
)

var Event1 *data1.SystemEventV1 = &data1.SystemEventV1{
	Id:       "1",
	Time:     time.Now(),
	Source:   "test",
	Type:     data1.Restart,
	Severity: data1.Important,
	Message:  "test restart #1",
}

var Event2 *data1.SystemEventV1 = &data1.SystemEventV1{
	Id:       "2",
	Time:     time.Now(),
	Source:   "test",
	Type:     data1.Failure,
	Severity: data1.Critical,
	Message:  "test error",
}

func TestEventLogHttpServiceV1(t *testing.T) {
	var persistence *persist.EventLogMemoryPersistence
	var controller *logic.EventLogController
	var service *services1.EventLogHttpServiceV1
	var url string = "http://localhost:3000"

	persistence = persist.NewEventLogMemoryPersistence()
	persistence.Configure(cconf.NewEmptyConfigParams())

	controller = logic.NewEventLogController()
	controller.Configure(cconf.NewEmptyConfigParams())
	service = services1.NewEventLogHttpServiceV1()
	service.Configure(cconf.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.port", "3000",
		"connection.host", "localhost",
	))

	references := cref.NewReferencesFromTuples(
		cref.NewDescriptor("pip-services-eventlog", "persistence", "memory", "default", "1.0"), persistence,
		cref.NewDescriptor("pip-services-eventlog", "controller", "default", "default", "1.0"), controller,
		cref.NewDescriptor("pip-services-eventlog", "service", "http", "default", "1.0"), service,
	)

	controller.SetReferences(references)
	service.SetReferences(references)

	opnErr := persistence.Open("")
	if opnErr != nil {
		panic("Can't open persistence")
	}
	service.Open("")
	defer service.Close("")
	defer persistence.Close("")

	// Create the first event
	body := cdata.NewAnyValueMapFromTuples(
		"event", Event1,
	)
	err := invoke(url+"/v1/eventlog/log_event", body, nil)
	assert.Nil(t, err)

	// Create the second event
	body = cdata.NewAnyValueMapFromTuples(
		"event", Event2,
	)
	err = invoke(url+"/v1/eventlog/log_event", body, nil)
	assert.Nil(t, err)

	// Get all events
	body = cdata.NewAnyValueMapFromTuples(
		"filter", cdata.NewFilterParamsFromTuples("source", "test"),
	)
	page := data1.SystemEventV1DataPage{}
	err = invoke(url+"/v1/eventlog/get_events", body, &page)
	assert.Nil(t, err)
	assert.NotNil(t, page)
	assert.Len(t, page.Data, 2)
}

func invoke(url string, body *cdata.AnyValueMap, result interface{}) error {
	var bodyReader *bytes.Reader = nil
	if body != nil {
		jsonBody, _ := json.Marshal(body.Value())
		bodyReader = bytes.NewReader(jsonBody)
	}

	postResponse, postErr := http.Post(url, "application/json", bodyReader)

	if postErr != nil {
		return postErr
	}

	if postResponse.StatusCode == 204 {
		return nil
	}

	resBody, bodyErr := ioutil.ReadAll(postResponse.Body)
	if bodyErr != nil {
		return bodyErr
	}

	if postResponse.StatusCode >= 400 {
		appErr := cerr.ApplicationError{}
		json.Unmarshal(resBody, &appErr)
		return &appErr
	}

	if result == nil {
		return nil
	}

	jsonErr := json.Unmarshal(resBody, result)
	return jsonErr
}
