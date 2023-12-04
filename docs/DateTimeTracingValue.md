# DateTimeTracingValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **time.Time** | Timestamp in UTC in the [ISO 8601 / RFC3339](https://tools.ietf.org/html/rfc3339) format. | 

## Methods

### NewDateTimeTracingValue

`func NewDateTimeTracingValue(value time.Time, ) *DateTimeTracingValue`

NewDateTimeTracingValue instantiates a new DateTimeTracingValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDateTimeTracingValueWithDefaults

`func NewDateTimeTracingValueWithDefaults() *DateTimeTracingValue`

NewDateTimeTracingValueWithDefaults instantiates a new DateTimeTracingValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *DateTimeTracingValue) GetValue() time.Time`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DateTimeTracingValue) GetValueOk() (*time.Time, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DateTimeTracingValue) SetValue(v time.Time)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


