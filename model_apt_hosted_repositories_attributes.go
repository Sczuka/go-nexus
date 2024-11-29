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

// checks if the AptHostedRepositoriesAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AptHostedRepositoriesAttributes{}

// AptHostedRepositoriesAttributes struct for AptHostedRepositoriesAttributes
type AptHostedRepositoriesAttributes struct {
	// Distribution to fetch
	Distribution *string `json:"distribution,omitempty"`
}

// NewAptHostedRepositoriesAttributes instantiates a new AptHostedRepositoriesAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAptHostedRepositoriesAttributes() *AptHostedRepositoriesAttributes {
	this := AptHostedRepositoriesAttributes{}
	return &this
}

// NewAptHostedRepositoriesAttributesWithDefaults instantiates a new AptHostedRepositoriesAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAptHostedRepositoriesAttributesWithDefaults() *AptHostedRepositoriesAttributes {
	this := AptHostedRepositoriesAttributes{}
	return &this
}

// GetDistribution returns the Distribution field value if set, zero value otherwise.
func (o *AptHostedRepositoriesAttributes) GetDistribution() string {
	if o == nil || IsNil(o.Distribution) {
		var ret string
		return ret
	}
	return *o.Distribution
}

// GetDistributionOk returns a tuple with the Distribution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AptHostedRepositoriesAttributes) GetDistributionOk() (*string, bool) {
	if o == nil || IsNil(o.Distribution) {
		return nil, false
	}
	return o.Distribution, true
}

// HasDistribution returns a boolean if a field has been set.
func (o *AptHostedRepositoriesAttributes) HasDistribution() bool {
	if o != nil && !IsNil(o.Distribution) {
		return true
	}

	return false
}

// SetDistribution gets a reference to the given string and assigns it to the Distribution field.
func (o *AptHostedRepositoriesAttributes) SetDistribution(v string) {
	o.Distribution = &v
}

func (o AptHostedRepositoriesAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AptHostedRepositoriesAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Distribution) {
		toSerialize["distribution"] = o.Distribution
	}
	return toSerialize, nil
}

type NullableAptHostedRepositoriesAttributes struct {
	value *AptHostedRepositoriesAttributes
	isSet bool
}

func (v NullableAptHostedRepositoriesAttributes) Get() *AptHostedRepositoriesAttributes {
	return v.value
}

func (v *NullableAptHostedRepositoriesAttributes) Set(val *AptHostedRepositoriesAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAptHostedRepositoriesAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAptHostedRepositoriesAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAptHostedRepositoriesAttributes(val *AptHostedRepositoriesAttributes) *NullableAptHostedRepositoriesAttributes {
	return &NullableAptHostedRepositoriesAttributes{value: val, isSet: true}
}

func (v NullableAptHostedRepositoriesAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAptHostedRepositoriesAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

