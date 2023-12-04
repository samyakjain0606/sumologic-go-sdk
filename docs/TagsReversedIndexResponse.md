# TagsReversedIndexResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TagsIndices** | Pointer to [**[]TagReversedIndex**](TagReversedIndex.md) | List of spans tag values indices. | [optional] 
**AttributesIndices** | Pointer to [**[]AttributeReversedIndex**](AttributeReversedIndex.md) | List of spans attribute values indices. | [optional] 
**Next** | Pointer to **string** | Next continuation token. | [optional] 

## Methods

### NewTagsReversedIndexResponse

`func NewTagsReversedIndexResponse() *TagsReversedIndexResponse`

NewTagsReversedIndexResponse instantiates a new TagsReversedIndexResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagsReversedIndexResponseWithDefaults

`func NewTagsReversedIndexResponseWithDefaults() *TagsReversedIndexResponse`

NewTagsReversedIndexResponseWithDefaults instantiates a new TagsReversedIndexResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTagsIndices

`func (o *TagsReversedIndexResponse) GetTagsIndices() []TagReversedIndex`

GetTagsIndices returns the TagsIndices field if non-nil, zero value otherwise.

### GetTagsIndicesOk

`func (o *TagsReversedIndexResponse) GetTagsIndicesOk() (*[]TagReversedIndex, bool)`

GetTagsIndicesOk returns a tuple with the TagsIndices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagsIndices

`func (o *TagsReversedIndexResponse) SetTagsIndices(v []TagReversedIndex)`

SetTagsIndices sets TagsIndices field to given value.

### HasTagsIndices

`func (o *TagsReversedIndexResponse) HasTagsIndices() bool`

HasTagsIndices returns a boolean if a field has been set.

### GetAttributesIndices

`func (o *TagsReversedIndexResponse) GetAttributesIndices() []AttributeReversedIndex`

GetAttributesIndices returns the AttributesIndices field if non-nil, zero value otherwise.

### GetAttributesIndicesOk

`func (o *TagsReversedIndexResponse) GetAttributesIndicesOk() (*[]AttributeReversedIndex, bool)`

GetAttributesIndicesOk returns a tuple with the AttributesIndices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributesIndices

`func (o *TagsReversedIndexResponse) SetAttributesIndices(v []AttributeReversedIndex)`

SetAttributesIndices sets AttributesIndices field to given value.

### HasAttributesIndices

`func (o *TagsReversedIndexResponse) HasAttributesIndices() bool`

HasAttributesIndices returns a boolean if a field has been set.

### GetNext

`func (o *TagsReversedIndexResponse) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *TagsReversedIndexResponse) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *TagsReversedIndexResponse) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *TagsReversedIndexResponse) HasNext() bool`

HasNext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


