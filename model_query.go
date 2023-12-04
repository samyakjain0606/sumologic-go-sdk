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

// checks if the Query type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Query{}

// Query struct for Query
type Query struct {
	// The metrics, traces or logs query.
	QueryString string `json:"queryString"`
	// The type of the query, either `Metrics`, `Traces`, `Spans` or `Logs`.
	QueryType string `json:"queryType"`
	// The key for metric, traces or log queries. Used as an identifier for queries. It is displayed on the panel builder and used for display overrides and query toggling. 
	QueryKey string `json:"queryKey"`
	// The mode of the metrics query that the user was editing. Can be `Basic` or `Advanced`. Will ONLY be specified for metrics queries. 
	MetricsQueryMode *string `json:"metricsQueryMode,omitempty"`
	MetricsQueryData *MetricsQueryData `json:"metricsQueryData,omitempty"`
	TracesQueryData *TracesQueryData `json:"tracesQueryData,omitempty"`
	SpansQueryData *SpansQueryData `json:"spansQueryData,omitempty"`
	// This field only applies for queryType of `Logs` but other query types may be supported in the future. Define the parsing mode to scan the JSON format log messages. Possible values are:   1. `Auto`   2. `Manual` In AutoParse mode, the system automatically figures out fields to parse based on the search query. While in the Manual mode, no fields are parsed out automatically. For more information see [Dynamic Parsing](https://help.sumologic.com/?cid=0011).
	ParseMode *string `json:"parseMode,omitempty"`
	// This field only applies for queryType of `Logs` but other query types may be supported in the future. Define the time source of this query. Possible values are `Message` and `Receipt`. `Message` will use the timeStamp on the message, while `Receipt` will use the timestamp it was received by Sumo.
	TimeSource *string `json:"timeSource,omitempty"`
	// This field only applies for queryType of `Metrics` but other query types may be supported in the future. Determines if the row should be returned in the response. Can be used in conjunction with a join, if only the result of the join is needed, and not the intermediate rows. Setting `transient` to `true`  wherever the intermediate results aren't required speeds up the computation and reduces the amount of data  transferred over the network.
	Transient *bool `json:"transient,omitempty"`
	// This field only applies for queryType of `Metrics` but other query types may be supported in the future. Specifies the output cardinality limitations for the query, which is the maximum number of timeseries returned in the result.
	OutputCardinalityLimit *int32 `json:"outputCardinalityLimit,omitempty"`
}

type _Query Query

// NewQuery instantiates a new Query object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuery(queryString string, queryType string, queryKey string) *Query {
	this := Query{}
	this.QueryString = queryString
	this.QueryType = queryType
	this.QueryKey = queryKey
	var parseMode string = "Auto"
	this.ParseMode = &parseMode
	var timeSource string = "Message"
	this.TimeSource = &timeSource
	var transient bool = false
	this.Transient = &transient
	var outputCardinalityLimit int32 = 1000
	this.OutputCardinalityLimit = &outputCardinalityLimit
	return &this
}

// NewQueryWithDefaults instantiates a new Query object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryWithDefaults() *Query {
	this := Query{}
	var parseMode string = "Auto"
	this.ParseMode = &parseMode
	var timeSource string = "Message"
	this.TimeSource = &timeSource
	var transient bool = false
	this.Transient = &transient
	var outputCardinalityLimit int32 = 1000
	this.OutputCardinalityLimit = &outputCardinalityLimit
	return &this
}

// GetQueryString returns the QueryString field value
func (o *Query) GetQueryString() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QueryString
}

// GetQueryStringOk returns a tuple with the QueryString field value
// and a boolean to check if the value has been set.
func (o *Query) GetQueryStringOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryString, true
}

// SetQueryString sets field value
func (o *Query) SetQueryString(v string) {
	o.QueryString = v
}

// GetQueryType returns the QueryType field value
func (o *Query) GetQueryType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QueryType
}

// GetQueryTypeOk returns a tuple with the QueryType field value
// and a boolean to check if the value has been set.
func (o *Query) GetQueryTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryType, true
}

// SetQueryType sets field value
func (o *Query) SetQueryType(v string) {
	o.QueryType = v
}

// GetQueryKey returns the QueryKey field value
func (o *Query) GetQueryKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QueryKey
}

// GetQueryKeyOk returns a tuple with the QueryKey field value
// and a boolean to check if the value has been set.
func (o *Query) GetQueryKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryKey, true
}

// SetQueryKey sets field value
func (o *Query) SetQueryKey(v string) {
	o.QueryKey = v
}

// GetMetricsQueryMode returns the MetricsQueryMode field value if set, zero value otherwise.
func (o *Query) GetMetricsQueryMode() string {
	if o == nil || IsNil(o.MetricsQueryMode) {
		var ret string
		return ret
	}
	return *o.MetricsQueryMode
}

// GetMetricsQueryModeOk returns a tuple with the MetricsQueryMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Query) GetMetricsQueryModeOk() (*string, bool) {
	if o == nil || IsNil(o.MetricsQueryMode) {
		return nil, false
	}
	return o.MetricsQueryMode, true
}

// HasMetricsQueryMode returns a boolean if a field has been set.
func (o *Query) HasMetricsQueryMode() bool {
	if o != nil && !IsNil(o.MetricsQueryMode) {
		return true
	}

	return false
}

// SetMetricsQueryMode gets a reference to the given string and assigns it to the MetricsQueryMode field.
func (o *Query) SetMetricsQueryMode(v string) {
	o.MetricsQueryMode = &v
}

// GetMetricsQueryData returns the MetricsQueryData field value if set, zero value otherwise.
func (o *Query) GetMetricsQueryData() MetricsQueryData {
	if o == nil || IsNil(o.MetricsQueryData) {
		var ret MetricsQueryData
		return ret
	}
	return *o.MetricsQueryData
}

// GetMetricsQueryDataOk returns a tuple with the MetricsQueryData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Query) GetMetricsQueryDataOk() (*MetricsQueryData, bool) {
	if o == nil || IsNil(o.MetricsQueryData) {
		return nil, false
	}
	return o.MetricsQueryData, true
}

// HasMetricsQueryData returns a boolean if a field has been set.
func (o *Query) HasMetricsQueryData() bool {
	if o != nil && !IsNil(o.MetricsQueryData) {
		return true
	}

	return false
}

// SetMetricsQueryData gets a reference to the given MetricsQueryData and assigns it to the MetricsQueryData field.
func (o *Query) SetMetricsQueryData(v MetricsQueryData) {
	o.MetricsQueryData = &v
}

// GetTracesQueryData returns the TracesQueryData field value if set, zero value otherwise.
func (o *Query) GetTracesQueryData() TracesQueryData {
	if o == nil || IsNil(o.TracesQueryData) {
		var ret TracesQueryData
		return ret
	}
	return *o.TracesQueryData
}

// GetTracesQueryDataOk returns a tuple with the TracesQueryData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Query) GetTracesQueryDataOk() (*TracesQueryData, bool) {
	if o == nil || IsNil(o.TracesQueryData) {
		return nil, false
	}
	return o.TracesQueryData, true
}

// HasTracesQueryData returns a boolean if a field has been set.
func (o *Query) HasTracesQueryData() bool {
	if o != nil && !IsNil(o.TracesQueryData) {
		return true
	}

	return false
}

// SetTracesQueryData gets a reference to the given TracesQueryData and assigns it to the TracesQueryData field.
func (o *Query) SetTracesQueryData(v TracesQueryData) {
	o.TracesQueryData = &v
}

// GetSpansQueryData returns the SpansQueryData field value if set, zero value otherwise.
func (o *Query) GetSpansQueryData() SpansQueryData {
	if o == nil || IsNil(o.SpansQueryData) {
		var ret SpansQueryData
		return ret
	}
	return *o.SpansQueryData
}

// GetSpansQueryDataOk returns a tuple with the SpansQueryData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Query) GetSpansQueryDataOk() (*SpansQueryData, bool) {
	if o == nil || IsNil(o.SpansQueryData) {
		return nil, false
	}
	return o.SpansQueryData, true
}

// HasSpansQueryData returns a boolean if a field has been set.
func (o *Query) HasSpansQueryData() bool {
	if o != nil && !IsNil(o.SpansQueryData) {
		return true
	}

	return false
}

// SetSpansQueryData gets a reference to the given SpansQueryData and assigns it to the SpansQueryData field.
func (o *Query) SetSpansQueryData(v SpansQueryData) {
	o.SpansQueryData = &v
}

// GetParseMode returns the ParseMode field value if set, zero value otherwise.
func (o *Query) GetParseMode() string {
	if o == nil || IsNil(o.ParseMode) {
		var ret string
		return ret
	}
	return *o.ParseMode
}

// GetParseModeOk returns a tuple with the ParseMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Query) GetParseModeOk() (*string, bool) {
	if o == nil || IsNil(o.ParseMode) {
		return nil, false
	}
	return o.ParseMode, true
}

// HasParseMode returns a boolean if a field has been set.
func (o *Query) HasParseMode() bool {
	if o != nil && !IsNil(o.ParseMode) {
		return true
	}

	return false
}

// SetParseMode gets a reference to the given string and assigns it to the ParseMode field.
func (o *Query) SetParseMode(v string) {
	o.ParseMode = &v
}

// GetTimeSource returns the TimeSource field value if set, zero value otherwise.
func (o *Query) GetTimeSource() string {
	if o == nil || IsNil(o.TimeSource) {
		var ret string
		return ret
	}
	return *o.TimeSource
}

// GetTimeSourceOk returns a tuple with the TimeSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Query) GetTimeSourceOk() (*string, bool) {
	if o == nil || IsNil(o.TimeSource) {
		return nil, false
	}
	return o.TimeSource, true
}

// HasTimeSource returns a boolean if a field has been set.
func (o *Query) HasTimeSource() bool {
	if o != nil && !IsNil(o.TimeSource) {
		return true
	}

	return false
}

// SetTimeSource gets a reference to the given string and assigns it to the TimeSource field.
func (o *Query) SetTimeSource(v string) {
	o.TimeSource = &v
}

// GetTransient returns the Transient field value if set, zero value otherwise.
func (o *Query) GetTransient() bool {
	if o == nil || IsNil(o.Transient) {
		var ret bool
		return ret
	}
	return *o.Transient
}

// GetTransientOk returns a tuple with the Transient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Query) GetTransientOk() (*bool, bool) {
	if o == nil || IsNil(o.Transient) {
		return nil, false
	}
	return o.Transient, true
}

// HasTransient returns a boolean if a field has been set.
func (o *Query) HasTransient() bool {
	if o != nil && !IsNil(o.Transient) {
		return true
	}

	return false
}

// SetTransient gets a reference to the given bool and assigns it to the Transient field.
func (o *Query) SetTransient(v bool) {
	o.Transient = &v
}

// GetOutputCardinalityLimit returns the OutputCardinalityLimit field value if set, zero value otherwise.
func (o *Query) GetOutputCardinalityLimit() int32 {
	if o == nil || IsNil(o.OutputCardinalityLimit) {
		var ret int32
		return ret
	}
	return *o.OutputCardinalityLimit
}

// GetOutputCardinalityLimitOk returns a tuple with the OutputCardinalityLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Query) GetOutputCardinalityLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.OutputCardinalityLimit) {
		return nil, false
	}
	return o.OutputCardinalityLimit, true
}

// HasOutputCardinalityLimit returns a boolean if a field has been set.
func (o *Query) HasOutputCardinalityLimit() bool {
	if o != nil && !IsNil(o.OutputCardinalityLimit) {
		return true
	}

	return false
}

// SetOutputCardinalityLimit gets a reference to the given int32 and assigns it to the OutputCardinalityLimit field.
func (o *Query) SetOutputCardinalityLimit(v int32) {
	o.OutputCardinalityLimit = &v
}

func (o Query) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Query) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["queryString"] = o.QueryString
	toSerialize["queryType"] = o.QueryType
	toSerialize["queryKey"] = o.QueryKey
	if !IsNil(o.MetricsQueryMode) {
		toSerialize["metricsQueryMode"] = o.MetricsQueryMode
	}
	if !IsNil(o.MetricsQueryData) {
		toSerialize["metricsQueryData"] = o.MetricsQueryData
	}
	if !IsNil(o.TracesQueryData) {
		toSerialize["tracesQueryData"] = o.TracesQueryData
	}
	if !IsNil(o.SpansQueryData) {
		toSerialize["spansQueryData"] = o.SpansQueryData
	}
	if !IsNil(o.ParseMode) {
		toSerialize["parseMode"] = o.ParseMode
	}
	if !IsNil(o.TimeSource) {
		toSerialize["timeSource"] = o.TimeSource
	}
	if !IsNil(o.Transient) {
		toSerialize["transient"] = o.Transient
	}
	if !IsNil(o.OutputCardinalityLimit) {
		toSerialize["outputCardinalityLimit"] = o.OutputCardinalityLimit
	}
	return toSerialize, nil
}

func (o *Query) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"queryString",
		"queryType",
		"queryKey",
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

	varQuery := _Query{}

	err = json.Unmarshal(bytes, &varQuery)

	if err != nil {
		return err
	}

	*o = Query(varQuery)

	return err
}

type NullableQuery struct {
	value *Query
	isSet bool
}

func (v NullableQuery) Get() *Query {
	return v.value
}

func (v *NullableQuery) Set(val *Query) {
	v.value = val
	v.isSet = true
}

func (v NullableQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuery(val *Query) *NullableQuery {
	return &NullableQuery{value: val, isSet: true}
}

func (v NullableQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


