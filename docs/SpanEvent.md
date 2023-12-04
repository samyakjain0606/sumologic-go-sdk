# SpanEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **time.Time** | Time when an event happened in the [ISO 8601 / RFC3339](https://tools.ietf.org/html/rfc3339) format. | 
**Name** | **string** | Name of the event. | 
**Attributes** | Pointer to [**[]SpanEventAttribute**](SpanEventAttribute.md) | Span event attributes. | [optional] 

## Methods

### NewSpanEvent

`func NewSpanEvent(timestamp time.Time, name string, ) *SpanEvent`

NewSpanEvent instantiates a new SpanEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpanEventWithDefaults

`func NewSpanEventWithDefaults() *SpanEvent`

NewSpanEventWithDefaults instantiates a new SpanEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *SpanEvent) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SpanEvent) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SpanEvent) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetName

`func (o *SpanEvent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SpanEvent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SpanEvent) SetName(v string)`

SetName sets Name field to given value.


### GetAttributes

`func (o *SpanEvent) GetAttributes() []SpanEventAttribute`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SpanEvent) GetAttributesOk() (*[]SpanEventAttribute, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SpanEvent) SetAttributes(v []SpanEventAttribute)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *SpanEvent) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


