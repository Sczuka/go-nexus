/*
Nexus Repository Manager REST API

Testing ProductLicensingAPIService

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

func Test_nexus_ProductLicensingAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ProductLicensingAPIService GetLicenseStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ProductLicensingAPI.GetLicenseStatus(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProductLicensingAPIService RemoveLicense", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ProductLicensingAPI.RemoveLicense(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProductLicensingAPIService SetLicense", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ProductLicensingAPI.SetLicense(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
