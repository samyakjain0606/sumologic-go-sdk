# DashboardTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Id of the dashboard. | 
**PanelToSessionIdMap** | Pointer to **map[string]string** | A map of panel to session id. The session id will be used to fetch data of the panel for the report. If not specified, a new session id will be created for the panel.  | [optional] 
**TimeRange** | Pointer to [**ResolvableTimeRange**](ResolvableTimeRange.md) |  | [optional] 
**VariableValues** | Pointer to [**VariablesValuesData**](VariablesValuesData.md) |  | [optional] 

## Methods

### NewDashboardTemplate

`func NewDashboardTemplate(id string, ) *DashboardTemplate`

NewDashboardTemplate instantiates a new DashboardTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardTemplateWithDefaults

`func NewDashboardTemplateWithDefaults() *DashboardTemplate`

NewDashboardTemplateWithDefaults instantiates a new DashboardTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DashboardTemplate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DashboardTemplate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DashboardTemplate) SetId(v string)`

SetId sets Id field to given value.


### GetPanelToSessionIdMap

`func (o *DashboardTemplate) GetPanelToSessionIdMap() map[string]string`

GetPanelToSessionIdMap returns the PanelToSessionIdMap field if non-nil, zero value otherwise.

### GetPanelToSessionIdMapOk

`func (o *DashboardTemplate) GetPanelToSessionIdMapOk() (*map[string]string, bool)`

GetPanelToSessionIdMapOk returns a tuple with the PanelToSessionIdMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPanelToSessionIdMap

`func (o *DashboardTemplate) SetPanelToSessionIdMap(v map[string]string)`

SetPanelToSessionIdMap sets PanelToSessionIdMap field to given value.

### HasPanelToSessionIdMap

`func (o *DashboardTemplate) HasPanelToSessionIdMap() bool`

HasPanelToSessionIdMap returns a boolean if a field has been set.

### GetTimeRange

`func (o *DashboardTemplate) GetTimeRange() ResolvableTimeRange`

GetTimeRange returns the TimeRange field if non-nil, zero value otherwise.

### GetTimeRangeOk

`func (o *DashboardTemplate) GetTimeRangeOk() (*ResolvableTimeRange, bool)`

GetTimeRangeOk returns a tuple with the TimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRange

`func (o *DashboardTemplate) SetTimeRange(v ResolvableTimeRange)`

SetTimeRange sets TimeRange field to given value.

### HasTimeRange

`func (o *DashboardTemplate) HasTimeRange() bool`

HasTimeRange returns a boolean if a field has been set.

### GetVariableValues

`func (o *DashboardTemplate) GetVariableValues() VariablesValuesData`

GetVariableValues returns the VariableValues field if non-nil, zero value otherwise.

### GetVariableValuesOk

`func (o *DashboardTemplate) GetVariableValuesOk() (*VariablesValuesData, bool)`

GetVariableValuesOk returns a tuple with the VariableValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableValues

`func (o *DashboardTemplate) SetVariableValues(v VariablesValuesData)`

SetVariableValues sets VariableValues field to given value.

### HasVariableValues

`func (o *DashboardTemplate) HasVariableValues() bool`

HasVariableValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


