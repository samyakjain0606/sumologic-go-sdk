# SpanQueryRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryString** | **string** | Query string using the log search syntax. | 
**RowId** | **string** | An identifier used to reference this particular row of the query request. Within a query, row ids must have distinct values. | 

## Methods

### NewSpanQueryRow

`func NewSpanQueryRow(queryString string, rowId string, ) *SpanQueryRow`

NewSpanQueryRow instantiates a new SpanQueryRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpanQueryRowWithDefaults

`func NewSpanQueryRowWithDefaults() *SpanQueryRow`

NewSpanQueryRowWithDefaults instantiates a new SpanQueryRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueryString

`func (o *SpanQueryRow) GetQueryString() string`

GetQueryString returns the QueryString field if non-nil, zero value otherwise.

### GetQueryStringOk

`func (o *SpanQueryRow) GetQueryStringOk() (*string, bool)`

GetQueryStringOk returns a tuple with the QueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryString

`func (o *SpanQueryRow) SetQueryString(v string)`

SetQueryString sets QueryString field to given value.


### GetRowId

`func (o *SpanQueryRow) GetRowId() string`

GetRowId returns the RowId field if non-nil, zero value otherwise.

### GetRowIdOk

`func (o *SpanQueryRow) GetRowIdOk() (*string, bool)`

GetRowIdOk returns a tuple with the RowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowId

`func (o *SpanQueryRow) SetRowId(v string)`

SetRowId sets RowId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


