# AggregationQueryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryRows** | [**[]AggregationQueryRow**](AggregationQueryRow.md) | A list of tracing queries. | 
**TimeRange** | [**ResolvableTimeRange**](ResolvableTimeRange.md) |  | 
**XAxisGroupByAttribute** | [**AggregationGroupByAttribute**](AggregationGroupByAttribute.md) |  | 

## Methods

### NewAggregationQueryRequest

`func NewAggregationQueryRequest(queryRows []AggregationQueryRow, timeRange ResolvableTimeRange, xAxisGroupByAttribute AggregationGroupByAttribute, ) *AggregationQueryRequest`

NewAggregationQueryRequest instantiates a new AggregationQueryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregationQueryRequestWithDefaults

`func NewAggregationQueryRequestWithDefaults() *AggregationQueryRequest`

NewAggregationQueryRequestWithDefaults instantiates a new AggregationQueryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueryRows

`func (o *AggregationQueryRequest) GetQueryRows() []AggregationQueryRow`

GetQueryRows returns the QueryRows field if non-nil, zero value otherwise.

### GetQueryRowsOk

`func (o *AggregationQueryRequest) GetQueryRowsOk() (*[]AggregationQueryRow, bool)`

GetQueryRowsOk returns a tuple with the QueryRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryRows

`func (o *AggregationQueryRequest) SetQueryRows(v []AggregationQueryRow)`

SetQueryRows sets QueryRows field to given value.


### GetTimeRange

`func (o *AggregationQueryRequest) GetTimeRange() ResolvableTimeRange`

GetTimeRange returns the TimeRange field if non-nil, zero value otherwise.

### GetTimeRangeOk

`func (o *AggregationQueryRequest) GetTimeRangeOk() (*ResolvableTimeRange, bool)`

GetTimeRangeOk returns a tuple with the TimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRange

`func (o *AggregationQueryRequest) SetTimeRange(v ResolvableTimeRange)`

SetTimeRange sets TimeRange field to given value.


### GetXAxisGroupByAttribute

`func (o *AggregationQueryRequest) GetXAxisGroupByAttribute() AggregationGroupByAttribute`

GetXAxisGroupByAttribute returns the XAxisGroupByAttribute field if non-nil, zero value otherwise.

### GetXAxisGroupByAttributeOk

`func (o *AggregationQueryRequest) GetXAxisGroupByAttributeOk() (*AggregationGroupByAttribute, bool)`

GetXAxisGroupByAttributeOk returns a tuple with the XAxisGroupByAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXAxisGroupByAttribute

`func (o *AggregationQueryRequest) SetXAxisGroupByAttribute(v AggregationGroupByAttribute)`

SetXAxisGroupByAttribute sets XAxisGroupByAttribute field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


