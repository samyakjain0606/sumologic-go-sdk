# NoTraceFieldValuesReason

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | A code uniquely identifying the reason for the lack of trace field values. Possible values: &#x60;HighCardinalityField&#x60;, &#x60;AutocompleteDisabled&#x60;. | 
**Message** | **string** | A short English-language description of the reason. | 

## Methods

### NewNoTraceFieldValuesReason

`func NewNoTraceFieldValuesReason(code string, message string, ) *NoTraceFieldValuesReason`

NewNoTraceFieldValuesReason instantiates a new NoTraceFieldValuesReason object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNoTraceFieldValuesReasonWithDefaults

`func NewNoTraceFieldValuesReasonWithDefaults() *NoTraceFieldValuesReason`

NewNoTraceFieldValuesReasonWithDefaults instantiates a new NoTraceFieldValuesReason object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *NoTraceFieldValuesReason) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *NoTraceFieldValuesReason) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *NoTraceFieldValuesReason) SetCode(v string)`

SetCode sets Code field to given value.


### GetMessage

`func (o *NoTraceFieldValuesReason) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *NoTraceFieldValuesReason) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *NoTraceFieldValuesReason) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


