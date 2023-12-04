# FieldTracingFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** | The field name to filter by. The list of supported field names can be retrieved using the [Trace Query Fields](#operation/getTraceQueryFields) endpoint. | 
**Operator** | **string** | The operator to use. Accepted values:   &lt;table&gt;     &lt;tr&gt;       &lt;th&gt;Operator&lt;/th&gt;       &lt;th&gt;Accepted value types&lt;/th&gt;     &lt;/tr&gt;     &lt;tr&gt;       &lt;th&gt;&amp;lt; &amp;lt;&#x3D; &amp;gt; &amp;gt;&#x3D; &#x3D; !&#x3D;&lt;/th&gt;       &lt;th&gt;StringTracingValue DoubleTracingValue IntegerTracingValue DateTimeTracingValue&lt;/th&gt;     &lt;/tr&gt;     &lt;tr&gt;       &lt;th&gt;in&lt;/th&gt;       &lt;th&gt;ArrayTracingValue of StringTracingValue / DoubleTracingValue / IntegerTracingValue / DateTimeTracingValue&lt;/th&gt;     &lt;/tr&gt;     &lt;tr&gt;       &lt;th&gt;between&lt;/th&gt;       &lt;th&gt;RangeTracingValue of StringTracingValue / DoubleTracingValue / IntegerTracingValue / DateTimeTracingValue&lt;/th&gt;     &lt;/tr&gt;   &lt;/table&gt; | 
**Value** | Pointer to [**TracingValue**](TracingValue.md) |  | [optional] 

## Methods

### NewFieldTracingFilter

`func NewFieldTracingFilter(field string, operator string, ) *FieldTracingFilter`

NewFieldTracingFilter instantiates a new FieldTracingFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldTracingFilterWithDefaults

`func NewFieldTracingFilterWithDefaults() *FieldTracingFilter`

NewFieldTracingFilterWithDefaults instantiates a new FieldTracingFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *FieldTracingFilter) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *FieldTracingFilter) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *FieldTracingFilter) SetField(v string)`

SetField sets Field field to given value.


### GetOperator

`func (o *FieldTracingFilter) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *FieldTracingFilter) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *FieldTracingFilter) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetValue

`func (o *FieldTracingFilter) GetValue() TracingValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FieldTracingFilter) GetValueOk() (*TracingValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FieldTracingFilter) SetValue(v TracingValue)`

SetValue sets Value field to given value.

### HasValue

`func (o *FieldTracingFilter) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


