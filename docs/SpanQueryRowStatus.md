# SpanQueryRowStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RowId** | **string** | A unique identifier of the query. | 
**Status** | **string** | Status of the query. Possible values: &#x60;Processing&#x60;, &#x60;Finished&#x60;, &#x60;Error&#x60;, &#x60;Paused&#x60;. | 
**StatusMessage** | Pointer to **string** | Descriptive message of the status. | [optional] 
**Count** | **int64** | Number of results matching the query | 
**ApproximatedFieldCounts** | Pointer to **bool** | Indicates whether facet field cardinality counts are approximated or not. | [optional] 
**FacetsCompleted** | Pointer to **bool** | Indicates whether facets calculation has completed. | [optional] 

## Methods

### NewSpanQueryRowStatus

`func NewSpanQueryRowStatus(rowId string, status string, count int64, ) *SpanQueryRowStatus`

NewSpanQueryRowStatus instantiates a new SpanQueryRowStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpanQueryRowStatusWithDefaults

`func NewSpanQueryRowStatusWithDefaults() *SpanQueryRowStatus`

NewSpanQueryRowStatusWithDefaults instantiates a new SpanQueryRowStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRowId

`func (o *SpanQueryRowStatus) GetRowId() string`

GetRowId returns the RowId field if non-nil, zero value otherwise.

### GetRowIdOk

`func (o *SpanQueryRowStatus) GetRowIdOk() (*string, bool)`

GetRowIdOk returns a tuple with the RowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowId

`func (o *SpanQueryRowStatus) SetRowId(v string)`

SetRowId sets RowId field to given value.


### GetStatus

`func (o *SpanQueryRowStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SpanQueryRowStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SpanQueryRowStatus) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStatusMessage

`func (o *SpanQueryRowStatus) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *SpanQueryRowStatus) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *SpanQueryRowStatus) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *SpanQueryRowStatus) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetCount

`func (o *SpanQueryRowStatus) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *SpanQueryRowStatus) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *SpanQueryRowStatus) SetCount(v int64)`

SetCount sets Count field to given value.


### GetApproximatedFieldCounts

`func (o *SpanQueryRowStatus) GetApproximatedFieldCounts() bool`

GetApproximatedFieldCounts returns the ApproximatedFieldCounts field if non-nil, zero value otherwise.

### GetApproximatedFieldCountsOk

`func (o *SpanQueryRowStatus) GetApproximatedFieldCountsOk() (*bool, bool)`

GetApproximatedFieldCountsOk returns a tuple with the ApproximatedFieldCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproximatedFieldCounts

`func (o *SpanQueryRowStatus) SetApproximatedFieldCounts(v bool)`

SetApproximatedFieldCounts sets ApproximatedFieldCounts field to given value.

### HasApproximatedFieldCounts

`func (o *SpanQueryRowStatus) HasApproximatedFieldCounts() bool`

HasApproximatedFieldCounts returns a boolean if a field has been set.

### GetFacetsCompleted

`func (o *SpanQueryRowStatus) GetFacetsCompleted() bool`

GetFacetsCompleted returns the FacetsCompleted field if non-nil, zero value otherwise.

### GetFacetsCompletedOk

`func (o *SpanQueryRowStatus) GetFacetsCompletedOk() (*bool, bool)`

GetFacetsCompletedOk returns a tuple with the FacetsCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacetsCompleted

`func (o *SpanQueryRowStatus) SetFacetsCompleted(v bool)`

SetFacetsCompleted sets FacetsCompleted field to given value.

### HasFacetsCompleted

`func (o *SpanQueryRowStatus) HasFacetsCompleted() bool`

HasFacetsCompleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


