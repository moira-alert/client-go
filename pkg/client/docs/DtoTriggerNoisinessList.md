# DtoTriggerNoisinessList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**List** | Pointer to [**[]DtoTriggerNoisiness**](DtoTriggerNoisiness.md) | List of entities. | [optional] 
**Page** | Pointer to **int64** | Page number. | [optional] 
**Size** | Pointer to **int64** | Size is the amount of entities per Page. | [optional] 
**Total** | Pointer to **int64** | Total amount of entities in the database. | [optional] 

## Methods

### NewDtoTriggerNoisinessList

`func NewDtoTriggerNoisinessList() *DtoTriggerNoisinessList`

NewDtoTriggerNoisinessList instantiates a new DtoTriggerNoisinessList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoTriggerNoisinessListWithDefaults

`func NewDtoTriggerNoisinessListWithDefaults() *DtoTriggerNoisinessList`

NewDtoTriggerNoisinessListWithDefaults instantiates a new DtoTriggerNoisinessList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetList

`func (o *DtoTriggerNoisinessList) GetList() []DtoTriggerNoisiness`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *DtoTriggerNoisinessList) GetListOk() (*[]DtoTriggerNoisiness, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *DtoTriggerNoisinessList) SetList(v []DtoTriggerNoisiness)`

SetList sets List field to given value.

### HasList

`func (o *DtoTriggerNoisinessList) HasList() bool`

HasList returns a boolean if a field has been set.

### GetPage

`func (o *DtoTriggerNoisinessList) GetPage() int64`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *DtoTriggerNoisinessList) GetPageOk() (*int64, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *DtoTriggerNoisinessList) SetPage(v int64)`

SetPage sets Page field to given value.

### HasPage

`func (o *DtoTriggerNoisinessList) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetSize

`func (o *DtoTriggerNoisinessList) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *DtoTriggerNoisinessList) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *DtoTriggerNoisinessList) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *DtoTriggerNoisinessList) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetTotal

`func (o *DtoTriggerNoisinessList) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *DtoTriggerNoisinessList) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *DtoTriggerNoisinessList) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *DtoTriggerNoisinessList) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


