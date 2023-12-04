# SpansCalculationVisualization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** | A field by which the spans are aggregated. | 
**Aggregator** | [**SpanCalculationAggregator**](SpanCalculationAggregator.md) |  | 

## Methods

### NewSpansCalculationVisualization

`func NewSpansCalculationVisualization(field string, aggregator SpanCalculationAggregator, ) *SpansCalculationVisualization`

NewSpansCalculationVisualization instantiates a new SpansCalculationVisualization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpansCalculationVisualizationWithDefaults

`func NewSpansCalculationVisualizationWithDefaults() *SpansCalculationVisualization`

NewSpansCalculationVisualizationWithDefaults instantiates a new SpansCalculationVisualization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *SpansCalculationVisualization) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *SpansCalculationVisualization) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *SpansCalculationVisualization) SetField(v string)`

SetField sets Field field to given value.


### GetAggregator

`func (o *SpansCalculationVisualization) GetAggregator() SpanCalculationAggregator`

GetAggregator returns the Aggregator field if non-nil, zero value otherwise.

### GetAggregatorOk

`func (o *SpansCalculationVisualization) GetAggregatorOk() (*SpanCalculationAggregator, bool)`

GetAggregatorOk returns a tuple with the Aggregator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregator

`func (o *SpansCalculationVisualization) SetAggregator(v SpanCalculationAggregator)`

SetAggregator sets Aggregator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


