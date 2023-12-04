# MonitorTemplatesLibraryMonitorTemplateResponse

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

### NewMonitorTemplatesLibraryMonitorTemplateResponse

`func NewMonitorTemplatesLibraryMonitorTemplateResponse(monitorType string, queries []MonitorQuery, triggers []TriggerCondition, ) *MonitorTemplatesLibraryMonitorTemplateResponse`

NewMonitorTemplatesLibraryMonitorTemplateResponse instantiates a new MonitorTemplatesLibraryMonitorTemplateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorTemplatesLibraryMonitorTemplateResponseWithDefaults

`func NewMonitorTemplatesLibraryMonitorTemplateResponseWithDefaults() *MonitorTemplatesLibraryMonitorTemplateResponse`

NewMonitorTemplatesLibraryMonitorTemplateResponseWithDefaults instantiates a new MonitorTemplatesLibraryMonitorTemplateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonitorType

`func (o *MonitorTemplatesLibraryMonitorTemplateResponse) GetMonitorType() string`

GetMonitorType returns the MonitorType field if non-nil, zero value otherwise.

### GetMonitorTypeOk

`func (o *MonitorTemplatesLibraryMonitorTemplateResponse) GetMonitorTypeOk() (*string, bool)`

GetMonitorTypeOk returns a tuple with the MonitorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorType

`func (o *MonitorTemplatesLibraryMonitorTemplateResponse) SetMonitorType(v string)`

SetMonitorType sets MonitorType field to given value.


### GetEvaluationDelay

`func (o *MonitorTemplatesLibraryMonitorTemplateResponse) GetEvaluationDelay() string`

GetEvaluationDelay returns the EvaluationDelay field if non-nil, zero value otherwise.

### GetEvaluationDelayOk

`func (o *MonitorTemplatesLibraryMonitorTemplateResponse) GetEvaluationDelayOk() (*string, bool)`

GetEvaluationDelayOk returns a tuple with the EvaluationDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationDelay

`func (o *MonitorTemplatesLibraryMonitorTemplateResponse) SetEvaluationDelay(v string)`

SetEvaluationDelay sets EvaluationDelay field to given value.

### HasEvaluationDelay

`func (o *MonitorTemplatesLibraryMonitorTemplateResponse) HasEvaluationDelay() bool`

HasEvaluationDelay returns a boolean if a field has been set.

### GetAlertName

`func (o *MonitorTemplatesLibraryMonitorTemplateResponse) GetAlertName() string`

GetAlertName returns the AlertName field if non-nil, zero value otherwise.

### GetAlertNameOk

`func (o *MonitorTemplatesLibraryMonitorTemplateResponse) GetAlertNameOk() (*string, bool)`

GetAlertNameOk returns a tuple with the AlertName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertName

`func (o *MonitorTemplatesLibraryMonitorTemplateResponse) SetAlertName(v string)`

SetAlertName sets AlertName field to given value.

### HasAlertName

`func (o *MonitorTemplatesLibraryMonitorTemplateResponse) HasAlertName() bool`

HasAlertName returns a boolean if a field has been set.

### GetQueries

`func (o *MonitorTemplatesLibraryMonitorTemplateResponse) GetQueries() []MonitorQuery`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *MonitorTemplatesLibraryMonitorTemplateResponse) GetQueriesOk() (*[]MonitorQuery, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *MonitorTemplatesLibraryMonitorTemplateResponse) SetQueries(v []MonitorQuery)`

SetQueries sets Queries field to given value.


### GetTriggers

`func (o *MonitorTemplatesLibraryMonitorTemplateResponse) GetTriggers() []TriggerCondition`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *MonitorTemplatesLibraryMonitorTemplateResponse) GetTriggersOk() (*[]TriggerCondition, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *MonitorTemplatesLibraryMonitorTemplateResponse) SetTriggers(v []TriggerCondition)`

SetTriggers sets Triggers field to given value.


### GetIsDisabled

`func (o *MonitorTemplatesLibraryMonitorTemplateResponse) GetIsDisabled() bool`

GetIsDisabled returns the IsDisabled field if non-nil, zero value otherwise.

### GetIsDisabledOk

`func (o *MonitorTemplatesLibraryMonitorTemplateResponse) GetIsDisabledOk() (*bool, bool)`

GetIsDisabledOk returns a tuple with the IsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisabled

`func (o *MonitorTemplatesLibraryMonitorTemplateResponse) SetIsDisabled(v bool)`

SetIsDisabled sets IsDisabled field to given value.

### HasIsDisabled

`func (o *MonitorTemplatesLibraryMonitorTemplateResponse) HasIsDisabled() bool`

HasIsDisabled returns a boolean if a field has been set.

### GetGroupNotifications

`func (o *MonitorTemplatesLibraryMonitorTemplateResponse) GetGroupNotifications() bool`

GetGroupNotifications returns the GroupNotifications field if non-nil, zero value otherwise.

### GetGroupNotificationsOk

`func (o *MonitorTemplatesLibraryMonitorTemplateResponse) GetGroupNotificationsOk() (*bool, bool)`

GetGroupNotificationsOk returns a tuple with the GroupNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupNotifications

`func (o *MonitorTemplatesLibraryMonitorTemplateResponse) SetGroupNotifications(v bool)`

SetGroupNotifications sets GroupNotifications field to given value.

### HasGroupNotifications

`func (o *MonitorTemplatesLibraryMonitorTemplateResponse) HasGroupNotifications() bool`

HasGroupNotifications returns a boolean if a field has been set.

### GetPlaybook

`func (o *MonitorTemplatesLibraryMonitorTemplateResponse) GetPlaybook() string`

GetPlaybook returns the Playbook field if non-nil, zero value otherwise.

### GetPlaybookOk

`func (o *MonitorTemplatesLibraryMonitorTemplateResponse) GetPlaybookOk() (*string, bool)`

GetPlaybookOk returns a tuple with the Playbook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaybook

`func (o *MonitorTemplatesLibraryMonitorTemplateResponse) SetPlaybook(v string)`

SetPlaybook sets Playbook field to given value.

### HasPlaybook

`func (o *MonitorTemplatesLibraryMonitorTemplateResponse) HasPlaybook() bool`

HasPlaybook returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


