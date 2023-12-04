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

// checks if the AlertsLibraryAlertUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlertsLibraryAlertUpdate{}

// AlertsLibraryAlertUpdate struct for AlertsLibraryAlertUpdate
type AlertsLibraryAlertUpdate struct {
	AlertsLibraryBaseUpdate
	// The Id of the associated monitor.
	MonitorId *string `json:"monitorId,omitempty"`
	// The time at which the alert was resolved.
	ResolvedAt NullableTime `json:"resolvedAt,omitempty"`
	// The time at which the incident started.
	AbnormalityStartTime *time.Time `json:"abnormalityStartTime,omitempty"`
	// The severity of the Alert. Valid values:   1. `Critical`   2. `Warning`   3. `MissingData`
	AlertType *string `json:"alertType,omitempty"`
	// The status of the Alert. Valid values:   1. `Triggered`   2. `Resolved`
	Status *string `json:"status,omitempty"`
	// All queries from the monitor relevant to the alert.
	MonitorQueries []AlertMonitorQuery `json:"monitorQueries,omitempty"`
	// All queries from the monitor relevant to the alert with triggered time series filters.
	TriggerQueries []AlertMonitorQuery `json:"triggerQueries,omitempty"`
	// URL for this monitor's view page
	MonitorUrl *string `json:"monitorUrl,omitempty"`
	// A link to search with the triggering data and time range
	TriggerQueryUrl *string `json:"triggerQueryUrl,omitempty"`
	// Trigger conditions which were breached to create this Alert.
	TriggerConditions []TriggerCondition `json:"triggerConditions,omitempty"`
	// The of the query result which breached the trigger condition.
	TriggerValue *float64 `json:"triggerValue,omitempty"`
	// The type of monitor. Valid values:   1. `Logs`: A logs query monitor.   2. `Metrics`: A metrics query monitor.
	MonitorType *string `json:"monitorType,omitempty"`
	// One or more primary entity identifiers involved in this Alert. Primary/secondary entities are explained in description for `entities`. DEPRECATED, USE `entities` INSTEAD. 
	// Deprecated
	EntityIds []string `json:"entityIds,omitempty"`
	// One or more primary entities involved in this Alert. Primary entity is the most concrete entity that can be assigned per time series or log group (e.g. k8s container), secondary entities are the less specific ones that can be assigned per that notification (e.g. k8s cluster or EC2 host). 
	Entities []AlertEntityInfo `json:"entities,omitempty"`
	// One or more secondary entity involved in this Alert. Primary/secondary entities are explained in description for `entities` 
	SecondaryEntities []AlertEntityInfo `json:"secondaryEntities,omitempty"`
	Notes *string `json:"notes,omitempty"`
	ExtraDetails *ExtraDetails `json:"extraDetails,omitempty"`
	// The condition which triggered this alert.
	AlertCondition NullableString `json:"alertCondition,omitempty"`
	// Flag of the alerts muting status.
	IsMuted *bool `json:"isMuted,omitempty"`
}

type _AlertsLibraryAlertUpdate AlertsLibraryAlertUpdate

// NewAlertsLibraryAlertUpdate instantiates a new AlertsLibraryAlertUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertsLibraryAlertUpdate(name string, version int64, type_ string) *AlertsLibraryAlertUpdate {
	this := AlertsLibraryAlertUpdate{}
	this.Name = name
	var description string = ""
	this.Description = &description
	this.Version = version
	this.Type = type_
	return &this
}

// NewAlertsLibraryAlertUpdateWithDefaults instantiates a new AlertsLibraryAlertUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertsLibraryAlertUpdateWithDefaults() *AlertsLibraryAlertUpdate {
	this := AlertsLibraryAlertUpdate{}
	return &this
}

// GetMonitorId returns the MonitorId field value if set, zero value otherwise.
func (o *AlertsLibraryAlertUpdate) GetMonitorId() string {
	if o == nil || IsNil(o.MonitorId) {
		var ret string
		return ret
	}
	return *o.MonitorId
}

// GetMonitorIdOk returns a tuple with the MonitorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsLibraryAlertUpdate) GetMonitorIdOk() (*string, bool) {
	if o == nil || IsNil(o.MonitorId) {
		return nil, false
	}
	return o.MonitorId, true
}

// HasMonitorId returns a boolean if a field has been set.
func (o *AlertsLibraryAlertUpdate) HasMonitorId() bool {
	if o != nil && !IsNil(o.MonitorId) {
		return true
	}

	return false
}

// SetMonitorId gets a reference to the given string and assigns it to the MonitorId field.
func (o *AlertsLibraryAlertUpdate) SetMonitorId(v string) {
	o.MonitorId = &v
}

// GetResolvedAt returns the ResolvedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlertsLibraryAlertUpdate) GetResolvedAt() time.Time {
	if o == nil || IsNil(o.ResolvedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ResolvedAt.Get()
}

// GetResolvedAtOk returns a tuple with the ResolvedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlertsLibraryAlertUpdate) GetResolvedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResolvedAt.Get(), o.ResolvedAt.IsSet()
}

// HasResolvedAt returns a boolean if a field has been set.
func (o *AlertsLibraryAlertUpdate) HasResolvedAt() bool {
	if o != nil && o.ResolvedAt.IsSet() {
		return true
	}

	return false
}

// SetResolvedAt gets a reference to the given NullableTime and assigns it to the ResolvedAt field.
func (o *AlertsLibraryAlertUpdate) SetResolvedAt(v time.Time) {
	o.ResolvedAt.Set(&v)
}
// SetResolvedAtNil sets the value for ResolvedAt to be an explicit nil
func (o *AlertsLibraryAlertUpdate) SetResolvedAtNil() {
	o.ResolvedAt.Set(nil)
}

// UnsetResolvedAt ensures that no value is present for ResolvedAt, not even an explicit nil
func (o *AlertsLibraryAlertUpdate) UnsetResolvedAt() {
	o.ResolvedAt.Unset()
}

// GetAbnormalityStartTime returns the AbnormalityStartTime field value if set, zero value otherwise.
func (o *AlertsLibraryAlertUpdate) GetAbnormalityStartTime() time.Time {
	if o == nil || IsNil(o.AbnormalityStartTime) {
		var ret time.Time
		return ret
	}
	return *o.AbnormalityStartTime
}

// GetAbnormalityStartTimeOk returns a tuple with the AbnormalityStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsLibraryAlertUpdate) GetAbnormalityStartTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.AbnormalityStartTime) {
		return nil, false
	}
	return o.AbnormalityStartTime, true
}

// HasAbnormalityStartTime returns a boolean if a field has been set.
func (o *AlertsLibraryAlertUpdate) HasAbnormalityStartTime() bool {
	if o != nil && !IsNil(o.AbnormalityStartTime) {
		return true
	}

	return false
}

// SetAbnormalityStartTime gets a reference to the given time.Time and assigns it to the AbnormalityStartTime field.
func (o *AlertsLibraryAlertUpdate) SetAbnormalityStartTime(v time.Time) {
	o.AbnormalityStartTime = &v
}

// GetAlertType returns the AlertType field value if set, zero value otherwise.
func (o *AlertsLibraryAlertUpdate) GetAlertType() string {
	if o == nil || IsNil(o.AlertType) {
		var ret string
		return ret
	}
	return *o.AlertType
}

// GetAlertTypeOk returns a tuple with the AlertType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsLibraryAlertUpdate) GetAlertTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AlertType) {
		return nil, false
	}
	return o.AlertType, true
}

// HasAlertType returns a boolean if a field has been set.
func (o *AlertsLibraryAlertUpdate) HasAlertType() bool {
	if o != nil && !IsNil(o.AlertType) {
		return true
	}

	return false
}

// SetAlertType gets a reference to the given string and assigns it to the AlertType field.
func (o *AlertsLibraryAlertUpdate) SetAlertType(v string) {
	o.AlertType = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AlertsLibraryAlertUpdate) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsLibraryAlertUpdate) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AlertsLibraryAlertUpdate) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AlertsLibraryAlertUpdate) SetStatus(v string) {
	o.Status = &v
}

// GetMonitorQueries returns the MonitorQueries field value if set, zero value otherwise.
func (o *AlertsLibraryAlertUpdate) GetMonitorQueries() []AlertMonitorQuery {
	if o == nil || IsNil(o.MonitorQueries) {
		var ret []AlertMonitorQuery
		return ret
	}
	return o.MonitorQueries
}

// GetMonitorQueriesOk returns a tuple with the MonitorQueries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsLibraryAlertUpdate) GetMonitorQueriesOk() ([]AlertMonitorQuery, bool) {
	if o == nil || IsNil(o.MonitorQueries) {
		return nil, false
	}
	return o.MonitorQueries, true
}

// HasMonitorQueries returns a boolean if a field has been set.
func (o *AlertsLibraryAlertUpdate) HasMonitorQueries() bool {
	if o != nil && !IsNil(o.MonitorQueries) {
		return true
	}

	return false
}

// SetMonitorQueries gets a reference to the given []AlertMonitorQuery and assigns it to the MonitorQueries field.
func (o *AlertsLibraryAlertUpdate) SetMonitorQueries(v []AlertMonitorQuery) {
	o.MonitorQueries = v
}

// GetTriggerQueries returns the TriggerQueries field value if set, zero value otherwise.
func (o *AlertsLibraryAlertUpdate) GetTriggerQueries() []AlertMonitorQuery {
	if o == nil || IsNil(o.TriggerQueries) {
		var ret []AlertMonitorQuery
		return ret
	}
	return o.TriggerQueries
}

// GetTriggerQueriesOk returns a tuple with the TriggerQueries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsLibraryAlertUpdate) GetTriggerQueriesOk() ([]AlertMonitorQuery, bool) {
	if o == nil || IsNil(o.TriggerQueries) {
		return nil, false
	}
	return o.TriggerQueries, true
}

// HasTriggerQueries returns a boolean if a field has been set.
func (o *AlertsLibraryAlertUpdate) HasTriggerQueries() bool {
	if o != nil && !IsNil(o.TriggerQueries) {
		return true
	}

	return false
}

// SetTriggerQueries gets a reference to the given []AlertMonitorQuery and assigns it to the TriggerQueries field.
func (o *AlertsLibraryAlertUpdate) SetTriggerQueries(v []AlertMonitorQuery) {
	o.TriggerQueries = v
}

// GetMonitorUrl returns the MonitorUrl field value if set, zero value otherwise.
func (o *AlertsLibraryAlertUpdate) GetMonitorUrl() string {
	if o == nil || IsNil(o.MonitorUrl) {
		var ret string
		return ret
	}
	return *o.MonitorUrl
}

// GetMonitorUrlOk returns a tuple with the MonitorUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsLibraryAlertUpdate) GetMonitorUrlOk() (*string, bool) {
	if o == nil || IsNil(o.MonitorUrl) {
		return nil, false
	}
	return o.MonitorUrl, true
}

// HasMonitorUrl returns a boolean if a field has been set.
func (o *AlertsLibraryAlertUpdate) HasMonitorUrl() bool {
	if o != nil && !IsNil(o.MonitorUrl) {
		return true
	}

	return false
}

// SetMonitorUrl gets a reference to the given string and assigns it to the MonitorUrl field.
func (o *AlertsLibraryAlertUpdate) SetMonitorUrl(v string) {
	o.MonitorUrl = &v
}

// GetTriggerQueryUrl returns the TriggerQueryUrl field value if set, zero value otherwise.
func (o *AlertsLibraryAlertUpdate) GetTriggerQueryUrl() string {
	if o == nil || IsNil(o.TriggerQueryUrl) {
		var ret string
		return ret
	}
	return *o.TriggerQueryUrl
}

// GetTriggerQueryUrlOk returns a tuple with the TriggerQueryUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsLibraryAlertUpdate) GetTriggerQueryUrlOk() (*string, bool) {
	if o == nil || IsNil(o.TriggerQueryUrl) {
		return nil, false
	}
	return o.TriggerQueryUrl, true
}

// HasTriggerQueryUrl returns a boolean if a field has been set.
func (o *AlertsLibraryAlertUpdate) HasTriggerQueryUrl() bool {
	if o != nil && !IsNil(o.TriggerQueryUrl) {
		return true
	}

	return false
}

// SetTriggerQueryUrl gets a reference to the given string and assigns it to the TriggerQueryUrl field.
func (o *AlertsLibraryAlertUpdate) SetTriggerQueryUrl(v string) {
	o.TriggerQueryUrl = &v
}

// GetTriggerConditions returns the TriggerConditions field value if set, zero value otherwise.
func (o *AlertsLibraryAlertUpdate) GetTriggerConditions() []TriggerCondition {
	if o == nil || IsNil(o.TriggerConditions) {
		var ret []TriggerCondition
		return ret
	}
	return o.TriggerConditions
}

// GetTriggerConditionsOk returns a tuple with the TriggerConditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsLibraryAlertUpdate) GetTriggerConditionsOk() ([]TriggerCondition, bool) {
	if o == nil || IsNil(o.TriggerConditions) {
		return nil, false
	}
	return o.TriggerConditions, true
}

// HasTriggerConditions returns a boolean if a field has been set.
func (o *AlertsLibraryAlertUpdate) HasTriggerConditions() bool {
	if o != nil && !IsNil(o.TriggerConditions) {
		return true
	}

	return false
}

// SetTriggerConditions gets a reference to the given []TriggerCondition and assigns it to the TriggerConditions field.
func (o *AlertsLibraryAlertUpdate) SetTriggerConditions(v []TriggerCondition) {
	o.TriggerConditions = v
}

// GetTriggerValue returns the TriggerValue field value if set, zero value otherwise.
func (o *AlertsLibraryAlertUpdate) GetTriggerValue() float64 {
	if o == nil || IsNil(o.TriggerValue) {
		var ret float64
		return ret
	}
	return *o.TriggerValue
}

// GetTriggerValueOk returns a tuple with the TriggerValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsLibraryAlertUpdate) GetTriggerValueOk() (*float64, bool) {
	if o == nil || IsNil(o.TriggerValue) {
		return nil, false
	}
	return o.TriggerValue, true
}

// HasTriggerValue returns a boolean if a field has been set.
func (o *AlertsLibraryAlertUpdate) HasTriggerValue() bool {
	if o != nil && !IsNil(o.TriggerValue) {
		return true
	}

	return false
}

// SetTriggerValue gets a reference to the given float64 and assigns it to the TriggerValue field.
func (o *AlertsLibraryAlertUpdate) SetTriggerValue(v float64) {
	o.TriggerValue = &v
}

// GetMonitorType returns the MonitorType field value if set, zero value otherwise.
func (o *AlertsLibraryAlertUpdate) GetMonitorType() string {
	if o == nil || IsNil(o.MonitorType) {
		var ret string
		return ret
	}
	return *o.MonitorType
}

// GetMonitorTypeOk returns a tuple with the MonitorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsLibraryAlertUpdate) GetMonitorTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MonitorType) {
		return nil, false
	}
	return o.MonitorType, true
}

// HasMonitorType returns a boolean if a field has been set.
func (o *AlertsLibraryAlertUpdate) HasMonitorType() bool {
	if o != nil && !IsNil(o.MonitorType) {
		return true
	}

	return false
}

// SetMonitorType gets a reference to the given string and assigns it to the MonitorType field.
func (o *AlertsLibraryAlertUpdate) SetMonitorType(v string) {
	o.MonitorType = &v
}

// GetEntityIds returns the EntityIds field value if set, zero value otherwise.
// Deprecated
func (o *AlertsLibraryAlertUpdate) GetEntityIds() []string {
	if o == nil || IsNil(o.EntityIds) {
		var ret []string
		return ret
	}
	return o.EntityIds
}

// GetEntityIdsOk returns a tuple with the EntityIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *AlertsLibraryAlertUpdate) GetEntityIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.EntityIds) {
		return nil, false
	}
	return o.EntityIds, true
}

// HasEntityIds returns a boolean if a field has been set.
func (o *AlertsLibraryAlertUpdate) HasEntityIds() bool {
	if o != nil && !IsNil(o.EntityIds) {
		return true
	}

	return false
}

// SetEntityIds gets a reference to the given []string and assigns it to the EntityIds field.
// Deprecated
func (o *AlertsLibraryAlertUpdate) SetEntityIds(v []string) {
	o.EntityIds = v
}

// GetEntities returns the Entities field value if set, zero value otherwise.
func (o *AlertsLibraryAlertUpdate) GetEntities() []AlertEntityInfo {
	if o == nil || IsNil(o.Entities) {
		var ret []AlertEntityInfo
		return ret
	}
	return o.Entities
}

// GetEntitiesOk returns a tuple with the Entities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsLibraryAlertUpdate) GetEntitiesOk() ([]AlertEntityInfo, bool) {
	if o == nil || IsNil(o.Entities) {
		return nil, false
	}
	return o.Entities, true
}

// HasEntities returns a boolean if a field has been set.
func (o *AlertsLibraryAlertUpdate) HasEntities() bool {
	if o != nil && !IsNil(o.Entities) {
		return true
	}

	return false
}

// SetEntities gets a reference to the given []AlertEntityInfo and assigns it to the Entities field.
func (o *AlertsLibraryAlertUpdate) SetEntities(v []AlertEntityInfo) {
	o.Entities = v
}

// GetSecondaryEntities returns the SecondaryEntities field value if set, zero value otherwise.
func (o *AlertsLibraryAlertUpdate) GetSecondaryEntities() []AlertEntityInfo {
	if o == nil || IsNil(o.SecondaryEntities) {
		var ret []AlertEntityInfo
		return ret
	}
	return o.SecondaryEntities
}

// GetSecondaryEntitiesOk returns a tuple with the SecondaryEntities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsLibraryAlertUpdate) GetSecondaryEntitiesOk() ([]AlertEntityInfo, bool) {
	if o == nil || IsNil(o.SecondaryEntities) {
		return nil, false
	}
	return o.SecondaryEntities, true
}

// HasSecondaryEntities returns a boolean if a field has been set.
func (o *AlertsLibraryAlertUpdate) HasSecondaryEntities() bool {
	if o != nil && !IsNil(o.SecondaryEntities) {
		return true
	}

	return false
}

// SetSecondaryEntities gets a reference to the given []AlertEntityInfo and assigns it to the SecondaryEntities field.
func (o *AlertsLibraryAlertUpdate) SetSecondaryEntities(v []AlertEntityInfo) {
	o.SecondaryEntities = v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *AlertsLibraryAlertUpdate) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsLibraryAlertUpdate) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *AlertsLibraryAlertUpdate) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *AlertsLibraryAlertUpdate) SetNotes(v string) {
	o.Notes = &v
}

// GetExtraDetails returns the ExtraDetails field value if set, zero value otherwise.
func (o *AlertsLibraryAlertUpdate) GetExtraDetails() ExtraDetails {
	if o == nil || IsNil(o.ExtraDetails) {
		var ret ExtraDetails
		return ret
	}
	return *o.ExtraDetails
}

// GetExtraDetailsOk returns a tuple with the ExtraDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsLibraryAlertUpdate) GetExtraDetailsOk() (*ExtraDetails, bool) {
	if o == nil || IsNil(o.ExtraDetails) {
		return nil, false
	}
	return o.ExtraDetails, true
}

// HasExtraDetails returns a boolean if a field has been set.
func (o *AlertsLibraryAlertUpdate) HasExtraDetails() bool {
	if o != nil && !IsNil(o.ExtraDetails) {
		return true
	}

	return false
}

// SetExtraDetails gets a reference to the given ExtraDetails and assigns it to the ExtraDetails field.
func (o *AlertsLibraryAlertUpdate) SetExtraDetails(v ExtraDetails) {
	o.ExtraDetails = &v
}

// GetAlertCondition returns the AlertCondition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlertsLibraryAlertUpdate) GetAlertCondition() string {
	if o == nil || IsNil(o.AlertCondition.Get()) {
		var ret string
		return ret
	}
	return *o.AlertCondition.Get()
}

// GetAlertConditionOk returns a tuple with the AlertCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlertsLibraryAlertUpdate) GetAlertConditionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AlertCondition.Get(), o.AlertCondition.IsSet()
}

// HasAlertCondition returns a boolean if a field has been set.
func (o *AlertsLibraryAlertUpdate) HasAlertCondition() bool {
	if o != nil && o.AlertCondition.IsSet() {
		return true
	}

	return false
}

// SetAlertCondition gets a reference to the given NullableString and assigns it to the AlertCondition field.
func (o *AlertsLibraryAlertUpdate) SetAlertCondition(v string) {
	o.AlertCondition.Set(&v)
}
// SetAlertConditionNil sets the value for AlertCondition to be an explicit nil
func (o *AlertsLibraryAlertUpdate) SetAlertConditionNil() {
	o.AlertCondition.Set(nil)
}

// UnsetAlertCondition ensures that no value is present for AlertCondition, not even an explicit nil
func (o *AlertsLibraryAlertUpdate) UnsetAlertCondition() {
	o.AlertCondition.Unset()
}

// GetIsMuted returns the IsMuted field value if set, zero value otherwise.
func (o *AlertsLibraryAlertUpdate) GetIsMuted() bool {
	if o == nil || IsNil(o.IsMuted) {
		var ret bool
		return ret
	}
	return *o.IsMuted
}

// GetIsMutedOk returns a tuple with the IsMuted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertsLibraryAlertUpdate) GetIsMutedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsMuted) {
		return nil, false
	}
	return o.IsMuted, true
}

// HasIsMuted returns a boolean if a field has been set.
func (o *AlertsLibraryAlertUpdate) HasIsMuted() bool {
	if o != nil && !IsNil(o.IsMuted) {
		return true
	}

	return false
}

// SetIsMuted gets a reference to the given bool and assigns it to the IsMuted field.
func (o *AlertsLibraryAlertUpdate) SetIsMuted(v bool) {
	o.IsMuted = &v
}

func (o AlertsLibraryAlertUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlertsLibraryAlertUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedAlertsLibraryBaseUpdate, errAlertsLibraryBaseUpdate := json.Marshal(o.AlertsLibraryBaseUpdate)
	if errAlertsLibraryBaseUpdate != nil {
		return map[string]interface{}{}, errAlertsLibraryBaseUpdate
	}
	errAlertsLibraryBaseUpdate = json.Unmarshal([]byte(serializedAlertsLibraryBaseUpdate), &toSerialize)
	if errAlertsLibraryBaseUpdate != nil {
		return map[string]interface{}{}, errAlertsLibraryBaseUpdate
	}
	if !IsNil(o.MonitorId) {
		toSerialize["monitorId"] = o.MonitorId
	}
	if o.ResolvedAt.IsSet() {
		toSerialize["resolvedAt"] = o.ResolvedAt.Get()
	}
	if !IsNil(o.AbnormalityStartTime) {
		toSerialize["abnormalityStartTime"] = o.AbnormalityStartTime
	}
	if !IsNil(o.AlertType) {
		toSerialize["alertType"] = o.AlertType
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.MonitorQueries) {
		toSerialize["monitorQueries"] = o.MonitorQueries
	}
	if !IsNil(o.TriggerQueries) {
		toSerialize["triggerQueries"] = o.TriggerQueries
	}
	if !IsNil(o.MonitorUrl) {
		toSerialize["monitorUrl"] = o.MonitorUrl
	}
	if !IsNil(o.TriggerQueryUrl) {
		toSerialize["triggerQueryUrl"] = o.TriggerQueryUrl
	}
	if !IsNil(o.TriggerConditions) {
		toSerialize["triggerConditions"] = o.TriggerConditions
	}
	if !IsNil(o.TriggerValue) {
		toSerialize["triggerValue"] = o.TriggerValue
	}
	if !IsNil(o.MonitorType) {
		toSerialize["monitorType"] = o.MonitorType
	}
	if !IsNil(o.EntityIds) {
		toSerialize["entityIds"] = o.EntityIds
	}
	if !IsNil(o.Entities) {
		toSerialize["entities"] = o.Entities
	}
	if !IsNil(o.SecondaryEntities) {
		toSerialize["secondaryEntities"] = o.SecondaryEntities
	}
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if !IsNil(o.ExtraDetails) {
		toSerialize["extraDetails"] = o.ExtraDetails
	}
	if o.AlertCondition.IsSet() {
		toSerialize["alertCondition"] = o.AlertCondition.Get()
	}
	if !IsNil(o.IsMuted) {
		toSerialize["isMuted"] = o.IsMuted
	}
	return toSerialize, nil
}

func (o *AlertsLibraryAlertUpdate) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"version",
		"type",
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

	varAlertsLibraryAlertUpdate := _AlertsLibraryAlertUpdate{}

	err = json.Unmarshal(bytes, &varAlertsLibraryAlertUpdate)

	if err != nil {
		return err
	}

	*o = AlertsLibraryAlertUpdate(varAlertsLibraryAlertUpdate)

	return err
}

type NullableAlertsLibraryAlertUpdate struct {
	value *AlertsLibraryAlertUpdate
	isSet bool
}

func (v NullableAlertsLibraryAlertUpdate) Get() *AlertsLibraryAlertUpdate {
	return v.value
}

func (v *NullableAlertsLibraryAlertUpdate) Set(val *AlertsLibraryAlertUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertsLibraryAlertUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertsLibraryAlertUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertsLibraryAlertUpdate(val *AlertsLibraryAlertUpdate) *NullableAlertsLibraryAlertUpdate {
	return &NullableAlertsLibraryAlertUpdate{value: val, isSet: true}
}

func (v NullableAlertsLibraryAlertUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertsLibraryAlertUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


