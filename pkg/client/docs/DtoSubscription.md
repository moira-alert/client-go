# DtoSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnyTags** | Pointer to **bool** |  | [optional] 
**Contacts** | Pointer to **[]string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IgnoreRecoverings** | Pointer to **bool** |  | [optional] 
**IgnoreWarnings** | Pointer to **bool** |  | [optional] 
**Plotting** | Pointer to [**MoiraPlottingData**](MoiraPlottingData.md) |  | [optional] 
**Sched** | Pointer to [**NullableMoiraScheduleData**](MoiraScheduleData.md) |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**TeamId** | Pointer to **string** |  | [optional] 
**Throttling** | Pointer to **bool** |  | [optional] 
**User** | Pointer to **string** |  | [optional] 

## Methods

### NewDtoSubscription

`func NewDtoSubscription() *DtoSubscription`

NewDtoSubscription instantiates a new DtoSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoSubscriptionWithDefaults

`func NewDtoSubscriptionWithDefaults() *DtoSubscription`

NewDtoSubscriptionWithDefaults instantiates a new DtoSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnyTags

`func (o *DtoSubscription) GetAnyTags() bool`

GetAnyTags returns the AnyTags field if non-nil, zero value otherwise.

### GetAnyTagsOk

`func (o *DtoSubscription) GetAnyTagsOk() (*bool, bool)`

GetAnyTagsOk returns a tuple with the AnyTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyTags

`func (o *DtoSubscription) SetAnyTags(v bool)`

SetAnyTags sets AnyTags field to given value.

### HasAnyTags

`func (o *DtoSubscription) HasAnyTags() bool`

HasAnyTags returns a boolean if a field has been set.

### GetContacts

`func (o *DtoSubscription) GetContacts() []string`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *DtoSubscription) GetContactsOk() (*[]string, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *DtoSubscription) SetContacts(v []string)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *DtoSubscription) HasContacts() bool`

HasContacts returns a boolean if a field has been set.

### GetEnabled

`func (o *DtoSubscription) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DtoSubscription) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DtoSubscription) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *DtoSubscription) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetId

`func (o *DtoSubscription) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DtoSubscription) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DtoSubscription) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DtoSubscription) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIgnoreRecoverings

`func (o *DtoSubscription) GetIgnoreRecoverings() bool`

GetIgnoreRecoverings returns the IgnoreRecoverings field if non-nil, zero value otherwise.

### GetIgnoreRecoveringsOk

`func (o *DtoSubscription) GetIgnoreRecoveringsOk() (*bool, bool)`

GetIgnoreRecoveringsOk returns a tuple with the IgnoreRecoverings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreRecoverings

`func (o *DtoSubscription) SetIgnoreRecoverings(v bool)`

SetIgnoreRecoverings sets IgnoreRecoverings field to given value.

### HasIgnoreRecoverings

`func (o *DtoSubscription) HasIgnoreRecoverings() bool`

HasIgnoreRecoverings returns a boolean if a field has been set.

### GetIgnoreWarnings

`func (o *DtoSubscription) GetIgnoreWarnings() bool`

GetIgnoreWarnings returns the IgnoreWarnings field if non-nil, zero value otherwise.

### GetIgnoreWarningsOk

`func (o *DtoSubscription) GetIgnoreWarningsOk() (*bool, bool)`

GetIgnoreWarningsOk returns a tuple with the IgnoreWarnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreWarnings

`func (o *DtoSubscription) SetIgnoreWarnings(v bool)`

SetIgnoreWarnings sets IgnoreWarnings field to given value.

### HasIgnoreWarnings

`func (o *DtoSubscription) HasIgnoreWarnings() bool`

HasIgnoreWarnings returns a boolean if a field has been set.

### GetPlotting

`func (o *DtoSubscription) GetPlotting() MoiraPlottingData`

GetPlotting returns the Plotting field if non-nil, zero value otherwise.

### GetPlottingOk

`func (o *DtoSubscription) GetPlottingOk() (*MoiraPlottingData, bool)`

GetPlottingOk returns a tuple with the Plotting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlotting

`func (o *DtoSubscription) SetPlotting(v MoiraPlottingData)`

SetPlotting sets Plotting field to given value.

### HasPlotting

`func (o *DtoSubscription) HasPlotting() bool`

HasPlotting returns a boolean if a field has been set.

### GetSched

`func (o *DtoSubscription) GetSched() MoiraScheduleData`

GetSched returns the Sched field if non-nil, zero value otherwise.

### GetSchedOk

`func (o *DtoSubscription) GetSchedOk() (*MoiraScheduleData, bool)`

GetSchedOk returns a tuple with the Sched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSched

`func (o *DtoSubscription) SetSched(v MoiraScheduleData)`

SetSched sets Sched field to given value.

### HasSched

`func (o *DtoSubscription) HasSched() bool`

HasSched returns a boolean if a field has been set.

### SetSchedNil

`func (o *DtoSubscription) SetSchedNil(b bool)`

 SetSchedNil sets the value for Sched to be an explicit nil

### UnsetSched
`func (o *DtoSubscription) UnsetSched()`

UnsetSched ensures that no value is present for Sched, not even an explicit nil
### GetTags

`func (o *DtoSubscription) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DtoSubscription) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DtoSubscription) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DtoSubscription) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTeamId

`func (o *DtoSubscription) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *DtoSubscription) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *DtoSubscription) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *DtoSubscription) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.

### GetThrottling

`func (o *DtoSubscription) GetThrottling() bool`

GetThrottling returns the Throttling field if non-nil, zero value otherwise.

### GetThrottlingOk

`func (o *DtoSubscription) GetThrottlingOk() (*bool, bool)`

GetThrottlingOk returns a tuple with the Throttling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottling

`func (o *DtoSubscription) SetThrottling(v bool)`

SetThrottling sets Throttling field to given value.

### HasThrottling

`func (o *DtoSubscription) HasThrottling() bool`

HasThrottling returns a boolean if a field has been set.

### GetUser

`func (o *DtoSubscription) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *DtoSubscription) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *DtoSubscription) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *DtoSubscription) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


