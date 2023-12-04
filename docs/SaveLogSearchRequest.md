# SaveLogSearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryString** | **string** | Query to perform. | 
**TimeRange** | [**ResolvableTimeRange**](ResolvableTimeRange.md) |  | 
**RunByReceiptTime** | Pointer to **bool** | This has the value &#x60;true&#x60; if the search is to be run by receipt time and &#x60;false&#x60; if it is to be run by message time. | [optional] [default to false]
**QueryParameters** | Pointer to [**[]LogSearchQueryParameterSyncDefinitionBase**](LogSearchQueryParameterSyncDefinitionBase.md) | Values for search template used in the search query. Learn more about the search templates here : https://help.sumologic.com/docs/search/get-started-with-search/build-search/search-templates/ | [optional] 
**ParsingMode** | Pointer to **string** | Define the parsing mode to scan the JSON format log messages. Possible values are:   1. &#x60;AutoParse&#x60;   2. &#x60;Manual&#x60; In AutoParse mode, the system automatically figures out fields to parse based on the search query. While in the Manual mode, no fields are parsed out automatically. For more information see [Dynamic Parsing](https://help.sumologic.com/?cid&#x3D;0011). | [optional] [default to "Manual"]
**Name** | **string** | Name of the item in the content library. | 
**Description** | Pointer to **string** | Item description in the content library. | [optional] 
**Schedule** | Pointer to [**LogSearchScheduleSyncDefinition**](LogSearchScheduleSyncDefinition.md) |  | [optional] 
**Properties** | Pointer to **string** | Aggregate Results Settings and View configurations, Legends settings, and different visualisation settings overrides. Leave this field empty to use the defaults. This property contains JSON object encoded as a string.  | [optional] 
**ParentId** | **string** | Identifier of a folder where to save the log search. | 

## Methods

### NewSaveLogSearchRequest

`func NewSaveLogSearchRequest(queryString string, timeRange ResolvableTimeRange, name string, parentId string, ) *SaveLogSearchRequest`

NewSaveLogSearchRequest instantiates a new SaveLogSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaveLogSearchRequestWithDefaults

`func NewSaveLogSearchRequestWithDefaults() *SaveLogSearchRequest`

NewSaveLogSearchRequestWithDefaults instantiates a new SaveLogSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueryString

`func (o *SaveLogSearchRequest) GetQueryString() string`

GetQueryString returns the QueryString field if non-nil, zero value otherwise.

### GetQueryStringOk

`func (o *SaveLogSearchRequest) GetQueryStringOk() (*string, bool)`

GetQueryStringOk returns a tuple with the QueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryString

`func (o *SaveLogSearchRequest) SetQueryString(v string)`

SetQueryString sets QueryString field to given value.


### GetTimeRange

`func (o *SaveLogSearchRequest) GetTimeRange() ResolvableTimeRange`

GetTimeRange returns the TimeRange field if non-nil, zero value otherwise.

### GetTimeRangeOk

`func (o *SaveLogSearchRequest) GetTimeRangeOk() (*ResolvableTimeRange, bool)`

GetTimeRangeOk returns a tuple with the TimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRange

`func (o *SaveLogSearchRequest) SetTimeRange(v ResolvableTimeRange)`

SetTimeRange sets TimeRange field to given value.


### GetRunByReceiptTime

`func (o *SaveLogSearchRequest) GetRunByReceiptTime() bool`

GetRunByReceiptTime returns the RunByReceiptTime field if non-nil, zero value otherwise.

### GetRunByReceiptTimeOk

`func (o *SaveLogSearchRequest) GetRunByReceiptTimeOk() (*bool, bool)`

GetRunByReceiptTimeOk returns a tuple with the RunByReceiptTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunByReceiptTime

`func (o *SaveLogSearchRequest) SetRunByReceiptTime(v bool)`

SetRunByReceiptTime sets RunByReceiptTime field to given value.

### HasRunByReceiptTime

`func (o *SaveLogSearchRequest) HasRunByReceiptTime() bool`

HasRunByReceiptTime returns a boolean if a field has been set.

### GetQueryParameters

`func (o *SaveLogSearchRequest) GetQueryParameters() []LogSearchQueryParameterSyncDefinitionBase`

GetQueryParameters returns the QueryParameters field if non-nil, zero value otherwise.

### GetQueryParametersOk

`func (o *SaveLogSearchRequest) GetQueryParametersOk() (*[]LogSearchQueryParameterSyncDefinitionBase, bool)`

GetQueryParametersOk returns a tuple with the QueryParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryParameters

`func (o *SaveLogSearchRequest) SetQueryParameters(v []LogSearchQueryParameterSyncDefinitionBase)`

SetQueryParameters sets QueryParameters field to given value.

### HasQueryParameters

`func (o *SaveLogSearchRequest) HasQueryParameters() bool`

HasQueryParameters returns a boolean if a field has been set.

### GetParsingMode

`func (o *SaveLogSearchRequest) GetParsingMode() string`

GetParsingMode returns the ParsingMode field if non-nil, zero value otherwise.

### GetParsingModeOk

`func (o *SaveLogSearchRequest) GetParsingModeOk() (*string, bool)`

GetParsingModeOk returns a tuple with the ParsingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParsingMode

`func (o *SaveLogSearchRequest) SetParsingMode(v string)`

SetParsingMode sets ParsingMode field to given value.

### HasParsingMode

`func (o *SaveLogSearchRequest) HasParsingMode() bool`

HasParsingMode returns a boolean if a field has been set.

### GetName

`func (o *SaveLogSearchRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SaveLogSearchRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SaveLogSearchRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *SaveLogSearchRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SaveLogSearchRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SaveLogSearchRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SaveLogSearchRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSchedule

`func (o *SaveLogSearchRequest) GetSchedule() LogSearchScheduleSyncDefinition`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *SaveLogSearchRequest) GetScheduleOk() (*LogSearchScheduleSyncDefinition, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *SaveLogSearchRequest) SetSchedule(v LogSearchScheduleSyncDefinition)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *SaveLogSearchRequest) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetProperties

`func (o *SaveLogSearchRequest) GetProperties() string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *SaveLogSearchRequest) GetPropertiesOk() (*string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *SaveLogSearchRequest) SetProperties(v string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *SaveLogSearchRequest) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetParentId

`func (o *SaveLogSearchRequest) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *SaveLogSearchRequest) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *SaveLogSearchRequest) SetParentId(v string)`

SetParentId sets ParentId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


