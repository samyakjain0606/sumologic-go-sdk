/*
Sumo Logic API

# Getting Started Welcome to the Sumo Logic API reference. You can use these APIs to interact with the Sumo Logic platform. For information on the collector and search APIs, see our [API home page](https://help.sumologic.com/docs/api). ## API Endpoints Sumo Logic has several deployments in different geographic locations. You'll need to use the Sumo Logic API endpoint corresponding to your geographic location. See the table below for the different API endpoints by deployment. For details determining your account's deployment, see [API endpoints](https://help.sumologic.com/?cid=3011).    <table>     <tr>       <td> <strong>Deployment</strong> </td>       <td> <strong>Endpoint</strong> </td>     </tr>     <tr>       <td> AU </td>       <td> https://api.au.sumologic.com/api/ </td>     </tr>     <tr>       <td> CA </td>       <td> https://api.ca.sumologic.com/api/ </td>     </tr>     <tr>       <td> DE </td>       <td> https://api.de.sumologic.com/api/ </td>     </tr>     <tr>       <td> EU </td>       <td> https://api.eu.sumologic.com/api/ </td>     </tr>     <tr>       <td> FED </td>       <td> https://api.fed.sumologic.com/api/ </td>     </tr>     <tr>       <td> IN </td>       <td> https://api.in.sumologic.com/api/ </td>     </tr>     <tr>       <td> JP </td>       <td> https://api.jp.sumologic.com/api/ </td>     </tr>     <tr>       <td> US1 </td>       <td> https://api.sumologic.com/api/ </td>     </tr>     <tr>       <td> US2 </td>       <td> https://api.us2.sumologic.com/api/ </td>     </tr>   </table>  ## Authentication Sumo Logic supports the following options for API authentication: - Access ID and Access Key - Base64 encoded Access ID and Access Key  See [Access Keys](https://help.sumologic.com/docs/manage/security/access-keys) to generate an Access Key. Make sure to copy the key you create, because it is displayed only once. When you have an Access ID and Access Key you can execute requests such as the following:   ```bash   curl -u \"<accessId>:<accessKey>\" -X GET https://api.<deployment>.sumologic.com/api/v1/users   ```  Where `deployment` is either `au`, `ca`, `de`, `eu`, `fed`, `in`, `jp`, `us1`, or `us2`. See [API endpoints](#section/API-Endpoints) for details.  If you prefer to use basic access authentication, you can do a Base64 encoding of your `<accessId>:<accessKey>` to authenticate your HTTPS request. The following is an example request, replace the placeholder `<encoded>` with your encoded Access ID and Access Key string:   ```bash   curl -H \"Authorization: Basic <encoded>\" -X GET https://api.<deployment>.sumologic.com/api/v1/users   ```   Refer to [API Authentication](https://help.sumologic.com/?cid=3012) for a Base64 example.  ## Status Codes Generic status codes that apply to all our APIs. See the [HTTP status code registry](https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml) for reference.   <table>     <tr>       <td> <strong>HTTP Status Code</strong> </td>       <td> <strong>Error Code</strong> </td>       <td> <strong>Description</strong> </td>     </tr>     <tr>       <td> 301 </td>       <td> moved </td>       <td> The requested resource SHOULD be accessed through returned URI in Location Header. See [troubleshooting](https://help.sumologic.com/docs/api/troubleshooting/#api---301-error---moved) for details.</td>     </tr>     <tr>       <td> 401 </td>       <td> unauthorized </td>       <td> Credential could not be verified.</td>     </tr>     <tr>       <td> 403 </td>       <td> forbidden </td>       <td> This operation is not allowed for your account type or the user doesn't have the role capability to perform this action. See [troubleshooting](https://help.sumologic.com/docs/api/troubleshooting/#api---401-error---credential-could-not-be-verified) for details.</td>     </tr>     <tr>       <td> 404 </td>       <td> notfound </td>       <td> Requested resource could not be found. </td>     </tr>     <tr>       <td> 405 </td>       <td> method.unsupported </td>       <td> Unsupported method for URL. </td>     </tr>     <tr>       <td> 415 </td>       <td> contenttype.invalid </td>       <td> Invalid content type. </td>     </tr>     <tr>       <td> 429 </td>       <td> rate.limit.exceeded </td>       <td> The API request rate is higher than 4 request per second or inflight API requests are higher than 10 request per second. </td>     </tr>     <tr>       <td> 500 </td>       <td> internal.error </td>       <td> Internal server error. </td>     </tr>     <tr>       <td> 503 </td>       <td> service.unavailable </td>       <td> Service is currently unavailable. </td>     </tr>   </table>  ## Filtering Some API endpoints support filtering results on a specified set of fields. Each endpoint that supports filtering will list the fields that can be filtered. Multiple fields can be combined by using an ampersand `&` character.  For example, to get 20 users whose `firstName` is `John` and `lastName` is `Doe`:   ```bash   api.sumologic.com/v1/users?limit=20&firstName=John&lastName=Doe   ```  ## Sorting Some API endpoints support sorting fields by using the `sortBy` query parameter. The default sort order is ascending. Prefix the field with a minus sign `-` to sort in descending order.  For example, to get 20 users sorted by their `email` in descending order:   ```bash   api.sumologic.com/v1/users?limit=20&sort=-email   ```  ## Asynchronous Request Asynchronous requests do not wait for results, instead they immediately respond back with a job identifier while the job runs in the background. You can use the job identifier to track the status of the asynchronous job request. Here is a typical flow for an asynchronous request. 1. Start an asynchronous job. On success, a job identifier is returned. The job identifier uniquely identifies   your asynchronous job.  2. Once started, use the job identifier from step 1 to track the status of your asynchronous job. An asynchronous   request will typically provide an endpoint to poll for the status of asynchronous job. A successful response   from the status endpoint will have the following structure:   ```json   {       \"status\": \"Status of asynchronous request\",       \"statusMessage\": \"Optional message with additional information in case request succeeds\",       \"error\": \"Error object in case request fails\"   }   ```   The `status` field can have one of the following values:     1. `Success`: The job succeeded. The `statusMessage` field might have additional information.     2. `InProgress`: The job is still running.     3. `Failed`: The job failed. The `error` field in the response will have more information about the failure.  3. Some asynchronous APIs may provide a third endpoint (like [export result](#operation/getAsyncExportResult))   to fetch the result of an asynchronous job.   ### Example Let's say we want to export a folder with the identifier `0000000006A2E86F`. We will use the [async export](#operation/beginAsyncExport) API to export all the content under the folder with `id=0000000006A2E86F`. 1. Start an export job for the folder   ```bash   curl -X POST -u \"<accessId>:<accessKey>\" https://api.<deployment>.sumologic.com/api/v2/content/0000000006A2E86F/export   ```   See [authentication section](#section/Authentication) for more details about `accessId`, `accessKey`, and   `deployment`.   On success, you will get back a job identifier. In the response below, `C03E086C137F38B4` is the job identifier.   ```bash   {       \"id\": \"C03E086C137F38B4\"   }   ```  2. Now poll for the status of the asynchronous job with the [status](#operation/getAsyncExportStatus) endpoint.   ```bash   curl -X GET -u \"<accessId>:<accessKey>\" https://api.<deployment>.sumologic.com/api/v2/content/0000000006A2E86F/export/C03E086C137F38B4/status   ```   You may get a response like   ```json   {       \"status\": \"InProgress\",       \"statusMessage\": null,       \"error\": null   }   ```   It implies the job is still in progress. Keep polling till the status is either `Success` or `Failed`.  3. When the asynchronous job completes (`status != \"InProgress\"`), you can fetch the results with the   [export result](#operation/getAsyncExportResult) endpoint.   ```bash   curl -X GET -u \"<accessId>:<accessKey>\" https://api.<deployment>.sumologic.com/api/v2/content/0000000006A2E86F/export/C03E086C137F38B4/result   ```    The asynchronous job may fail (`status == \"Failed\"`). You can look at the `error` field for more details.   ```json   {       \"status\": \"Failed\",       \"errors\": {           \"code\": \"content1:too_many_items\",           \"message\": \"Too many objects: object count(1100) was greater than limit 1000\"       }   }   ```   ## Rate Limiting * A rate limit of four API requests per second (240 requests per minute) applies to all API calls from a user. * A rate limit of 10 concurrent requests to any API endpoint applies to an access key.  If a rate is exceeded, a rate limit exceeded 429 status code is returned.  ## Generating Clients You can use [OpenAPI Generator](https://openapi-generator.tech) to generate clients from the YAML file to access the API.  ### Using [NPM](https://www.npmjs.com/get-npm) 1. Install [NPM package wrapper](https://github.com/openapitools/openapi-generator-cli) globally, exposing the CLI   on the command line:   ```bash   npm install @openapitools/openapi-generator-cli -g   ```   You can see detailed instructions [here](https://openapi-generator.tech/docs/installation#npm).  2. Download the [YAML file](/docs/sumologic-api.yaml) and save it locally. Let's say the file is saved as `sumologic-api.yaml`. 3. Use the following command to generate `python` client inside the `sumo/client/python` directory:   ```bash   openapi-generator generate -i sumologic-api.yaml -g python -o sumo/client/python   ```   ### Using [Homebrew](https://brew.sh/) 1. Install OpenAPI Generator   ```bash   brew install openapi-generator   ```  2. Download the [YAML file](/docs/sumologic-api.yaml) and save it locally. Let's say the file is saved as `sumologic-api.yaml`. 3. Use the following command to generate `python` client side code inside the `sumo/client/python` directory:   ```bash   openapi-generator generate -i sumologic-api.yaml -g python -o sumo/client/python   ``` 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the GetRoleDefinitionV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRoleDefinitionV2{}

// GetRoleDefinitionV2 struct for GetRoleDefinitionV2
type GetRoleDefinitionV2 struct {
	// Name of the role.
	Name string `json:"name"`
	// Description of the role.
	Description *string `json:"description,omitempty"`
	// A search filter which would be applied on partitions which belong to Log Analytics product area. Applicable with only `All` selectionType
	LogAnalyticsFilter *string `json:"logAnalyticsFilter,omitempty"`
	// A search filter which would be applied on partitions which belong to Audit Data product area. Help Doc : (https://help.sumologic.com/docs/manage/security/audit-index/). Applicable with only `All` selectionType
	AuditDataFilter *string `json:"auditDataFilter,omitempty"`
	// A search filter which would be applied on partitions which belong to Security Data product area. Applicable with only `All` selectionType
	SecurityDataFilter *string `json:"securityDataFilter,omitempty"`
	// Describes the Permission Construct for the list of views in \"selectedViews\" parameter.  ### Valid Values are :    - `All` selectionType would allow access to all views in the org.   - `Allow` selectionType would allow access to specific views mentioned in \"selectedViews\" parameter.   - `Deny` selectionType would deny access to specific views mentioned in \"selectedViews\" parameter.
	SelectionType *string `json:"selectionType,omitempty"`
	// List of views which with specific view level filters in accordance to the selectionType chosen.
	SelectedViews []GetViewFilterDefinition `json:"selectedViews,omitempty"`
	// List of user identifiers to assign the role to.
	Users []string `json:"users,omitempty"`
	// List of [capabilities](https://help.sumologic.com/Manage/Users-and-Roles/Manage-Roles/Role-Capabilities) associated with this role. Valid values are ### Data Management   - viewCollectors   - manageCollectors   - manageBudgets   - manageDataVolumeFeed   - viewFieldExtraction   - manageFieldExtractionRules   - manageS3DataForwarding   - manageContent   - manageApps   - dataVolumeIndex   - manageConnections   - viewScheduledViews   - manageScheduledViews   - viewPartitions   - managePartitions   - viewFields   - manageFields   - viewAccountOverview   - manageTokens   - downloadSearchResults  ### Entity management   - manageEntityTypeConfig  ### Metrics   - metricsTransformation   - metricsExtraction   - metricsRules  ### Security   - managePasswordPolicy   - ipAllowlisting   - createAccessKeys   - manageAccessKeys   - manageSupportAccountAccess   - manageAuditDataFeed   - manageSaml   - shareDashboardOutsideOrg   - manageOrgSettings   - changeDataAccessLevel  ### Dashboards   - shareDashboardWorld   - shareDashboardAllowlist  ### UserManagement   - manageUsersAndRoles  ### Observability   - searchAuditIndex   - auditEventIndex  ### Cloud SIEM Enterprise   - viewCse  ### Alerting   - viewMonitorsV2   - manageMonitorsV2   - viewAlerts
	Capabilities []string `json:"capabilities,omitempty"`
	// Set this to true if you want to automatically append all missing capability requirements. If set to false an error will be thrown if any capabilities are missing their dependencies.
	AutofillDependencies *bool `json:"autofillDependencies,omitempty"`
	// Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format.
	CreatedAt time.Time `json:"createdAt"`
	// Identifier of the user who created the resource.
	CreatedBy string `json:"createdBy"`
	// Last modification timestamp in UTC.
	ModifiedAt time.Time `json:"modifiedAt"`
	// Identifier of the user who last modified the resource.
	ModifiedBy string `json:"modifiedBy"`
	// Unique identifier for the role.
	Id string `json:"id"`
	// Role is system or user defined.
	SystemDefined *bool `json:"systemDefined,omitempty"`
}

type _GetRoleDefinitionV2 GetRoleDefinitionV2

// NewGetRoleDefinitionV2 instantiates a new GetRoleDefinitionV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRoleDefinitionV2(name string, createdAt time.Time, createdBy string, modifiedAt time.Time, modifiedBy string, id string) *GetRoleDefinitionV2 {
	this := GetRoleDefinitionV2{}
	this.Name = name
	var autofillDependencies bool = true
	this.AutofillDependencies = &autofillDependencies
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.ModifiedAt = modifiedAt
	this.ModifiedBy = modifiedBy
	this.Id = id
	return &this
}

// NewGetRoleDefinitionV2WithDefaults instantiates a new GetRoleDefinitionV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRoleDefinitionV2WithDefaults() *GetRoleDefinitionV2 {
	this := GetRoleDefinitionV2{}
	var autofillDependencies bool = true
	this.AutofillDependencies = &autofillDependencies
	return &this
}

// GetName returns the Name field value
func (o *GetRoleDefinitionV2) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetRoleDefinitionV2) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetRoleDefinitionV2) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GetRoleDefinitionV2) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRoleDefinitionV2) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GetRoleDefinitionV2) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GetRoleDefinitionV2) SetDescription(v string) {
	o.Description = &v
}

// GetLogAnalyticsFilter returns the LogAnalyticsFilter field value if set, zero value otherwise.
func (o *GetRoleDefinitionV2) GetLogAnalyticsFilter() string {
	if o == nil || IsNil(o.LogAnalyticsFilter) {
		var ret string
		return ret
	}
	return *o.LogAnalyticsFilter
}

// GetLogAnalyticsFilterOk returns a tuple with the LogAnalyticsFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRoleDefinitionV2) GetLogAnalyticsFilterOk() (*string, bool) {
	if o == nil || IsNil(o.LogAnalyticsFilter) {
		return nil, false
	}
	return o.LogAnalyticsFilter, true
}

// HasLogAnalyticsFilter returns a boolean if a field has been set.
func (o *GetRoleDefinitionV2) HasLogAnalyticsFilter() bool {
	if o != nil && !IsNil(o.LogAnalyticsFilter) {
		return true
	}

	return false
}

// SetLogAnalyticsFilter gets a reference to the given string and assigns it to the LogAnalyticsFilter field.
func (o *GetRoleDefinitionV2) SetLogAnalyticsFilter(v string) {
	o.LogAnalyticsFilter = &v
}

// GetAuditDataFilter returns the AuditDataFilter field value if set, zero value otherwise.
func (o *GetRoleDefinitionV2) GetAuditDataFilter() string {
	if o == nil || IsNil(o.AuditDataFilter) {
		var ret string
		return ret
	}
	return *o.AuditDataFilter
}

// GetAuditDataFilterOk returns a tuple with the AuditDataFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRoleDefinitionV2) GetAuditDataFilterOk() (*string, bool) {
	if o == nil || IsNil(o.AuditDataFilter) {
		return nil, false
	}
	return o.AuditDataFilter, true
}

// HasAuditDataFilter returns a boolean if a field has been set.
func (o *GetRoleDefinitionV2) HasAuditDataFilter() bool {
	if o != nil && !IsNil(o.AuditDataFilter) {
		return true
	}

	return false
}

// SetAuditDataFilter gets a reference to the given string and assigns it to the AuditDataFilter field.
func (o *GetRoleDefinitionV2) SetAuditDataFilter(v string) {
	o.AuditDataFilter = &v
}

// GetSecurityDataFilter returns the SecurityDataFilter field value if set, zero value otherwise.
func (o *GetRoleDefinitionV2) GetSecurityDataFilter() string {
	if o == nil || IsNil(o.SecurityDataFilter) {
		var ret string
		return ret
	}
	return *o.SecurityDataFilter
}

// GetSecurityDataFilterOk returns a tuple with the SecurityDataFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRoleDefinitionV2) GetSecurityDataFilterOk() (*string, bool) {
	if o == nil || IsNil(o.SecurityDataFilter) {
		return nil, false
	}
	return o.SecurityDataFilter, true
}

// HasSecurityDataFilter returns a boolean if a field has been set.
func (o *GetRoleDefinitionV2) HasSecurityDataFilter() bool {
	if o != nil && !IsNil(o.SecurityDataFilter) {
		return true
	}

	return false
}

// SetSecurityDataFilter gets a reference to the given string and assigns it to the SecurityDataFilter field.
func (o *GetRoleDefinitionV2) SetSecurityDataFilter(v string) {
	o.SecurityDataFilter = &v
}

// GetSelectionType returns the SelectionType field value if set, zero value otherwise.
func (o *GetRoleDefinitionV2) GetSelectionType() string {
	if o == nil || IsNil(o.SelectionType) {
		var ret string
		return ret
	}
	return *o.SelectionType
}

// GetSelectionTypeOk returns a tuple with the SelectionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRoleDefinitionV2) GetSelectionTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SelectionType) {
		return nil, false
	}
	return o.SelectionType, true
}

// HasSelectionType returns a boolean if a field has been set.
func (o *GetRoleDefinitionV2) HasSelectionType() bool {
	if o != nil && !IsNil(o.SelectionType) {
		return true
	}

	return false
}

// SetSelectionType gets a reference to the given string and assigns it to the SelectionType field.
func (o *GetRoleDefinitionV2) SetSelectionType(v string) {
	o.SelectionType = &v
}

// GetSelectedViews returns the SelectedViews field value if set, zero value otherwise.
func (o *GetRoleDefinitionV2) GetSelectedViews() []GetViewFilterDefinition {
	if o == nil || IsNil(o.SelectedViews) {
		var ret []GetViewFilterDefinition
		return ret
	}
	return o.SelectedViews
}

// GetSelectedViewsOk returns a tuple with the SelectedViews field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRoleDefinitionV2) GetSelectedViewsOk() ([]GetViewFilterDefinition, bool) {
	if o == nil || IsNil(o.SelectedViews) {
		return nil, false
	}
	return o.SelectedViews, true
}

// HasSelectedViews returns a boolean if a field has been set.
func (o *GetRoleDefinitionV2) HasSelectedViews() bool {
	if o != nil && !IsNil(o.SelectedViews) {
		return true
	}

	return false
}

// SetSelectedViews gets a reference to the given []GetViewFilterDefinition and assigns it to the SelectedViews field.
func (o *GetRoleDefinitionV2) SetSelectedViews(v []GetViewFilterDefinition) {
	o.SelectedViews = v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *GetRoleDefinitionV2) GetUsers() []string {
	if o == nil || IsNil(o.Users) {
		var ret []string
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRoleDefinitionV2) GetUsersOk() ([]string, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *GetRoleDefinitionV2) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []string and assigns it to the Users field.
func (o *GetRoleDefinitionV2) SetUsers(v []string) {
	o.Users = v
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise.
func (o *GetRoleDefinitionV2) GetCapabilities() []string {
	if o == nil || IsNil(o.Capabilities) {
		var ret []string
		return ret
	}
	return o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRoleDefinitionV2) GetCapabilitiesOk() ([]string, bool) {
	if o == nil || IsNil(o.Capabilities) {
		return nil, false
	}
	return o.Capabilities, true
}

// HasCapabilities returns a boolean if a field has been set.
func (o *GetRoleDefinitionV2) HasCapabilities() bool {
	if o != nil && !IsNil(o.Capabilities) {
		return true
	}

	return false
}

// SetCapabilities gets a reference to the given []string and assigns it to the Capabilities field.
func (o *GetRoleDefinitionV2) SetCapabilities(v []string) {
	o.Capabilities = v
}

// GetAutofillDependencies returns the AutofillDependencies field value if set, zero value otherwise.
func (o *GetRoleDefinitionV2) GetAutofillDependencies() bool {
	if o == nil || IsNil(o.AutofillDependencies) {
		var ret bool
		return ret
	}
	return *o.AutofillDependencies
}

// GetAutofillDependenciesOk returns a tuple with the AutofillDependencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRoleDefinitionV2) GetAutofillDependenciesOk() (*bool, bool) {
	if o == nil || IsNil(o.AutofillDependencies) {
		return nil, false
	}
	return o.AutofillDependencies, true
}

// HasAutofillDependencies returns a boolean if a field has been set.
func (o *GetRoleDefinitionV2) HasAutofillDependencies() bool {
	if o != nil && !IsNil(o.AutofillDependencies) {
		return true
	}

	return false
}

// SetAutofillDependencies gets a reference to the given bool and assigns it to the AutofillDependencies field.
func (o *GetRoleDefinitionV2) SetAutofillDependencies(v bool) {
	o.AutofillDependencies = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *GetRoleDefinitionV2) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *GetRoleDefinitionV2) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *GetRoleDefinitionV2) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *GetRoleDefinitionV2) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *GetRoleDefinitionV2) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *GetRoleDefinitionV2) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetModifiedAt returns the ModifiedAt field value
func (o *GetRoleDefinitionV2) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *GetRoleDefinitionV2) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value
func (o *GetRoleDefinitionV2) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// GetModifiedBy returns the ModifiedBy field value
func (o *GetRoleDefinitionV2) GetModifiedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value
// and a boolean to check if the value has been set.
func (o *GetRoleDefinitionV2) GetModifiedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedBy, true
}

// SetModifiedBy sets field value
func (o *GetRoleDefinitionV2) SetModifiedBy(v string) {
	o.ModifiedBy = v
}

// GetId returns the Id field value
func (o *GetRoleDefinitionV2) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GetRoleDefinitionV2) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GetRoleDefinitionV2) SetId(v string) {
	o.Id = v
}

// GetSystemDefined returns the SystemDefined field value if set, zero value otherwise.
func (o *GetRoleDefinitionV2) GetSystemDefined() bool {
	if o == nil || IsNil(o.SystemDefined) {
		var ret bool
		return ret
	}
	return *o.SystemDefined
}

// GetSystemDefinedOk returns a tuple with the SystemDefined field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRoleDefinitionV2) GetSystemDefinedOk() (*bool, bool) {
	if o == nil || IsNil(o.SystemDefined) {
		return nil, false
	}
	return o.SystemDefined, true
}

// HasSystemDefined returns a boolean if a field has been set.
func (o *GetRoleDefinitionV2) HasSystemDefined() bool {
	if o != nil && !IsNil(o.SystemDefined) {
		return true
	}

	return false
}

// SetSystemDefined gets a reference to the given bool and assigns it to the SystemDefined field.
func (o *GetRoleDefinitionV2) SetSystemDefined(v bool) {
	o.SystemDefined = &v
}

func (o GetRoleDefinitionV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRoleDefinitionV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.LogAnalyticsFilter) {
		toSerialize["logAnalyticsFilter"] = o.LogAnalyticsFilter
	}
	if !IsNil(o.AuditDataFilter) {
		toSerialize["auditDataFilter"] = o.AuditDataFilter
	}
	if !IsNil(o.SecurityDataFilter) {
		toSerialize["securityDataFilter"] = o.SecurityDataFilter
	}
	if !IsNil(o.SelectionType) {
		toSerialize["selectionType"] = o.SelectionType
	}
	if !IsNil(o.SelectedViews) {
		toSerialize["selectedViews"] = o.SelectedViews
	}
	if !IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}
	if !IsNil(o.Capabilities) {
		toSerialize["capabilities"] = o.Capabilities
	}
	if !IsNil(o.AutofillDependencies) {
		toSerialize["autofillDependencies"] = o.AutofillDependencies
	}
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["createdBy"] = o.CreatedBy
	toSerialize["modifiedAt"] = o.ModifiedAt
	toSerialize["modifiedBy"] = o.ModifiedBy
	toSerialize["id"] = o.Id
	if !IsNil(o.SystemDefined) {
		toSerialize["systemDefined"] = o.SystemDefined
	}
	return toSerialize, nil
}

func (o *GetRoleDefinitionV2) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"createdAt",
		"createdBy",
		"modifiedAt",
		"modifiedBy",
		"id",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varGetRoleDefinitionV2 := _GetRoleDefinitionV2{}

	err = json.Unmarshal(bytes, &varGetRoleDefinitionV2)

	if err != nil {
		return err
	}

	*o = GetRoleDefinitionV2(varGetRoleDefinitionV2)

	return err
}

type NullableGetRoleDefinitionV2 struct {
	value *GetRoleDefinitionV2
	isSet bool
}

func (v NullableGetRoleDefinitionV2) Get() *GetRoleDefinitionV2 {
	return v.value
}

func (v *NullableGetRoleDefinitionV2) Set(val *GetRoleDefinitionV2) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRoleDefinitionV2) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRoleDefinitionV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRoleDefinitionV2(val *GetRoleDefinitionV2) *NullableGetRoleDefinitionV2 {
	return &NullableGetRoleDefinitionV2{value: val, isSet: true}
}

func (v NullableGetRoleDefinitionV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRoleDefinitionV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


