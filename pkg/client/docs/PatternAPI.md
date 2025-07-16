# internal\PatternAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePattern**](PatternAPI.md#DeletePattern) | **Delete** /pattern/{pattern} | Deletes a Moira pattern
[**GetAllPatterns**](PatternAPI.md#GetAllPatterns) | **Get** /pattern | Get all patterns



## DeletePattern

> map[string]interface{} DeletePattern(ctx, pattern).Execute()

Deletes a Moira pattern

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
	pattern := "pattern_example" // string | Trigger pattern to operate on (default to "DevOps.my_server.hdd.freespace_mbytes")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatternAPI.DeletePattern(context.Background(), pattern).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatternAPI.DeletePattern``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePattern`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PatternAPI.DeletePattern`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pattern** | **string** | Trigger pattern to operate on | [default to &quot;DevOps.my_server.hdd.freespace_mbytes&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePatternRequest struct via the builder pattern


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


## GetAllPatterns

> DtoPatternList GetAllPatterns(ctx).Execute()

Get all patterns

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
	resp, r, err := apiClient.PatternAPI.GetAllPatterns(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatternAPI.GetAllPatterns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllPatterns`: DtoPatternList
	fmt.Fprintf(os.Stdout, "Response from `PatternAPI.GetAllPatterns`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllPatternsRequest struct via the builder pattern


### Return type

[**DtoPatternList**](DtoPatternList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

