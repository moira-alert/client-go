# internal\SubscriptionAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSubscription**](SubscriptionAPI.md#CreateSubscription) | **Put** /subscription | Create a new subscription
[**GetSubscription**](SubscriptionAPI.md#GetSubscription) | **Get** /subscription/{subscriptionID} | Get subscription by id
[**GetUserSubscriptions**](SubscriptionAPI.md#GetUserSubscriptions) | **Get** /subscription | Get all subscriptions
[**RemoveSubscription**](SubscriptionAPI.md#RemoveSubscription) | **Delete** /subscription/{subscriptionID} | Delete a subscription
[**SendTestNotification**](SubscriptionAPI.md#SendTestNotification) | **Put** /subscription/{subscriptionID}/test | Send a test notification for a subscription
[**UpdateSubscription**](SubscriptionAPI.md#UpdateSubscription) | **Put** /subscription/{subscriptionID} | Update a subscription



## CreateSubscription

> DtoSubscription CreateSubscription(ctx).DtoSubscription(dtoSubscription).Execute()

Create a new subscription

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
	dtoSubscription := *openapiclient.NewDtoSubscription() // DtoSubscription | Subscription data

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionAPI.CreateSubscription(context.Background()).DtoSubscription(dtoSubscription).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionAPI.CreateSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSubscription`: DtoSubscription
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionAPI.CreateSubscription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dtoSubscription** | [**DtoSubscription**](DtoSubscription.md) | Subscription data | 

### Return type

[**DtoSubscription**](DtoSubscription.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubscription

> DtoSubscription GetSubscription(ctx, subscriptionID).Execute()

Get subscription by id

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
	subscriptionID := "subscriptionID_example" // string | ID of the subscription to get (default to "bcba82f5-48cf-44c0-b7d6-e1d32c64a88c")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionAPI.GetSubscription(context.Background(), subscriptionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionAPI.GetSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubscription`: DtoSubscription
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionAPI.GetSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionID** | **string** | ID of the subscription to get | [default to &quot;bcba82f5-48cf-44c0-b7d6-e1d32c64a88c&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DtoSubscription**](DtoSubscription.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserSubscriptions

> DtoSubscriptionList GetUserSubscriptions(ctx).Execute()

Get all subscriptions

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
	resp, r, err := apiClient.SubscriptionAPI.GetUserSubscriptions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionAPI.GetUserSubscriptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserSubscriptions`: DtoSubscriptionList
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionAPI.GetUserSubscriptions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserSubscriptionsRequest struct via the builder pattern


### Return type

[**DtoSubscriptionList**](DtoSubscriptionList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveSubscription

> map[string]interface{} RemoveSubscription(ctx, subscriptionID).Execute()

Delete a subscription

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
	subscriptionID := "subscriptionID_example" // string | ID of the subscription to remove (default to "bcba82f5-48cf-44c0-b7d6-e1d32c64a88c")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionAPI.RemoveSubscription(context.Background(), subscriptionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionAPI.RemoveSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveSubscription`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionAPI.RemoveSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionID** | **string** | ID of the subscription to remove | [default to &quot;bcba82f5-48cf-44c0-b7d6-e1d32c64a88c&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendTestNotification

> map[string]interface{} SendTestNotification(ctx, subscriptionID).Execute()

Send a test notification for a subscription

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
	subscriptionID := "subscriptionID_example" // string | ID of the subscription to send the test notification (default to "bcba82f5-48cf-44c0-b7d6-e1d32c64a88c")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionAPI.SendTestNotification(context.Background(), subscriptionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionAPI.SendTestNotification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendTestNotification`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionAPI.SendTestNotification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionID** | **string** | ID of the subscription to send the test notification | [default to &quot;bcba82f5-48cf-44c0-b7d6-e1d32c64a88c&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiSendTestNotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSubscription

> DtoSubscription UpdateSubscription(ctx, subscriptionID).DtoSubscription(dtoSubscription).Execute()

Update a subscription

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
	subscriptionID := "subscriptionID_example" // string | ID of the subscription to update (default to "bcba82f5-48cf-44c0-b7d6-e1d32c64a88c")
	dtoSubscription := *openapiclient.NewDtoSubscription() // DtoSubscription | Updated subscription data

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionAPI.UpdateSubscription(context.Background(), subscriptionID).DtoSubscription(dtoSubscription).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionAPI.UpdateSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSubscription`: DtoSubscription
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionAPI.UpdateSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionID** | **string** | ID of the subscription to update | [default to &quot;bcba82f5-48cf-44c0-b7d6-e1d32c64a88c&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dtoSubscription** | [**DtoSubscription**](DtoSubscription.md) | Updated subscription data | 

### Return type

[**DtoSubscription**](DtoSubscription.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

