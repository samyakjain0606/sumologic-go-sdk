# SliQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RowId** | **string** | Unique id of the row. Used for query arithmetic, only for metric queries. | 
**Query** | **string** | Query String. | 
**UseRowCount** | **bool** | Determines whether to use count of rows (for logs) or data points (for metrics) in query result or specific field. | 
**Field** | Pointer to **string** | Field of log query output to compare against. To be used only for logs based data type when &#x60;useRowCount&#x60; is false. | [optional] 

## Methods

### NewSliQuery

`func NewSliQuery(rowId string, query string, useRowCount bool, ) *SliQuery`

NewSliQuery instantiates a new SliQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSliQueryWithDefaults

`func NewSliQueryWithDefaults() *SliQuery`

NewSliQueryWithDefaults instantiates a new SliQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRowId

`func (o *SliQuery) GetRowId() string`

GetRowId returns the RowId field if non-nil, zero value otherwise.

### GetRowIdOk

`func (o *SliQuery) GetRowIdOk() (*string, bool)`

GetRowIdOk returns a tuple with the RowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowId

`func (o *SliQuery) SetRowId(v string)`

SetRowId sets RowId field to given value.


### GetQuery

`func (o *SliQuery) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *SliQuery) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *SliQuery) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetUseRowCount

`func (o *SliQuery) GetUseRowCount() bool`

GetUseRowCount returns the UseRowCount field if non-nil, zero value otherwise.

### GetUseRowCountOk

`func (o *SliQuery) GetUseRowCountOk() (*bool, bool)`

GetUseRowCountOk returns a tuple with the UseRowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRowCount

`func (o *SliQuery) SetUseRowCount(v bool)`

SetUseRowCount sets UseRowCount field to given value.


### GetField

`func (o *SliQuery) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *SliQuery) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *SliQuery) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *SliQuery) HasField() bool`

HasField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


