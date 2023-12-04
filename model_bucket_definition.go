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

// checks if the BucketDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BucketDefinition{}

// BucketDefinition struct for BucketDefinition
type BucketDefinition struct {
	// Name of the S3 data forwarding destination.
	DestinationName string `json:"destinationName"`
	// Description of the S3 data forwarding destination.
	Description *string `json:"description,omitempty"`
	// AWS IAM authentication method used for access. Possible values are: 1. `AccessKey` 2. `RoleBased`
	AuthenticationMode string `json:"authenticationMode"`
	// The AWS Access ID to access the S3 bucket.
	AccessKeyId *string `json:"accessKeyId,omitempty"`
	// The AWS Secret Key to access the S3 bucket.
	SecretAccessKey *string `json:"secretAccessKey,omitempty"`
	// The AWS Role ARN to access the S3 bucket.
	RoleArn *string `json:"roleArn,omitempty"`
	// The region where the S3 bucket is located.
	Region *string `json:"region,omitempty"`
	// Enable S3 server-side encryption.
	Encrypted *bool `json:"encrypted,omitempty"`
	// True if the destination is Active.
	Enabled *bool `json:"enabled,omitempty"`
	// The name of the Amazon S3 bucket.
	BucketName string `json:"bucketName"`
	// Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format.
	CreatedAt time.Time `json:"createdAt"`
	// Identifier of the user who created the resource.
	CreatedBy string `json:"createdBy"`
	// Last modification timestamp in UTC.
	ModifiedAt time.Time `json:"modifiedAt"`
	// Identifier of the user who last modified the resource.
	ModifiedBy string `json:"modifiedBy"`
	// The unique identifier of the data forwarding destination.
	Id string `json:"id"`
	// True if invalidated by the system.
	InvalidatedBySystem *bool `json:"invalidatedBySystem,omitempty"`
}

type _BucketDefinition BucketDefinition

// NewBucketDefinition instantiates a new BucketDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBucketDefinition(destinationName string, authenticationMode string, bucketName string, createdAt time.Time, createdBy string, modifiedAt time.Time, modifiedBy string, id string) *BucketDefinition {
	this := BucketDefinition{}
	this.DestinationName = destinationName
	this.AuthenticationMode = authenticationMode
	this.BucketName = bucketName
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.ModifiedAt = modifiedAt
	this.ModifiedBy = modifiedBy
	this.Id = id
	return &this
}

// NewBucketDefinitionWithDefaults instantiates a new BucketDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBucketDefinitionWithDefaults() *BucketDefinition {
	this := BucketDefinition{}
	return &this
}

// GetDestinationName returns the DestinationName field value
func (o *BucketDefinition) GetDestinationName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DestinationName
}

// GetDestinationNameOk returns a tuple with the DestinationName field value
// and a boolean to check if the value has been set.
func (o *BucketDefinition) GetDestinationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestinationName, true
}

// SetDestinationName sets field value
func (o *BucketDefinition) SetDestinationName(v string) {
	o.DestinationName = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BucketDefinition) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BucketDefinition) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BucketDefinition) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BucketDefinition) SetDescription(v string) {
	o.Description = &v
}

// GetAuthenticationMode returns the AuthenticationMode field value
func (o *BucketDefinition) GetAuthenticationMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthenticationMode
}

// GetAuthenticationModeOk returns a tuple with the AuthenticationMode field value
// and a boolean to check if the value has been set.
func (o *BucketDefinition) GetAuthenticationModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthenticationMode, true
}

// SetAuthenticationMode sets field value
func (o *BucketDefinition) SetAuthenticationMode(v string) {
	o.AuthenticationMode = v
}

// GetAccessKeyId returns the AccessKeyId field value if set, zero value otherwise.
func (o *BucketDefinition) GetAccessKeyId() string {
	if o == nil || IsNil(o.AccessKeyId) {
		var ret string
		return ret
	}
	return *o.AccessKeyId
}

// GetAccessKeyIdOk returns a tuple with the AccessKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BucketDefinition) GetAccessKeyIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccessKeyId) {
		return nil, false
	}
	return o.AccessKeyId, true
}

// HasAccessKeyId returns a boolean if a field has been set.
func (o *BucketDefinition) HasAccessKeyId() bool {
	if o != nil && !IsNil(o.AccessKeyId) {
		return true
	}

	return false
}

// SetAccessKeyId gets a reference to the given string and assigns it to the AccessKeyId field.
func (o *BucketDefinition) SetAccessKeyId(v string) {
	o.AccessKeyId = &v
}

// GetSecretAccessKey returns the SecretAccessKey field value if set, zero value otherwise.
func (o *BucketDefinition) GetSecretAccessKey() string {
	if o == nil || IsNil(o.SecretAccessKey) {
		var ret string
		return ret
	}
	return *o.SecretAccessKey
}

// GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BucketDefinition) GetSecretAccessKeyOk() (*string, bool) {
	if o == nil || IsNil(o.SecretAccessKey) {
		return nil, false
	}
	return o.SecretAccessKey, true
}

// HasSecretAccessKey returns a boolean if a field has been set.
func (o *BucketDefinition) HasSecretAccessKey() bool {
	if o != nil && !IsNil(o.SecretAccessKey) {
		return true
	}

	return false
}

// SetSecretAccessKey gets a reference to the given string and assigns it to the SecretAccessKey field.
func (o *BucketDefinition) SetSecretAccessKey(v string) {
	o.SecretAccessKey = &v
}

// GetRoleArn returns the RoleArn field value if set, zero value otherwise.
func (o *BucketDefinition) GetRoleArn() string {
	if o == nil || IsNil(o.RoleArn) {
		var ret string
		return ret
	}
	return *o.RoleArn
}

// GetRoleArnOk returns a tuple with the RoleArn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BucketDefinition) GetRoleArnOk() (*string, bool) {
	if o == nil || IsNil(o.RoleArn) {
		return nil, false
	}
	return o.RoleArn, true
}

// HasRoleArn returns a boolean if a field has been set.
func (o *BucketDefinition) HasRoleArn() bool {
	if o != nil && !IsNil(o.RoleArn) {
		return true
	}

	return false
}

// SetRoleArn gets a reference to the given string and assigns it to the RoleArn field.
func (o *BucketDefinition) SetRoleArn(v string) {
	o.RoleArn = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *BucketDefinition) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BucketDefinition) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *BucketDefinition) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *BucketDefinition) SetRegion(v string) {
	o.Region = &v
}

// GetEncrypted returns the Encrypted field value if set, zero value otherwise.
func (o *BucketDefinition) GetEncrypted() bool {
	if o == nil || IsNil(o.Encrypted) {
		var ret bool
		return ret
	}
	return *o.Encrypted
}

// GetEncryptedOk returns a tuple with the Encrypted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BucketDefinition) GetEncryptedOk() (*bool, bool) {
	if o == nil || IsNil(o.Encrypted) {
		return nil, false
	}
	return o.Encrypted, true
}

// HasEncrypted returns a boolean if a field has been set.
func (o *BucketDefinition) HasEncrypted() bool {
	if o != nil && !IsNil(o.Encrypted) {
		return true
	}

	return false
}

// SetEncrypted gets a reference to the given bool and assigns it to the Encrypted field.
func (o *BucketDefinition) SetEncrypted(v bool) {
	o.Encrypted = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *BucketDefinition) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BucketDefinition) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *BucketDefinition) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *BucketDefinition) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetBucketName returns the BucketName field value
func (o *BucketDefinition) GetBucketName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BucketName
}

// GetBucketNameOk returns a tuple with the BucketName field value
// and a boolean to check if the value has been set.
func (o *BucketDefinition) GetBucketNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BucketName, true
}

// SetBucketName sets field value
func (o *BucketDefinition) SetBucketName(v string) {
	o.BucketName = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *BucketDefinition) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *BucketDefinition) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *BucketDefinition) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *BucketDefinition) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *BucketDefinition) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *BucketDefinition) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetModifiedAt returns the ModifiedAt field value
func (o *BucketDefinition) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *BucketDefinition) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value
func (o *BucketDefinition) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// GetModifiedBy returns the ModifiedBy field value
func (o *BucketDefinition) GetModifiedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value
// and a boolean to check if the value has been set.
func (o *BucketDefinition) GetModifiedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedBy, true
}

// SetModifiedBy sets field value
func (o *BucketDefinition) SetModifiedBy(v string) {
	o.ModifiedBy = v
}

// GetId returns the Id field value
func (o *BucketDefinition) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BucketDefinition) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BucketDefinition) SetId(v string) {
	o.Id = v
}

// GetInvalidatedBySystem returns the InvalidatedBySystem field value if set, zero value otherwise.
func (o *BucketDefinition) GetInvalidatedBySystem() bool {
	if o == nil || IsNil(o.InvalidatedBySystem) {
		var ret bool
		return ret
	}
	return *o.InvalidatedBySystem
}

// GetInvalidatedBySystemOk returns a tuple with the InvalidatedBySystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BucketDefinition) GetInvalidatedBySystemOk() (*bool, bool) {
	if o == nil || IsNil(o.InvalidatedBySystem) {
		return nil, false
	}
	return o.InvalidatedBySystem, true
}

// HasInvalidatedBySystem returns a boolean if a field has been set.
func (o *BucketDefinition) HasInvalidatedBySystem() bool {
	if o != nil && !IsNil(o.InvalidatedBySystem) {
		return true
	}

	return false
}

// SetInvalidatedBySystem gets a reference to the given bool and assigns it to the InvalidatedBySystem field.
func (o *BucketDefinition) SetInvalidatedBySystem(v bool) {
	o.InvalidatedBySystem = &v
}

func (o BucketDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BucketDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["destinationName"] = o.DestinationName
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["authenticationMode"] = o.AuthenticationMode
	if !IsNil(o.AccessKeyId) {
		toSerialize["accessKeyId"] = o.AccessKeyId
	}
	if !IsNil(o.SecretAccessKey) {
		toSerialize["secretAccessKey"] = o.SecretAccessKey
	}
	if !IsNil(o.RoleArn) {
		toSerialize["roleArn"] = o.RoleArn
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.Encrypted) {
		toSerialize["encrypted"] = o.Encrypted
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	toSerialize["bucketName"] = o.BucketName
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["createdBy"] = o.CreatedBy
	toSerialize["modifiedAt"] = o.ModifiedAt
	toSerialize["modifiedBy"] = o.ModifiedBy
	toSerialize["id"] = o.Id
	if !IsNil(o.InvalidatedBySystem) {
		toSerialize["invalidatedBySystem"] = o.InvalidatedBySystem
	}
	return toSerialize, nil
}

func (o *BucketDefinition) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"destinationName",
		"authenticationMode",
		"bucketName",
		"createdAt",
		"createdBy",
		"modifiedAt",
		"modifiedBy",
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

	varBucketDefinition := _BucketDefinition{}

	err = json.Unmarshal(bytes, &varBucketDefinition)

	if err != nil {
		return err
	}

	*o = BucketDefinition(varBucketDefinition)

	return err
}

type NullableBucketDefinition struct {
	value *BucketDefinition
	isSet bool
}

func (v NullableBucketDefinition) Get() *BucketDefinition {
	return v.value
}

func (v *NullableBucketDefinition) Set(val *BucketDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableBucketDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableBucketDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBucketDefinition(val *BucketDefinition) *NullableBucketDefinition {
	return &NullableBucketDefinition{value: val, isSet: true}
}

func (v NullableBucketDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBucketDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


