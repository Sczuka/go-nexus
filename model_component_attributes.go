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

// checks if the ComponentAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ComponentAttributes{}

// ComponentAttributes struct for ComponentAttributes
type ComponentAttributes struct {
	// Components in this repository count as proprietary for namespace conflict attacks (requires Sonatype Nexus Firewall)
	ProprietaryComponents *bool `json:"proprietaryComponents,omitempty"`
}

// NewComponentAttributes instantiates a new ComponentAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComponentAttributes() *ComponentAttributes {
	this := ComponentAttributes{}
	return &this
}

// NewComponentAttributesWithDefaults instantiates a new ComponentAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComponentAttributesWithDefaults() *ComponentAttributes {
	this := ComponentAttributes{}
	return &this
}

// GetProprietaryComponents returns the ProprietaryComponents field value if set, zero value otherwise.
func (o *ComponentAttributes) GetProprietaryComponents() bool {
	if o == nil || IsNil(o.ProprietaryComponents) {
		var ret bool
		return ret
	}
	return *o.ProprietaryComponents
}

// GetProprietaryComponentsOk returns a tuple with the ProprietaryComponents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentAttributes) GetProprietaryComponentsOk() (*bool, bool) {
	if o == nil || IsNil(o.ProprietaryComponents) {
		return nil, false
	}
	return o.ProprietaryComponents, true
}

// HasProprietaryComponents returns a boolean if a field has been set.
func (o *ComponentAttributes) HasProprietaryComponents() bool {
	if o != nil && !IsNil(o.ProprietaryComponents) {
		return true
	}

	return false
}

// SetProprietaryComponents gets a reference to the given bool and assigns it to the ProprietaryComponents field.
func (o *ComponentAttributes) SetProprietaryComponents(v bool) {
	o.ProprietaryComponents = &v
}

func (o ComponentAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ComponentAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProprietaryComponents) {
		toSerialize["proprietaryComponents"] = o.ProprietaryComponents
	}
	return toSerialize, nil
}

type NullableComponentAttributes struct {
	value *ComponentAttributes
	isSet bool
}

func (v NullableComponentAttributes) Get() *ComponentAttributes {
	return v.value
}

func (v *NullableComponentAttributes) Set(val *ComponentAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableComponentAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableComponentAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComponentAttributes(val *ComponentAttributes) *NullableComponentAttributes {
	return &NullableComponentAttributes{value: val, isSet: true}
}

func (v NullableComponentAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComponentAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


