# Sli

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EvaluationType** | **string** | Evaluate SLI using successful/total windows, or occurrence of successful events over entire compliance period, or based on monitor evaluation. | 

## Methods

### NewSli

`func NewSli(evaluationType string, ) *Sli`

NewSli instantiates a new Sli object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSliWithDefaults

`func NewSliWithDefaults() *Sli`

NewSliWithDefaults instantiates a new Sli object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvaluationType

`func (o *Sli) GetEvaluationType() string`

GetEvaluationType returns the EvaluationType field if non-nil, zero value otherwise.

### GetEvaluationTypeOk

`func (o *Sli) GetEvaluationTypeOk() (*string, bool)`

GetEvaluationTypeOk returns a tuple with the EvaluationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationType

`func (o *Sli) SetEvaluationType(v string)`

SetEvaluationType sets EvaluationType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


