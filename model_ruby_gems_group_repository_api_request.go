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

// checks if the RubyGemsGroupRepositoryApiRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RubyGemsGroupRepositoryApiRequest{}

// RubyGemsGroupRepositoryApiRequest struct for RubyGemsGroupRepositoryApiRequest
type RubyGemsGroupRepositoryApiRequest struct {
	// A unique identifier for this repository
	Name string `json:"name" validate:"regexp=^[a-zA-Z0-9\\\\-]{1}[a-zA-Z0-9_\\\\-\\\\.]*$"`
	// Whether this repository accepts incoming requests
	Online bool `json:"online"`
	Storage StorageAttributes `json:"storage"`
	Group GroupAttributes `json:"group"`
}

type _RubyGemsGroupRepositoryApiRequest RubyGemsGroupRepositoryApiRequest

// NewRubyGemsGroupRepositoryApiRequest instantiates a new RubyGemsGroupRepositoryApiRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRubyGemsGroupRepositoryApiRequest(name string, online bool, storage StorageAttributes, group GroupAttributes) *RubyGemsGroupRepositoryApiRequest {
	this := RubyGemsGroupRepositoryApiRequest{}
	this.Name = name
	this.Online = online
	this.Storage = storage
	this.Group = group
	return &this
}

// NewRubyGemsGroupRepositoryApiRequestWithDefaults instantiates a new RubyGemsGroupRepositoryApiRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRubyGemsGroupRepositoryApiRequestWithDefaults() *RubyGemsGroupRepositoryApiRequest {
	this := RubyGemsGroupRepositoryApiRequest{}
	return &this
}

// GetName returns the Name field value
func (o *RubyGemsGroupRepositoryApiRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RubyGemsGroupRepositoryApiRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RubyGemsGroupRepositoryApiRequest) SetName(v string) {
	o.Name = v
}

// GetOnline returns the Online field value
func (o *RubyGemsGroupRepositoryApiRequest) GetOnline() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Online
}

// GetOnlineOk returns a tuple with the Online field value
// and a boolean to check if the value has been set.
func (o *RubyGemsGroupRepositoryApiRequest) GetOnlineOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Online, true
}

// SetOnline sets field value
func (o *RubyGemsGroupRepositoryApiRequest) SetOnline(v bool) {
	o.Online = v
}

// GetStorage returns the Storage field value
func (o *RubyGemsGroupRepositoryApiRequest) GetStorage() StorageAttributes {
	if o == nil {
		var ret StorageAttributes
		return ret
	}

	return o.Storage
}

// GetStorageOk returns a tuple with the Storage field value
// and a boolean to check if the value has been set.
func (o *RubyGemsGroupRepositoryApiRequest) GetStorageOk() (*StorageAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Storage, true
}

// SetStorage sets field value
func (o *RubyGemsGroupRepositoryApiRequest) SetStorage(v StorageAttributes) {
	o.Storage = v
}

// GetGroup returns the Group field value
func (o *RubyGemsGroupRepositoryApiRequest) GetGroup() GroupAttributes {
	if o == nil {
		var ret GroupAttributes
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *RubyGemsGroupRepositoryApiRequest) GetGroupOk() (*GroupAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *RubyGemsGroupRepositoryApiRequest) SetGroup(v GroupAttributes) {
	o.Group = v
}

func (o RubyGemsGroupRepositoryApiRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RubyGemsGroupRepositoryApiRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["online"] = o.Online
	toSerialize["storage"] = o.Storage
	toSerialize["group"] = o.Group
	return toSerialize, nil
}

func (o *RubyGemsGroupRepositoryApiRequest) UnmarshalJSON(data []byte) (err error) {
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

	varRubyGemsGroupRepositoryApiRequest := _RubyGemsGroupRepositoryApiRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRubyGemsGroupRepositoryApiRequest)

	if err != nil {
		return err
	}

	*o = RubyGemsGroupRepositoryApiRequest(varRubyGemsGroupRepositoryApiRequest)

	return err
}

type NullableRubyGemsGroupRepositoryApiRequest struct {
	value *RubyGemsGroupRepositoryApiRequest
	isSet bool
}

func (v NullableRubyGemsGroupRepositoryApiRequest) Get() *RubyGemsGroupRepositoryApiRequest {
	return v.value
}

func (v *NullableRubyGemsGroupRepositoryApiRequest) Set(val *RubyGemsGroupRepositoryApiRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRubyGemsGroupRepositoryApiRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRubyGemsGroupRepositoryApiRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRubyGemsGroupRepositoryApiRequest(val *RubyGemsGroupRepositoryApiRequest) *NullableRubyGemsGroupRepositoryApiRequest {
	return &NullableRubyGemsGroupRepositoryApiRequest{value: val, isSet: true}
}

func (v NullableRubyGemsGroupRepositoryApiRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRubyGemsGroupRepositoryApiRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

