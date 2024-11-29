# nexus\ManageSonatypeRepositoryFirewallConfigurationAPI

All URIs are relative to *http://localhost/service/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DisableIq**](ManageSonatypeRepositoryFirewallConfigurationAPI.md#DisableIq) | **Post** /v1/iq/disable | Disable Sonatype Repository Firewall
[**EnableIq**](ManageSonatypeRepositoryFirewallConfigurationAPI.md#EnableIq) | **Post** /v1/iq/enable | Enable Sonatype Repository Firewall
[**GetAllAuditStatus**](ManageSonatypeRepositoryFirewallConfigurationAPI.md#GetAllAuditStatus) | **Get** /v1/iq/audit | List repositories audit statuses.
[**GetAuditStatus**](ManageSonatypeRepositoryFirewallConfigurationAPI.md#GetAuditStatus) | **Get** /v1/iq/audit/{repositoryName} | Get audit status for the repository
[**GetConfiguration**](ManageSonatypeRepositoryFirewallConfigurationAPI.md#GetConfiguration) | **Get** /v1/iq | Get Sonatype Repository Firewall configuration
[**ManageAudit**](ManageSonatypeRepositoryFirewallConfigurationAPI.md#ManageAudit) | **Put** /v1/iq/audit | Manage audit
[**UpdateConfiguration**](ManageSonatypeRepositoryFirewallConfigurationAPI.md#UpdateConfiguration) | **Put** /v1/iq | Update Sonatype Repository Firewall configuration
[**VerifyConnection**](ManageSonatypeRepositoryFirewallConfigurationAPI.md#VerifyConnection) | **Post** /v1/iq/verify-connection | Verify Sonatype Repository Firewall connection



## DisableIq

> DisableIq(ctx).Execute()

Disable Sonatype Repository Firewall

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sczuka/go-nexus"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ManageSonatypeRepositoryFirewallConfigurationAPI.DisableIq(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManageSonatypeRepositoryFirewallConfigurationAPI.DisableIq``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDisableIqRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableIq

> EnableIq(ctx).Execute()

Enable Sonatype Repository Firewall

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sczuka/go-nexus"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ManageSonatypeRepositoryFirewallConfigurationAPI.EnableIq(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManageSonatypeRepositoryFirewallConfigurationAPI.EnableIq``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiEnableIqRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllAuditStatus

> []IqAuditXo GetAllAuditStatus(ctx).Execute()

List repositories audit statuses.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sczuka/go-nexus"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManageSonatypeRepositoryFirewallConfigurationAPI.GetAllAuditStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManageSonatypeRepositoryFirewallConfigurationAPI.GetAllAuditStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllAuditStatus`: []IqAuditXo
	fmt.Fprintf(os.Stdout, "Response from `ManageSonatypeRepositoryFirewallConfigurationAPI.GetAllAuditStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllAuditStatusRequest struct via the builder pattern


### Return type

[**[]IqAuditXo**](IqAuditXo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuditStatus

> GetAuditStatus(ctx, repositoryName).Execute()

Get audit status for the repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sczuka/go-nexus"
)

func main() {
	repositoryName := "repositoryName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ManageSonatypeRepositoryFirewallConfigurationAPI.GetAuditStatus(context.Background(), repositoryName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManageSonatypeRepositoryFirewallConfigurationAPI.GetAuditStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuditStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfiguration

> GetConfiguration(ctx).Execute()

Get Sonatype Repository Firewall configuration

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sczuka/go-nexus"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ManageSonatypeRepositoryFirewallConfigurationAPI.GetConfiguration(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManageSonatypeRepositoryFirewallConfigurationAPI.GetConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigurationRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManageAudit

> ManageAudit(ctx).Body(body).Execute()

Manage audit

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sczuka/go-nexus"
)

func main() {
	body := *openapiclient.NewIqAuditXo("RepositoryName_example", false) // IqAuditXo |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ManageSonatypeRepositoryFirewallConfigurationAPI.ManageAudit(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManageSonatypeRepositoryFirewallConfigurationAPI.ManageAudit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiManageAuditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**IqAuditXo**](IqAuditXo.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConfiguration

> UpdateConfiguration(ctx).Body(body).Execute()

Update Sonatype Repository Firewall configuration

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sczuka/go-nexus"
)

func main() {
	body := *openapiclient.NewIqConnectionXo("AuthenticationType_example") // IqConnectionXo |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ManageSonatypeRepositoryFirewallConfigurationAPI.UpdateConfiguration(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManageSonatypeRepositoryFirewallConfigurationAPI.UpdateConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**IqConnectionXo**](IqConnectionXo.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyConnection

> VerifyConnection(ctx).Execute()

Verify Sonatype Repository Firewall connection

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sczuka/go-nexus"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ManageSonatypeRepositoryFirewallConfigurationAPI.VerifyConnection(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManageSonatypeRepositoryFirewallConfigurationAPI.VerifyConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyConnectionRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
