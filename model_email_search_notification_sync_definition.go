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

// checks if the EmailSearchNotificationSyncDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailSearchNotificationSyncDefinition{}

// EmailSearchNotificationSyncDefinition struct for EmailSearchNotificationSyncDefinition
type EmailSearchNotificationSyncDefinition struct {
	ScheduleNotificationSyncDefinition
	// A list of email recipients.
	ToList []string `json:"toList"`
	// If the notification is scheduled with a threshold, the default subject template will be \"Search Alert: {{AlertCondition}} results found for {{SearchName}}\". For email notifications without a threshold, the default subject template is \"Search Results: {{SearchName}}\".
	SubjectTemplate *string `json:"subjectTemplate,omitempty"`
	// A boolean value to indicate if the search query should be included in the notification email.
	IncludeQuery *bool `json:"includeQuery,omitempty"`
	// A boolean value to indicate if the search result set should be included in the notification email.
	IncludeResultSet *bool `json:"includeResultSet,omitempty"`
	// A boolean value to indicate if the search result histogram should be included in the notification email.
	IncludeHistogram *bool `json:"includeHistogram,omitempty"`
	// A boolean value to indicate if the search results should be included in the notification email as a CSV attachment.
	IncludeCsvAttachment *bool `json:"includeCsvAttachment,omitempty"`
}

type _EmailSearchNotificationSyncDefinition EmailSearchNotificationSyncDefinition

// NewEmailSearchNotificationSyncDefinition instantiates a new EmailSearchNotificationSyncDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailSearchNotificationSyncDefinition(toList []string, taskType string) *EmailSearchNotificationSyncDefinition {
	this := EmailSearchNotificationSyncDefinition{}
	this.TaskType = taskType
	this.ToList = toList
	var includeQuery bool = true
	this.IncludeQuery = &includeQuery
	var includeResultSet bool = true
	this.IncludeResultSet = &includeResultSet
	var includeHistogram bool = true
	this.IncludeHistogram = &includeHistogram
	var includeCsvAttachment bool = false
	this.IncludeCsvAttachment = &includeCsvAttachment
	return &this
}

// NewEmailSearchNotificationSyncDefinitionWithDefaults instantiates a new EmailSearchNotificationSyncDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailSearchNotificationSyncDefinitionWithDefaults() *EmailSearchNotificationSyncDefinition {
	this := EmailSearchNotificationSyncDefinition{}
	var includeQuery bool = true
	this.IncludeQuery = &includeQuery
	var includeResultSet bool = true
	this.IncludeResultSet = &includeResultSet
	var includeHistogram bool = true
	this.IncludeHistogram = &includeHistogram
	var includeCsvAttachment bool = false
	this.IncludeCsvAttachment = &includeCsvAttachment
	return &this
}

// GetToList returns the ToList field value
func (o *EmailSearchNotificationSyncDefinition) GetToList() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ToList
}

// GetToListOk returns a tuple with the ToList field value
// and a boolean to check if the value has been set.
func (o *EmailSearchNotificationSyncDefinition) GetToListOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ToList, true
}

// SetToList sets field value
func (o *EmailSearchNotificationSyncDefinition) SetToList(v []string) {
	o.ToList = v
}

// GetSubjectTemplate returns the SubjectTemplate field value if set, zero value otherwise.
func (o *EmailSearchNotificationSyncDefinition) GetSubjectTemplate() string {
	if o == nil || IsNil(o.SubjectTemplate) {
		var ret string
		return ret
	}
	return *o.SubjectTemplate
}

// GetSubjectTemplateOk returns a tuple with the SubjectTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailSearchNotificationSyncDefinition) GetSubjectTemplateOk() (*string, bool) {
	if o == nil || IsNil(o.SubjectTemplate) {
		return nil, false
	}
	return o.SubjectTemplate, true
}

// HasSubjectTemplate returns a boolean if a field has been set.
func (o *EmailSearchNotificationSyncDefinition) HasSubjectTemplate() bool {
	if o != nil && !IsNil(o.SubjectTemplate) {
		return true
	}

	return false
}

// SetSubjectTemplate gets a reference to the given string and assigns it to the SubjectTemplate field.
func (o *EmailSearchNotificationSyncDefinition) SetSubjectTemplate(v string) {
	o.SubjectTemplate = &v
}

// GetIncludeQuery returns the IncludeQuery field value if set, zero value otherwise.
func (o *EmailSearchNotificationSyncDefinition) GetIncludeQuery() bool {
	if o == nil || IsNil(o.IncludeQuery) {
		var ret bool
		return ret
	}
	return *o.IncludeQuery
}

// GetIncludeQueryOk returns a tuple with the IncludeQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailSearchNotificationSyncDefinition) GetIncludeQueryOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeQuery) {
		return nil, false
	}
	return o.IncludeQuery, true
}

// HasIncludeQuery returns a boolean if a field has been set.
func (o *EmailSearchNotificationSyncDefinition) HasIncludeQuery() bool {
	if o != nil && !IsNil(o.IncludeQuery) {
		return true
	}

	return false
}

// SetIncludeQuery gets a reference to the given bool and assigns it to the IncludeQuery field.
func (o *EmailSearchNotificationSyncDefinition) SetIncludeQuery(v bool) {
	o.IncludeQuery = &v
}

// GetIncludeResultSet returns the IncludeResultSet field value if set, zero value otherwise.
func (o *EmailSearchNotificationSyncDefinition) GetIncludeResultSet() bool {
	if o == nil || IsNil(o.IncludeResultSet) {
		var ret bool
		return ret
	}
	return *o.IncludeResultSet
}

// GetIncludeResultSetOk returns a tuple with the IncludeResultSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailSearchNotificationSyncDefinition) GetIncludeResultSetOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeResultSet) {
		return nil, false
	}
	return o.IncludeResultSet, true
}

// HasIncludeResultSet returns a boolean if a field has been set.
func (o *EmailSearchNotificationSyncDefinition) HasIncludeResultSet() bool {
	if o != nil && !IsNil(o.IncludeResultSet) {
		return true
	}

	return false
}

// SetIncludeResultSet gets a reference to the given bool and assigns it to the IncludeResultSet field.
func (o *EmailSearchNotificationSyncDefinition) SetIncludeResultSet(v bool) {
	o.IncludeResultSet = &v
}

// GetIncludeHistogram returns the IncludeHistogram field value if set, zero value otherwise.
func (o *EmailSearchNotificationSyncDefinition) GetIncludeHistogram() bool {
	if o == nil || IsNil(o.IncludeHistogram) {
		var ret bool
		return ret
	}
	return *o.IncludeHistogram
}

// GetIncludeHistogramOk returns a tuple with the IncludeHistogram field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailSearchNotificationSyncDefinition) GetIncludeHistogramOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeHistogram) {
		return nil, false
	}
	return o.IncludeHistogram, true
}

// HasIncludeHistogram returns a boolean if a field has been set.
func (o *EmailSearchNotificationSyncDefinition) HasIncludeHistogram() bool {
	if o != nil && !IsNil(o.IncludeHistogram) {
		return true
	}

	return false
}

// SetIncludeHistogram gets a reference to the given bool and assigns it to the IncludeHistogram field.
func (o *EmailSearchNotificationSyncDefinition) SetIncludeHistogram(v bool) {
	o.IncludeHistogram = &v
}

// GetIncludeCsvAttachment returns the IncludeCsvAttachment field value if set, zero value otherwise.
func (o *EmailSearchNotificationSyncDefinition) GetIncludeCsvAttachment() bool {
	if o == nil || IsNil(o.IncludeCsvAttachment) {
		var ret bool
		return ret
	}
	return *o.IncludeCsvAttachment
}

// GetIncludeCsvAttachmentOk returns a tuple with the IncludeCsvAttachment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailSearchNotificationSyncDefinition) GetIncludeCsvAttachmentOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeCsvAttachment) {
		return nil, false
	}
	return o.IncludeCsvAttachment, true
}

// HasIncludeCsvAttachment returns a boolean if a field has been set.
func (o *EmailSearchNotificationSyncDefinition) HasIncludeCsvAttachment() bool {
	if o != nil && !IsNil(o.IncludeCsvAttachment) {
		return true
	}

	return false
}

// SetIncludeCsvAttachment gets a reference to the given bool and assigns it to the IncludeCsvAttachment field.
func (o *EmailSearchNotificationSyncDefinition) SetIncludeCsvAttachment(v bool) {
	o.IncludeCsvAttachment = &v
}

func (o EmailSearchNotificationSyncDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailSearchNotificationSyncDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedScheduleNotificationSyncDefinition, errScheduleNotificationSyncDefinition := json.Marshal(o.ScheduleNotificationSyncDefinition)
	if errScheduleNotificationSyncDefinition != nil {
		return map[string]interface{}{}, errScheduleNotificationSyncDefinition
	}
	errScheduleNotificationSyncDefinition = json.Unmarshal([]byte(serializedScheduleNotificationSyncDefinition), &toSerialize)
	if errScheduleNotificationSyncDefinition != nil {
		return map[string]interface{}{}, errScheduleNotificationSyncDefinition
	}
	toSerialize["toList"] = o.ToList
	if !IsNil(o.SubjectTemplate) {
		toSerialize["subjectTemplate"] = o.SubjectTemplate
	}
	if !IsNil(o.IncludeQuery) {
		toSerialize["includeQuery"] = o.IncludeQuery
	}
	if !IsNil(o.IncludeResultSet) {
		toSerialize["includeResultSet"] = o.IncludeResultSet
	}
	if !IsNil(o.IncludeHistogram) {
		toSerialize["includeHistogram"] = o.IncludeHistogram
	}
	if !IsNil(o.IncludeCsvAttachment) {
		toSerialize["includeCsvAttachment"] = o.IncludeCsvAttachment
	}
	return toSerialize, nil
}

func (o *EmailSearchNotificationSyncDefinition) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"toList",
		"taskType",
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

	varEmailSearchNotificationSyncDefinition := _EmailSearchNotificationSyncDefinition{}

	err = json.Unmarshal(bytes, &varEmailSearchNotificationSyncDefinition)

	if err != nil {
		return err
	}

	*o = EmailSearchNotificationSyncDefinition(varEmailSearchNotificationSyncDefinition)

	return err
}

type NullableEmailSearchNotificationSyncDefinition struct {
	value *EmailSearchNotificationSyncDefinition
	isSet bool
}

func (v NullableEmailSearchNotificationSyncDefinition) Get() *EmailSearchNotificationSyncDefinition {
	return v.value
}

func (v *NullableEmailSearchNotificationSyncDefinition) Set(val *EmailSearchNotificationSyncDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailSearchNotificationSyncDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailSearchNotificationSyncDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailSearchNotificationSyncDefinition(val *EmailSearchNotificationSyncDefinition) *NullableEmailSearchNotificationSyncDefinition {
	return &NullableEmailSearchNotificationSyncDefinition{value: val, isSet: true}
}

func (v NullableEmailSearchNotificationSyncDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailSearchNotificationSyncDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


