package test_persistence

import (
	"os"
	"testing"

	persist "github.com/pip-services-infrastructure/pip-services-eventlog-go/persistence"
	cconf "github.com/pip-services3-go/pip-services3-commons-go/config"
)

func TestEventLogMongoDbPersistence(t *testing.T) {
	var persistence *persist.EventLogMongoDbPersistence
	var fixture *EventLogPersistenceFixture

	mongoUri := os.Getenv("MONGO_SERVICE_URI")
	mongoHost := os.Getenv("MONGO_SERVICE_HOST")

	if mongoHost == "" {
		mongoHost = "localhost"
	}

	mongoPort := os.Getenv("MONGO_SERVICE_PORT")
	if mongoPort == "" {
		mongoPort = "27017"
	}

	mongoDatabase := os.Getenv("MONGO_SERVICE_DB")
	if mongoDatabase == "" {
		mongoDatabase = "test"
	}

	// Exit if mongo connection is not set
	if mongoUri == "" && mongoHost == "" {
		return
	}

	persistence = persist.NewEventLogMongoDbPersistence()
	persistence.Configure(cconf.NewConfigParamsFromTuples(
		"connection.uri", mongoUri,
		"connection.host", mongoHost,
		"connection.port", mongoPort,
		"connection.database", mongoDatabase,
	))

	fixture = NewEventLogPersistenceFixture(persistence)

	opnErr := persistence.Open("")
	if opnErr == nil {
		persistence.Clear("")
	}

	defer persistence.Close("")

	t.Run("EventLogMongoDbPersistence:CRUD Operations", fixture.TestCrudOperations)
	// persistence.Clear("")
	// t.Run("EventLogMongoDbPersistence:Get with Filters", fixture.TestGetWithFilters)

}
