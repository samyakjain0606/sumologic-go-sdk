# Query

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryString** | **string** | The metrics, traces or logs query. | 
**QueryType** | **string** | The type of the query, either &#x60;Metrics&#x60;, &#x60;Traces&#x60;, &#x60;Spans&#x60; or &#x60;Logs&#x60;. | 
**QueryKey** | **string** | The key for metric, traces or log queries. Used as an identifier for queries. It is displayed on the panel builder and used for display overrides and query toggling.  | 
**MetricsQueryMode** | Pointer to **string** | The mode of the metrics query that the user was editing. Can be &#x60;Basic&#x60; or &#x60;Advanced&#x60;. Will ONLY be specified for metrics queries.  | [optional] 
**MetricsQueryData** | Pointer to [**MetricsQueryData**](MetricsQueryData.md) |  | [optional] 
**TracesQueryData** | Pointer to [**TracesQueryData**](TracesQueryData.md) |  | [optional] 
**SpansQueryData** | Pointer to [**SpansQueryData**](SpansQueryData.md) |  | [optional] 
**ParseMode** | Pointer to **string** | This field only applies for queryType of &#x60;Logs&#x60; but other query types may be supported in the future. Define the parsing mode to scan the JSON format log messages. Possible values are:   1. &#x60;Auto&#x60;   2. &#x60;Manual&#x60; In AutoParse mode, the system automatically figures out fields to parse based on the search query. While in the Manual mode, no fields are parsed out automatically. For more information see [Dynamic Parsing](https://help.sumologic.com/?cid&#x3D;0011). | [optional] [default to "Auto"]
**TimeSource** | Pointer to **string** | This field only applies for queryType of &#x60;Logs&#x60; but other query types may be supported in the future. Define the time source of this query. Possible values are &#x60;Message&#x60; and &#x60;Receipt&#x60;. &#x60;Message&#x60; will use the timeStamp on the message, while &#x60;Receipt&#x60; will use the timestamp it was received by Sumo. | [optional] [default to "Message"]
**Transient** | Pointer to **bool** | This field only applies for queryType of &#x60;Metrics&#x60; but other query types may be supported in the future. Determines if the row should be returned in the response. Can be used in conjunction with a join, if only the result of the join is needed, and not the intermediate rows. Setting &#x60;transient&#x60; to &#x60;true&#x60;  wherever the intermediate results aren&#39;t required speeds up the computation and reduces the amount of data  transferred over the network. | [optional] [default to false]
**OutputCardinalityLimit** | Pointer to **int32** | This field only applies for queryType of &#x60;Metrics&#x60; but other query types may be supported in the future. Specifies the output cardinality limitations for the query, which is the maximum number of timeseries returned in the result. | [optional] [default to 1000]

## Methods

### NewQuery

`func NewQuery(queryString string, queryType string, queryKey string, ) *Query`

NewQuery instantiates a new Query object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryWithDefaults

`func NewQueryWithDefaults() *Query`

NewQueryWithDefaults instantiates a new Query object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueryString

`func (o *Query) GetQueryString() string`

GetQueryString returns the QueryString field if non-nil, zero value otherwise.

### GetQueryStringOk

`func (o *Query) GetQueryStringOk() (*string, bool)`

GetQueryStringOk returns a tuple with the QueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryString

`func (o *Query) SetQueryString(v string)`

SetQueryString sets QueryString field to given value.


### GetQueryType

`func (o *Query) GetQueryType() string`

GetQueryType returns the QueryType field if non-nil, zero value otherwise.

### GetQueryTypeOk

`func (o *Query) GetQueryTypeOk() (*string, bool)`

GetQueryTypeOk returns a tuple with the QueryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryType

`func (o *Query) SetQueryType(v string)`

SetQueryType sets QueryType field to given value.


### GetQueryKey

`func (o *Query) GetQueryKey() string`

GetQueryKey returns the QueryKey field if non-nil, zero value otherwise.

### GetQueryKeyOk

`func (o *Query) GetQueryKeyOk() (*string, bool)`

GetQueryKeyOk returns a tuple with the QueryKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryKey

`func (o *Query) SetQueryKey(v string)`

SetQueryKey sets QueryKey field to given value.


### GetMetricsQueryMode

`func (o *Query) GetMetricsQueryMode() string`

GetMetricsQueryMode returns the MetricsQueryMode field if non-nil, zero value otherwise.

### GetMetricsQueryModeOk

`func (o *Query) GetMetricsQueryModeOk() (*string, bool)`

GetMetricsQueryModeOk returns a tuple with the MetricsQueryMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsQueryMode

`func (o *Query) SetMetricsQueryMode(v string)`

SetMetricsQueryMode sets MetricsQueryMode field to given value.

### HasMetricsQueryMode

`func (o *Query) HasMetricsQueryMode() bool`

HasMetricsQueryMode returns a boolean if a field has been set.

### GetMetricsQueryData

`func (o *Query) GetMetricsQueryData() MetricsQueryData`

GetMetricsQueryData returns the MetricsQueryData field if non-nil, zero value otherwise.

### GetMetricsQueryDataOk

`func (o *Query) GetMetricsQueryDataOk() (*MetricsQueryData, bool)`

GetMetricsQueryDataOk returns a tuple with the MetricsQueryData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsQueryData

`func (o *Query) SetMetricsQueryData(v MetricsQueryData)`

SetMetricsQueryData sets MetricsQueryData field to given value.

### HasMetricsQueryData

`func (o *Query) HasMetricsQueryData() bool`

HasMetricsQueryData returns a boolean if a field has been set.

### GetTracesQueryData

`func (o *Query) GetTracesQueryData() TracesQueryData`

GetTracesQueryData returns the TracesQueryData field if non-nil, zero value otherwise.

### GetTracesQueryDataOk

`func (o *Query) GetTracesQueryDataOk() (*TracesQueryData, bool)`

GetTracesQueryDataOk returns a tuple with the TracesQueryData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTracesQueryData

`func (o *Query) SetTracesQueryData(v TracesQueryData)`

SetTracesQueryData sets TracesQueryData field to given value.

### HasTracesQueryData

`func (o *Query) HasTracesQueryData() bool`

HasTracesQueryData returns a boolean if a field has been set.

### GetSpansQueryData

`func (o *Query) GetSpansQueryData() SpansQueryData`

GetSpansQueryData returns the SpansQueryData field if non-nil, zero value otherwise.

### GetSpansQueryDataOk

`func (o *Query) GetSpansQueryDataOk() (*SpansQueryData, bool)`

GetSpansQueryDataOk returns a tuple with the SpansQueryData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpansQueryData

`func (o *Query) SetSpansQueryData(v SpansQueryData)`

SetSpansQueryData sets SpansQueryData field to given value.

### HasSpansQueryData

`func (o *Query) HasSpansQueryData() bool`

HasSpansQueryData returns a boolean if a field has been set.

### GetParseMode

`func (o *Query) GetParseMode() string`

GetParseMode returns the ParseMode field if non-nil, zero value otherwise.

### GetParseModeOk

`func (o *Query) GetParseModeOk() (*string, bool)`

GetParseModeOk returns a tuple with the ParseMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParseMode

`func (o *Query) SetParseMode(v string)`

SetParseMode sets ParseMode field to given value.

### HasParseMode

`func (o *Query) HasParseMode() bool`

HasParseMode returns a boolean if a field has been set.

### GetTimeSource

`func (o *Query) GetTimeSource() string`

GetTimeSource returns the TimeSource field if non-nil, zero value otherwise.

### GetTimeSourceOk

`func (o *Query) GetTimeSourceOk() (*string, bool)`

GetTimeSourceOk returns a tuple with the TimeSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSource

`func (o *Query) SetTimeSource(v string)`

SetTimeSource sets TimeSource field to given value.

### HasTimeSource

`func (o *Query) HasTimeSource() bool`

HasTimeSource returns a boolean if a field has been set.

### GetTransient

`func (o *Query) GetTransient() bool`

GetTransient returns the Transient field if non-nil, zero value otherwise.

### GetTransientOk

`func (o *Query) GetTransientOk() (*bool, bool)`

GetTransientOk returns a tuple with the Transient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransient

`func (o *Query) SetTransient(v bool)`

SetTransient sets Transient field to given value.

### HasTransient

`func (o *Query) HasTransient() bool`

HasTransient returns a boolean if a field has been set.

### GetOutputCardinalityLimit

`func (o *Query) GetOutputCardinalityLimit() int32`

GetOutputCardinalityLimit returns the OutputCardinalityLimit field if non-nil, zero value otherwise.

### GetOutputCardinalityLimitOk

`func (o *Query) GetOutputCardinalityLimitOk() (*int32, bool)`

GetOutputCardinalityLimitOk returns a tuple with the OutputCardinalityLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputCardinalityLimit

`func (o *Query) SetOutputCardinalityLimit(v int32)`

SetOutputCardinalityLimit sets OutputCardinalityLimit field to given value.

### HasOutputCardinalityLimit

`func (o *Query) HasOutputCardinalityLimit() bool`

HasOutputCardinalityLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


