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

// checks if the MonitorTemplatesLibraryMonitorTemplateUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MonitorTemplatesLibraryMonitorTemplateUpdate{}

// MonitorTemplatesLibraryMonitorTemplateUpdate struct for MonitorTemplatesLibraryMonitorTemplateUpdate
type MonitorTemplatesLibraryMonitorTemplateUpdate struct {
	MonitorTemplatesLibraryBaseUpdate
	// The type of monitor template. Valid values:   1. `Logs`: A logs query monitor template.   2. `Metrics`: A metrics query monitor template.   3. `Slo`: A SLO based monitor template.
	MonitorType string `json:"monitorType"`
	// The delay duration for evaluating the monitor (relative to current time). The timerange of monitor will be shifted in the past by this delay time.
	EvaluationDelay *string `json:"evaluationDelay,omitempty"`
	// The name of the alert(s) triggered from the monitor created based on the template. Monitor name will be used if not specified.
	AlertName *string `json:"alertName,omitempty"`
	// All queries from the monitor.
	Queries []MonitorQuery `json:"queries"`
	// Defines the conditions of when to send notifications.
	Triggers []TriggerCondition `json:"triggers"`
	// Whether or not the monitor template is disabled.
	IsDisabled *bool `json:"isDisabled,omitempty"`
	// Whether or not to group notifications for individual items that meet the trigger condition.
	GroupNotifications *bool `json:"groupNotifications,omitempty"`
	// Notes such as links and instruction to help you resolve alerts triggered by this monitor template. {{Markdown}} supported. It will be enabled only if available for your organization. Please contact your Sumo Logic account team to learn more.
	Playbook *string `json:"playbook,omitempty"`
}

type _MonitorTemplatesLibraryMonitorTemplateUpdate MonitorTemplatesLibraryMonitorTemplateUpdate

// NewMonitorTemplatesLibraryMonitorTemplateUpdate instantiates a new MonitorTemplatesLibraryMonitorTemplateUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitorTemplatesLibraryMonitorTemplateUpdate(monitorType string, queries []MonitorQuery, triggers []TriggerCondition, name string, version int64, type_ string) *MonitorTemplatesLibraryMonitorTemplateUpdate {
	this := MonitorTemplatesLibraryMonitorTemplateUpdate{}
	this.Name = name
	var description string = ""
	this.Description = &description
	this.Version = version
	this.Type = type_
	this.MonitorType = monitorType
	var evaluationDelay string = "0m"
	this.EvaluationDelay = &evaluationDelay
	this.Queries = queries
	this.Triggers = triggers
	var isDisabled bool = false
	this.IsDisabled = &isDisabled
	var groupNotifications bool = true
	this.GroupNotifications = &groupNotifications
	var playbook string = ""
	this.Playbook = &playbook
	return &this
}

// NewMonitorTemplatesLibraryMonitorTemplateUpdateWithDefaults instantiates a new MonitorTemplatesLibraryMonitorTemplateUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitorTemplatesLibraryMonitorTemplateUpdateWithDefaults() *MonitorTemplatesLibraryMonitorTemplateUpdate {
	this := MonitorTemplatesLibraryMonitorTemplateUpdate{}
	var evaluationDelay string = "0m"
	this.EvaluationDelay = &evaluationDelay
	var isDisabled bool = false
	this.IsDisabled = &isDisabled
	var groupNotifications bool = true
	this.GroupNotifications = &groupNotifications
	var playbook string = ""
	this.Playbook = &playbook
	return &this
}

// GetMonitorType returns the MonitorType field value
func (o *MonitorTemplatesLibraryMonitorTemplateUpdate) GetMonitorType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MonitorType
}

// GetMonitorTypeOk returns a tuple with the MonitorType field value
// and a boolean to check if the value has been set.
func (o *MonitorTemplatesLibraryMonitorTemplateUpdate) GetMonitorTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MonitorType, true
}

// SetMonitorType sets field value
func (o *MonitorTemplatesLibraryMonitorTemplateUpdate) SetMonitorType(v string) {
	o.MonitorType = v
}

// GetEvaluationDelay returns the EvaluationDelay field value if set, zero value otherwise.
func (o *MonitorTemplatesLibraryMonitorTemplateUpdate) GetEvaluationDelay() string {
	if o == nil || IsNil(o.EvaluationDelay) {
		var ret string
		return ret
	}
	return *o.EvaluationDelay
}

// GetEvaluationDelayOk returns a tuple with the EvaluationDelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorTemplatesLibraryMonitorTemplateUpdate) GetEvaluationDelayOk() (*string, bool) {
	if o == nil || IsNil(o.EvaluationDelay) {
		return nil, false
	}
	return o.EvaluationDelay, true
}

// HasEvaluationDelay returns a boolean if a field has been set.
func (o *MonitorTemplatesLibraryMonitorTemplateUpdate) HasEvaluationDelay() bool {
	if o != nil && !IsNil(o.EvaluationDelay) {
		return true
	}

	return false
}

// SetEvaluationDelay gets a reference to the given string and assigns it to the EvaluationDelay field.
func (o *MonitorTemplatesLibraryMonitorTemplateUpdate) SetEvaluationDelay(v string) {
	o.EvaluationDelay = &v
}

// GetAlertName returns the AlertName field value if set, zero value otherwise.
func (o *MonitorTemplatesLibraryMonitorTemplateUpdate) GetAlertName() string {
	if o == nil || IsNil(o.AlertName) {
		var ret string
		return ret
	}
	return *o.AlertName
}

// GetAlertNameOk returns a tuple with the AlertName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorTemplatesLibraryMonitorTemplateUpdate) GetAlertNameOk() (*string, bool) {
	if o == nil || IsNil(o.AlertName) {
		return nil, false
	}
	return o.AlertName, true
}

// HasAlertName returns a boolean if a field has been set.
func (o *MonitorTemplatesLibraryMonitorTemplateUpdate) HasAlertName() bool {
	if o != nil && !IsNil(o.AlertName) {
		return true
	}

	return false
}

// SetAlertName gets a reference to the given string and assigns it to the AlertName field.
func (o *MonitorTemplatesLibraryMonitorTemplateUpdate) SetAlertName(v string) {
	o.AlertName = &v
}

// GetQueries returns the Queries field value
func (o *MonitorTemplatesLibraryMonitorTemplateUpdate) GetQueries() []MonitorQuery {
	if o == nil {
		var ret []MonitorQuery
		return ret
	}

	return o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value
// and a boolean to check if the value has been set.
func (o *MonitorTemplatesLibraryMonitorTemplateUpdate) GetQueriesOk() ([]MonitorQuery, bool) {
	if o == nil {
		return nil, false
	}
	return o.Queries, true
}

// SetQueries sets field value
func (o *MonitorTemplatesLibraryMonitorTemplateUpdate) SetQueries(v []MonitorQuery) {
	o.Queries = v
}

// GetTriggers returns the Triggers field value
func (o *MonitorTemplatesLibraryMonitorTemplateUpdate) GetTriggers() []TriggerCondition {
	if o == nil {
		var ret []TriggerCondition
		return ret
	}

	return o.Triggers
}

// GetTriggersOk returns a tuple with the Triggers field value
// and a boolean to check if the value has been set.
func (o *MonitorTemplatesLibraryMonitorTemplateUpdate) GetTriggersOk() ([]TriggerCondition, bool) {
	if o == nil {
		return nil, false
	}
	return o.Triggers, true
}

// SetTriggers sets field value
func (o *MonitorTemplatesLibraryMonitorTemplateUpdate) SetTriggers(v []TriggerCondition) {
	o.Triggers = v
}

// GetIsDisabled returns the IsDisabled field value if set, zero value otherwise.
func (o *MonitorTemplatesLibraryMonitorTemplateUpdate) GetIsDisabled() bool {
	if o == nil || IsNil(o.IsDisabled) {
		var ret bool
		return ret
	}
	return *o.IsDisabled
}

// GetIsDisabledOk returns a tuple with the IsDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorTemplatesLibraryMonitorTemplateUpdate) GetIsDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDisabled) {
		return nil, false
	}
	return o.IsDisabled, true
}

// HasIsDisabled returns a boolean if a field has been set.
func (o *MonitorTemplatesLibraryMonitorTemplateUpdate) HasIsDisabled() bool {
	if o != nil && !IsNil(o.IsDisabled) {
		return true
	}

	return false
}

// SetIsDisabled gets a reference to the given bool and assigns it to the IsDisabled field.
func (o *MonitorTemplatesLibraryMonitorTemplateUpdate) SetIsDisabled(v bool) {
	o.IsDisabled = &v
}

// GetGroupNotifications returns the GroupNotifications field value if set, zero value otherwise.
func (o *MonitorTemplatesLibraryMonitorTemplateUpdate) GetGroupNotifications() bool {
	if o == nil || IsNil(o.GroupNotifications) {
		var ret bool
		return ret
	}
	return *o.GroupNotifications
}

// GetGroupNotificationsOk returns a tuple with the GroupNotifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorTemplatesLibraryMonitorTemplateUpdate) GetGroupNotificationsOk() (*bool, bool) {
	if o == nil || IsNil(o.GroupNotifications) {
		return nil, false
	}
	return o.GroupNotifications, true
}

// HasGroupNotifications returns a boolean if a field has been set.
func (o *MonitorTemplatesLibraryMonitorTemplateUpdate) HasGroupNotifications() bool {
	if o != nil && !IsNil(o.GroupNotifications) {
		return true
	}

	return false
}

// SetGroupNotifications gets a reference to the given bool and assigns it to the GroupNotifications field.
func (o *MonitorTemplatesLibraryMonitorTemplateUpdate) SetGroupNotifications(v bool) {
	o.GroupNotifications = &v
}

// GetPlaybook returns the Playbook field value if set, zero value otherwise.
func (o *MonitorTemplatesLibraryMonitorTemplateUpdate) GetPlaybook() string {
	if o == nil || IsNil(o.Playbook) {
		var ret string
		return ret
	}
	return *o.Playbook
}

// GetPlaybookOk returns a tuple with the Playbook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorTemplatesLibraryMonitorTemplateUpdate) GetPlaybookOk() (*string, bool) {
	if o == nil || IsNil(o.Playbook) {
		return nil, false
	}
	return o.Playbook, true
}

// HasPlaybook returns a boolean if a field has been set.
func (o *MonitorTemplatesLibraryMonitorTemplateUpdate) HasPlaybook() bool {
	if o != nil && !IsNil(o.Playbook) {
		return true
	}

	return false
}

// SetPlaybook gets a reference to the given string and assigns it to the Playbook field.
func (o *MonitorTemplatesLibraryMonitorTemplateUpdate) SetPlaybook(v string) {
	o.Playbook = &v
}

func (o MonitorTemplatesLibraryMonitorTemplateUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MonitorTemplatesLibraryMonitorTemplateUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMonitorTemplatesLibraryBaseUpdate, errMonitorTemplatesLibraryBaseUpdate := json.Marshal(o.MonitorTemplatesLibraryBaseUpdate)
	if errMonitorTemplatesLibraryBaseUpdate != nil {
		return map[string]interface{}{}, errMonitorTemplatesLibraryBaseUpdate
	}
	errMonitorTemplatesLibraryBaseUpdate = json.Unmarshal([]byte(serializedMonitorTemplatesLibraryBaseUpdate), &toSerialize)
	if errMonitorTemplatesLibraryBaseUpdate != nil {
		return map[string]interface{}{}, errMonitorTemplatesLibraryBaseUpdate
	}
	toSerialize["monitorType"] = o.MonitorType
	if !IsNil(o.EvaluationDelay) {
		toSerialize["evaluationDelay"] = o.EvaluationDelay
	}
	if !IsNil(o.AlertName) {
		toSerialize["alertName"] = o.AlertName
	}
	toSerialize["queries"] = o.Queries
	toSerialize["triggers"] = o.Triggers
	if !IsNil(o.IsDisabled) {
		toSerialize["isDisabled"] = o.IsDisabled
	}
	if !IsNil(o.GroupNotifications) {
		toSerialize["groupNotifications"] = o.GroupNotifications
	}
	if !IsNil(o.Playbook) {
		toSerialize["playbook"] = o.Playbook
	}
	return toSerialize, nil
}

func (o *MonitorTemplatesLibraryMonitorTemplateUpdate) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"monitorType",
		"queries",
		"triggers",
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

	varMonitorTemplatesLibraryMonitorTemplateUpdate := _MonitorTemplatesLibraryMonitorTemplateUpdate{}

	err = json.Unmarshal(bytes, &varMonitorTemplatesLibraryMonitorTemplateUpdate)

	if err != nil {
		return err
	}

	*o = MonitorTemplatesLibraryMonitorTemplateUpdate(varMonitorTemplatesLibraryMonitorTemplateUpdate)

	return err
}

type NullableMonitorTemplatesLibraryMonitorTemplateUpdate struct {
	value *MonitorTemplatesLibraryMonitorTemplateUpdate
	isSet bool
}

func (v NullableMonitorTemplatesLibraryMonitorTemplateUpdate) Get() *MonitorTemplatesLibraryMonitorTemplateUpdate {
	return v.value
}

func (v *NullableMonitorTemplatesLibraryMonitorTemplateUpdate) Set(val *MonitorTemplatesLibraryMonitorTemplateUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitorTemplatesLibraryMonitorTemplateUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitorTemplatesLibraryMonitorTemplateUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitorTemplatesLibraryMonitorTemplateUpdate(val *MonitorTemplatesLibraryMonitorTemplateUpdate) *NullableMonitorTemplatesLibraryMonitorTemplateUpdate {
	return &NullableMonitorTemplatesLibraryMonitorTemplateUpdate{value: val, isSet: true}
}

func (v NullableMonitorTemplatesLibraryMonitorTemplateUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitorTemplatesLibraryMonitorTemplateUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


