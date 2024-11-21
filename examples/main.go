package main

import (
	"log"
	"os"

	"github.com/TykTechnologies/dashboard-sdk/examples/oas"
	"github.com/TykTechnologies/dashboard-sdk/pkg/dashboard"
)

var BaseUrl = "http://localhost:3000"

func main() {
	token := os.Getenv("TYK_DASHBOARD_SDK_TOKEN")
	client := createDashBoardClient(token)
	apiResp, err := oas.CreateOASAPI(client)
	if err != nil {
		log.Fatal(err)
		return
	}
	_, err = oas.GetOASApiByID(client, apiResp.GetID())
	if err != nil {
		log.Fatal(err)
		return
	}
}

func createDashBoardClient(token string) *dashboard.APIClient {
	dashConfig := dashboard.Configuration{
		DefaultHeader: map[string]string{},
		Debug:         false,
		Servers: dashboard.ServerConfigurations{
			{
				URL: BaseUrl,
			},
		},
	}
	dashConfig.AddDefaultHeader("authorization", "Bearer "+token)

	client := dashboard.NewAPIClient(&dashConfig)

	return client
}
