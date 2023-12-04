# Iso8601TimeRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | **time.Time** | Start time in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format | 
**End** | **time.Time** | End time in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format | 

## Methods

### NewIso8601TimeRange

`func NewIso8601TimeRange(start time.Time, end time.Time, ) *Iso8601TimeRange`

NewIso8601TimeRange instantiates a new Iso8601TimeRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIso8601TimeRangeWithDefaults

`func NewIso8601TimeRangeWithDefaults() *Iso8601TimeRange`

NewIso8601TimeRangeWithDefaults instantiates a new Iso8601TimeRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *Iso8601TimeRange) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *Iso8601TimeRange) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *Iso8601TimeRange) SetStart(v time.Time)`

SetStart sets Start field to given value.


### GetEnd

`func (o *Iso8601TimeRange) GetEnd() time.Time`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *Iso8601TimeRange) GetEndOk() (*time.Time, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *Iso8601TimeRange) SetEnd(v time.Time)`

SetEnd sets End field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


