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

// checks if the YumGroupRepositoryApiRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &YumGroupRepositoryApiRequest{}

// YumGroupRepositoryApiRequest struct for YumGroupRepositoryApiRequest
type YumGroupRepositoryApiRequest struct {
	// A unique identifier for this repository
	Name string `json:"name" validate:"regexp=^[a-zA-Z0-9\\\\-]{1}[a-zA-Z0-9_\\\\-\\\\.]*$"`
	// Whether this repository accepts incoming requests
	Online bool `json:"online"`
	Storage StorageAttributes `json:"storage"`
	Group GroupAttributes `json:"group"`
	YumSigning *YumSigningRepositoriesAttributes `json:"yumSigning,omitempty"`
}

type _YumGroupRepositoryApiRequest YumGroupRepositoryApiRequest

// NewYumGroupRepositoryApiRequest instantiates a new YumGroupRepositoryApiRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewYumGroupRepositoryApiRequest(name string, online bool, storage StorageAttributes, group GroupAttributes) *YumGroupRepositoryApiRequest {
	this := YumGroupRepositoryApiRequest{}
	this.Name = name
	this.Online = online
	this.Storage = storage
	this.Group = group
	return &this
}

// NewYumGroupRepositoryApiRequestWithDefaults instantiates a new YumGroupRepositoryApiRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewYumGroupRepositoryApiRequestWithDefaults() *YumGroupRepositoryApiRequest {
	this := YumGroupRepositoryApiRequest{}
	return &this
}

// GetName returns the Name field value
func (o *YumGroupRepositoryApiRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *YumGroupRepositoryApiRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *YumGroupRepositoryApiRequest) SetName(v string) {
	o.Name = v
}

// GetOnline returns the Online field value
func (o *YumGroupRepositoryApiRequest) GetOnline() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Online
}

// GetOnlineOk returns a tuple with the Online field value
// and a boolean to check if the value has been set.
func (o *YumGroupRepositoryApiRequest) GetOnlineOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Online, true
}

// SetOnline sets field value
func (o *YumGroupRepositoryApiRequest) SetOnline(v bool) {
	o.Online = v
}

// GetStorage returns the Storage field value
func (o *YumGroupRepositoryApiRequest) GetStorage() StorageAttributes {
	if o == nil {
		var ret StorageAttributes
		return ret
	}

	return o.Storage
}

// GetStorageOk returns a tuple with the Storage field value
// and a boolean to check if the value has been set.
func (o *YumGroupRepositoryApiRequest) GetStorageOk() (*StorageAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Storage, true
}

// SetStorage sets field value
func (o *YumGroupRepositoryApiRequest) SetStorage(v StorageAttributes) {
	o.Storage = v
}

// GetGroup returns the Group field value
func (o *YumGroupRepositoryApiRequest) GetGroup() GroupAttributes {
	if o == nil {
		var ret GroupAttributes
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *YumGroupRepositoryApiRequest) GetGroupOk() (*GroupAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *YumGroupRepositoryApiRequest) SetGroup(v GroupAttributes) {
	o.Group = v
}

// GetYumSigning returns the YumSigning field value if set, zero value otherwise.
func (o *YumGroupRepositoryApiRequest) GetYumSigning() YumSigningRepositoriesAttributes {
	if o == nil || IsNil(o.YumSigning) {
		var ret YumSigningRepositoriesAttributes
		return ret
	}
	return *o.YumSigning
}

// GetYumSigningOk returns a tuple with the YumSigning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *YumGroupRepositoryApiRequest) GetYumSigningOk() (*YumSigningRepositoriesAttributes, bool) {
	if o == nil || IsNil(o.YumSigning) {
		return nil, false
	}
	return o.YumSigning, true
}

// HasYumSigning returns a boolean if a field has been set.
func (o *YumGroupRepositoryApiRequest) HasYumSigning() bool {
	if o != nil && !IsNil(o.YumSigning) {
		return true
	}

	return false
}

// SetYumSigning gets a reference to the given YumSigningRepositoriesAttributes and assigns it to the YumSigning field.
func (o *YumGroupRepositoryApiRequest) SetYumSigning(v YumSigningRepositoriesAttributes) {
	o.YumSigning = &v
}

func (o YumGroupRepositoryApiRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o YumGroupRepositoryApiRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["online"] = o.Online
	toSerialize["storage"] = o.Storage
	toSerialize["group"] = o.Group
	if !IsNil(o.YumSigning) {
		toSerialize["yumSigning"] = o.YumSigning
	}
	return toSerialize, nil
}

func (o *YumGroupRepositoryApiRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"online",
		"storage",
		"group",
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

	varYumGroupRepositoryApiRequest := _YumGroupRepositoryApiRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varYumGroupRepositoryApiRequest)

	if err != nil {
		return err
	}

	*o = YumGroupRepositoryApiRequest(varYumGroupRepositoryApiRequest)

	return err
}

type NullableYumGroupRepositoryApiRequest struct {
	value *YumGroupRepositoryApiRequest
	isSet bool
}

func (v NullableYumGroupRepositoryApiRequest) Get() *YumGroupRepositoryApiRequest {
	return v.value
}

func (v *NullableYumGroupRepositoryApiRequest) Set(val *YumGroupRepositoryApiRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableYumGroupRepositoryApiRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableYumGroupRepositoryApiRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableYumGroupRepositoryApiRequest(val *YumGroupRepositoryApiRequest) *NullableYumGroupRepositoryApiRequest {
	return &NullableYumGroupRepositoryApiRequest{value: val, isSet: true}
}

func (v NullableYumGroupRepositoryApiRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableYumGroupRepositoryApiRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

