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

// checks if the TraceDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TraceDetail{}

// TraceDetail struct for TraceDetail
type TraceDetail struct {
	// Trace identifier.
	Id string `json:"id"`
	// Root service which started the trace. Examples: `user-service`, `authentication-service`, `payment-service`, `/shopping-cart`
	RootService *string `json:"rootService,omitempty"`
	// Root resource on which the trace was started. Examples: `db.query`, `http.request`, `rpc.call`, `container`
	RootResource *string `json:"rootResource,omitempty"`
	RootStatus *TraceSpanStatus `json:"rootStatus,omitempty"`
	// The name of the operation given to the root span.
	RootOperationName *string `json:"rootOperationName,omitempty"`
	// Calculated trace metrics.
	Metrics *map[string]DoubleTracingValue `json:"metrics,omitempty"`
	// Date and time the trace was started in [ISO 8601 / RFC3339](https://tools.ietf.org/html/rfc3339) format.
	StartedAt *time.Time `json:"startedAt,omitempty"`
	CriticalPathServiceBreakdownSummary *CriticalPathServiceBreakdownSummary `json:"criticalPathServiceBreakdownSummary,omitempty"`
}

type _TraceDetail TraceDetail

// NewTraceDetail instantiates a new TraceDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTraceDetail(id string) *TraceDetail {
	this := TraceDetail{}
	this.Id = id
	return &this
}

// NewTraceDetailWithDefaults instantiates a new TraceDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTraceDetailWithDefaults() *TraceDetail {
	this := TraceDetail{}
	return &this
}

// GetId returns the Id field value
func (o *TraceDetail) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TraceDetail) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TraceDetail) SetId(v string) {
	o.Id = v
}

// GetRootService returns the RootService field value if set, zero value otherwise.
func (o *TraceDetail) GetRootService() string {
	if o == nil || IsNil(o.RootService) {
		var ret string
		return ret
	}
	return *o.RootService
}

// GetRootServiceOk returns a tuple with the RootService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceDetail) GetRootServiceOk() (*string, bool) {
	if o == nil || IsNil(o.RootService) {
		return nil, false
	}
	return o.RootService, true
}

// HasRootService returns a boolean if a field has been set.
func (o *TraceDetail) HasRootService() bool {
	if o != nil && !IsNil(o.RootService) {
		return true
	}

	return false
}

// SetRootService gets a reference to the given string and assigns it to the RootService field.
func (o *TraceDetail) SetRootService(v string) {
	o.RootService = &v
}

// GetRootResource returns the RootResource field value if set, zero value otherwise.
func (o *TraceDetail) GetRootResource() string {
	if o == nil || IsNil(o.RootResource) {
		var ret string
		return ret
	}
	return *o.RootResource
}

// GetRootResourceOk returns a tuple with the RootResource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceDetail) GetRootResourceOk() (*string, bool) {
	if o == nil || IsNil(o.RootResource) {
		return nil, false
	}
	return o.RootResource, true
}

// HasRootResource returns a boolean if a field has been set.
func (o *TraceDetail) HasRootResource() bool {
	if o != nil && !IsNil(o.RootResource) {
		return true
	}

	return false
}

// SetRootResource gets a reference to the given string and assigns it to the RootResource field.
func (o *TraceDetail) SetRootResource(v string) {
	o.RootResource = &v
}

// GetRootStatus returns the RootStatus field value if set, zero value otherwise.
func (o *TraceDetail) GetRootStatus() TraceSpanStatus {
	if o == nil || IsNil(o.RootStatus) {
		var ret TraceSpanStatus
		return ret
	}
	return *o.RootStatus
}

// GetRootStatusOk returns a tuple with the RootStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceDetail) GetRootStatusOk() (*TraceSpanStatus, bool) {
	if o == nil || IsNil(o.RootStatus) {
		return nil, false
	}
	return o.RootStatus, true
}

// HasRootStatus returns a boolean if a field has been set.
func (o *TraceDetail) HasRootStatus() bool {
	if o != nil && !IsNil(o.RootStatus) {
		return true
	}

	return false
}

// SetRootStatus gets a reference to the given TraceSpanStatus and assigns it to the RootStatus field.
func (o *TraceDetail) SetRootStatus(v TraceSpanStatus) {
	o.RootStatus = &v
}

// GetRootOperationName returns the RootOperationName field value if set, zero value otherwise.
func (o *TraceDetail) GetRootOperationName() string {
	if o == nil || IsNil(o.RootOperationName) {
		var ret string
		return ret
	}
	return *o.RootOperationName
}

// GetRootOperationNameOk returns a tuple with the RootOperationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceDetail) GetRootOperationNameOk() (*string, bool) {
	if o == nil || IsNil(o.RootOperationName) {
		return nil, false
	}
	return o.RootOperationName, true
}

// HasRootOperationName returns a boolean if a field has been set.
func (o *TraceDetail) HasRootOperationName() bool {
	if o != nil && !IsNil(o.RootOperationName) {
		return true
	}

	return false
}

// SetRootOperationName gets a reference to the given string and assigns it to the RootOperationName field.
func (o *TraceDetail) SetRootOperationName(v string) {
	o.RootOperationName = &v
}

// GetMetrics returns the Metrics field value if set, zero value otherwise.
func (o *TraceDetail) GetMetrics() map[string]DoubleTracingValue {
	if o == nil || IsNil(o.Metrics) {
		var ret map[string]DoubleTracingValue
		return ret
	}
	return *o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceDetail) GetMetricsOk() (*map[string]DoubleTracingValue, bool) {
	if o == nil || IsNil(o.Metrics) {
		return nil, false
	}
	return o.Metrics, true
}

// HasMetrics returns a boolean if a field has been set.
func (o *TraceDetail) HasMetrics() bool {
	if o != nil && !IsNil(o.Metrics) {
		return true
	}

	return false
}

// SetMetrics gets a reference to the given map[string]DoubleTracingValue and assigns it to the Metrics field.
func (o *TraceDetail) SetMetrics(v map[string]DoubleTracingValue) {
	o.Metrics = &v
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise.
func (o *TraceDetail) GetStartedAt() time.Time {
	if o == nil || IsNil(o.StartedAt) {
		var ret time.Time
		return ret
	}
	return *o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceDetail) GetStartedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartedAt) {
		return nil, false
	}
	return o.StartedAt, true
}

// HasStartedAt returns a boolean if a field has been set.
func (o *TraceDetail) HasStartedAt() bool {
	if o != nil && !IsNil(o.StartedAt) {
		return true
	}

	return false
}

// SetStartedAt gets a reference to the given time.Time and assigns it to the StartedAt field.
func (o *TraceDetail) SetStartedAt(v time.Time) {
	o.StartedAt = &v
}

// GetCriticalPathServiceBreakdownSummary returns the CriticalPathServiceBreakdownSummary field value if set, zero value otherwise.
func (o *TraceDetail) GetCriticalPathServiceBreakdownSummary() CriticalPathServiceBreakdownSummary {
	if o == nil || IsNil(o.CriticalPathServiceBreakdownSummary) {
		var ret CriticalPathServiceBreakdownSummary
		return ret
	}
	return *o.CriticalPathServiceBreakdownSummary
}

// GetCriticalPathServiceBreakdownSummaryOk returns a tuple with the CriticalPathServiceBreakdownSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceDetail) GetCriticalPathServiceBreakdownSummaryOk() (*CriticalPathServiceBreakdownSummary, bool) {
	if o == nil || IsNil(o.CriticalPathServiceBreakdownSummary) {
		return nil, false
	}
	return o.CriticalPathServiceBreakdownSummary, true
}

// HasCriticalPathServiceBreakdownSummary returns a boolean if a field has been set.
func (o *TraceDetail) HasCriticalPathServiceBreakdownSummary() bool {
	if o != nil && !IsNil(o.CriticalPathServiceBreakdownSummary) {
		return true
	}

	return false
}

// SetCriticalPathServiceBreakdownSummary gets a reference to the given CriticalPathServiceBreakdownSummary and assigns it to the CriticalPathServiceBreakdownSummary field.
func (o *TraceDetail) SetCriticalPathServiceBreakdownSummary(v CriticalPathServiceBreakdownSummary) {
	o.CriticalPathServiceBreakdownSummary = &v
}

func (o TraceDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TraceDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.RootService) {
		toSerialize["rootService"] = o.RootService
	}
	if !IsNil(o.RootResource) {
		toSerialize["rootResource"] = o.RootResource
	}
	if !IsNil(o.RootStatus) {
		toSerialize["rootStatus"] = o.RootStatus
	}
	if !IsNil(o.RootOperationName) {
		toSerialize["rootOperationName"] = o.RootOperationName
	}
	if !IsNil(o.Metrics) {
		toSerialize["metrics"] = o.Metrics
	}
	if !IsNil(o.StartedAt) {
		toSerialize["startedAt"] = o.StartedAt
	}
	if !IsNil(o.CriticalPathServiceBreakdownSummary) {
		toSerialize["criticalPathServiceBreakdownSummary"] = o.CriticalPathServiceBreakdownSummary
	}
	return toSerialize, nil
}

func (o *TraceDetail) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varTraceDetail := _TraceDetail{}

	err = json.Unmarshal(bytes, &varTraceDetail)

	if err != nil {
		return err
	}

	*o = TraceDetail(varTraceDetail)

	return err
}

type NullableTraceDetail struct {
	value *TraceDetail
	isSet bool
}

func (v NullableTraceDetail) Get() *TraceDetail {
	return v.value
}

func (v *NullableTraceDetail) Set(val *TraceDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableTraceDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableTraceDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTraceDetail(val *TraceDetail) *NullableTraceDetail {
	return &NullableTraceDetail{value: val, isSet: true}
}

func (v NullableTraceDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTraceDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


