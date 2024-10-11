# Dashboard Sdks (Alpha)

- In an effort to reduce duplication of work we have generated the dashboard sdks that different teams can use instead
  of creating they own client .

- This is the dashboard sdk generated from the open Api specs.
- Report any issue you may encounter will using this Sdk as we plan to perfect them before making them public.


## How to install
To run the sdk please run 

`github.com/TykTechnologies/dashboard-sdk`

## Sample Of how to create and get OAS API
```go
func main() {
	data := `{"openapi":"3.0.3","info":{"title":"OAS Sample","description":"This is a sample OAS.","version":"1.0.0"},"servers":[{"url":"https://localhost:8080"}],"security":[{"bearerAuth":[]}],"paths":{"/api/sample/users":{"get":{"tags":["users"],"summary":"Get users","operationId":"getUsers","responses":{"200":{"description":"fetched users","content":{"application/json":{"schema":{"type":"array","items":{"type":"object","properties":{"name":{"type":"string"}}}}}}}}}}},"components":{"securitySchemes":{"bearerAuth":{"type":"http","scheme":"bearer","description":"The API Access Credentials"}}},"x-tyk-api-gateway":{"info":{"name":"user","state":{"active":true}},"upstream":{"url":"https://localhost:8080"},"server":{"listenPath":{"value":"/user-test/","strip":true}}}}`
	var schema dashboard.OAS
	err := json.Unmarshal([]byte(data), &schema)
	if err != nil {
		log.Fatal(err)
	}
	client := configDashboard()
	var dt dashboard.CreateApiOASRequest
	err = dt.UnmarshalJSON([]byte(data))
	if err != nil {
		log.Fatal(err)
	}
	execute, h, err := client.OASAPIsAPI.CreateApiOAS(context.Background()).CreateApiOASRequest(dt).Execute()
	if err != nil {
		log.Fatal(err)
	}
	if h.StatusCode != 200 {
		log.Fatal(errors.New("test here"))
	}
	log.Println(*execute.ID)
	api, resp, err := client.OASAPIsAPI.GetOASAPIDetails(context.Background(), execute.GetID()).Execute()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(resp.StatusCode)
	marshalJSON, err := api.MarshalJSON()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(marshalJSON))
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
	dashConfig.AddDefaultHeader("authorization", "Bearer "+secret)

	client := dashboard.NewAPIClient(&dashConfig)

	return client
}


```

## Documentation

For documentation [check here.](https://github.com/TykTechnologies/dashboard-sdk/blob/main/pkg/dashboard/README.md)
