# LogSearch

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
**CreatedAt** | **time.Time** | Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format. | 
**CreatedBy** | **string** | Identifier of the user who created the resource. | 
**ModifiedAt** | **time.Time** | Last modification timestamp in UTC. | 
**ModifiedBy** | **string** | Identifier of the user who last modified the resource. | 
**Id** | **string** | Identifier of the saved log search. | 
**ParentId** | Pointer to **string** | Identifier of the parent element in the content library, such as folder. | [optional] 

## Methods

### NewLogSearch

`func NewLogSearch(queryString string, timeRange ResolvableTimeRange, name string, createdAt time.Time, createdBy string, modifiedAt time.Time, modifiedBy string, id string, ) *LogSearch`

NewLogSearch instantiates a new LogSearch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogSearchWithDefaults

`func NewLogSearchWithDefaults() *LogSearch`

NewLogSearchWithDefaults instantiates a new LogSearch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueryString

`func (o *LogSearch) GetQueryString() string`

GetQueryString returns the QueryString field if non-nil, zero value otherwise.

### GetQueryStringOk

`func (o *LogSearch) GetQueryStringOk() (*string, bool)`

GetQueryStringOk returns a tuple with the QueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryString

`func (o *LogSearch) SetQueryString(v string)`

SetQueryString sets QueryString field to given value.


### GetTimeRange

`func (o *LogSearch) GetTimeRange() ResolvableTimeRange`

GetTimeRange returns the TimeRange field if non-nil, zero value otherwise.

### GetTimeRangeOk

`func (o *LogSearch) GetTimeRangeOk() (*ResolvableTimeRange, bool)`

GetTimeRangeOk returns a tuple with the TimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRange

`func (o *LogSearch) SetTimeRange(v ResolvableTimeRange)`

SetTimeRange sets TimeRange field to given value.


### GetRunByReceiptTime

`func (o *LogSearch) GetRunByReceiptTime() bool`

GetRunByReceiptTime returns the RunByReceiptTime field if non-nil, zero value otherwise.

### GetRunByReceiptTimeOk

`func (o *LogSearch) GetRunByReceiptTimeOk() (*bool, bool)`

GetRunByReceiptTimeOk returns a tuple with the RunByReceiptTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunByReceiptTime

`func (o *LogSearch) SetRunByReceiptTime(v bool)`

SetRunByReceiptTime sets RunByReceiptTime field to given value.

### HasRunByReceiptTime

`func (o *LogSearch) HasRunByReceiptTime() bool`

HasRunByReceiptTime returns a boolean if a field has been set.

### GetQueryParameters

`func (o *LogSearch) GetQueryParameters() []LogSearchQueryParameterSyncDefinitionBase`

GetQueryParameters returns the QueryParameters field if non-nil, zero value otherwise.

### GetQueryParametersOk

`func (o *LogSearch) GetQueryParametersOk() (*[]LogSearchQueryParameterSyncDefinitionBase, bool)`

GetQueryParametersOk returns a tuple with the QueryParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryParameters

`func (o *LogSearch) SetQueryParameters(v []LogSearchQueryParameterSyncDefinitionBase)`

SetQueryParameters sets QueryParameters field to given value.

### HasQueryParameters

`func (o *LogSearch) HasQueryParameters() bool`

HasQueryParameters returns a boolean if a field has been set.

### GetParsingMode

`func (o *LogSearch) GetParsingMode() string`

GetParsingMode returns the ParsingMode field if non-nil, zero value otherwise.

### GetParsingModeOk

`func (o *LogSearch) GetParsingModeOk() (*string, bool)`

GetParsingModeOk returns a tuple with the ParsingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParsingMode

`func (o *LogSearch) SetParsingMode(v string)`

SetParsingMode sets ParsingMode field to given value.

### HasParsingMode

`func (o *LogSearch) HasParsingMode() bool`

HasParsingMode returns a boolean if a field has been set.

### GetName

`func (o *LogSearch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LogSearch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LogSearch) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *LogSearch) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LogSearch) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LogSearch) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LogSearch) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSchedule

`func (o *LogSearch) GetSchedule() LogSearchScheduleSyncDefinition`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *LogSearch) GetScheduleOk() (*LogSearchScheduleSyncDefinition, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *LogSearch) SetSchedule(v LogSearchScheduleSyncDefinition)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *LogSearch) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetProperties

`func (o *LogSearch) GetProperties() string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *LogSearch) GetPropertiesOk() (*string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *LogSearch) SetProperties(v string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *LogSearch) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetCreatedAt

`func (o *LogSearch) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LogSearch) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LogSearch) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *LogSearch) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *LogSearch) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *LogSearch) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetModifiedAt

`func (o *LogSearch) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *LogSearch) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *LogSearch) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *LogSearch) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *LogSearch) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *LogSearch) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetId

`func (o *LogSearch) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LogSearch) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LogSearch) SetId(v string)`

SetId sets Id field to given value.


### GetParentId

`func (o *LogSearch) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *LogSearch) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *LogSearch) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *LogSearch) HasParentId() bool`

HasParentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


