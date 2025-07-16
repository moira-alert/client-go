# ApiSentry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dsn** | Pointer to **string** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 

## Methods

### NewApiSentry

`func NewApiSentry() *ApiSentry`

NewApiSentry instantiates a new ApiSentry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiSentryWithDefaults

`func NewApiSentryWithDefaults() *ApiSentry`

NewApiSentryWithDefaults instantiates a new ApiSentry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDsn

`func (o *ApiSentry) GetDsn() string`

GetDsn returns the Dsn field if non-nil, zero value otherwise.

### GetDsnOk

`func (o *ApiSentry) GetDsnOk() (*string, bool)`

GetDsnOk returns a tuple with the Dsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsn

`func (o *ApiSentry) SetDsn(v string)`

SetDsn sets Dsn field to given value.

### HasDsn

`func (o *ApiSentry) HasDsn() bool`

HasDsn returns a boolean if a field has been set.

### GetPlatform

`func (o *ApiSentry) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ApiSentry) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ApiSentry) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *ApiSentry) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


