# {{classname}}

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNewContact**](ContactApi.md#CreateNewContact) | **Put** /contact | Creates a new contact notification for the current user
[**GetAllContacts**](ContactApi.md#GetAllContacts) | **Get** /contact | Gets all Moira contacts
[**GetContactById**](ContactApi.md#GetContactById) | **Get** /contact/{contactID} | Get contact by ID
[**GetContactEventsById**](ContactApi.md#GetContactEventsById) | **Get** /contact/{contactID}/events | Get contact events by ID with time range
[**RemoveContact**](ContactApi.md#RemoveContact) | **Delete** /contact/{contactID} | Deletes notification contact for the current user and remove the contact ID from all subscriptions
[**SendTestContactNotification**](ContactApi.md#SendTestContactNotification) | **Post** /contact/{contactID}/test | Push a test notification to verify that the contact is properly set up
[**UpdateContact**](ContactApi.md#UpdateContact) | **Put** /contact/{contactID} | Updates an existing notification contact to the values passed in the request body

# **CreateNewContact**
> DtoContact CreateNewContact(ctx, body)
Creates a new contact notification for the current user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DtoContact**](DtoContact.md)| Contact data | 

### Return type

[**DtoContact**](dto.Contact.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllContacts**
> DtoContactList GetAllContacts(ctx, )
Gets all Moira contacts

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**DtoContactList**](dto.ContactList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContactById**
> DtoContact GetContactById(ctx, contactID)
Get contact by ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contactID** | **string**| Contact ID | 

### Return type

[**DtoContact**](dto.Contact.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContactEventsById**
> DtoContactEventItemList GetContactEventsById(ctx, contactID, optional)
Get contact events by ID with time range

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contactID** | **string**| Contact ID | 
 **optional** | ***ContactApiGetContactEventsByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContactApiGetContactEventsByIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **optional.String**| Start time of the time range | [default to -3hour]
 **to** | **optional.String**| End time of the time range | [default to now]

### Return type

[**DtoContactEventItemList**](dto.ContactEventItemList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveContact**
> RemoveContact(ctx, contactID)
Deletes notification contact for the current user and remove the contact ID from all subscriptions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contactID** | **string**| ID of the contact to remove | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendTestContactNotification**
> SendTestContactNotification(ctx, contactID)
Push a test notification to verify that the contact is properly set up

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contactID** | **string**| The ID of the target contact | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateContact**
> DtoContact UpdateContact(ctx, body, contactID)
Updates an existing notification contact to the values passed in the request body

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DtoContact**](DtoContact.md)| Updated contact data | 
  **contactID** | **string**| ID of the contact to update | 

### Return type

[**DtoContact**](dto.Contact.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

