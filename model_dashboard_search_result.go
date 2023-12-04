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

// checks if the DashboardSearchResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DashboardSearchResult{}

// DashboardSearchResult struct for DashboardSearchResult
type DashboardSearchResult struct {
	Status DashboardSearchStatus `json:"status"`
	Axes VisualDataAxes `json:"axes"`
	// The series returned from a search.
	Series []VisualDataSeries `json:"series"`
	// Errors returned by backend.
	Errors []ErrorDescription `json:"errors,omitempty"`
	TimeRange *BeginBoundedTimeRange `json:"timeRange,omitempty"`
	// A user-generated string to uniquely identify the search request. This field can be safely ignored if you don't intend to identify a search request.
	RequestToken *string `json:"requestToken,omitempty"`
	// The expected ordering of the column fields in tabular format. If null or empty, the ordering is unknown or indeterminate. 
	FieldOrdering []string `json:"fieldOrdering,omitempty"`
	// The total number of scanned bytes from infrequent tier data for the query in bytes.
	InfrequentScannedBytes *float32 `json:"infrequentScannedBytes,omitempty"`
	ScannedBytes *ScannedBytes `json:"scannedBytes,omitempty"`
	// The backfill percentage of a continuous query.
	BackfillPercent *float32 `json:"backfillPercent,omitempty"`
}

type _DashboardSearchResult DashboardSearchResult

// NewDashboardSearchResult instantiates a new DashboardSearchResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDashboardSearchResult(status DashboardSearchStatus, axes VisualDataAxes, series []VisualDataSeries) *DashboardSearchResult {
	this := DashboardSearchResult{}
	this.Status = status
	this.Axes = axes
	this.Series = series
	return &this
}

// NewDashboardSearchResultWithDefaults instantiates a new DashboardSearchResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDashboardSearchResultWithDefaults() *DashboardSearchResult {
	this := DashboardSearchResult{}
	return &this
}

// GetStatus returns the Status field value
func (o *DashboardSearchResult) GetStatus() DashboardSearchStatus {
	if o == nil {
		var ret DashboardSearchStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *DashboardSearchResult) GetStatusOk() (*DashboardSearchStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *DashboardSearchResult) SetStatus(v DashboardSearchStatus) {
	o.Status = v
}

// GetAxes returns the Axes field value
func (o *DashboardSearchResult) GetAxes() VisualDataAxes {
	if o == nil {
		var ret VisualDataAxes
		return ret
	}

	return o.Axes
}

// GetAxesOk returns a tuple with the Axes field value
// and a boolean to check if the value has been set.
func (o *DashboardSearchResult) GetAxesOk() (*VisualDataAxes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Axes, true
}

// SetAxes sets field value
func (o *DashboardSearchResult) SetAxes(v VisualDataAxes) {
	o.Axes = v
}

// GetSeries returns the Series field value
func (o *DashboardSearchResult) GetSeries() []VisualDataSeries {
	if o == nil {
		var ret []VisualDataSeries
		return ret
	}

	return o.Series
}

// GetSeriesOk returns a tuple with the Series field value
// and a boolean to check if the value has been set.
func (o *DashboardSearchResult) GetSeriesOk() ([]VisualDataSeries, bool) {
	if o == nil {
		return nil, false
	}
	return o.Series, true
}

// SetSeries sets field value
func (o *DashboardSearchResult) SetSeries(v []VisualDataSeries) {
	o.Series = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *DashboardSearchResult) GetErrors() []ErrorDescription {
	if o == nil || IsNil(o.Errors) {
		var ret []ErrorDescription
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardSearchResult) GetErrorsOk() ([]ErrorDescription, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *DashboardSearchResult) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []ErrorDescription and assigns it to the Errors field.
func (o *DashboardSearchResult) SetErrors(v []ErrorDescription) {
	o.Errors = v
}

// GetTimeRange returns the TimeRange field value if set, zero value otherwise.
func (o *DashboardSearchResult) GetTimeRange() BeginBoundedTimeRange {
	if o == nil || IsNil(o.TimeRange) {
		var ret BeginBoundedTimeRange
		return ret
	}
	return *o.TimeRange
}

// GetTimeRangeOk returns a tuple with the TimeRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardSearchResult) GetTimeRangeOk() (*BeginBoundedTimeRange, bool) {
	if o == nil || IsNil(o.TimeRange) {
		return nil, false
	}
	return o.TimeRange, true
}

// HasTimeRange returns a boolean if a field has been set.
func (o *DashboardSearchResult) HasTimeRange() bool {
	if o != nil && !IsNil(o.TimeRange) {
		return true
	}

	return false
}

// SetTimeRange gets a reference to the given BeginBoundedTimeRange and assigns it to the TimeRange field.
func (o *DashboardSearchResult) SetTimeRange(v BeginBoundedTimeRange) {
	o.TimeRange = &v
}

// GetRequestToken returns the RequestToken field value if set, zero value otherwise.
func (o *DashboardSearchResult) GetRequestToken() string {
	if o == nil || IsNil(o.RequestToken) {
		var ret string
		return ret
	}
	return *o.RequestToken
}

// GetRequestTokenOk returns a tuple with the RequestToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardSearchResult) GetRequestTokenOk() (*string, bool) {
	if o == nil || IsNil(o.RequestToken) {
		return nil, false
	}
	return o.RequestToken, true
}

// HasRequestToken returns a boolean if a field has been set.
func (o *DashboardSearchResult) HasRequestToken() bool {
	if o != nil && !IsNil(o.RequestToken) {
		return true
	}

	return false
}

// SetRequestToken gets a reference to the given string and assigns it to the RequestToken field.
func (o *DashboardSearchResult) SetRequestToken(v string) {
	o.RequestToken = &v
}

// GetFieldOrdering returns the FieldOrdering field value if set, zero value otherwise.
func (o *DashboardSearchResult) GetFieldOrdering() []string {
	if o == nil || IsNil(o.FieldOrdering) {
		var ret []string
		return ret
	}
	return o.FieldOrdering
}

// GetFieldOrderingOk returns a tuple with the FieldOrdering field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardSearchResult) GetFieldOrderingOk() ([]string, bool) {
	if o == nil || IsNil(o.FieldOrdering) {
		return nil, false
	}
	return o.FieldOrdering, true
}

// HasFieldOrdering returns a boolean if a field has been set.
func (o *DashboardSearchResult) HasFieldOrdering() bool {
	if o != nil && !IsNil(o.FieldOrdering) {
		return true
	}

	return false
}

// SetFieldOrdering gets a reference to the given []string and assigns it to the FieldOrdering field.
func (o *DashboardSearchResult) SetFieldOrdering(v []string) {
	o.FieldOrdering = v
}

// GetInfrequentScannedBytes returns the InfrequentScannedBytes field value if set, zero value otherwise.
func (o *DashboardSearchResult) GetInfrequentScannedBytes() float32 {
	if o == nil || IsNil(o.InfrequentScannedBytes) {
		var ret float32
		return ret
	}
	return *o.InfrequentScannedBytes
}

// GetInfrequentScannedBytesOk returns a tuple with the InfrequentScannedBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardSearchResult) GetInfrequentScannedBytesOk() (*float32, bool) {
	if o == nil || IsNil(o.InfrequentScannedBytes) {
		return nil, false
	}
	return o.InfrequentScannedBytes, true
}

// HasInfrequentScannedBytes returns a boolean if a field has been set.
func (o *DashboardSearchResult) HasInfrequentScannedBytes() bool {
	if o != nil && !IsNil(o.InfrequentScannedBytes) {
		return true
	}

	return false
}

// SetInfrequentScannedBytes gets a reference to the given float32 and assigns it to the InfrequentScannedBytes field.
func (o *DashboardSearchResult) SetInfrequentScannedBytes(v float32) {
	o.InfrequentScannedBytes = &v
}

// GetScannedBytes returns the ScannedBytes field value if set, zero value otherwise.
func (o *DashboardSearchResult) GetScannedBytes() ScannedBytes {
	if o == nil || IsNil(o.ScannedBytes) {
		var ret ScannedBytes
		return ret
	}
	return *o.ScannedBytes
}

// GetScannedBytesOk returns a tuple with the ScannedBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardSearchResult) GetScannedBytesOk() (*ScannedBytes, bool) {
	if o == nil || IsNil(o.ScannedBytes) {
		return nil, false
	}
	return o.ScannedBytes, true
}

// HasScannedBytes returns a boolean if a field has been set.
func (o *DashboardSearchResult) HasScannedBytes() bool {
	if o != nil && !IsNil(o.ScannedBytes) {
		return true
	}

	return false
}

// SetScannedBytes gets a reference to the given ScannedBytes and assigns it to the ScannedBytes field.
func (o *DashboardSearchResult) SetScannedBytes(v ScannedBytes) {
	o.ScannedBytes = &v
}

// GetBackfillPercent returns the BackfillPercent field value if set, zero value otherwise.
func (o *DashboardSearchResult) GetBackfillPercent() float32 {
	if o == nil || IsNil(o.BackfillPercent) {
		var ret float32
		return ret
	}
	return *o.BackfillPercent
}

// GetBackfillPercentOk returns a tuple with the BackfillPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardSearchResult) GetBackfillPercentOk() (*float32, bool) {
	if o == nil || IsNil(o.BackfillPercent) {
		return nil, false
	}
	return o.BackfillPercent, true
}

// HasBackfillPercent returns a boolean if a field has been set.
func (o *DashboardSearchResult) HasBackfillPercent() bool {
	if o != nil && !IsNil(o.BackfillPercent) {
		return true
	}

	return false
}

// SetBackfillPercent gets a reference to the given float32 and assigns it to the BackfillPercent field.
func (o *DashboardSearchResult) SetBackfillPercent(v float32) {
	o.BackfillPercent = &v
}

func (o DashboardSearchResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DashboardSearchResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	toSerialize["axes"] = o.Axes
	toSerialize["series"] = o.Series
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !IsNil(o.TimeRange) {
		toSerialize["timeRange"] = o.TimeRange
	}
	if !IsNil(o.RequestToken) {
		toSerialize["requestToken"] = o.RequestToken
	}
	if !IsNil(o.FieldOrdering) {
		toSerialize["fieldOrdering"] = o.FieldOrdering
	}
	if !IsNil(o.InfrequentScannedBytes) {
		toSerialize["infrequentScannedBytes"] = o.InfrequentScannedBytes
	}
	if !IsNil(o.ScannedBytes) {
		toSerialize["scannedBytes"] = o.ScannedBytes
	}
	if !IsNil(o.BackfillPercent) {
		toSerialize["backfillPercent"] = o.BackfillPercent
	}
	return toSerialize, nil
}

func (o *DashboardSearchResult) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
		"axes",
		"series",
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

	varDashboardSearchResult := _DashboardSearchResult{}

	err = json.Unmarshal(bytes, &varDashboardSearchResult)

	if err != nil {
		return err
	}

	*o = DashboardSearchResult(varDashboardSearchResult)

	return err
}

type NullableDashboardSearchResult struct {
	value *DashboardSearchResult
	isSet bool
}

func (v NullableDashboardSearchResult) Get() *DashboardSearchResult {
	return v.value
}

func (v *NullableDashboardSearchResult) Set(val *DashboardSearchResult) {
	v.value = val
	v.isSet = true
}

func (v NullableDashboardSearchResult) IsSet() bool {
	return v.isSet
}

func (v *NullableDashboardSearchResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDashboardSearchResult(val *DashboardSearchResult) *NullableDashboardSearchResult {
	return &NullableDashboardSearchResult{value: val, isSet: true}
}

func (v NullableDashboardSearchResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDashboardSearchResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


