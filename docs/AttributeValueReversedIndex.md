# AttributeValueReversedIndex

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttributeValue** | **string** | Value of the attribute. | 
**SpanIds** | **[]string** | List of span ids which have the given attribute and value. | 

## Methods

### NewAttributeValueReversedIndex

`func NewAttributeValueReversedIndex(attributeValue string, spanIds []string, ) *AttributeValueReversedIndex`

NewAttributeValueReversedIndex instantiates a new AttributeValueReversedIndex object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttributeValueReversedIndexWithDefaults

`func NewAttributeValueReversedIndexWithDefaults() *AttributeValueReversedIndex`

NewAttributeValueReversedIndexWithDefaults instantiates a new AttributeValueReversedIndex object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributeValue

`func (o *AttributeValueReversedIndex) GetAttributeValue() string`

GetAttributeValue returns the AttributeValue field if non-nil, zero value otherwise.

### GetAttributeValueOk

`func (o *AttributeValueReversedIndex) GetAttributeValueOk() (*string, bool)`

GetAttributeValueOk returns a tuple with the AttributeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeValue

`func (o *AttributeValueReversedIndex) SetAttributeValue(v string)`

SetAttributeValue sets AttributeValue field to given value.


### GetSpanIds

`func (o *AttributeValueReversedIndex) GetSpanIds() []string`

GetSpanIds returns the SpanIds field if non-nil, zero value otherwise.

### GetSpanIdsOk

`func (o *AttributeValueReversedIndex) GetSpanIdsOk() (*[]string, bool)`

GetSpanIdsOk returns a tuple with the SpanIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanIds

`func (o *AttributeValueReversedIndex) SetSpanIds(v []string)`

SetSpanIds sets SpanIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


