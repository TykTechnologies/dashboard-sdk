package oas

import (
	"context"
	"log"

	"github.com/TykTechnologies/dashboard-sdk/examples/assets"
	"github.com/TykTechnologies/dashboard-sdk/pkg/dashboard"
)

var sampleOAS = `{"openapi":"3.0.3","info":{"title":"OAS Sample","description":"This is a sample OAS.","version":"1.0.0"},"servers":[{"url":"https://localhost:8080"}],"security":[{"bearerAuth":[]}],"paths":{"/api/sample/users":{"get":{"tags":["users"],"summary":"Get users","operationId":"getUsers","responses":{"200":{"description":"fetched users","content":{"application/json":{"schema":{"type":"array","items":{"type":"object","properties":{"name":{"type":"string"}}}}}}}}}}},"components":{"securitySchemes":{"bearerAuth":{"type":"http","scheme":"bearer","description":"The API Access Credentials"}}},"x-tyk-api-gateway":{"info":{"name":"user","state":{"active":true}},"upstream":{"url":"https://localhost:8080"},"server":{"listenPath":{"value":"/user-test/","strip":true}}}}`

func CreateOASAPI(ctx context.Context, client *dashboard.APIClient) (*dashboard.ApiResponse, error) {
	var oasData dashboard.CreateApiOASRequest
	err := oasData.UnmarshalJSON([]byte(sampleOAS))
	if err != nil {
		return nil, err
	}
	apiResponse, resp, err := client.OASAPIsAPI.CreateApiOAS(ctx).CreateApiOASRequest(oasData).Execute()
	if err != nil {
		assets.ExtractErrMessage(resp)
		return nil, err
	}
	log.Printf("OAS API created successfully with ID: %s", apiResponse.GetID())
	log.Println("-------------------------------------------------------------")
	return apiResponse, nil
}

func GetOASApiByID(ctx context.Context, client *dashboard.APIClient, apiID string) (*dashboard.CreateApiOASRequest, error) {
	apiResponse, resp, err := client.OASAPIsAPI.GetOASAPIDetails(ctx, apiID).Execute()
	if err != nil {
		assets.ExtractErrMessage(resp)
		return nil, err
	}
	log.Printf("OAS API fetched successfully with ID: %s and Name: %s", apiResponse.XTykApiGateway.Info.GetId(), apiResponse.XTykApiGateway.Info.GetName())
	log.Println("-------------------------------------------------------------")
	return apiResponse, nil
}
