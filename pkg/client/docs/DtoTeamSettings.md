# DtoTeamSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contacts** | Pointer to [**[]DtoTeamContact**](DtoTeamContact.md) |  | [optional] 
**Subscriptions** | Pointer to [**[]MoiraSubscriptionData**](MoiraSubscriptionData.md) |  | [optional] 
**TeamId** | Pointer to **string** |  | [optional] 

## Methods

### NewDtoTeamSettings

`func NewDtoTeamSettings() *DtoTeamSettings`

NewDtoTeamSettings instantiates a new DtoTeamSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoTeamSettingsWithDefaults

`func NewDtoTeamSettingsWithDefaults() *DtoTeamSettings`

NewDtoTeamSettingsWithDefaults instantiates a new DtoTeamSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContacts

`func (o *DtoTeamSettings) GetContacts() []DtoTeamContact`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *DtoTeamSettings) GetContactsOk() (*[]DtoTeamContact, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *DtoTeamSettings) SetContacts(v []DtoTeamContact)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *DtoTeamSettings) HasContacts() bool`

HasContacts returns a boolean if a field has been set.

### GetSubscriptions

`func (o *DtoTeamSettings) GetSubscriptions() []MoiraSubscriptionData`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *DtoTeamSettings) GetSubscriptionsOk() (*[]MoiraSubscriptionData, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *DtoTeamSettings) SetSubscriptions(v []MoiraSubscriptionData)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *DtoTeamSettings) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.

### GetTeamId

`func (o *DtoTeamSettings) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *DtoTeamSettings) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *DtoTeamSettings) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *DtoTeamSettings) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


