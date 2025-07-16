# internal\TeamAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTeamUsers**](TeamAPI.md#AddTeamUsers) | **Post** /teams/{teamID}/users | Add users to a team
[**CreateTeam**](TeamAPI.md#CreateTeam) | **Post** /teams | Create a new team
[**DeleteTeam**](TeamAPI.md#DeleteTeam) | **Delete** /teams/{teamID} | Delete a team
[**DeleteTeamUser**](TeamAPI.md#DeleteTeamUser) | **Delete** /teams/{teamID}/users/{teamUserID} | Delete a user from a team
[**GetAllTeams**](TeamAPI.md#GetAllTeams) | **Get** /teams/all | Get all Moira teams
[**GetAllTeamsForUser**](TeamAPI.md#GetAllTeamsForUser) | **Get** /teams | Get all teams for user
[**GetTeam**](TeamAPI.md#GetTeam) | **Get** /teams/{teamID} | Get a team by ID
[**GetTeamSettings**](TeamAPI.md#GetTeamSettings) | **Get** /teams/{teamID}/settings | Get team settings
[**GetTeamUsers**](TeamAPI.md#GetTeamUsers) | **Get** /teams/{teamID}/users | Get users of a team
[**SetTeamUsers**](TeamAPI.md#SetTeamUsers) | **Put** /teams/{teamID}/users | Set users of a team
[**UpdateTeam**](TeamAPI.md#UpdateTeam) | **Patch** /teams/{teamID} | Update existing team



## AddTeamUsers

> DtoTeamMembers AddTeamUsers(ctx, teamID).DtoTeamMembers(dtoTeamMembers).Execute()

Add users to a team

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
	teamID := "teamID_example" // string | ID of the team (default to "bcba82f5-48cf-44c0-b7d6-e1d32c64a88c")
	dtoTeamMembers := *openapiclient.NewDtoTeamMembers() // DtoTeamMembers | Usernames to add to the team

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamAPI.AddTeamUsers(context.Background(), teamID).DtoTeamMembers(dtoTeamMembers).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.AddTeamUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddTeamUsers`: DtoTeamMembers
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.AddTeamUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamID** | **string** | ID of the team | [default to &quot;bcba82f5-48cf-44c0-b7d6-e1d32c64a88c&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiAddTeamUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dtoTeamMembers** | [**DtoTeamMembers**](DtoTeamMembers.md) | Usernames to add to the team | 

### Return type

[**DtoTeamMembers**](DtoTeamMembers.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTeam

> DtoSaveTeamResponse CreateTeam(ctx).DtoTeamModel(dtoTeamModel).Execute()

Create a new team

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
	dtoTeamModel := *openapiclient.NewDtoTeamModel() // DtoTeamModel | Team data

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamAPI.CreateTeam(context.Background()).DtoTeamModel(dtoTeamModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.CreateTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTeam`: DtoSaveTeamResponse
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.CreateTeam`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dtoTeamModel** | [**DtoTeamModel**](DtoTeamModel.md) | Team data | 

### Return type

[**DtoSaveTeamResponse**](DtoSaveTeamResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTeam

> DtoSaveTeamResponse DeleteTeam(ctx, teamID).Execute()

Delete a team

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
	teamID := "teamID_example" // string | ID of the team (default to "bcba82f5-48cf-44c0-b7d6-e1d32c64a88c")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamAPI.DeleteTeam(context.Background(), teamID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.DeleteTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTeam`: DtoSaveTeamResponse
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.DeleteTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamID** | **string** | ID of the team | [default to &quot;bcba82f5-48cf-44c0-b7d6-e1d32c64a88c&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DtoSaveTeamResponse**](DtoSaveTeamResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTeamUser

> DtoTeamMembers DeleteTeamUser(ctx, teamID, teamUserID).Execute()

Delete a user from a team

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
	teamID := "teamID_example" // string | ID of the team (default to "bcba82f5-48cf-44c0-b7d6-e1d32c64a88c")
	teamUserID := "teamUserID_example" // string | User login in methods related to teams manipulation (default to "anonymous")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamAPI.DeleteTeamUser(context.Background(), teamID, teamUserID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.DeleteTeamUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTeamUser`: DtoTeamMembers
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.DeleteTeamUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamID** | **string** | ID of the team | [default to &quot;bcba82f5-48cf-44c0-b7d6-e1d32c64a88c&quot;]
**teamUserID** | **string** | User login in methods related to teams manipulation | [default to &quot;anonymous&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTeamUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DtoTeamMembers**](DtoTeamMembers.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllTeams

> DtoTeamsList GetAllTeams(ctx).Size(size).P(p).SearchText(searchText).Sort(sort).Execute()

Get all Moira teams

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
	size := int32(56) // int32 | Number of items to be displayed on one page. if size = -1 then all teams returned (optional) (default to -1)
	p := int32(56) // int32 | Defines the number of the displayed page. E.g, p=2 would display the 2nd page (optional) (default to 0)
	searchText := "searchText_example" // string | Regular expression which will be applied to team id or team name than filtering teams (optional) (default to ".*")
	sort := "sort_example" // string | String to set sort order (by name). On empty - no order, asc - ascending, desc - descending (optional) (default to "asc")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamAPI.GetAllTeams(context.Background()).Size(size).P(p).SearchText(searchText).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.GetAllTeams``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllTeams`: DtoTeamsList
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.GetAllTeams`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllTeamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32** | Number of items to be displayed on one page. if size &#x3D; -1 then all teams returned | [default to -1]
 **p** | **int32** | Defines the number of the displayed page. E.g, p&#x3D;2 would display the 2nd page | [default to 0]
 **searchText** | **string** | Regular expression which will be applied to team id or team name than filtering teams | [default to &quot;.*&quot;]
 **sort** | **string** | String to set sort order (by name). On empty - no order, asc - ascending, desc - descending | [default to &quot;asc&quot;]

### Return type

[**DtoTeamsList**](DtoTeamsList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllTeamsForUser

> DtoUserTeams GetAllTeamsForUser(ctx).Execute()

Get all teams for user

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
	resp, r, err := apiClient.TeamAPI.GetAllTeamsForUser(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.GetAllTeamsForUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllTeamsForUser`: DtoUserTeams
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.GetAllTeamsForUser`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllTeamsForUserRequest struct via the builder pattern


### Return type

[**DtoUserTeams**](DtoUserTeams.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeam

> DtoTeamModel GetTeam(ctx, teamID).Execute()

Get a team by ID

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
	teamID := "teamID_example" // string | ID of the team (default to "bcba82f5-48cf-44c0-b7d6-e1d32c64a88c")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamAPI.GetTeam(context.Background(), teamID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.GetTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTeam`: DtoTeamModel
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.GetTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamID** | **string** | ID of the team | [default to &quot;bcba82f5-48cf-44c0-b7d6-e1d32c64a88c&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DtoTeamModel**](DtoTeamModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeamSettings

> DtoTeamSettings GetTeamSettings(ctx, teamID).Execute()

Get team settings

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
	teamID := "teamID_example" // string | ID of the team (default to "bcba82f5-48cf-44c0-b7d6-e1d32c64a88c")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamAPI.GetTeamSettings(context.Background(), teamID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.GetTeamSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTeamSettings`: DtoTeamSettings
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.GetTeamSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamID** | **string** | ID of the team | [default to &quot;bcba82f5-48cf-44c0-b7d6-e1d32c64a88c&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DtoTeamSettings**](DtoTeamSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeamUsers

> DtoTeamMembers GetTeamUsers(ctx, teamID).Execute()

Get users of a team

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
	teamID := "teamID_example" // string | ID of the team (default to "bcba82f5-48cf-44c0-b7d6-e1d32c64a88c")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamAPI.GetTeamUsers(context.Background(), teamID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.GetTeamUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTeamUsers`: DtoTeamMembers
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.GetTeamUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamID** | **string** | ID of the team | [default to &quot;bcba82f5-48cf-44c0-b7d6-e1d32c64a88c&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DtoTeamMembers**](DtoTeamMembers.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetTeamUsers

> DtoTeamMembers SetTeamUsers(ctx, teamID).DtoTeamMembers(dtoTeamMembers).Execute()

Set users of a team

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
	teamID := "teamID_example" // string | ID of the team (default to "bcba82f5-48cf-44c0-b7d6-e1d32c64a88c")
	dtoTeamMembers := *openapiclient.NewDtoTeamMembers() // DtoTeamMembers | Usernames to set as team members

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamAPI.SetTeamUsers(context.Background(), teamID).DtoTeamMembers(dtoTeamMembers).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.SetTeamUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetTeamUsers`: DtoTeamMembers
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.SetTeamUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamID** | **string** | ID of the team | [default to &quot;bcba82f5-48cf-44c0-b7d6-e1d32c64a88c&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiSetTeamUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dtoTeamMembers** | [**DtoTeamMembers**](DtoTeamMembers.md) | Usernames to set as team members | 

### Return type

[**DtoTeamMembers**](DtoTeamMembers.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTeam

> DtoSaveTeamResponse UpdateTeam(ctx, teamID).DtoTeamModel(dtoTeamModel).Execute()

Update existing team

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
	teamID := "teamID_example" // string | ID of the team (default to "bcba82f5-48cf-44c0-b7d6-e1d32c64a88c")
	dtoTeamModel := *openapiclient.NewDtoTeamModel() // DtoTeamModel | Updated team data

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamAPI.UpdateTeam(context.Background(), teamID).DtoTeamModel(dtoTeamModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.UpdateTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTeam`: DtoSaveTeamResponse
	fmt.Fprintf(os.Stdout, "Response from `TeamAPI.UpdateTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamID** | **string** | ID of the team | [default to &quot;bcba82f5-48cf-44c0-b7d6-e1d32c64a88c&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dtoTeamModel** | [**DtoTeamModel**](DtoTeamModel.md) | Updated team data | 

### Return type

[**DtoSaveTeamResponse**](DtoSaveTeamResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

