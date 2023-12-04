# TraceLightEventsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SpanEvents** | Pointer to [**map[string][]LightSpanEvent**](array.md) | Map of span ids to lists of their events, without their attributes. | [optional] 
**Next** | Pointer to **string** | Next continuation token. | [optional] 

## Methods

### NewTraceLightEventsResponse

`func NewTraceLightEventsResponse() *TraceLightEventsResponse`

NewTraceLightEventsResponse instantiates a new TraceLightEventsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTraceLightEventsResponseWithDefaults

`func NewTraceLightEventsResponseWithDefaults() *TraceLightEventsResponse`

NewTraceLightEventsResponseWithDefaults instantiates a new TraceLightEventsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpanEvents

`func (o *TraceLightEventsResponse) GetSpanEvents() map[string][]LightSpanEvent`

GetSpanEvents returns the SpanEvents field if non-nil, zero value otherwise.

### GetSpanEventsOk

`func (o *TraceLightEventsResponse) GetSpanEventsOk() (*map[string][]LightSpanEvent, bool)`

GetSpanEventsOk returns a tuple with the SpanEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanEvents

`func (o *TraceLightEventsResponse) SetSpanEvents(v map[string][]LightSpanEvent)`

SetSpanEvents sets SpanEvents field to given value.

### HasSpanEvents

`func (o *TraceLightEventsResponse) HasSpanEvents() bool`

HasSpanEvents returns a boolean if a field has been set.

### GetNext

`func (o *TraceLightEventsResponse) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *TraceLightEventsResponse) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *TraceLightEventsResponse) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *TraceLightEventsResponse) HasNext() bool`

HasNext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


