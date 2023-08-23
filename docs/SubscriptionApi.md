# {{classname}}

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSubscription**](SubscriptionApi.md#CreateSubscription) | **Put** /subscription | Create a new subscription
[**GetUserSubscriptions**](SubscriptionApi.md#GetUserSubscriptions) | **Get** /subscription | Get all subscriptions
[**RemoveSubscription**](SubscriptionApi.md#RemoveSubscription) | **Delete** /subscription/{subscriptionID} | Delete a subscription
[**SendTestNotification**](SubscriptionApi.md#SendTestNotification) | **Put** /subscription/{subscriptionID}/test | Send a test notification for a subscription
[**UpdateSubscription**](SubscriptionApi.md#UpdateSubscription) | **Put** /subscription/{subscriptionID} | Update a subscription

# **CreateSubscription**
> DtoSubscription CreateSubscription(ctx, body)
Create a new subscription

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DtoSubscription**](DtoSubscription.md)| Subscription data | 

### Return type

[**DtoSubscription**](dto.Subscription.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserSubscriptions**
> DtoSubscriptionList GetUserSubscriptions(ctx, )
Get all subscriptions

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**DtoSubscriptionList**](dto.SubscriptionList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveSubscription**
> RemoveSubscription(ctx, subscriptionID)
Delete a subscription

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subscriptionID** | **string**| ID of the subscription to remove | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendTestNotification**
> SendTestNotification(ctx, subscriptionID)
Send a test notification for a subscription

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subscriptionID** | **string**| ID of the subscription to send the test notification | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSubscription**
> DtoSubscription UpdateSubscription(ctx, body, subscriptionID)
Update a subscription

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DtoSubscription**](DtoSubscription.md)| Updated subscription data | 
  **subscriptionID** | **string**| ID of the subscription to update | 

### Return type

[**DtoSubscription**](dto.Subscription.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

