# DtoNotificationsList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**List** | Pointer to [**[]MoiraScheduledNotification**](MoiraScheduledNotification.md) |  | [optional] 
**Total** | Pointer to **int64** |  | [optional] 

## Methods

### NewDtoNotificationsList

`func NewDtoNotificationsList() *DtoNotificationsList`

NewDtoNotificationsList instantiates a new DtoNotificationsList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoNotificationsListWithDefaults

`func NewDtoNotificationsListWithDefaults() *DtoNotificationsList`

NewDtoNotificationsListWithDefaults instantiates a new DtoNotificationsList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetList

`func (o *DtoNotificationsList) GetList() []MoiraScheduledNotification`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *DtoNotificationsList) GetListOk() (*[]MoiraScheduledNotification, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *DtoNotificationsList) SetList(v []MoiraScheduledNotification)`

SetList sets List field to given value.

### HasList

`func (o *DtoNotificationsList) HasList() bool`

HasList returns a boolean if a field has been set.

### GetTotal

`func (o *DtoNotificationsList) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *DtoNotificationsList) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *DtoNotificationsList) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *DtoNotificationsList) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


