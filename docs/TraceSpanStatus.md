# TraceSpanStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Status code of the span. Possible values: &#x60;OK&#x60;, &#x60;ERROR&#x60;, &#x60;UNKNOWN&#x60;. | 
**Message** | Pointer to **string** | Optional descriptive message about the status, could be an http status code or the kind of an error, e.g. OSError. | [optional] 

## Methods

### NewTraceSpanStatus

`func NewTraceSpanStatus(code string, ) *TraceSpanStatus`

NewTraceSpanStatus instantiates a new TraceSpanStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTraceSpanStatusWithDefaults

`func NewTraceSpanStatusWithDefaults() *TraceSpanStatus`

NewTraceSpanStatusWithDefaults instantiates a new TraceSpanStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *TraceSpanStatus) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *TraceSpanStatus) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *TraceSpanStatus) SetCode(v string)`

SetCode sets Code field to given value.


### GetMessage

`func (o *TraceSpanStatus) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TraceSpanStatus) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TraceSpanStatus) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *TraceSpanStatus) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


