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

// checks if the SpanQuerySpanData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpanQuerySpanData{}

// SpanQuerySpanData struct for SpanQuerySpanData
type SpanQuerySpanData struct {
	// Identifier of the span.
	SpanId *string `json:"spanId,omitempty"`
	// Identifier of the trace.
	TraceId *string `json:"traceId,omitempty"`
	// Identifier of the parent span, if any. If the span has no parent it's considered a root span.
	ParentSpanId *string `json:"parentSpanId,omitempty"`
	// The name of the operation given to the span.
	OperationName *string `json:"operationName,omitempty"`
	// The name of the service this span is part of.
	Service *string `json:"service,omitempty"`
	// Name of the possible remote span's service.
	RemoteService *string `json:"remoteService,omitempty"`
	// Number of nanoseconds the span lasted.
	Duration int64 `json:"duration"`
	// Date and time the span was started in [ISO 8601 / RFC3339](https://tools.ietf.org/html/rfc3339) format.
	StartedAt time.Time `json:"startedAt"`
	Status *TraceSpanStatus `json:"status,omitempty"`
	// Span kind describes the relationship between the Span, its parents, and its children in a Trace. Possible values: `CLIENT`, `SERVER`, `PRODUCER`, `CONSUMER`, `INTERNAL`.
	Kind *string `json:"kind,omitempty"`
	// Tags attached to this span as JSON.
	TagsJSON *string `json:"tagsJSON,omitempty"`
	// Metadata attached to the span.
	Metadata *map[string]string `json:"metadata,omitempty"`
}

type _SpanQuerySpanData SpanQuerySpanData

// NewSpanQuerySpanData instantiates a new SpanQuerySpanData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpanQuerySpanData(duration int64, startedAt time.Time) *SpanQuerySpanData {
	this := SpanQuerySpanData{}
	this.Duration = duration
	this.StartedAt = startedAt
	return &this
}

// NewSpanQuerySpanDataWithDefaults instantiates a new SpanQuerySpanData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpanQuerySpanDataWithDefaults() *SpanQuerySpanData {
	this := SpanQuerySpanData{}
	return &this
}

// GetSpanId returns the SpanId field value if set, zero value otherwise.
func (o *SpanQuerySpanData) GetSpanId() string {
	if o == nil || IsNil(o.SpanId) {
		var ret string
		return ret
	}
	return *o.SpanId
}

// GetSpanIdOk returns a tuple with the SpanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpanQuerySpanData) GetSpanIdOk() (*string, bool) {
	if o == nil || IsNil(o.SpanId) {
		return nil, false
	}
	return o.SpanId, true
}

// HasSpanId returns a boolean if a field has been set.
func (o *SpanQuerySpanData) HasSpanId() bool {
	if o != nil && !IsNil(o.SpanId) {
		return true
	}

	return false
}

// SetSpanId gets a reference to the given string and assigns it to the SpanId field.
func (o *SpanQuerySpanData) SetSpanId(v string) {
	o.SpanId = &v
}

// GetTraceId returns the TraceId field value if set, zero value otherwise.
func (o *SpanQuerySpanData) GetTraceId() string {
	if o == nil || IsNil(o.TraceId) {
		var ret string
		return ret
	}
	return *o.TraceId
}

// GetTraceIdOk returns a tuple with the TraceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpanQuerySpanData) GetTraceIdOk() (*string, bool) {
	if o == nil || IsNil(o.TraceId) {
		return nil, false
	}
	return o.TraceId, true
}

// HasTraceId returns a boolean if a field has been set.
func (o *SpanQuerySpanData) HasTraceId() bool {
	if o != nil && !IsNil(o.TraceId) {
		return true
	}

	return false
}

// SetTraceId gets a reference to the given string and assigns it to the TraceId field.
func (o *SpanQuerySpanData) SetTraceId(v string) {
	o.TraceId = &v
}

// GetParentSpanId returns the ParentSpanId field value if set, zero value otherwise.
func (o *SpanQuerySpanData) GetParentSpanId() string {
	if o == nil || IsNil(o.ParentSpanId) {
		var ret string
		return ret
	}
	return *o.ParentSpanId
}

// GetParentSpanIdOk returns a tuple with the ParentSpanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpanQuerySpanData) GetParentSpanIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentSpanId) {
		return nil, false
	}
	return o.ParentSpanId, true
}

// HasParentSpanId returns a boolean if a field has been set.
func (o *SpanQuerySpanData) HasParentSpanId() bool {
	if o != nil && !IsNil(o.ParentSpanId) {
		return true
	}

	return false
}

// SetParentSpanId gets a reference to the given string and assigns it to the ParentSpanId field.
func (o *SpanQuerySpanData) SetParentSpanId(v string) {
	o.ParentSpanId = &v
}

// GetOperationName returns the OperationName field value if set, zero value otherwise.
func (o *SpanQuerySpanData) GetOperationName() string {
	if o == nil || IsNil(o.OperationName) {
		var ret string
		return ret
	}
	return *o.OperationName
}

// GetOperationNameOk returns a tuple with the OperationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpanQuerySpanData) GetOperationNameOk() (*string, bool) {
	if o == nil || IsNil(o.OperationName) {
		return nil, false
	}
	return o.OperationName, true
}

// HasOperationName returns a boolean if a field has been set.
func (o *SpanQuerySpanData) HasOperationName() bool {
	if o != nil && !IsNil(o.OperationName) {
		return true
	}

	return false
}

// SetOperationName gets a reference to the given string and assigns it to the OperationName field.
func (o *SpanQuerySpanData) SetOperationName(v string) {
	o.OperationName = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *SpanQuerySpanData) GetService() string {
	if o == nil || IsNil(o.Service) {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpanQuerySpanData) GetServiceOk() (*string, bool) {
	if o == nil || IsNil(o.Service) {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *SpanQuerySpanData) HasService() bool {
	if o != nil && !IsNil(o.Service) {
		return true
	}

	return false
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *SpanQuerySpanData) SetService(v string) {
	o.Service = &v
}

// GetRemoteService returns the RemoteService field value if set, zero value otherwise.
func (o *SpanQuerySpanData) GetRemoteService() string {
	if o == nil || IsNil(o.RemoteService) {
		var ret string
		return ret
	}
	return *o.RemoteService
}

// GetRemoteServiceOk returns a tuple with the RemoteService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpanQuerySpanData) GetRemoteServiceOk() (*string, bool) {
	if o == nil || IsNil(o.RemoteService) {
		return nil, false
	}
	return o.RemoteService, true
}

// HasRemoteService returns a boolean if a field has been set.
func (o *SpanQuerySpanData) HasRemoteService() bool {
	if o != nil && !IsNil(o.RemoteService) {
		return true
	}

	return false
}

// SetRemoteService gets a reference to the given string and assigns it to the RemoteService field.
func (o *SpanQuerySpanData) SetRemoteService(v string) {
	o.RemoteService = &v
}

// GetDuration returns the Duration field value
func (o *SpanQuerySpanData) GetDuration() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *SpanQuerySpanData) GetDurationOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *SpanQuerySpanData) SetDuration(v int64) {
	o.Duration = v
}

// GetStartedAt returns the StartedAt field value
func (o *SpanQuerySpanData) GetStartedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value
// and a boolean to check if the value has been set.
func (o *SpanQuerySpanData) GetStartedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartedAt, true
}

// SetStartedAt sets field value
func (o *SpanQuerySpanData) SetStartedAt(v time.Time) {
	o.StartedAt = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SpanQuerySpanData) GetStatus() TraceSpanStatus {
	if o == nil || IsNil(o.Status) {
		var ret TraceSpanStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpanQuerySpanData) GetStatusOk() (*TraceSpanStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SpanQuerySpanData) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given TraceSpanStatus and assigns it to the Status field.
func (o *SpanQuerySpanData) SetStatus(v TraceSpanStatus) {
	o.Status = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *SpanQuerySpanData) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpanQuerySpanData) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *SpanQuerySpanData) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *SpanQuerySpanData) SetKind(v string) {
	o.Kind = &v
}

// GetTagsJSON returns the TagsJSON field value if set, zero value otherwise.
func (o *SpanQuerySpanData) GetTagsJSON() string {
	if o == nil || IsNil(o.TagsJSON) {
		var ret string
		return ret
	}
	return *o.TagsJSON
}

// GetTagsJSONOk returns a tuple with the TagsJSON field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpanQuerySpanData) GetTagsJSONOk() (*string, bool) {
	if o == nil || IsNil(o.TagsJSON) {
		return nil, false
	}
	return o.TagsJSON, true
}

// HasTagsJSON returns a boolean if a field has been set.
func (o *SpanQuerySpanData) HasTagsJSON() bool {
	if o != nil && !IsNil(o.TagsJSON) {
		return true
	}

	return false
}

// SetTagsJSON gets a reference to the given string and assigns it to the TagsJSON field.
func (o *SpanQuerySpanData) SetTagsJSON(v string) {
	o.TagsJSON = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *SpanQuerySpanData) GetMetadata() map[string]string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpanQuerySpanData) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *SpanQuerySpanData) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *SpanQuerySpanData) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

func (o SpanQuerySpanData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpanQuerySpanData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SpanId) {
		toSerialize["spanId"] = o.SpanId
	}
	if !IsNil(o.TraceId) {
		toSerialize["traceId"] = o.TraceId
	}
	if !IsNil(o.ParentSpanId) {
		toSerialize["parentSpanId"] = o.ParentSpanId
	}
	if !IsNil(o.OperationName) {
		toSerialize["operationName"] = o.OperationName
	}
	if !IsNil(o.Service) {
		toSerialize["service"] = o.Service
	}
	if !IsNil(o.RemoteService) {
		toSerialize["remoteService"] = o.RemoteService
	}
	toSerialize["duration"] = o.Duration
	toSerialize["startedAt"] = o.StartedAt
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if !IsNil(o.TagsJSON) {
		toSerialize["tagsJSON"] = o.TagsJSON
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

func (o *SpanQuerySpanData) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"duration",
		"startedAt",
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

	varSpanQuerySpanData := _SpanQuerySpanData{}

	err = json.Unmarshal(bytes, &varSpanQuerySpanData)

	if err != nil {
		return err
	}

	*o = SpanQuerySpanData(varSpanQuerySpanData)

	return err
}

type NullableSpanQuerySpanData struct {
	value *SpanQuerySpanData
	isSet bool
}

func (v NullableSpanQuerySpanData) Get() *SpanQuerySpanData {
	return v.value
}

func (v *NullableSpanQuerySpanData) Set(val *SpanQuerySpanData) {
	v.value = val
	v.isSet = true
}

func (v NullableSpanQuerySpanData) IsSet() bool {
	return v.isSet
}

func (v *NullableSpanQuerySpanData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpanQuerySpanData(val *SpanQuerySpanData) *NullableSpanQuerySpanData {
	return &NullableSpanQuerySpanData{value: val, isSet: true}
}

func (v NullableSpanQuerySpanData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpanQuerySpanData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


