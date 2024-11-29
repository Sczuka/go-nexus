/*
Nexus Repository Manager REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.70.3-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nexus

import (
	"encoding/json"
)

// checks if the ApiUserSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiUserSource{}

// ApiUserSource struct for ApiUserSource
type ApiUserSource struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewApiUserSource instantiates a new ApiUserSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiUserSource() *ApiUserSource {
	this := ApiUserSource{}
	return &this
}

// NewApiUserSourceWithDefaults instantiates a new ApiUserSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiUserSourceWithDefaults() *ApiUserSource {
	this := ApiUserSource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApiUserSource) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUserSource) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApiUserSource) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApiUserSource) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApiUserSource) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUserSource) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApiUserSource) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApiUserSource) SetName(v string) {
	o.Name = &v
}

func (o ApiUserSource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiUserSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableApiUserSource struct {
	value *ApiUserSource
	isSet bool
}

func (v NullableApiUserSource) Get() *ApiUserSource {
	return v.value
}

func (v *NullableApiUserSource) Set(val *ApiUserSource) {
	v.value = val
	v.isSet = true
}

func (v NullableApiUserSource) IsSet() bool {
	return v.isSet
}

func (v *NullableApiUserSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiUserSource(val *ApiUserSource) *NullableApiUserSource {
	return &NullableApiUserSource{value: val, isSet: true}
}

func (v NullableApiUserSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiUserSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

