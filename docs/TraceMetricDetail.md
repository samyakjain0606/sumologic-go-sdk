# TraceMetricDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metric** | **string** | Trace metric name. In trace queries it can be used in &#x60;MetricTracingFilter.metric&#x60;. | 
**Description** | Pointer to **string** | Short description of the metric. | [optional] 
**Type** | **string** | The type the values of this field will have. Possible values: &#x60;DoubleTracingValue&#x60;, &#x60;IntegerTracingValue&#x60;. | 

## Methods

### NewTraceMetricDetail

`func NewTraceMetricDetail(metric string, type_ string, ) *TraceMetricDetail`

NewTraceMetricDetail instantiates a new TraceMetricDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTraceMetricDetailWithDefaults

`func NewTraceMetricDetailWithDefaults() *TraceMetricDetail`

NewTraceMetricDetailWithDefaults instantiates a new TraceMetricDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetric

`func (o *TraceMetricDetail) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *TraceMetricDetail) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *TraceMetricDetail) SetMetric(v string)`

SetMetric sets Metric field to given value.


### GetDescription

`func (o *TraceMetricDetail) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TraceMetricDetail) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TraceMetricDetail) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TraceMetricDetail) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *TraceMetricDetail) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TraceMetricDetail) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TraceMetricDetail) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


