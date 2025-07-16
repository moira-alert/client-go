# MoiraMetricState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeletedButKept** | Pointer to **bool** | DeletedButKept controls whether the metric is shown to the user if the trigger has ttlState &#x3D; Del and the metric is in Maintenance. The metric remains in the database | [optional] 
**EventTimestamp** | Pointer to **int64** |  | [optional] 
**Maintenance** | Pointer to **int64** |  | [optional] 
**MaintenanceInfo** | Pointer to [**NullableMoiraMaintenanceInfo**](MoiraMaintenanceInfo.md) |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Suppressed** | Pointer to **bool** |  | [optional] 
**SuppressedState** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **int64** |  | [optional] 
**Value** | Pointer to **NullableFloat32** |  | [optional] 
**Values** | Pointer to **map[string]float32** |  | [optional] 

## Methods

### NewMoiraMetricState

`func NewMoiraMetricState() *MoiraMetricState`

NewMoiraMetricState instantiates a new MoiraMetricState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoiraMetricStateWithDefaults

`func NewMoiraMetricStateWithDefaults() *MoiraMetricState`

NewMoiraMetricStateWithDefaults instantiates a new MoiraMetricState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeletedButKept

`func (o *MoiraMetricState) GetDeletedButKept() bool`

GetDeletedButKept returns the DeletedButKept field if non-nil, zero value otherwise.

### GetDeletedButKeptOk

`func (o *MoiraMetricState) GetDeletedButKeptOk() (*bool, bool)`

GetDeletedButKeptOk returns a tuple with the DeletedButKept field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedButKept

`func (o *MoiraMetricState) SetDeletedButKept(v bool)`

SetDeletedButKept sets DeletedButKept field to given value.

### HasDeletedButKept

`func (o *MoiraMetricState) HasDeletedButKept() bool`

HasDeletedButKept returns a boolean if a field has been set.

### GetEventTimestamp

`func (o *MoiraMetricState) GetEventTimestamp() int64`

GetEventTimestamp returns the EventTimestamp field if non-nil, zero value otherwise.

### GetEventTimestampOk

`func (o *MoiraMetricState) GetEventTimestampOk() (*int64, bool)`

GetEventTimestampOk returns a tuple with the EventTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTimestamp

`func (o *MoiraMetricState) SetEventTimestamp(v int64)`

SetEventTimestamp sets EventTimestamp field to given value.

### HasEventTimestamp

`func (o *MoiraMetricState) HasEventTimestamp() bool`

HasEventTimestamp returns a boolean if a field has been set.

### GetMaintenance

`func (o *MoiraMetricState) GetMaintenance() int64`

GetMaintenance returns the Maintenance field if non-nil, zero value otherwise.

### GetMaintenanceOk

`func (o *MoiraMetricState) GetMaintenanceOk() (*int64, bool)`

GetMaintenanceOk returns a tuple with the Maintenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenance

`func (o *MoiraMetricState) SetMaintenance(v int64)`

SetMaintenance sets Maintenance field to given value.

### HasMaintenance

`func (o *MoiraMetricState) HasMaintenance() bool`

HasMaintenance returns a boolean if a field has been set.

### GetMaintenanceInfo

`func (o *MoiraMetricState) GetMaintenanceInfo() MoiraMaintenanceInfo`

GetMaintenanceInfo returns the MaintenanceInfo field if non-nil, zero value otherwise.

### GetMaintenanceInfoOk

`func (o *MoiraMetricState) GetMaintenanceInfoOk() (*MoiraMaintenanceInfo, bool)`

GetMaintenanceInfoOk returns a tuple with the MaintenanceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceInfo

`func (o *MoiraMetricState) SetMaintenanceInfo(v MoiraMaintenanceInfo)`

SetMaintenanceInfo sets MaintenanceInfo field to given value.

### HasMaintenanceInfo

`func (o *MoiraMetricState) HasMaintenanceInfo() bool`

HasMaintenanceInfo returns a boolean if a field has been set.

### SetMaintenanceInfoNil

`func (o *MoiraMetricState) SetMaintenanceInfoNil(b bool)`

 SetMaintenanceInfoNil sets the value for MaintenanceInfo to be an explicit nil

### UnsetMaintenanceInfo
`func (o *MoiraMetricState) UnsetMaintenanceInfo()`

UnsetMaintenanceInfo ensures that no value is present for MaintenanceInfo, not even an explicit nil
### GetState

`func (o *MoiraMetricState) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MoiraMetricState) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MoiraMetricState) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *MoiraMetricState) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSuppressed

`func (o *MoiraMetricState) GetSuppressed() bool`

GetSuppressed returns the Suppressed field if non-nil, zero value otherwise.

### GetSuppressedOk

`func (o *MoiraMetricState) GetSuppressedOk() (*bool, bool)`

GetSuppressedOk returns a tuple with the Suppressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressed

`func (o *MoiraMetricState) SetSuppressed(v bool)`

SetSuppressed sets Suppressed field to given value.

### HasSuppressed

`func (o *MoiraMetricState) HasSuppressed() bool`

HasSuppressed returns a boolean if a field has been set.

### GetSuppressedState

`func (o *MoiraMetricState) GetSuppressedState() string`

GetSuppressedState returns the SuppressedState field if non-nil, zero value otherwise.

### GetSuppressedStateOk

`func (o *MoiraMetricState) GetSuppressedStateOk() (*string, bool)`

GetSuppressedStateOk returns a tuple with the SuppressedState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressedState

`func (o *MoiraMetricState) SetSuppressedState(v string)`

SetSuppressedState sets SuppressedState field to given value.

### HasSuppressedState

`func (o *MoiraMetricState) HasSuppressedState() bool`

HasSuppressedState returns a boolean if a field has been set.

### GetTimestamp

`func (o *MoiraMetricState) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *MoiraMetricState) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *MoiraMetricState) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *MoiraMetricState) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetValue

`func (o *MoiraMetricState) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MoiraMetricState) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MoiraMetricState) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *MoiraMetricState) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *MoiraMetricState) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *MoiraMetricState) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetValues

`func (o *MoiraMetricState) GetValues() map[string]float32`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *MoiraMetricState) GetValuesOk() (*map[string]float32, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *MoiraMetricState) SetValues(v map[string]float32)`

SetValues sets Values field to given value.

### HasValues

`func (o *MoiraMetricState) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


