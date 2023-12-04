# SavedSearchSyncDefinitionBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryText** | **string** | The text of a Sumo Logic query. | 
**ByReceiptTime** | **bool** | Set it to true to run the search using receipt time. By default, searches do not run by receipt time. | [default to false]
**ViewName** | Pointer to **string** | The name of the Scheduled View that has indexed the data you want to search. | [optional] 
**ViewStartTime** | Pointer to **time.Time** | Start timestamp of the Scheduled View in UTC format. | [optional] 
**QueryParameters** | [**[]QueryParameterSyncDefinition**](QueryParameterSyncDefinition.md) | An array of search query parameter objects. | 
**ParsingMode** | Pointer to **string** | Define the parsing mode to scan the JSON format log messages. Possible values are:   1. &#x60;AutoParse&#x60;   2. &#x60;Manual&#x60; In AutoParse mode, the system automatically figures out fields to parse based on the search query. While in the Manual mode, no fields are parsed out automatically. For more information see [Dynamic Parsing](https://help.sumologic.com/?cid&#x3D;0011). | [optional] [default to "Manual"]

## Methods

### NewSavedSearchSyncDefinitionBase

`func NewSavedSearchSyncDefinitionBase(queryText string, byReceiptTime bool, queryParameters []QueryParameterSyncDefinition, ) *SavedSearchSyncDefinitionBase`

NewSavedSearchSyncDefinitionBase instantiates a new SavedSearchSyncDefinitionBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSavedSearchSyncDefinitionBaseWithDefaults

`func NewSavedSearchSyncDefinitionBaseWithDefaults() *SavedSearchSyncDefinitionBase`

NewSavedSearchSyncDefinitionBaseWithDefaults instantiates a new SavedSearchSyncDefinitionBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueryText

`func (o *SavedSearchSyncDefinitionBase) GetQueryText() string`

GetQueryText returns the QueryText field if non-nil, zero value otherwise.

### GetQueryTextOk

`func (o *SavedSearchSyncDefinitionBase) GetQueryTextOk() (*string, bool)`

GetQueryTextOk returns a tuple with the QueryText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryText

`func (o *SavedSearchSyncDefinitionBase) SetQueryText(v string)`

SetQueryText sets QueryText field to given value.


### GetByReceiptTime

`func (o *SavedSearchSyncDefinitionBase) GetByReceiptTime() bool`

GetByReceiptTime returns the ByReceiptTime field if non-nil, zero value otherwise.

### GetByReceiptTimeOk

`func (o *SavedSearchSyncDefinitionBase) GetByReceiptTimeOk() (*bool, bool)`

GetByReceiptTimeOk returns a tuple with the ByReceiptTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByReceiptTime

`func (o *SavedSearchSyncDefinitionBase) SetByReceiptTime(v bool)`

SetByReceiptTime sets ByReceiptTime field to given value.


### GetViewName

`func (o *SavedSearchSyncDefinitionBase) GetViewName() string`

GetViewName returns the ViewName field if non-nil, zero value otherwise.

### GetViewNameOk

`func (o *SavedSearchSyncDefinitionBase) GetViewNameOk() (*string, bool)`

GetViewNameOk returns a tuple with the ViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewName

`func (o *SavedSearchSyncDefinitionBase) SetViewName(v string)`

SetViewName sets ViewName field to given value.

### HasViewName

`func (o *SavedSearchSyncDefinitionBase) HasViewName() bool`

HasViewName returns a boolean if a field has been set.

### GetViewStartTime

`func (o *SavedSearchSyncDefinitionBase) GetViewStartTime() time.Time`

GetViewStartTime returns the ViewStartTime field if non-nil, zero value otherwise.

### GetViewStartTimeOk

`func (o *SavedSearchSyncDefinitionBase) GetViewStartTimeOk() (*time.Time, bool)`

GetViewStartTimeOk returns a tuple with the ViewStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewStartTime

`func (o *SavedSearchSyncDefinitionBase) SetViewStartTime(v time.Time)`

SetViewStartTime sets ViewStartTime field to given value.

### HasViewStartTime

`func (o *SavedSearchSyncDefinitionBase) HasViewStartTime() bool`

HasViewStartTime returns a boolean if a field has been set.

### GetQueryParameters

`func (o *SavedSearchSyncDefinitionBase) GetQueryParameters() []QueryParameterSyncDefinition`

GetQueryParameters returns the QueryParameters field if non-nil, zero value otherwise.

### GetQueryParametersOk

`func (o *SavedSearchSyncDefinitionBase) GetQueryParametersOk() (*[]QueryParameterSyncDefinition, bool)`

GetQueryParametersOk returns a tuple with the QueryParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryParameters

`func (o *SavedSearchSyncDefinitionBase) SetQueryParameters(v []QueryParameterSyncDefinition)`

SetQueryParameters sets QueryParameters field to given value.


### GetParsingMode

`func (o *SavedSearchSyncDefinitionBase) GetParsingMode() string`

GetParsingMode returns the ParsingMode field if non-nil, zero value otherwise.

### GetParsingModeOk

`func (o *SavedSearchSyncDefinitionBase) GetParsingModeOk() (*string, bool)`

GetParsingModeOk returns a tuple with the ParsingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParsingMode

`func (o *SavedSearchSyncDefinitionBase) SetParsingMode(v string)`

SetParsingMode sets ParsingMode field to given value.

### HasParsingMode

`func (o *SavedSearchSyncDefinitionBase) HasParsingMode() bool`

HasParsingMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


