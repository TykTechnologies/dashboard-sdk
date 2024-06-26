/*
NEW Tyk DASH API

Testing AssetsAPIService

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

func Test_dashboard_AssetsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AssetsAPIService AddAsset", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.AssetsAPI.AddAsset(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AssetsAPIService DeleteAsset", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var assetID string

		resp, httpRes, err := apiClient.AssetsAPI.DeleteAsset(context.Background(), assetID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AssetsAPIService GetAsset", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var assetID string

		resp, httpRes, err := apiClient.AssetsAPI.GetAsset(context.Background(), assetID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AssetsAPIService ListAssets", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.AssetsAPI.ListAssets(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AssetsAPIService UpdateAsset", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var assetID string

		resp, httpRes, err := apiClient.AssetsAPI.UpdateAsset(context.Background(), assetID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
