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

// checks if the ContentSelectorApiUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContentSelectorApiUpdateRequest{}

// ContentSelectorApiUpdateRequest struct for ContentSelectorApiUpdateRequest
type ContentSelectorApiUpdateRequest struct {
	// An optional description of this content selector
	Description *string `json:"description,omitempty"`
	// The expression used to identify content
	Expression *string `json:"expression,omitempty"`
}

// NewContentSelectorApiUpdateRequest instantiates a new ContentSelectorApiUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentSelectorApiUpdateRequest() *ContentSelectorApiUpdateRequest {
	this := ContentSelectorApiUpdateRequest{}
	return &this
}

// NewContentSelectorApiUpdateRequestWithDefaults instantiates a new ContentSelectorApiUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentSelectorApiUpdateRequestWithDefaults() *ContentSelectorApiUpdateRequest {
	this := ContentSelectorApiUpdateRequest{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ContentSelectorApiUpdateRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentSelectorApiUpdateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ContentSelectorApiUpdateRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ContentSelectorApiUpdateRequest) SetDescription(v string) {
	o.Description = &v
}

// GetExpression returns the Expression field value if set, zero value otherwise.
func (o *ContentSelectorApiUpdateRequest) GetExpression() string {
	if o == nil || IsNil(o.Expression) {
		var ret string
		return ret
	}
	return *o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentSelectorApiUpdateRequest) GetExpressionOk() (*string, bool) {
	if o == nil || IsNil(o.Expression) {
		return nil, false
	}
	return o.Expression, true
}

// HasExpression returns a boolean if a field has been set.
func (o *ContentSelectorApiUpdateRequest) HasExpression() bool {
	if o != nil && !IsNil(o.Expression) {
		return true
	}

	return false
}

// SetExpression gets a reference to the given string and assigns it to the Expression field.
func (o *ContentSelectorApiUpdateRequest) SetExpression(v string) {
	o.Expression = &v
}

func (o ContentSelectorApiUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContentSelectorApiUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Expression) {
		toSerialize["expression"] = o.Expression
	}
	return toSerialize, nil
}

type NullableContentSelectorApiUpdateRequest struct {
	value *ContentSelectorApiUpdateRequest
	isSet bool
}

func (v NullableContentSelectorApiUpdateRequest) Get() *ContentSelectorApiUpdateRequest {
	return v.value
}

func (v *NullableContentSelectorApiUpdateRequest) Set(val *ContentSelectorApiUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableContentSelectorApiUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableContentSelectorApiUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentSelectorApiUpdateRequest(val *ContentSelectorApiUpdateRequest) *NullableContentSelectorApiUpdateRequest {
	return &NullableContentSelectorApiUpdateRequest{value: val, isSet: true}
}

func (v NullableContentSelectorApiUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentSelectorApiUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


