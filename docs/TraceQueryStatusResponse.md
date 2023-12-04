# TraceQueryStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryRows** | [**[]TraceQueryRowStatus**](TraceQueryRowStatus.md) | A list of trace queries. | 
**Status** | **string** | Status of the query. Possible values: &#x60;Processing&#x60;, &#x60;Finished&#x60;, &#x60;Error&#x60;, &#x60;Canceled&#x60;. | 

## Methods

### NewTraceQueryStatusResponse

`func NewTraceQueryStatusResponse(queryRows []TraceQueryRowStatus, status string, ) *TraceQueryStatusResponse`

NewTraceQueryStatusResponse instantiates a new TraceQueryStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTraceQueryStatusResponseWithDefaults

`func NewTraceQueryStatusResponseWithDefaults() *TraceQueryStatusResponse`

NewTraceQueryStatusResponseWithDefaults instantiates a new TraceQueryStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueryRows

`func (o *TraceQueryStatusResponse) GetQueryRows() []TraceQueryRowStatus`

GetQueryRows returns the QueryRows field if non-nil, zero value otherwise.

### GetQueryRowsOk

`func (o *TraceQueryStatusResponse) GetQueryRowsOk() (*[]TraceQueryRowStatus, bool)`

GetQueryRowsOk returns a tuple with the QueryRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryRows

`func (o *TraceQueryStatusResponse) SetQueryRows(v []TraceQueryRowStatus)`

SetQueryRows sets QueryRows field to given value.


### GetStatus

`func (o *TraceQueryStatusResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TraceQueryStatusResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TraceQueryStatusResponse) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


