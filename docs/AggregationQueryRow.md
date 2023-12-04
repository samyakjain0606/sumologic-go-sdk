# AggregationQueryRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | [**TraceQueryExpression**](TraceQueryExpression.md) |  | 
**RowId** | **string** | An identifier used to reference this particular row of the query request while fetching a query result. Within a query, row ids must have distinct values. | 

## Methods

### NewAggregationQueryRow

`func NewAggregationQueryRow(query TraceQueryExpression, rowId string, ) *AggregationQueryRow`

NewAggregationQueryRow instantiates a new AggregationQueryRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregationQueryRowWithDefaults

`func NewAggregationQueryRowWithDefaults() *AggregationQueryRow`

NewAggregationQueryRowWithDefaults instantiates a new AggregationQueryRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *AggregationQueryRow) GetQuery() TraceQueryExpression`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *AggregationQueryRow) GetQueryOk() (*TraceQueryExpression, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *AggregationQueryRow) SetQuery(v TraceQueryExpression)`

SetQuery sets Query field to given value.


### GetRowId

`func (o *AggregationQueryRow) GetRowId() string`

GetRowId returns the RowId field if non-nil, zero value otherwise.

### GetRowIdOk

`func (o *AggregationQueryRow) GetRowIdOk() (*string, bool)`

GetRowIdOk returns a tuple with the RowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowId

`func (o *AggregationQueryRow) SetRowId(v string)`

SetRowId sets RowId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


