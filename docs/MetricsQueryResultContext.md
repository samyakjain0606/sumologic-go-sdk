# MetricsQueryResultContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QuantizationGranularity** | Pointer to **int64** | Quantization granularity. Size of the quantization bucket/quant in milliseconds. | [optional] 
**Rollup** | Pointer to **string** | We use the term rollup to refer to the aggregation function Sumo Logic uses when quantizing metrics. Can be &#x60;Avg&#x60;, &#x60;Sum&#x60;, &#x60;Min&#x60;, &#x60;Max&#x60;, &#x60;Count&#x60; or &#x60;Rate&#x60;. | [optional] 
**ActualQueryTimeRange** | Pointer to [**Iso8601TimeRange**](Iso8601TimeRange.md) |  | [optional] 

## Methods

### NewMetricsQueryResultContext

`func NewMetricsQueryResultContext() *MetricsQueryResultContext`

NewMetricsQueryResultContext instantiates a new MetricsQueryResultContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsQueryResultContextWithDefaults

`func NewMetricsQueryResultContextWithDefaults() *MetricsQueryResultContext`

NewMetricsQueryResultContextWithDefaults instantiates a new MetricsQueryResultContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuantizationGranularity

`func (o *MetricsQueryResultContext) GetQuantizationGranularity() int64`

GetQuantizationGranularity returns the QuantizationGranularity field if non-nil, zero value otherwise.

### GetQuantizationGranularityOk

`func (o *MetricsQueryResultContext) GetQuantizationGranularityOk() (*int64, bool)`

GetQuantizationGranularityOk returns a tuple with the QuantizationGranularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantizationGranularity

`func (o *MetricsQueryResultContext) SetQuantizationGranularity(v int64)`

SetQuantizationGranularity sets QuantizationGranularity field to given value.

### HasQuantizationGranularity

`func (o *MetricsQueryResultContext) HasQuantizationGranularity() bool`

HasQuantizationGranularity returns a boolean if a field has been set.

### GetRollup

`func (o *MetricsQueryResultContext) GetRollup() string`

GetRollup returns the Rollup field if non-nil, zero value otherwise.

### GetRollupOk

`func (o *MetricsQueryResultContext) GetRollupOk() (*string, bool)`

GetRollupOk returns a tuple with the Rollup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollup

`func (o *MetricsQueryResultContext) SetRollup(v string)`

SetRollup sets Rollup field to given value.

### HasRollup

`func (o *MetricsQueryResultContext) HasRollup() bool`

HasRollup returns a boolean if a field has been set.

### GetActualQueryTimeRange

`func (o *MetricsQueryResultContext) GetActualQueryTimeRange() Iso8601TimeRange`

GetActualQueryTimeRange returns the ActualQueryTimeRange field if non-nil, zero value otherwise.

### GetActualQueryTimeRangeOk

`func (o *MetricsQueryResultContext) GetActualQueryTimeRangeOk() (*Iso8601TimeRange, bool)`

GetActualQueryTimeRangeOk returns a tuple with the ActualQueryTimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualQueryTimeRange

`func (o *MetricsQueryResultContext) SetActualQueryTimeRange(v Iso8601TimeRange)`

SetActualQueryTimeRange sets ActualQueryTimeRange field to given value.

### HasActualQueryTimeRange

`func (o *MetricsQueryResultContext) HasActualQueryTimeRange() bool`

HasActualQueryTimeRange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


