# CreateRoleDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the role. | 
**Description** | Pointer to **string** | Description of the role. | [optional] 
**FilterPredicate** | Pointer to **string** | A search filter to restrict access to specific logs. The filter is silently added to the beginning of each query a user runs. For example, using &#39;!_sourceCategory&#x3D;billing&#39; as a filter predicate will prevent users assigned to the role from viewing logs from the source category named &#39;billing&#39;. | [optional] 
**Users** | Pointer to **[]string** | List of user identifiers to assign the role to. | [optional] 
**Capabilities** | Pointer to **[]string** | List of [capabilities](https://help.sumologic.com/docs/manage/users-roles/roles/role-capabilities/) associated with this role. Valid values are ### Data Management   - viewCollectors   - manageCollectors   - manageBudgets   - manageDataVolumeFeed   - viewFieldExtraction   - manageFieldExtractionRules   - manageS3DataForwarding   - manageContent   - manageApps   - dataVolumeIndex   - manageConnections   - viewScheduledViews   - manageScheduledViews   - viewPartitions   - managePartitions   - viewFields   - manageFields   - viewAccountOverview   - manageTokens   - downloadSearchResults   - manageIndexes   - manageDataStreams   - viewParsers   - viewDataStreams  ### Entity management   - manageEntityTypeConfig  ### Metrics   - metricsTransformation   - metricsExtraction   - metricsRules  ### Security   - managePasswordPolicy   - ipAllowlisting   - ipWhitelisting   - createAccessKeys   - manageAccessKeys   - manageSupportAccountAccess   - manageAuditDataFeed   - manageSaml   - shareDashboardOutsideOrg   - manageOrgSettings   - changeDataAccessLevel  ### Dashboards   - shareDashboardWorld   - shareDashboardAllowlist   - shareDashboardWhitelist  ### UserManagement   - manageUsersAndRoles  ### Observability   - searchAuditIndex   - auditEventIndex  ### Cloud SIEM Enterprise   - viewCse   - cseViewAutomations   - cseManageContextActions   - cseViewNetworkBlocks   - cseManageInsightTags   - cseViewRules   - cseViewThreatIntelligence   - cseCommentOnInsights   - cseViewEntityGroups   - cseManageEntityConfiguration   - cseManageNetworkBlocks   - cseManageMatchLists   - cseViewCustomInsights   - cseManageActions   - cseManageAutomations   - cseManageMappings   - cseManageThreatIntelligence   - cseViewActions   - cseCreateInsights   - cseManageTagSchemas   - cseInvokeInsights   - cseManageCustomEntityType   - cseViewTagSchemas   - cseDeleteInsights   - cseManageCustomInsights   - cseViewFileAnalysis   - cseManageFileAnalysis   - cseManageEntityCriticality   - cseViewEntityCriticality   - cseViewEntity   - cseManageCustomInsightStatuses   - cseViewContextActions   - cseViewMappings   - cseViewCustomEntityType   - cseManageEntityGroups   - cseViewCustomInsightStatuses   - cseViewEnrichments   - cseManageInsightSignals   - cseManageRules   - cseManageArtifacts   - cseViewMatchLists   - cseManageInsightPolicy   - cseManageEnrichments   - cseViewEntityConfiguration   - cseManageEntity   - cseExecuteAutomations   - cseManageSuppressedEntities   - cseManageInsightStatus     - cseManageInsightAssignee   - cseManageFavoriteFields   - cseViewSuppressedEntities  ### Alerting   - viewMonitorsV2   - manageMonitorsV2   - viewAlerts   - viewMutingSchedules   - manageMutingSchedules   - adminMonitorsV2  ### SLO   - viewSlos   - manageSlos  ### CloudSoar   - cloudSoarPlaybooksAccess   - cloudSoarNotificationConfigure   - cloudSoarReportAll   - cloudSoarIncidentTriageAccess   - cloudSoarIncidentTaskView   - cloudSoarIncidentChangeOwnership   - cloudSoarIncidentNotesEdit   - cloudSoarAPIEmailEdit   - cloudSoarIncidentTemplatesAccess   - cloudSoarIncidentPlaybooksManage   - cloudSoarGeneralConfigure   - cloudSoarEntitiesAccess   - cloudSoarEntitiesBulkPhysicalDelete   - cloudSoarIncidentAttachmentsAccess   - cloudSoarAppCentralAccess   - cloudSoarBridgeMonitoringAccess   - viewCloudSoar   - cloudSoarIncidentView   - cloudSoarObservabilityAccess   - cloudSoarAPIEmailRead   - cloudSoarAppCentralExport   - cloudSoarWidgetsAll   - cloudSoarIncidentTaskReassign   - cloudSoarIntegrationsAccess   - cloudSoarCustomizationIncidentLabels   - cloudSoarAutomationRulesConfigure   - cloudSoarIncidentTaskAccessAll   - cloudSoarAuditAndInformationConfigureAuditTrail   - cloudSoarIncidentTriageEdit   - cloudSoarIncidentEdit   - cloudSoarNotificationTriage   - cloudSoarIncidentTriageBulkPhysicalDelete   - cloudSoarIncidentNotesAccess   - cloudSoarAPIUse   - cloudSoarIncidentPlaybooksEdit   - cloudSoarDashboardAll   - cloudSoarEntitiesManage   - cloudSoarIncidentTemplatesConfigure   - cloudSoarIncidentTriageAccessAll   - cloudSoarPlaybooksConfigure   - cloudSoarIncidentAccessAll   - cloudSoarCustomizationLogo   - cloudSoarIncidentTaskAccess   - cloudSoarIncidentTriageView   - cloudSoarIntegrationsConfigure   - cloudSoarIncidentManageInvestigators   - cloudSoarIncidentAccess   - cloudSoarAuditAndInformationLicenseInformation   - cloudSoarIncidentBulkOperations   - cloudSoarCustomizationFields   - cloudSoarIncidentTaskEdit   - cloudSoarDashboardAccess   - cloudSoarIncidentAttachmentsEdit   - cloudSoarIncidentFoldersEdit   - cloudSoarUserManagementGroups   - cloudSoarIncidentPlaybooksAccess   - cloudSoarIncidentWarRoomUse   - cloudSoarReportAccess   - cloudSoarAuditAndInformationAuditTrail   - cloudSoarAutomationRulesAccess   - cloudSoarIncidentTriageChangeOwnership   - cloudSoarObservabilityManagement | [optional] 
**AutofillDependencies** | Pointer to **bool** | Set this to true if you want to automatically append all missing capability requirements. If set to false an error will be thrown if any capabilities are missing their dependencies. | [optional] [default to true]

## Methods

### NewCreateRoleDefinition

`func NewCreateRoleDefinition(name string, ) *CreateRoleDefinition`

NewCreateRoleDefinition instantiates a new CreateRoleDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRoleDefinitionWithDefaults

`func NewCreateRoleDefinitionWithDefaults() *CreateRoleDefinition`

NewCreateRoleDefinitionWithDefaults instantiates a new CreateRoleDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateRoleDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateRoleDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateRoleDefinition) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateRoleDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateRoleDefinition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateRoleDefinition) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateRoleDefinition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFilterPredicate

`func (o *CreateRoleDefinition) GetFilterPredicate() string`

GetFilterPredicate returns the FilterPredicate field if non-nil, zero value otherwise.

### GetFilterPredicateOk

`func (o *CreateRoleDefinition) GetFilterPredicateOk() (*string, bool)`

GetFilterPredicateOk returns a tuple with the FilterPredicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterPredicate

`func (o *CreateRoleDefinition) SetFilterPredicate(v string)`

SetFilterPredicate sets FilterPredicate field to given value.

### HasFilterPredicate

`func (o *CreateRoleDefinition) HasFilterPredicate() bool`

HasFilterPredicate returns a boolean if a field has been set.

### GetUsers

`func (o *CreateRoleDefinition) GetUsers() []string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *CreateRoleDefinition) GetUsersOk() (*[]string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *CreateRoleDefinition) SetUsers(v []string)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *CreateRoleDefinition) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetCapabilities

`func (o *CreateRoleDefinition) GetCapabilities() []string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *CreateRoleDefinition) GetCapabilitiesOk() (*[]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *CreateRoleDefinition) SetCapabilities(v []string)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *CreateRoleDefinition) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetAutofillDependencies

`func (o *CreateRoleDefinition) GetAutofillDependencies() bool`

GetAutofillDependencies returns the AutofillDependencies field if non-nil, zero value otherwise.

### GetAutofillDependenciesOk

`func (o *CreateRoleDefinition) GetAutofillDependenciesOk() (*bool, bool)`

GetAutofillDependenciesOk returns a tuple with the AutofillDependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutofillDependencies

`func (o *CreateRoleDefinition) SetAutofillDependencies(v bool)`

SetAutofillDependencies sets AutofillDependencies field to given value.

### HasAutofillDependencies

`func (o *CreateRoleDefinition) HasAutofillDependencies() bool`

HasAutofillDependencies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


