# {{classname}}

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteNotification**](NotificationApi.md#DeleteNotification) | **Delete** /notification | Delete a notification by id
[**GetNotifications**](NotificationApi.md#GetNotifications) | **Get** /notification | Gets a paginated list of notifications, all notifications are fetched if end &#x3D; -1 and start &#x3D; 0

# **DeleteNotification**
> DtoNotificationDeleteResponse DeleteNotification(ctx, id)
Delete a notification by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of deleted notification | [default to bcba82f5-48cf-44c0-b7d6-e1d32c64a88c]

### Return type

[**DtoNotificationDeleteResponse**](dto.NotificationDeleteResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNotifications**
> DtoNotificationsList GetNotifications(ctx, optional)
Gets a paginated list of notifications, all notifications are fetched if end = -1 and start = 0

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NotificationApiGetNotificationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotificationApiGetNotificationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **optional.Int32**| Default Value: 0 | [default to 0]
 **end** | **optional.Int32**| Default Value: -1 | [default to -1]

### Return type

[**DtoNotificationsList**](dto.NotificationsList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

