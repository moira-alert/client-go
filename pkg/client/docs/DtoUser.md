# DtoUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthEnabled** | Pointer to **bool** |  | [optional] 
**Login** | Pointer to **string** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 

## Methods

### NewDtoUser

`func NewDtoUser() *DtoUser`

NewDtoUser instantiates a new DtoUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoUserWithDefaults

`func NewDtoUserWithDefaults() *DtoUser`

NewDtoUserWithDefaults instantiates a new DtoUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthEnabled

`func (o *DtoUser) GetAuthEnabled() bool`

GetAuthEnabled returns the AuthEnabled field if non-nil, zero value otherwise.

### GetAuthEnabledOk

`func (o *DtoUser) GetAuthEnabledOk() (*bool, bool)`

GetAuthEnabledOk returns a tuple with the AuthEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthEnabled

`func (o *DtoUser) SetAuthEnabled(v bool)`

SetAuthEnabled sets AuthEnabled field to given value.

### HasAuthEnabled

`func (o *DtoUser) HasAuthEnabled() bool`

HasAuthEnabled returns a boolean if a field has been set.

### GetLogin

`func (o *DtoUser) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *DtoUser) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *DtoUser) SetLogin(v string)`

SetLogin sets Login field to given value.

### HasLogin

`func (o *DtoUser) HasLogin() bool`

HasLogin returns a boolean if a field has been set.

### GetRole

`func (o *DtoUser) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *DtoUser) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *DtoUser) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *DtoUser) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


