# MoiraScheduledNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contact** | Pointer to [**MoiraContactData**](MoiraContactData.md) |  | [optional] 
**CreatedAt** | Pointer to **int64** |  | [optional] 
**Event** | Pointer to [**MoiraNotificationEvent**](MoiraNotificationEvent.md) |  | [optional] 
**Plotting** | Pointer to [**MoiraPlottingData**](MoiraPlottingData.md) |  | [optional] 
**SendFail** | Pointer to **int32** |  | [optional] 
**Throttled** | Pointer to **bool** |  | [optional] 
**Timestamp** | Pointer to **int64** |  | [optional] 
**Trigger** | Pointer to [**MoiraTriggerData**](MoiraTriggerData.md) |  | [optional] 

## Methods

### NewMoiraScheduledNotification

`func NewMoiraScheduledNotification() *MoiraScheduledNotification`

NewMoiraScheduledNotification instantiates a new MoiraScheduledNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoiraScheduledNotificationWithDefaults

`func NewMoiraScheduledNotificationWithDefaults() *MoiraScheduledNotification`

NewMoiraScheduledNotificationWithDefaults instantiates a new MoiraScheduledNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContact

`func (o *MoiraScheduledNotification) GetContact() MoiraContactData`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *MoiraScheduledNotification) GetContactOk() (*MoiraContactData, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *MoiraScheduledNotification) SetContact(v MoiraContactData)`

SetContact sets Contact field to given value.

### HasContact

`func (o *MoiraScheduledNotification) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MoiraScheduledNotification) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MoiraScheduledNotification) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MoiraScheduledNotification) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MoiraScheduledNotification) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEvent

`func (o *MoiraScheduledNotification) GetEvent() MoiraNotificationEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *MoiraScheduledNotification) GetEventOk() (*MoiraNotificationEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *MoiraScheduledNotification) SetEvent(v MoiraNotificationEvent)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *MoiraScheduledNotification) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetPlotting

`func (o *MoiraScheduledNotification) GetPlotting() MoiraPlottingData`

GetPlotting returns the Plotting field if non-nil, zero value otherwise.

### GetPlottingOk

`func (o *MoiraScheduledNotification) GetPlottingOk() (*MoiraPlottingData, bool)`

GetPlottingOk returns a tuple with the Plotting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlotting

`func (o *MoiraScheduledNotification) SetPlotting(v MoiraPlottingData)`

SetPlotting sets Plotting field to given value.

### HasPlotting

`func (o *MoiraScheduledNotification) HasPlotting() bool`

HasPlotting returns a boolean if a field has been set.

### GetSendFail

`func (o *MoiraScheduledNotification) GetSendFail() int32`

GetSendFail returns the SendFail field if non-nil, zero value otherwise.

### GetSendFailOk

`func (o *MoiraScheduledNotification) GetSendFailOk() (*int32, bool)`

GetSendFailOk returns a tuple with the SendFail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendFail

`func (o *MoiraScheduledNotification) SetSendFail(v int32)`

SetSendFail sets SendFail field to given value.

### HasSendFail

`func (o *MoiraScheduledNotification) HasSendFail() bool`

HasSendFail returns a boolean if a field has been set.

### GetThrottled

`func (o *MoiraScheduledNotification) GetThrottled() bool`

GetThrottled returns the Throttled field if non-nil, zero value otherwise.

### GetThrottledOk

`func (o *MoiraScheduledNotification) GetThrottledOk() (*bool, bool)`

GetThrottledOk returns a tuple with the Throttled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottled

`func (o *MoiraScheduledNotification) SetThrottled(v bool)`

SetThrottled sets Throttled field to given value.

### HasThrottled

`func (o *MoiraScheduledNotification) HasThrottled() bool`

HasThrottled returns a boolean if a field has been set.

### GetTimestamp

`func (o *MoiraScheduledNotification) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *MoiraScheduledNotification) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *MoiraScheduledNotification) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *MoiraScheduledNotification) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTrigger

`func (o *MoiraScheduledNotification) GetTrigger() MoiraTriggerData`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *MoiraScheduledNotification) GetTriggerOk() (*MoiraTriggerData, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *MoiraScheduledNotification) SetTrigger(v MoiraTriggerData)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *MoiraScheduledNotification) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


