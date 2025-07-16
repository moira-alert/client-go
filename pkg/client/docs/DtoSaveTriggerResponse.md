# DtoSaveTriggerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckResult** | Pointer to [**DtoTriggerCheckResponse**](DtoTriggerCheckResponse.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewDtoSaveTriggerResponse

`func NewDtoSaveTriggerResponse() *DtoSaveTriggerResponse`

NewDtoSaveTriggerResponse instantiates a new DtoSaveTriggerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoSaveTriggerResponseWithDefaults

`func NewDtoSaveTriggerResponseWithDefaults() *DtoSaveTriggerResponse`

NewDtoSaveTriggerResponseWithDefaults instantiates a new DtoSaveTriggerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckResult

`func (o *DtoSaveTriggerResponse) GetCheckResult() DtoTriggerCheckResponse`

GetCheckResult returns the CheckResult field if non-nil, zero value otherwise.

### GetCheckResultOk

`func (o *DtoSaveTriggerResponse) GetCheckResultOk() (*DtoTriggerCheckResponse, bool)`

GetCheckResultOk returns a tuple with the CheckResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckResult

`func (o *DtoSaveTriggerResponse) SetCheckResult(v DtoTriggerCheckResponse)`

SetCheckResult sets CheckResult field to given value.

### HasCheckResult

`func (o *DtoSaveTriggerResponse) HasCheckResult() bool`

HasCheckResult returns a boolean if a field has been set.

### GetId

`func (o *DtoSaveTriggerResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DtoSaveTriggerResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DtoSaveTriggerResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DtoSaveTriggerResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessage

`func (o *DtoSaveTriggerResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DtoSaveTriggerResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DtoSaveTriggerResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *DtoSaveTriggerResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


