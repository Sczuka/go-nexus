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

// checks if the BlobStoreQuotaResultXO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlobStoreQuotaResultXO{}

// BlobStoreQuotaResultXO struct for BlobStoreQuotaResultXO
type BlobStoreQuotaResultXO struct {
	IsViolation *bool `json:"isViolation,omitempty"`
	Message *string `json:"message,omitempty"`
	BlobStoreName *string `json:"blobStoreName,omitempty"`
}

// NewBlobStoreQuotaResultXO instantiates a new BlobStoreQuotaResultXO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlobStoreQuotaResultXO() *BlobStoreQuotaResultXO {
	this := BlobStoreQuotaResultXO{}
	return &this
}

// NewBlobStoreQuotaResultXOWithDefaults instantiates a new BlobStoreQuotaResultXO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlobStoreQuotaResultXOWithDefaults() *BlobStoreQuotaResultXO {
	this := BlobStoreQuotaResultXO{}
	return &this
}

// GetIsViolation returns the IsViolation field value if set, zero value otherwise.
func (o *BlobStoreQuotaResultXO) GetIsViolation() bool {
	if o == nil || IsNil(o.IsViolation) {
		var ret bool
		return ret
	}
	return *o.IsViolation
}

// GetIsViolationOk returns a tuple with the IsViolation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlobStoreQuotaResultXO) GetIsViolationOk() (*bool, bool) {
	if o == nil || IsNil(o.IsViolation) {
		return nil, false
	}
	return o.IsViolation, true
}

// HasIsViolation returns a boolean if a field has been set.
func (o *BlobStoreQuotaResultXO) HasIsViolation() bool {
	if o != nil && !IsNil(o.IsViolation) {
		return true
	}

	return false
}

// SetIsViolation gets a reference to the given bool and assigns it to the IsViolation field.
func (o *BlobStoreQuotaResultXO) SetIsViolation(v bool) {
	o.IsViolation = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *BlobStoreQuotaResultXO) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlobStoreQuotaResultXO) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *BlobStoreQuotaResultXO) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *BlobStoreQuotaResultXO) SetMessage(v string) {
	o.Message = &v
}

// GetBlobStoreName returns the BlobStoreName field value if set, zero value otherwise.
func (o *BlobStoreQuotaResultXO) GetBlobStoreName() string {
	if o == nil || IsNil(o.BlobStoreName) {
		var ret string
		return ret
	}
	return *o.BlobStoreName
}

// GetBlobStoreNameOk returns a tuple with the BlobStoreName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlobStoreQuotaResultXO) GetBlobStoreNameOk() (*string, bool) {
	if o == nil || IsNil(o.BlobStoreName) {
		return nil, false
	}
	return o.BlobStoreName, true
}

// HasBlobStoreName returns a boolean if a field has been set.
func (o *BlobStoreQuotaResultXO) HasBlobStoreName() bool {
	if o != nil && !IsNil(o.BlobStoreName) {
		return true
	}

	return false
}

// SetBlobStoreName gets a reference to the given string and assigns it to the BlobStoreName field.
func (o *BlobStoreQuotaResultXO) SetBlobStoreName(v string) {
	o.BlobStoreName = &v
}

func (o BlobStoreQuotaResultXO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlobStoreQuotaResultXO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsViolation) {
		toSerialize["isViolation"] = o.IsViolation
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.BlobStoreName) {
		toSerialize["blobStoreName"] = o.BlobStoreName
	}
	return toSerialize, nil
}

type NullableBlobStoreQuotaResultXO struct {
	value *BlobStoreQuotaResultXO
	isSet bool
}

func (v NullableBlobStoreQuotaResultXO) Get() *BlobStoreQuotaResultXO {
	return v.value
}

func (v *NullableBlobStoreQuotaResultXO) Set(val *BlobStoreQuotaResultXO) {
	v.value = val
	v.isSet = true
}

func (v NullableBlobStoreQuotaResultXO) IsSet() bool {
	return v.isSet
}

func (v *NullableBlobStoreQuotaResultXO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlobStoreQuotaResultXO(val *BlobStoreQuotaResultXO) *NullableBlobStoreQuotaResultXO {
	return &NullableBlobStoreQuotaResultXO{value: val, isSet: true}
}

func (v NullableBlobStoreQuotaResultXO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlobStoreQuotaResultXO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

