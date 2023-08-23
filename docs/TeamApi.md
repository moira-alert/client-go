# {{classname}}

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTeamUsers**](TeamApi.md#AddTeamUsers) | **Post** /teams/{teamID}/users | Add users to a team
[**CreateTeam**](TeamApi.md#CreateTeam) | **Post** /teams | Create a new team
[**DeleteTeam**](TeamApi.md#DeleteTeam) | **Delete** /teams/{teamID} | Delete a team
[**DeleteTeamUser**](TeamApi.md#DeleteTeamUser) | **Delete** /teams/{teamID}/users/{teamUserID} | Delete a user from a team
[**GetAllTeams**](TeamApi.md#GetAllTeams) | **Get** /teams | Get all teams
[**GetTeam**](TeamApi.md#GetTeam) | **Get** /teams/{teamID} | Get a team by ID
[**GetTeamSettings**](TeamApi.md#GetTeamSettings) | **Get** /teams/{teamID}/settings | Get team settings
[**GetTeamUsers**](TeamApi.md#GetTeamUsers) | **Get** /teams/{teamID}/users | Get users of a team
[**SetTeamUsers**](TeamApi.md#SetTeamUsers) | **Put** /teams/{teamID}/users | Set users of a team
[**UpdateTeam**](TeamApi.md#UpdateTeam) | **Patch** /teams/{teamID} | Update existing team

# **AddTeamUsers**
> DtoTeamMembers AddTeamUsers(ctx, body, teamID)
Add users to a team

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DtoTeamMembers**](DtoTeamMembers.md)| Usernames to add to the team | 
  **teamID** | **string**| ID of the team | 

### Return type

[**DtoTeamMembers**](dto.TeamMembers.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateTeam**
> DtoSaveTeamResponse CreateTeam(ctx, body)
Create a new team

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DtoTeamModel**](DtoTeamModel.md)| Team data | 

### Return type

[**DtoSaveTeamResponse**](dto.SaveTeamResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTeam**
> DtoSaveTeamResponse DeleteTeam(ctx, teamID)
Delete a team

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **teamID** | **string**| ID of the team | 

### Return type

[**DtoSaveTeamResponse**](dto.SaveTeamResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTeamUser**
> DtoTeamMembers DeleteTeamUser(ctx, teamID, teamUserID)
Delete a user from a team

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **teamID** | **string**| ID of the team | 
  **teamUserID** | **string**| User login in methods related to teams manipulation | 

### Return type

[**DtoTeamMembers**](dto.TeamMembers.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllTeams**
> DtoUserTeams GetAllTeams(ctx, )
Get all teams

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**DtoUserTeams**](dto.UserTeams.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTeam**
> DtoTeamModel GetTeam(ctx, teamID)
Get a team by ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **teamID** | **string**| ID of the team | 

### Return type

[**DtoTeamModel**](dto.TeamModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTeamSettings**
> DtoTeamSettings GetTeamSettings(ctx, teamID)
Get team settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **teamID** | **string**| ID of the team | 

### Return type

[**DtoTeamSettings**](dto.TeamSettings.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTeamUsers**
> DtoTeamMembers GetTeamUsers(ctx, teamID)
Get users of a team

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **teamID** | **string**| ID of the team | 

### Return type

[**DtoTeamMembers**](dto.TeamMembers.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetTeamUsers**
> DtoTeamMembers SetTeamUsers(ctx, body, teamID)
Set users of a team

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DtoTeamMembers**](DtoTeamMembers.md)| Usernames to set as team members | 
  **teamID** | **string**| ID of the team | 

### Return type

[**DtoTeamMembers**](dto.TeamMembers.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTeam**
> DtoSaveTeamResponse UpdateTeam(ctx, body, teamID)
Update existing team

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DtoTeamModel**](DtoTeamModel.md)| Updated team data | 
  **teamID** | **string**| ID of the team | 

### Return type

[**DtoSaveTeamResponse**](dto.SaveTeamResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

