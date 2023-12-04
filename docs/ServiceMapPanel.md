# ServiceMapPanel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Application** | Pointer to **string** | Filter services by the application custom tag. | [optional] 
**Service** | Pointer to **string** | Show only the specific service and its connections to other services. | [optional] 
**ShowRemoteServices** | Pointer to **bool** | Show remote services, like databases or external calls, automatically detected in client traffic. | [optional] 
**Environment** | Pointer to **string** | Show only service map data specific to the provided environment. | [optional] 

## Methods

### NewServiceMapPanel

`func NewServiceMapPanel() *ServiceMapPanel`

NewServiceMapPanel instantiates a new ServiceMapPanel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceMapPanelWithDefaults

`func NewServiceMapPanelWithDefaults() *ServiceMapPanel`

NewServiceMapPanelWithDefaults instantiates a new ServiceMapPanel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplication

`func (o *ServiceMapPanel) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *ServiceMapPanel) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *ServiceMapPanel) SetApplication(v string)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *ServiceMapPanel) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetService

`func (o *ServiceMapPanel) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *ServiceMapPanel) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *ServiceMapPanel) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *ServiceMapPanel) HasService() bool`

HasService returns a boolean if a field has been set.

### GetShowRemoteServices

`func (o *ServiceMapPanel) GetShowRemoteServices() bool`

GetShowRemoteServices returns the ShowRemoteServices field if non-nil, zero value otherwise.

### GetShowRemoteServicesOk

`func (o *ServiceMapPanel) GetShowRemoteServicesOk() (*bool, bool)`

GetShowRemoteServicesOk returns a tuple with the ShowRemoteServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowRemoteServices

`func (o *ServiceMapPanel) SetShowRemoteServices(v bool)`

SetShowRemoteServices sets ShowRemoteServices field to given value.

### HasShowRemoteServices

`func (o *ServiceMapPanel) HasShowRemoteServices() bool`

HasShowRemoteServices returns a boolean if a field has been set.

### GetEnvironment

`func (o *ServiceMapPanel) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ServiceMapPanel) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ServiceMapPanel) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ServiceMapPanel) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


