# MoiraScheduleData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Days** | Pointer to [**[]MoiraScheduleDataDay**](MoiraScheduleDataDay.md) |  | [optional] 
**EndOffset** | Pointer to **int64** |  | [optional] 
**StartOffset** | Pointer to **int64** |  | [optional] 
**TzOffset** | Pointer to **int64** |  | [optional] 

## Methods

### NewMoiraScheduleData

`func NewMoiraScheduleData() *MoiraScheduleData`

NewMoiraScheduleData instantiates a new MoiraScheduleData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoiraScheduleDataWithDefaults

`func NewMoiraScheduleDataWithDefaults() *MoiraScheduleData`

NewMoiraScheduleDataWithDefaults instantiates a new MoiraScheduleData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDays

`func (o *MoiraScheduleData) GetDays() []MoiraScheduleDataDay`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *MoiraScheduleData) GetDaysOk() (*[]MoiraScheduleDataDay, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *MoiraScheduleData) SetDays(v []MoiraScheduleDataDay)`

SetDays sets Days field to given value.

### HasDays

`func (o *MoiraScheduleData) HasDays() bool`

HasDays returns a boolean if a field has been set.

### GetEndOffset

`func (o *MoiraScheduleData) GetEndOffset() int64`

GetEndOffset returns the EndOffset field if non-nil, zero value otherwise.

### GetEndOffsetOk

`func (o *MoiraScheduleData) GetEndOffsetOk() (*int64, bool)`

GetEndOffsetOk returns a tuple with the EndOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndOffset

`func (o *MoiraScheduleData) SetEndOffset(v int64)`

SetEndOffset sets EndOffset field to given value.

### HasEndOffset

`func (o *MoiraScheduleData) HasEndOffset() bool`

HasEndOffset returns a boolean if a field has been set.

### GetStartOffset

`func (o *MoiraScheduleData) GetStartOffset() int64`

GetStartOffset returns the StartOffset field if non-nil, zero value otherwise.

### GetStartOffsetOk

`func (o *MoiraScheduleData) GetStartOffsetOk() (*int64, bool)`

GetStartOffsetOk returns a tuple with the StartOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartOffset

`func (o *MoiraScheduleData) SetStartOffset(v int64)`

SetStartOffset sets StartOffset field to given value.

### HasStartOffset

`func (o *MoiraScheduleData) HasStartOffset() bool`

HasStartOffset returns a boolean if a field has been set.

### GetTzOffset

`func (o *MoiraScheduleData) GetTzOffset() int64`

GetTzOffset returns the TzOffset field if non-nil, zero value otherwise.

### GetTzOffsetOk

`func (o *MoiraScheduleData) GetTzOffsetOk() (*int64, bool)`

GetTzOffsetOk returns a tuple with the TzOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTzOffset

`func (o *MoiraScheduleData) SetTzOffset(v int64)`

SetTzOffset sets TzOffset field to given value.

### HasTzOffset

`func (o *MoiraScheduleData) HasTzOffset() bool`

HasTzOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


