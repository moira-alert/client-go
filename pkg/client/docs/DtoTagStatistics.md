# DtoTagStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Subscriptions** | Pointer to [**[]MoiraSubscriptionData**](MoiraSubscriptionData.md) |  | [optional] 
**Triggers** | Pointer to **[]string** |  | [optional] 

## Methods

### NewDtoTagStatistics

`func NewDtoTagStatistics() *DtoTagStatistics`

NewDtoTagStatistics instantiates a new DtoTagStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoTagStatisticsWithDefaults

`func NewDtoTagStatisticsWithDefaults() *DtoTagStatistics`

NewDtoTagStatisticsWithDefaults instantiates a new DtoTagStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DtoTagStatistics) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DtoTagStatistics) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DtoTagStatistics) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DtoTagStatistics) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSubscriptions

`func (o *DtoTagStatistics) GetSubscriptions() []MoiraSubscriptionData`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *DtoTagStatistics) GetSubscriptionsOk() (*[]MoiraSubscriptionData, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *DtoTagStatistics) SetSubscriptions(v []MoiraSubscriptionData)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *DtoTagStatistics) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.

### GetTriggers

`func (o *DtoTagStatistics) GetTriggers() []string`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *DtoTagStatistics) GetTriggersOk() (*[]string, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *DtoTagStatistics) SetTriggers(v []string)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *DtoTagStatistics) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


