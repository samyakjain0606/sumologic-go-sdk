# SpanEventAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttributeName** | Pointer to **string** | Name of the attribute. | [optional] 
**AttributeValue** | Pointer to [**EventAttributeValue**](EventAttributeValue.md) |  | [optional] 

## Methods

### NewSpanEventAttribute

`func NewSpanEventAttribute() *SpanEventAttribute`

NewSpanEventAttribute instantiates a new SpanEventAttribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpanEventAttributeWithDefaults

`func NewSpanEventAttributeWithDefaults() *SpanEventAttribute`

NewSpanEventAttributeWithDefaults instantiates a new SpanEventAttribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributeName

`func (o *SpanEventAttribute) GetAttributeName() string`

GetAttributeName returns the AttributeName field if non-nil, zero value otherwise.

### GetAttributeNameOk

`func (o *SpanEventAttribute) GetAttributeNameOk() (*string, bool)`

GetAttributeNameOk returns a tuple with the AttributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeName

`func (o *SpanEventAttribute) SetAttributeName(v string)`

SetAttributeName sets AttributeName field to given value.

### HasAttributeName

`func (o *SpanEventAttribute) HasAttributeName() bool`

HasAttributeName returns a boolean if a field has been set.

### GetAttributeValue

`func (o *SpanEventAttribute) GetAttributeValue() EventAttributeValue`

GetAttributeValue returns the AttributeValue field if non-nil, zero value otherwise.

### GetAttributeValueOk

`func (o *SpanEventAttribute) GetAttributeValueOk() (*EventAttributeValue, bool)`

GetAttributeValueOk returns a tuple with the AttributeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeValue

`func (o *SpanEventAttribute) SetAttributeValue(v EventAttributeValue)`

SetAttributeValue sets AttributeValue field to given value.

### HasAttributeValue

`func (o *SpanEventAttribute) HasAttributeValue() bool`

HasAttributeValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


