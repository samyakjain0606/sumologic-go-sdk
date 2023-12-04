# Window

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryType** | **string** | Type of Raw Data Queries for SLI (Logs/Metrics). | 
**Queries** | [**[]SliQueryGroup**](SliQueryGroup.md) | Queries for defining SLI. | 
**Threshold** | **float32** | Threshold for classifying window as successful or unsuccessful. | 
**Op** | **string** | Comparison function with window threshold (LessThan/GreaterThan/LessThanOrEqual/GreaterThanOrEqual). | 
**Aggregation** | Pointer to **string** | Aggregation function applied over each window to arrive at SLI. Must be &#x60;Avg&#x60;, &#x60;Min&#x60;, &#x60;Max&#x60;, &#x60;Sum&#x60;, or percentile of the form &#x60;pX&#x60; where &#x60;X&#x60; is an integer between 0 and 100. | [optional] 
**Size** | **string** | Size of the aggregation window (minimum of 1m and maximum of 1h). | 

## Methods

### NewWindow

`func NewWindow(queryType string, queries []SliQueryGroup, threshold float32, op string, size string, ) *Window`

NewWindow instantiates a new Window object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWindowWithDefaults

`func NewWindowWithDefaults() *Window`

NewWindowWithDefaults instantiates a new Window object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueryType

`func (o *Window) GetQueryType() string`

GetQueryType returns the QueryType field if non-nil, zero value otherwise.

### GetQueryTypeOk

`func (o *Window) GetQueryTypeOk() (*string, bool)`

GetQueryTypeOk returns a tuple with the QueryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryType

`func (o *Window) SetQueryType(v string)`

SetQueryType sets QueryType field to given value.


### GetQueries

`func (o *Window) GetQueries() []SliQueryGroup`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *Window) GetQueriesOk() (*[]SliQueryGroup, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *Window) SetQueries(v []SliQueryGroup)`

SetQueries sets Queries field to given value.


### GetThreshold

`func (o *Window) GetThreshold() float32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *Window) GetThresholdOk() (*float32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *Window) SetThreshold(v float32)`

SetThreshold sets Threshold field to given value.


### GetOp

`func (o *Window) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *Window) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *Window) SetOp(v string)`

SetOp sets Op field to given value.


### GetAggregation

`func (o *Window) GetAggregation() string`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *Window) GetAggregationOk() (*string, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *Window) SetAggregation(v string)`

SetAggregation sets Aggregation field to given value.

### HasAggregation

`func (o *Window) HasAggregation() bool`

HasAggregation returns a boolean if a field has been set.

### GetSize

`func (o *Window) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Window) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Window) SetSize(v string)`

SetSize sets Size field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


