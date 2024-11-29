# BowerProxyApiRepository

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
**Bower** | Pointer to [**BowerAttributes**](BowerAttributes.md) |  | [optional] 

## Methods

### NewBowerProxyApiRepository

`func NewBowerProxyApiRepository(online bool, storage StorageAttributes, proxy ProxyAttributes, negativeCache NegativeCacheAttributes, httpClient HttpClientAttributes, ) *BowerProxyApiRepository`

NewBowerProxyApiRepository instantiates a new BowerProxyApiRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBowerProxyApiRepositoryWithDefaults

`func NewBowerProxyApiRepositoryWithDefaults() *BowerProxyApiRepository`

NewBowerProxyApiRepositoryWithDefaults instantiates a new BowerProxyApiRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BowerProxyApiRepository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BowerProxyApiRepository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BowerProxyApiRepository) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BowerProxyApiRepository) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOnline

`func (o *BowerProxyApiRepository) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *BowerProxyApiRepository) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *BowerProxyApiRepository) SetOnline(v bool)`

SetOnline sets Online field to given value.


### GetStorage

`func (o *BowerProxyApiRepository) GetStorage() StorageAttributes`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *BowerProxyApiRepository) GetStorageOk() (*StorageAttributes, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *BowerProxyApiRepository) SetStorage(v StorageAttributes)`

SetStorage sets Storage field to given value.


### GetCleanup

`func (o *BowerProxyApiRepository) GetCleanup() CleanupPolicyAttributes`

GetCleanup returns the Cleanup field if non-nil, zero value otherwise.

### GetCleanupOk

`func (o *BowerProxyApiRepository) GetCleanupOk() (*CleanupPolicyAttributes, bool)`

GetCleanupOk returns a tuple with the Cleanup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanup

`func (o *BowerProxyApiRepository) SetCleanup(v CleanupPolicyAttributes)`

SetCleanup sets Cleanup field to given value.

### HasCleanup

`func (o *BowerProxyApiRepository) HasCleanup() bool`

HasCleanup returns a boolean if a field has been set.

### GetProxy

`func (o *BowerProxyApiRepository) GetProxy() ProxyAttributes`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *BowerProxyApiRepository) GetProxyOk() (*ProxyAttributes, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *BowerProxyApiRepository) SetProxy(v ProxyAttributes)`

SetProxy sets Proxy field to given value.


### GetNegativeCache

`func (o *BowerProxyApiRepository) GetNegativeCache() NegativeCacheAttributes`

GetNegativeCache returns the NegativeCache field if non-nil, zero value otherwise.

### GetNegativeCacheOk

`func (o *BowerProxyApiRepository) GetNegativeCacheOk() (*NegativeCacheAttributes, bool)`

GetNegativeCacheOk returns a tuple with the NegativeCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegativeCache

`func (o *BowerProxyApiRepository) SetNegativeCache(v NegativeCacheAttributes)`

SetNegativeCache sets NegativeCache field to given value.


### GetHttpClient

`func (o *BowerProxyApiRepository) GetHttpClient() HttpClientAttributes`

GetHttpClient returns the HttpClient field if non-nil, zero value otherwise.

### GetHttpClientOk

`func (o *BowerProxyApiRepository) GetHttpClientOk() (*HttpClientAttributes, bool)`

GetHttpClientOk returns a tuple with the HttpClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpClient

`func (o *BowerProxyApiRepository) SetHttpClient(v HttpClientAttributes)`

SetHttpClient sets HttpClient field to given value.


### GetRoutingRuleName

`func (o *BowerProxyApiRepository) GetRoutingRuleName() string`

GetRoutingRuleName returns the RoutingRuleName field if non-nil, zero value otherwise.

### GetRoutingRuleNameOk

`func (o *BowerProxyApiRepository) GetRoutingRuleNameOk() (*string, bool)`

GetRoutingRuleNameOk returns a tuple with the RoutingRuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingRuleName

`func (o *BowerProxyApiRepository) SetRoutingRuleName(v string)`

SetRoutingRuleName sets RoutingRuleName field to given value.

### HasRoutingRuleName

`func (o *BowerProxyApiRepository) HasRoutingRuleName() bool`

HasRoutingRuleName returns a boolean if a field has been set.

### GetReplication

`func (o *BowerProxyApiRepository) GetReplication() ReplicationAttributes`

GetReplication returns the Replication field if non-nil, zero value otherwise.

### GetReplicationOk

`func (o *BowerProxyApiRepository) GetReplicationOk() (*ReplicationAttributes, bool)`

GetReplicationOk returns a tuple with the Replication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplication

`func (o *BowerProxyApiRepository) SetReplication(v ReplicationAttributes)`

SetReplication sets Replication field to given value.

### HasReplication

`func (o *BowerProxyApiRepository) HasReplication() bool`

HasReplication returns a boolean if a field has been set.

### GetBower

`func (o *BowerProxyApiRepository) GetBower() BowerAttributes`

GetBower returns the Bower field if non-nil, zero value otherwise.

### GetBowerOk

`func (o *BowerProxyApiRepository) GetBowerOk() (*BowerAttributes, bool)`

GetBowerOk returns a tuple with the Bower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBower

`func (o *BowerProxyApiRepository) SetBower(v BowerAttributes)`

SetBower sets Bower field to given value.

### HasBower

`func (o *BowerProxyApiRepository) HasBower() bool`

HasBower returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


