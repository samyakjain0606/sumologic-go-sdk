# MetricTracingFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metric** | **string** | The name of the metric to filter by. The list of supported metrics can be retrieved using the [Trace Metrics](#operation/getMetrics) endpoint. | 
**Operator** | **string** | The operator to use. Accepted values:   &lt;table&gt;     &lt;tr&gt;       &lt;th&gt;Operator&lt;/th&gt;       &lt;th&gt;Accepted value types&lt;/th&gt;     &lt;/tr&gt;     &lt;tr&gt;       &lt;th&gt;&amp;lt; &amp;lt;&#x3D; &amp;gt; &amp;gt;&#x3D; &#x3D;&lt;/th&gt;       &lt;th&gt;DoubleTracingValue IntegerTracingValue&lt;/th&gt;     &lt;/tr&gt;     &lt;tr&gt;       &lt;th&gt;between&lt;/th&gt;       &lt;th&gt;RangeTracingValue of DoubleTracingValue / IntegerTracingValue&lt;/th&gt;     &lt;/tr&gt;   &lt;/table&gt; | 
**Value** | Pointer to [**TracingValue**](TracingValue.md) |  | [optional] 

## Methods

### NewMetricTracingFilter

`func NewMetricTracingFilter(metric string, operator string, ) *MetricTracingFilter`

NewMetricTracingFilter instantiates a new MetricTracingFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricTracingFilterWithDefaults

`func NewMetricTracingFilterWithDefaults() *MetricTracingFilter`

NewMetricTracingFilterWithDefaults instantiates a new MetricTracingFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetric

`func (o *MetricTracingFilter) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *MetricTracingFilter) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *MetricTracingFilter) SetMetric(v string)`

SetMetric sets Metric field to given value.


### GetOperator

`func (o *MetricTracingFilter) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *MetricTracingFilter) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *MetricTracingFilter) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetValue

`func (o *MetricTracingFilter) GetValue() TracingValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MetricTracingFilter) GetValueOk() (*TracingValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MetricTracingFilter) SetValue(v TracingValue)`

SetValue sets Value field to given value.

### HasValue

`func (o *MetricTracingFilter) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


