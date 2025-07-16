# ApiWebConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contacts** | Pointer to [**[]ApiWebContact**](ApiWebContact.md) |  | [optional] 
**FeatureFlags** | Pointer to [**ApiFeatureFlags**](ApiFeatureFlags.md) |  | [optional] 
**MetricSourceClusters** | Pointer to [**[]ApiMetricSourceCluster**](ApiMetricSourceCluster.md) |  | [optional] 
**RemoteAllowed** | Pointer to **bool** |  | [optional] 
**Sentry** | Pointer to [**ApiSentry**](ApiSentry.md) |  | [optional] 
**SupportEmail** | Pointer to **string** |  | [optional] 

## Methods

### NewApiWebConfig

`func NewApiWebConfig() *ApiWebConfig`

NewApiWebConfig instantiates a new ApiWebConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiWebConfigWithDefaults

`func NewApiWebConfigWithDefaults() *ApiWebConfig`

NewApiWebConfigWithDefaults instantiates a new ApiWebConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContacts

`func (o *ApiWebConfig) GetContacts() []ApiWebContact`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *ApiWebConfig) GetContactsOk() (*[]ApiWebContact, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *ApiWebConfig) SetContacts(v []ApiWebContact)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *ApiWebConfig) HasContacts() bool`

HasContacts returns a boolean if a field has been set.

### GetFeatureFlags

`func (o *ApiWebConfig) GetFeatureFlags() ApiFeatureFlags`

GetFeatureFlags returns the FeatureFlags field if non-nil, zero value otherwise.

### GetFeatureFlagsOk

`func (o *ApiWebConfig) GetFeatureFlagsOk() (*ApiFeatureFlags, bool)`

GetFeatureFlagsOk returns a tuple with the FeatureFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureFlags

`func (o *ApiWebConfig) SetFeatureFlags(v ApiFeatureFlags)`

SetFeatureFlags sets FeatureFlags field to given value.

### HasFeatureFlags

`func (o *ApiWebConfig) HasFeatureFlags() bool`

HasFeatureFlags returns a boolean if a field has been set.

### GetMetricSourceClusters

`func (o *ApiWebConfig) GetMetricSourceClusters() []ApiMetricSourceCluster`

GetMetricSourceClusters returns the MetricSourceClusters field if non-nil, zero value otherwise.

### GetMetricSourceClustersOk

`func (o *ApiWebConfig) GetMetricSourceClustersOk() (*[]ApiMetricSourceCluster, bool)`

GetMetricSourceClustersOk returns a tuple with the MetricSourceClusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricSourceClusters

`func (o *ApiWebConfig) SetMetricSourceClusters(v []ApiMetricSourceCluster)`

SetMetricSourceClusters sets MetricSourceClusters field to given value.

### HasMetricSourceClusters

`func (o *ApiWebConfig) HasMetricSourceClusters() bool`

HasMetricSourceClusters returns a boolean if a field has been set.

### GetRemoteAllowed

`func (o *ApiWebConfig) GetRemoteAllowed() bool`

GetRemoteAllowed returns the RemoteAllowed field if non-nil, zero value otherwise.

### GetRemoteAllowedOk

`func (o *ApiWebConfig) GetRemoteAllowedOk() (*bool, bool)`

GetRemoteAllowedOk returns a tuple with the RemoteAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAllowed

`func (o *ApiWebConfig) SetRemoteAllowed(v bool)`

SetRemoteAllowed sets RemoteAllowed field to given value.

### HasRemoteAllowed

`func (o *ApiWebConfig) HasRemoteAllowed() bool`

HasRemoteAllowed returns a boolean if a field has been set.

### GetSentry

`func (o *ApiWebConfig) GetSentry() ApiSentry`

GetSentry returns the Sentry field if non-nil, zero value otherwise.

### GetSentryOk

`func (o *ApiWebConfig) GetSentryOk() (*ApiSentry, bool)`

GetSentryOk returns a tuple with the Sentry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentry

`func (o *ApiWebConfig) SetSentry(v ApiSentry)`

SetSentry sets Sentry field to given value.

### HasSentry

`func (o *ApiWebConfig) HasSentry() bool`

HasSentry returns a boolean if a field has been set.

### GetSupportEmail

`func (o *ApiWebConfig) GetSupportEmail() string`

GetSupportEmail returns the SupportEmail field if non-nil, zero value otherwise.

### GetSupportEmailOk

`func (o *ApiWebConfig) GetSupportEmailOk() (*string, bool)`

GetSupportEmailOk returns a tuple with the SupportEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportEmail

`func (o *ApiWebConfig) SetSupportEmail(v string)`

SetSupportEmail sets SupportEmail field to given value.

### HasSupportEmail

`func (o *ApiWebConfig) HasSupportEmail() bool`

HasSupportEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


