# RoleDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the role. | 
**Description** | Pointer to **string** | Description of the role. | [optional] 
**LogAnalyticsFilter** | Pointer to **string** | A search filter which would be applied on partitions which belong to Log Analytics product area. Applicable with only &#x60;All&#x60; selectionType | [optional] 
**AuditDataFilter** | Pointer to **string** | A search filter which would be applied on partitions which belong to Audit Data product area. Help Doc : (https://help.sumologic.com/docs/manage/security/audit-index/). Applicable with only &#x60;All&#x60; selectionType | [optional] 
**SecurityDataFilter** | Pointer to **string** | A search filter which would be applied on partitions which belong to Security Data product area. Applicable with only &#x60;All&#x60; selectionType | [optional] 
**SelectionType** | Pointer to **string** | Describes the Permission Construct for the list of views in \&quot;selectedViews\&quot; parameter.  ### Valid Values are :    - &#x60;All&#x60; selectionType would allow access to all views in the org.   - &#x60;Allow&#x60; selectionType would allow access to specific views mentioned in \&quot;selectedViews\&quot; parameter.   - &#x60;Deny&#x60; selectionType would deny access to specific views mentioned in \&quot;selectedViews\&quot; parameter. | [optional] 
**SelectedViews** | Pointer to [**[]GetViewFilterDefinition**](GetViewFilterDefinition.md) | List of views which with specific view level filters in accordance to the selectionType chosen. | [optional] 
**Users** | Pointer to **[]string** | List of user identifiers to assign the role to. | [optional] 
**Capabilities** | Pointer to **[]string** | List of [capabilities](https://help.sumologic.com/Manage/Users-and-Roles/Manage-Roles/Role-Capabilities) associated with this role. Valid values are ### Data Management   - viewCollectors   - manageCollectors   - manageBudgets   - manageDataVolumeFeed   - viewFieldExtraction   - manageFieldExtractionRules   - manageS3DataForwarding   - manageContent   - manageApps   - dataVolumeIndex   - manageConnections   - viewScheduledViews   - manageScheduledViews   - viewPartitions   - managePartitions   - viewFields   - manageFields   - viewAccountOverview   - manageTokens   - downloadSearchResults  ### Entity management   - manageEntityTypeConfig  ### Metrics   - metricsTransformation   - metricsExtraction   - metricsRules  ### Security   - managePasswordPolicy   - ipAllowlisting   - createAccessKeys   - manageAccessKeys   - manageSupportAccountAccess   - manageAuditDataFeed   - manageSaml   - shareDashboardOutsideOrg   - manageOrgSettings   - changeDataAccessLevel  ### Dashboards   - shareDashboardWorld   - shareDashboardAllowlist  ### UserManagement   - manageUsersAndRoles  ### Observability   - searchAuditIndex   - auditEventIndex  ### Cloud SIEM Enterprise   - viewCse  ### Alerting   - viewMonitorsV2   - manageMonitorsV2   - viewAlerts | [optional] 
**AutofillDependencies** | Pointer to **bool** | Set this to true if you want to automatically append all missing capability requirements. If set to false an error will be thrown if any capabilities are missing their dependencies. | [optional] [default to true]

## Methods

### NewRoleDefinition

`func NewRoleDefinition(name string, ) *RoleDefinition`

NewRoleDefinition instantiates a new RoleDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleDefinitionWithDefaults

`func NewRoleDefinitionWithDefaults() *RoleDefinition`

NewRoleDefinitionWithDefaults instantiates a new RoleDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RoleDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RoleDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RoleDefinition) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *RoleDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RoleDefinition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RoleDefinition) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RoleDefinition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLogAnalyticsFilter

`func (o *RoleDefinition) GetLogAnalyticsFilter() string`

GetLogAnalyticsFilter returns the LogAnalyticsFilter field if non-nil, zero value otherwise.

### GetLogAnalyticsFilterOk

`func (o *RoleDefinition) GetLogAnalyticsFilterOk() (*string, bool)`

GetLogAnalyticsFilterOk returns a tuple with the LogAnalyticsFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogAnalyticsFilter

`func (o *RoleDefinition) SetLogAnalyticsFilter(v string)`

SetLogAnalyticsFilter sets LogAnalyticsFilter field to given value.

### HasLogAnalyticsFilter

`func (o *RoleDefinition) HasLogAnalyticsFilter() bool`

HasLogAnalyticsFilter returns a boolean if a field has been set.

### GetAuditDataFilter

`func (o *RoleDefinition) GetAuditDataFilter() string`

GetAuditDataFilter returns the AuditDataFilter field if non-nil, zero value otherwise.

### GetAuditDataFilterOk

`func (o *RoleDefinition) GetAuditDataFilterOk() (*string, bool)`

GetAuditDataFilterOk returns a tuple with the AuditDataFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditDataFilter

`func (o *RoleDefinition) SetAuditDataFilter(v string)`

SetAuditDataFilter sets AuditDataFilter field to given value.

### HasAuditDataFilter

`func (o *RoleDefinition) HasAuditDataFilter() bool`

HasAuditDataFilter returns a boolean if a field has been set.

### GetSecurityDataFilter

`func (o *RoleDefinition) GetSecurityDataFilter() string`

GetSecurityDataFilter returns the SecurityDataFilter field if non-nil, zero value otherwise.

### GetSecurityDataFilterOk

`func (o *RoleDefinition) GetSecurityDataFilterOk() (*string, bool)`

GetSecurityDataFilterOk returns a tuple with the SecurityDataFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityDataFilter

`func (o *RoleDefinition) SetSecurityDataFilter(v string)`

SetSecurityDataFilter sets SecurityDataFilter field to given value.

### HasSecurityDataFilter

`func (o *RoleDefinition) HasSecurityDataFilter() bool`

HasSecurityDataFilter returns a boolean if a field has been set.

### GetSelectionType

`func (o *RoleDefinition) GetSelectionType() string`

GetSelectionType returns the SelectionType field if non-nil, zero value otherwise.

### GetSelectionTypeOk

`func (o *RoleDefinition) GetSelectionTypeOk() (*string, bool)`

GetSelectionTypeOk returns a tuple with the SelectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectionType

`func (o *RoleDefinition) SetSelectionType(v string)`

SetSelectionType sets SelectionType field to given value.

### HasSelectionType

`func (o *RoleDefinition) HasSelectionType() bool`

HasSelectionType returns a boolean if a field has been set.

### GetSelectedViews

`func (o *RoleDefinition) GetSelectedViews() []GetViewFilterDefinition`

GetSelectedViews returns the SelectedViews field if non-nil, zero value otherwise.

### GetSelectedViewsOk

`func (o *RoleDefinition) GetSelectedViewsOk() (*[]GetViewFilterDefinition, bool)`

GetSelectedViewsOk returns a tuple with the SelectedViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedViews

`func (o *RoleDefinition) SetSelectedViews(v []GetViewFilterDefinition)`

SetSelectedViews sets SelectedViews field to given value.

### HasSelectedViews

`func (o *RoleDefinition) HasSelectedViews() bool`

HasSelectedViews returns a boolean if a field has been set.

### GetUsers

`func (o *RoleDefinition) GetUsers() []string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *RoleDefinition) GetUsersOk() (*[]string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *RoleDefinition) SetUsers(v []string)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *RoleDefinition) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetCapabilities

`func (o *RoleDefinition) GetCapabilities() []string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *RoleDefinition) GetCapabilitiesOk() (*[]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *RoleDefinition) SetCapabilities(v []string)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *RoleDefinition) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetAutofillDependencies

`func (o *RoleDefinition) GetAutofillDependencies() bool`

GetAutofillDependencies returns the AutofillDependencies field if non-nil, zero value otherwise.

### GetAutofillDependenciesOk

`func (o *RoleDefinition) GetAutofillDependenciesOk() (*bool, bool)`

GetAutofillDependenciesOk returns a tuple with the AutofillDependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutofillDependencies

`func (o *RoleDefinition) SetAutofillDependencies(v bool)`

SetAutofillDependencies sets AutofillDependencies field to given value.

### HasAutofillDependencies

`func (o *RoleDefinition) HasAutofillDependencies() bool`

HasAutofillDependencies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


