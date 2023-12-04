# CriticalPathServiceBreakdownSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Elements** | [**[]CriticalPathServiceBreakdownElementBase**](CriticalPathServiceBreakdownElementBase.md) | List of the elements representing the critical path service duration breakdown - contains the first few services with the longest overall duration of the spans contributing to the critical path. | 
**OtherServicesDuration** | **int64** | Overall processing time in nanoseconds consumed by the rest of the spans in the critical path (a sum of the duration times of the spans&#39; critical path segments). | 
**IdleTime** | **int64** | Overall time in nanoseconds when no particular operation was in progress. | 

## Methods

### NewCriticalPathServiceBreakdownSummary

`func NewCriticalPathServiceBreakdownSummary(elements []CriticalPathServiceBreakdownElementBase, otherServicesDuration int64, idleTime int64, ) *CriticalPathServiceBreakdownSummary`

NewCriticalPathServiceBreakdownSummary instantiates a new CriticalPathServiceBreakdownSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCriticalPathServiceBreakdownSummaryWithDefaults

`func NewCriticalPathServiceBreakdownSummaryWithDefaults() *CriticalPathServiceBreakdownSummary`

NewCriticalPathServiceBreakdownSummaryWithDefaults instantiates a new CriticalPathServiceBreakdownSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetElements

`func (o *CriticalPathServiceBreakdownSummary) GetElements() []CriticalPathServiceBreakdownElementBase`

GetElements returns the Elements field if non-nil, zero value otherwise.

### GetElementsOk

`func (o *CriticalPathServiceBreakdownSummary) GetElementsOk() (*[]CriticalPathServiceBreakdownElementBase, bool)`

GetElementsOk returns a tuple with the Elements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElements

`func (o *CriticalPathServiceBreakdownSummary) SetElements(v []CriticalPathServiceBreakdownElementBase)`

SetElements sets Elements field to given value.


### GetOtherServicesDuration

`func (o *CriticalPathServiceBreakdownSummary) GetOtherServicesDuration() int64`

GetOtherServicesDuration returns the OtherServicesDuration field if non-nil, zero value otherwise.

### GetOtherServicesDurationOk

`func (o *CriticalPathServiceBreakdownSummary) GetOtherServicesDurationOk() (*int64, bool)`

GetOtherServicesDurationOk returns a tuple with the OtherServicesDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherServicesDuration

`func (o *CriticalPathServiceBreakdownSummary) SetOtherServicesDuration(v int64)`

SetOtherServicesDuration sets OtherServicesDuration field to given value.


### GetIdleTime

`func (o *CriticalPathServiceBreakdownSummary) GetIdleTime() int64`

GetIdleTime returns the IdleTime field if non-nil, zero value otherwise.

### GetIdleTimeOk

`func (o *CriticalPathServiceBreakdownSummary) GetIdleTimeOk() (*int64, bool)`

GetIdleTimeOk returns a tuple with the IdleTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleTime

`func (o *CriticalPathServiceBreakdownSummary) SetIdleTime(v int64)`

SetIdleTime sets IdleTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


