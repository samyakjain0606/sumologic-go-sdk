# SpanQueryRowError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | The error code. | 
**Message** | **string** | Short description of the occured error. | 
**Details** | Pointer to **string** | Details about the occured error. | [optional] 

## Methods

### NewSpanQueryRowError

`func NewSpanQueryRowError(code string, message string, ) *SpanQueryRowError`

NewSpanQueryRowError instantiates a new SpanQueryRowError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpanQueryRowErrorWithDefaults

`func NewSpanQueryRowErrorWithDefaults() *SpanQueryRowError`

NewSpanQueryRowErrorWithDefaults instantiates a new SpanQueryRowError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *SpanQueryRowError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *SpanQueryRowError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *SpanQueryRowError) SetCode(v string)`

SetCode sets Code field to given value.


### GetMessage

`func (o *SpanQueryRowError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SpanQueryRowError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SpanQueryRowError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetDetails

`func (o *SpanQueryRowError) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *SpanQueryRowError) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *SpanQueryRowError) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *SpanQueryRowError) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


