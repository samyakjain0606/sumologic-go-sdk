# AlertsLibraryAlertUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MonitorId** | Pointer to **string** | The Id of the associated monitor. | [optional] 
**ResolvedAt** | Pointer to **NullableTime** | The time at which the alert was resolved. | [optional] 
**AbnormalityStartTime** | Pointer to **time.Time** | The time at which the incident started. | [optional] 
**AlertType** | Pointer to **string** | The severity of the Alert. Valid values:   1. &#x60;Critical&#x60;   2. &#x60;Warning&#x60;   3. &#x60;MissingData&#x60; | [optional] 
**Status** | Pointer to **string** | The status of the Alert. Valid values:   1. &#x60;Triggered&#x60;   2. &#x60;Resolved&#x60; | [optional] 
**MonitorQueries** | Pointer to [**[]AlertMonitorQuery**](AlertMonitorQuery.md) | All queries from the monitor relevant to the alert. | [optional] 
**TriggerQueries** | Pointer to [**[]AlertMonitorQuery**](AlertMonitorQuery.md) | All queries from the monitor relevant to the alert with triggered time series filters. | [optional] 
**MonitorUrl** | Pointer to **string** | URL for this monitor&#39;s view page | [optional] 
**TriggerQueryUrl** | Pointer to **string** | A link to search with the triggering data and time range | [optional] 
**TriggerConditions** | Pointer to [**[]TriggerCondition**](TriggerCondition.md) | Trigger conditions which were breached to create this Alert. | [optional] 
**TriggerValue** | Pointer to **float64** | The of the query result which breached the trigger condition. | [optional] 
**MonitorType** | Pointer to **string** | The type of monitor. Valid values:   1. &#x60;Logs&#x60;: A logs query monitor.   2. &#x60;Metrics&#x60;: A metrics query monitor. | [optional] 
**EntityIds** | Pointer to **[]string** | One or more primary entity identifiers involved in this Alert. Primary/secondary entities are explained in description for &#x60;entities&#x60;. DEPRECATED, USE &#x60;entities&#x60; INSTEAD.  | [optional] 
**Entities** | Pointer to [**[]AlertEntityInfo**](AlertEntityInfo.md) | One or more primary entities involved in this Alert. Primary entity is the most concrete entity that can be assigned per time series or log group (e.g. k8s container), secondary entities are the less specific ones that can be assigned per that notification (e.g. k8s cluster or EC2 host).  | [optional] 
**SecondaryEntities** | Pointer to [**[]AlertEntityInfo**](AlertEntityInfo.md) | One or more secondary entity involved in this Alert. Primary/secondary entities are explained in description for &#x60;entities&#x60;  | [optional] [default to []]
**Notes** | Pointer to **string** |  | [optional] 
**ExtraDetails** | Pointer to [**ExtraDetails**](ExtraDetails.md) |  | [optional] 
**AlertCondition** | Pointer to **NullableString** | The condition which triggered this alert. | [optional] 
**IsMuted** | Pointer to **bool** | Flag of the alerts muting status. | [optional] 

## Methods

### NewAlertsLibraryAlertUpdate

`func NewAlertsLibraryAlertUpdate() *AlertsLibraryAlertUpdate`

NewAlertsLibraryAlertUpdate instantiates a new AlertsLibraryAlertUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertsLibraryAlertUpdateWithDefaults

`func NewAlertsLibraryAlertUpdateWithDefaults() *AlertsLibraryAlertUpdate`

NewAlertsLibraryAlertUpdateWithDefaults instantiates a new AlertsLibraryAlertUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonitorId

`func (o *AlertsLibraryAlertUpdate) GetMonitorId() string`

GetMonitorId returns the MonitorId field if non-nil, zero value otherwise.

### GetMonitorIdOk

`func (o *AlertsLibraryAlertUpdate) GetMonitorIdOk() (*string, bool)`

GetMonitorIdOk returns a tuple with the MonitorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorId

`func (o *AlertsLibraryAlertUpdate) SetMonitorId(v string)`

SetMonitorId sets MonitorId field to given value.

### HasMonitorId

`func (o *AlertsLibraryAlertUpdate) HasMonitorId() bool`

HasMonitorId returns a boolean if a field has been set.

### GetResolvedAt

`func (o *AlertsLibraryAlertUpdate) GetResolvedAt() time.Time`

GetResolvedAt returns the ResolvedAt field if non-nil, zero value otherwise.

### GetResolvedAtOk

`func (o *AlertsLibraryAlertUpdate) GetResolvedAtOk() (*time.Time, bool)`

GetResolvedAtOk returns a tuple with the ResolvedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedAt

`func (o *AlertsLibraryAlertUpdate) SetResolvedAt(v time.Time)`

SetResolvedAt sets ResolvedAt field to given value.

### HasResolvedAt

`func (o *AlertsLibraryAlertUpdate) HasResolvedAt() bool`

HasResolvedAt returns a boolean if a field has been set.

### SetResolvedAtNil

`func (o *AlertsLibraryAlertUpdate) SetResolvedAtNil(b bool)`

 SetResolvedAtNil sets the value for ResolvedAt to be an explicit nil

### UnsetResolvedAt
`func (o *AlertsLibraryAlertUpdate) UnsetResolvedAt()`

UnsetResolvedAt ensures that no value is present for ResolvedAt, not even an explicit nil
### GetAbnormalityStartTime

`func (o *AlertsLibraryAlertUpdate) GetAbnormalityStartTime() time.Time`

GetAbnormalityStartTime returns the AbnormalityStartTime field if non-nil, zero value otherwise.

### GetAbnormalityStartTimeOk

`func (o *AlertsLibraryAlertUpdate) GetAbnormalityStartTimeOk() (*time.Time, bool)`

GetAbnormalityStartTimeOk returns a tuple with the AbnormalityStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbnormalityStartTime

`func (o *AlertsLibraryAlertUpdate) SetAbnormalityStartTime(v time.Time)`

SetAbnormalityStartTime sets AbnormalityStartTime field to given value.

### HasAbnormalityStartTime

`func (o *AlertsLibraryAlertUpdate) HasAbnormalityStartTime() bool`

HasAbnormalityStartTime returns a boolean if a field has been set.

### GetAlertType

`func (o *AlertsLibraryAlertUpdate) GetAlertType() string`

GetAlertType returns the AlertType field if non-nil, zero value otherwise.

### GetAlertTypeOk

`func (o *AlertsLibraryAlertUpdate) GetAlertTypeOk() (*string, bool)`

GetAlertTypeOk returns a tuple with the AlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertType

`func (o *AlertsLibraryAlertUpdate) SetAlertType(v string)`

SetAlertType sets AlertType field to given value.

### HasAlertType

`func (o *AlertsLibraryAlertUpdate) HasAlertType() bool`

HasAlertType returns a boolean if a field has been set.

### GetStatus

`func (o *AlertsLibraryAlertUpdate) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AlertsLibraryAlertUpdate) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AlertsLibraryAlertUpdate) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AlertsLibraryAlertUpdate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMonitorQueries

`func (o *AlertsLibraryAlertUpdate) GetMonitorQueries() []AlertMonitorQuery`

GetMonitorQueries returns the MonitorQueries field if non-nil, zero value otherwise.

### GetMonitorQueriesOk

`func (o *AlertsLibraryAlertUpdate) GetMonitorQueriesOk() (*[]AlertMonitorQuery, bool)`

GetMonitorQueriesOk returns a tuple with the MonitorQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorQueries

`func (o *AlertsLibraryAlertUpdate) SetMonitorQueries(v []AlertMonitorQuery)`

SetMonitorQueries sets MonitorQueries field to given value.

### HasMonitorQueries

`func (o *AlertsLibraryAlertUpdate) HasMonitorQueries() bool`

HasMonitorQueries returns a boolean if a field has been set.

### GetTriggerQueries

`func (o *AlertsLibraryAlertUpdate) GetTriggerQueries() []AlertMonitorQuery`

GetTriggerQueries returns the TriggerQueries field if non-nil, zero value otherwise.

### GetTriggerQueriesOk

`func (o *AlertsLibraryAlertUpdate) GetTriggerQueriesOk() (*[]AlertMonitorQuery, bool)`

GetTriggerQueriesOk returns a tuple with the TriggerQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerQueries

`func (o *AlertsLibraryAlertUpdate) SetTriggerQueries(v []AlertMonitorQuery)`

SetTriggerQueries sets TriggerQueries field to given value.

### HasTriggerQueries

`func (o *AlertsLibraryAlertUpdate) HasTriggerQueries() bool`

HasTriggerQueries returns a boolean if a field has been set.

### GetMonitorUrl

`func (o *AlertsLibraryAlertUpdate) GetMonitorUrl() string`

GetMonitorUrl returns the MonitorUrl field if non-nil, zero value otherwise.

### GetMonitorUrlOk

`func (o *AlertsLibraryAlertUpdate) GetMonitorUrlOk() (*string, bool)`

GetMonitorUrlOk returns a tuple with the MonitorUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorUrl

`func (o *AlertsLibraryAlertUpdate) SetMonitorUrl(v string)`

SetMonitorUrl sets MonitorUrl field to given value.

### HasMonitorUrl

`func (o *AlertsLibraryAlertUpdate) HasMonitorUrl() bool`

HasMonitorUrl returns a boolean if a field has been set.

### GetTriggerQueryUrl

`func (o *AlertsLibraryAlertUpdate) GetTriggerQueryUrl() string`

GetTriggerQueryUrl returns the TriggerQueryUrl field if non-nil, zero value otherwise.

### GetTriggerQueryUrlOk

`func (o *AlertsLibraryAlertUpdate) GetTriggerQueryUrlOk() (*string, bool)`

GetTriggerQueryUrlOk returns a tuple with the TriggerQueryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerQueryUrl

`func (o *AlertsLibraryAlertUpdate) SetTriggerQueryUrl(v string)`

SetTriggerQueryUrl sets TriggerQueryUrl field to given value.

### HasTriggerQueryUrl

`func (o *AlertsLibraryAlertUpdate) HasTriggerQueryUrl() bool`

HasTriggerQueryUrl returns a boolean if a field has been set.

### GetTriggerConditions

`func (o *AlertsLibraryAlertUpdate) GetTriggerConditions() []TriggerCondition`

GetTriggerConditions returns the TriggerConditions field if non-nil, zero value otherwise.

### GetTriggerConditionsOk

`func (o *AlertsLibraryAlertUpdate) GetTriggerConditionsOk() (*[]TriggerCondition, bool)`

GetTriggerConditionsOk returns a tuple with the TriggerConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerConditions

`func (o *AlertsLibraryAlertUpdate) SetTriggerConditions(v []TriggerCondition)`

SetTriggerConditions sets TriggerConditions field to given value.

### HasTriggerConditions

`func (o *AlertsLibraryAlertUpdate) HasTriggerConditions() bool`

HasTriggerConditions returns a boolean if a field has been set.

### GetTriggerValue

`func (o *AlertsLibraryAlertUpdate) GetTriggerValue() float64`

GetTriggerValue returns the TriggerValue field if non-nil, zero value otherwise.

### GetTriggerValueOk

`func (o *AlertsLibraryAlertUpdate) GetTriggerValueOk() (*float64, bool)`

GetTriggerValueOk returns a tuple with the TriggerValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerValue

`func (o *AlertsLibraryAlertUpdate) SetTriggerValue(v float64)`

SetTriggerValue sets TriggerValue field to given value.

### HasTriggerValue

`func (o *AlertsLibraryAlertUpdate) HasTriggerValue() bool`

HasTriggerValue returns a boolean if a field has been set.

### GetMonitorType

`func (o *AlertsLibraryAlertUpdate) GetMonitorType() string`

GetMonitorType returns the MonitorType field if non-nil, zero value otherwise.

### GetMonitorTypeOk

`func (o *AlertsLibraryAlertUpdate) GetMonitorTypeOk() (*string, bool)`

GetMonitorTypeOk returns a tuple with the MonitorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorType

`func (o *AlertsLibraryAlertUpdate) SetMonitorType(v string)`

SetMonitorType sets MonitorType field to given value.

### HasMonitorType

`func (o *AlertsLibraryAlertUpdate) HasMonitorType() bool`

HasMonitorType returns a boolean if a field has been set.

### GetEntityIds

`func (o *AlertsLibraryAlertUpdate) GetEntityIds() []string`

GetEntityIds returns the EntityIds field if non-nil, zero value otherwise.

### GetEntityIdsOk

`func (o *AlertsLibraryAlertUpdate) GetEntityIdsOk() (*[]string, bool)`

GetEntityIdsOk returns a tuple with the EntityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityIds

`func (o *AlertsLibraryAlertUpdate) SetEntityIds(v []string)`

SetEntityIds sets EntityIds field to given value.

### HasEntityIds

`func (o *AlertsLibraryAlertUpdate) HasEntityIds() bool`

HasEntityIds returns a boolean if a field has been set.

### GetEntities

`func (o *AlertsLibraryAlertUpdate) GetEntities() []AlertEntityInfo`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *AlertsLibraryAlertUpdate) GetEntitiesOk() (*[]AlertEntityInfo, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *AlertsLibraryAlertUpdate) SetEntities(v []AlertEntityInfo)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *AlertsLibraryAlertUpdate) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetSecondaryEntities

`func (o *AlertsLibraryAlertUpdate) GetSecondaryEntities() []AlertEntityInfo`

GetSecondaryEntities returns the SecondaryEntities field if non-nil, zero value otherwise.

### GetSecondaryEntitiesOk

`func (o *AlertsLibraryAlertUpdate) GetSecondaryEntitiesOk() (*[]AlertEntityInfo, bool)`

GetSecondaryEntitiesOk returns a tuple with the SecondaryEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryEntities

`func (o *AlertsLibraryAlertUpdate) SetSecondaryEntities(v []AlertEntityInfo)`

SetSecondaryEntities sets SecondaryEntities field to given value.

### HasSecondaryEntities

`func (o *AlertsLibraryAlertUpdate) HasSecondaryEntities() bool`

HasSecondaryEntities returns a boolean if a field has been set.

### GetNotes

`func (o *AlertsLibraryAlertUpdate) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *AlertsLibraryAlertUpdate) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *AlertsLibraryAlertUpdate) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *AlertsLibraryAlertUpdate) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetExtraDetails

`func (o *AlertsLibraryAlertUpdate) GetExtraDetails() ExtraDetails`

GetExtraDetails returns the ExtraDetails field if non-nil, zero value otherwise.

### GetExtraDetailsOk

`func (o *AlertsLibraryAlertUpdate) GetExtraDetailsOk() (*ExtraDetails, bool)`

GetExtraDetailsOk returns a tuple with the ExtraDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraDetails

`func (o *AlertsLibraryAlertUpdate) SetExtraDetails(v ExtraDetails)`

SetExtraDetails sets ExtraDetails field to given value.

### HasExtraDetails

`func (o *AlertsLibraryAlertUpdate) HasExtraDetails() bool`

HasExtraDetails returns a boolean if a field has been set.

### GetAlertCondition

`func (o *AlertsLibraryAlertUpdate) GetAlertCondition() string`

GetAlertCondition returns the AlertCondition field if non-nil, zero value otherwise.

### GetAlertConditionOk

`func (o *AlertsLibraryAlertUpdate) GetAlertConditionOk() (*string, bool)`

GetAlertConditionOk returns a tuple with the AlertCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertCondition

`func (o *AlertsLibraryAlertUpdate) SetAlertCondition(v string)`

SetAlertCondition sets AlertCondition field to given value.

### HasAlertCondition

`func (o *AlertsLibraryAlertUpdate) HasAlertCondition() bool`

HasAlertCondition returns a boolean if a field has been set.

### SetAlertConditionNil

`func (o *AlertsLibraryAlertUpdate) SetAlertConditionNil(b bool)`

 SetAlertConditionNil sets the value for AlertCondition to be an explicit nil

### UnsetAlertCondition
`func (o *AlertsLibraryAlertUpdate) UnsetAlertCondition()`

UnsetAlertCondition ensures that no value is present for AlertCondition, not even an explicit nil
### GetIsMuted

`func (o *AlertsLibraryAlertUpdate) GetIsMuted() bool`

GetIsMuted returns the IsMuted field if non-nil, zero value otherwise.

### GetIsMutedOk

`func (o *AlertsLibraryAlertUpdate) GetIsMutedOk() (*bool, bool)`

GetIsMutedOk returns a tuple with the IsMuted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMuted

`func (o *AlertsLibraryAlertUpdate) SetIsMuted(v bool)`

SetIsMuted sets IsMuted field to given value.

### HasIsMuted

`func (o *AlertsLibraryAlertUpdate) HasIsMuted() bool`

HasIsMuted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


