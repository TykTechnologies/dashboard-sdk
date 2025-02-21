# Dashboard Sdks (Alpha) 

- In an effort to reduce duplication of work we have generated the dashboard sdks that different teams can use instead
  of creating they own client .

- This is the dashboard sdk generated from the open Api specs.
- Report any issue you may encounter while using this Sdk .

## Versioning

Since the SDK is still experimental, we adopt a versioning scheme that reflects both its developmental status and the specific dashboard version used to generate it. For example, a tag like `v0.0.1-dashboard.5.7.2` indicates that:

- **v0.0.1** – The SDK is in an early, experimental stage and may undergo breaking changes.
- **-dashboard.5.7.2** – The SDK was generated from Dashboard version 5.7.2.

This versioning strategy ensures clear traceability, allowing users to easily identify which Gateway version the SDK is compatible with.


## How to install

### To install the SDK generated for a specific dashboard version (e.g., 5.7.2), run:

  `go get github.com/TykTechnologies/dashboard-sdk@v0.0.1-dashboard.5.7.2`

## Example of creating an OAS API using SDKs.

```go
package main

import (
	"context"
	"log"

	"github.com/TykTechnologies/dashboard-sdk/pkg/dashboard"
)

var (
	BaseUrl = "http://localhost:3000"
)
var sampleOAS = `{"openapi":"3.0.3","info":{"title":"OAS Sample","description":"This is a sample OAS.","version":"1.0.0"},"servers":[{"url":"https://localhost:8080"}],"security":[{"bearerAuth":[]}],"paths":{"/api/sample/users":{"get":{"tags":["users"],"summary":"Get users","operationId":"getUsers","responses":{"200":{"description":"fetched users","content":{"application/json":{"schema":{"type":"array","items":{"type":"object","properties":{"name":{"type":"string"}}}}}}}}}}},"components":{"securitySchemes":{"bearerAuth":{"type":"http","scheme":"bearer","description":"The API Access Credentials"}}},"x-tyk-api-gateway":{"info":{"name":"user","state":{"active":true}},"upstream":{"url":"https://localhost:8080"},"server":{"listenPath":{"value":"/user-test/","strip":true}}}}`

func main() {
	client := configDashboard()

	apiResp, err := createOASAPI(client)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println(apiResp.GetID())
}

func createOASAPI(client *dashboard.APIClient) (*dashboard.ApiResponse, error) {
	var oasData dashboard.CreateApiOASRequest
	err := oasData.UnmarshalJSON([]byte(sampleOAS))
	if err != nil {
		return nil, err
	}
	apiResponse, rep, err := client.OASAPIsAPI.CreateApiOAS(context.Background()).CreateApiOASRequest(oasData).Execute()
	if err != nil {
		log.Println(rep.StatusCode)
		log.Println(rep.Body)
		return nil, err
	}
	return apiResponse, nil
}

func configDashboard() *dashboard.APIClient {
	dashConfig := dashboard.Configuration{
		DefaultHeader: map[string]string{},
		Debug:         false,
		Servers: dashboard.ServerConfigurations{
			{
				URL: BaseUrl,
			},
		},
	}
	dashConfig.AddDefaultHeader("authorization", "Bearer "+<YOUR SECRET HERE>)

	client := dashboard.NewAPIClient(&dashConfig)

	return client
}
```

## Documentation

For documentation please [look here](https://github.com/TykTechnologies/dashboard-sdk/blob/main/pkg/dashboard/README.md#documentation-for-api-endpoints).

