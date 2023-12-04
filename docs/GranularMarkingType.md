# GranularMarkingType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Lang** | Pointer to **string** | The lang property identifies the language of the text identified by this marking | [optional] 
**MarkingRef** | Pointer to **string** | The marking_ref property specifies the ID of the marking-definition object that describes the marking | [optional] 
**Selectors** | **[]string** | The selectors property specifies a list of selectors for content contained within the STIX Object in which this property appears | 

## Methods

### NewGranularMarkingType

`func NewGranularMarkingType(selectors []string, ) *GranularMarkingType`

NewGranularMarkingType instantiates a new GranularMarkingType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGranularMarkingTypeWithDefaults

`func NewGranularMarkingTypeWithDefaults() *GranularMarkingType`

NewGranularMarkingTypeWithDefaults instantiates a new GranularMarkingType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLang

`func (o *GranularMarkingType) GetLang() string`

GetLang returns the Lang field if non-nil, zero value otherwise.

### GetLangOk

`func (o *GranularMarkingType) GetLangOk() (*string, bool)`

GetLangOk returns a tuple with the Lang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLang

`func (o *GranularMarkingType) SetLang(v string)`

SetLang sets Lang field to given value.

### HasLang

`func (o *GranularMarkingType) HasLang() bool`

HasLang returns a boolean if a field has been set.

### GetMarkingRef

`func (o *GranularMarkingType) GetMarkingRef() string`

GetMarkingRef returns the MarkingRef field if non-nil, zero value otherwise.

### GetMarkingRefOk

`func (o *GranularMarkingType) GetMarkingRefOk() (*string, bool)`

GetMarkingRefOk returns a tuple with the MarkingRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkingRef

`func (o *GranularMarkingType) SetMarkingRef(v string)`

SetMarkingRef sets MarkingRef field to given value.

### HasMarkingRef

`func (o *GranularMarkingType) HasMarkingRef() bool`

HasMarkingRef returns a boolean if a field has been set.

### GetSelectors

`func (o *GranularMarkingType) GetSelectors() []string`

GetSelectors returns the Selectors field if non-nil, zero value otherwise.

### GetSelectorsOk

`func (o *GranularMarkingType) GetSelectorsOk() (*[]string, bool)`

GetSelectorsOk returns a tuple with the Selectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectors

`func (o *GranularMarkingType) SetSelectors(v []string)`

SetSelectors sets Selectors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


