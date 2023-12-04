# SpansFilterKeyValuePair

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | **string** | A symbol that indicates an operation to be performed between a &#x60;fieldName&#x60; and &#x60;fieldValue&#x60;. | 
**FieldValue** | **string** | The second argument of the operation applied to a &#x60;fieldName&#x60;. | 

## Methods

### NewSpansFilterKeyValuePair

`func NewSpansFilterKeyValuePair(operator string, fieldValue string, ) *SpansFilterKeyValuePair`

NewSpansFilterKeyValuePair instantiates a new SpansFilterKeyValuePair object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpansFilterKeyValuePairWithDefaults

`func NewSpansFilterKeyValuePairWithDefaults() *SpansFilterKeyValuePair`

NewSpansFilterKeyValuePairWithDefaults instantiates a new SpansFilterKeyValuePair object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *SpansFilterKeyValuePair) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *SpansFilterKeyValuePair) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *SpansFilterKeyValuePair) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetFieldValue

`func (o *SpansFilterKeyValuePair) GetFieldValue() string`

GetFieldValue returns the FieldValue field if non-nil, zero value otherwise.

### GetFieldValueOk

`func (o *SpansFilterKeyValuePair) GetFieldValueOk() (*string, bool)`

GetFieldValueOk returns a tuple with the FieldValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldValue

`func (o *SpansFilterKeyValuePair) SetFieldValue(v string)`

SetFieldValue sets FieldValue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


