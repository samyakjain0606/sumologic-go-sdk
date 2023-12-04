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

// checks if the StaticCondition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StaticCondition{}

// StaticCondition struct for StaticCondition
type StaticCondition struct {
	TriggerCondition
	// The relative time range of the monitor. Valid values of time ranges are `-5m`, `-10m`, `-15m`, `-30m`, `-1h`, `-3h`, `-6h`, `-12h`, or `-24h`.
	TimeRange string `json:"timeRange"`
	// The data value for the condition. This defines the threshold for when to trigger. Threshold value is not applicable for `MissingData` and `ResolvedMissingData` triggerTypes and will be ignored if specified.
	Threshold *float64 `json:"threshold,omitempty"`
	// The comparison type for the `threshold` evaluation. This defines how you want the data value compared. Valid values:   1. `LessThan`: Less than than the configured threshold.   2. `GreaterThan`: Greater than the configured threshold.   3. `LessThanOrEqual`: Less than or equal to the configured threshold.   4. `GreaterThanOrEqual`: Greater than or equal to the configured threshold. ThresholdType value is not applicable for `MissingData` and `ResolvedMissingData` triggerTypes and will be ignored if specified.
	ThresholdType *string `json:"thresholdType,omitempty"`
	// The name of the field that the trigger condition will alert on. The trigger could compare the value of specified field with the threshold. If `field` is not specified, monitor would default to result count instead.
	Field *string `json:"field,omitempty"`
	// The criteria to evaluate the threshold and thresholdType in the given time range. Valid values:   1. `AtLeastOnce`: Trigger if the threshold is met at least once. (NOTE: This is the only valid value if monitorType is `Metrics`.)   2. `Always`: Trigger if the threshold is met continuously. (NOTE: This is the only valid value if monitorType is `Metrics`.)   3. `ResultCount`: Trigger if the threshold is met against the count of results. (NOTE: This is the only valid value if monitorType is `Logs`.)   4. `MissingData`: Trigger if the data is missing. (NOTE: This is valid for both `Logs` and `Metrics` monitorTypes)
	OccurrenceType string `json:"occurrenceType"`
	// Determines which time series from queries to use for Metrics MissingData and ResolvedMissingData triggers Valid values:   1. `AllTimeSeries`: Evaluate the condition against all time series. (NOTE: This option is only valid if monitorType is `Metrics`)   2. `AnyTimeSeries`: Evaluate the condition against any time series. (NOTE: This option is only valid if monitorType is `Metrics`)   3. `AllResults`: Evaluate the condition against results from all queries. (NOTE: This option is only valid if monitorType is `Logs`)
	TriggerSource string `json:"triggerSource"`
	// The minimum number of data points to alert or resolve a metrics monitor within the time range. This field is only valid for Metrics Monitor, it will always be set to 1 for `AtleastOnce` occurrence type and for `Always`, if not specified by user it will default to 2.
	MinDataPoints *int32 `json:"minDataPoints,omitempty"`
}

type _StaticCondition StaticCondition

// NewStaticCondition instantiates a new StaticCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStaticCondition(timeRange string, occurrenceType string, triggerSource string, triggerType string) *StaticCondition {
	this := StaticCondition{}
	var detectionMethod string = "StaticCondition"
	this.DetectionMethod = &detectionMethod
	this.TriggerType = triggerType
	this.TimeRange = timeRange
	var threshold float64 = 0.0
	this.Threshold = &threshold
	var thresholdType string = "GreaterThanOrEqual"
	this.ThresholdType = &thresholdType
	this.OccurrenceType = occurrenceType
	this.TriggerSource = triggerSource
	return &this
}

// NewStaticConditionWithDefaults instantiates a new StaticCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStaticConditionWithDefaults() *StaticCondition {
	this := StaticCondition{}
	var threshold float64 = 0.0
	this.Threshold = &threshold
	var thresholdType string = "GreaterThanOrEqual"
	this.ThresholdType = &thresholdType
	return &this
}

// GetTimeRange returns the TimeRange field value
func (o *StaticCondition) GetTimeRange() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TimeRange
}

// GetTimeRangeOk returns a tuple with the TimeRange field value
// and a boolean to check if the value has been set.
func (o *StaticCondition) GetTimeRangeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeRange, true
}

// SetTimeRange sets field value
func (o *StaticCondition) SetTimeRange(v string) {
	o.TimeRange = v
}

// GetThreshold returns the Threshold field value if set, zero value otherwise.
func (o *StaticCondition) GetThreshold() float64 {
	if o == nil || IsNil(o.Threshold) {
		var ret float64
		return ret
	}
	return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaticCondition) GetThresholdOk() (*float64, bool) {
	if o == nil || IsNil(o.Threshold) {
		return nil, false
	}
	return o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *StaticCondition) HasThreshold() bool {
	if o != nil && !IsNil(o.Threshold) {
		return true
	}

	return false
}

// SetThreshold gets a reference to the given float64 and assigns it to the Threshold field.
func (o *StaticCondition) SetThreshold(v float64) {
	o.Threshold = &v
}

// GetThresholdType returns the ThresholdType field value if set, zero value otherwise.
func (o *StaticCondition) GetThresholdType() string {
	if o == nil || IsNil(o.ThresholdType) {
		var ret string
		return ret
	}
	return *o.ThresholdType
}

// GetThresholdTypeOk returns a tuple with the ThresholdType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaticCondition) GetThresholdTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ThresholdType) {
		return nil, false
	}
	return o.ThresholdType, true
}

// HasThresholdType returns a boolean if a field has been set.
func (o *StaticCondition) HasThresholdType() bool {
	if o != nil && !IsNil(o.ThresholdType) {
		return true
	}

	return false
}

// SetThresholdType gets a reference to the given string and assigns it to the ThresholdType field.
func (o *StaticCondition) SetThresholdType(v string) {
	o.ThresholdType = &v
}

// GetField returns the Field field value if set, zero value otherwise.
func (o *StaticCondition) GetField() string {
	if o == nil || IsNil(o.Field) {
		var ret string
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaticCondition) GetFieldOk() (*string, bool) {
	if o == nil || IsNil(o.Field) {
		return nil, false
	}
	return o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *StaticCondition) HasField() bool {
	if o != nil && !IsNil(o.Field) {
		return true
	}

	return false
}

// SetField gets a reference to the given string and assigns it to the Field field.
func (o *StaticCondition) SetField(v string) {
	o.Field = &v
}

// GetOccurrenceType returns the OccurrenceType field value
func (o *StaticCondition) GetOccurrenceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OccurrenceType
}

// GetOccurrenceTypeOk returns a tuple with the OccurrenceType field value
// and a boolean to check if the value has been set.
func (o *StaticCondition) GetOccurrenceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OccurrenceType, true
}

// SetOccurrenceType sets field value
func (o *StaticCondition) SetOccurrenceType(v string) {
	o.OccurrenceType = v
}

// GetTriggerSource returns the TriggerSource field value
func (o *StaticCondition) GetTriggerSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TriggerSource
}

// GetTriggerSourceOk returns a tuple with the TriggerSource field value
// and a boolean to check if the value has been set.
func (o *StaticCondition) GetTriggerSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TriggerSource, true
}

// SetTriggerSource sets field value
func (o *StaticCondition) SetTriggerSource(v string) {
	o.TriggerSource = v
}

// GetMinDataPoints returns the MinDataPoints field value if set, zero value otherwise.
func (o *StaticCondition) GetMinDataPoints() int32 {
	if o == nil || IsNil(o.MinDataPoints) {
		var ret int32
		return ret
	}
	return *o.MinDataPoints
}

// GetMinDataPointsOk returns a tuple with the MinDataPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaticCondition) GetMinDataPointsOk() (*int32, bool) {
	if o == nil || IsNil(o.MinDataPoints) {
		return nil, false
	}
	return o.MinDataPoints, true
}

// HasMinDataPoints returns a boolean if a field has been set.
func (o *StaticCondition) HasMinDataPoints() bool {
	if o != nil && !IsNil(o.MinDataPoints) {
		return true
	}

	return false
}

// SetMinDataPoints gets a reference to the given int32 and assigns it to the MinDataPoints field.
func (o *StaticCondition) SetMinDataPoints(v int32) {
	o.MinDataPoints = &v
}

func (o StaticCondition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StaticCondition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedTriggerCondition, errTriggerCondition := json.Marshal(o.TriggerCondition)
	if errTriggerCondition != nil {
		return map[string]interface{}{}, errTriggerCondition
	}
	errTriggerCondition = json.Unmarshal([]byte(serializedTriggerCondition), &toSerialize)
	if errTriggerCondition != nil {
		return map[string]interface{}{}, errTriggerCondition
	}
	toSerialize["timeRange"] = o.TimeRange
	if !IsNil(o.Threshold) {
		toSerialize["threshold"] = o.Threshold
	}
	if !IsNil(o.ThresholdType) {
		toSerialize["thresholdType"] = o.ThresholdType
	}
	if !IsNil(o.Field) {
		toSerialize["field"] = o.Field
	}
	toSerialize["occurrenceType"] = o.OccurrenceType
	toSerialize["triggerSource"] = o.TriggerSource
	if !IsNil(o.MinDataPoints) {
		toSerialize["minDataPoints"] = o.MinDataPoints
	}
	return toSerialize, nil
}

func (o *StaticCondition) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"timeRange",
		"occurrenceType",
		"triggerSource",
		"triggerType",
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

	varStaticCondition := _StaticCondition{}

	err = json.Unmarshal(bytes, &varStaticCondition)

	if err != nil {
		return err
	}

	*o = StaticCondition(varStaticCondition)

	return err
}

type NullableStaticCondition struct {
	value *StaticCondition
	isSet bool
}

func (v NullableStaticCondition) Get() *StaticCondition {
	return v.value
}

func (v *NullableStaticCondition) Set(val *StaticCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableStaticCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableStaticCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStaticCondition(val *StaticCondition) *NullableStaticCondition {
	return &NullableStaticCondition{value: val, isSet: true}
}

func (v NullableStaticCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStaticCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


