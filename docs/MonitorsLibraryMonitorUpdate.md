# MonitorsLibraryMonitorUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MonitorType** | **string** | The type of monitor. Valid values:   1. &#x60;Logs&#x60;: A logs query monitor.   2. &#x60;Metrics&#x60;: A metrics query monitor.   3. &#x60;Slo&#x60;: A SLO based monitor. Currently SLO based monitor is available in closed beta (Notify your Sumo Logic representative in order to get the early access). | 
**EvaluationDelay** | Pointer to **string** | The delay duration for evaluating the monitor (relative to current time). The timerange of monitor will be shifted in the past by this delay time. | [optional] [default to "0m"]
**AlertName** | Pointer to **string** | The name of the alert(s) triggered from this monitor. Monitor name will be used if not specified. All template variables can be used here except {{AlertName}}, {{AlertResponseURL}}, {{ResultsJson}}, and {{Playbook}}. | [optional] 
**RunAs** | Pointer to [**RunAs**](RunAs.md) |  | [optional] 
**NotificationGroupFields** | Pointer to **[]string** | The set of fields to be used to group alert notifications for a monitor. The value of this field will be considered only when &#39;groupNotifications&#39; is true. The fields with very high cardinality such as &#x60;_blockid&#x60;, &#x60;_raw&#x60;, &#x60;_messagetime&#x60;, &#x60;_receipttime&#x60;, and &#x60;_messageid&#x60; are not allowed for Alert Grouping. | [optional] 
**Queries** | [**[]MonitorQuery**](MonitorQuery.md) | All queries from the monitor. | 
**Triggers** | [**[]TriggerCondition**](TriggerCondition.md) | Defines the conditions of when to send notifications. | 
**Notifications** | Pointer to [**[]MonitorNotification**](MonitorNotification.md) | The notifications the monitor will send when the respective trigger condition is met. | [optional] [default to []]
**IsDisabled** | Pointer to **bool** | Whether or not the monitor is disabled. Disabled monitors will not run, and will not generate or send notifications. | [optional] [default to false]
**GroupNotifications** | Pointer to **bool** | Whether or not to group notifications for individual items that meet the trigger condition. | [optional] [default to true]
**Playbook** | Pointer to **string** | Notes such as links and instruction to help you resolve alerts triggered by this monitor. {{Markdown}} supported. It will be enabled only if available for your organization. Please contact your Sumo Logic account team to learn more. | [optional] [default to ""]
**SloId** | Pointer to **string** | Identifier of the SLO definition for the monitor. This is only applicable for SLO type monitors. | [optional] 
**AutomatedPlaybookIds** | Pointer to **[]string** | The set of automated playbook ids for a monitor. | [optional] [default to []]

## Methods

### NewMonitorsLibraryMonitorUpdate

`func NewMonitorsLibraryMonitorUpdate(monitorType string, queries []MonitorQuery, triggers []TriggerCondition, ) *MonitorsLibraryMonitorUpdate`

NewMonitorsLibraryMonitorUpdate instantiates a new MonitorsLibraryMonitorUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorsLibraryMonitorUpdateWithDefaults

`func NewMonitorsLibraryMonitorUpdateWithDefaults() *MonitorsLibraryMonitorUpdate`

NewMonitorsLibraryMonitorUpdateWithDefaults instantiates a new MonitorsLibraryMonitorUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonitorType

`func (o *MonitorsLibraryMonitorUpdate) GetMonitorType() string`

GetMonitorType returns the MonitorType field if non-nil, zero value otherwise.

### GetMonitorTypeOk

`func (o *MonitorsLibraryMonitorUpdate) GetMonitorTypeOk() (*string, bool)`

GetMonitorTypeOk returns a tuple with the MonitorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorType

`func (o *MonitorsLibraryMonitorUpdate) SetMonitorType(v string)`

SetMonitorType sets MonitorType field to given value.


### GetEvaluationDelay

`func (o *MonitorsLibraryMonitorUpdate) GetEvaluationDelay() string`

GetEvaluationDelay returns the EvaluationDelay field if non-nil, zero value otherwise.

### GetEvaluationDelayOk

`func (o *MonitorsLibraryMonitorUpdate) GetEvaluationDelayOk() (*string, bool)`

GetEvaluationDelayOk returns a tuple with the EvaluationDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationDelay

`func (o *MonitorsLibraryMonitorUpdate) SetEvaluationDelay(v string)`

SetEvaluationDelay sets EvaluationDelay field to given value.

### HasEvaluationDelay

`func (o *MonitorsLibraryMonitorUpdate) HasEvaluationDelay() bool`

HasEvaluationDelay returns a boolean if a field has been set.

### GetAlertName

`func (o *MonitorsLibraryMonitorUpdate) GetAlertName() string`

GetAlertName returns the AlertName field if non-nil, zero value otherwise.

### GetAlertNameOk

`func (o *MonitorsLibraryMonitorUpdate) GetAlertNameOk() (*string, bool)`

GetAlertNameOk returns a tuple with the AlertName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertName

`func (o *MonitorsLibraryMonitorUpdate) SetAlertName(v string)`

SetAlertName sets AlertName field to given value.

### HasAlertName

`func (o *MonitorsLibraryMonitorUpdate) HasAlertName() bool`

HasAlertName returns a boolean if a field has been set.

### GetRunAs

`func (o *MonitorsLibraryMonitorUpdate) GetRunAs() RunAs`

GetRunAs returns the RunAs field if non-nil, zero value otherwise.

### GetRunAsOk

`func (o *MonitorsLibraryMonitorUpdate) GetRunAsOk() (*RunAs, bool)`

GetRunAsOk returns a tuple with the RunAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunAs

`func (o *MonitorsLibraryMonitorUpdate) SetRunAs(v RunAs)`

SetRunAs sets RunAs field to given value.

### HasRunAs

`func (o *MonitorsLibraryMonitorUpdate) HasRunAs() bool`

HasRunAs returns a boolean if a field has been set.

### GetNotificationGroupFields

`func (o *MonitorsLibraryMonitorUpdate) GetNotificationGroupFields() []string`

GetNotificationGroupFields returns the NotificationGroupFields field if non-nil, zero value otherwise.

### GetNotificationGroupFieldsOk

`func (o *MonitorsLibraryMonitorUpdate) GetNotificationGroupFieldsOk() (*[]string, bool)`

GetNotificationGroupFieldsOk returns a tuple with the NotificationGroupFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationGroupFields

`func (o *MonitorsLibraryMonitorUpdate) SetNotificationGroupFields(v []string)`

SetNotificationGroupFields sets NotificationGroupFields field to given value.

### HasNotificationGroupFields

`func (o *MonitorsLibraryMonitorUpdate) HasNotificationGroupFields() bool`

HasNotificationGroupFields returns a boolean if a field has been set.

### GetQueries

`func (o *MonitorsLibraryMonitorUpdate) GetQueries() []MonitorQuery`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *MonitorsLibraryMonitorUpdate) GetQueriesOk() (*[]MonitorQuery, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *MonitorsLibraryMonitorUpdate) SetQueries(v []MonitorQuery)`

SetQueries sets Queries field to given value.


### GetTriggers

`func (o *MonitorsLibraryMonitorUpdate) GetTriggers() []TriggerCondition`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *MonitorsLibraryMonitorUpdate) GetTriggersOk() (*[]TriggerCondition, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *MonitorsLibraryMonitorUpdate) SetTriggers(v []TriggerCondition)`

SetTriggers sets Triggers field to given value.


### GetNotifications

`func (o *MonitorsLibraryMonitorUpdate) GetNotifications() []MonitorNotification`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *MonitorsLibraryMonitorUpdate) GetNotificationsOk() (*[]MonitorNotification, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *MonitorsLibraryMonitorUpdate) SetNotifications(v []MonitorNotification)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *MonitorsLibraryMonitorUpdate) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetIsDisabled

`func (o *MonitorsLibraryMonitorUpdate) GetIsDisabled() bool`

GetIsDisabled returns the IsDisabled field if non-nil, zero value otherwise.

### GetIsDisabledOk

`func (o *MonitorsLibraryMonitorUpdate) GetIsDisabledOk() (*bool, bool)`

GetIsDisabledOk returns a tuple with the IsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisabled

`func (o *MonitorsLibraryMonitorUpdate) SetIsDisabled(v bool)`

SetIsDisabled sets IsDisabled field to given value.

### HasIsDisabled

`func (o *MonitorsLibraryMonitorUpdate) HasIsDisabled() bool`

HasIsDisabled returns a boolean if a field has been set.

### GetGroupNotifications

`func (o *MonitorsLibraryMonitorUpdate) GetGroupNotifications() bool`

GetGroupNotifications returns the GroupNotifications field if non-nil, zero value otherwise.

### GetGroupNotificationsOk

`func (o *MonitorsLibraryMonitorUpdate) GetGroupNotificationsOk() (*bool, bool)`

GetGroupNotificationsOk returns a tuple with the GroupNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupNotifications

`func (o *MonitorsLibraryMonitorUpdate) SetGroupNotifications(v bool)`

SetGroupNotifications sets GroupNotifications field to given value.

### HasGroupNotifications

`func (o *MonitorsLibraryMonitorUpdate) HasGroupNotifications() bool`

HasGroupNotifications returns a boolean if a field has been set.

### GetPlaybook

`func (o *MonitorsLibraryMonitorUpdate) GetPlaybook() string`

GetPlaybook returns the Playbook field if non-nil, zero value otherwise.

### GetPlaybookOk

`func (o *MonitorsLibraryMonitorUpdate) GetPlaybookOk() (*string, bool)`

GetPlaybookOk returns a tuple with the Playbook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaybook

`func (o *MonitorsLibraryMonitorUpdate) SetPlaybook(v string)`

SetPlaybook sets Playbook field to given value.

### HasPlaybook

`func (o *MonitorsLibraryMonitorUpdate) HasPlaybook() bool`

HasPlaybook returns a boolean if a field has been set.

### GetSloId

`func (o *MonitorsLibraryMonitorUpdate) GetSloId() string`

GetSloId returns the SloId field if non-nil, zero value otherwise.

### GetSloIdOk

`func (o *MonitorsLibraryMonitorUpdate) GetSloIdOk() (*string, bool)`

GetSloIdOk returns a tuple with the SloId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSloId

`func (o *MonitorsLibraryMonitorUpdate) SetSloId(v string)`

SetSloId sets SloId field to given value.

### HasSloId

`func (o *MonitorsLibraryMonitorUpdate) HasSloId() bool`

HasSloId returns a boolean if a field has been set.

### GetAutomatedPlaybookIds

`func (o *MonitorsLibraryMonitorUpdate) GetAutomatedPlaybookIds() []string`

GetAutomatedPlaybookIds returns the AutomatedPlaybookIds field if non-nil, zero value otherwise.

### GetAutomatedPlaybookIdsOk

`func (o *MonitorsLibraryMonitorUpdate) GetAutomatedPlaybookIdsOk() (*[]string, bool)`

GetAutomatedPlaybookIdsOk returns a tuple with the AutomatedPlaybookIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomatedPlaybookIds

`func (o *MonitorsLibraryMonitorUpdate) SetAutomatedPlaybookIds(v []string)`

SetAutomatedPlaybookIds sets AutomatedPlaybookIds field to given value.

### HasAutomatedPlaybookIds

`func (o *MonitorsLibraryMonitorUpdate) HasAutomatedPlaybookIds() bool`

HasAutomatedPlaybookIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


