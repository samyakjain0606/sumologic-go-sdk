# CriticalPathServiceBreakdownResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Elements** | [**[]CriticalPathServiceBreakdownElementDetail**](CriticalPathServiceBreakdownElementDetail.md) | List of elements representing the critical path service breakdown. | 
**IdleTime** | **int64** | Overall time in nanoseconds when no particular operation was in progress. | 

## Methods

### NewCriticalPathServiceBreakdownResponse

`func NewCriticalPathServiceBreakdownResponse(elements []CriticalPathServiceBreakdownElementDetail, idleTime int64, ) *CriticalPathServiceBreakdownResponse`

NewCriticalPathServiceBreakdownResponse instantiates a new CriticalPathServiceBreakdownResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCriticalPathServiceBreakdownResponseWithDefaults

`func NewCriticalPathServiceBreakdownResponseWithDefaults() *CriticalPathServiceBreakdownResponse`

NewCriticalPathServiceBreakdownResponseWithDefaults instantiates a new CriticalPathServiceBreakdownResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetElements

`func (o *CriticalPathServiceBreakdownResponse) GetElements() []CriticalPathServiceBreakdownElementDetail`

GetElements returns the Elements field if non-nil, zero value otherwise.

### GetElementsOk

`func (o *CriticalPathServiceBreakdownResponse) GetElementsOk() (*[]CriticalPathServiceBreakdownElementDetail, bool)`

GetElementsOk returns a tuple with the Elements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElements

`func (o *CriticalPathServiceBreakdownResponse) SetElements(v []CriticalPathServiceBreakdownElementDetail)`

SetElements sets Elements field to given value.


### GetIdleTime

`func (o *CriticalPathServiceBreakdownResponse) GetIdleTime() int64`

GetIdleTime returns the IdleTime field if non-nil, zero value otherwise.

### GetIdleTimeOk

`func (o *CriticalPathServiceBreakdownResponse) GetIdleTimeOk() (*int64, bool)`

GetIdleTimeOk returns a tuple with the IdleTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleTime

`func (o *CriticalPathServiceBreakdownResponse) SetIdleTime(v int64)`

SetIdleTime sets IdleTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


