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

// checks if the DashboardRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DashboardRequest{}

// DashboardRequest struct for DashboardRequest
type DashboardRequest struct {
	// Title of the dashboard.
	Title string `json:"title"`
	// Description of the dashboard.
	Description *string `json:"description,omitempty"`
	// The identifier of the folder to save the dashboard in. By default it is saved in your personal folder. 
	FolderId *string `json:"folderId,omitempty"`
	TopologyLabelMap *TopologyLabelMap `json:"topologyLabelMap,omitempty"`
	// If set denotes that the dashboard concerns a given domain (e.g. `aws`, `k8s`, `app`).
	Domain *string `json:"domain,omitempty"`
	// If set to non-empty array denotes that the dashboard concerns given hierarchies.
	Hierarchies []string `json:"hierarchies,omitempty"`
	// Interval of time (in seconds) to automatically refresh the dashboard. A value of 0 means we never automatically refresh the dashboard. Allowed values are `0`, `30`, `60`, `120`, `300`, `900`, `1800`, `3600`, `7200`, `86400`. 
	RefreshInterval *int32 `json:"refreshInterval,omitempty"`
	TimeRange ResolvableTimeRange `json:"timeRange"`
	// Panels in the dashboard.
	Panels []Panel `json:"panels,omitempty"`
	Layout *Layout `json:"layout,omitempty"`
	// Variables to apply to the panels.
	Variables []Variable `json:"variables,omitempty"`
	// Theme for the dashboard. Either `Light` or `Dark`.
	Theme *string `json:"theme,omitempty"`
	// Is the dashboard public
	IsPublic *bool `json:"isPublic,omitempty"`
	// Whether to highlight threshold violations.
	HighlightViolations *bool `json:"highlightViolations,omitempty"`
}

type _DashboardRequest DashboardRequest

// NewDashboardRequest instantiates a new DashboardRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDashboardRequest(title string, timeRange ResolvableTimeRange) *DashboardRequest {
	this := DashboardRequest{}
	this.Title = title
	var domain string = ""
	this.Domain = &domain
	this.TimeRange = timeRange
	var theme string = "Light"
	this.Theme = &theme
	var isPublic bool = false
	this.IsPublic = &isPublic
	var highlightViolations bool = false
	this.HighlightViolations = &highlightViolations
	return &this
}

// NewDashboardRequestWithDefaults instantiates a new DashboardRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDashboardRequestWithDefaults() *DashboardRequest {
	this := DashboardRequest{}
	var domain string = ""
	this.Domain = &domain
	var theme string = "Light"
	this.Theme = &theme
	var isPublic bool = false
	this.IsPublic = &isPublic
	var highlightViolations bool = false
	this.HighlightViolations = &highlightViolations
	return &this
}

// GetTitle returns the Title field value
func (o *DashboardRequest) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *DashboardRequest) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *DashboardRequest) SetTitle(v string) {
	o.Title = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DashboardRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DashboardRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DashboardRequest) SetDescription(v string) {
	o.Description = &v
}

// GetFolderId returns the FolderId field value if set, zero value otherwise.
func (o *DashboardRequest) GetFolderId() string {
	if o == nil || IsNil(o.FolderId) {
		var ret string
		return ret
	}
	return *o.FolderId
}

// GetFolderIdOk returns a tuple with the FolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardRequest) GetFolderIdOk() (*string, bool) {
	if o == nil || IsNil(o.FolderId) {
		return nil, false
	}
	return o.FolderId, true
}

// HasFolderId returns a boolean if a field has been set.
func (o *DashboardRequest) HasFolderId() bool {
	if o != nil && !IsNil(o.FolderId) {
		return true
	}

	return false
}

// SetFolderId gets a reference to the given string and assigns it to the FolderId field.
func (o *DashboardRequest) SetFolderId(v string) {
	o.FolderId = &v
}

// GetTopologyLabelMap returns the TopologyLabelMap field value if set, zero value otherwise.
func (o *DashboardRequest) GetTopologyLabelMap() TopologyLabelMap {
	if o == nil || IsNil(o.TopologyLabelMap) {
		var ret TopologyLabelMap
		return ret
	}
	return *o.TopologyLabelMap
}

// GetTopologyLabelMapOk returns a tuple with the TopologyLabelMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardRequest) GetTopologyLabelMapOk() (*TopologyLabelMap, bool) {
	if o == nil || IsNil(o.TopologyLabelMap) {
		return nil, false
	}
	return o.TopologyLabelMap, true
}

// HasTopologyLabelMap returns a boolean if a field has been set.
func (o *DashboardRequest) HasTopologyLabelMap() bool {
	if o != nil && !IsNil(o.TopologyLabelMap) {
		return true
	}

	return false
}

// SetTopologyLabelMap gets a reference to the given TopologyLabelMap and assigns it to the TopologyLabelMap field.
func (o *DashboardRequest) SetTopologyLabelMap(v TopologyLabelMap) {
	o.TopologyLabelMap = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *DashboardRequest) GetDomain() string {
	if o == nil || IsNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardRequest) GetDomainOk() (*string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *DashboardRequest) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *DashboardRequest) SetDomain(v string) {
	o.Domain = &v
}

// GetHierarchies returns the Hierarchies field value if set, zero value otherwise.
func (o *DashboardRequest) GetHierarchies() []string {
	if o == nil || IsNil(o.Hierarchies) {
		var ret []string
		return ret
	}
	return o.Hierarchies
}

// GetHierarchiesOk returns a tuple with the Hierarchies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardRequest) GetHierarchiesOk() ([]string, bool) {
	if o == nil || IsNil(o.Hierarchies) {
		return nil, false
	}
	return o.Hierarchies, true
}

// HasHierarchies returns a boolean if a field has been set.
func (o *DashboardRequest) HasHierarchies() bool {
	if o != nil && !IsNil(o.Hierarchies) {
		return true
	}

	return false
}

// SetHierarchies gets a reference to the given []string and assigns it to the Hierarchies field.
func (o *DashboardRequest) SetHierarchies(v []string) {
	o.Hierarchies = v
}

// GetRefreshInterval returns the RefreshInterval field value if set, zero value otherwise.
func (o *DashboardRequest) GetRefreshInterval() int32 {
	if o == nil || IsNil(o.RefreshInterval) {
		var ret int32
		return ret
	}
	return *o.RefreshInterval
}

// GetRefreshIntervalOk returns a tuple with the RefreshInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardRequest) GetRefreshIntervalOk() (*int32, bool) {
	if o == nil || IsNil(o.RefreshInterval) {
		return nil, false
	}
	return o.RefreshInterval, true
}

// HasRefreshInterval returns a boolean if a field has been set.
func (o *DashboardRequest) HasRefreshInterval() bool {
	if o != nil && !IsNil(o.RefreshInterval) {
		return true
	}

	return false
}

// SetRefreshInterval gets a reference to the given int32 and assigns it to the RefreshInterval field.
func (o *DashboardRequest) SetRefreshInterval(v int32) {
	o.RefreshInterval = &v
}

// GetTimeRange returns the TimeRange field value
func (o *DashboardRequest) GetTimeRange() ResolvableTimeRange {
	if o == nil {
		var ret ResolvableTimeRange
		return ret
	}

	return o.TimeRange
}

// GetTimeRangeOk returns a tuple with the TimeRange field value
// and a boolean to check if the value has been set.
func (o *DashboardRequest) GetTimeRangeOk() (*ResolvableTimeRange, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeRange, true
}

// SetTimeRange sets field value
func (o *DashboardRequest) SetTimeRange(v ResolvableTimeRange) {
	o.TimeRange = v
}

// GetPanels returns the Panels field value if set, zero value otherwise.
func (o *DashboardRequest) GetPanels() []Panel {
	if o == nil || IsNil(o.Panels) {
		var ret []Panel
		return ret
	}
	return o.Panels
}

// GetPanelsOk returns a tuple with the Panels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardRequest) GetPanelsOk() ([]Panel, bool) {
	if o == nil || IsNil(o.Panels) {
		return nil, false
	}
	return o.Panels, true
}

// HasPanels returns a boolean if a field has been set.
func (o *DashboardRequest) HasPanels() bool {
	if o != nil && !IsNil(o.Panels) {
		return true
	}

	return false
}

// SetPanels gets a reference to the given []Panel and assigns it to the Panels field.
func (o *DashboardRequest) SetPanels(v []Panel) {
	o.Panels = v
}

// GetLayout returns the Layout field value if set, zero value otherwise.
func (o *DashboardRequest) GetLayout() Layout {
	if o == nil || IsNil(o.Layout) {
		var ret Layout
		return ret
	}
	return *o.Layout
}

// GetLayoutOk returns a tuple with the Layout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardRequest) GetLayoutOk() (*Layout, bool) {
	if o == nil || IsNil(o.Layout) {
		return nil, false
	}
	return o.Layout, true
}

// HasLayout returns a boolean if a field has been set.
func (o *DashboardRequest) HasLayout() bool {
	if o != nil && !IsNil(o.Layout) {
		return true
	}

	return false
}

// SetLayout gets a reference to the given Layout and assigns it to the Layout field.
func (o *DashboardRequest) SetLayout(v Layout) {
	o.Layout = &v
}

// GetVariables returns the Variables field value if set, zero value otherwise.
func (o *DashboardRequest) GetVariables() []Variable {
	if o == nil || IsNil(o.Variables) {
		var ret []Variable
		return ret
	}
	return o.Variables
}

// GetVariablesOk returns a tuple with the Variables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardRequest) GetVariablesOk() ([]Variable, bool) {
	if o == nil || IsNil(o.Variables) {
		return nil, false
	}
	return o.Variables, true
}

// HasVariables returns a boolean if a field has been set.
func (o *DashboardRequest) HasVariables() bool {
	if o != nil && !IsNil(o.Variables) {
		return true
	}

	return false
}

// SetVariables gets a reference to the given []Variable and assigns it to the Variables field.
func (o *DashboardRequest) SetVariables(v []Variable) {
	o.Variables = v
}

// GetTheme returns the Theme field value if set, zero value otherwise.
func (o *DashboardRequest) GetTheme() string {
	if o == nil || IsNil(o.Theme) {
		var ret string
		return ret
	}
	return *o.Theme
}

// GetThemeOk returns a tuple with the Theme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardRequest) GetThemeOk() (*string, bool) {
	if o == nil || IsNil(o.Theme) {
		return nil, false
	}
	return o.Theme, true
}

// HasTheme returns a boolean if a field has been set.
func (o *DashboardRequest) HasTheme() bool {
	if o != nil && !IsNil(o.Theme) {
		return true
	}

	return false
}

// SetTheme gets a reference to the given string and assigns it to the Theme field.
func (o *DashboardRequest) SetTheme(v string) {
	o.Theme = &v
}

// GetIsPublic returns the IsPublic field value if set, zero value otherwise.
func (o *DashboardRequest) GetIsPublic() bool {
	if o == nil || IsNil(o.IsPublic) {
		var ret bool
		return ret
	}
	return *o.IsPublic
}

// GetIsPublicOk returns a tuple with the IsPublic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardRequest) GetIsPublicOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPublic) {
		return nil, false
	}
	return o.IsPublic, true
}

// HasIsPublic returns a boolean if a field has been set.
func (o *DashboardRequest) HasIsPublic() bool {
	if o != nil && !IsNil(o.IsPublic) {
		return true
	}

	return false
}

// SetIsPublic gets a reference to the given bool and assigns it to the IsPublic field.
func (o *DashboardRequest) SetIsPublic(v bool) {
	o.IsPublic = &v
}

// GetHighlightViolations returns the HighlightViolations field value if set, zero value otherwise.
func (o *DashboardRequest) GetHighlightViolations() bool {
	if o == nil || IsNil(o.HighlightViolations) {
		var ret bool
		return ret
	}
	return *o.HighlightViolations
}

// GetHighlightViolationsOk returns a tuple with the HighlightViolations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardRequest) GetHighlightViolationsOk() (*bool, bool) {
	if o == nil || IsNil(o.HighlightViolations) {
		return nil, false
	}
	return o.HighlightViolations, true
}

// HasHighlightViolations returns a boolean if a field has been set.
func (o *DashboardRequest) HasHighlightViolations() bool {
	if o != nil && !IsNil(o.HighlightViolations) {
		return true
	}

	return false
}

// SetHighlightViolations gets a reference to the given bool and assigns it to the HighlightViolations field.
func (o *DashboardRequest) SetHighlightViolations(v bool) {
	o.HighlightViolations = &v
}

func (o DashboardRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DashboardRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["title"] = o.Title
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.FolderId) {
		toSerialize["folderId"] = o.FolderId
	}
	if !IsNil(o.TopologyLabelMap) {
		toSerialize["topologyLabelMap"] = o.TopologyLabelMap
	}
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	if !IsNil(o.Hierarchies) {
		toSerialize["hierarchies"] = o.Hierarchies
	}
	if !IsNil(o.RefreshInterval) {
		toSerialize["refreshInterval"] = o.RefreshInterval
	}
	toSerialize["timeRange"] = o.TimeRange
	if !IsNil(o.Panels) {
		toSerialize["panels"] = o.Panels
	}
	if !IsNil(o.Layout) {
		toSerialize["layout"] = o.Layout
	}
	if !IsNil(o.Variables) {
		toSerialize["variables"] = o.Variables
	}
	if !IsNil(o.Theme) {
		toSerialize["theme"] = o.Theme
	}
	if !IsNil(o.IsPublic) {
		toSerialize["isPublic"] = o.IsPublic
	}
	if !IsNil(o.HighlightViolations) {
		toSerialize["highlightViolations"] = o.HighlightViolations
	}
	return toSerialize, nil
}

func (o *DashboardRequest) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"title",
		"timeRange",
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

	varDashboardRequest := _DashboardRequest{}

	err = json.Unmarshal(bytes, &varDashboardRequest)

	if err != nil {
		return err
	}

	*o = DashboardRequest(varDashboardRequest)

	return err
}

type NullableDashboardRequest struct {
	value *DashboardRequest
	isSet bool
}

func (v NullableDashboardRequest) Get() *DashboardRequest {
	return v.value
}

func (v *NullableDashboardRequest) Set(val *DashboardRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDashboardRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDashboardRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDashboardRequest(val *DashboardRequest) *NullableDashboardRequest {
	return &NullableDashboardRequest{value: val, isSet: true}
}

func (v NullableDashboardRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDashboardRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


