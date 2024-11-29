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

// checks if the SamlConfigurationXO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SamlConfigurationXO{}

// SamlConfigurationXO struct for SamlConfigurationXO
type SamlConfigurationXO struct {
	// SAML Service Provider's unique identifying URI
	EntityId *string `json:"entityId,omitempty"`
	// SAML Identity Provider Metadata XML
	IdpMetadata string `json:"idpMetadata"`
	// SAML attribute name for the username
	UsernameAttribute string `json:"usernameAttribute"`
	// SAML attribute name for the first name
	FirstNameAttribute *string `json:"firstNameAttribute,omitempty"`
	// SAML attribute name for the last name
	LastNameAttribute *string `json:"lastNameAttribute,omitempty"`
	// SAML attribute name for email
	EmailAttribute *string `json:"emailAttribute,omitempty"`
	// SAML attribute name for groups which maps the Identity Provider groups to a Nexus Repository Manager role
	GroupsAttribute *string `json:"groupsAttribute,omitempty"`
	// Validate signatures on responses from Identity Provider
	ValidateResponseSignature *bool `json:"validateResponseSignature,omitempty"`
	// Validate signatures on assertions from Identity Provider
	ValidateAssertionSignature *bool `json:"validateAssertionSignature,omitempty"`
}

type _SamlConfigurationXO SamlConfigurationXO

// NewSamlConfigurationXO instantiates a new SamlConfigurationXO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSamlConfigurationXO(idpMetadata string, usernameAttribute string) *SamlConfigurationXO {
	this := SamlConfigurationXO{}
	this.IdpMetadata = idpMetadata
	this.UsernameAttribute = usernameAttribute
	return &this
}

// NewSamlConfigurationXOWithDefaults instantiates a new SamlConfigurationXO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSamlConfigurationXOWithDefaults() *SamlConfigurationXO {
	this := SamlConfigurationXO{}
	return &this
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *SamlConfigurationXO) GetEntityId() string {
	if o == nil || IsNil(o.EntityId) {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlConfigurationXO) GetEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.EntityId) {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *SamlConfigurationXO) HasEntityId() bool {
	if o != nil && !IsNil(o.EntityId) {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *SamlConfigurationXO) SetEntityId(v string) {
	o.EntityId = &v
}

// GetIdpMetadata returns the IdpMetadata field value
func (o *SamlConfigurationXO) GetIdpMetadata() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdpMetadata
}

// GetIdpMetadataOk returns a tuple with the IdpMetadata field value
// and a boolean to check if the value has been set.
func (o *SamlConfigurationXO) GetIdpMetadataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdpMetadata, true
}

// SetIdpMetadata sets field value
func (o *SamlConfigurationXO) SetIdpMetadata(v string) {
	o.IdpMetadata = v
}

// GetUsernameAttribute returns the UsernameAttribute field value
func (o *SamlConfigurationXO) GetUsernameAttribute() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UsernameAttribute
}

// GetUsernameAttributeOk returns a tuple with the UsernameAttribute field value
// and a boolean to check if the value has been set.
func (o *SamlConfigurationXO) GetUsernameAttributeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UsernameAttribute, true
}

// SetUsernameAttribute sets field value
func (o *SamlConfigurationXO) SetUsernameAttribute(v string) {
	o.UsernameAttribute = v
}

// GetFirstNameAttribute returns the FirstNameAttribute field value if set, zero value otherwise.
func (o *SamlConfigurationXO) GetFirstNameAttribute() string {
	if o == nil || IsNil(o.FirstNameAttribute) {
		var ret string
		return ret
	}
	return *o.FirstNameAttribute
}

// GetFirstNameAttributeOk returns a tuple with the FirstNameAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlConfigurationXO) GetFirstNameAttributeOk() (*string, bool) {
	if o == nil || IsNil(o.FirstNameAttribute) {
		return nil, false
	}
	return o.FirstNameAttribute, true
}

// HasFirstNameAttribute returns a boolean if a field has been set.
func (o *SamlConfigurationXO) HasFirstNameAttribute() bool {
	if o != nil && !IsNil(o.FirstNameAttribute) {
		return true
	}

	return false
}

// SetFirstNameAttribute gets a reference to the given string and assigns it to the FirstNameAttribute field.
func (o *SamlConfigurationXO) SetFirstNameAttribute(v string) {
	o.FirstNameAttribute = &v
}

// GetLastNameAttribute returns the LastNameAttribute field value if set, zero value otherwise.
func (o *SamlConfigurationXO) GetLastNameAttribute() string {
	if o == nil || IsNil(o.LastNameAttribute) {
		var ret string
		return ret
	}
	return *o.LastNameAttribute
}

// GetLastNameAttributeOk returns a tuple with the LastNameAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlConfigurationXO) GetLastNameAttributeOk() (*string, bool) {
	if o == nil || IsNil(o.LastNameAttribute) {
		return nil, false
	}
	return o.LastNameAttribute, true
}

// HasLastNameAttribute returns a boolean if a field has been set.
func (o *SamlConfigurationXO) HasLastNameAttribute() bool {
	if o != nil && !IsNil(o.LastNameAttribute) {
		return true
	}

	return false
}

// SetLastNameAttribute gets a reference to the given string and assigns it to the LastNameAttribute field.
func (o *SamlConfigurationXO) SetLastNameAttribute(v string) {
	o.LastNameAttribute = &v
}

// GetEmailAttribute returns the EmailAttribute field value if set, zero value otherwise.
func (o *SamlConfigurationXO) GetEmailAttribute() string {
	if o == nil || IsNil(o.EmailAttribute) {
		var ret string
		return ret
	}
	return *o.EmailAttribute
}

// GetEmailAttributeOk returns a tuple with the EmailAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlConfigurationXO) GetEmailAttributeOk() (*string, bool) {
	if o == nil || IsNil(o.EmailAttribute) {
		return nil, false
	}
	return o.EmailAttribute, true
}

// HasEmailAttribute returns a boolean if a field has been set.
func (o *SamlConfigurationXO) HasEmailAttribute() bool {
	if o != nil && !IsNil(o.EmailAttribute) {
		return true
	}

	return false
}

// SetEmailAttribute gets a reference to the given string and assigns it to the EmailAttribute field.
func (o *SamlConfigurationXO) SetEmailAttribute(v string) {
	o.EmailAttribute = &v
}

// GetGroupsAttribute returns the GroupsAttribute field value if set, zero value otherwise.
func (o *SamlConfigurationXO) GetGroupsAttribute() string {
	if o == nil || IsNil(o.GroupsAttribute) {
		var ret string
		return ret
	}
	return *o.GroupsAttribute
}

// GetGroupsAttributeOk returns a tuple with the GroupsAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlConfigurationXO) GetGroupsAttributeOk() (*string, bool) {
	if o == nil || IsNil(o.GroupsAttribute) {
		return nil, false
	}
	return o.GroupsAttribute, true
}

// HasGroupsAttribute returns a boolean if a field has been set.
func (o *SamlConfigurationXO) HasGroupsAttribute() bool {
	if o != nil && !IsNil(o.GroupsAttribute) {
		return true
	}

	return false
}

// SetGroupsAttribute gets a reference to the given string and assigns it to the GroupsAttribute field.
func (o *SamlConfigurationXO) SetGroupsAttribute(v string) {
	o.GroupsAttribute = &v
}

// GetValidateResponseSignature returns the ValidateResponseSignature field value if set, zero value otherwise.
func (o *SamlConfigurationXO) GetValidateResponseSignature() bool {
	if o == nil || IsNil(o.ValidateResponseSignature) {
		var ret bool
		return ret
	}
	return *o.ValidateResponseSignature
}

// GetValidateResponseSignatureOk returns a tuple with the ValidateResponseSignature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlConfigurationXO) GetValidateResponseSignatureOk() (*bool, bool) {
	if o == nil || IsNil(o.ValidateResponseSignature) {
		return nil, false
	}
	return o.ValidateResponseSignature, true
}

// HasValidateResponseSignature returns a boolean if a field has been set.
func (o *SamlConfigurationXO) HasValidateResponseSignature() bool {
	if o != nil && !IsNil(o.ValidateResponseSignature) {
		return true
	}

	return false
}

// SetValidateResponseSignature gets a reference to the given bool and assigns it to the ValidateResponseSignature field.
func (o *SamlConfigurationXO) SetValidateResponseSignature(v bool) {
	o.ValidateResponseSignature = &v
}

// GetValidateAssertionSignature returns the ValidateAssertionSignature field value if set, zero value otherwise.
func (o *SamlConfigurationXO) GetValidateAssertionSignature() bool {
	if o == nil || IsNil(o.ValidateAssertionSignature) {
		var ret bool
		return ret
	}
	return *o.ValidateAssertionSignature
}

// GetValidateAssertionSignatureOk returns a tuple with the ValidateAssertionSignature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlConfigurationXO) GetValidateAssertionSignatureOk() (*bool, bool) {
	if o == nil || IsNil(o.ValidateAssertionSignature) {
		return nil, false
	}
	return o.ValidateAssertionSignature, true
}

// HasValidateAssertionSignature returns a boolean if a field has been set.
func (o *SamlConfigurationXO) HasValidateAssertionSignature() bool {
	if o != nil && !IsNil(o.ValidateAssertionSignature) {
		return true
	}

	return false
}

// SetValidateAssertionSignature gets a reference to the given bool and assigns it to the ValidateAssertionSignature field.
func (o *SamlConfigurationXO) SetValidateAssertionSignature(v bool) {
	o.ValidateAssertionSignature = &v
}

func (o SamlConfigurationXO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SamlConfigurationXO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EntityId) {
		toSerialize["entityId"] = o.EntityId
	}
	toSerialize["idpMetadata"] = o.IdpMetadata
	toSerialize["usernameAttribute"] = o.UsernameAttribute
	if !IsNil(o.FirstNameAttribute) {
		toSerialize["firstNameAttribute"] = o.FirstNameAttribute
	}
	if !IsNil(o.LastNameAttribute) {
		toSerialize["lastNameAttribute"] = o.LastNameAttribute
	}
	if !IsNil(o.EmailAttribute) {
		toSerialize["emailAttribute"] = o.EmailAttribute
	}
	if !IsNil(o.GroupsAttribute) {
		toSerialize["groupsAttribute"] = o.GroupsAttribute
	}
	if !IsNil(o.ValidateResponseSignature) {
		toSerialize["validateResponseSignature"] = o.ValidateResponseSignature
	}
	if !IsNil(o.ValidateAssertionSignature) {
		toSerialize["validateAssertionSignature"] = o.ValidateAssertionSignature
	}
	return toSerialize, nil
}

func (o *SamlConfigurationXO) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"idpMetadata",
		"usernameAttribute",
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

	varSamlConfigurationXO := _SamlConfigurationXO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSamlConfigurationXO)

	if err != nil {
		return err
	}

	*o = SamlConfigurationXO(varSamlConfigurationXO)

	return err
}

type NullableSamlConfigurationXO struct {
	value *SamlConfigurationXO
	isSet bool
}

func (v NullableSamlConfigurationXO) Get() *SamlConfigurationXO {
	return v.value
}

func (v *NullableSamlConfigurationXO) Set(val *SamlConfigurationXO) {
	v.value = val
	v.isSet = true
}

func (v NullableSamlConfigurationXO) IsSet() bool {
	return v.isSet
}

func (v *NullableSamlConfigurationXO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSamlConfigurationXO(val *SamlConfigurationXO) *NullableSamlConfigurationXO {
	return &NullableSamlConfigurationXO{value: val, isSet: true}
}

func (v NullableSamlConfigurationXO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSamlConfigurationXO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

