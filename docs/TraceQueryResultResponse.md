# TraceQueryResultResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | [**[]TraceDetail**](TraceDetail.md) | List of traces matching the query. | 
**Next** | Pointer to **string** | Next continuation token. | [optional] 

## Methods

### NewTraceQueryResultResponse

`func NewTraceQueryResultResponse(results []TraceDetail, ) *TraceQueryResultResponse`

NewTraceQueryResultResponse instantiates a new TraceQueryResultResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTraceQueryResultResponseWithDefaults

`func NewTraceQueryResultResponseWithDefaults() *TraceQueryResultResponse`

NewTraceQueryResultResponseWithDefaults instantiates a new TraceQueryResultResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *TraceQueryResultResponse) GetResults() []TraceDetail`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *TraceQueryResultResponse) GetResultsOk() (*[]TraceDetail, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *TraceQueryResultResponse) SetResults(v []TraceDetail)`

SetResults sets Results field to given value.


### GetNext

`func (o *TraceQueryResultResponse) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *TraceQueryResultResponse) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *TraceQueryResultResponse) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *TraceQueryResultResponse) HasNext() bool`

HasNext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


