# EventLog Microservice

This is a system event logging microservice from Pip.Services library. 
It logs important system events like starts and stops of servers,
upgrades to a new version, fatal system errors or key business transactions.

The microservice currently supports the following deployment options:
* Deployment platforms: Standalone Process, Seneca Plugin
* External APIs: HTTP/REST
* Persistence: Memory, Flat Files, MongoDB

This microservice has no dependencies on other microservices.

<a name="links"></a> Quick Links:

* [Download Links](docs/Downloads.md)
* [Development Guide](docs/Development.md)
* [Configuration Guide](docs/Configuration.md)
* [Deployment Guide](docs/Deployment.md)
* Client SDKs
  - [Node.js SDK](https://github.com/pip-services/pip-clients-eventlog-node)
  - [Golang SDK](https://github.com/pip-services/pip-clients-eventlog-go)
  - [Dart SDK](https://github.com/pip-services/pip-clients-eventlog-dart)
* Communication Protocols
  - [HTTP Version 1](doc/HttpProtocolV1.md)

##  Contract

Logical contract of the microservice is presented below. For physical implementation (HTTP/REST, Thrift, Seneca, Lambda, etc.),
please, refer to documentation of the specific protocol.

```golang
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

// EventLogTypeV1
const Restart = "restart"
const Failure = "failure"
const Warning = "warning"
const Transaction = "transaction"
const Other = "other"

// EventLogSeverityV1
const Critical = 0
const Important = 500
const Informational = 1000

interface IEventLogV1 {
    getEvents(correlationId: string, filter: FilterParams, paging: PagingParams, 
        callback: (err: any, page: DataPage<SystemEventV1>) => void): void;
    
    logEvent(correlationId: string, event: SystemEventV1, 
        callback?: (err: any, event: SystemEventV1) => void): void;
}
```

## Download

Right now the only way to get the microservice is to check it out directly from github repository
```bash
git clone git@github.com:pip-services-infrastructure/pip-services-eventlog-go.git
```

Pip.Service team is working to implement packaging and make stable releases available for your 
as zip downloadable archieves.

## Run

Add **config.json** file to the root of the microservice folder and set configuration parameters.
As the starting point you can use example configuration from **config.example.yml** file. 

Example of microservice configuration
```yaml
- descriptor: "pip-services-container:container-info:default:default:1.0"
  name: "pip-services-eventlog"
  description: "EventLog microservice"

- descriptor: "pip-services-commons:logger:console:default:1.0"
  level: "trace"

- descriptor: "pip-services-eventlog:persistence:file:default:1.0"
  path: "./data/eventlog.json"

- descriptor: "pip-services-eventlog:controller:default:default:1.0"

- descriptor: "pip-services-eventlog:service:http:default:1.0"
  connection:
    protocol: "http"
    host: "0.0.0.0"
    port: 8080
```
 
For more information on the microservice configuration see [Configuration Guide](Configuration.md).

Start the microservice using the command:
```bash
go run ./bin/run.go
```

## Use

The easiest way to work with the microservice is to use client SDK. 
The complete list of available client SDKs for different languages is listed in the [Quick Links](#links)

If you use Golang then you should add dependency to the client SDK into **go.mod** file of your project
```golang
...
require (

    github.com/pip-services-infrastructure/pip-services-eventlog-go v1.0.0
    ....
)

```

Inside your code get the reference to the client SDK
```golang
import (
	clients1 "github.com/pip-services-infrastructure/pip-clients-eventlog-go/version1"
)

var client *clients1.EventLogHttpClientV1
```

Define client configuration parameters that match configuration of the microservice external API
```golang
// Client configuration
httpConfig := cconf.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.port", "3000",
		"connection.host", "localhost",
	)

	client = clients1.NewEventLogHttpClientV1()
	client.Configure(httpConfig)
```

Instantiate the client and open connection to the microservice
```golang

// Connect to the microservice
err := client.Open("")
 if (err) {
        panic("Connection to the microservice failed");
    }
defer client.Close("")
// Work with the microservice

```

Now the client is ready to perform operations
```golang
// Log system event
event1:=&clients1.SystemEventV1{
        Type: "restart",
        source: "server1",
        Message: "Restarted server",
    }

err := client.LogEvent(
    "",
    event1,
);
```

```golang
var now = time.Now();

// Get the list system events
page, err1 := client.getEvents(
    "",
    cdata.NewFilterParamsFromTuples(
        "from_time": new Date(now.getTime() - 24 * 3600 * 1000),
        "to_time": now,
        "source": "server1"
    ), cdata.NewEmptyPagingParams(),
);

```    

## Acknowledgements

This microservice was created and currently maintained by *Sergey Seroukhov*.

