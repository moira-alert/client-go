# {{classname}}

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTrigger**](TriggerApi.md#CreateTrigger) | **Put** /trigger | Create a new trigger
[**DeleteTriggerMetric**](TriggerApi.md#DeleteTriggerMetric) | **Delete** /trigger/{triggerID}/metrics | Delete metric from last check and all trigger pattern metrics
[**DeleteTriggerNodataMetrics**](TriggerApi.md#DeleteTriggerNodataMetrics) | **Delete** /trigger/{triggerID}/metrics/nodata | Delete all metrics from last data which are in NODATA state. It also deletes all trigger patterns of those metrics
[**DeleteTriggerThrottling**](TriggerApi.md#DeleteTriggerThrottling) | **Delete** /trigger/{triggerID}/throttling | Deletes throttling for a trigger
[**GetAllTriggers**](TriggerApi.md#GetAllTriggers) | **Get** /trigger | Get all triggers
[**GetTrigger**](TriggerApi.md#GetTrigger) | **Get** /trigger/{triggerID} | Get an existing trigger
[**GetTriggerDump**](TriggerApi.md#GetTriggerDump) | **Get** /trigger/{triggerID}/dump | Get trigger dump
[**GetTriggerMetrics**](TriggerApi.md#GetTriggerMetrics) | **Get** /trigger/{triggerID}/metrics | Get metrics associated with certain trigger
[**GetTriggerState**](TriggerApi.md#GetTriggerState) | **Get** /trigger/{triggerID}/state | Get the trigger state as at last check
[**GetTriggerThrottling**](TriggerApi.md#GetTriggerThrottling) | **Get** /trigger/{triggerID}/throttling | Get a trigger with its throttling i.e its next allowed message time
[**RemoveTrigger**](TriggerApi.md#RemoveTrigger) | **Delete** /trigger/{triggerID} | Remove trigger
[**RenderTriggerMetrics**](TriggerApi.md#RenderTriggerMetrics) | **Get** /trigger/{triggerID}/render | Render trigger metrics plot
[**SearchTriggers**](TriggerApi.md#SearchTriggers) | **Get** /trigger/search | Search triggers. Replaces the deprecated &#x60;page&#x60; path
[**SetTriggerMaintenance**](TriggerApi.md#SetTriggerMaintenance) | **Put** /trigger/{triggerID}/setMaintenance | Set metrics and the trigger itself to maintenance mode
[**TriggerCheck**](TriggerApi.md#TriggerCheck) | **Put** /trigger/check | Validates trigger target
[**TriggerPagersPagerIDDelete**](TriggerApi.md#TriggerPagersPagerIDDelete) | **Delete** /trigger/pagers/{pagerID} | Delete triggers pager
[**UpdateTrigger**](TriggerApi.md#UpdateTrigger) | **Put** /trigger/{triggerID} | Update existing trigger

# **CreateTrigger**
> DtoSaveTriggerResponse CreateTrigger(ctx, body, optional)
Create a new trigger

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DtoTrigger**](DtoTrigger.md)| Trigger data | 
 **optional** | ***TriggerApiCreateTriggerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TriggerApiCreateTriggerOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **validate** | **optional.**| For validating targets | 

### Return type

[**DtoSaveTriggerResponse**](dto.SaveTriggerResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTriggerMetric**
> DeleteTriggerMetric(ctx, triggerID, optional)
Delete metric from last check and all trigger pattern metrics

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **triggerID** | **string**| Trigger ID | 
 **optional** | ***TriggerApiDeleteTriggerMetricOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TriggerApiDeleteTriggerMetricOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| Name of the target metric | [default to DevOps.my_server.hdd.freespace_mbytes]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTriggerNodataMetrics**
> DeleteTriggerNodataMetrics(ctx, triggerID)
Delete all metrics from last data which are in NODATA state. It also deletes all trigger patterns of those metrics

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **triggerID** | **string**| Trigger ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTriggerThrottling**
> DeleteTriggerThrottling(ctx, triggerID)
Deletes throttling for a trigger

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **triggerID** | **string**| Trigger ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllTriggers**
> DtoTriggersList GetAllTriggers(ctx, )
Get all triggers

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**DtoTriggersList**](dto.TriggersList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTrigger**
> DtoTrigger GetTrigger(ctx, triggerID, optional)
Get an existing trigger

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **triggerID** | **string**| Trigger ID | 
 **optional** | ***TriggerApiGetTriggerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TriggerApiGetTriggerOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **populated** | **optional.Bool**| Populated | [default to false]

### Return type

[**DtoTrigger**](dto.Trigger.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTriggerDump**
> DtoTriggerDump GetTriggerDump(ctx, triggerID)
Get trigger dump

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **triggerID** | **string**| Trigger ID | 

### Return type

[**DtoTriggerDump**](dto.TriggerDump.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTriggerMetrics**
> map[string]map[string][]MoiraMetricValue GetTriggerMetrics(ctx, triggerID, optional)
Get metrics associated with certain trigger

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **triggerID** | **string**| Trigger ID | 
 **optional** | ***TriggerApiGetTriggerMetricsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TriggerApiGetTriggerMetricsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **optional.String**| Start time for metrics retrieval | [default to -10minutes]
 **to** | **optional.String**| End time for metrics retrieval | [default to now]

### Return type

[**map[string]map[string][]MoiraMetricValue**](map.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTriggerState**
> DtoTriggerCheck GetTriggerState(ctx, triggerID)
Get the trigger state as at last check

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **triggerID** | **string**| Trigger ID | 

### Return type

[**DtoTriggerCheck**](dto.TriggerCheck.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTriggerThrottling**
> DtoThrottlingResponse GetTriggerThrottling(ctx, triggerID)
Get a trigger with its throttling i.e its next allowed message time

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **triggerID** | **string**| Trigger ID | 

### Return type

[**DtoThrottlingResponse**](dto.ThrottlingResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveTrigger**
> RemoveTrigger(ctx, triggerID)
Remove trigger

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **triggerID** | **string**| Trigger ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RenderTriggerMetrics**
> RenderTriggerMetrics(ctx, triggerID, optional)
Render trigger metrics plot

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **triggerID** | **string**| Trigger ID | 
 **optional** | ***TriggerApiRenderTriggerMetricsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TriggerApiRenderTriggerMetricsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **target** | **optional.String**| Target metric name | [default to t1]
 **from** | **optional.String**| Start time for metrics retrieval | [default to -1hour]
 **to** | **optional.String**| End time for metrics retrieval | [default to now]
 **timezone** | **optional.String**| Timezone for rendering | [default to UTC]
 **theme** | **optional.String**| Plot theme | [default to light]
 **realtime** | **optional.Bool**| Fetch real-time data | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: image/png

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchTriggers**
> DtoTriggersList SearchTriggers(ctx, optional)
Search triggers. Replaces the deprecated `page` path

You can also add filtering by tags, for this purpose add query parameters tags[0]=test, tags[1]=test1 and so on For example, `/api/trigger/search?tags[0]=test&tags[1]=test1`

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TriggerApiSearchTriggersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TriggerApiSearchTriggersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **onlyProblems** | **optional.Bool**| Only include problems | [default to false]
 **text** | **optional.String**| Search text | [default to cpu]
 **p** | **optional.Int32**| Page number | [default to 0]
 **size** | **optional.Int32**| Page size | [default to 10]
 **createPager** | **optional.Bool**| Create pager | [default to false]
 **pagerID** | **optional.String**| Pager ID | [default to bcba82f5-48cf-44c0-b7d6-e1d32c64a88c]

### Return type

[**DtoTriggersList**](dto.TriggersList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetTriggerMaintenance**
> SetTriggerMaintenance(ctx, body, triggerID)
Set metrics and the trigger itself to maintenance mode

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DtoTriggerMaintenance**](DtoTriggerMaintenance.md)| Maintenance data | 
  **triggerID** | **string**| Trigger ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TriggerCheck**
> DtoTriggerCheckResponse TriggerCheck(ctx, body)
Validates trigger target

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DtoTrigger**](DtoTrigger.md)| Trigger data | 

### Return type

[**DtoTriggerCheckResponse**](dto.TriggerCheckResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TriggerPagersPagerIDDelete**
> DtoTriggersSearchResultDeleteResponse TriggerPagersPagerIDDelete(ctx, pagerID)
Delete triggers pager

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pagerID** | **string**| Pager ID to delete | 

### Return type

[**DtoTriggersSearchResultDeleteResponse**](dto.TriggersSearchResultDeleteResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTrigger**
> DtoSaveTriggerResponse UpdateTrigger(ctx, body, triggerID, optional)
Update existing trigger

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DtoTrigger**](DtoTrigger.md)| Trigger data | 
  **triggerID** | **string**| Trigger ID | 
 **optional** | ***TriggerApiUpdateTriggerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TriggerApiUpdateTriggerOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **validate** | **optional.**| For validating targets | 

### Return type

[**DtoSaveTriggerResponse**](dto.SaveTriggerResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

