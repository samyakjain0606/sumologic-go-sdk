# SpanQueryResultSpansResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SpanPage** | [**[]SpanQuerySpanData**](SpanQuerySpanData.md) | List of trace spans. | 
**Next** | Pointer to **string** | Next continuation token. | [optional] 

## Methods

### NewSpanQueryResultSpansResponse

`func NewSpanQueryResultSpansResponse(spanPage []SpanQuerySpanData, ) *SpanQueryResultSpansResponse`

NewSpanQueryResultSpansResponse instantiates a new SpanQueryResultSpansResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpanQueryResultSpansResponseWithDefaults

`func NewSpanQueryResultSpansResponseWithDefaults() *SpanQueryResultSpansResponse`

NewSpanQueryResultSpansResponseWithDefaults instantiates a new SpanQueryResultSpansResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpanPage

`func (o *SpanQueryResultSpansResponse) GetSpanPage() []SpanQuerySpanData`

GetSpanPage returns the SpanPage field if non-nil, zero value otherwise.

### GetSpanPageOk

`func (o *SpanQueryResultSpansResponse) GetSpanPageOk() (*[]SpanQuerySpanData, bool)`

GetSpanPageOk returns a tuple with the SpanPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanPage

`func (o *SpanQueryResultSpansResponse) SetSpanPage(v []SpanQuerySpanData)`

SetSpanPage sets SpanPage field to given value.


### GetNext

`func (o *SpanQueryResultSpansResponse) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *SpanQueryResultSpansResponse) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *SpanQueryResultSpansResponse) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *SpanQueryResultSpansResponse) HasNext() bool`

HasNext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


