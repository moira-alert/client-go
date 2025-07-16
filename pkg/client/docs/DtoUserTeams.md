# DtoUserTeams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Teams** | Pointer to [**[]DtoTeamModel**](DtoTeamModel.md) |  | [optional] 

## Methods

### NewDtoUserTeams

`func NewDtoUserTeams() *DtoUserTeams`

NewDtoUserTeams instantiates a new DtoUserTeams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoUserTeamsWithDefaults

`func NewDtoUserTeamsWithDefaults() *DtoUserTeams`

NewDtoUserTeamsWithDefaults instantiates a new DtoUserTeams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTeams

`func (o *DtoUserTeams) GetTeams() []DtoTeamModel`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *DtoUserTeams) GetTeamsOk() (*[]DtoTeamModel, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *DtoUserTeams) SetTeams(v []DtoTeamModel)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *DtoUserTeams) HasTeams() bool`

HasTeams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


