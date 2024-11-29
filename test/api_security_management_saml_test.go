/*
Nexus Repository Manager REST API

Testing SecurityManagementSAMLAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package nexus

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/sczuka/go-nexus"
)

func Test_nexus_SecurityManagementSAMLAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SecurityManagementSAMLAPIService DeleteSamlConfiguration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.SecurityManagementSAMLAPI.DeleteSamlConfiguration(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SecurityManagementSAMLAPIService GetMetadata", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.SecurityManagementSAMLAPI.GetMetadata(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SecurityManagementSAMLAPIService GetPublicCertificateInPemFormat", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.SecurityManagementSAMLAPI.GetPublicCertificateInPemFormat(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SecurityManagementSAMLAPIService GetSamlConfiguration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.SecurityManagementSAMLAPI.GetSamlConfiguration(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SecurityManagementSAMLAPIService PutSamlConfiguration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.SecurityManagementSAMLAPI.PutSamlConfiguration(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
