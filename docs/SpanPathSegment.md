# SpanPathSegment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SpanId** | **string** | Span identifier. | 
**Service** | Pointer to **string** | The name of the service this span is part of. | [optional] 
**ServiceColor** | Pointer to **string** | Color hex code assigned to the service. | [optional] 
**StartOffset** | **int64** | Number of nanoseconds from the span startedAt the segment started. | 
**Duration** | **int64** | Number of nanoseconds the span segment lasted. | 
**Fraction** | Pointer to **float64** | The fraction (value between 0.0 and 1.0) from the trace duration time this segment took. | [optional] 

## Methods

### NewSpanPathSegment

`func NewSpanPathSegment(spanId string, startOffset int64, duration int64, ) *SpanPathSegment`

NewSpanPathSegment instantiates a new SpanPathSegment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpanPathSegmentWithDefaults

`func NewSpanPathSegmentWithDefaults() *SpanPathSegment`

NewSpanPathSegmentWithDefaults instantiates a new SpanPathSegment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpanId

`func (o *SpanPathSegment) GetSpanId() string`

GetSpanId returns the SpanId field if non-nil, zero value otherwise.

### GetSpanIdOk

`func (o *SpanPathSegment) GetSpanIdOk() (*string, bool)`

GetSpanIdOk returns a tuple with the SpanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanId

`func (o *SpanPathSegment) SetSpanId(v string)`

SetSpanId sets SpanId field to given value.


### GetService

`func (o *SpanPathSegment) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *SpanPathSegment) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *SpanPathSegment) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *SpanPathSegment) HasService() bool`

HasService returns a boolean if a field has been set.

### GetServiceColor

`func (o *SpanPathSegment) GetServiceColor() string`

GetServiceColor returns the ServiceColor field if non-nil, zero value otherwise.

### GetServiceColorOk

`func (o *SpanPathSegment) GetServiceColorOk() (*string, bool)`

GetServiceColorOk returns a tuple with the ServiceColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceColor

`func (o *SpanPathSegment) SetServiceColor(v string)`

SetServiceColor sets ServiceColor field to given value.

### HasServiceColor

`func (o *SpanPathSegment) HasServiceColor() bool`

HasServiceColor returns a boolean if a field has been set.

### GetStartOffset

`func (o *SpanPathSegment) GetStartOffset() int64`

GetStartOffset returns the StartOffset field if non-nil, zero value otherwise.

### GetStartOffsetOk

`func (o *SpanPathSegment) GetStartOffsetOk() (*int64, bool)`

GetStartOffsetOk returns a tuple with the StartOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartOffset

`func (o *SpanPathSegment) SetStartOffset(v int64)`

SetStartOffset sets StartOffset field to given value.


### GetDuration

`func (o *SpanPathSegment) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *SpanPathSegment) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *SpanPathSegment) SetDuration(v int64)`

SetDuration sets Duration field to given value.


### GetFraction

`func (o *SpanPathSegment) GetFraction() float64`

GetFraction returns the Fraction field if non-nil, zero value otherwise.

### GetFractionOk

`func (o *SpanPathSegment) GetFractionOk() (*float64, bool)`

GetFractionOk returns a tuple with the Fraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFraction

`func (o *SpanPathSegment) SetFraction(v float64)`

SetFraction sets Fraction field to given value.

### HasFraction

`func (o *SpanPathSegment) HasFraction() bool`

HasFraction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


