# SpanQueryStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryRows** | [**[]SpanQueryRowStatus**](SpanQueryRowStatus.md) | A list of span analytics queries. | 
**Status** | **string** | Status of the query. Possible values: &#x60;Processing&#x60;, &#x60;Finished&#x60;, &#x60;Error&#x60;, &#x60;Paused&#x60; | 

## Methods

### NewSpanQueryStatusResponse

`func NewSpanQueryStatusResponse(queryRows []SpanQueryRowStatus, status string, ) *SpanQueryStatusResponse`

NewSpanQueryStatusResponse instantiates a new SpanQueryStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpanQueryStatusResponseWithDefaults

`func NewSpanQueryStatusResponseWithDefaults() *SpanQueryStatusResponse`

NewSpanQueryStatusResponseWithDefaults instantiates a new SpanQueryStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueryRows

`func (o *SpanQueryStatusResponse) GetQueryRows() []SpanQueryRowStatus`

GetQueryRows returns the QueryRows field if non-nil, zero value otherwise.

### GetQueryRowsOk

`func (o *SpanQueryStatusResponse) GetQueryRowsOk() (*[]SpanQueryRowStatus, bool)`

GetQueryRowsOk returns a tuple with the QueryRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryRows

`func (o *SpanQueryStatusResponse) SetQueryRows(v []SpanQueryRowStatus)`

SetQueryRows sets QueryRows field to given value.


### GetStatus

`func (o *SpanQueryStatusResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SpanQueryStatusResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SpanQueryStatusResponse) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


