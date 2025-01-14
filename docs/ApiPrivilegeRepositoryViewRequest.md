# ApiPrivilegeRepositoryViewRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the privilege.  This value cannot be changed. | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Actions** | Pointer to **[]string** | A collection of actions to associate with the privilege, using BREAD syntax (browse,read,edit,add,delete,all) as well as &#39;run&#39; for script privileges. | [optional] 
**Format** | Pointer to **string** | The repository format (i.e &#39;nuget&#39;, &#39;npm&#39;) this privilege will grant access to (or * for all). | [optional] 
**Repository** | Pointer to **string** | The name of the repository this privilege will grant access to (or * for all). | [optional] 

## Methods

### NewApiPrivilegeRepositoryViewRequest

`func NewApiPrivilegeRepositoryViewRequest() *ApiPrivilegeRepositoryViewRequest`

NewApiPrivilegeRepositoryViewRequest instantiates a new ApiPrivilegeRepositoryViewRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiPrivilegeRepositoryViewRequestWithDefaults

`func NewApiPrivilegeRepositoryViewRequestWithDefaults() *ApiPrivilegeRepositoryViewRequest`

NewApiPrivilegeRepositoryViewRequestWithDefaults instantiates a new ApiPrivilegeRepositoryViewRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiPrivilegeRepositoryViewRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiPrivilegeRepositoryViewRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiPrivilegeRepositoryViewRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiPrivilegeRepositoryViewRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ApiPrivilegeRepositoryViewRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiPrivilegeRepositoryViewRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiPrivilegeRepositoryViewRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiPrivilegeRepositoryViewRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetActions

`func (o *ApiPrivilegeRepositoryViewRequest) GetActions() []string`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ApiPrivilegeRepositoryViewRequest) GetActionsOk() (*[]string, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ApiPrivilegeRepositoryViewRequest) SetActions(v []string)`

SetActions sets Actions field to given value.

### HasActions

`func (o *ApiPrivilegeRepositoryViewRequest) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetFormat

`func (o *ApiPrivilegeRepositoryViewRequest) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *ApiPrivilegeRepositoryViewRequest) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *ApiPrivilegeRepositoryViewRequest) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *ApiPrivilegeRepositoryViewRequest) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetRepository

`func (o *ApiPrivilegeRepositoryViewRequest) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *ApiPrivilegeRepositoryViewRequest) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *ApiPrivilegeRepositoryViewRequest) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *ApiPrivilegeRepositoryViewRequest) HasRepository() bool`

HasRepository returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


