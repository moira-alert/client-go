# DtoTriggerDump

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **string** |  | [optional] 
**LastCheck** | Pointer to [**MoiraCheckData**](MoiraCheckData.md) |  | [optional] 
**Metrics** | Pointer to [**[]DtoPatternMetrics**](DtoPatternMetrics.md) |  | [optional] 
**Trigger** | Pointer to [**MoiraTrigger**](MoiraTrigger.md) |  | [optional] 

## Methods

### NewDtoTriggerDump

`func NewDtoTriggerDump() *DtoTriggerDump`

NewDtoTriggerDump instantiates a new DtoTriggerDump object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoTriggerDumpWithDefaults

`func NewDtoTriggerDumpWithDefaults() *DtoTriggerDump`

NewDtoTriggerDumpWithDefaults instantiates a new DtoTriggerDump object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *DtoTriggerDump) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *DtoTriggerDump) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *DtoTriggerDump) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *DtoTriggerDump) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLastCheck

`func (o *DtoTriggerDump) GetLastCheck() MoiraCheckData`

GetLastCheck returns the LastCheck field if non-nil, zero value otherwise.

### GetLastCheckOk

`func (o *DtoTriggerDump) GetLastCheckOk() (*MoiraCheckData, bool)`

GetLastCheckOk returns a tuple with the LastCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheck

`func (o *DtoTriggerDump) SetLastCheck(v MoiraCheckData)`

SetLastCheck sets LastCheck field to given value.

### HasLastCheck

`func (o *DtoTriggerDump) HasLastCheck() bool`

HasLastCheck returns a boolean if a field has been set.

### GetMetrics

`func (o *DtoTriggerDump) GetMetrics() []DtoPatternMetrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *DtoTriggerDump) GetMetricsOk() (*[]DtoPatternMetrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *DtoTriggerDump) SetMetrics(v []DtoPatternMetrics)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *DtoTriggerDump) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetTrigger

`func (o *DtoTriggerDump) GetTrigger() MoiraTrigger`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *DtoTriggerDump) GetTriggerOk() (*MoiraTrigger, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *DtoTriggerDump) SetTrigger(v MoiraTrigger)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *DtoTriggerDump) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


