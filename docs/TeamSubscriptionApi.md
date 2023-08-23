# {{classname}}

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNewTeamSubscription**](TeamSubscriptionApi.md#CreateNewTeamSubscription) | **Post** /teams/{teamID}/subscriptions | Create a new team subscription

# **CreateNewTeamSubscription**
> DtoSubscription CreateNewTeamSubscription(ctx, body, teamID)
Create a new team subscription

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DtoSubscription**](DtoSubscription.md)| Team subscription data | 
  **teamID** | **string**| The ID of team | 

### Return type

[**DtoSubscription**](dto.Subscription.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

