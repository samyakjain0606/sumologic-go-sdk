# AttributeReversedIndex

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttributeName** | **string** | Name of the attribute. | 
**AttributeValueStatistics** | [**[]AttributeValueReversedIndex**](AttributeValueReversedIndex.md) | List of value statistics of the given attribute. | 

## Methods

### NewAttributeReversedIndex

`func NewAttributeReversedIndex(attributeName string, attributeValueStatistics []AttributeValueReversedIndex, ) *AttributeReversedIndex`

NewAttributeReversedIndex instantiates a new AttributeReversedIndex object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttributeReversedIndexWithDefaults

`func NewAttributeReversedIndexWithDefaults() *AttributeReversedIndex`

NewAttributeReversedIndexWithDefaults instantiates a new AttributeReversedIndex object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributeName

`func (o *AttributeReversedIndex) GetAttributeName() string`

GetAttributeName returns the AttributeName field if non-nil, zero value otherwise.

### GetAttributeNameOk

`func (o *AttributeReversedIndex) GetAttributeNameOk() (*string, bool)`

GetAttributeNameOk returns a tuple with the AttributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeName

`func (o *AttributeReversedIndex) SetAttributeName(v string)`

SetAttributeName sets AttributeName field to given value.


### GetAttributeValueStatistics

`func (o *AttributeReversedIndex) GetAttributeValueStatistics() []AttributeValueReversedIndex`

GetAttributeValueStatistics returns the AttributeValueStatistics field if non-nil, zero value otherwise.

### GetAttributeValueStatisticsOk

`func (o *AttributeReversedIndex) GetAttributeValueStatisticsOk() (*[]AttributeValueReversedIndex, bool)`

GetAttributeValueStatisticsOk returns a tuple with the AttributeValueStatistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeValueStatistics

`func (o *AttributeReversedIndex) SetAttributeValueStatistics(v []AttributeValueReversedIndex)`

SetAttributeValueStatistics sets AttributeValueStatistics field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


