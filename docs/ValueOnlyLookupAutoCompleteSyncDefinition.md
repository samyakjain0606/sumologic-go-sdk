# ValueOnlyLookupAutoCompleteSyncDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoCompleteKey** | **string** | The autocomplete key to be used to fetch autocomplete values. | 
**LookupFileName** | **string** | The lookup file to use as a source for autocomplete values. | 
**LookupValueColumn** | **string** | The column from the lookup file to fill the actual value when a particular label is selected. | 

## Methods

### NewValueOnlyLookupAutoCompleteSyncDefinition

`func NewValueOnlyLookupAutoCompleteSyncDefinition(autoCompleteKey string, lookupFileName string, lookupValueColumn string, ) *ValueOnlyLookupAutoCompleteSyncDefinition`

NewValueOnlyLookupAutoCompleteSyncDefinition instantiates a new ValueOnlyLookupAutoCompleteSyncDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValueOnlyLookupAutoCompleteSyncDefinitionWithDefaults

`func NewValueOnlyLookupAutoCompleteSyncDefinitionWithDefaults() *ValueOnlyLookupAutoCompleteSyncDefinition`

NewValueOnlyLookupAutoCompleteSyncDefinitionWithDefaults instantiates a new ValueOnlyLookupAutoCompleteSyncDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoCompleteKey

`func (o *ValueOnlyLookupAutoCompleteSyncDefinition) GetAutoCompleteKey() string`

GetAutoCompleteKey returns the AutoCompleteKey field if non-nil, zero value otherwise.

### GetAutoCompleteKeyOk

`func (o *ValueOnlyLookupAutoCompleteSyncDefinition) GetAutoCompleteKeyOk() (*string, bool)`

GetAutoCompleteKeyOk returns a tuple with the AutoCompleteKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCompleteKey

`func (o *ValueOnlyLookupAutoCompleteSyncDefinition) SetAutoCompleteKey(v string)`

SetAutoCompleteKey sets AutoCompleteKey field to given value.


### GetLookupFileName

`func (o *ValueOnlyLookupAutoCompleteSyncDefinition) GetLookupFileName() string`

GetLookupFileName returns the LookupFileName field if non-nil, zero value otherwise.

### GetLookupFileNameOk

`func (o *ValueOnlyLookupAutoCompleteSyncDefinition) GetLookupFileNameOk() (*string, bool)`

GetLookupFileNameOk returns a tuple with the LookupFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookupFileName

`func (o *ValueOnlyLookupAutoCompleteSyncDefinition) SetLookupFileName(v string)`

SetLookupFileName sets LookupFileName field to given value.


### GetLookupValueColumn

`func (o *ValueOnlyLookupAutoCompleteSyncDefinition) GetLookupValueColumn() string`

GetLookupValueColumn returns the LookupValueColumn field if non-nil, zero value otherwise.

### GetLookupValueColumnOk

`func (o *ValueOnlyLookupAutoCompleteSyncDefinition) GetLookupValueColumnOk() (*string, bool)`

GetLookupValueColumnOk returns a tuple with the LookupValueColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookupValueColumn

`func (o *ValueOnlyLookupAutoCompleteSyncDefinition) SetLookupValueColumn(v string)`

SetLookupValueColumn sets LookupValueColumn field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


