# DtoContactNoisinessList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**List** | Pointer to [**[]DtoContactNoisiness**](DtoContactNoisiness.md) | List of entities. | [optional] 
**Page** | Pointer to **int64** | Page number. | [optional] 
**Size** | Pointer to **int64** | Size is the amount of entities per Page. | [optional] 
**Total** | Pointer to **int64** | Total amount of entities in the database. | [optional] 

## Methods

### NewDtoContactNoisinessList

`func NewDtoContactNoisinessList() *DtoContactNoisinessList`

NewDtoContactNoisinessList instantiates a new DtoContactNoisinessList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoContactNoisinessListWithDefaults

`func NewDtoContactNoisinessListWithDefaults() *DtoContactNoisinessList`

NewDtoContactNoisinessListWithDefaults instantiates a new DtoContactNoisinessList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetList

`func (o *DtoContactNoisinessList) GetList() []DtoContactNoisiness`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *DtoContactNoisinessList) GetListOk() (*[]DtoContactNoisiness, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *DtoContactNoisinessList) SetList(v []DtoContactNoisiness)`

SetList sets List field to given value.

### HasList

`func (o *DtoContactNoisinessList) HasList() bool`

HasList returns a boolean if a field has been set.

### GetPage

`func (o *DtoContactNoisinessList) GetPage() int64`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *DtoContactNoisinessList) GetPageOk() (*int64, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *DtoContactNoisinessList) SetPage(v int64)`

SetPage sets Page field to given value.

### HasPage

`func (o *DtoContactNoisinessList) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetSize

`func (o *DtoContactNoisinessList) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *DtoContactNoisinessList) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *DtoContactNoisinessList) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *DtoContactNoisinessList) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetTotal

`func (o *DtoContactNoisinessList) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *DtoContactNoisinessList) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *DtoContactNoisinessList) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *DtoContactNoisinessList) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


