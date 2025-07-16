# MoiraCheckData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventTimestamp** | Pointer to **int64** |  | [optional] 
**LastSuccessfulCheckTimestamp** | Pointer to **int64** | LastSuccessfulCheckTimestamp - time of the last check of the trigger, during which there were no errors | [optional] 
**Maintenance** | Pointer to **int64** |  | [optional] 
**MaintenanceInfo** | Pointer to [**NullableMoiraMaintenanceInfo**](MoiraMaintenanceInfo.md) |  | [optional] 
**Metrics** | Pointer to [**map[string]MoiraMetricState**](MoiraMetricState.md) |  | [optional] 
**MetricsToTargetRelation** | Pointer to **map[string]string** | MetricsToTargetRelation is a map that holds relation between metric names that was alone during last check and targets that fetched this metric  {\&quot;t1\&quot;: \&quot;metric.name.1\&quot;, \&quot;t2\&quot;: \&quot;metric.name.2\&quot;} | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**Score** | Pointer to **int64** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Suppressed** | Pointer to **bool** |  | [optional] 
**SuppressedState** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **int64** | Timestamp - time, which means when the checker last checked this trigger, this value stops updating if the trigger does not receive metrics | [optional] 

## Methods

### NewMoiraCheckData

`func NewMoiraCheckData() *MoiraCheckData`

NewMoiraCheckData instantiates a new MoiraCheckData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoiraCheckDataWithDefaults

`func NewMoiraCheckDataWithDefaults() *MoiraCheckData`

NewMoiraCheckDataWithDefaults instantiates a new MoiraCheckData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventTimestamp

`func (o *MoiraCheckData) GetEventTimestamp() int64`

GetEventTimestamp returns the EventTimestamp field if non-nil, zero value otherwise.

### GetEventTimestampOk

`func (o *MoiraCheckData) GetEventTimestampOk() (*int64, bool)`

GetEventTimestampOk returns a tuple with the EventTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTimestamp

`func (o *MoiraCheckData) SetEventTimestamp(v int64)`

SetEventTimestamp sets EventTimestamp field to given value.

### HasEventTimestamp

`func (o *MoiraCheckData) HasEventTimestamp() bool`

HasEventTimestamp returns a boolean if a field has been set.

### GetLastSuccessfulCheckTimestamp

`func (o *MoiraCheckData) GetLastSuccessfulCheckTimestamp() int64`

GetLastSuccessfulCheckTimestamp returns the LastSuccessfulCheckTimestamp field if non-nil, zero value otherwise.

### GetLastSuccessfulCheckTimestampOk

`func (o *MoiraCheckData) GetLastSuccessfulCheckTimestampOk() (*int64, bool)`

GetLastSuccessfulCheckTimestampOk returns a tuple with the LastSuccessfulCheckTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccessfulCheckTimestamp

`func (o *MoiraCheckData) SetLastSuccessfulCheckTimestamp(v int64)`

SetLastSuccessfulCheckTimestamp sets LastSuccessfulCheckTimestamp field to given value.

### HasLastSuccessfulCheckTimestamp

`func (o *MoiraCheckData) HasLastSuccessfulCheckTimestamp() bool`

HasLastSuccessfulCheckTimestamp returns a boolean if a field has been set.

### GetMaintenance

`func (o *MoiraCheckData) GetMaintenance() int64`

GetMaintenance returns the Maintenance field if non-nil, zero value otherwise.

### GetMaintenanceOk

`func (o *MoiraCheckData) GetMaintenanceOk() (*int64, bool)`

GetMaintenanceOk returns a tuple with the Maintenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenance

`func (o *MoiraCheckData) SetMaintenance(v int64)`

SetMaintenance sets Maintenance field to given value.

### HasMaintenance

`func (o *MoiraCheckData) HasMaintenance() bool`

HasMaintenance returns a boolean if a field has been set.

### GetMaintenanceInfo

`func (o *MoiraCheckData) GetMaintenanceInfo() MoiraMaintenanceInfo`

GetMaintenanceInfo returns the MaintenanceInfo field if non-nil, zero value otherwise.

### GetMaintenanceInfoOk

`func (o *MoiraCheckData) GetMaintenanceInfoOk() (*MoiraMaintenanceInfo, bool)`

GetMaintenanceInfoOk returns a tuple with the MaintenanceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceInfo

`func (o *MoiraCheckData) SetMaintenanceInfo(v MoiraMaintenanceInfo)`

SetMaintenanceInfo sets MaintenanceInfo field to given value.

### HasMaintenanceInfo

`func (o *MoiraCheckData) HasMaintenanceInfo() bool`

HasMaintenanceInfo returns a boolean if a field has been set.

### SetMaintenanceInfoNil

`func (o *MoiraCheckData) SetMaintenanceInfoNil(b bool)`

 SetMaintenanceInfoNil sets the value for MaintenanceInfo to be an explicit nil

### UnsetMaintenanceInfo
`func (o *MoiraCheckData) UnsetMaintenanceInfo()`

UnsetMaintenanceInfo ensures that no value is present for MaintenanceInfo, not even an explicit nil
### GetMetrics

`func (o *MoiraCheckData) GetMetrics() map[string]MoiraMetricState`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *MoiraCheckData) GetMetricsOk() (*map[string]MoiraMetricState, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *MoiraCheckData) SetMetrics(v map[string]MoiraMetricState)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *MoiraCheckData) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetMetricsToTargetRelation

`func (o *MoiraCheckData) GetMetricsToTargetRelation() map[string]string`

GetMetricsToTargetRelation returns the MetricsToTargetRelation field if non-nil, zero value otherwise.

### GetMetricsToTargetRelationOk

`func (o *MoiraCheckData) GetMetricsToTargetRelationOk() (*map[string]string, bool)`

GetMetricsToTargetRelationOk returns a tuple with the MetricsToTargetRelation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsToTargetRelation

`func (o *MoiraCheckData) SetMetricsToTargetRelation(v map[string]string)`

SetMetricsToTargetRelation sets MetricsToTargetRelation field to given value.

### HasMetricsToTargetRelation

`func (o *MoiraCheckData) HasMetricsToTargetRelation() bool`

HasMetricsToTargetRelation returns a boolean if a field has been set.

### GetMsg

`func (o *MoiraCheckData) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *MoiraCheckData) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *MoiraCheckData) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *MoiraCheckData) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetScore

`func (o *MoiraCheckData) GetScore() int64`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *MoiraCheckData) GetScoreOk() (*int64, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *MoiraCheckData) SetScore(v int64)`

SetScore sets Score field to given value.

### HasScore

`func (o *MoiraCheckData) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetState

`func (o *MoiraCheckData) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MoiraCheckData) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MoiraCheckData) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *MoiraCheckData) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSuppressed

`func (o *MoiraCheckData) GetSuppressed() bool`

GetSuppressed returns the Suppressed field if non-nil, zero value otherwise.

### GetSuppressedOk

`func (o *MoiraCheckData) GetSuppressedOk() (*bool, bool)`

GetSuppressedOk returns a tuple with the Suppressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressed

`func (o *MoiraCheckData) SetSuppressed(v bool)`

SetSuppressed sets Suppressed field to given value.

### HasSuppressed

`func (o *MoiraCheckData) HasSuppressed() bool`

HasSuppressed returns a boolean if a field has been set.

### GetSuppressedState

`func (o *MoiraCheckData) GetSuppressedState() string`

GetSuppressedState returns the SuppressedState field if non-nil, zero value otherwise.

### GetSuppressedStateOk

`func (o *MoiraCheckData) GetSuppressedStateOk() (*string, bool)`

GetSuppressedStateOk returns a tuple with the SuppressedState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressedState

`func (o *MoiraCheckData) SetSuppressedState(v string)`

SetSuppressedState sets SuppressedState field to given value.

### HasSuppressedState

`func (o *MoiraCheckData) HasSuppressedState() bool`

HasSuppressedState returns a boolean if a field has been set.

### GetTimestamp

`func (o *MoiraCheckData) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *MoiraCheckData) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *MoiraCheckData) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *MoiraCheckData) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


