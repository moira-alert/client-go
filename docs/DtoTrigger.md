# DtoTrigger

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AloneMetrics** | **map[string]bool** | A list of targets that have only alone metrics | [optional] [default to null]
**CreatedAt** | **string** | Datetime when the trigger was created | [optional] [default to null]
**CreatedBy** | **string** | Username who created trigger | [optional] [default to null]
**Desc** | **string** | Description string | [optional] [default to null]
**ErrorValue** | **float64** | ERROR threshold | [optional] [default to null]
**Expression** | **string** | Used if you need more complex logic than provided by WARN/ERROR values | [optional] [default to null]
**Id** | **string** | Trigger unique ID | [optional] [default to null]
**IsRemote** | **bool** | Shows if trigger is remote (graphite-backend) based or stored inside Moira-Redis DB | [optional] [default to null]
**MuteNewMetrics** | **bool** | If true, first event NODATA → OK will be omitted | [optional] [default to null]
**Name** | **string** | Trigger name | [optional] [default to null]
**Patterns** | **[]string** | Graphite patterns for trigger | [optional] [default to null]
**Sched** | [***AllOfdtoTriggerSched**](AllOfdtoTriggerSched.md) | Determines when Moira should monitor trigger | [optional] [default to null]
**Tags** | **[]string** | Set of tags to manipulate subscriptions | [optional] [default to null]
**Targets** | **[]string** | Graphite-like targets: t1, t2, ... | [optional] [default to null]
**Throttling** | **int32** |  | [optional] [default to null]
**TriggerType** | **string** | Could be: rising, falling, expression | [optional] [default to null]
**Ttl** | **int32** | When there are no metrics for trigger, Moira will switch metric to TTLState state after TTL seconds | [optional] [default to null]
**TtlState** | **string** | When there are no metrics for trigger, Moira will switch metric to TTLState state after TTL seconds | [optional] [default to null]
**UpdatedAt** | **string** | Datetime  when the trigger was updated | [optional] [default to null]
**UpdatedBy** | **string** | Username who updated trigger | [optional] [default to null]
**WarnValue** | **float64** | WARN threshold | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

