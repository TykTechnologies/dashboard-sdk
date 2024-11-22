package main

import (
	"context"
	"log"
	"os"

	"github.com/TykTechnologies/dashboard-sdk/examples/assets"
	"github.com/TykTechnologies/dashboard-sdk/examples/oas"
	"github.com/TykTechnologies/dashboard-sdk/pkg/dashboard"
)

var BaseUrl = "http://localhost:3000"

func main() {
	ctx := context.Background()
	token := os.Getenv("TYK_DASHBOARD_SDK_TOKEN")
	if token == "" {
		log.Println("You have not provided a token")
		return
	}
	client := createDashBoardClient(token)
	apiResp, err := oas.CreateOASAPI(ctx, client)
	if err != nil {
		log.Fatal(err)
		return
	}
	_, err = oas.GetOASApiByID(ctx, client, apiResp.GetID())
	if err != nil {
		log.Fatal(err)
		return
	}
	createAssetsResponse, err := assets.CreateAssets(ctx, client)
	if err != nil {
		log.Fatal(err)
		return
	}
	_, err = assets.GetAssetsByID(ctx, client, createAssetsResponse.GetID())
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
