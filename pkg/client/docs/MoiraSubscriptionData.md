# MoiraSubscriptionData

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

### NewMoiraSubscriptionData

`func NewMoiraSubscriptionData() *MoiraSubscriptionData`

NewMoiraSubscriptionData instantiates a new MoiraSubscriptionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoiraSubscriptionDataWithDefaults

`func NewMoiraSubscriptionDataWithDefaults() *MoiraSubscriptionData`

NewMoiraSubscriptionDataWithDefaults instantiates a new MoiraSubscriptionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnyTags

`func (o *MoiraSubscriptionData) GetAnyTags() bool`

GetAnyTags returns the AnyTags field if non-nil, zero value otherwise.

### GetAnyTagsOk

`func (o *MoiraSubscriptionData) GetAnyTagsOk() (*bool, bool)`

GetAnyTagsOk returns a tuple with the AnyTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyTags

`func (o *MoiraSubscriptionData) SetAnyTags(v bool)`

SetAnyTags sets AnyTags field to given value.

### HasAnyTags

`func (o *MoiraSubscriptionData) HasAnyTags() bool`

HasAnyTags returns a boolean if a field has been set.

### GetContacts

`func (o *MoiraSubscriptionData) GetContacts() []string`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *MoiraSubscriptionData) GetContactsOk() (*[]string, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *MoiraSubscriptionData) SetContacts(v []string)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *MoiraSubscriptionData) HasContacts() bool`

HasContacts returns a boolean if a field has been set.

### GetEnabled

`func (o *MoiraSubscriptionData) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MoiraSubscriptionData) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MoiraSubscriptionData) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *MoiraSubscriptionData) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetId

`func (o *MoiraSubscriptionData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MoiraSubscriptionData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MoiraSubscriptionData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MoiraSubscriptionData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIgnoreRecoverings

`func (o *MoiraSubscriptionData) GetIgnoreRecoverings() bool`

GetIgnoreRecoverings returns the IgnoreRecoverings field if non-nil, zero value otherwise.

### GetIgnoreRecoveringsOk

`func (o *MoiraSubscriptionData) GetIgnoreRecoveringsOk() (*bool, bool)`

GetIgnoreRecoveringsOk returns a tuple with the IgnoreRecoverings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreRecoverings

`func (o *MoiraSubscriptionData) SetIgnoreRecoverings(v bool)`

SetIgnoreRecoverings sets IgnoreRecoverings field to given value.

### HasIgnoreRecoverings

`func (o *MoiraSubscriptionData) HasIgnoreRecoverings() bool`

HasIgnoreRecoverings returns a boolean if a field has been set.

### GetIgnoreWarnings

`func (o *MoiraSubscriptionData) GetIgnoreWarnings() bool`

GetIgnoreWarnings returns the IgnoreWarnings field if non-nil, zero value otherwise.

### GetIgnoreWarningsOk

`func (o *MoiraSubscriptionData) GetIgnoreWarningsOk() (*bool, bool)`

GetIgnoreWarningsOk returns a tuple with the IgnoreWarnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreWarnings

`func (o *MoiraSubscriptionData) SetIgnoreWarnings(v bool)`

SetIgnoreWarnings sets IgnoreWarnings field to given value.

### HasIgnoreWarnings

`func (o *MoiraSubscriptionData) HasIgnoreWarnings() bool`

HasIgnoreWarnings returns a boolean if a field has been set.

### GetPlotting

`func (o *MoiraSubscriptionData) GetPlotting() MoiraPlottingData`

GetPlotting returns the Plotting field if non-nil, zero value otherwise.

### GetPlottingOk

`func (o *MoiraSubscriptionData) GetPlottingOk() (*MoiraPlottingData, bool)`

GetPlottingOk returns a tuple with the Plotting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlotting

`func (o *MoiraSubscriptionData) SetPlotting(v MoiraPlottingData)`

SetPlotting sets Plotting field to given value.

### HasPlotting

`func (o *MoiraSubscriptionData) HasPlotting() bool`

HasPlotting returns a boolean if a field has been set.

### GetSched

`func (o *MoiraSubscriptionData) GetSched() MoiraScheduleData`

GetSched returns the Sched field if non-nil, zero value otherwise.

### GetSchedOk

`func (o *MoiraSubscriptionData) GetSchedOk() (*MoiraScheduleData, bool)`

GetSchedOk returns a tuple with the Sched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSched

`func (o *MoiraSubscriptionData) SetSched(v MoiraScheduleData)`

SetSched sets Sched field to given value.

### HasSched

`func (o *MoiraSubscriptionData) HasSched() bool`

HasSched returns a boolean if a field has been set.

### SetSchedNil

`func (o *MoiraSubscriptionData) SetSchedNil(b bool)`

 SetSchedNil sets the value for Sched to be an explicit nil

### UnsetSched
`func (o *MoiraSubscriptionData) UnsetSched()`

UnsetSched ensures that no value is present for Sched, not even an explicit nil
### GetTags

`func (o *MoiraSubscriptionData) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MoiraSubscriptionData) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MoiraSubscriptionData) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *MoiraSubscriptionData) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTeamId

`func (o *MoiraSubscriptionData) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *MoiraSubscriptionData) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *MoiraSubscriptionData) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *MoiraSubscriptionData) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.

### GetThrottling

`func (o *MoiraSubscriptionData) GetThrottling() bool`

GetThrottling returns the Throttling field if non-nil, zero value otherwise.

### GetThrottlingOk

`func (o *MoiraSubscriptionData) GetThrottlingOk() (*bool, bool)`

GetThrottlingOk returns a tuple with the Throttling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottling

`func (o *MoiraSubscriptionData) SetThrottling(v bool)`

SetThrottling sets Throttling field to given value.

### HasThrottling

`func (o *MoiraSubscriptionData) HasThrottling() bool`

HasThrottling returns a boolean if a field has been set.

### GetUser

`func (o *MoiraSubscriptionData) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *MoiraSubscriptionData) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *MoiraSubscriptionData) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *MoiraSubscriptionData) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


