# internal\TeamSubscriptionAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNewTeamSubscription**](TeamSubscriptionAPI.md#CreateNewTeamSubscription) | **Post** /teams/{teamID}/subscriptions | Create a new team subscription



## CreateNewTeamSubscription

> DtoSubscription CreateNewTeamSubscription(ctx, teamID).DtoSubscription(dtoSubscription).Execute()

Create a new team subscription

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
	teamID := "teamID_example" // string | The ID of team (default to "bcba82f5-48cf-44c0-b7d6-e1d32c64a88c")
	dtoSubscription := *openapiclient.NewDtoSubscription() // DtoSubscription | Team subscription data

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamSubscriptionAPI.CreateNewTeamSubscription(context.Background(), teamID).DtoSubscription(dtoSubscription).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamSubscriptionAPI.CreateNewTeamSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNewTeamSubscription`: DtoSubscription
	fmt.Fprintf(os.Stdout, "Response from `TeamSubscriptionAPI.CreateNewTeamSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamID** | **string** | The ID of team | [default to &quot;bcba82f5-48cf-44c0-b7d6-e1d32c64a88c&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNewTeamSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dtoSubscription** | [**DtoSubscription**](DtoSubscription.md) | Team subscription data | 

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

