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

// checks if the MavenHostedApiRepository type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MavenHostedApiRepository{}

// MavenHostedApiRepository struct for MavenHostedApiRepository
type MavenHostedApiRepository struct {
	// A unique identifier for this repository
	Name *string `json:"name,omitempty" validate:"regexp=^[a-zA-Z0-9\\\\-]{1}[a-zA-Z0-9_\\\\-\\\\.]*$"`
	// Whether this repository accepts incoming requests
	Online bool `json:"online"`
	Storage HostedStorageAttributes `json:"storage"`
	Cleanup *CleanupPolicyAttributes `json:"cleanup,omitempty"`
	Component *ComponentAttributes `json:"component,omitempty"`
	Maven MavenAttributes `json:"maven"`
}

type _MavenHostedApiRepository MavenHostedApiRepository

// NewMavenHostedApiRepository instantiates a new MavenHostedApiRepository object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMavenHostedApiRepository(online bool, storage HostedStorageAttributes, maven MavenAttributes) *MavenHostedApiRepository {
	this := MavenHostedApiRepository{}
	this.Online = online
	this.Storage = storage
	this.Maven = maven
	return &this
}

// NewMavenHostedApiRepositoryWithDefaults instantiates a new MavenHostedApiRepository object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMavenHostedApiRepositoryWithDefaults() *MavenHostedApiRepository {
	this := MavenHostedApiRepository{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MavenHostedApiRepository) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MavenHostedApiRepository) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MavenHostedApiRepository) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MavenHostedApiRepository) SetName(v string) {
	o.Name = &v
}

// GetOnline returns the Online field value
func (o *MavenHostedApiRepository) GetOnline() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Online
}

// GetOnlineOk returns a tuple with the Online field value
// and a boolean to check if the value has been set.
func (o *MavenHostedApiRepository) GetOnlineOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Online, true
}

// SetOnline sets field value
func (o *MavenHostedApiRepository) SetOnline(v bool) {
	o.Online = v
}

// GetStorage returns the Storage field value
func (o *MavenHostedApiRepository) GetStorage() HostedStorageAttributes {
	if o == nil {
		var ret HostedStorageAttributes
		return ret
	}

	return o.Storage
}

// GetStorageOk returns a tuple with the Storage field value
// and a boolean to check if the value has been set.
func (o *MavenHostedApiRepository) GetStorageOk() (*HostedStorageAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Storage, true
}

// SetStorage sets field value
func (o *MavenHostedApiRepository) SetStorage(v HostedStorageAttributes) {
	o.Storage = v
}

// GetCleanup returns the Cleanup field value if set, zero value otherwise.
func (o *MavenHostedApiRepository) GetCleanup() CleanupPolicyAttributes {
	if o == nil || IsNil(o.Cleanup) {
		var ret CleanupPolicyAttributes
		return ret
	}
	return *o.Cleanup
}

// GetCleanupOk returns a tuple with the Cleanup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MavenHostedApiRepository) GetCleanupOk() (*CleanupPolicyAttributes, bool) {
	if o == nil || IsNil(o.Cleanup) {
		return nil, false
	}
	return o.Cleanup, true
}

// HasCleanup returns a boolean if a field has been set.
func (o *MavenHostedApiRepository) HasCleanup() bool {
	if o != nil && !IsNil(o.Cleanup) {
		return true
	}

	return false
}

// SetCleanup gets a reference to the given CleanupPolicyAttributes and assigns it to the Cleanup field.
func (o *MavenHostedApiRepository) SetCleanup(v CleanupPolicyAttributes) {
	o.Cleanup = &v
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *MavenHostedApiRepository) GetComponent() ComponentAttributes {
	if o == nil || IsNil(o.Component) {
		var ret ComponentAttributes
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MavenHostedApiRepository) GetComponentOk() (*ComponentAttributes, bool) {
	if o == nil || IsNil(o.Component) {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *MavenHostedApiRepository) HasComponent() bool {
	if o != nil && !IsNil(o.Component) {
		return true
	}

	return false
}

// SetComponent gets a reference to the given ComponentAttributes and assigns it to the Component field.
func (o *MavenHostedApiRepository) SetComponent(v ComponentAttributes) {
	o.Component = &v
}

// GetMaven returns the Maven field value
func (o *MavenHostedApiRepository) GetMaven() MavenAttributes {
	if o == nil {
		var ret MavenAttributes
		return ret
	}

	return o.Maven
}

// GetMavenOk returns a tuple with the Maven field value
// and a boolean to check if the value has been set.
func (o *MavenHostedApiRepository) GetMavenOk() (*MavenAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Maven, true
}

// SetMaven sets field value
func (o *MavenHostedApiRepository) SetMaven(v MavenAttributes) {
	o.Maven = v
}

func (o MavenHostedApiRepository) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MavenHostedApiRepository) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["online"] = o.Online
	toSerialize["storage"] = o.Storage
	if !IsNil(o.Cleanup) {
		toSerialize["cleanup"] = o.Cleanup
	}
	if !IsNil(o.Component) {
		toSerialize["component"] = o.Component
	}
	toSerialize["maven"] = o.Maven
	return toSerialize, nil
}

func (o *MavenHostedApiRepository) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"online",
		"storage",
		"maven",
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

	varMavenHostedApiRepository := _MavenHostedApiRepository{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMavenHostedApiRepository)

	if err != nil {
		return err
	}

	*o = MavenHostedApiRepository(varMavenHostedApiRepository)

	return err
}

type NullableMavenHostedApiRepository struct {
	value *MavenHostedApiRepository
	isSet bool
}

func (v NullableMavenHostedApiRepository) Get() *MavenHostedApiRepository {
	return v.value
}

func (v *NullableMavenHostedApiRepository) Set(val *MavenHostedApiRepository) {
	v.value = val
	v.isSet = true
}

func (v NullableMavenHostedApiRepository) IsSet() bool {
	return v.isSet
}

func (v *NullableMavenHostedApiRepository) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMavenHostedApiRepository(val *MavenHostedApiRepository) *NullableMavenHostedApiRepository {
	return &NullableMavenHostedApiRepository{value: val, isSet: true}
}

func (v NullableMavenHostedApiRepository) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMavenHostedApiRepository) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

