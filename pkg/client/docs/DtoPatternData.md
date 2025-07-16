# DtoPatternData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metrics** | Pointer to **[]string** |  | [optional] 
**Pattern** | Pointer to **string** |  | [optional] 
**Triggers** | Pointer to [**[]DtoTriggerModel**](DtoTriggerModel.md) |  | [optional] 

## Methods

### NewDtoPatternData

`func NewDtoPatternData() *DtoPatternData`

NewDtoPatternData instantiates a new DtoPatternData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoPatternDataWithDefaults

`func NewDtoPatternDataWithDefaults() *DtoPatternData`

NewDtoPatternDataWithDefaults instantiates a new DtoPatternData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetrics

`func (o *DtoPatternData) GetMetrics() []string`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *DtoPatternData) GetMetricsOk() (*[]string, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *DtoPatternData) SetMetrics(v []string)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *DtoPatternData) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetPattern

`func (o *DtoPatternData) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *DtoPatternData) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *DtoPatternData) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *DtoPatternData) HasPattern() bool`

HasPattern returns a boolean if a field has been set.

### GetTriggers

`func (o *DtoPatternData) GetTriggers() []DtoTriggerModel`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *DtoPatternData) GetTriggersOk() (*[]DtoTriggerModel, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *DtoPatternData) SetTriggers(v []DtoTriggerModel)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *DtoPatternData) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


