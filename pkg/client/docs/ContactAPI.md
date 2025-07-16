# internal\ContactAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNewContact**](ContactAPI.md#CreateNewContact) | **Put** /contact | Creates a new contact notification for the current user
[**GetAllContacts**](ContactAPI.md#GetAllContacts) | **Get** /contact | Gets all Moira contacts
[**GetContactById**](ContactAPI.md#GetContactById) | **Get** /contact/{contactID} | Get contact by ID
[**GetContactEventsById**](ContactAPI.md#GetContactEventsById) | **Get** /contact/{contactID}/events | Get contact events by ID with time range
[**GetContactsNoisiness**](ContactAPI.md#GetContactsNoisiness) | **Get** /contact/noisiness | Get contacts noisiness
[**RemoveContact**](ContactAPI.md#RemoveContact) | **Delete** /contact/{contactID} | Deletes notification contact for the current user and remove the contact ID from all subscriptions
[**SendTestContactNotification**](ContactAPI.md#SendTestContactNotification) | **Post** /contact/{contactID}/test | Push a test notification to verify that the contact is properly set up
[**UpdateContact**](ContactAPI.md#UpdateContact) | **Put** /contact/{contactID} | Updates an existing notification contact to the values passed in the request body



## CreateNewContact

> DtoContact CreateNewContact(ctx).DtoContact(dtoContact).Execute()

Creates a new contact notification for the current user

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	dtoContact := *openapiclient.NewDtoContact() // DtoContact | Contact data

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactAPI.CreateNewContact(context.Background()).DtoContact(dtoContact).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactAPI.CreateNewContact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNewContact`: DtoContact
	fmt.Fprintf(os.Stdout, "Response from `ContactAPI.CreateNewContact`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNewContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dtoContact** | [**DtoContact**](DtoContact.md) | Contact data | 

### Return type

[**DtoContact**](DtoContact.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllContacts

> DtoContactList GetAllContacts(ctx).Execute()

Gets all Moira contacts

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactAPI.GetAllContacts(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactAPI.GetAllContacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllContacts`: DtoContactList
	fmt.Fprintf(os.Stdout, "Response from `ContactAPI.GetAllContacts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllContactsRequest struct via the builder pattern


### Return type

[**DtoContactList**](DtoContactList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContactById

> DtoContact GetContactById(ctx, contactID).Execute()

Get contact by ID

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	contactID := "contactID_example" // string | Contact ID (default to "bcba82f5-48cf-44c0-b7d6-e1d32c64a88c")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactAPI.GetContactById(context.Background(), contactID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactAPI.GetContactById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContactById`: DtoContact
	fmt.Fprintf(os.Stdout, "Response from `ContactAPI.GetContactById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactID** | **string** | Contact ID | [default to &quot;bcba82f5-48cf-44c0-b7d6-e1d32c64a88c&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetContactByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DtoContact**](DtoContact.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContactEventsById

> DtoContactEventItemList GetContactEventsById(ctx, contactID).From(from).To(to).Size(size).P(p).Execute()

Get contact events by ID with time range

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	contactID := "contactID_example" // string | Contact ID (default to "bcba82f5-48cf-44c0-b7d6-e1d32c64a88c")
	from := "from_example" // string | Start time of the time range (optional) (default to "-3hour")
	to := "to_example" // string | End time of the time range (optional) (default to "now")
	size := int32(56) // int32 | Number of items to return or all items if size == -1 (if size == -1 p should be zero for correct work) (optional) (default to 100)
	p := int32(56) // int32 | Defines the index of data portion (combined with size). E.g, p=2, size=100 will return records from 200 (including), to 300 (not including) (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactAPI.GetContactEventsById(context.Background(), contactID).From(from).To(to).Size(size).P(p).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactAPI.GetContactEventsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContactEventsById`: DtoContactEventItemList
	fmt.Fprintf(os.Stdout, "Response from `ContactAPI.GetContactEventsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactID** | **string** | Contact ID | [default to &quot;bcba82f5-48cf-44c0-b7d6-e1d32c64a88c&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetContactEventsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **string** | Start time of the time range | [default to &quot;-3hour&quot;]
 **to** | **string** | End time of the time range | [default to &quot;now&quot;]
 **size** | **int32** | Number of items to return or all items if size &#x3D;&#x3D; -1 (if size &#x3D;&#x3D; -1 p should be zero for correct work) | [default to 100]
 **p** | **int32** | Defines the index of data portion (combined with size). E.g, p&#x3D;2, size&#x3D;100 will return records from 200 (including), to 300 (not including) | [default to 0]

### Return type

[**DtoContactEventItemList**](DtoContactEventItemList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContactsNoisiness

> DtoContactNoisinessList GetContactsNoisiness(ctx).Size(size).P(p).From(from).To(to).Sort(sort).Execute()

Get contacts noisiness

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	size := int32(56) // int32 | Number of items to be displayed on one page. if size = -1 then all events returned (optional) (default to 100)
	p := int32(56) // int32 | Defines the number of the displayed page. E.g, p=2 would display the 2nd page (optional) (default to 0)
	from := "from_example" // string | Start time of the time range (optional) (default to "-3hours")
	to := "to_example" // string | End time of the time range (optional) (default to "now")
	sort := "sort_example" // string | String to set sort order (by events_count). On empty - no order, asc - ascending, desc - descending (optional) (default to "desc")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactAPI.GetContactsNoisiness(context.Background()).Size(size).P(p).From(from).To(to).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactAPI.GetContactsNoisiness``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContactsNoisiness`: DtoContactNoisinessList
	fmt.Fprintf(os.Stdout, "Response from `ContactAPI.GetContactsNoisiness`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetContactsNoisinessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32** | Number of items to be displayed on one page. if size &#x3D; -1 then all events returned | [default to 100]
 **p** | **int32** | Defines the number of the displayed page. E.g, p&#x3D;2 would display the 2nd page | [default to 0]
 **from** | **string** | Start time of the time range | [default to &quot;-3hours&quot;]
 **to** | **string** | End time of the time range | [default to &quot;now&quot;]
 **sort** | **string** | String to set sort order (by events_count). On empty - no order, asc - ascending, desc - descending | [default to &quot;desc&quot;]

### Return type

[**DtoContactNoisinessList**](DtoContactNoisinessList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveContact

> map[string]interface{} RemoveContact(ctx, contactID).Body(body).Execute()

Deletes notification contact for the current user and remove the contact ID from all subscriptions

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	contactID := "contactID_example" // string | ID of the contact to remove (default to "bcba82f5-48cf-44c0-b7d6-e1d32c64a88c")
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactAPI.RemoveContact(context.Background(), contactID).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactAPI.RemoveContact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveContact`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ContactAPI.RemoveContact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactID** | **string** | ID of the contact to remove | [default to &quot;bcba82f5-48cf-44c0-b7d6-e1d32c64a88c&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendTestContactNotification

> map[string]interface{} SendTestContactNotification(ctx, contactID).Body(body).Execute()

Push a test notification to verify that the contact is properly set up

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	contactID := "contactID_example" // string | The ID of the target contact (default to "bcba82f5-48cf-44c0-b7d6-e1d32c64a88c")
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactAPI.SendTestContactNotification(context.Background(), contactID).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactAPI.SendTestContactNotification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendTestContactNotification`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ContactAPI.SendTestContactNotification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactID** | **string** | The ID of the target contact | [default to &quot;bcba82f5-48cf-44c0-b7d6-e1d32c64a88c&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiSendTestContactNotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateContact

> DtoContact UpdateContact(ctx, contactID).DtoContact(dtoContact).Execute()

Updates an existing notification contact to the values passed in the request body

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	contactID := "contactID_example" // string | ID of the contact to update (default to "bcba82f5-48cf-44c0-b7d6-e1d32c64a88c")
	dtoContact := *openapiclient.NewDtoContact() // DtoContact | Updated contact data

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactAPI.UpdateContact(context.Background(), contactID).DtoContact(dtoContact).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactAPI.UpdateContact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateContact`: DtoContact
	fmt.Fprintf(os.Stdout, "Response from `ContactAPI.UpdateContact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactID** | **string** | ID of the contact to update | [default to &quot;bcba82f5-48cf-44c0-b7d6-e1d32c64a88c&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dtoContact** | [**DtoContact**](DtoContact.md) | Updated contact data | 

### Return type

[**DtoContact**](DtoContact.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

