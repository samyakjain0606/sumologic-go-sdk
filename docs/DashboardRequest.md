# DashboardRequest

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

## Methods

### NewDashboardRequest

`func NewDashboardRequest(title string, timeRange ResolvableTimeRange, ) *DashboardRequest`

NewDashboardRequest instantiates a new DashboardRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardRequestWithDefaults

`func NewDashboardRequestWithDefaults() *DashboardRequest`

NewDashboardRequestWithDefaults instantiates a new DashboardRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *DashboardRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DashboardRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DashboardRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *DashboardRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DashboardRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DashboardRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DashboardRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFolderId

`func (o *DashboardRequest) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *DashboardRequest) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *DashboardRequest) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *DashboardRequest) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetTopologyLabelMap

`func (o *DashboardRequest) GetTopologyLabelMap() TopologyLabelMap`

GetTopologyLabelMap returns the TopologyLabelMap field if non-nil, zero value otherwise.

### GetTopologyLabelMapOk

`func (o *DashboardRequest) GetTopologyLabelMapOk() (*TopologyLabelMap, bool)`

GetTopologyLabelMapOk returns a tuple with the TopologyLabelMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologyLabelMap

`func (o *DashboardRequest) SetTopologyLabelMap(v TopologyLabelMap)`

SetTopologyLabelMap sets TopologyLabelMap field to given value.

### HasTopologyLabelMap

`func (o *DashboardRequest) HasTopologyLabelMap() bool`

HasTopologyLabelMap returns a boolean if a field has been set.

### GetDomain

`func (o *DashboardRequest) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *DashboardRequest) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *DashboardRequest) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *DashboardRequest) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetHierarchies

`func (o *DashboardRequest) GetHierarchies() []string`

GetHierarchies returns the Hierarchies field if non-nil, zero value otherwise.

### GetHierarchiesOk

`func (o *DashboardRequest) GetHierarchiesOk() (*[]string, bool)`

GetHierarchiesOk returns a tuple with the Hierarchies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHierarchies

`func (o *DashboardRequest) SetHierarchies(v []string)`

SetHierarchies sets Hierarchies field to given value.

### HasHierarchies

`func (o *DashboardRequest) HasHierarchies() bool`

HasHierarchies returns a boolean if a field has been set.

### GetRefreshInterval

`func (o *DashboardRequest) GetRefreshInterval() int32`

GetRefreshInterval returns the RefreshInterval field if non-nil, zero value otherwise.

### GetRefreshIntervalOk

`func (o *DashboardRequest) GetRefreshIntervalOk() (*int32, bool)`

GetRefreshIntervalOk returns a tuple with the RefreshInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshInterval

`func (o *DashboardRequest) SetRefreshInterval(v int32)`

SetRefreshInterval sets RefreshInterval field to given value.

### HasRefreshInterval

`func (o *DashboardRequest) HasRefreshInterval() bool`

HasRefreshInterval returns a boolean if a field has been set.

### GetTimeRange

`func (o *DashboardRequest) GetTimeRange() ResolvableTimeRange`

GetTimeRange returns the TimeRange field if non-nil, zero value otherwise.

### GetTimeRangeOk

`func (o *DashboardRequest) GetTimeRangeOk() (*ResolvableTimeRange, bool)`

GetTimeRangeOk returns a tuple with the TimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRange

`func (o *DashboardRequest) SetTimeRange(v ResolvableTimeRange)`

SetTimeRange sets TimeRange field to given value.


### GetPanels

`func (o *DashboardRequest) GetPanels() []Panel`

GetPanels returns the Panels field if non-nil, zero value otherwise.

### GetPanelsOk

`func (o *DashboardRequest) GetPanelsOk() (*[]Panel, bool)`

GetPanelsOk returns a tuple with the Panels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPanels

`func (o *DashboardRequest) SetPanels(v []Panel)`

SetPanels sets Panels field to given value.

### HasPanels

`func (o *DashboardRequest) HasPanels() bool`

HasPanels returns a boolean if a field has been set.

### GetLayout

`func (o *DashboardRequest) GetLayout() Layout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *DashboardRequest) GetLayoutOk() (*Layout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *DashboardRequest) SetLayout(v Layout)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *DashboardRequest) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### GetVariables

`func (o *DashboardRequest) GetVariables() []Variable`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *DashboardRequest) GetVariablesOk() (*[]Variable, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *DashboardRequest) SetVariables(v []Variable)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *DashboardRequest) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### GetTheme

`func (o *DashboardRequest) GetTheme() string`

GetTheme returns the Theme field if non-nil, zero value otherwise.

### GetThemeOk

`func (o *DashboardRequest) GetThemeOk() (*string, bool)`

GetThemeOk returns a tuple with the Theme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheme

`func (o *DashboardRequest) SetTheme(v string)`

SetTheme sets Theme field to given value.

### HasTheme

`func (o *DashboardRequest) HasTheme() bool`

HasTheme returns a boolean if a field has been set.

### GetIsPublic

`func (o *DashboardRequest) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *DashboardRequest) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *DashboardRequest) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *DashboardRequest) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetHighlightViolations

`func (o *DashboardRequest) GetHighlightViolations() bool`

GetHighlightViolations returns the HighlightViolations field if non-nil, zero value otherwise.

### GetHighlightViolationsOk

`func (o *DashboardRequest) GetHighlightViolationsOk() (*bool, bool)`

GetHighlightViolationsOk returns a tuple with the HighlightViolations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlightViolations

`func (o *DashboardRequest) SetHighlightViolations(v bool)`

SetHighlightViolations sets HighlightViolations field to given value.

### HasHighlightViolations

`func (o *DashboardRequest) HasHighlightViolations() bool`

HasHighlightViolations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


