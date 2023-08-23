# MoiraCheckData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventTimestamp** | **int32** |  | [optional] [default to null]
**LastSuccessfulCheckTimestamp** | **int32** |  | [optional] [default to null]
**Maintenance** | **int32** |  | [optional] [default to null]
**MaintenanceInfo** | [***MoiraMaintenanceInfo**](moira.MaintenanceInfo.md) |  | [optional] [default to null]
**Metrics** | [**map[string]MoiraMetricState**](moira.MetricState.md) |  | [optional] [default to null]
**MetricsToTargetRelation** | **map[string]string** | MetricsToTargetRelation is a map that holds relation between metric names that was alone during last check and targets that fetched this metric  {\&quot;t1\&quot;: \&quot;metric.name.1\&quot;, \&quot;t2\&quot;: \&quot;metric.name.2\&quot;} | [optional] [default to null]
**Msg** | **string** |  | [optional] [default to null]
**Score** | **int32** |  | [optional] [default to null]
**State** | **string** |  | [optional] [default to null]
**Suppressed** | **bool** |  | [optional] [default to null]
**SuppressedState** | **string** |  | [optional] [default to null]
**Timestamp** | **int32** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

