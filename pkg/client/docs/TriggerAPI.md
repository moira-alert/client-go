# internal\TriggerAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTrigger**](TriggerAPI.md#CreateTrigger) | **Put** /trigger | Create a new trigger
[**DeletePager**](TriggerAPI.md#DeletePager) | **Delete** /trigger/search/pager | Delete triggers pager
[**DeleteTriggerMetric**](TriggerAPI.md#DeleteTriggerMetric) | **Delete** /trigger/{triggerID}/metrics | Delete metric from last check and all trigger pattern metrics
[**DeleteTriggerNodataMetrics**](TriggerAPI.md#DeleteTriggerNodataMetrics) | **Delete** /trigger/{triggerID}/metrics/nodata | Delete all metrics from last data which are in NODATA state. It also deletes all trigger patterns of those metrics
[**DeleteTriggerThrottling**](TriggerAPI.md#DeleteTriggerThrottling) | **Delete** /trigger/{triggerID}/throttling | Deletes throttling for a trigger
[**GetAllTriggers**](TriggerAPI.md#GetAllTriggers) | **Get** /trigger | Get all triggers
[**GetTrigger**](TriggerAPI.md#GetTrigger) | **Get** /trigger/{triggerID} | Get an existing trigger
[**GetTriggerDump**](TriggerAPI.md#GetTriggerDump) | **Get** /trigger/{triggerID}/dump | Get trigger dump
[**GetTriggerMetrics**](TriggerAPI.md#GetTriggerMetrics) | **Get** /trigger/{triggerID}/metrics | Get metrics associated with certain trigger
[**GetTriggerState**](TriggerAPI.md#GetTriggerState) | **Get** /trigger/{triggerID}/state | Get the trigger state as at last check
[**GetTriggerThrottling**](TriggerAPI.md#GetTriggerThrottling) | **Get** /trigger/{triggerID}/throttling | Get a trigger with its throttling i.e its next allowed message time
[**GetTriggersNoisiness**](TriggerAPI.md#GetTriggersNoisiness) | **Get** /trigger/noisiness | Get triggers noisiness
[**GetUnusedTriggers**](TriggerAPI.md#GetUnusedTriggers) | **Get** /trigger/unused | Get unused triggers
[**RemoveTrigger**](TriggerAPI.md#RemoveTrigger) | **Delete** /trigger/{triggerID} | Remove trigger
[**RenderTriggerMetrics**](TriggerAPI.md#RenderTriggerMetrics) | **Get** /trigger/{triggerID}/render | Render trigger metrics plot
[**SearchTriggers**](TriggerAPI.md#SearchTriggers) | **Get** /trigger/search | Search triggers. Replaces the deprecated &#x60;page&#x60; path
[**SetTriggerMaintenance**](TriggerAPI.md#SetTriggerMaintenance) | **Put** /trigger/{triggerID}/setMaintenance | Set metrics and the trigger itself to maintenance mode
[**TriggerCheck**](TriggerAPI.md#TriggerCheck) | **Put** /trigger/check | Validates trigger target
[**UpdateTrigger**](TriggerAPI.md#UpdateTrigger) | **Put** /trigger/{triggerID} | Update existing trigger



## CreateTrigger

> DtoSaveTriggerResponse CreateTrigger(ctx).DtoTrigger(dtoTrigger).Validate(validate).Execute()

Create a new trigger

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
	dtoTrigger := *openapiclient.NewDtoTrigger() // DtoTrigger | Trigger data
	validate := true // bool | For validating targets (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TriggerAPI.CreateTrigger(context.Background()).DtoTrigger(dtoTrigger).Validate(validate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TriggerAPI.CreateTrigger``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTrigger`: DtoSaveTriggerResponse
	fmt.Fprintf(os.Stdout, "Response from `TriggerAPI.CreateTrigger`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTriggerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dtoTrigger** | [**DtoTrigger**](DtoTrigger.md) | Trigger data | 
 **validate** | **bool** | For validating targets | 

### Return type

[**DtoSaveTriggerResponse**](DtoSaveTriggerResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePager

> DtoTriggersSearchResultDeleteResponse DeletePager(ctx).PagerID(pagerID).Execute()

Delete triggers pager

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
	pagerID := "pagerID_example" // string | Pager ID (optional) (default to "bcba82f5-48cf-44c0-b7d6-e1d32c64a88c")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TriggerAPI.DeletePager(context.Background()).PagerID(pagerID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TriggerAPI.DeletePager``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePager`: DtoTriggersSearchResultDeleteResponse
	fmt.Fprintf(os.Stdout, "Response from `TriggerAPI.DeletePager`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeletePagerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pagerID** | **string** | Pager ID | [default to &quot;bcba82f5-48cf-44c0-b7d6-e1d32c64a88c&quot;]

### Return type

[**DtoTriggersSearchResultDeleteResponse**](DtoTriggersSearchResultDeleteResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTriggerMetric

> map[string]interface{} DeleteTriggerMetric(ctx, triggerID).Name(name).Execute()

Delete metric from last check and all trigger pattern metrics

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
	triggerID := "triggerID_example" // string | Trigger ID (default to "bcba82f5-48cf-44c0-b7d6-e1d32c64a88c")
	name := "name_example" // string | Name of the target metric (optional) (default to "DevOps.my_server.hdd.freespace_mbytes")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TriggerAPI.DeleteTriggerMetric(context.Background(), triggerID).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TriggerAPI.DeleteTriggerMetric``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTriggerMetric`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `TriggerAPI.DeleteTriggerMetric`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**triggerID** | **string** | Trigger ID | [default to &quot;bcba82f5-48cf-44c0-b7d6-e1d32c64a88c&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTriggerMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | Name of the target metric | [default to &quot;DevOps.my_server.hdd.freespace_mbytes&quot;]

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


## DeleteTriggerNodataMetrics

> map[string]interface{} DeleteTriggerNodataMetrics(ctx, triggerID).Execute()

Delete all metrics from last data which are in NODATA state. It also deletes all trigger patterns of those metrics

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
	triggerID := "triggerID_example" // string | Trigger ID (default to "bcba82f5-48cf-44c0-b7d6-e1d32c64a88c")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TriggerAPI.DeleteTriggerNodataMetrics(context.Background(), triggerID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TriggerAPI.DeleteTriggerNodataMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTriggerNodataMetrics`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `TriggerAPI.DeleteTriggerNodataMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**triggerID** | **string** | Trigger ID | [default to &quot;bcba82f5-48cf-44c0-b7d6-e1d32c64a88c&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTriggerNodataMetricsRequest struct via the builder pattern


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


## DeleteTriggerThrottling

> DeleteTriggerThrottling(ctx, triggerID).Execute()

Deletes throttling for a trigger

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
	triggerID := "triggerID_example" // string | Trigger ID (default to "bcba82f5-48cf-44c0-b7d6-e1d32c64a88c")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TriggerAPI.DeleteTriggerThrottling(context.Background(), triggerID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TriggerAPI.DeleteTriggerThrottling``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**triggerID** | **string** | Trigger ID | [default to &quot;bcba82f5-48cf-44c0-b7d6-e1d32c64a88c&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTriggerThrottlingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllTriggers

> DtoTriggersList GetAllTriggers(ctx).Execute()

Get all triggers

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
	resp, r, err := apiClient.TriggerAPI.GetAllTriggers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TriggerAPI.GetAllTriggers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllTriggers`: DtoTriggersList
	fmt.Fprintf(os.Stdout, "Response from `TriggerAPI.GetAllTriggers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllTriggersRequest struct via the builder pattern


### Return type

[**DtoTriggersList**](DtoTriggersList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrigger

> DtoTrigger GetTrigger(ctx, triggerID).Populated(populated).Execute()

Get an existing trigger

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
	triggerID := "triggerID_example" // string | Trigger ID (default to "bcba82f5-48cf-44c0-b7d6-e1d32c64a88c")
	populated := true // bool | Populated (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TriggerAPI.GetTrigger(context.Background(), triggerID).Populated(populated).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TriggerAPI.GetTrigger``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTrigger`: DtoTrigger
	fmt.Fprintf(os.Stdout, "Response from `TriggerAPI.GetTrigger`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**triggerID** | **string** | Trigger ID | [default to &quot;bcba82f5-48cf-44c0-b7d6-e1d32c64a88c&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetTriggerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **populated** | **bool** | Populated | [default to false]

### Return type

[**DtoTrigger**](DtoTrigger.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTriggerDump

> DtoTriggerDump GetTriggerDump(ctx, triggerID).Execute()

Get trigger dump

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
	triggerID := "triggerID_example" // string | Trigger ID (default to "bcba82f5-48cf-44c0-b7d6-e1d32c64a88c")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TriggerAPI.GetTriggerDump(context.Background(), triggerID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TriggerAPI.GetTriggerDump``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTriggerDump`: DtoTriggerDump
	fmt.Fprintf(os.Stdout, "Response from `TriggerAPI.GetTriggerDump`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**triggerID** | **string** | Trigger ID | [default to &quot;bcba82f5-48cf-44c0-b7d6-e1d32c64a88c&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetTriggerDumpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DtoTriggerDump**](DtoTriggerDump.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTriggerMetrics

> map[string]map[string][]MoiraMetricValue GetTriggerMetrics(ctx, triggerID).From(from).To(to).Execute()

Get metrics associated with certain trigger

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
	triggerID := "triggerID_example" // string | Trigger ID (default to "bcba82f5-48cf-44c0-b7d6-e1d32c64a88c")
	from := "from_example" // string | Start time for metrics retrieval (optional) (default to "-10minutes")
	to := "to_example" // string | End time for metrics retrieval (optional) (default to "now")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TriggerAPI.GetTriggerMetrics(context.Background(), triggerID).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TriggerAPI.GetTriggerMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTriggerMetrics`: map[string]map[string][]MoiraMetricValue
	fmt.Fprintf(os.Stdout, "Response from `TriggerAPI.GetTriggerMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**triggerID** | **string** | Trigger ID | [default to &quot;bcba82f5-48cf-44c0-b7d6-e1d32c64a88c&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetTriggerMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **string** | Start time for metrics retrieval | [default to &quot;-10minutes&quot;]
 **to** | **string** | End time for metrics retrieval | [default to &quot;now&quot;]

### Return type

[**map[string]map[string][]MoiraMetricValue**](map.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTriggerState

> DtoTriggerCheck GetTriggerState(ctx, triggerID).Execute()

Get the trigger state as at last check

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
	triggerID := "triggerID_example" // string | Trigger ID (default to "bcba82f5-48cf-44c0-b7d6-e1d32c64a88c")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TriggerAPI.GetTriggerState(context.Background(), triggerID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TriggerAPI.GetTriggerState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTriggerState`: DtoTriggerCheck
	fmt.Fprintf(os.Stdout, "Response from `TriggerAPI.GetTriggerState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**triggerID** | **string** | Trigger ID | [default to &quot;bcba82f5-48cf-44c0-b7d6-e1d32c64a88c&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetTriggerStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DtoTriggerCheck**](DtoTriggerCheck.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTriggerThrottling

> DtoThrottlingResponse GetTriggerThrottling(ctx, triggerID).Execute()

Get a trigger with its throttling i.e its next allowed message time

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
	triggerID := "triggerID_example" // string | Trigger ID (default to "bcba82f5-48cf-44c0-b7d6-e1d32c64a88c")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TriggerAPI.GetTriggerThrottling(context.Background(), triggerID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TriggerAPI.GetTriggerThrottling``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTriggerThrottling`: DtoThrottlingResponse
	fmt.Fprintf(os.Stdout, "Response from `TriggerAPI.GetTriggerThrottling`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**triggerID** | **string** | Trigger ID | [default to &quot;bcba82f5-48cf-44c0-b7d6-e1d32c64a88c&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetTriggerThrottlingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DtoThrottlingResponse**](DtoThrottlingResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTriggersNoisiness

> DtoTriggerNoisinessList GetTriggersNoisiness(ctx).Size(size).P(p).From(from).To(to).Sort(sort).Execute()

Get triggers noisiness

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
	resp, r, err := apiClient.TriggerAPI.GetTriggersNoisiness(context.Background()).Size(size).P(p).From(from).To(to).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TriggerAPI.GetTriggersNoisiness``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTriggersNoisiness`: DtoTriggerNoisinessList
	fmt.Fprintf(os.Stdout, "Response from `TriggerAPI.GetTriggersNoisiness`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTriggersNoisinessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32** | Number of items to be displayed on one page. if size &#x3D; -1 then all events returned | [default to 100]
 **p** | **int32** | Defines the number of the displayed page. E.g, p&#x3D;2 would display the 2nd page | [default to 0]
 **from** | **string** | Start time of the time range | [default to &quot;-3hours&quot;]
 **to** | **string** | End time of the time range | [default to &quot;now&quot;]
 **sort** | **string** | String to set sort order (by events_count). On empty - no order, asc - ascending, desc - descending | [default to &quot;desc&quot;]

### Return type

[**DtoTriggerNoisinessList**](DtoTriggerNoisinessList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUnusedTriggers

> DtoTriggersList GetUnusedTriggers(ctx).Execute()

Get unused triggers

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
	resp, r, err := apiClient.TriggerAPI.GetUnusedTriggers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TriggerAPI.GetUnusedTriggers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUnusedTriggers`: DtoTriggersList
	fmt.Fprintf(os.Stdout, "Response from `TriggerAPI.GetUnusedTriggers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUnusedTriggersRequest struct via the builder pattern


### Return type

[**DtoTriggersList**](DtoTriggersList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveTrigger

> RemoveTrigger(ctx, triggerID).Execute()

Remove trigger

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
	triggerID := "triggerID_example" // string | Trigger ID (default to "bcba82f5-48cf-44c0-b7d6-e1d32c64a88c")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TriggerAPI.RemoveTrigger(context.Background(), triggerID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TriggerAPI.RemoveTrigger``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**triggerID** | **string** | Trigger ID | [default to &quot;bcba82f5-48cf-44c0-b7d6-e1d32c64a88c&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveTriggerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RenderTriggerMetrics

> *os.File RenderTriggerMetrics(ctx, triggerID).Target(target).From(from).To(to).Timezone(timezone).Theme(theme).Realtime(realtime).Execute()

Render trigger metrics plot

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
	triggerID := "triggerID_example" // string | Trigger ID (default to "bcba82f5-48cf-44c0-b7d6-e1d32c64a88c")
	target := "target_example" // string | Target metric name (optional) (default to "t1")
	from := "from_example" // string | Start time for metrics retrieval (optional) (default to "-1hour")
	to := "to_example" // string | End time for metrics retrieval (optional) (default to "now")
	timezone := "timezone_example" // string | Timezone for rendering (optional) (default to "UTC")
	theme := "theme_example" // string | Plot theme (optional) (default to "light")
	realtime := true // bool | Fetch real-time data (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TriggerAPI.RenderTriggerMetrics(context.Background(), triggerID).Target(target).From(from).To(to).Timezone(timezone).Theme(theme).Realtime(realtime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TriggerAPI.RenderTriggerMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RenderTriggerMetrics`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `TriggerAPI.RenderTriggerMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**triggerID** | **string** | Trigger ID | [default to &quot;bcba82f5-48cf-44c0-b7d6-e1d32c64a88c&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiRenderTriggerMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **target** | **string** | Target metric name | [default to &quot;t1&quot;]
 **from** | **string** | Start time for metrics retrieval | [default to &quot;-1hour&quot;]
 **to** | **string** | End time for metrics retrieval | [default to &quot;now&quot;]
 **timezone** | **string** | Timezone for rendering | [default to &quot;UTC&quot;]
 **theme** | **string** | Plot theme | [default to &quot;light&quot;]
 **realtime** | **bool** | Fetch real-time data | [default to false]

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/png, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchTriggers

> DtoTriggersList SearchTriggers(ctx).OnlyProblems(onlyProblems).Text(text).P(p).Size(size).Tags(tags).CreatePager(createPager).PagerID(pagerID).CreatedBy(createdBy).Execute()

Search triggers. Replaces the deprecated `page` path



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
	onlyProblems := true // bool | Only include problems (optional) (default to false)
	text := "text_example" // string | Search text (optional) (default to "cpu")
	p := int32(56) // int32 | Page number (optional) (default to 0)
	size := int32(56) // int32 | Page size (optional) (default to 10)
	tags := []string{[]string{"Tags_example"}} // []string | Search tag (optional)
	createPager := true // bool | Create pager (optional) (default to false)
	pagerID := "pagerID_example" // string | Pager ID (optional) (default to "bcba82f5-48cf-44c0-b7d6-e1d32c64a88c")
	createdBy := "createdBy_example" // string | Created By (optional) (default to "moira.team")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TriggerAPI.SearchTriggers(context.Background()).OnlyProblems(onlyProblems).Text(text).P(p).Size(size).Tags(tags).CreatePager(createPager).PagerID(pagerID).CreatedBy(createdBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TriggerAPI.SearchTriggers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchTriggers`: DtoTriggersList
	fmt.Fprintf(os.Stdout, "Response from `TriggerAPI.SearchTriggers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchTriggersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **onlyProblems** | **bool** | Only include problems | [default to false]
 **text** | **string** | Search text | [default to &quot;cpu&quot;]
 **p** | **int32** | Page number | [default to 0]
 **size** | **int32** | Page size | [default to 10]
 **tags** | **[][]string** | Search tag | 
 **createPager** | **bool** | Create pager | [default to false]
 **pagerID** | **string** | Pager ID | [default to &quot;bcba82f5-48cf-44c0-b7d6-e1d32c64a88c&quot;]
 **createdBy** | **string** | Created By | [default to &quot;moira.team&quot;]

### Return type

[**DtoTriggersList**](DtoTriggersList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetTriggerMaintenance

> map[string]interface{} SetTriggerMaintenance(ctx, triggerID).DtoTriggerMaintenance(dtoTriggerMaintenance).Execute()

Set metrics and the trigger itself to maintenance mode

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
	triggerID := "triggerID_example" // string | Trigger ID (default to "bcba82f5-48cf-44c0-b7d6-e1d32c64a88c")
	dtoTriggerMaintenance := *openapiclient.NewDtoTriggerMaintenance() // DtoTriggerMaintenance | Maintenance data

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TriggerAPI.SetTriggerMaintenance(context.Background(), triggerID).DtoTriggerMaintenance(dtoTriggerMaintenance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TriggerAPI.SetTriggerMaintenance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetTriggerMaintenance`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `TriggerAPI.SetTriggerMaintenance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**triggerID** | **string** | Trigger ID | [default to &quot;bcba82f5-48cf-44c0-b7d6-e1d32c64a88c&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiSetTriggerMaintenanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dtoTriggerMaintenance** | [**DtoTriggerMaintenance**](DtoTriggerMaintenance.md) | Maintenance data | 

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


## TriggerCheck

> DtoTriggerCheckResponse TriggerCheck(ctx).DtoTrigger(dtoTrigger).Execute()

Validates trigger target

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
	dtoTrigger := *openapiclient.NewDtoTrigger() // DtoTrigger | Trigger data

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TriggerAPI.TriggerCheck(context.Background()).DtoTrigger(dtoTrigger).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TriggerAPI.TriggerCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TriggerCheck`: DtoTriggerCheckResponse
	fmt.Fprintf(os.Stdout, "Response from `TriggerAPI.TriggerCheck`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTriggerCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dtoTrigger** | [**DtoTrigger**](DtoTrigger.md) | Trigger data | 

### Return type

[**DtoTriggerCheckResponse**](DtoTriggerCheckResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTrigger

> DtoSaveTriggerResponse UpdateTrigger(ctx, triggerID).DtoTrigger(dtoTrigger).Validate(validate).Execute()

Update existing trigger

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
	triggerID := "triggerID_example" // string | Trigger ID (default to "bcba82f5-48cf-44c0-b7d6-e1d32c64a88c")
	dtoTrigger := *openapiclient.NewDtoTrigger() // DtoTrigger | Trigger data
	validate := true // bool | For validating targets (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TriggerAPI.UpdateTrigger(context.Background(), triggerID).DtoTrigger(dtoTrigger).Validate(validate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TriggerAPI.UpdateTrigger``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTrigger`: DtoSaveTriggerResponse
	fmt.Fprintf(os.Stdout, "Response from `TriggerAPI.UpdateTrigger`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**triggerID** | **string** | Trigger ID | [default to &quot;bcba82f5-48cf-44c0-b7d6-e1d32c64a88c&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTriggerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dtoTrigger** | [**DtoTrigger**](DtoTrigger.md) | Trigger data | 
 **validate** | **bool** | For validating targets | 

### Return type

[**DtoSaveTriggerResponse**](DtoSaveTriggerResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

