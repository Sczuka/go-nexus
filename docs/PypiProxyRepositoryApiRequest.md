# PypiProxyRepositoryApiRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A unique identifier for this repository | 
**Online** | **bool** | Whether this repository accepts incoming requests | 
**Storage** | [**StorageAttributes**](StorageAttributes.md) |  | 
**Cleanup** | Pointer to [**CleanupPolicyAttributes**](CleanupPolicyAttributes.md) |  | [optional] 
**Proxy** | [**ProxyAttributes**](ProxyAttributes.md) |  | 
**NegativeCache** | [**NegativeCacheAttributes**](NegativeCacheAttributes.md) |  | 
**HttpClient** | [**HttpClientAttributes**](HttpClientAttributes.md) |  | 
**RoutingRule** | Pointer to **string** |  | [optional] 
**Replication** | Pointer to [**ReplicationAttributes**](ReplicationAttributes.md) |  | [optional] 
**Pypi** | Pointer to [**PyPiProxyAttributes**](PyPiProxyAttributes.md) |  | [optional] 

## Methods

### NewPypiProxyRepositoryApiRequest

`func NewPypiProxyRepositoryApiRequest(name string, online bool, storage StorageAttributes, proxy ProxyAttributes, negativeCache NegativeCacheAttributes, httpClient HttpClientAttributes, ) *PypiProxyRepositoryApiRequest`

NewPypiProxyRepositoryApiRequest instantiates a new PypiProxyRepositoryApiRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPypiProxyRepositoryApiRequestWithDefaults

`func NewPypiProxyRepositoryApiRequestWithDefaults() *PypiProxyRepositoryApiRequest`

NewPypiProxyRepositoryApiRequestWithDefaults instantiates a new PypiProxyRepositoryApiRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PypiProxyRepositoryApiRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PypiProxyRepositoryApiRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PypiProxyRepositoryApiRequest) SetName(v string)`

SetName sets Name field to given value.


### GetOnline

`func (o *PypiProxyRepositoryApiRequest) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *PypiProxyRepositoryApiRequest) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *PypiProxyRepositoryApiRequest) SetOnline(v bool)`

SetOnline sets Online field to given value.


### GetStorage

`func (o *PypiProxyRepositoryApiRequest) GetStorage() StorageAttributes`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *PypiProxyRepositoryApiRequest) GetStorageOk() (*StorageAttributes, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *PypiProxyRepositoryApiRequest) SetStorage(v StorageAttributes)`

SetStorage sets Storage field to given value.


### GetCleanup

`func (o *PypiProxyRepositoryApiRequest) GetCleanup() CleanupPolicyAttributes`

GetCleanup returns the Cleanup field if non-nil, zero value otherwise.

### GetCleanupOk

`func (o *PypiProxyRepositoryApiRequest) GetCleanupOk() (*CleanupPolicyAttributes, bool)`

GetCleanupOk returns a tuple with the Cleanup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanup

`func (o *PypiProxyRepositoryApiRequest) SetCleanup(v CleanupPolicyAttributes)`

SetCleanup sets Cleanup field to given value.

### HasCleanup

`func (o *PypiProxyRepositoryApiRequest) HasCleanup() bool`

HasCleanup returns a boolean if a field has been set.

### GetProxy

`func (o *PypiProxyRepositoryApiRequest) GetProxy() ProxyAttributes`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *PypiProxyRepositoryApiRequest) GetProxyOk() (*ProxyAttributes, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *PypiProxyRepositoryApiRequest) SetProxy(v ProxyAttributes)`

SetProxy sets Proxy field to given value.


### GetNegativeCache

`func (o *PypiProxyRepositoryApiRequest) GetNegativeCache() NegativeCacheAttributes`

GetNegativeCache returns the NegativeCache field if non-nil, zero value otherwise.

### GetNegativeCacheOk

`func (o *PypiProxyRepositoryApiRequest) GetNegativeCacheOk() (*NegativeCacheAttributes, bool)`

GetNegativeCacheOk returns a tuple with the NegativeCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegativeCache

`func (o *PypiProxyRepositoryApiRequest) SetNegativeCache(v NegativeCacheAttributes)`

SetNegativeCache sets NegativeCache field to given value.


### GetHttpClient

`func (o *PypiProxyRepositoryApiRequest) GetHttpClient() HttpClientAttributes`

GetHttpClient returns the HttpClient field if non-nil, zero value otherwise.

### GetHttpClientOk

`func (o *PypiProxyRepositoryApiRequest) GetHttpClientOk() (*HttpClientAttributes, bool)`

GetHttpClientOk returns a tuple with the HttpClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpClient

`func (o *PypiProxyRepositoryApiRequest) SetHttpClient(v HttpClientAttributes)`

SetHttpClient sets HttpClient field to given value.


### GetRoutingRule

`func (o *PypiProxyRepositoryApiRequest) GetRoutingRule() string`

GetRoutingRule returns the RoutingRule field if non-nil, zero value otherwise.

### GetRoutingRuleOk

`func (o *PypiProxyRepositoryApiRequest) GetRoutingRuleOk() (*string, bool)`

GetRoutingRuleOk returns a tuple with the RoutingRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingRule

`func (o *PypiProxyRepositoryApiRequest) SetRoutingRule(v string)`

SetRoutingRule sets RoutingRule field to given value.

### HasRoutingRule

`func (o *PypiProxyRepositoryApiRequest) HasRoutingRule() bool`

HasRoutingRule returns a boolean if a field has been set.

### GetReplication

`func (o *PypiProxyRepositoryApiRequest) GetReplication() ReplicationAttributes`

GetReplication returns the Replication field if non-nil, zero value otherwise.

### GetReplicationOk

`func (o *PypiProxyRepositoryApiRequest) GetReplicationOk() (*ReplicationAttributes, bool)`

GetReplicationOk returns a tuple with the Replication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplication

`func (o *PypiProxyRepositoryApiRequest) SetReplication(v ReplicationAttributes)`

SetReplication sets Replication field to given value.

### HasReplication

`func (o *PypiProxyRepositoryApiRequest) HasReplication() bool`

HasReplication returns a boolean if a field has been set.

### GetPypi

`func (o *PypiProxyRepositoryApiRequest) GetPypi() PyPiProxyAttributes`

GetPypi returns the Pypi field if non-nil, zero value otherwise.

### GetPypiOk

`func (o *PypiProxyRepositoryApiRequest) GetPypiOk() (*PyPiProxyAttributes, bool)`

GetPypiOk returns a tuple with the Pypi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPypi

`func (o *PypiProxyRepositoryApiRequest) SetPypi(v PyPiProxyAttributes)`

SetPypi sets Pypi field to given value.

### HasPypi

`func (o *PypiProxyRepositoryApiRequest) HasPypi() bool`

HasPypi returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


