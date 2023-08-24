# {{classname}}

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNewTeamContact**](TeamContactApi.md#CreateNewTeamContact) | **Post** /teams/{teamID}/contacts | Create a new team contact

# **CreateNewTeamContact**
> DtoContact CreateNewTeamContact(ctx, body, teamID)
Create a new team contact

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DtoContact**](DtoContact.md)| Team contact data | 
  **teamID** | **string**| The ID of team | 

### Return type

[**DtoContact**](dto.Contact.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

