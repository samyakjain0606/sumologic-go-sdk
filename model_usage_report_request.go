/*
Sumo Logic API

# Getting Started Welcome to the Sumo Logic API reference. You can use these APIs to interact with the Sumo Logic platform. For information on the collector and search APIs, see our [API home page](https://help.sumologic.com/docs/api). ## API Endpoints Sumo Logic has several deployments in different geographic locations. You'll need to use the Sumo Logic API endpoint corresponding to your geographic location. See the table below for the different API endpoints by deployment. For details determining your account's deployment, see [API endpoints](https://help.sumologic.com/?cid=3011).    <table>     <tr>       <td> <strong>Deployment</strong> </td>       <td> <strong>Endpoint</strong> </td>     </tr>     <tr>       <td> AU </td>       <td> https://api.au.sumologic.com/api/ </td>     </tr>     <tr>       <td> CA </td>       <td> https://api.ca.sumologic.com/api/ </td>     </tr>     <tr>       <td> DE </td>       <td> https://api.de.sumologic.com/api/ </td>     </tr>     <tr>       <td> EU </td>       <td> https://api.eu.sumologic.com/api/ </td>     </tr>     <tr>       <td> FED </td>       <td> https://api.fed.sumologic.com/api/ </td>     </tr>     <tr>       <td> IN </td>       <td> https://api.in.sumologic.com/api/ </td>     </tr>     <tr>       <td> JP </td>       <td> https://api.jp.sumologic.com/api/ </td>     </tr>     <tr>       <td> US1 </td>       <td> https://api.sumologic.com/api/ </td>     </tr>     <tr>       <td> US2 </td>       <td> https://api.us2.sumologic.com/api/ </td>     </tr>   </table>  ## Authentication Sumo Logic supports the following options for API authentication: - Access ID and Access Key - Base64 encoded Access ID and Access Key  See [Access Keys](https://help.sumologic.com/docs/manage/security/access-keys) to generate an Access Key. Make sure to copy the key you create, because it is displayed only once. When you have an Access ID and Access Key you can execute requests such as the following:   ```bash   curl -u \"<accessId>:<accessKey>\" -X GET https://api.<deployment>.sumologic.com/api/v1/users   ```  Where `deployment` is either `au`, `ca`, `de`, `eu`, `fed`, `in`, `jp`, `us1`, or `us2`. See [API endpoints](#section/API-Endpoints) for details.  If you prefer to use basic access authentication, you can do a Base64 encoding of your `<accessId>:<accessKey>` to authenticate your HTTPS request. The following is an example request, replace the placeholder `<encoded>` with your encoded Access ID and Access Key string:   ```bash   curl -H \"Authorization: Basic <encoded>\" -X GET https://api.<deployment>.sumologic.com/api/v1/users   ```   Refer to [API Authentication](https://help.sumologic.com/?cid=3012) for a Base64 example.  ## Status Codes Generic status codes that apply to all our APIs. See the [HTTP status code registry](https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml) for reference.   <table>     <tr>       <td> <strong>HTTP Status Code</strong> </td>       <td> <strong>Error Code</strong> </td>       <td> <strong>Description</strong> </td>     </tr>     <tr>       <td> 301 </td>       <td> moved </td>       <td> The requested resource SHOULD be accessed through returned URI in Location Header. See [troubleshooting](https://help.sumologic.com/docs/api/troubleshooting/#api---301-error---moved) for details.</td>     </tr>     <tr>       <td> 401 </td>       <td> unauthorized </td>       <td> Credential could not be verified.</td>     </tr>     <tr>       <td> 403 </td>       <td> forbidden </td>       <td> This operation is not allowed for your account type or the user doesn't have the role capability to perform this action. See [troubleshooting](https://help.sumologic.com/docs/api/troubleshooting/#api---401-error---credential-could-not-be-verified) for details.</td>     </tr>     <tr>       <td> 404 </td>       <td> notfound </td>       <td> Requested resource could not be found. </td>     </tr>     <tr>       <td> 405 </td>       <td> method.unsupported </td>       <td> Unsupported method for URL. </td>     </tr>     <tr>       <td> 415 </td>       <td> contenttype.invalid </td>       <td> Invalid content type. </td>     </tr>     <tr>       <td> 429 </td>       <td> rate.limit.exceeded </td>       <td> The API request rate is higher than 4 request per second or inflight API requests are higher than 10 request per second. </td>     </tr>     <tr>       <td> 500 </td>       <td> internal.error </td>       <td> Internal server error. </td>     </tr>     <tr>       <td> 503 </td>       <td> service.unavailable </td>       <td> Service is currently unavailable. </td>     </tr>   </table>  ## Filtering Some API endpoints support filtering results on a specified set of fields. Each endpoint that supports filtering will list the fields that can be filtered. Multiple fields can be combined by using an ampersand `&` character.  For example, to get 20 users whose `firstName` is `John` and `lastName` is `Doe`:   ```bash   api.sumologic.com/v1/users?limit=20&firstName=John&lastName=Doe   ```  ## Sorting Some API endpoints support sorting fields by using the `sortBy` query parameter. The default sort order is ascending. Prefix the field with a minus sign `-` to sort in descending order.  For example, to get 20 users sorted by their `email` in descending order:   ```bash   api.sumologic.com/v1/users?limit=20&sort=-email   ```  ## Asynchronous Request Asynchronous requests do not wait for results, instead they immediately respond back with a job identifier while the job runs in the background. You can use the job identifier to track the status of the asynchronous job request. Here is a typical flow for an asynchronous request. 1. Start an asynchronous job. On success, a job identifier is returned. The job identifier uniquely identifies   your asynchronous job.  2. Once started, use the job identifier from step 1 to track the status of your asynchronous job. An asynchronous   request will typically provide an endpoint to poll for the status of asynchronous job. A successful response   from the status endpoint will have the following structure:   ```json   {       \"status\": \"Status of asynchronous request\",       \"statusMessage\": \"Optional message with additional information in case request succeeds\",       \"error\": \"Error object in case request fails\"   }   ```   The `status` field can have one of the following values:     1. `Success`: The job succeeded. The `statusMessage` field might have additional information.     2. `InProgress`: The job is still running.     3. `Failed`: The job failed. The `error` field in the response will have more information about the failure.  3. Some asynchronous APIs may provide a third endpoint (like [export result](#operation/getAsyncExportResult))   to fetch the result of an asynchronous job.   ### Example Let's say we want to export a folder with the identifier `0000000006A2E86F`. We will use the [async export](#operation/beginAsyncExport) API to export all the content under the folder with `id=0000000006A2E86F`. 1. Start an export job for the folder   ```bash   curl -X POST -u \"<accessId>:<accessKey>\" https://api.<deployment>.sumologic.com/api/v2/content/0000000006A2E86F/export   ```   See [authentication section](#section/Authentication) for more details about `accessId`, `accessKey`, and   `deployment`.   On success, you will get back a job identifier. In the response below, `C03E086C137F38B4` is the job identifier.   ```bash   {       \"id\": \"C03E086C137F38B4\"   }   ```  2. Now poll for the status of the asynchronous job with the [status](#operation/getAsyncExportStatus) endpoint.   ```bash   curl -X GET -u \"<accessId>:<accessKey>\" https://api.<deployment>.sumologic.com/api/v2/content/0000000006A2E86F/export/C03E086C137F38B4/status   ```   You may get a response like   ```json   {       \"status\": \"InProgress\",       \"statusMessage\": null,       \"error\": null   }   ```   It implies the job is still in progress. Keep polling till the status is either `Success` or `Failed`.  3. When the asynchronous job completes (`status != \"InProgress\"`), you can fetch the results with the   [export result](#operation/getAsyncExportResult) endpoint.   ```bash   curl -X GET -u \"<accessId>:<accessKey>\" https://api.<deployment>.sumologic.com/api/v2/content/0000000006A2E86F/export/C03E086C137F38B4/result   ```    The asynchronous job may fail (`status == \"Failed\"`). You can look at the `error` field for more details.   ```json   {       \"status\": \"Failed\",       \"errors\": {           \"code\": \"content1:too_many_items\",           \"message\": \"Too many objects: object count(1100) was greater than limit 1000\"       }   }   ```   ## Rate Limiting * A rate limit of four API requests per second (240 requests per minute) applies to all API calls from a user. * A rate limit of 10 concurrent requests to any API endpoint applies to an access key.  If a rate is exceeded, a rate limit exceeded 429 status code is returned.  ## Generating Clients You can use [OpenAPI Generator](https://openapi-generator.tech) to generate clients from the YAML file to access the API.  ### Using [NPM](https://www.npmjs.com/get-npm) 1. Install [NPM package wrapper](https://github.com/openapitools/openapi-generator-cli) globally, exposing the CLI   on the command line:   ```bash   npm install @openapitools/openapi-generator-cli -g   ```   You can see detailed instructions [here](https://openapi-generator.tech/docs/installation#npm).  2. Download the [YAML file](/docs/sumologic-api.yaml) and save it locally. Let's say the file is saved as `sumologic-api.yaml`. 3. Use the following command to generate `python` client inside the `sumo/client/python` directory:   ```bash   openapi-generator generate -i sumologic-api.yaml -g python -o sumo/client/python   ```   ### Using [Homebrew](https://brew.sh/) 1. Install OpenAPI Generator   ```bash   brew install openapi-generator   ```  2. Download the [YAML file](/docs/sumologic-api.yaml) and save it locally. Let's say the file is saved as `sumologic-api.yaml`. 3. Use the following command to generate `python` client side code inside the `sumo/client/python` directory:   ```bash   openapi-generator generate -i sumologic-api.yaml -g python -o sumo/client/python   ``` 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the UsageReportRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsageReportRequest{}

// UsageReportRequest Usage Export Report Request
type UsageReportRequest struct {
	// Start date, without the time, of the usage data to fetch. If no value is provided startDate is used as the start of the subscription. The start date cannot be before the start of the subscription.
	StartDate *string `json:"startDate,omitempty"`
	// End date, without the time, of usage data to fetch. If no value is provided endDate is used as the end of the subscription. The end date cannot be after the end of the subscription.
	EndDate *string `json:"endDate,omitempty"`
	// Perform a groupBy operation on the usage details. If no value is provided data is grouped by `Day` - `day`: Aggregate the data by day - `week`: Aggregate the data by week. Week starts at Monday and ends at sunday night. - `month`: Aggregate the data by calendar month.
	GroupBy *string `json:"groupBy,omitempty"`
	// Specifies the type of report to be exported. Available types are `standard` and `detailed`. An additional `childDetailed` type is available for Sumo Orgs parents. Detailed report will have raw consumption along with the credits breakdown. If no value is provided Standard reports will be exported.
	ReportType *string `json:"reportType,omitempty"`
	// Deployment charges will be applied to the returned usages csv if this is set to true and the organization  is a part of Sumo Organizations as a child organization.
	IncludeDeploymentCharge *bool `json:"includeDeploymentCharge,omitempty"`
}

// NewUsageReportRequest instantiates a new UsageReportRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageReportRequest() *UsageReportRequest {
	this := UsageReportRequest{}
	var groupBy string = "day"
	this.GroupBy = &groupBy
	var reportType string = "standard"
	this.ReportType = &reportType
	var includeDeploymentCharge bool = false
	this.IncludeDeploymentCharge = &includeDeploymentCharge
	return &this
}

// NewUsageReportRequestWithDefaults instantiates a new UsageReportRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageReportRequestWithDefaults() *UsageReportRequest {
	this := UsageReportRequest{}
	var groupBy string = "day"
	this.GroupBy = &groupBy
	var reportType string = "standard"
	this.ReportType = &reportType
	var includeDeploymentCharge bool = false
	this.IncludeDeploymentCharge = &includeDeploymentCharge
	return &this
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *UsageReportRequest) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageReportRequest) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *UsageReportRequest) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *UsageReportRequest) SetStartDate(v string) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *UsageReportRequest) GetEndDate() string {
	if o == nil || IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageReportRequest) GetEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *UsageReportRequest) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *UsageReportRequest) SetEndDate(v string) {
	o.EndDate = &v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *UsageReportRequest) GetGroupBy() string {
	if o == nil || IsNil(o.GroupBy) {
		var ret string
		return ret
	}
	return *o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageReportRequest) GetGroupByOk() (*string, bool) {
	if o == nil || IsNil(o.GroupBy) {
		return nil, false
	}
	return o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *UsageReportRequest) HasGroupBy() bool {
	if o != nil && !IsNil(o.GroupBy) {
		return true
	}

	return false
}

// SetGroupBy gets a reference to the given string and assigns it to the GroupBy field.
func (o *UsageReportRequest) SetGroupBy(v string) {
	o.GroupBy = &v
}

// GetReportType returns the ReportType field value if set, zero value otherwise.
func (o *UsageReportRequest) GetReportType() string {
	if o == nil || IsNil(o.ReportType) {
		var ret string
		return ret
	}
	return *o.ReportType
}

// GetReportTypeOk returns a tuple with the ReportType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageReportRequest) GetReportTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ReportType) {
		return nil, false
	}
	return o.ReportType, true
}

// HasReportType returns a boolean if a field has been set.
func (o *UsageReportRequest) HasReportType() bool {
	if o != nil && !IsNil(o.ReportType) {
		return true
	}

	return false
}

// SetReportType gets a reference to the given string and assigns it to the ReportType field.
func (o *UsageReportRequest) SetReportType(v string) {
	o.ReportType = &v
}

// GetIncludeDeploymentCharge returns the IncludeDeploymentCharge field value if set, zero value otherwise.
func (o *UsageReportRequest) GetIncludeDeploymentCharge() bool {
	if o == nil || IsNil(o.IncludeDeploymentCharge) {
		var ret bool
		return ret
	}
	return *o.IncludeDeploymentCharge
}

// GetIncludeDeploymentChargeOk returns a tuple with the IncludeDeploymentCharge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageReportRequest) GetIncludeDeploymentChargeOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeDeploymentCharge) {
		return nil, false
	}
	return o.IncludeDeploymentCharge, true
}

// HasIncludeDeploymentCharge returns a boolean if a field has been set.
func (o *UsageReportRequest) HasIncludeDeploymentCharge() bool {
	if o != nil && !IsNil(o.IncludeDeploymentCharge) {
		return true
	}

	return false
}

// SetIncludeDeploymentCharge gets a reference to the given bool and assigns it to the IncludeDeploymentCharge field.
func (o *UsageReportRequest) SetIncludeDeploymentCharge(v bool) {
	o.IncludeDeploymentCharge = &v
}

func (o UsageReportRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsageReportRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	if !IsNil(o.GroupBy) {
		toSerialize["groupBy"] = o.GroupBy
	}
	if !IsNil(o.ReportType) {
		toSerialize["reportType"] = o.ReportType
	}
	if !IsNil(o.IncludeDeploymentCharge) {
		toSerialize["includeDeploymentCharge"] = o.IncludeDeploymentCharge
	}
	return toSerialize, nil
}

type NullableUsageReportRequest struct {
	value *UsageReportRequest
	isSet bool
}

func (v NullableUsageReportRequest) Get() *UsageReportRequest {
	return v.value
}

func (v *NullableUsageReportRequest) Set(val *UsageReportRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageReportRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageReportRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageReportRequest(val *UsageReportRequest) *NullableUsageReportRequest {
	return &NullableUsageReportRequest{value: val, isSet: true}
}

func (v NullableUsageReportRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageReportRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


