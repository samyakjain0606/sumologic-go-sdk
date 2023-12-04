# MonitorTemplatesLibraryMonitorTemplateExport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MonitorType** | **string** | The type of monitor template. Valid values:   1. &#x60;Logs&#x60;: A logs query monitor template.   2. &#x60;Metrics&#x60;: A metrics query monitor template.   3. &#x60;Slo&#x60;: A SLO based monitor template. | 
**EvaluationDelay** | Pointer to **string** | The delay duration for evaluating the monitor (relative to current time). The timerange of monitor will be shifted in the past by this delay time. | [optional] [default to "0m"]
**AlertName** | Pointer to **string** | The name of the alert(s) triggered from the monitor created based on the template. Monitor name will be used if not specified. | [optional] 
**Queries** | [**[]MonitorQuery**](MonitorQuery.md) | All queries from the monitor. | 
**Triggers** | [**[]TriggerCondition**](TriggerCondition.md) | Defines the conditions of when to send notifications. | 
**IsDisabled** | Pointer to **bool** | Whether or not the monitor template is disabled. | [optional] [default to false]
**GroupNotifications** | Pointer to **bool** | Whether or not to group notifications for individual items that meet the trigger condition. | [optional] [default to true]
**Playbook** | Pointer to **string** | Notes such as links and instruction to help you resolve alerts triggered by this monitor template. {{Markdown}} supported. It will be enabled only if available for your organization. Please contact your Sumo Logic account team to learn more. | [optional] [default to ""]

## Methods

### NewMonitorTemplatesLibraryMonitorTemplateExport

`func NewMonitorTemplatesLibraryMonitorTemplateExport(monitorType string, queries []MonitorQuery, triggers []TriggerCondition, ) *MonitorTemplatesLibraryMonitorTemplateExport`

NewMonitorTemplatesLibraryMonitorTemplateExport instantiates a new MonitorTemplatesLibraryMonitorTemplateExport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorTemplatesLibraryMonitorTemplateExportWithDefaults

`func NewMonitorTemplatesLibraryMonitorTemplateExportWithDefaults() *MonitorTemplatesLibraryMonitorTemplateExport`

NewMonitorTemplatesLibraryMonitorTemplateExportWithDefaults instantiates a new MonitorTemplatesLibraryMonitorTemplateExport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonitorType

`func (o *MonitorTemplatesLibraryMonitorTemplateExport) GetMonitorType() string`

GetMonitorType returns the MonitorType field if non-nil, zero value otherwise.

### GetMonitorTypeOk

`func (o *MonitorTemplatesLibraryMonitorTemplateExport) GetMonitorTypeOk() (*string, bool)`

GetMonitorTypeOk returns a tuple with the MonitorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorType

`func (o *MonitorTemplatesLibraryMonitorTemplateExport) SetMonitorType(v string)`

SetMonitorType sets MonitorType field to given value.


### GetEvaluationDelay

`func (o *MonitorTemplatesLibraryMonitorTemplateExport) GetEvaluationDelay() string`

GetEvaluationDelay returns the EvaluationDelay field if non-nil, zero value otherwise.

### GetEvaluationDelayOk

`func (o *MonitorTemplatesLibraryMonitorTemplateExport) GetEvaluationDelayOk() (*string, bool)`

GetEvaluationDelayOk returns a tuple with the EvaluationDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationDelay

`func (o *MonitorTemplatesLibraryMonitorTemplateExport) SetEvaluationDelay(v string)`

SetEvaluationDelay sets EvaluationDelay field to given value.

### HasEvaluationDelay

`func (o *MonitorTemplatesLibraryMonitorTemplateExport) HasEvaluationDelay() bool`

HasEvaluationDelay returns a boolean if a field has been set.

### GetAlertName

`func (o *MonitorTemplatesLibraryMonitorTemplateExport) GetAlertName() string`

GetAlertName returns the AlertName field if non-nil, zero value otherwise.

### GetAlertNameOk

`func (o *MonitorTemplatesLibraryMonitorTemplateExport) GetAlertNameOk() (*string, bool)`

GetAlertNameOk returns a tuple with the AlertName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertName

`func (o *MonitorTemplatesLibraryMonitorTemplateExport) SetAlertName(v string)`

SetAlertName sets AlertName field to given value.

### HasAlertName

`func (o *MonitorTemplatesLibraryMonitorTemplateExport) HasAlertName() bool`

HasAlertName returns a boolean if a field has been set.

### GetQueries

`func (o *MonitorTemplatesLibraryMonitorTemplateExport) GetQueries() []MonitorQuery`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *MonitorTemplatesLibraryMonitorTemplateExport) GetQueriesOk() (*[]MonitorQuery, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *MonitorTemplatesLibraryMonitorTemplateExport) SetQueries(v []MonitorQuery)`

SetQueries sets Queries field to given value.


### GetTriggers

`func (o *MonitorTemplatesLibraryMonitorTemplateExport) GetTriggers() []TriggerCondition`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *MonitorTemplatesLibraryMonitorTemplateExport) GetTriggersOk() (*[]TriggerCondition, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *MonitorTemplatesLibraryMonitorTemplateExport) SetTriggers(v []TriggerCondition)`

SetTriggers sets Triggers field to given value.


### GetIsDisabled

`func (o *MonitorTemplatesLibraryMonitorTemplateExport) GetIsDisabled() bool`

GetIsDisabled returns the IsDisabled field if non-nil, zero value otherwise.

### GetIsDisabledOk

`func (o *MonitorTemplatesLibraryMonitorTemplateExport) GetIsDisabledOk() (*bool, bool)`

GetIsDisabledOk returns a tuple with the IsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisabled

`func (o *MonitorTemplatesLibraryMonitorTemplateExport) SetIsDisabled(v bool)`

SetIsDisabled sets IsDisabled field to given value.

### HasIsDisabled

`func (o *MonitorTemplatesLibraryMonitorTemplateExport) HasIsDisabled() bool`

HasIsDisabled returns a boolean if a field has been set.

### GetGroupNotifications

`func (o *MonitorTemplatesLibraryMonitorTemplateExport) GetGroupNotifications() bool`

GetGroupNotifications returns the GroupNotifications field if non-nil, zero value otherwise.

### GetGroupNotificationsOk

`func (o *MonitorTemplatesLibraryMonitorTemplateExport) GetGroupNotificationsOk() (*bool, bool)`

GetGroupNotificationsOk returns a tuple with the GroupNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupNotifications

`func (o *MonitorTemplatesLibraryMonitorTemplateExport) SetGroupNotifications(v bool)`

SetGroupNotifications sets GroupNotifications field to given value.

### HasGroupNotifications

`func (o *MonitorTemplatesLibraryMonitorTemplateExport) HasGroupNotifications() bool`

HasGroupNotifications returns a boolean if a field has been set.

### GetPlaybook

`func (o *MonitorTemplatesLibraryMonitorTemplateExport) GetPlaybook() string`

GetPlaybook returns the Playbook field if non-nil, zero value otherwise.

### GetPlaybookOk

`func (o *MonitorTemplatesLibraryMonitorTemplateExport) GetPlaybookOk() (*string, bool)`

GetPlaybookOk returns a tuple with the Playbook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaybook

`func (o *MonitorTemplatesLibraryMonitorTemplateExport) SetPlaybook(v string)`

SetPlaybook sets Playbook field to given value.

### HasPlaybook

`func (o *MonitorTemplatesLibraryMonitorTemplateExport) HasPlaybook() bool`

HasPlaybook returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


