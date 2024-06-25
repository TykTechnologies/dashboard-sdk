/*
NEW Tyk DASH API

Testing OASAPIsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package dashboard

import (
	"context"
	openapiclient "github.com/TykTechnologies/dashboard-sdk/pkg/dashboard"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_dashboard_OASAPIsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OASAPIsAPIService CreateApiOAS", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.OASAPIsAPI.CreateApiOAS(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OASAPIsAPIService DeleteOASApi", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var apiId string

		resp, httpRes, err := apiClient.OASAPIsAPI.DeleteOASApi(context.Background(), apiId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OASAPIsAPIService DownloadApiOASPublic", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var apiId string

		resp, httpRes, err := apiClient.OASAPIsAPI.DownloadApiOASPublic(context.Background(), apiId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OASAPIsAPIService GetApiCategories", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var apiId string

		resp, httpRes, err := apiClient.OASAPIsAPI.GetApiCategories(context.Background(), apiId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OASAPIsAPIService GetOASAPIDetails", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var apiId string

		resp, httpRes, err := apiClient.OASAPIsAPI.GetOASAPIDetails(context.Background(), apiId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OASAPIsAPIService ImportOAS", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.OASAPIsAPI.ImportOAS(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OASAPIsAPIService ListOASApiVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var apiId string

		resp, httpRes, err := apiClient.OASAPIsAPI.ListOASApiVersions(context.Background(), apiId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OASAPIsAPIService PatchApiOAS", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var apiId string

		resp, httpRes, err := apiClient.OASAPIsAPI.PatchApiOAS(context.Background(), apiId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OASAPIsAPIService UpdateApiCategories", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var apiId string

		resp, httpRes, err := apiClient.OASAPIsAPI.UpdateApiCategories(context.Background(), apiId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OASAPIsAPIService UpdateApiOAS", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var apiId string

		resp, httpRes, err := apiClient.OASAPIsAPI.UpdateApiOAS(context.Background(), apiId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
