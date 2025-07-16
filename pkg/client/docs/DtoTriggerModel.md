# DtoTriggerModel

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
**TriggerSource** | Pointer to **string** | Shows the type of source from where the metrics are fetched | [optional] 
**TriggerType** | Pointer to **string** | Could be: rising, falling, expression | [optional] 
**Ttl** | Pointer to **int64** | When there are no metrics for trigger, Moira will switch metric to TTLState state after TTL seconds | [optional] 
**TtlState** | Pointer to **NullableString** | When there are no metrics for trigger, Moira will switch metric to TTLState state after TTL seconds | [optional] 
**UpdatedAt** | Pointer to **NullableString** | Datetime  when the trigger was updated | [optional] 
**UpdatedBy** | Pointer to **string** | Username who updated trigger | [optional] 
**WarnValue** | Pointer to **NullableFloat32** | WARN threshold | [optional] 

## Methods

### NewDtoTriggerModel

`func NewDtoTriggerModel() *DtoTriggerModel`

NewDtoTriggerModel instantiates a new DtoTriggerModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoTriggerModelWithDefaults

`func NewDtoTriggerModelWithDefaults() *DtoTriggerModel`

NewDtoTriggerModelWithDefaults instantiates a new DtoTriggerModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAloneMetrics

`func (o *DtoTriggerModel) GetAloneMetrics() map[string]bool`

GetAloneMetrics returns the AloneMetrics field if non-nil, zero value otherwise.

### GetAloneMetricsOk

`func (o *DtoTriggerModel) GetAloneMetricsOk() (*map[string]bool, bool)`

GetAloneMetricsOk returns a tuple with the AloneMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAloneMetrics

`func (o *DtoTriggerModel) SetAloneMetrics(v map[string]bool)`

SetAloneMetrics sets AloneMetrics field to given value.

### HasAloneMetrics

`func (o *DtoTriggerModel) HasAloneMetrics() bool`

HasAloneMetrics returns a boolean if a field has been set.

### GetClusterId

`func (o *DtoTriggerModel) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *DtoTriggerModel) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *DtoTriggerModel) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *DtoTriggerModel) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DtoTriggerModel) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DtoTriggerModel) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DtoTriggerModel) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DtoTriggerModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *DtoTriggerModel) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *DtoTriggerModel) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetCreatedBy

`func (o *DtoTriggerModel) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *DtoTriggerModel) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *DtoTriggerModel) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *DtoTriggerModel) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDesc

`func (o *DtoTriggerModel) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *DtoTriggerModel) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *DtoTriggerModel) SetDesc(v string)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *DtoTriggerModel) HasDesc() bool`

HasDesc returns a boolean if a field has been set.

### SetDescNil

`func (o *DtoTriggerModel) SetDescNil(b bool)`

 SetDescNil sets the value for Desc to be an explicit nil

### UnsetDesc
`func (o *DtoTriggerModel) UnsetDesc()`

UnsetDesc ensures that no value is present for Desc, not even an explicit nil
### GetErrorValue

`func (o *DtoTriggerModel) GetErrorValue() float32`

GetErrorValue returns the ErrorValue field if non-nil, zero value otherwise.

### GetErrorValueOk

`func (o *DtoTriggerModel) GetErrorValueOk() (*float32, bool)`

GetErrorValueOk returns a tuple with the ErrorValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorValue

`func (o *DtoTriggerModel) SetErrorValue(v float32)`

SetErrorValue sets ErrorValue field to given value.

### HasErrorValue

`func (o *DtoTriggerModel) HasErrorValue() bool`

HasErrorValue returns a boolean if a field has been set.

### SetErrorValueNil

`func (o *DtoTriggerModel) SetErrorValueNil(b bool)`

 SetErrorValueNil sets the value for ErrorValue to be an explicit nil

### UnsetErrorValue
`func (o *DtoTriggerModel) UnsetErrorValue()`

UnsetErrorValue ensures that no value is present for ErrorValue, not even an explicit nil
### GetExpression

`func (o *DtoTriggerModel) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *DtoTriggerModel) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *DtoTriggerModel) SetExpression(v string)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *DtoTriggerModel) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### GetId

`func (o *DtoTriggerModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DtoTriggerModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DtoTriggerModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DtoTriggerModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsRemote

`func (o *DtoTriggerModel) GetIsRemote() bool`

GetIsRemote returns the IsRemote field if non-nil, zero value otherwise.

### GetIsRemoteOk

`func (o *DtoTriggerModel) GetIsRemoteOk() (*bool, bool)`

GetIsRemoteOk returns a tuple with the IsRemote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemote

`func (o *DtoTriggerModel) SetIsRemote(v bool)`

SetIsRemote sets IsRemote field to given value.

### HasIsRemote

`func (o *DtoTriggerModel) HasIsRemote() bool`

HasIsRemote returns a boolean if a field has been set.

### GetMuteNewMetrics

`func (o *DtoTriggerModel) GetMuteNewMetrics() bool`

GetMuteNewMetrics returns the MuteNewMetrics field if non-nil, zero value otherwise.

### GetMuteNewMetricsOk

`func (o *DtoTriggerModel) GetMuteNewMetricsOk() (*bool, bool)`

GetMuteNewMetricsOk returns a tuple with the MuteNewMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuteNewMetrics

`func (o *DtoTriggerModel) SetMuteNewMetrics(v bool)`

SetMuteNewMetrics sets MuteNewMetrics field to given value.

### HasMuteNewMetrics

`func (o *DtoTriggerModel) HasMuteNewMetrics() bool`

HasMuteNewMetrics returns a boolean if a field has been set.

### GetName

`func (o *DtoTriggerModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DtoTriggerModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DtoTriggerModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DtoTriggerModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPatterns

`func (o *DtoTriggerModel) GetPatterns() []string`

GetPatterns returns the Patterns field if non-nil, zero value otherwise.

### GetPatternsOk

`func (o *DtoTriggerModel) GetPatternsOk() (*[]string, bool)`

GetPatternsOk returns a tuple with the Patterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatterns

`func (o *DtoTriggerModel) SetPatterns(v []string)`

SetPatterns sets Patterns field to given value.

### HasPatterns

`func (o *DtoTriggerModel) HasPatterns() bool`

HasPatterns returns a boolean if a field has been set.

### GetSched

`func (o *DtoTriggerModel) GetSched() MoiraScheduleData`

GetSched returns the Sched field if non-nil, zero value otherwise.

### GetSchedOk

`func (o *DtoTriggerModel) GetSchedOk() (*MoiraScheduleData, bool)`

GetSchedOk returns a tuple with the Sched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSched

`func (o *DtoTriggerModel) SetSched(v MoiraScheduleData)`

SetSched sets Sched field to given value.

### HasSched

`func (o *DtoTriggerModel) HasSched() bool`

HasSched returns a boolean if a field has been set.

### SetSchedNil

`func (o *DtoTriggerModel) SetSchedNil(b bool)`

 SetSchedNil sets the value for Sched to be an explicit nil

### UnsetSched
`func (o *DtoTriggerModel) UnsetSched()`

UnsetSched ensures that no value is present for Sched, not even an explicit nil
### GetTags

`func (o *DtoTriggerModel) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DtoTriggerModel) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DtoTriggerModel) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DtoTriggerModel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargets

`func (o *DtoTriggerModel) GetTargets() []string`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *DtoTriggerModel) GetTargetsOk() (*[]string, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *DtoTriggerModel) SetTargets(v []string)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *DtoTriggerModel) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### GetTriggerSource

`func (o *DtoTriggerModel) GetTriggerSource() string`

GetTriggerSource returns the TriggerSource field if non-nil, zero value otherwise.

### GetTriggerSourceOk

`func (o *DtoTriggerModel) GetTriggerSourceOk() (*string, bool)`

GetTriggerSourceOk returns a tuple with the TriggerSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerSource

`func (o *DtoTriggerModel) SetTriggerSource(v string)`

SetTriggerSource sets TriggerSource field to given value.

### HasTriggerSource

`func (o *DtoTriggerModel) HasTriggerSource() bool`

HasTriggerSource returns a boolean if a field has been set.

### GetTriggerType

`func (o *DtoTriggerModel) GetTriggerType() string`

GetTriggerType returns the TriggerType field if non-nil, zero value otherwise.

### GetTriggerTypeOk

`func (o *DtoTriggerModel) GetTriggerTypeOk() (*string, bool)`

GetTriggerTypeOk returns a tuple with the TriggerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerType

`func (o *DtoTriggerModel) SetTriggerType(v string)`

SetTriggerType sets TriggerType field to given value.

### HasTriggerType

`func (o *DtoTriggerModel) HasTriggerType() bool`

HasTriggerType returns a boolean if a field has been set.

### GetTtl

`func (o *DtoTriggerModel) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *DtoTriggerModel) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *DtoTriggerModel) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *DtoTriggerModel) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetTtlState

`func (o *DtoTriggerModel) GetTtlState() string`

GetTtlState returns the TtlState field if non-nil, zero value otherwise.

### GetTtlStateOk

`func (o *DtoTriggerModel) GetTtlStateOk() (*string, bool)`

GetTtlStateOk returns a tuple with the TtlState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtlState

`func (o *DtoTriggerModel) SetTtlState(v string)`

SetTtlState sets TtlState field to given value.

### HasTtlState

`func (o *DtoTriggerModel) HasTtlState() bool`

HasTtlState returns a boolean if a field has been set.

### SetTtlStateNil

`func (o *DtoTriggerModel) SetTtlStateNil(b bool)`

 SetTtlStateNil sets the value for TtlState to be an explicit nil

### UnsetTtlState
`func (o *DtoTriggerModel) UnsetTtlState()`

UnsetTtlState ensures that no value is present for TtlState, not even an explicit nil
### GetUpdatedAt

`func (o *DtoTriggerModel) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DtoTriggerModel) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DtoTriggerModel) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DtoTriggerModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *DtoTriggerModel) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *DtoTriggerModel) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetUpdatedBy

`func (o *DtoTriggerModel) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *DtoTriggerModel) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *DtoTriggerModel) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *DtoTriggerModel) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetWarnValue

`func (o *DtoTriggerModel) GetWarnValue() float32`

GetWarnValue returns the WarnValue field if non-nil, zero value otherwise.

### GetWarnValueOk

`func (o *DtoTriggerModel) GetWarnValueOk() (*float32, bool)`

GetWarnValueOk returns a tuple with the WarnValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnValue

`func (o *DtoTriggerModel) SetWarnValue(v float32)`

SetWarnValue sets WarnValue field to given value.

### HasWarnValue

`func (o *DtoTriggerModel) HasWarnValue() bool`

HasWarnValue returns a boolean if a field has been set.

### SetWarnValueNil

`func (o *DtoTriggerModel) SetWarnValueNil(b bool)`

 SetWarnValueNil sets the value for WarnValue to be an explicit nil

### UnsetWarnValue
`func (o *DtoTriggerModel) UnsetWarnValue()`

UnsetWarnValue ensures that no value is present for WarnValue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


