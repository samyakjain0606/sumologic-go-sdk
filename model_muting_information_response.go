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

// checks if the MutingInformationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MutingInformationResponse{}

// MutingInformationResponse Muting information fields for the monitor.
type MutingInformationResponse struct {
	// Identifier of the monitor.
	Id *string `json:"id,omitempty"`
	// Flag to indicate the monitor muted or not.
	IsMuted bool `json:"isMuted"`
	// Timestamp in Epoch that this monitor is currently muted until.
	MutingEndTime *int64 `json:"mutingEndTime,omitempty"`
	// Array of muting schedules that this monitor is associated with.
	MutingSchedules []MutingScheduleResponse `json:"mutingSchedules,omitempty"`
	AdhocMuting *AdhocMutingResponse `json:"adhocMuting,omitempty"`
}

type _MutingInformationResponse MutingInformationResponse

// NewMutingInformationResponse instantiates a new MutingInformationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMutingInformationResponse(isMuted bool) *MutingInformationResponse {
	this := MutingInformationResponse{}
	this.IsMuted = isMuted
	return &this
}

// NewMutingInformationResponseWithDefaults instantiates a new MutingInformationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMutingInformationResponseWithDefaults() *MutingInformationResponse {
	this := MutingInformationResponse{}
	var isMuted bool = false
	this.IsMuted = isMuted
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MutingInformationResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutingInformationResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MutingInformationResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MutingInformationResponse) SetId(v string) {
	o.Id = &v
}

// GetIsMuted returns the IsMuted field value
func (o *MutingInformationResponse) GetIsMuted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsMuted
}

// GetIsMutedOk returns a tuple with the IsMuted field value
// and a boolean to check if the value has been set.
func (o *MutingInformationResponse) GetIsMutedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsMuted, true
}

// SetIsMuted sets field value
func (o *MutingInformationResponse) SetIsMuted(v bool) {
	o.IsMuted = v
}

// GetMutingEndTime returns the MutingEndTime field value if set, zero value otherwise.
func (o *MutingInformationResponse) GetMutingEndTime() int64 {
	if o == nil || IsNil(o.MutingEndTime) {
		var ret int64
		return ret
	}
	return *o.MutingEndTime
}

// GetMutingEndTimeOk returns a tuple with the MutingEndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutingInformationResponse) GetMutingEndTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.MutingEndTime) {
		return nil, false
	}
	return o.MutingEndTime, true
}

// HasMutingEndTime returns a boolean if a field has been set.
func (o *MutingInformationResponse) HasMutingEndTime() bool {
	if o != nil && !IsNil(o.MutingEndTime) {
		return true
	}

	return false
}

// SetMutingEndTime gets a reference to the given int64 and assigns it to the MutingEndTime field.
func (o *MutingInformationResponse) SetMutingEndTime(v int64) {
	o.MutingEndTime = &v
}

// GetMutingSchedules returns the MutingSchedules field value if set, zero value otherwise.
func (o *MutingInformationResponse) GetMutingSchedules() []MutingScheduleResponse {
	if o == nil || IsNil(o.MutingSchedules) {
		var ret []MutingScheduleResponse
		return ret
	}
	return o.MutingSchedules
}

// GetMutingSchedulesOk returns a tuple with the MutingSchedules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutingInformationResponse) GetMutingSchedulesOk() ([]MutingScheduleResponse, bool) {
	if o == nil || IsNil(o.MutingSchedules) {
		return nil, false
	}
	return o.MutingSchedules, true
}

// HasMutingSchedules returns a boolean if a field has been set.
func (o *MutingInformationResponse) HasMutingSchedules() bool {
	if o != nil && !IsNil(o.MutingSchedules) {
		return true
	}

	return false
}

// SetMutingSchedules gets a reference to the given []MutingScheduleResponse and assigns it to the MutingSchedules field.
func (o *MutingInformationResponse) SetMutingSchedules(v []MutingScheduleResponse) {
	o.MutingSchedules = v
}

// GetAdhocMuting returns the AdhocMuting field value if set, zero value otherwise.
func (o *MutingInformationResponse) GetAdhocMuting() AdhocMutingResponse {
	if o == nil || IsNil(o.AdhocMuting) {
		var ret AdhocMutingResponse
		return ret
	}
	return *o.AdhocMuting
}

// GetAdhocMutingOk returns a tuple with the AdhocMuting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutingInformationResponse) GetAdhocMutingOk() (*AdhocMutingResponse, bool) {
	if o == nil || IsNil(o.AdhocMuting) {
		return nil, false
	}
	return o.AdhocMuting, true
}

// HasAdhocMuting returns a boolean if a field has been set.
func (o *MutingInformationResponse) HasAdhocMuting() bool {
	if o != nil && !IsNil(o.AdhocMuting) {
		return true
	}

	return false
}

// SetAdhocMuting gets a reference to the given AdhocMutingResponse and assigns it to the AdhocMuting field.
func (o *MutingInformationResponse) SetAdhocMuting(v AdhocMutingResponse) {
	o.AdhocMuting = &v
}

func (o MutingInformationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MutingInformationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["isMuted"] = o.IsMuted
	if !IsNil(o.MutingEndTime) {
		toSerialize["mutingEndTime"] = o.MutingEndTime
	}
	if !IsNil(o.MutingSchedules) {
		toSerialize["mutingSchedules"] = o.MutingSchedules
	}
	if !IsNil(o.AdhocMuting) {
		toSerialize["adhocMuting"] = o.AdhocMuting
	}
	return toSerialize, nil
}

func (o *MutingInformationResponse) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"isMuted",
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

	varMutingInformationResponse := _MutingInformationResponse{}

	err = json.Unmarshal(bytes, &varMutingInformationResponse)

	if err != nil {
		return err
	}

	*o = MutingInformationResponse(varMutingInformationResponse)

	return err
}

type NullableMutingInformationResponse struct {
	value *MutingInformationResponse
	isSet bool
}

func (v NullableMutingInformationResponse) Get() *MutingInformationResponse {
	return v.value
}

func (v *NullableMutingInformationResponse) Set(val *MutingInformationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMutingInformationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMutingInformationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMutingInformationResponse(val *MutingInformationResponse) *NullableMutingInformationResponse {
	return &NullableMutingInformationResponse{value: val, isSet: true}
}

func (v NullableMutingInformationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMutingInformationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


