# SpanQueryAggregateResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | Status of the query. Possible values: &#x60;Processing&#x60;, &#x60;Finished&#x60;, &#x60;Error&#x60;, &#x60;Paused&#x60;. | 
**StatusMessage** | Pointer to **string** | Descriptive message of the status | [optional] 
**Series** | [**[]SpanQueryAggregateDataSeries**](SpanQueryAggregateDataSeries.md) | The series returned from a search. | 

## Methods

### NewSpanQueryAggregateResult

`func NewSpanQueryAggregateResult(status string, series []SpanQueryAggregateDataSeries, ) *SpanQueryAggregateResult`

NewSpanQueryAggregateResult instantiates a new SpanQueryAggregateResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpanQueryAggregateResultWithDefaults

`func NewSpanQueryAggregateResultWithDefaults() *SpanQueryAggregateResult`

NewSpanQueryAggregateResultWithDefaults instantiates a new SpanQueryAggregateResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *SpanQueryAggregateResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SpanQueryAggregateResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SpanQueryAggregateResult) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStatusMessage

`func (o *SpanQueryAggregateResult) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *SpanQueryAggregateResult) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *SpanQueryAggregateResult) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *SpanQueryAggregateResult) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetSeries

`func (o *SpanQueryAggregateResult) GetSeries() []SpanQueryAggregateDataSeries`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *SpanQueryAggregateResult) GetSeriesOk() (*[]SpanQueryAggregateDataSeries, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *SpanQueryAggregateResult) SetSeries(v []SpanQueryAggregateDataSeries)`

SetSeries sets Series field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


