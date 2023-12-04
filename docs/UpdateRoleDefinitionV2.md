# UpdateRoleDefinitionV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the role. | 
**Description** | **string** | Description of the role. | 
**LogAnalyticsFilter** | **string** | A search filter which would be applied on partitions which belong to Log Analytics product area. Applicable with only &#x60;All&#x60; selectionType. | 
**AuditDataFilter** | **string** | A search filter which would be applied on partitions which belong to Audit Data product area. Help Doc : (https://help.sumologic.com/docs/manage/security/audit-index/). Applicable with only &#x60;All&#x60; selectionType. | 
**SecurityDataFilter** | **string** | A search filter which would be applied on partitions which belong to Security Data product area. Applicable with only &#x60;All&#x60; selectionType. | 
**SelectionType** | **string** | Describes the Permission Construct for the list of views in \&quot;selectedViews\&quot; parameter.  ### Valid Values are :    - &#x60;All&#x60; selectionType would allow access to all views in the org.   - &#x60;Allow&#x60; selectionType would allow access to specific views mentioned in \&quot;selectedViews\&quot; parameter.   - &#x60;Deny&#x60; selectionType would deny access to specific views mentioned in \&quot;selectedViews\&quot; parameter. | 
**SelectedViews** | [**[]ViewFilterDefinition**](ViewFilterDefinition.md) | List of views which with specific view level filters in accordance to the selectionType chosen. | 
**Users** | **[]string** | List of user identifiers to assign the role to. | 
**Capabilities** | **[]string** | List of [capabilities](https://help.sumologic.com/Manage/Users-and-Roles/Manage-Roles/Role-Capabilities) associated with this role. Valid values are ### Data Management   - viewCollectors   - manageCollectors   - manageBudgets   - manageDataVolumeFeed   - viewFieldExtraction   - manageFieldExtractionRules   - manageS3DataForwarding   - manageContent   - manageApps   - dataVolumeIndex   - manageConnections   - viewScheduledViews   - manageScheduledViews   - viewPartitions   - managePartitions   - viewFields   - manageFields   - viewAccountOverview   - manageTokens   - downloadSearchResults  ### Entity management   - manageEntityTypeConfig  ### Metrics   - metricsTransformation   - metricsExtraction   - metricsRules  ### Security   - managePasswordPolicy   - ipAllowlisting   - createAccessKeys   - manageAccessKeys   - manageSupportAccountAccess   - manageAuditDataFeed   - manageSaml   - shareDashboardOutsideOrg   - manageOrgSettings   - changeDataAccessLevel  ### Dashboards   - shareDashboardWorld   - shareDashboardAllowlist  ### UserManagement   - manageUsersAndRoles  ### Observability   - searchAuditIndex   - auditEventIndex  ### Cloud SIEM Enterprise   - viewCse  ### Alerting   - viewMonitorsV2   - manageMonitorsV2   - viewAlerts | 
**AutofillDependencies** | Pointer to **bool** | Set this to true if you want to automatically append all missing capability requirements. If set to false an error will be thrown if any capabilities are missing their dependencies. | [optional] [default to true]

## Methods

### NewUpdateRoleDefinitionV2

`func NewUpdateRoleDefinitionV2(name string, description string, logAnalyticsFilter string, auditDataFilter string, securityDataFilter string, selectionType string, selectedViews []ViewFilterDefinition, users []string, capabilities []string, ) *UpdateRoleDefinitionV2`

NewUpdateRoleDefinitionV2 instantiates a new UpdateRoleDefinitionV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRoleDefinitionV2WithDefaults

`func NewUpdateRoleDefinitionV2WithDefaults() *UpdateRoleDefinitionV2`

NewUpdateRoleDefinitionV2WithDefaults instantiates a new UpdateRoleDefinitionV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateRoleDefinitionV2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateRoleDefinitionV2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateRoleDefinitionV2) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *UpdateRoleDefinitionV2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateRoleDefinitionV2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateRoleDefinitionV2) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetLogAnalyticsFilter

`func (o *UpdateRoleDefinitionV2) GetLogAnalyticsFilter() string`

GetLogAnalyticsFilter returns the LogAnalyticsFilter field if non-nil, zero value otherwise.

### GetLogAnalyticsFilterOk

`func (o *UpdateRoleDefinitionV2) GetLogAnalyticsFilterOk() (*string, bool)`

GetLogAnalyticsFilterOk returns a tuple with the LogAnalyticsFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogAnalyticsFilter

`func (o *UpdateRoleDefinitionV2) SetLogAnalyticsFilter(v string)`

SetLogAnalyticsFilter sets LogAnalyticsFilter field to given value.


### GetAuditDataFilter

`func (o *UpdateRoleDefinitionV2) GetAuditDataFilter() string`

GetAuditDataFilter returns the AuditDataFilter field if non-nil, zero value otherwise.

### GetAuditDataFilterOk

`func (o *UpdateRoleDefinitionV2) GetAuditDataFilterOk() (*string, bool)`

GetAuditDataFilterOk returns a tuple with the AuditDataFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditDataFilter

`func (o *UpdateRoleDefinitionV2) SetAuditDataFilter(v string)`

SetAuditDataFilter sets AuditDataFilter field to given value.


### GetSecurityDataFilter

`func (o *UpdateRoleDefinitionV2) GetSecurityDataFilter() string`

GetSecurityDataFilter returns the SecurityDataFilter field if non-nil, zero value otherwise.

### GetSecurityDataFilterOk

`func (o *UpdateRoleDefinitionV2) GetSecurityDataFilterOk() (*string, bool)`

GetSecurityDataFilterOk returns a tuple with the SecurityDataFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityDataFilter

`func (o *UpdateRoleDefinitionV2) SetSecurityDataFilter(v string)`

SetSecurityDataFilter sets SecurityDataFilter field to given value.


### GetSelectionType

`func (o *UpdateRoleDefinitionV2) GetSelectionType() string`

GetSelectionType returns the SelectionType field if non-nil, zero value otherwise.

### GetSelectionTypeOk

`func (o *UpdateRoleDefinitionV2) GetSelectionTypeOk() (*string, bool)`

GetSelectionTypeOk returns a tuple with the SelectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectionType

`func (o *UpdateRoleDefinitionV2) SetSelectionType(v string)`

SetSelectionType sets SelectionType field to given value.


### GetSelectedViews

`func (o *UpdateRoleDefinitionV2) GetSelectedViews() []ViewFilterDefinition`

GetSelectedViews returns the SelectedViews field if non-nil, zero value otherwise.

### GetSelectedViewsOk

`func (o *UpdateRoleDefinitionV2) GetSelectedViewsOk() (*[]ViewFilterDefinition, bool)`

GetSelectedViewsOk returns a tuple with the SelectedViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedViews

`func (o *UpdateRoleDefinitionV2) SetSelectedViews(v []ViewFilterDefinition)`

SetSelectedViews sets SelectedViews field to given value.


### GetUsers

`func (o *UpdateRoleDefinitionV2) GetUsers() []string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *UpdateRoleDefinitionV2) GetUsersOk() (*[]string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *UpdateRoleDefinitionV2) SetUsers(v []string)`

SetUsers sets Users field to given value.


### GetCapabilities

`func (o *UpdateRoleDefinitionV2) GetCapabilities() []string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *UpdateRoleDefinitionV2) GetCapabilitiesOk() (*[]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *UpdateRoleDefinitionV2) SetCapabilities(v []string)`

SetCapabilities sets Capabilities field to given value.


### GetAutofillDependencies

`func (o *UpdateRoleDefinitionV2) GetAutofillDependencies() bool`

GetAutofillDependencies returns the AutofillDependencies field if non-nil, zero value otherwise.

### GetAutofillDependenciesOk

`func (o *UpdateRoleDefinitionV2) GetAutofillDependenciesOk() (*bool, bool)`

GetAutofillDependenciesOk returns a tuple with the AutofillDependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutofillDependencies

`func (o *UpdateRoleDefinitionV2) SetAutofillDependencies(v bool)`

SetAutofillDependencies sets AutofillDependencies field to given value.

### HasAutofillDependencies

`func (o *UpdateRoleDefinitionV2) HasAutofillDependencies() bool`

HasAutofillDependencies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


