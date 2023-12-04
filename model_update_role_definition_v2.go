/*
Sumo Logic API

# Getting Started Welcome to the Sumo Logic API reference. You can use these APIs to interact with the Sumo Logic platform. For information on the collector and search APIs, see our [API home page](https://help.sumologic.com/docs/api). ## API Endpoints Sumo Logic has several deployments in different geographic locations. You'll need to use the Sumo Logic API endpoint corresponding to your geographic location. See the table below for the different API endpoints by deployment. For details determining your account's deployment, see [API endpoints](https://help.sumologic.com/?cid=3011).    <table>     <tr>       <td> <strong>Deployment</strong> </td>       <td> <strong>Endpoint</strong> </td>     </tr>     <tr>       <td> AU </td>       <td> https://api.au.sumologic.com/api/ </td>     </tr>     <tr>       <td> CA </td>       <td> https://api.ca.sumologic.com/api/ </td>     </tr>     <tr>       <td> DE </td>       <td> https://api.de.sumologic.com/api/ </td>     </tr>     <tr>       <td> EU </td>       <td> https://api.eu.sumologic.com/api/ </td>     </tr>     <tr>       <td> FED </td>       <td> https://api.fed.sumologic.com/api/ </td>     </tr>     <tr>       <td> IN </td>       <td> https://api.in.sumologic.com/api/ </td>     </tr>     <tr>       <td> JP </td>       <td> https://api.jp.sumologic.com/api/ </td>     </tr>     <tr>       <td> US1 </td>       <td> https://api.sumologic.com/api/ </td>     </tr>     <tr>       <td> US2 </td>       <td> https://api.us2.sumologic.com/api/ </td>     </tr>   </table>  ## Authentication Sumo Logic supports the following options for API authentication: - Access ID and Access Key - Base64 encoded Access ID and Access Key  See [Access Keys](https://help.sumologic.com/docs/manage/security/access-keys) to generate an Access Key. Make sure to copy the key you create, because it is displayed only once. When you have an Access ID and Access Key you can execute requests such as the following:   ```bash   curl -u \"<accessId>:<accessKey>\" -X GET https://api.<deployment>.sumologic.com/api/v1/users   ```  Where `deployment` is either `au`, `ca`, `de`, `eu`, `fed`, `in`, `jp`, `us1`, or `us2`. See [API endpoints](#section/API-Endpoints) for details.  If you prefer to use basic access authentication, you can do a Base64 encoding of your `<accessId>:<accessKey>` to authenticate your HTTPS request. The following is an example request, replace the placeholder `<encoded>` with your encoded Access ID and Access Key string:   ```bash   curl -H \"Authorization: Basic <encoded>\" -X GET https://api.<deployment>.sumologic.com/api/v1/users   ```   Refer to [API Authentication](https://help.sumologic.com/?cid=3012) for a Base64 example.  ## Status Codes Generic status codes that apply to all our APIs. See the [HTTP status code registry](https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml) for reference.   <table>     <tr>       <td> <strong>HTTP Status Code</strong> </td>       <td> <strong>Error Code</strong> </td>       <td> <strong>Description</strong> </td>     </tr>     <tr>       <td> 301 </td>       <td> moved </td>       <td> The requested resource SHOULD be accessed through returned URI in Location Header. See [troubleshooting](https://help.sumologic.com/docs/api/troubleshooting/#api---301-error---moved) for details.</td>     </tr>     <tr>       <td> 401 </td>       <td> unauthorized </td>       <td> Credential could not be verified.</td>     </tr>     <tr>       <td> 403 </td>       <td> forbidden </td>       <td> This operation is not allowed for your account type or the user doesn't have the role capability to perform this action. See [troubleshooting](https://help.sumologic.com/docs/api/troubleshooting/#api---401-error---credential-could-not-be-verified) for details.</td>     </tr>     <tr>       <td> 404 </td>       <td> notfound </td>       <td> Requested resource could not be found. </td>     </tr>     <tr>       <td> 405 </td>       <td> method.unsupported </td>       <td> Unsupported method for URL. </td>     </tr>     <tr>       <td> 415 </td>       <td> contenttype.invalid </td>       <td> Invalid content type. </td>     </tr>     <tr>       <td> 429 </td>       <td> rate.limit.exceeded </td>       <td> The API request rate is higher than 4 request per second or inflight API requests are higher than 10 request per second. </td>     </tr>     <tr>       <td> 500 </td>       <td> internal.error </td>       <td> Internal server error. </td>     </tr>     <tr>       <td> 503 </td>       <td> service.unavailable </td>       <td> Service is currently unavailable. </td>     </tr>   </table>  ## Filtering Some API endpoints support filtering results on a specified set of fields. Each endpoint that supports filtering will list the fields that can be filtered. Multiple fields can be combined by using an ampersand `&` character.  For example, to get 20 users whose `firstName` is `John` and `lastName` is `Doe`:   ```bash   api.sumologic.com/v1/users?limit=20&firstName=John&lastName=Doe   ```  ## Sorting Some API endpoints support sorting fields by using the `sortBy` query parameter. The default sort order is ascending. Prefix the field with a minus sign `-` to sort in descending order.  For example, to get 20 users sorted by their `email` in descending order:   ```bash   api.sumologic.com/v1/users?limit=20&sort=-email   ```  ## Asynchronous Request Asynchronous requests do not wait for results, instead they immediately respond back with a job identifier while the job runs in the background. You can use the job identifier to track the status of the asynchronous job request. Here is a typical flow for an asynchronous request. 1. Start an asynchronous job. On success, a job identifier is returned. The job identifier uniquely identifies   your asynchronous job.  2. Once started, use the job identifier from step 1 to track the status of your asynchronous job. An asynchronous   request will typically provide an endpoint to poll for the status of asynchronous job. A successful response   from the status endpoint will have the following structure:   ```json   {       \"status\": \"Status of asynchronous request\",       \"statusMessage\": \"Optional message with additional information in case request succeeds\",       \"error\": \"Error object in case request fails\"   }   ```   The `status` field can have one of the following values:     1. `Success`: The job succeeded. The `statusMessage` field might have additional information.     2. `InProgress`: The job is still running.     3. `Failed`: The job failed. The `error` field in the response will have more information about the failure.  3. Some asynchronous APIs may provide a third endpoint (like [export result](#operation/getAsyncExportResult))   to fetch the result of an asynchronous job.   ### Example Let's say we want to export a folder with the identifier `0000000006A2E86F`. We will use the [async export](#operation/beginAsyncExport) API to export all the content under the folder with `id=0000000006A2E86F`. 1. Start an export job for the folder   ```bash   curl -X POST -u \"<accessId>:<accessKey>\" https://api.<deployment>.sumologic.com/api/v2/content/0000000006A2E86F/export   ```   See [authentication section](#section/Authentication) for more details about `accessId`, `accessKey`, and   `deployment`.   On success, you will get back a job identifier. In the response below, `C03E086C137F38B4` is the job identifier.   ```bash   {       \"id\": \"C03E086C137F38B4\"   }   ```  2. Now poll for the status of the asynchronous job with the [status](#operation/getAsyncExportStatus) endpoint.   ```bash   curl -X GET -u \"<accessId>:<accessKey>\" https://api.<deployment>.sumologic.com/api/v2/content/0000000006A2E86F/export/C03E086C137F38B4/status   ```   You may get a response like   ```json   {       \"status\": \"InProgress\",       \"statusMessage\": null,       \"error\": null   }   ```   It implies the job is still in progress. Keep polling till the status is either `Success` or `Failed`.  3. When the asynchronous job completes (`status != \"InProgress\"`), you can fetch the results with the   [export result](#operation/getAsyncExportResult) endpoint.   ```bash   curl -X GET -u \"<accessId>:<accessKey>\" https://api.<deployment>.sumologic.com/api/v2/content/0000000006A2E86F/export/C03E086C137F38B4/result   ```    The asynchronous job may fail (`status == \"Failed\"`). You can look at the `error` field for more details.   ```json   {       \"status\": \"Failed\",       \"errors\": {           \"code\": \"content1:too_many_items\",           \"message\": \"Too many objects: object count(1100) was greater than limit 1000\"       }   }   ```   ## Rate Limiting * A rate limit of four API requests per second (240 requests per minute) applies to all API calls from a user. * A rate limit of 10 concurrent requests to any API endpoint applies to an access key.  If a rate is exceeded, a rate limit exceeded 429 status code is returned.  ## Generating Clients You can use [OpenAPI Generator](https://openapi-generator.tech) to generate clients from the YAML file to access the API.  ### Using [NPM](https://www.npmjs.com/get-npm) 1. Install [NPM package wrapper](https://github.com/openapitools/openapi-generator-cli) globally, exposing the CLI   on the command line:   ```bash   npm install @openapitools/openapi-generator-cli -g   ```   You can see detailed instructions [here](https://openapi-generator.tech/docs/installation#npm).  2. Download the [YAML file](/docs/sumologic-api.yaml) and save it locally. Let's say the file is saved as `sumologic-api.yaml`. 3. Use the following command to generate `python` client inside the `sumo/client/python` directory:   ```bash   openapi-generator generate -i sumologic-api.yaml -g python -o sumo/client/python   ```   ### Using [Homebrew](https://brew.sh/) 1. Install OpenAPI Generator   ```bash   brew install openapi-generator   ```  2. Download the [YAML file](/docs/sumologic-api.yaml) and save it locally. Let's say the file is saved as `sumologic-api.yaml`. 3. Use the following command to generate `python` client side code inside the `sumo/client/python` directory:   ```bash   openapi-generator generate -i sumologic-api.yaml -g python -o sumo/client/python   ``` 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// checks if the UpdateRoleDefinitionV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateRoleDefinitionV2{}

// UpdateRoleDefinitionV2 struct for UpdateRoleDefinitionV2
type UpdateRoleDefinitionV2 struct {
	// Name of the role.
	Name string `json:"name"`
	// Description of the role.
	Description string `json:"description"`
	// A search filter which would be applied on partitions which belong to Log Analytics product area. Applicable with only `All` selectionType.
	LogAnalyticsFilter string `json:"logAnalyticsFilter"`
	// A search filter which would be applied on partitions which belong to Audit Data product area. Help Doc : (https://help.sumologic.com/docs/manage/security/audit-index/). Applicable with only `All` selectionType.
	AuditDataFilter string `json:"auditDataFilter"`
	// A search filter which would be applied on partitions which belong to Security Data product area. Applicable with only `All` selectionType.
	SecurityDataFilter string `json:"securityDataFilter"`
	// Describes the Permission Construct for the list of views in \"selectedViews\" parameter.  ### Valid Values are :    - `All` selectionType would allow access to all views in the org.   - `Allow` selectionType would allow access to specific views mentioned in \"selectedViews\" parameter.   - `Deny` selectionType would deny access to specific views mentioned in \"selectedViews\" parameter.
	SelectionType string `json:"selectionType"`
	// List of views which with specific view level filters in accordance to the selectionType chosen.
	SelectedViews []ViewFilterDefinition `json:"selectedViews"`
	// List of user identifiers to assign the role to.
	Users []string `json:"users"`
	// List of [capabilities](https://help.sumologic.com/Manage/Users-and-Roles/Manage-Roles/Role-Capabilities) associated with this role. Valid values are ### Data Management   - viewCollectors   - manageCollectors   - manageBudgets   - manageDataVolumeFeed   - viewFieldExtraction   - manageFieldExtractionRules   - manageS3DataForwarding   - manageContent   - manageApps   - dataVolumeIndex   - manageConnections   - viewScheduledViews   - manageScheduledViews   - viewPartitions   - managePartitions   - viewFields   - manageFields   - viewAccountOverview   - manageTokens   - downloadSearchResults  ### Entity management   - manageEntityTypeConfig  ### Metrics   - metricsTransformation   - metricsExtraction   - metricsRules  ### Security   - managePasswordPolicy   - ipAllowlisting   - createAccessKeys   - manageAccessKeys   - manageSupportAccountAccess   - manageAuditDataFeed   - manageSaml   - shareDashboardOutsideOrg   - manageOrgSettings   - changeDataAccessLevel  ### Dashboards   - shareDashboardWorld   - shareDashboardAllowlist  ### UserManagement   - manageUsersAndRoles  ### Observability   - searchAuditIndex   - auditEventIndex  ### Cloud SIEM Enterprise   - viewCse  ### Alerting   - viewMonitorsV2   - manageMonitorsV2   - viewAlerts
	Capabilities []string `json:"capabilities"`
	// Set this to true if you want to automatically append all missing capability requirements. If set to false an error will be thrown if any capabilities are missing their dependencies.
	AutofillDependencies *bool `json:"autofillDependencies,omitempty"`
}

type _UpdateRoleDefinitionV2 UpdateRoleDefinitionV2

// NewUpdateRoleDefinitionV2 instantiates a new UpdateRoleDefinitionV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateRoleDefinitionV2(name string, description string, logAnalyticsFilter string, auditDataFilter string, securityDataFilter string, selectionType string, selectedViews []ViewFilterDefinition, users []string, capabilities []string) *UpdateRoleDefinitionV2 {
	this := UpdateRoleDefinitionV2{}
	this.Name = name
	this.Description = description
	this.LogAnalyticsFilter = logAnalyticsFilter
	this.AuditDataFilter = auditDataFilter
	this.SecurityDataFilter = securityDataFilter
	this.SelectionType = selectionType
	this.SelectedViews = selectedViews
	this.Users = users
	this.Capabilities = capabilities
	var autofillDependencies bool = true
	this.AutofillDependencies = &autofillDependencies
	return &this
}

// NewUpdateRoleDefinitionV2WithDefaults instantiates a new UpdateRoleDefinitionV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateRoleDefinitionV2WithDefaults() *UpdateRoleDefinitionV2 {
	this := UpdateRoleDefinitionV2{}
	var autofillDependencies bool = true
	this.AutofillDependencies = &autofillDependencies
	return &this
}

// GetName returns the Name field value
func (o *UpdateRoleDefinitionV2) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateRoleDefinitionV2) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateRoleDefinitionV2) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *UpdateRoleDefinitionV2) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *UpdateRoleDefinitionV2) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *UpdateRoleDefinitionV2) SetDescription(v string) {
	o.Description = v
}

// GetLogAnalyticsFilter returns the LogAnalyticsFilter field value
func (o *UpdateRoleDefinitionV2) GetLogAnalyticsFilter() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LogAnalyticsFilter
}

// GetLogAnalyticsFilterOk returns a tuple with the LogAnalyticsFilter field value
// and a boolean to check if the value has been set.
func (o *UpdateRoleDefinitionV2) GetLogAnalyticsFilterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogAnalyticsFilter, true
}

// SetLogAnalyticsFilter sets field value
func (o *UpdateRoleDefinitionV2) SetLogAnalyticsFilter(v string) {
	o.LogAnalyticsFilter = v
}

// GetAuditDataFilter returns the AuditDataFilter field value
func (o *UpdateRoleDefinitionV2) GetAuditDataFilter() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuditDataFilter
}

// GetAuditDataFilterOk returns a tuple with the AuditDataFilter field value
// and a boolean to check if the value has been set.
func (o *UpdateRoleDefinitionV2) GetAuditDataFilterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuditDataFilter, true
}

// SetAuditDataFilter sets field value
func (o *UpdateRoleDefinitionV2) SetAuditDataFilter(v string) {
	o.AuditDataFilter = v
}

// GetSecurityDataFilter returns the SecurityDataFilter field value
func (o *UpdateRoleDefinitionV2) GetSecurityDataFilter() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecurityDataFilter
}

// GetSecurityDataFilterOk returns a tuple with the SecurityDataFilter field value
// and a boolean to check if the value has been set.
func (o *UpdateRoleDefinitionV2) GetSecurityDataFilterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecurityDataFilter, true
}

// SetSecurityDataFilter sets field value
func (o *UpdateRoleDefinitionV2) SetSecurityDataFilter(v string) {
	o.SecurityDataFilter = v
}

// GetSelectionType returns the SelectionType field value
func (o *UpdateRoleDefinitionV2) GetSelectionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SelectionType
}

// GetSelectionTypeOk returns a tuple with the SelectionType field value
// and a boolean to check if the value has been set.
func (o *UpdateRoleDefinitionV2) GetSelectionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SelectionType, true
}

// SetSelectionType sets field value
func (o *UpdateRoleDefinitionV2) SetSelectionType(v string) {
	o.SelectionType = v
}

// GetSelectedViews returns the SelectedViews field value
func (o *UpdateRoleDefinitionV2) GetSelectedViews() []ViewFilterDefinition {
	if o == nil {
		var ret []ViewFilterDefinition
		return ret
	}

	return o.SelectedViews
}

// GetSelectedViewsOk returns a tuple with the SelectedViews field value
// and a boolean to check if the value has been set.
func (o *UpdateRoleDefinitionV2) GetSelectedViewsOk() ([]ViewFilterDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return o.SelectedViews, true
}

// SetSelectedViews sets field value
func (o *UpdateRoleDefinitionV2) SetSelectedViews(v []ViewFilterDefinition) {
	o.SelectedViews = v
}

// GetUsers returns the Users field value
func (o *UpdateRoleDefinitionV2) GetUsers() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Users
}

// GetUsersOk returns a tuple with the Users field value
// and a boolean to check if the value has been set.
func (o *UpdateRoleDefinitionV2) GetUsersOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Users, true
}

// SetUsers sets field value
func (o *UpdateRoleDefinitionV2) SetUsers(v []string) {
	o.Users = v
}

// GetCapabilities returns the Capabilities field value
func (o *UpdateRoleDefinitionV2) GetCapabilities() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value
// and a boolean to check if the value has been set.
func (o *UpdateRoleDefinitionV2) GetCapabilitiesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Capabilities, true
}

// SetCapabilities sets field value
func (o *UpdateRoleDefinitionV2) SetCapabilities(v []string) {
	o.Capabilities = v
}

// GetAutofillDependencies returns the AutofillDependencies field value if set, zero value otherwise.
func (o *UpdateRoleDefinitionV2) GetAutofillDependencies() bool {
	if o == nil || IsNil(o.AutofillDependencies) {
		var ret bool
		return ret
	}
	return *o.AutofillDependencies
}

// GetAutofillDependenciesOk returns a tuple with the AutofillDependencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRoleDefinitionV2) GetAutofillDependenciesOk() (*bool, bool) {
	if o == nil || IsNil(o.AutofillDependencies) {
		return nil, false
	}
	return o.AutofillDependencies, true
}

// HasAutofillDependencies returns a boolean if a field has been set.
func (o *UpdateRoleDefinitionV2) HasAutofillDependencies() bool {
	if o != nil && !IsNil(o.AutofillDependencies) {
		return true
	}

	return false
}

// SetAutofillDependencies gets a reference to the given bool and assigns it to the AutofillDependencies field.
func (o *UpdateRoleDefinitionV2) SetAutofillDependencies(v bool) {
	o.AutofillDependencies = &v
}

func (o UpdateRoleDefinitionV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateRoleDefinitionV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	toSerialize["logAnalyticsFilter"] = o.LogAnalyticsFilter
	toSerialize["auditDataFilter"] = o.AuditDataFilter
	toSerialize["securityDataFilter"] = o.SecurityDataFilter
	toSerialize["selectionType"] = o.SelectionType
	toSerialize["selectedViews"] = o.SelectedViews
	toSerialize["users"] = o.Users
	toSerialize["capabilities"] = o.Capabilities
	if !IsNil(o.AutofillDependencies) {
		toSerialize["autofillDependencies"] = o.AutofillDependencies
	}
	return toSerialize, nil
}

func (o *UpdateRoleDefinitionV2) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"description",
		"logAnalyticsFilter",
		"auditDataFilter",
		"securityDataFilter",
		"selectionType",
		"selectedViews",
		"users",
		"capabilities",
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

	varUpdateRoleDefinitionV2 := _UpdateRoleDefinitionV2{}

	err = json.Unmarshal(bytes, &varUpdateRoleDefinitionV2)

	if err != nil {
		return err
	}

	*o = UpdateRoleDefinitionV2(varUpdateRoleDefinitionV2)

	return err
}

type NullableUpdateRoleDefinitionV2 struct {
	value *UpdateRoleDefinitionV2
	isSet bool
}

func (v NullableUpdateRoleDefinitionV2) Get() *UpdateRoleDefinitionV2 {
	return v.value
}

func (v *NullableUpdateRoleDefinitionV2) Set(val *UpdateRoleDefinitionV2) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateRoleDefinitionV2) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateRoleDefinitionV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateRoleDefinitionV2(val *UpdateRoleDefinitionV2) *NullableUpdateRoleDefinitionV2 {
	return &NullableUpdateRoleDefinitionV2{value: val, isSet: true}
}

func (v NullableUpdateRoleDefinitionV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateRoleDefinitionV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


