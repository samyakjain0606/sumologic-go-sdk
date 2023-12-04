/*
Sumo Logic API

Testing LogsDataForwardingManagementAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_openapi_LogsDataForwardingManagementAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test LogsDataForwardingManagementAPIService CreateDataForwardingBucket", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.LogsDataForwardingManagementAPI.CreateDataForwardingBucket(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LogsDataForwardingManagementAPIService CreateDataForwardingRule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.LogsDataForwardingManagementAPI.CreateDataForwardingRule(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LogsDataForwardingManagementAPIService DeleteDataForwardingBucket", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.LogsDataForwardingManagementAPI.DeleteDataForwardingBucket(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LogsDataForwardingManagementAPIService DeleteDataForwardingRule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var indexId string

		httpRes, err := apiClient.LogsDataForwardingManagementAPI.DeleteDataForwardingRule(context.Background(), indexId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LogsDataForwardingManagementAPIService GetDataForwardingBuckets", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.LogsDataForwardingManagementAPI.GetDataForwardingBuckets(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LogsDataForwardingManagementAPIService GetDataForwardingDestination", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.LogsDataForwardingManagementAPI.GetDataForwardingDestination(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LogsDataForwardingManagementAPIService GetDataForwardingRule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var indexId string

		resp, httpRes, err := apiClient.LogsDataForwardingManagementAPI.GetDataForwardingRule(context.Background(), indexId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LogsDataForwardingManagementAPIService GetRulesAndBuckets", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.LogsDataForwardingManagementAPI.GetRulesAndBuckets(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LogsDataForwardingManagementAPIService UpdateDataForwardingBucket", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.LogsDataForwardingManagementAPI.UpdateDataForwardingBucket(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LogsDataForwardingManagementAPIService UpdateDataForwardingRule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var indexId string

		resp, httpRes, err := apiClient.LogsDataForwardingManagementAPI.UpdateDataForwardingRule(context.Background(), indexId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
