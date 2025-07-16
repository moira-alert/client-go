# DtoPatternMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metrics** | Pointer to [**map[string][]MoiraMetricValue**](array.md) |  | [optional] 
**Pattern** | Pointer to **string** |  | [optional] 
**Retention** | Pointer to **map[string]int32** |  | [optional] 

## Methods

### NewDtoPatternMetrics

`func NewDtoPatternMetrics() *DtoPatternMetrics`

NewDtoPatternMetrics instantiates a new DtoPatternMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoPatternMetricsWithDefaults

`func NewDtoPatternMetricsWithDefaults() *DtoPatternMetrics`

NewDtoPatternMetricsWithDefaults instantiates a new DtoPatternMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetrics

`func (o *DtoPatternMetrics) GetMetrics() map[string][]MoiraMetricValue`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *DtoPatternMetrics) GetMetricsOk() (*map[string][]MoiraMetricValue, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *DtoPatternMetrics) SetMetrics(v map[string][]MoiraMetricValue)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *DtoPatternMetrics) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetPattern

`func (o *DtoPatternMetrics) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *DtoPatternMetrics) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *DtoPatternMetrics) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *DtoPatternMetrics) HasPattern() bool`

HasPattern returns a boolean if a field has been set.

### GetRetention

`func (o *DtoPatternMetrics) GetRetention() map[string]int32`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *DtoPatternMetrics) GetRetentionOk() (*map[string]int32, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *DtoPatternMetrics) SetRetention(v map[string]int32)`

SetRetention sets Retention field to given value.

### HasRetention

`func (o *DtoPatternMetrics) HasRetention() bool`

HasRetention returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


