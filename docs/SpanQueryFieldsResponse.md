# SpanQueryFieldsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fields** | [**[]SpanQueryFieldDetail**](SpanQueryFieldDetail.md) | List of span fields. | 

## Methods

### NewSpanQueryFieldsResponse

`func NewSpanQueryFieldsResponse(fields []SpanQueryFieldDetail, ) *SpanQueryFieldsResponse`

NewSpanQueryFieldsResponse instantiates a new SpanQueryFieldsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpanQueryFieldsResponseWithDefaults

`func NewSpanQueryFieldsResponseWithDefaults() *SpanQueryFieldsResponse`

NewSpanQueryFieldsResponseWithDefaults instantiates a new SpanQueryFieldsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFields

`func (o *SpanQueryFieldsResponse) GetFields() []SpanQueryFieldDetail`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *SpanQueryFieldsResponse) GetFieldsOk() (*[]SpanQueryFieldDetail, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *SpanQueryFieldsResponse) SetFields(v []SpanQueryFieldDetail)`

SetFields sets Fields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


