# {{classname}}

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllTags**](TagApi.md#GetAllTags) | **Get** /tag | Get all tags
[**GetAllTagsAndSubscriptions**](TagApi.md#GetAllTagsAndSubscriptions) | **Get** /tag/stats | Get all tags and their subscriptions
[**RemoveTag**](TagApi.md#RemoveTag) | **Delete** /tag/{tag} | Remove a tag

# **GetAllTags**
> DtoTagsData GetAllTags(ctx, )
Get all tags

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**DtoTagsData**](dto.TagsData.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllTagsAndSubscriptions**
> DtoTagsStatistics GetAllTagsAndSubscriptions(ctx, )
Get all tags and their subscriptions

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**DtoTagsStatistics**](dto.TagsStatistics.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveTag**
> DtoMessageResponse RemoveTag(ctx, tag)
Remove a tag

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tag** | **string**| Name of the tag to remove | 

### Return type

[**DtoMessageResponse**](dto.MessageResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

