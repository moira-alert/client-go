# MoiraTriggerCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AloneMetrics** | Pointer to **map[string]bool** |  | [optional] 
**ClusterId** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **NullableInt64** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**Desc** | Pointer to **NullableString** |  | [optional] 
**ErrorValue** | Pointer to **NullableFloat32** |  | [optional] 
**Expression** | Pointer to **NullableString** |  | [optional] 
**Highlights** | Pointer to **map[string]string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LastCheck** | Pointer to [**MoiraCheckData**](MoiraCheckData.md) |  | [optional] 
**MuteNewMetrics** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Patterns** | Pointer to **[]string** |  | [optional] 
**PythonExpression** | Pointer to **NullableString** |  | [optional] 
**Sched** | Pointer to [**NullableMoiraScheduleData**](MoiraScheduleData.md) |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Targets** | Pointer to **[]string** |  | [optional] 
**Throttling** | Pointer to **int64** |  | [optional] 
**TriggerSource** | Pointer to **string** |  | [optional] 
**TriggerType** | Pointer to **string** |  | [optional] 
**Ttl** | Pointer to **int64** |  | [optional] 
**TtlState** | Pointer to **NullableString** |  | [optional] 
**UpdatedAt** | Pointer to **NullableInt64** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 
**WarnValue** | Pointer to **NullableFloat32** |  | [optional] 

## Methods

### NewMoiraTriggerCheck

`func NewMoiraTriggerCheck() *MoiraTriggerCheck`

NewMoiraTriggerCheck instantiates a new MoiraTriggerCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoiraTriggerCheckWithDefaults

`func NewMoiraTriggerCheckWithDefaults() *MoiraTriggerCheck`

NewMoiraTriggerCheckWithDefaults instantiates a new MoiraTriggerCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAloneMetrics

`func (o *MoiraTriggerCheck) GetAloneMetrics() map[string]bool`

GetAloneMetrics returns the AloneMetrics field if non-nil, zero value otherwise.

### GetAloneMetricsOk

`func (o *MoiraTriggerCheck) GetAloneMetricsOk() (*map[string]bool, bool)`

GetAloneMetricsOk returns a tuple with the AloneMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAloneMetrics

`func (o *MoiraTriggerCheck) SetAloneMetrics(v map[string]bool)`

SetAloneMetrics sets AloneMetrics field to given value.

### HasAloneMetrics

`func (o *MoiraTriggerCheck) HasAloneMetrics() bool`

HasAloneMetrics returns a boolean if a field has been set.

### GetClusterId

`func (o *MoiraTriggerCheck) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *MoiraTriggerCheck) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *MoiraTriggerCheck) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *MoiraTriggerCheck) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MoiraTriggerCheck) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MoiraTriggerCheck) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MoiraTriggerCheck) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MoiraTriggerCheck) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *MoiraTriggerCheck) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *MoiraTriggerCheck) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetCreatedBy

`func (o *MoiraTriggerCheck) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *MoiraTriggerCheck) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *MoiraTriggerCheck) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *MoiraTriggerCheck) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDesc

`func (o *MoiraTriggerCheck) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *MoiraTriggerCheck) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *MoiraTriggerCheck) SetDesc(v string)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *MoiraTriggerCheck) HasDesc() bool`

HasDesc returns a boolean if a field has been set.

### SetDescNil

`func (o *MoiraTriggerCheck) SetDescNil(b bool)`

 SetDescNil sets the value for Desc to be an explicit nil

### UnsetDesc
`func (o *MoiraTriggerCheck) UnsetDesc()`

UnsetDesc ensures that no value is present for Desc, not even an explicit nil
### GetErrorValue

`func (o *MoiraTriggerCheck) GetErrorValue() float32`

GetErrorValue returns the ErrorValue field if non-nil, zero value otherwise.

### GetErrorValueOk

`func (o *MoiraTriggerCheck) GetErrorValueOk() (*float32, bool)`

GetErrorValueOk returns a tuple with the ErrorValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorValue

`func (o *MoiraTriggerCheck) SetErrorValue(v float32)`

SetErrorValue sets ErrorValue field to given value.

### HasErrorValue

`func (o *MoiraTriggerCheck) HasErrorValue() bool`

HasErrorValue returns a boolean if a field has been set.

### SetErrorValueNil

`func (o *MoiraTriggerCheck) SetErrorValueNil(b bool)`

 SetErrorValueNil sets the value for ErrorValue to be an explicit nil

### UnsetErrorValue
`func (o *MoiraTriggerCheck) UnsetErrorValue()`

UnsetErrorValue ensures that no value is present for ErrorValue, not even an explicit nil
### GetExpression

`func (o *MoiraTriggerCheck) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *MoiraTriggerCheck) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *MoiraTriggerCheck) SetExpression(v string)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *MoiraTriggerCheck) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### SetExpressionNil

`func (o *MoiraTriggerCheck) SetExpressionNil(b bool)`

 SetExpressionNil sets the value for Expression to be an explicit nil

### UnsetExpression
`func (o *MoiraTriggerCheck) UnsetExpression()`

UnsetExpression ensures that no value is present for Expression, not even an explicit nil
### GetHighlights

`func (o *MoiraTriggerCheck) GetHighlights() map[string]string`

GetHighlights returns the Highlights field if non-nil, zero value otherwise.

### GetHighlightsOk

`func (o *MoiraTriggerCheck) GetHighlightsOk() (*map[string]string, bool)`

GetHighlightsOk returns a tuple with the Highlights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlights

`func (o *MoiraTriggerCheck) SetHighlights(v map[string]string)`

SetHighlights sets Highlights field to given value.

### HasHighlights

`func (o *MoiraTriggerCheck) HasHighlights() bool`

HasHighlights returns a boolean if a field has been set.

### GetId

`func (o *MoiraTriggerCheck) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MoiraTriggerCheck) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MoiraTriggerCheck) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MoiraTriggerCheck) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastCheck

`func (o *MoiraTriggerCheck) GetLastCheck() MoiraCheckData`

GetLastCheck returns the LastCheck field if non-nil, zero value otherwise.

### GetLastCheckOk

`func (o *MoiraTriggerCheck) GetLastCheckOk() (*MoiraCheckData, bool)`

GetLastCheckOk returns a tuple with the LastCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheck

`func (o *MoiraTriggerCheck) SetLastCheck(v MoiraCheckData)`

SetLastCheck sets LastCheck field to given value.

### HasLastCheck

`func (o *MoiraTriggerCheck) HasLastCheck() bool`

HasLastCheck returns a boolean if a field has been set.

### GetMuteNewMetrics

`func (o *MoiraTriggerCheck) GetMuteNewMetrics() bool`

GetMuteNewMetrics returns the MuteNewMetrics field if non-nil, zero value otherwise.

### GetMuteNewMetricsOk

`func (o *MoiraTriggerCheck) GetMuteNewMetricsOk() (*bool, bool)`

GetMuteNewMetricsOk returns a tuple with the MuteNewMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuteNewMetrics

`func (o *MoiraTriggerCheck) SetMuteNewMetrics(v bool)`

SetMuteNewMetrics sets MuteNewMetrics field to given value.

### HasMuteNewMetrics

`func (o *MoiraTriggerCheck) HasMuteNewMetrics() bool`

HasMuteNewMetrics returns a boolean if a field has been set.

### GetName

`func (o *MoiraTriggerCheck) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MoiraTriggerCheck) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MoiraTriggerCheck) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MoiraTriggerCheck) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPatterns

`func (o *MoiraTriggerCheck) GetPatterns() []string`

GetPatterns returns the Patterns field if non-nil, zero value otherwise.

### GetPatternsOk

`func (o *MoiraTriggerCheck) GetPatternsOk() (*[]string, bool)`

GetPatternsOk returns a tuple with the Patterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatterns

`func (o *MoiraTriggerCheck) SetPatterns(v []string)`

SetPatterns sets Patterns field to given value.

### HasPatterns

`func (o *MoiraTriggerCheck) HasPatterns() bool`

HasPatterns returns a boolean if a field has been set.

### GetPythonExpression

`func (o *MoiraTriggerCheck) GetPythonExpression() string`

GetPythonExpression returns the PythonExpression field if non-nil, zero value otherwise.

### GetPythonExpressionOk

`func (o *MoiraTriggerCheck) GetPythonExpressionOk() (*string, bool)`

GetPythonExpressionOk returns a tuple with the PythonExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPythonExpression

`func (o *MoiraTriggerCheck) SetPythonExpression(v string)`

SetPythonExpression sets PythonExpression field to given value.

### HasPythonExpression

`func (o *MoiraTriggerCheck) HasPythonExpression() bool`

HasPythonExpression returns a boolean if a field has been set.

### SetPythonExpressionNil

`func (o *MoiraTriggerCheck) SetPythonExpressionNil(b bool)`

 SetPythonExpressionNil sets the value for PythonExpression to be an explicit nil

### UnsetPythonExpression
`func (o *MoiraTriggerCheck) UnsetPythonExpression()`

UnsetPythonExpression ensures that no value is present for PythonExpression, not even an explicit nil
### GetSched

`func (o *MoiraTriggerCheck) GetSched() MoiraScheduleData`

GetSched returns the Sched field if non-nil, zero value otherwise.

### GetSchedOk

`func (o *MoiraTriggerCheck) GetSchedOk() (*MoiraScheduleData, bool)`

GetSchedOk returns a tuple with the Sched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSched

`func (o *MoiraTriggerCheck) SetSched(v MoiraScheduleData)`

SetSched sets Sched field to given value.

### HasSched

`func (o *MoiraTriggerCheck) HasSched() bool`

HasSched returns a boolean if a field has been set.

### SetSchedNil

`func (o *MoiraTriggerCheck) SetSchedNil(b bool)`

 SetSchedNil sets the value for Sched to be an explicit nil

### UnsetSched
`func (o *MoiraTriggerCheck) UnsetSched()`

UnsetSched ensures that no value is present for Sched, not even an explicit nil
### GetTags

`func (o *MoiraTriggerCheck) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MoiraTriggerCheck) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MoiraTriggerCheck) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *MoiraTriggerCheck) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargets

`func (o *MoiraTriggerCheck) GetTargets() []string`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *MoiraTriggerCheck) GetTargetsOk() (*[]string, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *MoiraTriggerCheck) SetTargets(v []string)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *MoiraTriggerCheck) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### GetThrottling

`func (o *MoiraTriggerCheck) GetThrottling() int64`

GetThrottling returns the Throttling field if non-nil, zero value otherwise.

### GetThrottlingOk

`func (o *MoiraTriggerCheck) GetThrottlingOk() (*int64, bool)`

GetThrottlingOk returns a tuple with the Throttling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottling

`func (o *MoiraTriggerCheck) SetThrottling(v int64)`

SetThrottling sets Throttling field to given value.

### HasThrottling

`func (o *MoiraTriggerCheck) HasThrottling() bool`

HasThrottling returns a boolean if a field has been set.

### GetTriggerSource

`func (o *MoiraTriggerCheck) GetTriggerSource() string`

GetTriggerSource returns the TriggerSource field if non-nil, zero value otherwise.

### GetTriggerSourceOk

`func (o *MoiraTriggerCheck) GetTriggerSourceOk() (*string, bool)`

GetTriggerSourceOk returns a tuple with the TriggerSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerSource

`func (o *MoiraTriggerCheck) SetTriggerSource(v string)`

SetTriggerSource sets TriggerSource field to given value.

### HasTriggerSource

`func (o *MoiraTriggerCheck) HasTriggerSource() bool`

HasTriggerSource returns a boolean if a field has been set.

### GetTriggerType

`func (o *MoiraTriggerCheck) GetTriggerType() string`

GetTriggerType returns the TriggerType field if non-nil, zero value otherwise.

### GetTriggerTypeOk

`func (o *MoiraTriggerCheck) GetTriggerTypeOk() (*string, bool)`

GetTriggerTypeOk returns a tuple with the TriggerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerType

`func (o *MoiraTriggerCheck) SetTriggerType(v string)`

SetTriggerType sets TriggerType field to given value.

### HasTriggerType

`func (o *MoiraTriggerCheck) HasTriggerType() bool`

HasTriggerType returns a boolean if a field has been set.

### GetTtl

`func (o *MoiraTriggerCheck) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *MoiraTriggerCheck) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *MoiraTriggerCheck) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *MoiraTriggerCheck) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetTtlState

`func (o *MoiraTriggerCheck) GetTtlState() string`

GetTtlState returns the TtlState field if non-nil, zero value otherwise.

### GetTtlStateOk

`func (o *MoiraTriggerCheck) GetTtlStateOk() (*string, bool)`

GetTtlStateOk returns a tuple with the TtlState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtlState

`func (o *MoiraTriggerCheck) SetTtlState(v string)`

SetTtlState sets TtlState field to given value.

### HasTtlState

`func (o *MoiraTriggerCheck) HasTtlState() bool`

HasTtlState returns a boolean if a field has been set.

### SetTtlStateNil

`func (o *MoiraTriggerCheck) SetTtlStateNil(b bool)`

 SetTtlStateNil sets the value for TtlState to be an explicit nil

### UnsetTtlState
`func (o *MoiraTriggerCheck) UnsetTtlState()`

UnsetTtlState ensures that no value is present for TtlState, not even an explicit nil
### GetUpdatedAt

`func (o *MoiraTriggerCheck) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MoiraTriggerCheck) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MoiraTriggerCheck) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *MoiraTriggerCheck) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *MoiraTriggerCheck) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *MoiraTriggerCheck) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetUpdatedBy

`func (o *MoiraTriggerCheck) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *MoiraTriggerCheck) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *MoiraTriggerCheck) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *MoiraTriggerCheck) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetWarnValue

`func (o *MoiraTriggerCheck) GetWarnValue() float32`

GetWarnValue returns the WarnValue field if non-nil, zero value otherwise.

### GetWarnValueOk

`func (o *MoiraTriggerCheck) GetWarnValueOk() (*float32, bool)`

GetWarnValueOk returns a tuple with the WarnValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnValue

`func (o *MoiraTriggerCheck) SetWarnValue(v float32)`

SetWarnValue sets WarnValue field to given value.

### HasWarnValue

`func (o *MoiraTriggerCheck) HasWarnValue() bool`

HasWarnValue returns a boolean if a field has been set.

### SetWarnValueNil

`func (o *MoiraTriggerCheck) SetWarnValueNil(b bool)`

 SetWarnValueNil sets the value for WarnValue to be an explicit nil

### UnsetWarnValue
`func (o *MoiraTriggerCheck) UnsetWarnValue()`

UnsetWarnValue ensures that no value is present for WarnValue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


