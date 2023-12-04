# GetRoleDefinitionV2

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
**CreatedAt** | **time.Time** | Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format. | 
**CreatedBy** | **string** | Identifier of the user who created the resource. | 
**ModifiedAt** | **time.Time** | Last modification timestamp in UTC. | 
**ModifiedBy** | **string** | Identifier of the user who last modified the resource. | 
**Id** | **string** | Unique identifier for the role. | 
**SystemDefined** | Pointer to **bool** | Role is system or user defined. | [optional] 

## Methods

### NewGetRoleDefinitionV2

`func NewGetRoleDefinitionV2(name string, createdAt time.Time, createdBy string, modifiedAt time.Time, modifiedBy string, id string, ) *GetRoleDefinitionV2`

NewGetRoleDefinitionV2 instantiates a new GetRoleDefinitionV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRoleDefinitionV2WithDefaults

`func NewGetRoleDefinitionV2WithDefaults() *GetRoleDefinitionV2`

NewGetRoleDefinitionV2WithDefaults instantiates a new GetRoleDefinitionV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GetRoleDefinitionV2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetRoleDefinitionV2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetRoleDefinitionV2) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *GetRoleDefinitionV2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetRoleDefinitionV2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetRoleDefinitionV2) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetRoleDefinitionV2) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLogAnalyticsFilter

`func (o *GetRoleDefinitionV2) GetLogAnalyticsFilter() string`

GetLogAnalyticsFilter returns the LogAnalyticsFilter field if non-nil, zero value otherwise.

### GetLogAnalyticsFilterOk

`func (o *GetRoleDefinitionV2) GetLogAnalyticsFilterOk() (*string, bool)`

GetLogAnalyticsFilterOk returns a tuple with the LogAnalyticsFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogAnalyticsFilter

`func (o *GetRoleDefinitionV2) SetLogAnalyticsFilter(v string)`

SetLogAnalyticsFilter sets LogAnalyticsFilter field to given value.

### HasLogAnalyticsFilter

`func (o *GetRoleDefinitionV2) HasLogAnalyticsFilter() bool`

HasLogAnalyticsFilter returns a boolean if a field has been set.

### GetAuditDataFilter

`func (o *GetRoleDefinitionV2) GetAuditDataFilter() string`

GetAuditDataFilter returns the AuditDataFilter field if non-nil, zero value otherwise.

### GetAuditDataFilterOk

`func (o *GetRoleDefinitionV2) GetAuditDataFilterOk() (*string, bool)`

GetAuditDataFilterOk returns a tuple with the AuditDataFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditDataFilter

`func (o *GetRoleDefinitionV2) SetAuditDataFilter(v string)`

SetAuditDataFilter sets AuditDataFilter field to given value.

### HasAuditDataFilter

`func (o *GetRoleDefinitionV2) HasAuditDataFilter() bool`

HasAuditDataFilter returns a boolean if a field has been set.

### GetSecurityDataFilter

`func (o *GetRoleDefinitionV2) GetSecurityDataFilter() string`

GetSecurityDataFilter returns the SecurityDataFilter field if non-nil, zero value otherwise.

### GetSecurityDataFilterOk

`func (o *GetRoleDefinitionV2) GetSecurityDataFilterOk() (*string, bool)`

GetSecurityDataFilterOk returns a tuple with the SecurityDataFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityDataFilter

`func (o *GetRoleDefinitionV2) SetSecurityDataFilter(v string)`

SetSecurityDataFilter sets SecurityDataFilter field to given value.

### HasSecurityDataFilter

`func (o *GetRoleDefinitionV2) HasSecurityDataFilter() bool`

HasSecurityDataFilter returns a boolean if a field has been set.

### GetSelectionType

`func (o *GetRoleDefinitionV2) GetSelectionType() string`

GetSelectionType returns the SelectionType field if non-nil, zero value otherwise.

### GetSelectionTypeOk

`func (o *GetRoleDefinitionV2) GetSelectionTypeOk() (*string, bool)`

GetSelectionTypeOk returns a tuple with the SelectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectionType

`func (o *GetRoleDefinitionV2) SetSelectionType(v string)`

SetSelectionType sets SelectionType field to given value.

### HasSelectionType

`func (o *GetRoleDefinitionV2) HasSelectionType() bool`

HasSelectionType returns a boolean if a field has been set.

### GetSelectedViews

`func (o *GetRoleDefinitionV2) GetSelectedViews() []GetViewFilterDefinition`

GetSelectedViews returns the SelectedViews field if non-nil, zero value otherwise.

### GetSelectedViewsOk

`func (o *GetRoleDefinitionV2) GetSelectedViewsOk() (*[]GetViewFilterDefinition, bool)`

GetSelectedViewsOk returns a tuple with the SelectedViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedViews

`func (o *GetRoleDefinitionV2) SetSelectedViews(v []GetViewFilterDefinition)`

SetSelectedViews sets SelectedViews field to given value.

### HasSelectedViews

`func (o *GetRoleDefinitionV2) HasSelectedViews() bool`

HasSelectedViews returns a boolean if a field has been set.

### GetUsers

`func (o *GetRoleDefinitionV2) GetUsers() []string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *GetRoleDefinitionV2) GetUsersOk() (*[]string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *GetRoleDefinitionV2) SetUsers(v []string)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *GetRoleDefinitionV2) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetCapabilities

`func (o *GetRoleDefinitionV2) GetCapabilities() []string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *GetRoleDefinitionV2) GetCapabilitiesOk() (*[]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *GetRoleDefinitionV2) SetCapabilities(v []string)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *GetRoleDefinitionV2) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetAutofillDependencies

`func (o *GetRoleDefinitionV2) GetAutofillDependencies() bool`

GetAutofillDependencies returns the AutofillDependencies field if non-nil, zero value otherwise.

### GetAutofillDependenciesOk

`func (o *GetRoleDefinitionV2) GetAutofillDependenciesOk() (*bool, bool)`

GetAutofillDependenciesOk returns a tuple with the AutofillDependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutofillDependencies

`func (o *GetRoleDefinitionV2) SetAutofillDependencies(v bool)`

SetAutofillDependencies sets AutofillDependencies field to given value.

### HasAutofillDependencies

`func (o *GetRoleDefinitionV2) HasAutofillDependencies() bool`

HasAutofillDependencies returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetRoleDefinitionV2) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetRoleDefinitionV2) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetRoleDefinitionV2) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *GetRoleDefinitionV2) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GetRoleDefinitionV2) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GetRoleDefinitionV2) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetModifiedAt

`func (o *GetRoleDefinitionV2) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *GetRoleDefinitionV2) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *GetRoleDefinitionV2) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *GetRoleDefinitionV2) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *GetRoleDefinitionV2) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *GetRoleDefinitionV2) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetId

`func (o *GetRoleDefinitionV2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetRoleDefinitionV2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetRoleDefinitionV2) SetId(v string)`

SetId sets Id field to given value.


### GetSystemDefined

`func (o *GetRoleDefinitionV2) GetSystemDefined() bool`

GetSystemDefined returns the SystemDefined field if non-nil, zero value otherwise.

### GetSystemDefinedOk

`func (o *GetRoleDefinitionV2) GetSystemDefinedOk() (*bool, bool)`

GetSystemDefinedOk returns a tuple with the SystemDefined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemDefined

`func (o *GetRoleDefinitionV2) SetSystemDefined(v bool)`

SetSystemDefined sets SystemDefined field to given value.

### HasSystemDefined

`func (o *GetRoleDefinitionV2) HasSystemDefined() bool`

HasSystemDefined returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


