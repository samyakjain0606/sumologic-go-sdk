# SpanQueryRowResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RowId** | **string** | A unique identifier of the query. | 
**Errors** | Pointer to [**[]SpanQueryRowError**](SpanQueryRowError.md) | List of errors which occured when executing the query | [optional] 
**IsAggregation** | **bool** | Indicates whether this query is an aggregation | [default to false]
**ExecutedQuery** | Pointer to **string** | The executed query after rewriting | [optional] 

## Methods

### NewSpanQueryRowResponse

`func NewSpanQueryRowResponse(rowId string, isAggregation bool, ) *SpanQueryRowResponse`

NewSpanQueryRowResponse instantiates a new SpanQueryRowResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpanQueryRowResponseWithDefaults

`func NewSpanQueryRowResponseWithDefaults() *SpanQueryRowResponse`

NewSpanQueryRowResponseWithDefaults instantiates a new SpanQueryRowResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRowId

`func (o *SpanQueryRowResponse) GetRowId() string`

GetRowId returns the RowId field if non-nil, zero value otherwise.

### GetRowIdOk

`func (o *SpanQueryRowResponse) GetRowIdOk() (*string, bool)`

GetRowIdOk returns a tuple with the RowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowId

`func (o *SpanQueryRowResponse) SetRowId(v string)`

SetRowId sets RowId field to given value.


### GetErrors

`func (o *SpanQueryRowResponse) GetErrors() []SpanQueryRowError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *SpanQueryRowResponse) GetErrorsOk() (*[]SpanQueryRowError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *SpanQueryRowResponse) SetErrors(v []SpanQueryRowError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *SpanQueryRowResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetIsAggregation

`func (o *SpanQueryRowResponse) GetIsAggregation() bool`

GetIsAggregation returns the IsAggregation field if non-nil, zero value otherwise.

### GetIsAggregationOk

`func (o *SpanQueryRowResponse) GetIsAggregationOk() (*bool, bool)`

GetIsAggregationOk returns a tuple with the IsAggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAggregation

`func (o *SpanQueryRowResponse) SetIsAggregation(v bool)`

SetIsAggregation sets IsAggregation field to given value.


### GetExecutedQuery

`func (o *SpanQueryRowResponse) GetExecutedQuery() string`

GetExecutedQuery returns the ExecutedQuery field if non-nil, zero value otherwise.

### GetExecutedQueryOk

`func (o *SpanQueryRowResponse) GetExecutedQueryOk() (*string, bool)`

GetExecutedQueryOk returns a tuple with the ExecutedQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedQuery

`func (o *SpanQueryRowResponse) SetExecutedQuery(v string)`

SetExecutedQuery sets ExecutedQuery field to given value.

### HasExecutedQuery

`func (o *SpanQueryRowResponse) HasExecutedQuery() bool`

HasExecutedQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


