# SpanQueryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryId** | **string** | Id of the created query | 
**QueryRows** | [**[]SpanQueryRowResponse**](SpanQueryRowResponse.md) | A list of row responses with details about individual queries. | 
**HasErrors** | Pointer to **bool** | Indicates whether there was an error while executing the query. | [optional] [default to false]
**TimeRange** | Pointer to [**BeginBoundedTimeRange**](BeginBoundedTimeRange.md) |  | [optional] 

## Methods

### NewSpanQueryResponse

`func NewSpanQueryResponse(queryId string, queryRows []SpanQueryRowResponse, ) *SpanQueryResponse`

NewSpanQueryResponse instantiates a new SpanQueryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpanQueryResponseWithDefaults

`func NewSpanQueryResponseWithDefaults() *SpanQueryResponse`

NewSpanQueryResponseWithDefaults instantiates a new SpanQueryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueryId

`func (o *SpanQueryResponse) GetQueryId() string`

GetQueryId returns the QueryId field if non-nil, zero value otherwise.

### GetQueryIdOk

`func (o *SpanQueryResponse) GetQueryIdOk() (*string, bool)`

GetQueryIdOk returns a tuple with the QueryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryId

`func (o *SpanQueryResponse) SetQueryId(v string)`

SetQueryId sets QueryId field to given value.


### GetQueryRows

`func (o *SpanQueryResponse) GetQueryRows() []SpanQueryRowResponse`

GetQueryRows returns the QueryRows field if non-nil, zero value otherwise.

### GetQueryRowsOk

`func (o *SpanQueryResponse) GetQueryRowsOk() (*[]SpanQueryRowResponse, bool)`

GetQueryRowsOk returns a tuple with the QueryRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryRows

`func (o *SpanQueryResponse) SetQueryRows(v []SpanQueryRowResponse)`

SetQueryRows sets QueryRows field to given value.


### GetHasErrors

`func (o *SpanQueryResponse) GetHasErrors() bool`

GetHasErrors returns the HasErrors field if non-nil, zero value otherwise.

### GetHasErrorsOk

`func (o *SpanQueryResponse) GetHasErrorsOk() (*bool, bool)`

GetHasErrorsOk returns a tuple with the HasErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasErrors

`func (o *SpanQueryResponse) SetHasErrors(v bool)`

SetHasErrors sets HasErrors field to given value.

### HasHasErrors

`func (o *SpanQueryResponse) HasHasErrors() bool`

HasHasErrors returns a boolean if a field has been set.

### GetTimeRange

`func (o *SpanQueryResponse) GetTimeRange() BeginBoundedTimeRange`

GetTimeRange returns the TimeRange field if non-nil, zero value otherwise.

### GetTimeRangeOk

`func (o *SpanQueryResponse) GetTimeRangeOk() (*BeginBoundedTimeRange, bool)`

GetTimeRangeOk returns a tuple with the TimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRange

`func (o *SpanQueryResponse) SetTimeRange(v BeginBoundedTimeRange)`

SetTimeRange sets TimeRange field to given value.

### HasTimeRange

`func (o *SpanQueryResponse) HasTimeRange() bool`

HasTimeRange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


