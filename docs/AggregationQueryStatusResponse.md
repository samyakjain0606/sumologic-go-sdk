# AggregationQueryStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryRows** | [**[]AggregationQueryRowStatus**](AggregationQueryRowStatus.md) | A list of statuses on a per query row basis. | 
**Status** | **string** | Status of the query. Possible values: &#x60;Processing&#x60;, &#x60;Finished&#x60;, &#x60;Error&#x60;, &#x60;Canceled&#x60;. | 
**StatusMessage** | Pointer to **string** | Descriptive message of the status. | [optional] 

## Methods

### NewAggregationQueryStatusResponse

`func NewAggregationQueryStatusResponse(queryRows []AggregationQueryRowStatus, status string, ) *AggregationQueryStatusResponse`

NewAggregationQueryStatusResponse instantiates a new AggregationQueryStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregationQueryStatusResponseWithDefaults

`func NewAggregationQueryStatusResponseWithDefaults() *AggregationQueryStatusResponse`

NewAggregationQueryStatusResponseWithDefaults instantiates a new AggregationQueryStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueryRows

`func (o *AggregationQueryStatusResponse) GetQueryRows() []AggregationQueryRowStatus`

GetQueryRows returns the QueryRows field if non-nil, zero value otherwise.

### GetQueryRowsOk

`func (o *AggregationQueryStatusResponse) GetQueryRowsOk() (*[]AggregationQueryRowStatus, bool)`

GetQueryRowsOk returns a tuple with the QueryRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryRows

`func (o *AggregationQueryStatusResponse) SetQueryRows(v []AggregationQueryRowStatus)`

SetQueryRows sets QueryRows field to given value.


### GetStatus

`func (o *AggregationQueryStatusResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AggregationQueryStatusResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AggregationQueryStatusResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStatusMessage

`func (o *AggregationQueryStatusResponse) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *AggregationQueryStatusResponse) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *AggregationQueryStatusResponse) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *AggregationQueryStatusResponse) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


