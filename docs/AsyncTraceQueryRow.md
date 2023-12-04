# AsyncTraceQueryRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | [**TraceQueryExpression**](TraceQueryExpression.md) |  | 
**RowId** | **string** | An identifier used to reference this particular row of the query request while fetching a query result. Within a query, row ids must have distinct values. | 
**OrderBy** | Pointer to [**OrderBy**](OrderBy.md) |  | [optional] 

## Methods

### NewAsyncTraceQueryRow

`func NewAsyncTraceQueryRow(query TraceQueryExpression, rowId string, ) *AsyncTraceQueryRow`

NewAsyncTraceQueryRow instantiates a new AsyncTraceQueryRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAsyncTraceQueryRowWithDefaults

`func NewAsyncTraceQueryRowWithDefaults() *AsyncTraceQueryRow`

NewAsyncTraceQueryRowWithDefaults instantiates a new AsyncTraceQueryRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *AsyncTraceQueryRow) GetQuery() TraceQueryExpression`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *AsyncTraceQueryRow) GetQueryOk() (*TraceQueryExpression, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *AsyncTraceQueryRow) SetQuery(v TraceQueryExpression)`

SetQuery sets Query field to given value.


### GetRowId

`func (o *AsyncTraceQueryRow) GetRowId() string`

GetRowId returns the RowId field if non-nil, zero value otherwise.

### GetRowIdOk

`func (o *AsyncTraceQueryRow) GetRowIdOk() (*string, bool)`

GetRowIdOk returns a tuple with the RowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowId

`func (o *AsyncTraceQueryRow) SetRowId(v string)`

SetRowId sets RowId field to given value.


### GetOrderBy

`func (o *AsyncTraceQueryRow) GetOrderBy() OrderBy`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *AsyncTraceQueryRow) GetOrderByOk() (*OrderBy, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *AsyncTraceQueryRow) SetOrderBy(v OrderBy)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *AsyncTraceQueryRow) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


