# DtoTrigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AloneMetrics** | Pointer to **map[string]bool** | A list of targets that have only alone metrics | [optional] 
**ClusterId** | Pointer to **string** | Shows the exact cluster from where the metrics are fetched | [optional] 
**CreatedAt** | Pointer to **NullableString** | Datetime when the trigger was created | [optional] 
**CreatedBy** | Pointer to **string** | Username who created trigger | [optional] 
**Desc** | Pointer to **NullableString** | Description string | [optional] 
**ErrorValue** | Pointer to **NullableFloat32** | ERROR threshold | [optional] 
**Expression** | Pointer to **string** | Used if you need more complex logic than provided by WARN/ERROR values | [optional] 
**Id** | Pointer to **string** | Trigger unique ID | [optional] 
**IsRemote** | Pointer to **bool** | Shows if trigger is remote (graphite-backend) based or stored inside Moira-Redis DB  Deprecated: Use TriggerSource field instead | [optional] 
**MuteNewMetrics** | Pointer to **bool** | If true, first event NODATA → OK will be omitted | [optional] 
**Name** | Pointer to **string** | Trigger name | [optional] 
**Patterns** | Pointer to **[]string** | Graphite patterns for trigger | [optional] 
**Sched** | Pointer to [**NullableMoiraScheduleData**](MoiraScheduleData.md) |  | [optional] 
**Tags** | Pointer to **[]string** | Set of tags to manipulate subscriptions | [optional] 
**Targets** | Pointer to **[]string** | Graphite-like targets: t1, t2, ... | [optional] 
**Throttling** | Pointer to **int64** |  | [optional] 
**TriggerSource** | Pointer to **string** | Shows the type of source from where the metrics are fetched | [optional] 
**TriggerType** | Pointer to **string** | Could be: rising, falling, expression | [optional] 
**Ttl** | Pointer to **int64** | When there are no metrics for trigger, Moira will switch metric to TTLState state after TTL seconds | [optional] 
**TtlState** | Pointer to **NullableString** | When there are no metrics for trigger, Moira will switch metric to TTLState state after TTL seconds | [optional] 
**UpdatedAt** | Pointer to **NullableString** | Datetime  when the trigger was updated | [optional] 
**UpdatedBy** | Pointer to **string** | Username who updated trigger | [optional] 
**WarnValue** | Pointer to **NullableFloat32** | WARN threshold | [optional] 

## Methods

### NewDtoTrigger

`func NewDtoTrigger() *DtoTrigger`

NewDtoTrigger instantiates a new DtoTrigger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoTriggerWithDefaults

`func NewDtoTriggerWithDefaults() *DtoTrigger`

NewDtoTriggerWithDefaults instantiates a new DtoTrigger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAloneMetrics

`func (o *DtoTrigger) GetAloneMetrics() map[string]bool`

GetAloneMetrics returns the AloneMetrics field if non-nil, zero value otherwise.

### GetAloneMetricsOk

`func (o *DtoTrigger) GetAloneMetricsOk() (*map[string]bool, bool)`

GetAloneMetricsOk returns a tuple with the AloneMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAloneMetrics

`func (o *DtoTrigger) SetAloneMetrics(v map[string]bool)`

SetAloneMetrics sets AloneMetrics field to given value.

### HasAloneMetrics

`func (o *DtoTrigger) HasAloneMetrics() bool`

HasAloneMetrics returns a boolean if a field has been set.

### GetClusterId

`func (o *DtoTrigger) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *DtoTrigger) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *DtoTrigger) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *DtoTrigger) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DtoTrigger) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DtoTrigger) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DtoTrigger) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DtoTrigger) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *DtoTrigger) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *DtoTrigger) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetCreatedBy

`func (o *DtoTrigger) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *DtoTrigger) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *DtoTrigger) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *DtoTrigger) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDesc

`func (o *DtoTrigger) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *DtoTrigger) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *DtoTrigger) SetDesc(v string)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *DtoTrigger) HasDesc() bool`

HasDesc returns a boolean if a field has been set.

### SetDescNil

`func (o *DtoTrigger) SetDescNil(b bool)`

 SetDescNil sets the value for Desc to be an explicit nil

### UnsetDesc
`func (o *DtoTrigger) UnsetDesc()`

UnsetDesc ensures that no value is present for Desc, not even an explicit nil
### GetErrorValue

`func (o *DtoTrigger) GetErrorValue() float32`

GetErrorValue returns the ErrorValue field if non-nil, zero value otherwise.

### GetErrorValueOk

`func (o *DtoTrigger) GetErrorValueOk() (*float32, bool)`

GetErrorValueOk returns a tuple with the ErrorValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorValue

`func (o *DtoTrigger) SetErrorValue(v float32)`

SetErrorValue sets ErrorValue field to given value.

### HasErrorValue

`func (o *DtoTrigger) HasErrorValue() bool`

HasErrorValue returns a boolean if a field has been set.

### SetErrorValueNil

`func (o *DtoTrigger) SetErrorValueNil(b bool)`

 SetErrorValueNil sets the value for ErrorValue to be an explicit nil

### UnsetErrorValue
`func (o *DtoTrigger) UnsetErrorValue()`

UnsetErrorValue ensures that no value is present for ErrorValue, not even an explicit nil
### GetExpression

`func (o *DtoTrigger) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *DtoTrigger) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *DtoTrigger) SetExpression(v string)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *DtoTrigger) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### GetId

`func (o *DtoTrigger) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DtoTrigger) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DtoTrigger) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DtoTrigger) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsRemote

`func (o *DtoTrigger) GetIsRemote() bool`

GetIsRemote returns the IsRemote field if non-nil, zero value otherwise.

### GetIsRemoteOk

`func (o *DtoTrigger) GetIsRemoteOk() (*bool, bool)`

GetIsRemoteOk returns a tuple with the IsRemote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemote

`func (o *DtoTrigger) SetIsRemote(v bool)`

SetIsRemote sets IsRemote field to given value.

### HasIsRemote

`func (o *DtoTrigger) HasIsRemote() bool`

HasIsRemote returns a boolean if a field has been set.

### GetMuteNewMetrics

`func (o *DtoTrigger) GetMuteNewMetrics() bool`

GetMuteNewMetrics returns the MuteNewMetrics field if non-nil, zero value otherwise.

### GetMuteNewMetricsOk

`func (o *DtoTrigger) GetMuteNewMetricsOk() (*bool, bool)`

GetMuteNewMetricsOk returns a tuple with the MuteNewMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuteNewMetrics

`func (o *DtoTrigger) SetMuteNewMetrics(v bool)`

SetMuteNewMetrics sets MuteNewMetrics field to given value.

### HasMuteNewMetrics

`func (o *DtoTrigger) HasMuteNewMetrics() bool`

HasMuteNewMetrics returns a boolean if a field has been set.

### GetName

`func (o *DtoTrigger) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DtoTrigger) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DtoTrigger) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DtoTrigger) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPatterns

`func (o *DtoTrigger) GetPatterns() []string`

GetPatterns returns the Patterns field if non-nil, zero value otherwise.

### GetPatternsOk

`func (o *DtoTrigger) GetPatternsOk() (*[]string, bool)`

GetPatternsOk returns a tuple with the Patterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatterns

`func (o *DtoTrigger) SetPatterns(v []string)`

SetPatterns sets Patterns field to given value.

### HasPatterns

`func (o *DtoTrigger) HasPatterns() bool`

HasPatterns returns a boolean if a field has been set.

### GetSched

`func (o *DtoTrigger) GetSched() MoiraScheduleData`

GetSched returns the Sched field if non-nil, zero value otherwise.

### GetSchedOk

`func (o *DtoTrigger) GetSchedOk() (*MoiraScheduleData, bool)`

GetSchedOk returns a tuple with the Sched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSched

`func (o *DtoTrigger) SetSched(v MoiraScheduleData)`

SetSched sets Sched field to given value.

### HasSched

`func (o *DtoTrigger) HasSched() bool`

HasSched returns a boolean if a field has been set.

### SetSchedNil

`func (o *DtoTrigger) SetSchedNil(b bool)`

 SetSchedNil sets the value for Sched to be an explicit nil

### UnsetSched
`func (o *DtoTrigger) UnsetSched()`

UnsetSched ensures that no value is present for Sched, not even an explicit nil
### GetTags

`func (o *DtoTrigger) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DtoTrigger) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DtoTrigger) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DtoTrigger) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargets

`func (o *DtoTrigger) GetTargets() []string`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *DtoTrigger) GetTargetsOk() (*[]string, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *DtoTrigger) SetTargets(v []string)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *DtoTrigger) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### GetThrottling

`func (o *DtoTrigger) GetThrottling() int64`

GetThrottling returns the Throttling field if non-nil, zero value otherwise.

### GetThrottlingOk

`func (o *DtoTrigger) GetThrottlingOk() (*int64, bool)`

GetThrottlingOk returns a tuple with the Throttling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottling

`func (o *DtoTrigger) SetThrottling(v int64)`

SetThrottling sets Throttling field to given value.

### HasThrottling

`func (o *DtoTrigger) HasThrottling() bool`

HasThrottling returns a boolean if a field has been set.

### GetTriggerSource

`func (o *DtoTrigger) GetTriggerSource() string`

GetTriggerSource returns the TriggerSource field if non-nil, zero value otherwise.

### GetTriggerSourceOk

`func (o *DtoTrigger) GetTriggerSourceOk() (*string, bool)`

GetTriggerSourceOk returns a tuple with the TriggerSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerSource

`func (o *DtoTrigger) SetTriggerSource(v string)`

SetTriggerSource sets TriggerSource field to given value.

### HasTriggerSource

`func (o *DtoTrigger) HasTriggerSource() bool`

HasTriggerSource returns a boolean if a field has been set.

### GetTriggerType

`func (o *DtoTrigger) GetTriggerType() string`

GetTriggerType returns the TriggerType field if non-nil, zero value otherwise.

### GetTriggerTypeOk

`func (o *DtoTrigger) GetTriggerTypeOk() (*string, bool)`

GetTriggerTypeOk returns a tuple with the TriggerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerType

`func (o *DtoTrigger) SetTriggerType(v string)`

SetTriggerType sets TriggerType field to given value.

### HasTriggerType

`func (o *DtoTrigger) HasTriggerType() bool`

HasTriggerType returns a boolean if a field has been set.

### GetTtl

`func (o *DtoTrigger) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *DtoTrigger) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *DtoTrigger) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *DtoTrigger) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetTtlState

`func (o *DtoTrigger) GetTtlState() string`

GetTtlState returns the TtlState field if non-nil, zero value otherwise.

### GetTtlStateOk

`func (o *DtoTrigger) GetTtlStateOk() (*string, bool)`

GetTtlStateOk returns a tuple with the TtlState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtlState

`func (o *DtoTrigger) SetTtlState(v string)`

SetTtlState sets TtlState field to given value.

### HasTtlState

`func (o *DtoTrigger) HasTtlState() bool`

HasTtlState returns a boolean if a field has been set.

### SetTtlStateNil

`func (o *DtoTrigger) SetTtlStateNil(b bool)`

 SetTtlStateNil sets the value for TtlState to be an explicit nil

### UnsetTtlState
`func (o *DtoTrigger) UnsetTtlState()`

UnsetTtlState ensures that no value is present for TtlState, not even an explicit nil
### GetUpdatedAt

`func (o *DtoTrigger) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DtoTrigger) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DtoTrigger) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DtoTrigger) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *DtoTrigger) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *DtoTrigger) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetUpdatedBy

`func (o *DtoTrigger) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *DtoTrigger) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *DtoTrigger) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *DtoTrigger) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetWarnValue

`func (o *DtoTrigger) GetWarnValue() float32`

GetWarnValue returns the WarnValue field if non-nil, zero value otherwise.

### GetWarnValueOk

`func (o *DtoTrigger) GetWarnValueOk() (*float32, bool)`

GetWarnValueOk returns a tuple with the WarnValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnValue

`func (o *DtoTrigger) SetWarnValue(v float32)`

SetWarnValue sets WarnValue field to given value.

### HasWarnValue

`func (o *DtoTrigger) HasWarnValue() bool`

HasWarnValue returns a boolean if a field has been set.

### SetWarnValueNil

`func (o *DtoTrigger) SetWarnValueNil(b bool)`

 SetWarnValueNil sets the value for WarnValue to be an explicit nil

### UnsetWarnValue
`func (o *DtoTrigger) UnsetWarnValue()`

UnsetWarnValue ensures that no value is present for WarnValue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


