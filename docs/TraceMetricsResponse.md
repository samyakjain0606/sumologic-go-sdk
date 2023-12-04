# TraceMetricsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metrics** | [**[]TraceMetricDetail**](TraceMetricDetail.md) | List of trace metrics. | 

## Methods

### NewTraceMetricsResponse

`func NewTraceMetricsResponse(metrics []TraceMetricDetail, ) *TraceMetricsResponse`

NewTraceMetricsResponse instantiates a new TraceMetricsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTraceMetricsResponseWithDefaults

`func NewTraceMetricsResponseWithDefaults() *TraceMetricsResponse`

NewTraceMetricsResponseWithDefaults instantiates a new TraceMetricsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetrics

`func (o *TraceMetricsResponse) GetMetrics() []TraceMetricDetail`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *TraceMetricsResponse) GetMetricsOk() (*[]TraceMetricDetail, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *TraceMetricsResponse) SetMetrics(v []TraceMetricDetail)`

SetMetrics sets Metrics field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


