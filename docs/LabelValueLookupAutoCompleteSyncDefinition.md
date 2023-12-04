# LabelValueLookupAutoCompleteSyncDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoCompleteKey** | **string** | The autocomplete key to be used to fetch autocomplete values. | 
**LookupFileName** | **string** | The lookup file to use as a source for autocomplete values. | 
**LookupLabelColumn** | **string** | The column from the lookup file to use for autocomplete labels. | 
**LookupValueColumn** | **string** | The column from the lookup file to fill the actual value when a particular label is selected. | 

## Methods

### NewLabelValueLookupAutoCompleteSyncDefinition

`func NewLabelValueLookupAutoCompleteSyncDefinition(autoCompleteKey string, lookupFileName string, lookupLabelColumn string, lookupValueColumn string, ) *LabelValueLookupAutoCompleteSyncDefinition`

NewLabelValueLookupAutoCompleteSyncDefinition instantiates a new LabelValueLookupAutoCompleteSyncDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLabelValueLookupAutoCompleteSyncDefinitionWithDefaults

`func NewLabelValueLookupAutoCompleteSyncDefinitionWithDefaults() *LabelValueLookupAutoCompleteSyncDefinition`

NewLabelValueLookupAutoCompleteSyncDefinitionWithDefaults instantiates a new LabelValueLookupAutoCompleteSyncDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoCompleteKey

`func (o *LabelValueLookupAutoCompleteSyncDefinition) GetAutoCompleteKey() string`

GetAutoCompleteKey returns the AutoCompleteKey field if non-nil, zero value otherwise.

### GetAutoCompleteKeyOk

`func (o *LabelValueLookupAutoCompleteSyncDefinition) GetAutoCompleteKeyOk() (*string, bool)`

GetAutoCompleteKeyOk returns a tuple with the AutoCompleteKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCompleteKey

`func (o *LabelValueLookupAutoCompleteSyncDefinition) SetAutoCompleteKey(v string)`

SetAutoCompleteKey sets AutoCompleteKey field to given value.


### GetLookupFileName

`func (o *LabelValueLookupAutoCompleteSyncDefinition) GetLookupFileName() string`

GetLookupFileName returns the LookupFileName field if non-nil, zero value otherwise.

### GetLookupFileNameOk

`func (o *LabelValueLookupAutoCompleteSyncDefinition) GetLookupFileNameOk() (*string, bool)`

GetLookupFileNameOk returns a tuple with the LookupFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookupFileName

`func (o *LabelValueLookupAutoCompleteSyncDefinition) SetLookupFileName(v string)`

SetLookupFileName sets LookupFileName field to given value.


### GetLookupLabelColumn

`func (o *LabelValueLookupAutoCompleteSyncDefinition) GetLookupLabelColumn() string`

GetLookupLabelColumn returns the LookupLabelColumn field if non-nil, zero value otherwise.

### GetLookupLabelColumnOk

`func (o *LabelValueLookupAutoCompleteSyncDefinition) GetLookupLabelColumnOk() (*string, bool)`

GetLookupLabelColumnOk returns a tuple with the LookupLabelColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookupLabelColumn

`func (o *LabelValueLookupAutoCompleteSyncDefinition) SetLookupLabelColumn(v string)`

SetLookupLabelColumn sets LookupLabelColumn field to given value.


### GetLookupValueColumn

`func (o *LabelValueLookupAutoCompleteSyncDefinition) GetLookupValueColumn() string`

GetLookupValueColumn returns the LookupValueColumn field if non-nil, zero value otherwise.

### GetLookupValueColumnOk

`func (o *LabelValueLookupAutoCompleteSyncDefinition) GetLookupValueColumnOk() (*string, bool)`

GetLookupValueColumnOk returns a tuple with the LookupValueColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookupValueColumn

`func (o *LabelValueLookupAutoCompleteSyncDefinition) SetLookupValueColumn(v string)`

SetLookupValueColumn sets LookupValueColumn field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


