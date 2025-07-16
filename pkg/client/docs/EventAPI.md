# internal\EventAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAllEvents**](EventAPI.md#DeleteAllEvents) | **Delete** /event/all | Deletes all notification events
[**GetEventsList**](EventAPI.md#GetEventsList) | **Get** /event/{triggerID} | Gets all trigger events for current page and their count



## DeleteAllEvents

> map[string]interface{} DeleteAllEvents(ctx).Execute()

Deletes all notification events

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
	resp, r, err := apiClient.EventAPI.DeleteAllEvents(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventAPI.DeleteAllEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAllEvents`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `EventAPI.DeleteAllEvents`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAllEventsRequest struct via the builder pattern


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


## GetEventsList

> DtoEventsList GetEventsList(ctx, triggerID).Size(size).P(p).From(from).To(to).Metric(metric).States(states).Execute()

Gets all trigger events for current page and their count

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
	triggerID := "triggerID_example" // string | The ID of updated trigger (default to "bcba82f5-48cf-44c0-b7d6-e1d32c64a88c")
	size := int32(56) // int32 | Number of items to be displayed on one page. if size = -1 then all events returned (optional) (default to 100)
	p := int32(56) // int32 | Defines the number of the displayed page. E.g, p=2 would display the 2nd page (optional) (default to 0)
	from := "from_example" // string | Start time of the time range (optional) (default to "-3hours")
	to := "to_example" // string | End time of the time range (optional) (default to "now")
	metric := "metric_example" // string | Regular expression that will be used to filter events (optional) (default to ".*")
	states := []string{"Inner_example"} // []string | String of ',' separated state names. If empty then all states will be used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventAPI.GetEventsList(context.Background(), triggerID).Size(size).P(p).From(from).To(to).Metric(metric).States(states).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventAPI.GetEventsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEventsList`: DtoEventsList
	fmt.Fprintf(os.Stdout, "Response from `EventAPI.GetEventsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**triggerID** | **string** | The ID of updated trigger | [default to &quot;bcba82f5-48cf-44c0-b7d6-e1d32c64a88c&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **size** | **int32** | Number of items to be displayed on one page. if size &#x3D; -1 then all events returned | [default to 100]
 **p** | **int32** | Defines the number of the displayed page. E.g, p&#x3D;2 would display the 2nd page | [default to 0]
 **from** | **string** | Start time of the time range | [default to &quot;-3hours&quot;]
 **to** | **string** | End time of the time range | [default to &quot;now&quot;]
 **metric** | **string** | Regular expression that will be used to filter events | [default to &quot;.*&quot;]
 **states** | **[]string** | String of &#39;,&#39; separated state names. If empty then all states will be used. | 

### Return type

[**DtoEventsList**](DtoEventsList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

