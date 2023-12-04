# CpcQueryRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | [**TraceQueryExpression**](TraceQueryExpression.md) |  | 
**RowId** | **string** | An identifier used to reference this particular row of the query request while fetching a query result. Within a query, row ids must have distinct values. | 

## Methods

### NewCpcQueryRow

`func NewCpcQueryRow(query TraceQueryExpression, rowId string, ) *CpcQueryRow`

NewCpcQueryRow instantiates a new CpcQueryRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCpcQueryRowWithDefaults

`func NewCpcQueryRowWithDefaults() *CpcQueryRow`

NewCpcQueryRowWithDefaults instantiates a new CpcQueryRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *CpcQueryRow) GetQuery() TraceQueryExpression`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *CpcQueryRow) GetQueryOk() (*TraceQueryExpression, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *CpcQueryRow) SetQuery(v TraceQueryExpression)`

SetQuery sets Query field to given value.


### GetRowId

`func (o *CpcQueryRow) GetRowId() string`

GetRowId returns the RowId field if non-nil, zero value otherwise.

### GetRowIdOk

`func (o *CpcQueryRow) GetRowIdOk() (*string, bool)`

GetRowIdOk returns a tuple with the RowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowId

`func (o *CpcQueryRow) SetRowId(v string)`

SetRowId sets RowId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


