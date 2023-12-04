# TagReversedIndex

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TagName** | **string** | Name of the tag. | 
**TagValueStatistics** | [**[]TagValueReversedIndex**](TagValueReversedIndex.md) | List of value statistics of the given tag. | 

## Methods

### NewTagReversedIndex

`func NewTagReversedIndex(tagName string, tagValueStatistics []TagValueReversedIndex, ) *TagReversedIndex`

NewTagReversedIndex instantiates a new TagReversedIndex object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagReversedIndexWithDefaults

`func NewTagReversedIndexWithDefaults() *TagReversedIndex`

NewTagReversedIndexWithDefaults instantiates a new TagReversedIndex object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTagName

`func (o *TagReversedIndex) GetTagName() string`

GetTagName returns the TagName field if non-nil, zero value otherwise.

### GetTagNameOk

`func (o *TagReversedIndex) GetTagNameOk() (*string, bool)`

GetTagNameOk returns a tuple with the TagName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagName

`func (o *TagReversedIndex) SetTagName(v string)`

SetTagName sets TagName field to given value.


### GetTagValueStatistics

`func (o *TagReversedIndex) GetTagValueStatistics() []TagValueReversedIndex`

GetTagValueStatistics returns the TagValueStatistics field if non-nil, zero value otherwise.

### GetTagValueStatisticsOk

`func (o *TagReversedIndex) GetTagValueStatisticsOk() (*[]TagValueReversedIndex, bool)`

GetTagValueStatisticsOk returns a tuple with the TagValueStatistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagValueStatistics

`func (o *TagReversedIndex) SetTagValueStatistics(v []TagValueReversedIndex)`

SetTagValueStatistics sets TagValueStatistics field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


