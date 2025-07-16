# MoiraEventInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Interval** | Pointer to **NullableInt64** |  | [optional] 
**Maintenance** | Pointer to [**NullableMoiraMaintenanceInfo**](MoiraMaintenanceInfo.md) |  | [optional] 

## Methods

### NewMoiraEventInfo

`func NewMoiraEventInfo() *MoiraEventInfo`

NewMoiraEventInfo instantiates a new MoiraEventInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoiraEventInfoWithDefaults

`func NewMoiraEventInfoWithDefaults() *MoiraEventInfo`

NewMoiraEventInfoWithDefaults instantiates a new MoiraEventInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterval

`func (o *MoiraEventInfo) GetInterval() int64`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *MoiraEventInfo) GetIntervalOk() (*int64, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *MoiraEventInfo) SetInterval(v int64)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *MoiraEventInfo) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### SetIntervalNil

`func (o *MoiraEventInfo) SetIntervalNil(b bool)`

 SetIntervalNil sets the value for Interval to be an explicit nil

### UnsetInterval
`func (o *MoiraEventInfo) UnsetInterval()`

UnsetInterval ensures that no value is present for Interval, not even an explicit nil
### GetMaintenance

`func (o *MoiraEventInfo) GetMaintenance() MoiraMaintenanceInfo`

GetMaintenance returns the Maintenance field if non-nil, zero value otherwise.

### GetMaintenanceOk

`func (o *MoiraEventInfo) GetMaintenanceOk() (*MoiraMaintenanceInfo, bool)`

GetMaintenanceOk returns a tuple with the Maintenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenance

`func (o *MoiraEventInfo) SetMaintenance(v MoiraMaintenanceInfo)`

SetMaintenance sets Maintenance field to given value.

### HasMaintenance

`func (o *MoiraEventInfo) HasMaintenance() bool`

HasMaintenance returns a boolean if a field has been set.

### SetMaintenanceNil

`func (o *MoiraEventInfo) SetMaintenanceNil(b bool)`

 SetMaintenanceNil sets the value for Maintenance to be an explicit nil

### UnsetMaintenance
`func (o *MoiraEventInfo) UnsetMaintenance()`

UnsetMaintenance ensures that no value is present for Maintenance, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


