# CriticalPathServiceBreakdownElementBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Service** | Pointer to **string** | The name of the service. | [optional] 
**ServiceColor** | Pointer to **string** | Color hex code assigned to the service. | [optional] 
**Duration** | **int64** | Overall processing time in nanoseconds consumed by the spans belonging to this service in the critical path (a sum of the duration times of the spans&#39; critical path segments). | 

## Methods

### NewCriticalPathServiceBreakdownElementBase

`func NewCriticalPathServiceBreakdownElementBase(duration int64, ) *CriticalPathServiceBreakdownElementBase`

NewCriticalPathServiceBreakdownElementBase instantiates a new CriticalPathServiceBreakdownElementBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCriticalPathServiceBreakdownElementBaseWithDefaults

`func NewCriticalPathServiceBreakdownElementBaseWithDefaults() *CriticalPathServiceBreakdownElementBase`

NewCriticalPathServiceBreakdownElementBaseWithDefaults instantiates a new CriticalPathServiceBreakdownElementBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetService

`func (o *CriticalPathServiceBreakdownElementBase) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *CriticalPathServiceBreakdownElementBase) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *CriticalPathServiceBreakdownElementBase) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *CriticalPathServiceBreakdownElementBase) HasService() bool`

HasService returns a boolean if a field has been set.

### GetServiceColor

`func (o *CriticalPathServiceBreakdownElementBase) GetServiceColor() string`

GetServiceColor returns the ServiceColor field if non-nil, zero value otherwise.

### GetServiceColorOk

`func (o *CriticalPathServiceBreakdownElementBase) GetServiceColorOk() (*string, bool)`

GetServiceColorOk returns a tuple with the ServiceColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceColor

`func (o *CriticalPathServiceBreakdownElementBase) SetServiceColor(v string)`

SetServiceColor sets ServiceColor field to given value.

### HasServiceColor

`func (o *CriticalPathServiceBreakdownElementBase) HasServiceColor() bool`

HasServiceColor returns a boolean if a field has been set.

### GetDuration

`func (o *CriticalPathServiceBreakdownElementBase) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *CriticalPathServiceBreakdownElementBase) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *CriticalPathServiceBreakdownElementBase) SetDuration(v int64)`

SetDuration sets Duration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


