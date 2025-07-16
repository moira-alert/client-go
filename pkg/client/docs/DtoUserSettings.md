# DtoUserSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthEnabled** | Pointer to **bool** |  | [optional] 
**Contacts** | Pointer to [**[]MoiraContactData**](MoiraContactData.md) |  | [optional] 
**Login** | Pointer to **string** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**Subscriptions** | Pointer to [**[]MoiraSubscriptionData**](MoiraSubscriptionData.md) |  | [optional] 

## Methods

### NewDtoUserSettings

`func NewDtoUserSettings() *DtoUserSettings`

NewDtoUserSettings instantiates a new DtoUserSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoUserSettingsWithDefaults

`func NewDtoUserSettingsWithDefaults() *DtoUserSettings`

NewDtoUserSettingsWithDefaults instantiates a new DtoUserSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthEnabled

`func (o *DtoUserSettings) GetAuthEnabled() bool`

GetAuthEnabled returns the AuthEnabled field if non-nil, zero value otherwise.

### GetAuthEnabledOk

`func (o *DtoUserSettings) GetAuthEnabledOk() (*bool, bool)`

GetAuthEnabledOk returns a tuple with the AuthEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthEnabled

`func (o *DtoUserSettings) SetAuthEnabled(v bool)`

SetAuthEnabled sets AuthEnabled field to given value.

### HasAuthEnabled

`func (o *DtoUserSettings) HasAuthEnabled() bool`

HasAuthEnabled returns a boolean if a field has been set.

### GetContacts

`func (o *DtoUserSettings) GetContacts() []MoiraContactData`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *DtoUserSettings) GetContactsOk() (*[]MoiraContactData, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *DtoUserSettings) SetContacts(v []MoiraContactData)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *DtoUserSettings) HasContacts() bool`

HasContacts returns a boolean if a field has been set.

### GetLogin

`func (o *DtoUserSettings) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *DtoUserSettings) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *DtoUserSettings) SetLogin(v string)`

SetLogin sets Login field to given value.

### HasLogin

`func (o *DtoUserSettings) HasLogin() bool`

HasLogin returns a boolean if a field has been set.

### GetRole

`func (o *DtoUserSettings) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *DtoUserSettings) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *DtoUserSettings) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *DtoUserSettings) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetSubscriptions

`func (o *DtoUserSettings) GetSubscriptions() []MoiraSubscriptionData`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *DtoUserSettings) GetSubscriptionsOk() (*[]MoiraSubscriptionData, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *DtoUserSettings) SetSubscriptions(v []MoiraSubscriptionData)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *DtoUserSettings) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


