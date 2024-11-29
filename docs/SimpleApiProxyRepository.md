# SimpleApiProxyRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A unique identifier for this repository | [optional] 
**Online** | **bool** | Whether this repository accepts incoming requests | 
**Storage** | [**StorageAttributes**](StorageAttributes.md) |  | 
**Cleanup** | Pointer to [**CleanupPolicyAttributes**](CleanupPolicyAttributes.md) |  | [optional] 
**Proxy** | [**ProxyAttributes**](ProxyAttributes.md) |  | 
**NegativeCache** | [**NegativeCacheAttributes**](NegativeCacheAttributes.md) |  | 
**HttpClient** | [**HttpClientAttributes**](HttpClientAttributes.md) |  | 
**RoutingRuleName** | Pointer to **string** | The name of the routing rule assigned to this repository | [optional] 
**Replication** | Pointer to [**ReplicationAttributes**](ReplicationAttributes.md) |  | [optional] 

## Methods

### NewSimpleApiProxyRepository

`func NewSimpleApiProxyRepository(online bool, storage StorageAttributes, proxy ProxyAttributes, negativeCache NegativeCacheAttributes, httpClient HttpClientAttributes, ) *SimpleApiProxyRepository`

NewSimpleApiProxyRepository instantiates a new SimpleApiProxyRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimpleApiProxyRepositoryWithDefaults

`func NewSimpleApiProxyRepositoryWithDefaults() *SimpleApiProxyRepository`

NewSimpleApiProxyRepositoryWithDefaults instantiates a new SimpleApiProxyRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SimpleApiProxyRepository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SimpleApiProxyRepository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SimpleApiProxyRepository) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SimpleApiProxyRepository) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOnline

`func (o *SimpleApiProxyRepository) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *SimpleApiProxyRepository) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *SimpleApiProxyRepository) SetOnline(v bool)`

SetOnline sets Online field to given value.


### GetStorage

`func (o *SimpleApiProxyRepository) GetStorage() StorageAttributes`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *SimpleApiProxyRepository) GetStorageOk() (*StorageAttributes, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *SimpleApiProxyRepository) SetStorage(v StorageAttributes)`

SetStorage sets Storage field to given value.


### GetCleanup

`func (o *SimpleApiProxyRepository) GetCleanup() CleanupPolicyAttributes`

GetCleanup returns the Cleanup field if non-nil, zero value otherwise.

### GetCleanupOk

`func (o *SimpleApiProxyRepository) GetCleanupOk() (*CleanupPolicyAttributes, bool)`

GetCleanupOk returns a tuple with the Cleanup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanup

`func (o *SimpleApiProxyRepository) SetCleanup(v CleanupPolicyAttributes)`

SetCleanup sets Cleanup field to given value.

### HasCleanup

`func (o *SimpleApiProxyRepository) HasCleanup() bool`

HasCleanup returns a boolean if a field has been set.

### GetProxy

`func (o *SimpleApiProxyRepository) GetProxy() ProxyAttributes`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *SimpleApiProxyRepository) GetProxyOk() (*ProxyAttributes, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *SimpleApiProxyRepository) SetProxy(v ProxyAttributes)`

SetProxy sets Proxy field to given value.


### GetNegativeCache

`func (o *SimpleApiProxyRepository) GetNegativeCache() NegativeCacheAttributes`

GetNegativeCache returns the NegativeCache field if non-nil, zero value otherwise.

### GetNegativeCacheOk

`func (o *SimpleApiProxyRepository) GetNegativeCacheOk() (*NegativeCacheAttributes, bool)`

GetNegativeCacheOk returns a tuple with the NegativeCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegativeCache

`func (o *SimpleApiProxyRepository) SetNegativeCache(v NegativeCacheAttributes)`

SetNegativeCache sets NegativeCache field to given value.


### GetHttpClient

`func (o *SimpleApiProxyRepository) GetHttpClient() HttpClientAttributes`

GetHttpClient returns the HttpClient field if non-nil, zero value otherwise.

### GetHttpClientOk

`func (o *SimpleApiProxyRepository) GetHttpClientOk() (*HttpClientAttributes, bool)`

GetHttpClientOk returns a tuple with the HttpClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpClient

`func (o *SimpleApiProxyRepository) SetHttpClient(v HttpClientAttributes)`

SetHttpClient sets HttpClient field to given value.


### GetRoutingRuleName

`func (o *SimpleApiProxyRepository) GetRoutingRuleName() string`

GetRoutingRuleName returns the RoutingRuleName field if non-nil, zero value otherwise.

### GetRoutingRuleNameOk

`func (o *SimpleApiProxyRepository) GetRoutingRuleNameOk() (*string, bool)`

GetRoutingRuleNameOk returns a tuple with the RoutingRuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingRuleName

`func (o *SimpleApiProxyRepository) SetRoutingRuleName(v string)`

SetRoutingRuleName sets RoutingRuleName field to given value.

### HasRoutingRuleName

`func (o *SimpleApiProxyRepository) HasRoutingRuleName() bool`

HasRoutingRuleName returns a boolean if a field has been set.

### GetReplication

`func (o *SimpleApiProxyRepository) GetReplication() ReplicationAttributes`

GetReplication returns the Replication field if non-nil, zero value otherwise.

### GetReplicationOk

`func (o *SimpleApiProxyRepository) GetReplicationOk() (*ReplicationAttributes, bool)`

GetReplicationOk returns a tuple with the Replication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplication

`func (o *SimpleApiProxyRepository) SetReplication(v ReplicationAttributes)`

SetReplication sets Replication field to given value.

### HasReplication

`func (o *SimpleApiProxyRepository) HasReplication() bool`

HasReplication returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


