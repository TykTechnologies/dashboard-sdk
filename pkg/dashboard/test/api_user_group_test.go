/*
NEW Tyk DASH API

Testing UserGroupAPIService

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

func Test_dashboard_UserGroupAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test UserGroupAPIService CreateUserGroup", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.UserGroupAPI.CreateUserGroup(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserGroupAPIService DeleteUserGroup", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var groupId string

		resp, httpRes, err := apiClient.UserGroupAPI.DeleteUserGroup(context.Background(), groupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserGroupAPIService GetUserGroup", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var groupId string

		resp, httpRes, err := apiClient.UserGroupAPI.GetUserGroup(context.Background(), groupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserGroupAPIService ListUserGroups", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.UserGroupAPI.ListUserGroups(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserGroupAPIService UpdateUserGroup", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var groupId string

		resp, httpRes, err := apiClient.UserGroupAPI.UpdateUserGroup(context.Background(), groupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
