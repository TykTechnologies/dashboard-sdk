package main

import (
	"context"
	"log"

	"github.com/TykTechnologies/dashboard-sdk/pkg/dashboard"
)

var (
	BaseUrl           = ""
	xTykAuthorization = "084e8c34e79447cc7b72cd189e1513dc"
)

func main() {
	client := configDashboard()
	policies, h, err := client.PoliciesAPI.GetPolicies(context.Background()).Execute()
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(h.StatusCode)
	log.Println(policies)
}

func configDashboard() *dashboard.APIClient {
	dashConfig := dashboard.Configuration{
		DefaultHeader: map[string]string{
			xTykAuthorization: xTykAuthorization,
		},
		Debug: false,
		Servers: dashboard.ServerConfigurations{
			{
				URL: BaseUrl,
			},
		},
	}

	client := dashboard.NewAPIClient(&dashConfig)

	return client
}
