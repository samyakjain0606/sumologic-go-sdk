# AsyncTraceQueryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryRows** | [**[]AsyncTraceQueryRow**](AsyncTraceQueryRow.md) | A list of trace queries. | 
**TimeRange** | [**ResolvableTimeRange**](ResolvableTimeRange.md) |  | 

## Methods

### NewAsyncTraceQueryRequest

`func NewAsyncTraceQueryRequest(queryRows []AsyncTraceQueryRow, timeRange ResolvableTimeRange, ) *AsyncTraceQueryRequest`

NewAsyncTraceQueryRequest instantiates a new AsyncTraceQueryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAsyncTraceQueryRequestWithDefaults

`func NewAsyncTraceQueryRequestWithDefaults() *AsyncTraceQueryRequest`

NewAsyncTraceQueryRequestWithDefaults instantiates a new AsyncTraceQueryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueryRows

`func (o *AsyncTraceQueryRequest) GetQueryRows() []AsyncTraceQueryRow`

GetQueryRows returns the QueryRows field if non-nil, zero value otherwise.

### GetQueryRowsOk

`func (o *AsyncTraceQueryRequest) GetQueryRowsOk() (*[]AsyncTraceQueryRow, bool)`

GetQueryRowsOk returns a tuple with the QueryRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryRows

`func (o *AsyncTraceQueryRequest) SetQueryRows(v []AsyncTraceQueryRow)`

SetQueryRows sets QueryRows field to given value.


### GetTimeRange

`func (o *AsyncTraceQueryRequest) GetTimeRange() ResolvableTimeRange`

GetTimeRange returns the TimeRange field if non-nil, zero value otherwise.

### GetTimeRangeOk

`func (o *AsyncTraceQueryRequest) GetTimeRangeOk() (*ResolvableTimeRange, bool)`

GetTimeRangeOk returns a tuple with the TimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRange

`func (o *AsyncTraceQueryRequest) SetTimeRange(v ResolvableTimeRange)`

SetTimeRange sets TimeRange field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


