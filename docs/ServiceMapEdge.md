# ServiceMapEdge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | **string** | Name of a source service. Edge is directed from source to target. | 
**Target** | **string** | Name of a target service. Edge is directed from source to target. | 
**LastSeenAt** | **time.Time** | The last time in UTC an edge has been seen. Formatted as defined by date-time - RFC3339. | 

## Methods

### NewServiceMapEdge

`func NewServiceMapEdge(source string, target string, lastSeenAt time.Time, ) *ServiceMapEdge`

NewServiceMapEdge instantiates a new ServiceMapEdge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceMapEdgeWithDefaults

`func NewServiceMapEdgeWithDefaults() *ServiceMapEdge`

NewServiceMapEdgeWithDefaults instantiates a new ServiceMapEdge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *ServiceMapEdge) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ServiceMapEdge) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ServiceMapEdge) SetSource(v string)`

SetSource sets Source field to given value.


### GetTarget

`func (o *ServiceMapEdge) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *ServiceMapEdge) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *ServiceMapEdge) SetTarget(v string)`

SetTarget sets Target field to given value.


### GetLastSeenAt

`func (o *ServiceMapEdge) GetLastSeenAt() time.Time`

GetLastSeenAt returns the LastSeenAt field if non-nil, zero value otherwise.

### GetLastSeenAtOk

`func (o *ServiceMapEdge) GetLastSeenAtOk() (*time.Time, bool)`

GetLastSeenAtOk returns a tuple with the LastSeenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenAt

`func (o *ServiceMapEdge) SetLastSeenAt(v time.Time)`

SetLastSeenAt sets LastSeenAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


