/*
NEW Tyk DASH API

Testing WebhooksAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package dashboard

import (
	"context"
	openapiclient "github.com/TykTechnologies/dashboard-sdk"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_dashboard_WebhooksAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test WebhooksAPIService CreateWebhook", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.WebhooksAPI.CreateWebhook(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WebhooksAPIService DeleteWebhook", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var hookId string

		resp, httpRes, err := apiClient.WebhooksAPI.DeleteWebhook(context.Background(), hookId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WebhooksAPIService GetWebhookDetail", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var hookId string

		resp, httpRes, err := apiClient.WebhooksAPI.GetWebhookDetail(context.Background(), hookId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WebhooksAPIService GetWebhookList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.WebhooksAPI.GetWebhookList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WebhooksAPIService UpdateWebhook", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var hookId string

		resp, httpRes, err := apiClient.WebhooksAPI.UpdateWebhook(context.Background(), hookId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
