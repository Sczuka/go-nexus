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

// checks if the MavenGroupRepositoryApiRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MavenGroupRepositoryApiRequest{}

// MavenGroupRepositoryApiRequest struct for MavenGroupRepositoryApiRequest
type MavenGroupRepositoryApiRequest struct {
	// A unique identifier for this repository
	Name string `json:"name" validate:"regexp=^[a-zA-Z0-9\\\\-]{1}[a-zA-Z0-9_\\\\-\\\\.]*$"`
	// Whether this repository accepts incoming requests
	Online bool `json:"online"`
	Storage StorageAttributes `json:"storage"`
	Group GroupAttributes `json:"group"`
}

type _MavenGroupRepositoryApiRequest MavenGroupRepositoryApiRequest

// NewMavenGroupRepositoryApiRequest instantiates a new MavenGroupRepositoryApiRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMavenGroupRepositoryApiRequest(name string, online bool, storage StorageAttributes, group GroupAttributes) *MavenGroupRepositoryApiRequest {
	this := MavenGroupRepositoryApiRequest{}
	this.Name = name
	this.Online = online
	this.Storage = storage
	this.Group = group
	return &this
}

// NewMavenGroupRepositoryApiRequestWithDefaults instantiates a new MavenGroupRepositoryApiRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMavenGroupRepositoryApiRequestWithDefaults() *MavenGroupRepositoryApiRequest {
	this := MavenGroupRepositoryApiRequest{}
	return &this
}

// GetName returns the Name field value
func (o *MavenGroupRepositoryApiRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MavenGroupRepositoryApiRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MavenGroupRepositoryApiRequest) SetName(v string) {
	o.Name = v
}

// GetOnline returns the Online field value
func (o *MavenGroupRepositoryApiRequest) GetOnline() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Online
}

// GetOnlineOk returns a tuple with the Online field value
// and a boolean to check if the value has been set.
func (o *MavenGroupRepositoryApiRequest) GetOnlineOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Online, true
}

// SetOnline sets field value
func (o *MavenGroupRepositoryApiRequest) SetOnline(v bool) {
	o.Online = v
}

// GetStorage returns the Storage field value
func (o *MavenGroupRepositoryApiRequest) GetStorage() StorageAttributes {
	if o == nil {
		var ret StorageAttributes
		return ret
	}

	return o.Storage
}

// GetStorageOk returns a tuple with the Storage field value
// and a boolean to check if the value has been set.
func (o *MavenGroupRepositoryApiRequest) GetStorageOk() (*StorageAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Storage, true
}

// SetStorage sets field value
func (o *MavenGroupRepositoryApiRequest) SetStorage(v StorageAttributes) {
	o.Storage = v
}

// GetGroup returns the Group field value
func (o *MavenGroupRepositoryApiRequest) GetGroup() GroupAttributes {
	if o == nil {
		var ret GroupAttributes
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *MavenGroupRepositoryApiRequest) GetGroupOk() (*GroupAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *MavenGroupRepositoryApiRequest) SetGroup(v GroupAttributes) {
	o.Group = v
}

func (o MavenGroupRepositoryApiRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MavenGroupRepositoryApiRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["online"] = o.Online
	toSerialize["storage"] = o.Storage
	toSerialize["group"] = o.Group
	return toSerialize, nil
}

func (o *MavenGroupRepositoryApiRequest) UnmarshalJSON(data []byte) (err error) {
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

	varMavenGroupRepositoryApiRequest := _MavenGroupRepositoryApiRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMavenGroupRepositoryApiRequest)

	if err != nil {
		return err
	}

	*o = MavenGroupRepositoryApiRequest(varMavenGroupRepositoryApiRequest)

	return err
}

type NullableMavenGroupRepositoryApiRequest struct {
	value *MavenGroupRepositoryApiRequest
	isSet bool
}

func (v NullableMavenGroupRepositoryApiRequest) Get() *MavenGroupRepositoryApiRequest {
	return v.value
}

func (v *NullableMavenGroupRepositoryApiRequest) Set(val *MavenGroupRepositoryApiRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMavenGroupRepositoryApiRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMavenGroupRepositoryApiRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMavenGroupRepositoryApiRequest(val *MavenGroupRepositoryApiRequest) *NullableMavenGroupRepositoryApiRequest {
	return &NullableMavenGroupRepositoryApiRequest{value: val, isSet: true}
}

func (v NullableMavenGroupRepositoryApiRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMavenGroupRepositoryApiRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


