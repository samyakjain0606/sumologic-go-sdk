# TraceQueryRowStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RowId** | **string** | A unique identifier of the query. | 
**Status** | **string** | Status of the query. Possible values: &#x60;Processing&#x60;, &#x60;Finished&#x60;, &#x60;Error&#x60;, &#x60;Canceled&#x60;. | 
**StatusMessage** | Pointer to **string** | Descriptive message of the status | [optional] 
**Count** | **int64** | Number of results matching the query | 

## Methods

### NewTraceQueryRowStatus

`func NewTraceQueryRowStatus(rowId string, status string, count int64, ) *TraceQueryRowStatus`

NewTraceQueryRowStatus instantiates a new TraceQueryRowStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTraceQueryRowStatusWithDefaults

`func NewTraceQueryRowStatusWithDefaults() *TraceQueryRowStatus`

NewTraceQueryRowStatusWithDefaults instantiates a new TraceQueryRowStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRowId

`func (o *TraceQueryRowStatus) GetRowId() string`

GetRowId returns the RowId field if non-nil, zero value otherwise.

### GetRowIdOk

`func (o *TraceQueryRowStatus) GetRowIdOk() (*string, bool)`

GetRowIdOk returns a tuple with the RowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowId

`func (o *TraceQueryRowStatus) SetRowId(v string)`

SetRowId sets RowId field to given value.


### GetStatus

`func (o *TraceQueryRowStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TraceQueryRowStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TraceQueryRowStatus) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStatusMessage

`func (o *TraceQueryRowStatus) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *TraceQueryRowStatus) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *TraceQueryRowStatus) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *TraceQueryRowStatus) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetCount

`func (o *TraceQueryRowStatus) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *TraceQueryRowStatus) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *TraceQueryRowStatus) SetCount(v int64)`

SetCount sets Count field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


