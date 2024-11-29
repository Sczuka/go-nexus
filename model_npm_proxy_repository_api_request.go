/*
Nexus Repository Manager REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.70.3-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nexus

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the NpmProxyRepositoryApiRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NpmProxyRepositoryApiRequest{}

// NpmProxyRepositoryApiRequest struct for NpmProxyRepositoryApiRequest
type NpmProxyRepositoryApiRequest struct {
	// A unique identifier for this repository
	Name string `json:"name" validate:"regexp=^[a-zA-Z0-9\\\\-]{1}[a-zA-Z0-9_\\\\-\\\\.]*$"`
	// Whether this repository accepts incoming requests
	Online bool `json:"online"`
	Storage StorageAttributes `json:"storage"`
	Cleanup *CleanupPolicyAttributes `json:"cleanup,omitempty"`
	Proxy ProxyAttributes `json:"proxy"`
	NegativeCache NegativeCacheAttributes `json:"negativeCache"`
	HttpClient HttpClientAttributes `json:"httpClient"`
	RoutingRule *string `json:"routingRule,omitempty"`
	Replication *ReplicationAttributes `json:"replication,omitempty"`
	Npm *NpmAttributes `json:"npm,omitempty"`
}

type _NpmProxyRepositoryApiRequest NpmProxyRepositoryApiRequest

// NewNpmProxyRepositoryApiRequest instantiates a new NpmProxyRepositoryApiRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNpmProxyRepositoryApiRequest(name string, online bool, storage StorageAttributes, proxy ProxyAttributes, negativeCache NegativeCacheAttributes, httpClient HttpClientAttributes) *NpmProxyRepositoryApiRequest {
	this := NpmProxyRepositoryApiRequest{}
	this.Name = name
	this.Online = online
	this.Storage = storage
	this.Proxy = proxy
	this.NegativeCache = negativeCache
	this.HttpClient = httpClient
	return &this
}

// NewNpmProxyRepositoryApiRequestWithDefaults instantiates a new NpmProxyRepositoryApiRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNpmProxyRepositoryApiRequestWithDefaults() *NpmProxyRepositoryApiRequest {
	this := NpmProxyRepositoryApiRequest{}
	return &this
}

// GetName returns the Name field value
func (o *NpmProxyRepositoryApiRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NpmProxyRepositoryApiRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NpmProxyRepositoryApiRequest) SetName(v string) {
	o.Name = v
}

// GetOnline returns the Online field value
func (o *NpmProxyRepositoryApiRequest) GetOnline() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Online
}

// GetOnlineOk returns a tuple with the Online field value
// and a boolean to check if the value has been set.
func (o *NpmProxyRepositoryApiRequest) GetOnlineOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Online, true
}

// SetOnline sets field value
func (o *NpmProxyRepositoryApiRequest) SetOnline(v bool) {
	o.Online = v
}

// GetStorage returns the Storage field value
func (o *NpmProxyRepositoryApiRequest) GetStorage() StorageAttributes {
	if o == nil {
		var ret StorageAttributes
		return ret
	}

	return o.Storage
}

// GetStorageOk returns a tuple with the Storage field value
// and a boolean to check if the value has been set.
func (o *NpmProxyRepositoryApiRequest) GetStorageOk() (*StorageAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Storage, true
}

// SetStorage sets field value
func (o *NpmProxyRepositoryApiRequest) SetStorage(v StorageAttributes) {
	o.Storage = v
}

// GetCleanup returns the Cleanup field value if set, zero value otherwise.
func (o *NpmProxyRepositoryApiRequest) GetCleanup() CleanupPolicyAttributes {
	if o == nil || IsNil(o.Cleanup) {
		var ret CleanupPolicyAttributes
		return ret
	}
	return *o.Cleanup
}

// GetCleanupOk returns a tuple with the Cleanup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NpmProxyRepositoryApiRequest) GetCleanupOk() (*CleanupPolicyAttributes, bool) {
	if o == nil || IsNil(o.Cleanup) {
		return nil, false
	}
	return o.Cleanup, true
}

// HasCleanup returns a boolean if a field has been set.
func (o *NpmProxyRepositoryApiRequest) HasCleanup() bool {
	if o != nil && !IsNil(o.Cleanup) {
		return true
	}

	return false
}

// SetCleanup gets a reference to the given CleanupPolicyAttributes and assigns it to the Cleanup field.
func (o *NpmProxyRepositoryApiRequest) SetCleanup(v CleanupPolicyAttributes) {
	o.Cleanup = &v
}

// GetProxy returns the Proxy field value
func (o *NpmProxyRepositoryApiRequest) GetProxy() ProxyAttributes {
	if o == nil {
		var ret ProxyAttributes
		return ret
	}

	return o.Proxy
}

// GetProxyOk returns a tuple with the Proxy field value
// and a boolean to check if the value has been set.
func (o *NpmProxyRepositoryApiRequest) GetProxyOk() (*ProxyAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Proxy, true
}

// SetProxy sets field value
func (o *NpmProxyRepositoryApiRequest) SetProxy(v ProxyAttributes) {
	o.Proxy = v
}

// GetNegativeCache returns the NegativeCache field value
func (o *NpmProxyRepositoryApiRequest) GetNegativeCache() NegativeCacheAttributes {
	if o == nil {
		var ret NegativeCacheAttributes
		return ret
	}

	return o.NegativeCache
}

// GetNegativeCacheOk returns a tuple with the NegativeCache field value
// and a boolean to check if the value has been set.
func (o *NpmProxyRepositoryApiRequest) GetNegativeCacheOk() (*NegativeCacheAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NegativeCache, true
}

// SetNegativeCache sets field value
func (o *NpmProxyRepositoryApiRequest) SetNegativeCache(v NegativeCacheAttributes) {
	o.NegativeCache = v
}

// GetHttpClient returns the HttpClient field value
func (o *NpmProxyRepositoryApiRequest) GetHttpClient() HttpClientAttributes {
	if o == nil {
		var ret HttpClientAttributes
		return ret
	}

	return o.HttpClient
}

// GetHttpClientOk returns a tuple with the HttpClient field value
// and a boolean to check if the value has been set.
func (o *NpmProxyRepositoryApiRequest) GetHttpClientOk() (*HttpClientAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HttpClient, true
}

// SetHttpClient sets field value
func (o *NpmProxyRepositoryApiRequest) SetHttpClient(v HttpClientAttributes) {
	o.HttpClient = v
}

// GetRoutingRule returns the RoutingRule field value if set, zero value otherwise.
func (o *NpmProxyRepositoryApiRequest) GetRoutingRule() string {
	if o == nil || IsNil(o.RoutingRule) {
		var ret string
		return ret
	}
	return *o.RoutingRule
}

// GetRoutingRuleOk returns a tuple with the RoutingRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NpmProxyRepositoryApiRequest) GetRoutingRuleOk() (*string, bool) {
	if o == nil || IsNil(o.RoutingRule) {
		return nil, false
	}
	return o.RoutingRule, true
}

// HasRoutingRule returns a boolean if a field has been set.
func (o *NpmProxyRepositoryApiRequest) HasRoutingRule() bool {
	if o != nil && !IsNil(o.RoutingRule) {
		return true
	}

	return false
}

// SetRoutingRule gets a reference to the given string and assigns it to the RoutingRule field.
func (o *NpmProxyRepositoryApiRequest) SetRoutingRule(v string) {
	o.RoutingRule = &v
}

// GetReplication returns the Replication field value if set, zero value otherwise.
func (o *NpmProxyRepositoryApiRequest) GetReplication() ReplicationAttributes {
	if o == nil || IsNil(o.Replication) {
		var ret ReplicationAttributes
		return ret
	}
	return *o.Replication
}

// GetReplicationOk returns a tuple with the Replication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NpmProxyRepositoryApiRequest) GetReplicationOk() (*ReplicationAttributes, bool) {
	if o == nil || IsNil(o.Replication) {
		return nil, false
	}
	return o.Replication, true
}

// HasReplication returns a boolean if a field has been set.
func (o *NpmProxyRepositoryApiRequest) HasReplication() bool {
	if o != nil && !IsNil(o.Replication) {
		return true
	}

	return false
}

// SetReplication gets a reference to the given ReplicationAttributes and assigns it to the Replication field.
func (o *NpmProxyRepositoryApiRequest) SetReplication(v ReplicationAttributes) {
	o.Replication = &v
}

// GetNpm returns the Npm field value if set, zero value otherwise.
func (o *NpmProxyRepositoryApiRequest) GetNpm() NpmAttributes {
	if o == nil || IsNil(o.Npm) {
		var ret NpmAttributes
		return ret
	}
	return *o.Npm
}

// GetNpmOk returns a tuple with the Npm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NpmProxyRepositoryApiRequest) GetNpmOk() (*NpmAttributes, bool) {
	if o == nil || IsNil(o.Npm) {
		return nil, false
	}
	return o.Npm, true
}

// HasNpm returns a boolean if a field has been set.
func (o *NpmProxyRepositoryApiRequest) HasNpm() bool {
	if o != nil && !IsNil(o.Npm) {
		return true
	}

	return false
}

// SetNpm gets a reference to the given NpmAttributes and assigns it to the Npm field.
func (o *NpmProxyRepositoryApiRequest) SetNpm(v NpmAttributes) {
	o.Npm = &v
}

func (o NpmProxyRepositoryApiRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NpmProxyRepositoryApiRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["online"] = o.Online
	toSerialize["storage"] = o.Storage
	if !IsNil(o.Cleanup) {
		toSerialize["cleanup"] = o.Cleanup
	}
	toSerialize["proxy"] = o.Proxy
	toSerialize["negativeCache"] = o.NegativeCache
	toSerialize["httpClient"] = o.HttpClient
	if !IsNil(o.RoutingRule) {
		toSerialize["routingRule"] = o.RoutingRule
	}
	if !IsNil(o.Replication) {
		toSerialize["replication"] = o.Replication
	}
	if !IsNil(o.Npm) {
		toSerialize["npm"] = o.Npm
	}
	return toSerialize, nil
}

func (o *NpmProxyRepositoryApiRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"online",
		"storage",
		"proxy",
		"negativeCache",
		"httpClient",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varNpmProxyRepositoryApiRequest := _NpmProxyRepositoryApiRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNpmProxyRepositoryApiRequest)

	if err != nil {
		return err
	}

	*o = NpmProxyRepositoryApiRequest(varNpmProxyRepositoryApiRequest)

	return err
}

type NullableNpmProxyRepositoryApiRequest struct {
	value *NpmProxyRepositoryApiRequest
	isSet bool
}

func (v NullableNpmProxyRepositoryApiRequest) Get() *NpmProxyRepositoryApiRequest {
	return v.value
}

func (v *NullableNpmProxyRepositoryApiRequest) Set(val *NpmProxyRepositoryApiRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNpmProxyRepositoryApiRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNpmProxyRepositoryApiRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNpmProxyRepositoryApiRequest(val *NpmProxyRepositoryApiRequest) *NullableNpmProxyRepositoryApiRequest {
	return &NullableNpmProxyRepositoryApiRequest{value: val, isSet: true}
}

func (v NullableNpmProxyRepositoryApiRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNpmProxyRepositoryApiRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


