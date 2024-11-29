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

// checks if the ApiEmailConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiEmailConfiguration{}

// ApiEmailConfiguration struct for ApiEmailConfiguration
type ApiEmailConfiguration struct {
	Enabled *bool `json:"enabled,omitempty"`
	Host *string `json:"host,omitempty"`
	Port int32 `json:"port"`
	Username *string `json:"username,omitempty"`
	Password *string `json:"password,omitempty"`
	FromAddress *string `json:"fromAddress,omitempty"`
	// A prefix to add to all email subjects to aid in identifying automated emails
	SubjectPrefix *string `json:"subjectPrefix,omitempty"`
	// Enable STARTTLS Support for Insecure Connections
	StartTlsEnabled *bool `json:"startTlsEnabled,omitempty"`
	// Require STARTTLS Support
	StartTlsRequired *bool `json:"startTlsRequired,omitempty"`
	// Enable SSL/TLS Encryption upon Connection
	SslOnConnectEnabled *bool `json:"sslOnConnectEnabled,omitempty"`
	// Verify the server certificate when using TLS or SSL
	SslServerIdentityCheckEnabled *bool `json:"sslServerIdentityCheckEnabled,omitempty"`
	// Use the Nexus Repository Manager's certificate truststore
	NexusTrustStoreEnabled *bool `json:"nexusTrustStoreEnabled,omitempty"`
}

type _ApiEmailConfiguration ApiEmailConfiguration

// NewApiEmailConfiguration instantiates a new ApiEmailConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiEmailConfiguration(port int32) *ApiEmailConfiguration {
	this := ApiEmailConfiguration{}
	this.Port = port
	return &this
}

// NewApiEmailConfigurationWithDefaults instantiates a new ApiEmailConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiEmailConfigurationWithDefaults() *ApiEmailConfiguration {
	this := ApiEmailConfiguration{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ApiEmailConfiguration) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiEmailConfiguration) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ApiEmailConfiguration) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ApiEmailConfiguration) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *ApiEmailConfiguration) GetHost() string {
	if o == nil || IsNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiEmailConfiguration) GetHostOk() (*string, bool) {
	if o == nil || IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *ApiEmailConfiguration) HasHost() bool {
	if o != nil && !IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *ApiEmailConfiguration) SetHost(v string) {
	o.Host = &v
}

// GetPort returns the Port field value
func (o *ApiEmailConfiguration) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *ApiEmailConfiguration) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *ApiEmailConfiguration) SetPort(v int32) {
	o.Port = v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *ApiEmailConfiguration) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiEmailConfiguration) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *ApiEmailConfiguration) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *ApiEmailConfiguration) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *ApiEmailConfiguration) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiEmailConfiguration) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *ApiEmailConfiguration) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *ApiEmailConfiguration) SetPassword(v string) {
	o.Password = &v
}

// GetFromAddress returns the FromAddress field value if set, zero value otherwise.
func (o *ApiEmailConfiguration) GetFromAddress() string {
	if o == nil || IsNil(o.FromAddress) {
		var ret string
		return ret
	}
	return *o.FromAddress
}

// GetFromAddressOk returns a tuple with the FromAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiEmailConfiguration) GetFromAddressOk() (*string, bool) {
	if o == nil || IsNil(o.FromAddress) {
		return nil, false
	}
	return o.FromAddress, true
}

// HasFromAddress returns a boolean if a field has been set.
func (o *ApiEmailConfiguration) HasFromAddress() bool {
	if o != nil && !IsNil(o.FromAddress) {
		return true
	}

	return false
}

// SetFromAddress gets a reference to the given string and assigns it to the FromAddress field.
func (o *ApiEmailConfiguration) SetFromAddress(v string) {
	o.FromAddress = &v
}

// GetSubjectPrefix returns the SubjectPrefix field value if set, zero value otherwise.
func (o *ApiEmailConfiguration) GetSubjectPrefix() string {
	if o == nil || IsNil(o.SubjectPrefix) {
		var ret string
		return ret
	}
	return *o.SubjectPrefix
}

// GetSubjectPrefixOk returns a tuple with the SubjectPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiEmailConfiguration) GetSubjectPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.SubjectPrefix) {
		return nil, false
	}
	return o.SubjectPrefix, true
}

// HasSubjectPrefix returns a boolean if a field has been set.
func (o *ApiEmailConfiguration) HasSubjectPrefix() bool {
	if o != nil && !IsNil(o.SubjectPrefix) {
		return true
	}

	return false
}

// SetSubjectPrefix gets a reference to the given string and assigns it to the SubjectPrefix field.
func (o *ApiEmailConfiguration) SetSubjectPrefix(v string) {
	o.SubjectPrefix = &v
}

// GetStartTlsEnabled returns the StartTlsEnabled field value if set, zero value otherwise.
func (o *ApiEmailConfiguration) GetStartTlsEnabled() bool {
	if o == nil || IsNil(o.StartTlsEnabled) {
		var ret bool
		return ret
	}
	return *o.StartTlsEnabled
}

// GetStartTlsEnabledOk returns a tuple with the StartTlsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiEmailConfiguration) GetStartTlsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.StartTlsEnabled) {
		return nil, false
	}
	return o.StartTlsEnabled, true
}

// HasStartTlsEnabled returns a boolean if a field has been set.
func (o *ApiEmailConfiguration) HasStartTlsEnabled() bool {
	if o != nil && !IsNil(o.StartTlsEnabled) {
		return true
	}

	return false
}

// SetStartTlsEnabled gets a reference to the given bool and assigns it to the StartTlsEnabled field.
func (o *ApiEmailConfiguration) SetStartTlsEnabled(v bool) {
	o.StartTlsEnabled = &v
}

// GetStartTlsRequired returns the StartTlsRequired field value if set, zero value otherwise.
func (o *ApiEmailConfiguration) GetStartTlsRequired() bool {
	if o == nil || IsNil(o.StartTlsRequired) {
		var ret bool
		return ret
	}
	return *o.StartTlsRequired
}

// GetStartTlsRequiredOk returns a tuple with the StartTlsRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiEmailConfiguration) GetStartTlsRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.StartTlsRequired) {
		return nil, false
	}
	return o.StartTlsRequired, true
}

// HasStartTlsRequired returns a boolean if a field has been set.
func (o *ApiEmailConfiguration) HasStartTlsRequired() bool {
	if o != nil && !IsNil(o.StartTlsRequired) {
		return true
	}

	return false
}

// SetStartTlsRequired gets a reference to the given bool and assigns it to the StartTlsRequired field.
func (o *ApiEmailConfiguration) SetStartTlsRequired(v bool) {
	o.StartTlsRequired = &v
}

// GetSslOnConnectEnabled returns the SslOnConnectEnabled field value if set, zero value otherwise.
func (o *ApiEmailConfiguration) GetSslOnConnectEnabled() bool {
	if o == nil || IsNil(o.SslOnConnectEnabled) {
		var ret bool
		return ret
	}
	return *o.SslOnConnectEnabled
}

// GetSslOnConnectEnabledOk returns a tuple with the SslOnConnectEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiEmailConfiguration) GetSslOnConnectEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.SslOnConnectEnabled) {
		return nil, false
	}
	return o.SslOnConnectEnabled, true
}

// HasSslOnConnectEnabled returns a boolean if a field has been set.
func (o *ApiEmailConfiguration) HasSslOnConnectEnabled() bool {
	if o != nil && !IsNil(o.SslOnConnectEnabled) {
		return true
	}

	return false
}

// SetSslOnConnectEnabled gets a reference to the given bool and assigns it to the SslOnConnectEnabled field.
func (o *ApiEmailConfiguration) SetSslOnConnectEnabled(v bool) {
	o.SslOnConnectEnabled = &v
}

// GetSslServerIdentityCheckEnabled returns the SslServerIdentityCheckEnabled field value if set, zero value otherwise.
func (o *ApiEmailConfiguration) GetSslServerIdentityCheckEnabled() bool {
	if o == nil || IsNil(o.SslServerIdentityCheckEnabled) {
		var ret bool
		return ret
	}
	return *o.SslServerIdentityCheckEnabled
}

// GetSslServerIdentityCheckEnabledOk returns a tuple with the SslServerIdentityCheckEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiEmailConfiguration) GetSslServerIdentityCheckEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.SslServerIdentityCheckEnabled) {
		return nil, false
	}
	return o.SslServerIdentityCheckEnabled, true
}

// HasSslServerIdentityCheckEnabled returns a boolean if a field has been set.
func (o *ApiEmailConfiguration) HasSslServerIdentityCheckEnabled() bool {
	if o != nil && !IsNil(o.SslServerIdentityCheckEnabled) {
		return true
	}

	return false
}

// SetSslServerIdentityCheckEnabled gets a reference to the given bool and assigns it to the SslServerIdentityCheckEnabled field.
func (o *ApiEmailConfiguration) SetSslServerIdentityCheckEnabled(v bool) {
	o.SslServerIdentityCheckEnabled = &v
}

// GetNexusTrustStoreEnabled returns the NexusTrustStoreEnabled field value if set, zero value otherwise.
func (o *ApiEmailConfiguration) GetNexusTrustStoreEnabled() bool {
	if o == nil || IsNil(o.NexusTrustStoreEnabled) {
		var ret bool
		return ret
	}
	return *o.NexusTrustStoreEnabled
}

// GetNexusTrustStoreEnabledOk returns a tuple with the NexusTrustStoreEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiEmailConfiguration) GetNexusTrustStoreEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.NexusTrustStoreEnabled) {
		return nil, false
	}
	return o.NexusTrustStoreEnabled, true
}

// HasNexusTrustStoreEnabled returns a boolean if a field has been set.
func (o *ApiEmailConfiguration) HasNexusTrustStoreEnabled() bool {
	if o != nil && !IsNil(o.NexusTrustStoreEnabled) {
		return true
	}

	return false
}

// SetNexusTrustStoreEnabled gets a reference to the given bool and assigns it to the NexusTrustStoreEnabled field.
func (o *ApiEmailConfiguration) SetNexusTrustStoreEnabled(v bool) {
	o.NexusTrustStoreEnabled = &v
}

func (o ApiEmailConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiEmailConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	toSerialize["port"] = o.Port
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.FromAddress) {
		toSerialize["fromAddress"] = o.FromAddress
	}
	if !IsNil(o.SubjectPrefix) {
		toSerialize["subjectPrefix"] = o.SubjectPrefix
	}
	if !IsNil(o.StartTlsEnabled) {
		toSerialize["startTlsEnabled"] = o.StartTlsEnabled
	}
	if !IsNil(o.StartTlsRequired) {
		toSerialize["startTlsRequired"] = o.StartTlsRequired
	}
	if !IsNil(o.SslOnConnectEnabled) {
		toSerialize["sslOnConnectEnabled"] = o.SslOnConnectEnabled
	}
	if !IsNil(o.SslServerIdentityCheckEnabled) {
		toSerialize["sslServerIdentityCheckEnabled"] = o.SslServerIdentityCheckEnabled
	}
	if !IsNil(o.NexusTrustStoreEnabled) {
		toSerialize["nexusTrustStoreEnabled"] = o.NexusTrustStoreEnabled
	}
	return toSerialize, nil
}

func (o *ApiEmailConfiguration) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"port",
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

	varApiEmailConfiguration := _ApiEmailConfiguration{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiEmailConfiguration)

	if err != nil {
		return err
	}

	*o = ApiEmailConfiguration(varApiEmailConfiguration)

	return err
}

type NullableApiEmailConfiguration struct {
	value *ApiEmailConfiguration
	isSet bool
}

func (v NullableApiEmailConfiguration) Get() *ApiEmailConfiguration {
	return v.value
}

func (v *NullableApiEmailConfiguration) Set(val *ApiEmailConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableApiEmailConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableApiEmailConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiEmailConfiguration(val *ApiEmailConfiguration) *NullableApiEmailConfiguration {
	return &NullableApiEmailConfiguration{value: val, isSet: true}
}

func (v NullableApiEmailConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiEmailConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

