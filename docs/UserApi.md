# {{classname}}

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUserName**](UserApi.md#GetUserName) | **Get** /user | Gets the username of the authenticated user if it is available
[**GetUserSettings**](UserApi.md#GetUserSettings) | **Get** /user/settings | Get the user&#x27;s contacts and subscriptions

# **GetUserName**
> DtoUser GetUserName(ctx, )
Gets the username of the authenticated user if it is available

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**DtoUser**](dto.User.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserSettings**
> DtoUserSettings GetUserSettings(ctx, )
Get the user's contacts and subscriptions

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**DtoUserSettings**](dto.UserSettings.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

