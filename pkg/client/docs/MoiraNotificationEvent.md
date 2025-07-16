# MoiraNotificationEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContactId** | Pointer to **string** |  | [optional] 
**EventMessage** | Pointer to [**NullableMoiraEventInfo**](MoiraEventInfo.md) |  | [optional] 
**Metric** | Pointer to **string** |  | [optional] 
**Msg** | Pointer to **NullableString** |  | [optional] 
**OldState** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**SubId** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **int64** |  | [optional] 
**TriggerEvent** | Pointer to **bool** |  | [optional] 
**TriggerId** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **NullableFloat32** |  | [optional] 
**Values** | Pointer to **map[string]float32** |  | [optional] 

## Methods

### NewMoiraNotificationEvent

`func NewMoiraNotificationEvent() *MoiraNotificationEvent`

NewMoiraNotificationEvent instantiates a new MoiraNotificationEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoiraNotificationEventWithDefaults

`func NewMoiraNotificationEventWithDefaults() *MoiraNotificationEvent`

NewMoiraNotificationEventWithDefaults instantiates a new MoiraNotificationEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContactId

`func (o *MoiraNotificationEvent) GetContactId() string`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *MoiraNotificationEvent) GetContactIdOk() (*string, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *MoiraNotificationEvent) SetContactId(v string)`

SetContactId sets ContactId field to given value.

### HasContactId

`func (o *MoiraNotificationEvent) HasContactId() bool`

HasContactId returns a boolean if a field has been set.

### GetEventMessage

`func (o *MoiraNotificationEvent) GetEventMessage() MoiraEventInfo`

GetEventMessage returns the EventMessage field if non-nil, zero value otherwise.

### GetEventMessageOk

`func (o *MoiraNotificationEvent) GetEventMessageOk() (*MoiraEventInfo, bool)`

GetEventMessageOk returns a tuple with the EventMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventMessage

`func (o *MoiraNotificationEvent) SetEventMessage(v MoiraEventInfo)`

SetEventMessage sets EventMessage field to given value.

### HasEventMessage

`func (o *MoiraNotificationEvent) HasEventMessage() bool`

HasEventMessage returns a boolean if a field has been set.

### SetEventMessageNil

`func (o *MoiraNotificationEvent) SetEventMessageNil(b bool)`

 SetEventMessageNil sets the value for EventMessage to be an explicit nil

### UnsetEventMessage
`func (o *MoiraNotificationEvent) UnsetEventMessage()`

UnsetEventMessage ensures that no value is present for EventMessage, not even an explicit nil
### GetMetric

`func (o *MoiraNotificationEvent) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *MoiraNotificationEvent) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *MoiraNotificationEvent) SetMetric(v string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *MoiraNotificationEvent) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetMsg

`func (o *MoiraNotificationEvent) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *MoiraNotificationEvent) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *MoiraNotificationEvent) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *MoiraNotificationEvent) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### SetMsgNil

`func (o *MoiraNotificationEvent) SetMsgNil(b bool)`

 SetMsgNil sets the value for Msg to be an explicit nil

### UnsetMsg
`func (o *MoiraNotificationEvent) UnsetMsg()`

UnsetMsg ensures that no value is present for Msg, not even an explicit nil
### GetOldState

`func (o *MoiraNotificationEvent) GetOldState() string`

GetOldState returns the OldState field if non-nil, zero value otherwise.

### GetOldStateOk

`func (o *MoiraNotificationEvent) GetOldStateOk() (*string, bool)`

GetOldStateOk returns a tuple with the OldState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldState

`func (o *MoiraNotificationEvent) SetOldState(v string)`

SetOldState sets OldState field to given value.

### HasOldState

`func (o *MoiraNotificationEvent) HasOldState() bool`

HasOldState returns a boolean if a field has been set.

### GetState

`func (o *MoiraNotificationEvent) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MoiraNotificationEvent) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MoiraNotificationEvent) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *MoiraNotificationEvent) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSubId

`func (o *MoiraNotificationEvent) GetSubId() string`

GetSubId returns the SubId field if non-nil, zero value otherwise.

### GetSubIdOk

`func (o *MoiraNotificationEvent) GetSubIdOk() (*string, bool)`

GetSubIdOk returns a tuple with the SubId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubId

`func (o *MoiraNotificationEvent) SetSubId(v string)`

SetSubId sets SubId field to given value.

### HasSubId

`func (o *MoiraNotificationEvent) HasSubId() bool`

HasSubId returns a boolean if a field has been set.

### SetSubIdNil

`func (o *MoiraNotificationEvent) SetSubIdNil(b bool)`

 SetSubIdNil sets the value for SubId to be an explicit nil

### UnsetSubId
`func (o *MoiraNotificationEvent) UnsetSubId()`

UnsetSubId ensures that no value is present for SubId, not even an explicit nil
### GetTimestamp

`func (o *MoiraNotificationEvent) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *MoiraNotificationEvent) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *MoiraNotificationEvent) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *MoiraNotificationEvent) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTriggerEvent

`func (o *MoiraNotificationEvent) GetTriggerEvent() bool`

GetTriggerEvent returns the TriggerEvent field if non-nil, zero value otherwise.

### GetTriggerEventOk

`func (o *MoiraNotificationEvent) GetTriggerEventOk() (*bool, bool)`

GetTriggerEventOk returns a tuple with the TriggerEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerEvent

`func (o *MoiraNotificationEvent) SetTriggerEvent(v bool)`

SetTriggerEvent sets TriggerEvent field to given value.

### HasTriggerEvent

`func (o *MoiraNotificationEvent) HasTriggerEvent() bool`

HasTriggerEvent returns a boolean if a field has been set.

### GetTriggerId

`func (o *MoiraNotificationEvent) GetTriggerId() string`

GetTriggerId returns the TriggerId field if non-nil, zero value otherwise.

### GetTriggerIdOk

`func (o *MoiraNotificationEvent) GetTriggerIdOk() (*string, bool)`

GetTriggerIdOk returns a tuple with the TriggerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerId

`func (o *MoiraNotificationEvent) SetTriggerId(v string)`

SetTriggerId sets TriggerId field to given value.

### HasTriggerId

`func (o *MoiraNotificationEvent) HasTriggerId() bool`

HasTriggerId returns a boolean if a field has been set.

### GetValue

`func (o *MoiraNotificationEvent) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MoiraNotificationEvent) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MoiraNotificationEvent) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *MoiraNotificationEvent) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *MoiraNotificationEvent) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *MoiraNotificationEvent) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetValues

`func (o *MoiraNotificationEvent) GetValues() map[string]float32`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *MoiraNotificationEvent) GetValuesOk() (*map[string]float32, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *MoiraNotificationEvent) SetValues(v map[string]float32)`

SetValues sets Values field to given value.

### HasValues

`func (o *MoiraNotificationEvent) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


