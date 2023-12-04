# CpcQueryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryRows** | [**[]CpcQueryRow**](CpcQueryRow.md) | A list of cpc queries. | 
**TimeRange** | [**ResolvableTimeRange**](ResolvableTimeRange.md) |  | 

## Methods

### NewCpcQueryRequest

`func NewCpcQueryRequest(queryRows []CpcQueryRow, timeRange ResolvableTimeRange, ) *CpcQueryRequest`

NewCpcQueryRequest instantiates a new CpcQueryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCpcQueryRequestWithDefaults

`func NewCpcQueryRequestWithDefaults() *CpcQueryRequest`

NewCpcQueryRequestWithDefaults instantiates a new CpcQueryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueryRows

`func (o *CpcQueryRequest) GetQueryRows() []CpcQueryRow`

GetQueryRows returns the QueryRows field if non-nil, zero value otherwise.

### GetQueryRowsOk

`func (o *CpcQueryRequest) GetQueryRowsOk() (*[]CpcQueryRow, bool)`

GetQueryRowsOk returns a tuple with the QueryRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryRows

`func (o *CpcQueryRequest) SetQueryRows(v []CpcQueryRow)`

SetQueryRows sets QueryRows field to given value.


### GetTimeRange

`func (o *CpcQueryRequest) GetTimeRange() ResolvableTimeRange`

GetTimeRange returns the TimeRange field if non-nil, zero value otherwise.

### GetTimeRangeOk

`func (o *CpcQueryRequest) GetTimeRangeOk() (*ResolvableTimeRange, bool)`

GetTimeRangeOk returns a tuple with the TimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRange

`func (o *CpcQueryRequest) SetTimeRange(v ResolvableTimeRange)`

SetTimeRange sets TimeRange field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


