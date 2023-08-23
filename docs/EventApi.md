# {{classname}}

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAllEvents**](EventApi.md#DeleteAllEvents) | **Delete** /event/all | Deletes all notification events
[**GetEventsList**](EventApi.md#GetEventsList) | **Get** /event/{triggerID} | Gets all trigger events for current page and their count

# **DeleteAllEvents**
> DeleteAllEvents(ctx, )
Deletes all notification events

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEventsList**
> DtoEventsList GetEventsList(ctx, triggerID, optional)
Gets all trigger events for current page and their count

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **triggerID** | **string**| The ID of updated trigger | 
 **optional** | ***EventApiGetEventsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EventApiGetEventsListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **size** | **optional.Int32**| Number of items to be displayed on one page | [default to 100]
 **p** | **optional.Int32**| Defines the number of the displayed page. E.g, p&#x3D;2 would display the 2nd page | [default to 0]

### Return type

[**DtoEventsList**](dto.EventsList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

