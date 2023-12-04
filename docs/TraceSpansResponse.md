# TraceSpansResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SpanPage** | Pointer to [**[]TraceSpan**](TraceSpan.md) | List of trace spans. | [optional] 
**TotalCount** | **int64** | Total count of spans for this trace. | 
**Next** | Pointer to **string** | Next continuation token. | [optional] 

## Methods

### NewTraceSpansResponse

`func NewTraceSpansResponse(totalCount int64, ) *TraceSpansResponse`

NewTraceSpansResponse instantiates a new TraceSpansResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTraceSpansResponseWithDefaults

`func NewTraceSpansResponseWithDefaults() *TraceSpansResponse`

NewTraceSpansResponseWithDefaults instantiates a new TraceSpansResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpanPage

`func (o *TraceSpansResponse) GetSpanPage() []TraceSpan`

GetSpanPage returns the SpanPage field if non-nil, zero value otherwise.

### GetSpanPageOk

`func (o *TraceSpansResponse) GetSpanPageOk() (*[]TraceSpan, bool)`

GetSpanPageOk returns a tuple with the SpanPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanPage

`func (o *TraceSpansResponse) SetSpanPage(v []TraceSpan)`

SetSpanPage sets SpanPage field to given value.

### HasSpanPage

`func (o *TraceSpansResponse) HasSpanPage() bool`

HasSpanPage returns a boolean if a field has been set.

### GetTotalCount

`func (o *TraceSpansResponse) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *TraceSpansResponse) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *TraceSpansResponse) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.


### GetNext

`func (o *TraceSpansResponse) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *TraceSpansResponse) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *TraceSpansResponse) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *TraceSpansResponse) HasNext() bool`

HasNext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


