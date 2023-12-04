# Dashboard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | Title of the dashboard. | 
**Description** | Pointer to **string** | Description of the dashboard. | [optional] 
**FolderId** | Pointer to **string** | The identifier of the folder to save the dashboard in. By default it is saved in your personal folder.  | [optional] 
**TopologyLabelMap** | Pointer to [**TopologyLabelMap**](TopologyLabelMap.md) |  | [optional] 
**Domain** | Pointer to **string** | If set denotes that the dashboard concerns a given domain (e.g. &#x60;aws&#x60;, &#x60;k8s&#x60;, &#x60;app&#x60;). | [optional] [default to ""]
**Hierarchies** | Pointer to **[]string** | If set to non-empty array denotes that the dashboard concerns given hierarchies. | [optional] [default to []]
**RefreshInterval** | Pointer to **int32** | Interval of time (in seconds) to automatically refresh the dashboard. A value of 0 means we never automatically refresh the dashboard. Allowed values are &#x60;0&#x60;, &#x60;30&#x60;, &#x60;60&#x60;, &#x60;120&#x60;, &#x60;300&#x60;, &#x60;900&#x60;, &#x60;1800&#x60;, &#x60;3600&#x60;, &#x60;7200&#x60;, &#x60;86400&#x60;.  | [optional] 
**TimeRange** | [**ResolvableTimeRange**](ResolvableTimeRange.md) |  | 
**Panels** | Pointer to [**[]Panel**](Panel.md) | Panels in the dashboard. | [optional] 
**Layout** | Pointer to [**Layout**](Layout.md) |  | [optional] 
**Variables** | Pointer to [**[]Variable**](Variable.md) | Variables to apply to the panels. | [optional] 
**Theme** | Pointer to **string** | Theme for the dashboard. Either &#x60;Light&#x60; or &#x60;Dark&#x60;. | [optional] [default to "Light"]
**IsPublic** | Pointer to **bool** | Is the dashboard public | [optional] [default to false]
**HighlightViolations** | Pointer to **bool** | Whether to highlight threshold violations. | [optional] [default to false]
**Id** | Pointer to **string** | Unique identifier for the dashboard. This id is used to get detailed information about the dashboard, such as panels, variables and the layout.  | [optional] 
**ContentId** | Pointer to **string** | Content identifier for the dashboard. This id is used to connect to the Sumo Content Library and get general metadata about the dashboard. Use this id if you want to search for dashboards in Sumo folders.  | [optional] 
**ScheduleId** | Pointer to **string** | Scheduled report identifier for the dashboard. Only most recently modified report schedule is rerun per dashboard. This id is used to manage the schedule details through the scheduled report API.  | [optional] 

## Methods

### NewDashboard

`func NewDashboard(title string, timeRange ResolvableTimeRange, ) *Dashboard`

NewDashboard instantiates a new Dashboard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardWithDefaults

`func NewDashboardWithDefaults() *Dashboard`

NewDashboardWithDefaults instantiates a new Dashboard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *Dashboard) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Dashboard) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Dashboard) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *Dashboard) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Dashboard) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Dashboard) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Dashboard) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFolderId

`func (o *Dashboard) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *Dashboard) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *Dashboard) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *Dashboard) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetTopologyLabelMap

`func (o *Dashboard) GetTopologyLabelMap() TopologyLabelMap`

GetTopologyLabelMap returns the TopologyLabelMap field if non-nil, zero value otherwise.

### GetTopologyLabelMapOk

`func (o *Dashboard) GetTopologyLabelMapOk() (*TopologyLabelMap, bool)`

GetTopologyLabelMapOk returns a tuple with the TopologyLabelMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologyLabelMap

`func (o *Dashboard) SetTopologyLabelMap(v TopologyLabelMap)`

SetTopologyLabelMap sets TopologyLabelMap field to given value.

### HasTopologyLabelMap

`func (o *Dashboard) HasTopologyLabelMap() bool`

HasTopologyLabelMap returns a boolean if a field has been set.

### GetDomain

`func (o *Dashboard) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *Dashboard) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *Dashboard) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *Dashboard) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetHierarchies

`func (o *Dashboard) GetHierarchies() []string`

GetHierarchies returns the Hierarchies field if non-nil, zero value otherwise.

### GetHierarchiesOk

`func (o *Dashboard) GetHierarchiesOk() (*[]string, bool)`

GetHierarchiesOk returns a tuple with the Hierarchies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHierarchies

`func (o *Dashboard) SetHierarchies(v []string)`

SetHierarchies sets Hierarchies field to given value.

### HasHierarchies

`func (o *Dashboard) HasHierarchies() bool`

HasHierarchies returns a boolean if a field has been set.

### GetRefreshInterval

`func (o *Dashboard) GetRefreshInterval() int32`

GetRefreshInterval returns the RefreshInterval field if non-nil, zero value otherwise.

### GetRefreshIntervalOk

`func (o *Dashboard) GetRefreshIntervalOk() (*int32, bool)`

GetRefreshIntervalOk returns a tuple with the RefreshInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshInterval

`func (o *Dashboard) SetRefreshInterval(v int32)`

SetRefreshInterval sets RefreshInterval field to given value.

### HasRefreshInterval

`func (o *Dashboard) HasRefreshInterval() bool`

HasRefreshInterval returns a boolean if a field has been set.

### GetTimeRange

`func (o *Dashboard) GetTimeRange() ResolvableTimeRange`

GetTimeRange returns the TimeRange field if non-nil, zero value otherwise.

### GetTimeRangeOk

`func (o *Dashboard) GetTimeRangeOk() (*ResolvableTimeRange, bool)`

GetTimeRangeOk returns a tuple with the TimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRange

`func (o *Dashboard) SetTimeRange(v ResolvableTimeRange)`

SetTimeRange sets TimeRange field to given value.


### GetPanels

`func (o *Dashboard) GetPanels() []Panel`

GetPanels returns the Panels field if non-nil, zero value otherwise.

### GetPanelsOk

`func (o *Dashboard) GetPanelsOk() (*[]Panel, bool)`

GetPanelsOk returns a tuple with the Panels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPanels

`func (o *Dashboard) SetPanels(v []Panel)`

SetPanels sets Panels field to given value.

### HasPanels

`func (o *Dashboard) HasPanels() bool`

HasPanels returns a boolean if a field has been set.

### GetLayout

`func (o *Dashboard) GetLayout() Layout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *Dashboard) GetLayoutOk() (*Layout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *Dashboard) SetLayout(v Layout)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *Dashboard) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### GetVariables

`func (o *Dashboard) GetVariables() []Variable`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *Dashboard) GetVariablesOk() (*[]Variable, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *Dashboard) SetVariables(v []Variable)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *Dashboard) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### GetTheme

`func (o *Dashboard) GetTheme() string`

GetTheme returns the Theme field if non-nil, zero value otherwise.

### GetThemeOk

`func (o *Dashboard) GetThemeOk() (*string, bool)`

GetThemeOk returns a tuple with the Theme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheme

`func (o *Dashboard) SetTheme(v string)`

SetTheme sets Theme field to given value.

### HasTheme

`func (o *Dashboard) HasTheme() bool`

HasTheme returns a boolean if a field has been set.

### GetIsPublic

`func (o *Dashboard) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *Dashboard) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *Dashboard) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *Dashboard) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetHighlightViolations

`func (o *Dashboard) GetHighlightViolations() bool`

GetHighlightViolations returns the HighlightViolations field if non-nil, zero value otherwise.

### GetHighlightViolationsOk

`func (o *Dashboard) GetHighlightViolationsOk() (*bool, bool)`

GetHighlightViolationsOk returns a tuple with the HighlightViolations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlightViolations

`func (o *Dashboard) SetHighlightViolations(v bool)`

SetHighlightViolations sets HighlightViolations field to given value.

### HasHighlightViolations

`func (o *Dashboard) HasHighlightViolations() bool`

HasHighlightViolations returns a boolean if a field has been set.

### GetId

`func (o *Dashboard) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Dashboard) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Dashboard) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Dashboard) HasId() bool`

HasId returns a boolean if a field has been set.

### GetContentId

`func (o *Dashboard) GetContentId() string`

GetContentId returns the ContentId field if non-nil, zero value otherwise.

### GetContentIdOk

`func (o *Dashboard) GetContentIdOk() (*string, bool)`

GetContentIdOk returns a tuple with the ContentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentId

`func (o *Dashboard) SetContentId(v string)`

SetContentId sets ContentId field to given value.

### HasContentId

`func (o *Dashboard) HasContentId() bool`

HasContentId returns a boolean if a field has been set.

### GetScheduleId

`func (o *Dashboard) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *Dashboard) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *Dashboard) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.

### HasScheduleId

`func (o *Dashboard) HasScheduleId() bool`

HasScheduleId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


