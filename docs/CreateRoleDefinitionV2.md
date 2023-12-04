# CreateRoleDefinitionV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the role. | 
**Description** | Pointer to **string** | Description of the role. | [optional] 
**LogAnalyticsFilter** | Pointer to **string** | A search filter which would be applied on partitions which belong to Log Analytics product area. Applicable with only &#x60;All&#x60; selectionType | [optional] 
**AuditDataFilter** | Pointer to **string** | A search filter which would be applied on partitions which belong to Audit Data product area. Help Doc : (https://help.sumologic.com/docs/manage/security/audit-index/). Applicable with only &#x60;All&#x60; selectionType | [optional] 
**SecurityDataFilter** | Pointer to **string** | A search filter which would be applied on partitions which belong to Security Data product area. Applicable with only &#x60;All&#x60; selectionType. | [optional] 
**SelectionType** | Pointer to **string** | Describes the Permission Construct for the list of views in \&quot;selectedViews\&quot; parameter.  ### Valid Values are :    - &#x60;All&#x60; selectionType would allow access to all views in the org.   - &#x60;Allow&#x60; selectionType would allow access to specific views mentioned in \&quot;selectedViews\&quot; parameter.   - &#x60;Deny&#x60; selectionType would deny access to specific views mentioned in \&quot;selectedViews\&quot; parameter. | [optional] 
**SelectedViews** | Pointer to [**[]ViewFilterDefinition**](ViewFilterDefinition.md) | List of views which with specific view level filters in accordance to the selectionType chosen. | [optional] 
**Users** | Pointer to **[]string** | List of user identifiers to assign the role to. | [optional] 
**Capabilities** | Pointer to **[]string** | List of [capabilities](https://help.sumologic.com/docs/manage/users-roles/roles/role-capabilities/) associated with this role. Valid values are ### Data Management   - viewCollectors   - manageCollectors   - manageBudgets   - manageDataVolumeFeed   - viewFieldExtraction   - manageFieldExtractionRules   - manageS3DataForwarding   - manageContent   - manageApps   - dataVolumeIndex   - manageConnections   - viewScheduledViews   - manageScheduledViews   - viewPartitions   - managePartitions   - viewFields   - manageFields   - viewAccountOverview   - manageTokens   - downloadSearchResults   - manageIndexes   - manageDataStreams   - viewParsers   - viewDataStreams  ### Entity management   - manageEntityTypeConfig  ### Metrics   - metricsTransformation   - metricsExtraction   - metricsRules  ### Security   - managePasswordPolicy   - ipAllowlisting   - ipWhitelisting   - createAccessKeys   - manageAccessKeys   - manageSupportAccountAccess   - manageAuditDataFeed   - manageSaml   - shareDashboardOutsideOrg   - manageOrgSettings   - changeDataAccessLevel  ### Dashboards   - shareDashboardWorld   - shareDashboardAllowlist   - shareDashboardWhitelist  ### UserManagement   - manageUsersAndRoles  ### Observability   - searchAuditIndex   - auditEventIndex  ### Cloud SIEM Enterprise   - viewCse   - cseViewAutomations   - cseManageContextActions   - cseViewNetworkBlocks   - cseManageInsightTags   - cseViewRules   - cseViewThreatIntelligence   - cseCommentOnInsights   - cseViewEntityGroups   - cseManageEntityConfiguration   - cseManageNetworkBlocks   - cseManageMatchLists   - cseViewCustomInsights   - cseManageActions   - cseManageAutomations   - cseManageMappings   - cseManageThreatIntelligence   - cseViewActions   - cseCreateInsights   - cseManageTagSchemas   - cseInvokeInsights   - cseManageCustomEntityType   - cseViewTagSchemas   - cseDeleteInsights   - cseManageCustomInsights   - cseViewFileAnalysis   - cseManageFileAnalysis   - cseManageEntityCriticality   - cseViewEntityCriticality   - cseViewEntity   - cseManageCustomInsightStatuses   - cseViewContextActions   - cseViewMappings   - cseViewCustomEntityType   - cseManageEntityGroups   - cseViewCustomInsightStatuses   - cseViewEnrichments   - cseManageInsightSignals   - cseManageRules   - cseManageArtifacts   - cseViewMatchLists   - cseManageInsightPolicy   - cseManageEnrichments   - cseViewEntityConfiguration   - cseManageEntity   - cseExecuteAutomations   - cseManageSuppressedEntities   - cseManageInsightStatus     - cseManageInsightAssignee   - cseManageFavoriteFields   - cseViewSuppressedEntities  ### Alerting   - viewMonitorsV2   - manageMonitorsV2   - viewAlerts   - viewMutingSchedules   - manageMutingSchedules   - adminMonitorsV2  ### SLO   - viewSlos   - manageSlos  ### CloudSoar   - cloudSoarPlaybooksAccess   - cloudSoarNotificationConfigure   - cloudSoarReportAll   - cloudSoarIncidentTriageAccess   - cloudSoarIncidentTaskView   - cloudSoarIncidentChangeOwnership   - cloudSoarIncidentNotesEdit   - cloudSoarAPIEmailEdit   - cloudSoarIncidentTemplatesAccess   - cloudSoarIncidentPlaybooksManage   - cloudSoarGeneralConfigure   - cloudSoarEntitiesAccess   - cloudSoarEntitiesBulkPhysicalDelete   - cloudSoarIncidentAttachmentsAccess   - cloudSoarAppCentralAccess   - cloudSoarBridgeMonitoringAccess   - viewCloudSoar   - cloudSoarIncidentView   - cloudSoarObservabilityAccess   - cloudSoarAPIEmailRead   - cloudSoarAppCentralExport   - cloudSoarWidgetsAll   - cloudSoarIncidentTaskReassign   - cloudSoarIntegrationsAccess   - cloudSoarCustomizationIncidentLabels   - cloudSoarAutomationRulesConfigure   - cloudSoarIncidentTaskAccessAll   - cloudSoarAuditAndInformationConfigureAuditTrail   - cloudSoarIncidentTriageEdit   - cloudSoarIncidentEdit   - cloudSoarNotificationTriage   - cloudSoarIncidentTriageBulkPhysicalDelete   - cloudSoarIncidentNotesAccess   - cloudSoarAPIUse   - cloudSoarIncidentPlaybooksEdit   - cloudSoarDashboardAll   - cloudSoarEntitiesManage   - cloudSoarIncidentTemplatesConfigure   - cloudSoarIncidentTriageAccessAll   - cloudSoarPlaybooksConfigure   - cloudSoarIncidentAccessAll   - cloudSoarCustomizationLogo   - cloudSoarIncidentTaskAccess   - cloudSoarIncidentTriageView   - cloudSoarIntegrationsConfigure   - cloudSoarIncidentManageInvestigators   - cloudSoarIncidentAccess   - cloudSoarAuditAndInformationLicenseInformation   - cloudSoarIncidentBulkOperations   - cloudSoarCustomizationFields   - cloudSoarIncidentTaskEdit   - cloudSoarDashboardAccess   - cloudSoarIncidentAttachmentsEdit   - cloudSoarIncidentFoldersEdit   - cloudSoarUserManagementGroups   - cloudSoarIncidentPlaybooksAccess   - cloudSoarIncidentWarRoomUse   - cloudSoarReportAccess   - cloudSoarAuditAndInformationAuditTrail   - cloudSoarAutomationRulesAccess   - cloudSoarIncidentTriageChangeOwnership   - cloudSoarObservabilityManagement | [optional] 
**AutofillDependencies** | Pointer to **bool** | Set this to true if you want to automatically append all missing capability requirements. If set to false an error will be thrown if any capabilities are missing their dependencies. | [optional] [default to true]

## Methods

### NewCreateRoleDefinitionV2

`func NewCreateRoleDefinitionV2(name string, ) *CreateRoleDefinitionV2`

NewCreateRoleDefinitionV2 instantiates a new CreateRoleDefinitionV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRoleDefinitionV2WithDefaults

`func NewCreateRoleDefinitionV2WithDefaults() *CreateRoleDefinitionV2`

NewCreateRoleDefinitionV2WithDefaults instantiates a new CreateRoleDefinitionV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateRoleDefinitionV2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateRoleDefinitionV2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateRoleDefinitionV2) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateRoleDefinitionV2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateRoleDefinitionV2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateRoleDefinitionV2) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateRoleDefinitionV2) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLogAnalyticsFilter

`func (o *CreateRoleDefinitionV2) GetLogAnalyticsFilter() string`

GetLogAnalyticsFilter returns the LogAnalyticsFilter field if non-nil, zero value otherwise.

### GetLogAnalyticsFilterOk

`func (o *CreateRoleDefinitionV2) GetLogAnalyticsFilterOk() (*string, bool)`

GetLogAnalyticsFilterOk returns a tuple with the LogAnalyticsFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogAnalyticsFilter

`func (o *CreateRoleDefinitionV2) SetLogAnalyticsFilter(v string)`

SetLogAnalyticsFilter sets LogAnalyticsFilter field to given value.

### HasLogAnalyticsFilter

`func (o *CreateRoleDefinitionV2) HasLogAnalyticsFilter() bool`

HasLogAnalyticsFilter returns a boolean if a field has been set.

### GetAuditDataFilter

`func (o *CreateRoleDefinitionV2) GetAuditDataFilter() string`

GetAuditDataFilter returns the AuditDataFilter field if non-nil, zero value otherwise.

### GetAuditDataFilterOk

`func (o *CreateRoleDefinitionV2) GetAuditDataFilterOk() (*string, bool)`

GetAuditDataFilterOk returns a tuple with the AuditDataFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditDataFilter

`func (o *CreateRoleDefinitionV2) SetAuditDataFilter(v string)`

SetAuditDataFilter sets AuditDataFilter field to given value.

### HasAuditDataFilter

`func (o *CreateRoleDefinitionV2) HasAuditDataFilter() bool`

HasAuditDataFilter returns a boolean if a field has been set.

### GetSecurityDataFilter

`func (o *CreateRoleDefinitionV2) GetSecurityDataFilter() string`

GetSecurityDataFilter returns the SecurityDataFilter field if non-nil, zero value otherwise.

### GetSecurityDataFilterOk

`func (o *CreateRoleDefinitionV2) GetSecurityDataFilterOk() (*string, bool)`

GetSecurityDataFilterOk returns a tuple with the SecurityDataFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityDataFilter

`func (o *CreateRoleDefinitionV2) SetSecurityDataFilter(v string)`

SetSecurityDataFilter sets SecurityDataFilter field to given value.

### HasSecurityDataFilter

`func (o *CreateRoleDefinitionV2) HasSecurityDataFilter() bool`

HasSecurityDataFilter returns a boolean if a field has been set.

### GetSelectionType

`func (o *CreateRoleDefinitionV2) GetSelectionType() string`

GetSelectionType returns the SelectionType field if non-nil, zero value otherwise.

### GetSelectionTypeOk

`func (o *CreateRoleDefinitionV2) GetSelectionTypeOk() (*string, bool)`

GetSelectionTypeOk returns a tuple with the SelectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectionType

`func (o *CreateRoleDefinitionV2) SetSelectionType(v string)`

SetSelectionType sets SelectionType field to given value.

### HasSelectionType

`func (o *CreateRoleDefinitionV2) HasSelectionType() bool`

HasSelectionType returns a boolean if a field has been set.

### GetSelectedViews

`func (o *CreateRoleDefinitionV2) GetSelectedViews() []ViewFilterDefinition`

GetSelectedViews returns the SelectedViews field if non-nil, zero value otherwise.

### GetSelectedViewsOk

`func (o *CreateRoleDefinitionV2) GetSelectedViewsOk() (*[]ViewFilterDefinition, bool)`

GetSelectedViewsOk returns a tuple with the SelectedViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedViews

`func (o *CreateRoleDefinitionV2) SetSelectedViews(v []ViewFilterDefinition)`

SetSelectedViews sets SelectedViews field to given value.

### HasSelectedViews

`func (o *CreateRoleDefinitionV2) HasSelectedViews() bool`

HasSelectedViews returns a boolean if a field has been set.

### GetUsers

`func (o *CreateRoleDefinitionV2) GetUsers() []string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *CreateRoleDefinitionV2) GetUsersOk() (*[]string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *CreateRoleDefinitionV2) SetUsers(v []string)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *CreateRoleDefinitionV2) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetCapabilities

`func (o *CreateRoleDefinitionV2) GetCapabilities() []string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *CreateRoleDefinitionV2) GetCapabilitiesOk() (*[]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *CreateRoleDefinitionV2) SetCapabilities(v []string)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *CreateRoleDefinitionV2) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetAutofillDependencies

`func (o *CreateRoleDefinitionV2) GetAutofillDependencies() bool`

GetAutofillDependencies returns the AutofillDependencies field if non-nil, zero value otherwise.

### GetAutofillDependenciesOk

`func (o *CreateRoleDefinitionV2) GetAutofillDependenciesOk() (*bool, bool)`

GetAutofillDependenciesOk returns a tuple with the AutofillDependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutofillDependencies

`func (o *CreateRoleDefinitionV2) SetAutofillDependencies(v bool)`

SetAutofillDependencies sets AutofillDependencies field to given value.

### HasAutofillDependencies

`func (o *CreateRoleDefinitionV2) HasAutofillDependencies() bool`

HasAutofillDependencies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


