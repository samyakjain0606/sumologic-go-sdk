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

// checks if the OTCollector type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OTCollector{}

// OTCollector An OT Collector definition.
type OTCollector struct {
	// Unique identifier of the OT Collector.
	Id string `json:"id"`
	// Name of the OT Collector.
	Name string `json:"name"`
	Version OTCollectorVersion `json:"version"`
	// Category of the OT Collector.
	Category *string `json:"category,omitempty"`
	// Description of the OT Collector.
	Description *string `json:"description,omitempty"`
	// Tags associated with the OT Collector.
	Tags *map[string]string `json:"tags,omitempty"`
	HealthIncidentsTracker *OTCollectorHealthIncidentsTracker `json:"healthIncidentsTracker,omitempty"`
	// Ephemeral Status of the OT Collector.
	Ephemeral *bool `json:"ephemeral,omitempty"`
	// Alive Status of the OT Collector based on heartbeat.
	Alive *bool `json:"alive,omitempty"`
	// Management Status of the OT Collector based on if it is remotely or locally managed.
	IsRemotelyManaged *bool `json:"isRemotelyManaged,omitempty"`
	// Config map that includes Base 64 Encoded Effective Configuration Yaml of the Remotely managed OT Collector.
	EffectiveConfig *map[string]string `json:"effectiveConfig,omitempty"`
	SystemInfo OTCollectorSystemInfo `json:"systemInfo"`
	// timezone of the collector
	TimeZone *string `json:"timeZone,omitempty"`
	// Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format.
	CreatedAt time.Time `json:"createdAt"`
	// Identifier of the user who created the resource.
	CreatedBy string `json:"createdBy"`
	// Last modification timestamp in UTC.
	ModifiedAt time.Time `json:"modifiedAt"`
	// Identifier of the user who last modified the resource.
	ModifiedBy string `json:"modifiedBy"`
	// Count of the source templates linked to a collector
	SourceTemplateLinkedCount *int32 `json:"sourceTemplateLinkedCount,omitempty"`
}

type _OTCollector OTCollector

// NewOTCollector instantiates a new OTCollector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOTCollector(id string, name string, version OTCollectorVersion, systemInfo OTCollectorSystemInfo, createdAt time.Time, createdBy string, modifiedAt time.Time, modifiedBy string) *OTCollector {
	this := OTCollector{}
	this.Id = id
	this.Name = name
	this.Version = version
	this.SystemInfo = systemInfo
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.ModifiedAt = modifiedAt
	this.ModifiedBy = modifiedBy
	return &this
}

// NewOTCollectorWithDefaults instantiates a new OTCollector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOTCollectorWithDefaults() *OTCollector {
	this := OTCollector{}
	return &this
}

// GetId returns the Id field value
func (o *OTCollector) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OTCollector) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OTCollector) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *OTCollector) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OTCollector) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OTCollector) SetName(v string) {
	o.Name = v
}

// GetVersion returns the Version field value
func (o *OTCollector) GetVersion() OTCollectorVersion {
	if o == nil {
		var ret OTCollectorVersion
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *OTCollector) GetVersionOk() (*OTCollectorVersion, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *OTCollector) SetVersion(v OTCollectorVersion) {
	o.Version = v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *OTCollector) GetCategory() string {
	if o == nil || IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OTCollector) GetCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *OTCollector) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *OTCollector) SetCategory(v string) {
	o.Category = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OTCollector) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OTCollector) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OTCollector) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OTCollector) SetDescription(v string) {
	o.Description = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *OTCollector) GetTags() map[string]string {
	if o == nil || IsNil(o.Tags) {
		var ret map[string]string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OTCollector) GetTagsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *OTCollector) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *OTCollector) SetTags(v map[string]string) {
	o.Tags = &v
}

// GetHealthIncidentsTracker returns the HealthIncidentsTracker field value if set, zero value otherwise.
func (o *OTCollector) GetHealthIncidentsTracker() OTCollectorHealthIncidentsTracker {
	if o == nil || IsNil(o.HealthIncidentsTracker) {
		var ret OTCollectorHealthIncidentsTracker
		return ret
	}
	return *o.HealthIncidentsTracker
}

// GetHealthIncidentsTrackerOk returns a tuple with the HealthIncidentsTracker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OTCollector) GetHealthIncidentsTrackerOk() (*OTCollectorHealthIncidentsTracker, bool) {
	if o == nil || IsNil(o.HealthIncidentsTracker) {
		return nil, false
	}
	return o.HealthIncidentsTracker, true
}

// HasHealthIncidentsTracker returns a boolean if a field has been set.
func (o *OTCollector) HasHealthIncidentsTracker() bool {
	if o != nil && !IsNil(o.HealthIncidentsTracker) {
		return true
	}

	return false
}

// SetHealthIncidentsTracker gets a reference to the given OTCollectorHealthIncidentsTracker and assigns it to the HealthIncidentsTracker field.
func (o *OTCollector) SetHealthIncidentsTracker(v OTCollectorHealthIncidentsTracker) {
	o.HealthIncidentsTracker = &v
}

// GetEphemeral returns the Ephemeral field value if set, zero value otherwise.
func (o *OTCollector) GetEphemeral() bool {
	if o == nil || IsNil(o.Ephemeral) {
		var ret bool
		return ret
	}
	return *o.Ephemeral
}

// GetEphemeralOk returns a tuple with the Ephemeral field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OTCollector) GetEphemeralOk() (*bool, bool) {
	if o == nil || IsNil(o.Ephemeral) {
		return nil, false
	}
	return o.Ephemeral, true
}

// HasEphemeral returns a boolean if a field has been set.
func (o *OTCollector) HasEphemeral() bool {
	if o != nil && !IsNil(o.Ephemeral) {
		return true
	}

	return false
}

// SetEphemeral gets a reference to the given bool and assigns it to the Ephemeral field.
func (o *OTCollector) SetEphemeral(v bool) {
	o.Ephemeral = &v
}

// GetAlive returns the Alive field value if set, zero value otherwise.
func (o *OTCollector) GetAlive() bool {
	if o == nil || IsNil(o.Alive) {
		var ret bool
		return ret
	}
	return *o.Alive
}

// GetAliveOk returns a tuple with the Alive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OTCollector) GetAliveOk() (*bool, bool) {
	if o == nil || IsNil(o.Alive) {
		return nil, false
	}
	return o.Alive, true
}

// HasAlive returns a boolean if a field has been set.
func (o *OTCollector) HasAlive() bool {
	if o != nil && !IsNil(o.Alive) {
		return true
	}

	return false
}

// SetAlive gets a reference to the given bool and assigns it to the Alive field.
func (o *OTCollector) SetAlive(v bool) {
	o.Alive = &v
}

// GetIsRemotelyManaged returns the IsRemotelyManaged field value if set, zero value otherwise.
func (o *OTCollector) GetIsRemotelyManaged() bool {
	if o == nil || IsNil(o.IsRemotelyManaged) {
		var ret bool
		return ret
	}
	return *o.IsRemotelyManaged
}

// GetIsRemotelyManagedOk returns a tuple with the IsRemotelyManaged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OTCollector) GetIsRemotelyManagedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsRemotelyManaged) {
		return nil, false
	}
	return o.IsRemotelyManaged, true
}

// HasIsRemotelyManaged returns a boolean if a field has been set.
func (o *OTCollector) HasIsRemotelyManaged() bool {
	if o != nil && !IsNil(o.IsRemotelyManaged) {
		return true
	}

	return false
}

// SetIsRemotelyManaged gets a reference to the given bool and assigns it to the IsRemotelyManaged field.
func (o *OTCollector) SetIsRemotelyManaged(v bool) {
	o.IsRemotelyManaged = &v
}

// GetEffectiveConfig returns the EffectiveConfig field value if set, zero value otherwise.
func (o *OTCollector) GetEffectiveConfig() map[string]string {
	if o == nil || IsNil(o.EffectiveConfig) {
		var ret map[string]string
		return ret
	}
	return *o.EffectiveConfig
}

// GetEffectiveConfigOk returns a tuple with the EffectiveConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OTCollector) GetEffectiveConfigOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.EffectiveConfig) {
		return nil, false
	}
	return o.EffectiveConfig, true
}

// HasEffectiveConfig returns a boolean if a field has been set.
func (o *OTCollector) HasEffectiveConfig() bool {
	if o != nil && !IsNil(o.EffectiveConfig) {
		return true
	}

	return false
}

// SetEffectiveConfig gets a reference to the given map[string]string and assigns it to the EffectiveConfig field.
func (o *OTCollector) SetEffectiveConfig(v map[string]string) {
	o.EffectiveConfig = &v
}

// GetSystemInfo returns the SystemInfo field value
func (o *OTCollector) GetSystemInfo() OTCollectorSystemInfo {
	if o == nil {
		var ret OTCollectorSystemInfo
		return ret
	}

	return o.SystemInfo
}

// GetSystemInfoOk returns a tuple with the SystemInfo field value
// and a boolean to check if the value has been set.
func (o *OTCollector) GetSystemInfoOk() (*OTCollectorSystemInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SystemInfo, true
}

// SetSystemInfo sets field value
func (o *OTCollector) SetSystemInfo(v OTCollectorSystemInfo) {
	o.SystemInfo = v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *OTCollector) GetTimeZone() string {
	if o == nil || IsNil(o.TimeZone) {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OTCollector) GetTimeZoneOk() (*string, bool) {
	if o == nil || IsNil(o.TimeZone) {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *OTCollector) HasTimeZone() bool {
	if o != nil && !IsNil(o.TimeZone) {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *OTCollector) SetTimeZone(v string) {
	o.TimeZone = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *OTCollector) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *OTCollector) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *OTCollector) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *OTCollector) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *OTCollector) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *OTCollector) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetModifiedAt returns the ModifiedAt field value
func (o *OTCollector) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *OTCollector) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value
func (o *OTCollector) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// GetModifiedBy returns the ModifiedBy field value
func (o *OTCollector) GetModifiedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value
// and a boolean to check if the value has been set.
func (o *OTCollector) GetModifiedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedBy, true
}

// SetModifiedBy sets field value
func (o *OTCollector) SetModifiedBy(v string) {
	o.ModifiedBy = v
}

// GetSourceTemplateLinkedCount returns the SourceTemplateLinkedCount field value if set, zero value otherwise.
func (o *OTCollector) GetSourceTemplateLinkedCount() int32 {
	if o == nil || IsNil(o.SourceTemplateLinkedCount) {
		var ret int32
		return ret
	}
	return *o.SourceTemplateLinkedCount
}

// GetSourceTemplateLinkedCountOk returns a tuple with the SourceTemplateLinkedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OTCollector) GetSourceTemplateLinkedCountOk() (*int32, bool) {
	if o == nil || IsNil(o.SourceTemplateLinkedCount) {
		return nil, false
	}
	return o.SourceTemplateLinkedCount, true
}

// HasSourceTemplateLinkedCount returns a boolean if a field has been set.
func (o *OTCollector) HasSourceTemplateLinkedCount() bool {
	if o != nil && !IsNil(o.SourceTemplateLinkedCount) {
		return true
	}

	return false
}

// SetSourceTemplateLinkedCount gets a reference to the given int32 and assigns it to the SourceTemplateLinkedCount field.
func (o *OTCollector) SetSourceTemplateLinkedCount(v int32) {
	o.SourceTemplateLinkedCount = &v
}

func (o OTCollector) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OTCollector) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["version"] = o.Version
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.HealthIncidentsTracker) {
		toSerialize["healthIncidentsTracker"] = o.HealthIncidentsTracker
	}
	if !IsNil(o.Ephemeral) {
		toSerialize["ephemeral"] = o.Ephemeral
	}
	if !IsNil(o.Alive) {
		toSerialize["alive"] = o.Alive
	}
	if !IsNil(o.IsRemotelyManaged) {
		toSerialize["isRemotelyManaged"] = o.IsRemotelyManaged
	}
	if !IsNil(o.EffectiveConfig) {
		toSerialize["effectiveConfig"] = o.EffectiveConfig
	}
	toSerialize["systemInfo"] = o.SystemInfo
	if !IsNil(o.TimeZone) {
		toSerialize["timeZone"] = o.TimeZone
	}
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["createdBy"] = o.CreatedBy
	toSerialize["modifiedAt"] = o.ModifiedAt
	toSerialize["modifiedBy"] = o.ModifiedBy
	if !IsNil(o.SourceTemplateLinkedCount) {
		toSerialize["sourceTemplateLinkedCount"] = o.SourceTemplateLinkedCount
	}
	return toSerialize, nil
}

func (o *OTCollector) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"version",
		"systemInfo",
		"createdAt",
		"createdBy",
		"modifiedAt",
		"modifiedBy",
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

	varOTCollector := _OTCollector{}

	err = json.Unmarshal(bytes, &varOTCollector)

	if err != nil {
		return err
	}

	*o = OTCollector(varOTCollector)

	return err
}

type NullableOTCollector struct {
	value *OTCollector
	isSet bool
}

func (v NullableOTCollector) Get() *OTCollector {
	return v.value
}

func (v *NullableOTCollector) Set(val *OTCollector) {
	v.value = val
	v.isSet = true
}

func (v NullableOTCollector) IsSet() bool {
	return v.isSet
}

func (v *NullableOTCollector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOTCollector(val *OTCollector) *NullableOTCollector {
	return &NullableOTCollector{value: val, isSet: true}
}

func (v NullableOTCollector) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOTCollector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


