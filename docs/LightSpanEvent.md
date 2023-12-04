# LightSpanEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **time.Time** | Time when an event happened in the [ISO 8601 / RFC3339](https://tools.ietf.org/html/rfc3339) format. | 
**Name** | **string** | Name of the event. | 

## Methods

### NewLightSpanEvent

`func NewLightSpanEvent(timestamp time.Time, name string, ) *LightSpanEvent`

NewLightSpanEvent instantiates a new LightSpanEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLightSpanEventWithDefaults

`func NewLightSpanEventWithDefaults() *LightSpanEvent`

NewLightSpanEventWithDefaults instantiates a new LightSpanEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *LightSpanEvent) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *LightSpanEvent) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *LightSpanEvent) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetName

`func (o *LightSpanEvent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LightSpanEvent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LightSpanEvent) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


