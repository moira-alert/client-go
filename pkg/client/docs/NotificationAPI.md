# internal\NotificationAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAllNotifications**](NotificationAPI.md#DeleteAllNotifications) | **Delete** /notification/all | Delete all notifications
[**DeleteNotification**](NotificationAPI.md#DeleteNotification) | **Delete** /notification | Delete a notification by id
[**DeleteNotificationsFiltered**](NotificationAPI.md#DeleteNotificationsFiltered) | **Delete** /notification/filtered | Delete notifications filtered by tags and timestamps
[**GetNotifications**](NotificationAPI.md#GetNotifications) | **Get** /notification | Gets a paginated list of notifications, all notifications are fetched if end &#x3D; -1 and start &#x3D; 0



## DeleteAllNotifications

> DtoNotificationsList DeleteAllNotifications(ctx).Execute()

Delete all notifications

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
	resp, r, err := apiClient.NotificationAPI.DeleteAllNotifications(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationAPI.DeleteAllNotifications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAllNotifications`: DtoNotificationsList
	fmt.Fprintf(os.Stdout, "Response from `NotificationAPI.DeleteAllNotifications`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAllNotificationsRequest struct via the builder pattern


### Return type

[**DtoNotificationsList**](DtoNotificationsList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNotification

> DtoNotificationDeleteResponse DeleteNotification(ctx).Id(id).Execute()

Delete a notification by id

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
	id := "id_example" // string | The ID of deleted notification (default to "bcba82f5-48cf-44c0-b7d6-e1d32c64a88c")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationAPI.DeleteNotification(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationAPI.DeleteNotification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNotification`: DtoNotificationDeleteResponse
	fmt.Fprintf(os.Stdout, "Response from `NotificationAPI.DeleteNotification`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The ID of deleted notification | [default to &quot;bcba82f5-48cf-44c0-b7d6-e1d32c64a88c&quot;]

### Return type

[**DtoNotificationDeleteResponse**](DtoNotificationDeleteResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNotificationsFiltered

> DtoNotificationsList DeleteNotificationsFiltered(ctx).Execute()

Delete notifications filtered by tags and timestamps

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
	resp, r, err := apiClient.NotificationAPI.DeleteNotificationsFiltered(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationAPI.DeleteNotificationsFiltered``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNotificationsFiltered`: DtoNotificationsList
	fmt.Fprintf(os.Stdout, "Response from `NotificationAPI.DeleteNotificationsFiltered`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNotificationsFilteredRequest struct via the builder pattern


### Return type

[**DtoNotificationsList**](DtoNotificationsList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotifications

> DtoNotificationsList GetNotifications(ctx).Start(start).End(end).Execute()

Gets a paginated list of notifications, all notifications are fetched if end = -1 and start = 0

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
	start := int32(56) // int32 | Default Value: 0 (optional) (default to 0)
	end := int32(56) // int32 | Default Value: -1 (optional) (default to -1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationAPI.GetNotifications(context.Background()).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationAPI.GetNotifications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNotifications`: DtoNotificationsList
	fmt.Fprintf(os.Stdout, "Response from `NotificationAPI.GetNotifications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **int32** | Default Value: 0 | [default to 0]
 **end** | **int32** | Default Value: -1 | [default to -1]

### Return type

[**DtoNotificationsList**](DtoNotificationsList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

