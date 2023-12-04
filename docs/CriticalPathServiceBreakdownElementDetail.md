# CriticalPathServiceBreakdownElementDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Service** | Pointer to **string** | The name of the service. | [optional] 
**ServiceColor** | Pointer to **string** | Color hex code assigned to the service. | [optional] 
**Duration** | **int64** | Overall processing time in nanoseconds consumed by the spans belonging to this service in the critical path (a sum of the duration times of the spans&#39; critical path segments). | 
**NumSpans** | **int32** | Number of spans that are part of this service. | 
**LongestSegmentDuration** | **int64** | Number of nanoseconds the longest span segment in the critical path lasted. | 

## Methods

### NewCriticalPathServiceBreakdownElementDetail

`func NewCriticalPathServiceBreakdownElementDetail(duration int64, numSpans int32, longestSegmentDuration int64, ) *CriticalPathServiceBreakdownElementDetail`

NewCriticalPathServiceBreakdownElementDetail instantiates a new CriticalPathServiceBreakdownElementDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCriticalPathServiceBreakdownElementDetailWithDefaults

`func NewCriticalPathServiceBreakdownElementDetailWithDefaults() *CriticalPathServiceBreakdownElementDetail`

NewCriticalPathServiceBreakdownElementDetailWithDefaults instantiates a new CriticalPathServiceBreakdownElementDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetService

`func (o *CriticalPathServiceBreakdownElementDetail) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *CriticalPathServiceBreakdownElementDetail) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *CriticalPathServiceBreakdownElementDetail) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *CriticalPathServiceBreakdownElementDetail) HasService() bool`

HasService returns a boolean if a field has been set.

### GetServiceColor

`func (o *CriticalPathServiceBreakdownElementDetail) GetServiceColor() string`

GetServiceColor returns the ServiceColor field if non-nil, zero value otherwise.

### GetServiceColorOk

`func (o *CriticalPathServiceBreakdownElementDetail) GetServiceColorOk() (*string, bool)`

GetServiceColorOk returns a tuple with the ServiceColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceColor

`func (o *CriticalPathServiceBreakdownElementDetail) SetServiceColor(v string)`

SetServiceColor sets ServiceColor field to given value.

### HasServiceColor

`func (o *CriticalPathServiceBreakdownElementDetail) HasServiceColor() bool`

HasServiceColor returns a boolean if a field has been set.

### GetDuration

`func (o *CriticalPathServiceBreakdownElementDetail) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *CriticalPathServiceBreakdownElementDetail) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *CriticalPathServiceBreakdownElementDetail) SetDuration(v int64)`

SetDuration sets Duration field to given value.


### GetNumSpans

`func (o *CriticalPathServiceBreakdownElementDetail) GetNumSpans() int32`

GetNumSpans returns the NumSpans field if non-nil, zero value otherwise.

### GetNumSpansOk

`func (o *CriticalPathServiceBreakdownElementDetail) GetNumSpansOk() (*int32, bool)`

GetNumSpansOk returns a tuple with the NumSpans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSpans

`func (o *CriticalPathServiceBreakdownElementDetail) SetNumSpans(v int32)`

SetNumSpans sets NumSpans field to given value.


### GetLongestSegmentDuration

`func (o *CriticalPathServiceBreakdownElementDetail) GetLongestSegmentDuration() int64`

GetLongestSegmentDuration returns the LongestSegmentDuration field if non-nil, zero value otherwise.

### GetLongestSegmentDurationOk

`func (o *CriticalPathServiceBreakdownElementDetail) GetLongestSegmentDurationOk() (*int64, bool)`

GetLongestSegmentDurationOk returns a tuple with the LongestSegmentDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongestSegmentDuration

`func (o *CriticalPathServiceBreakdownElementDetail) SetLongestSegmentDuration(v int64)`

SetLongestSegmentDuration sets LongestSegmentDuration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


