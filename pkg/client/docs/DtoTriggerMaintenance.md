# DtoTriggerMaintenance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metrics** | Pointer to **map[string]interface{}** |  | [optional] 
**Trigger** | Pointer to **NullableInt64** |  | [optional] 

## Methods

### NewDtoTriggerMaintenance

`func NewDtoTriggerMaintenance() *DtoTriggerMaintenance`

NewDtoTriggerMaintenance instantiates a new DtoTriggerMaintenance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoTriggerMaintenanceWithDefaults

`func NewDtoTriggerMaintenanceWithDefaults() *DtoTriggerMaintenance`

NewDtoTriggerMaintenanceWithDefaults instantiates a new DtoTriggerMaintenance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetrics

`func (o *DtoTriggerMaintenance) GetMetrics() map[string]interface{}`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *DtoTriggerMaintenance) GetMetricsOk() (*map[string]interface{}, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *DtoTriggerMaintenance) SetMetrics(v map[string]interface{})`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *DtoTriggerMaintenance) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetTrigger

`func (o *DtoTriggerMaintenance) GetTrigger() int64`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *DtoTriggerMaintenance) GetTriggerOk() (*int64, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *DtoTriggerMaintenance) SetTrigger(v int64)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *DtoTriggerMaintenance) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.

### SetTriggerNil

`func (o *DtoTriggerMaintenance) SetTriggerNil(b bool)`

 SetTriggerNil sets the value for Trigger to be an explicit nil

### UnsetTrigger
`func (o *DtoTriggerMaintenance) UnsetTrigger()`

UnsetTrigger ensures that no value is present for Trigger, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


