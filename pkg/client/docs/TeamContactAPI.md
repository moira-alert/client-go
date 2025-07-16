# internal\TeamContactAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNewTeamContact**](TeamContactAPI.md#CreateNewTeamContact) | **Post** /teams/{teamID}/contacts | Create a new team contact



## CreateNewTeamContact

> DtoContact CreateNewTeamContact(ctx, teamID).DtoContact(dtoContact).Execute()

Create a new team contact

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
	dtoContact := *openapiclient.NewDtoContact() // DtoContact | Team contact data

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamContactAPI.CreateNewTeamContact(context.Background(), teamID).DtoContact(dtoContact).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamContactAPI.CreateNewTeamContact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNewTeamContact`: DtoContact
	fmt.Fprintf(os.Stdout, "Response from `TeamContactAPI.CreateNewTeamContact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamID** | **string** | The ID of team | [default to &quot;bcba82f5-48cf-44c0-b7d6-e1d32c64a88c&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNewTeamContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dtoContact** | [**DtoContact**](DtoContact.md) | Team contact data | 

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

