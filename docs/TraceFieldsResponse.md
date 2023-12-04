# TraceFieldsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fields** | [**[]TraceFieldDetail**](TraceFieldDetail.md) | List of filter fields. | 

## Methods

### NewTraceFieldsResponse

`func NewTraceFieldsResponse(fields []TraceFieldDetail, ) *TraceFieldsResponse`

NewTraceFieldsResponse instantiates a new TraceFieldsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTraceFieldsResponseWithDefaults

`func NewTraceFieldsResponseWithDefaults() *TraceFieldsResponse`

NewTraceFieldsResponseWithDefaults instantiates a new TraceFieldsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFields

`func (o *TraceFieldsResponse) GetFields() []TraceFieldDetail`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *TraceFieldsResponse) GetFieldsOk() (*[]TraceFieldDetail, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *TraceFieldsResponse) SetFields(v []TraceFieldDetail)`

SetFields sets Fields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


