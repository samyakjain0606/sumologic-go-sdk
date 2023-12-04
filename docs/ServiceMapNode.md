# ServiceMapNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceName** | **string** | Name of a service in a service map. | 
**ServiceColor** | Pointer to **string** | Color hex code assigned to the service. | [optional] 
**LastSeenAt** | **time.Time** | The last time in UTC a service has been seen. Formatted as defined by date-time - RFC3339. | 
**IsRemote** | **bool** | Indicates whether node comes from inferred remote service or instrumented one. | 
**ServiceType** | **string** | Defines type of service. | 

## Methods

### NewServiceMapNode

`func NewServiceMapNode(serviceName string, lastSeenAt time.Time, isRemote bool, serviceType string, ) *ServiceMapNode`

NewServiceMapNode instantiates a new ServiceMapNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceMapNodeWithDefaults

`func NewServiceMapNodeWithDefaults() *ServiceMapNode`

NewServiceMapNodeWithDefaults instantiates a new ServiceMapNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceName

`func (o *ServiceMapNode) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *ServiceMapNode) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *ServiceMapNode) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.


### GetServiceColor

`func (o *ServiceMapNode) GetServiceColor() string`

GetServiceColor returns the ServiceColor field if non-nil, zero value otherwise.

### GetServiceColorOk

`func (o *ServiceMapNode) GetServiceColorOk() (*string, bool)`

GetServiceColorOk returns a tuple with the ServiceColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceColor

`func (o *ServiceMapNode) SetServiceColor(v string)`

SetServiceColor sets ServiceColor field to given value.

### HasServiceColor

`func (o *ServiceMapNode) HasServiceColor() bool`

HasServiceColor returns a boolean if a field has been set.

### GetLastSeenAt

`func (o *ServiceMapNode) GetLastSeenAt() time.Time`

GetLastSeenAt returns the LastSeenAt field if non-nil, zero value otherwise.

### GetLastSeenAtOk

`func (o *ServiceMapNode) GetLastSeenAtOk() (*time.Time, bool)`

GetLastSeenAtOk returns a tuple with the LastSeenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenAt

`func (o *ServiceMapNode) SetLastSeenAt(v time.Time)`

SetLastSeenAt sets LastSeenAt field to given value.


### GetIsRemote

`func (o *ServiceMapNode) GetIsRemote() bool`

GetIsRemote returns the IsRemote field if non-nil, zero value otherwise.

### GetIsRemoteOk

`func (o *ServiceMapNode) GetIsRemoteOk() (*bool, bool)`

GetIsRemoteOk returns a tuple with the IsRemote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemote

`func (o *ServiceMapNode) SetIsRemote(v bool)`

SetIsRemote sets IsRemote field to given value.


### GetServiceType

`func (o *ServiceMapNode) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *ServiceMapNode) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *ServiceMapNode) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


