# DtoEventsList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**List** | Pointer to [**[]MoiraNotificationEvent**](MoiraNotificationEvent.md) |  | [optional] 
**Page** | Pointer to **int64** |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**Total** | Pointer to **int64** |  | [optional] 

## Methods

### NewDtoEventsList

`func NewDtoEventsList() *DtoEventsList`

NewDtoEventsList instantiates a new DtoEventsList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoEventsListWithDefaults

`func NewDtoEventsListWithDefaults() *DtoEventsList`

NewDtoEventsListWithDefaults instantiates a new DtoEventsList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetList

`func (o *DtoEventsList) GetList() []MoiraNotificationEvent`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *DtoEventsList) GetListOk() (*[]MoiraNotificationEvent, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *DtoEventsList) SetList(v []MoiraNotificationEvent)`

SetList sets List field to given value.

### HasList

`func (o *DtoEventsList) HasList() bool`

HasList returns a boolean if a field has been set.

### GetPage

`func (o *DtoEventsList) GetPage() int64`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *DtoEventsList) GetPageOk() (*int64, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *DtoEventsList) SetPage(v int64)`

SetPage sets Page field to given value.

### HasPage

`func (o *DtoEventsList) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetSize

`func (o *DtoEventsList) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *DtoEventsList) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *DtoEventsList) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *DtoEventsList) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetTotal

`func (o *DtoEventsList) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *DtoEventsList) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *DtoEventsList) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *DtoEventsList) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


