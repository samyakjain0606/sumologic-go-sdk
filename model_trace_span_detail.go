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

// checks if the TraceSpanDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TraceSpanDetail{}

// TraceSpanDetail struct for TraceSpanDetail
type TraceSpanDetail struct {
	// Identifier of the span.
	Id string `json:"id"`
	// Identifier of the parent span, if any. If the span has no parent it's considered a root span.
	ParentId *string `json:"parentId,omitempty"`
	// The name of the operation given to the span.
	OperationName string `json:"operationName"`
	// The name of the resource attached to the span.
	Resource *string `json:"resource,omitempty"`
	// The name of the service this span is part of.
	Service *string `json:"service,omitempty"`
	// Color hex code assigned to the service.
	ServiceColor *string `json:"serviceColor,omitempty"`
	// Defines type of service.
	ServiceType *string `json:"serviceType,omitempty"`
	// Number of nanoseconds the span lasted.
	Duration int64 `json:"duration"`
	// Date and time the span was started in the [ISO 8601 / RFC3339](https://tools.ietf.org/html/rfc3339) format.
	StartedAt time.Time `json:"startedAt"`
	Status TraceSpanStatus `json:"status"`
	// Span kind describes the relationship between the Span, its parents, and its children in a Trace. Possible values: `CLIENT`, `SERVER`, `PRODUCER`, `CONSUMER`, `INTERNAL`.
	Kind *string `json:"kind,omitempty"`
	// Name of the possible remote span's service.
	RemoteService *string `json:"remoteService,omitempty"`
	// Color hex code assigned to the remote service.
	RemoteServiceColor *string `json:"remoteServiceColor,omitempty"`
	// Defines type of service.
	RemoteServiceType *string `json:"remoteServiceType,omitempty"`
	Info *TraceSpanInfo `json:"info,omitempty"`
	// Number of span links in this span.
	NumberOfLinks *int32 `json:"numberOfLinks,omitempty"`
	// Produced error message (could be a stack trace, database error code, ..)
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// Fields attached to this span.
	Fields *map[string]TracingValue `json:"fields,omitempty"`
	CriticalPathContribution *TraceSpanCriticalPathContribution `json:"criticalPathContribution,omitempty"`
	// Logs attached to this span.
	Logs []string `json:"logs,omitempty"`
	// Events attached to this span.
	Events []SpanEvent `json:"events,omitempty"`
	// List of casually related spans.
	Links []SpanLink `json:"links,omitempty"`
}

type _TraceSpanDetail TraceSpanDetail

// NewTraceSpanDetail instantiates a new TraceSpanDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTraceSpanDetail(id string, operationName string, duration int64, startedAt time.Time, status TraceSpanStatus) *TraceSpanDetail {
	this := TraceSpanDetail{}
	this.Id = id
	this.OperationName = operationName
	this.Duration = duration
	this.StartedAt = startedAt
	this.Status = status
	return &this
}

// NewTraceSpanDetailWithDefaults instantiates a new TraceSpanDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTraceSpanDetailWithDefaults() *TraceSpanDetail {
	this := TraceSpanDetail{}
	return &this
}

// GetId returns the Id field value
func (o *TraceSpanDetail) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TraceSpanDetail) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TraceSpanDetail) SetId(v string) {
	o.Id = v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *TraceSpanDetail) GetParentId() string {
	if o == nil || IsNil(o.ParentId) {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceSpanDetail) GetParentIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentId) {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *TraceSpanDetail) HasParentId() bool {
	if o != nil && !IsNil(o.ParentId) {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *TraceSpanDetail) SetParentId(v string) {
	o.ParentId = &v
}

// GetOperationName returns the OperationName field value
func (o *TraceSpanDetail) GetOperationName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OperationName
}

// GetOperationNameOk returns a tuple with the OperationName field value
// and a boolean to check if the value has been set.
func (o *TraceSpanDetail) GetOperationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OperationName, true
}

// SetOperationName sets field value
func (o *TraceSpanDetail) SetOperationName(v string) {
	o.OperationName = v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *TraceSpanDetail) GetResource() string {
	if o == nil || IsNil(o.Resource) {
		var ret string
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceSpanDetail) GetResourceOk() (*string, bool) {
	if o == nil || IsNil(o.Resource) {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *TraceSpanDetail) HasResource() bool {
	if o != nil && !IsNil(o.Resource) {
		return true
	}

	return false
}

// SetResource gets a reference to the given string and assigns it to the Resource field.
func (o *TraceSpanDetail) SetResource(v string) {
	o.Resource = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *TraceSpanDetail) GetService() string {
	if o == nil || IsNil(o.Service) {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceSpanDetail) GetServiceOk() (*string, bool) {
	if o == nil || IsNil(o.Service) {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *TraceSpanDetail) HasService() bool {
	if o != nil && !IsNil(o.Service) {
		return true
	}

	return false
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *TraceSpanDetail) SetService(v string) {
	o.Service = &v
}

// GetServiceColor returns the ServiceColor field value if set, zero value otherwise.
func (o *TraceSpanDetail) GetServiceColor() string {
	if o == nil || IsNil(o.ServiceColor) {
		var ret string
		return ret
	}
	return *o.ServiceColor
}

// GetServiceColorOk returns a tuple with the ServiceColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceSpanDetail) GetServiceColorOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceColor) {
		return nil, false
	}
	return o.ServiceColor, true
}

// HasServiceColor returns a boolean if a field has been set.
func (o *TraceSpanDetail) HasServiceColor() bool {
	if o != nil && !IsNil(o.ServiceColor) {
		return true
	}

	return false
}

// SetServiceColor gets a reference to the given string and assigns it to the ServiceColor field.
func (o *TraceSpanDetail) SetServiceColor(v string) {
	o.ServiceColor = &v
}

// GetServiceType returns the ServiceType field value if set, zero value otherwise.
func (o *TraceSpanDetail) GetServiceType() string {
	if o == nil || IsNil(o.ServiceType) {
		var ret string
		return ret
	}
	return *o.ServiceType
}

// GetServiceTypeOk returns a tuple with the ServiceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceSpanDetail) GetServiceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceType) {
		return nil, false
	}
	return o.ServiceType, true
}

// HasServiceType returns a boolean if a field has been set.
func (o *TraceSpanDetail) HasServiceType() bool {
	if o != nil && !IsNil(o.ServiceType) {
		return true
	}

	return false
}

// SetServiceType gets a reference to the given string and assigns it to the ServiceType field.
func (o *TraceSpanDetail) SetServiceType(v string) {
	o.ServiceType = &v
}

// GetDuration returns the Duration field value
func (o *TraceSpanDetail) GetDuration() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *TraceSpanDetail) GetDurationOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *TraceSpanDetail) SetDuration(v int64) {
	o.Duration = v
}

// GetStartedAt returns the StartedAt field value
func (o *TraceSpanDetail) GetStartedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value
// and a boolean to check if the value has been set.
func (o *TraceSpanDetail) GetStartedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartedAt, true
}

// SetStartedAt sets field value
func (o *TraceSpanDetail) SetStartedAt(v time.Time) {
	o.StartedAt = v
}

// GetStatus returns the Status field value
func (o *TraceSpanDetail) GetStatus() TraceSpanStatus {
	if o == nil {
		var ret TraceSpanStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TraceSpanDetail) GetStatusOk() (*TraceSpanStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *TraceSpanDetail) SetStatus(v TraceSpanStatus) {
	o.Status = v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *TraceSpanDetail) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceSpanDetail) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *TraceSpanDetail) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *TraceSpanDetail) SetKind(v string) {
	o.Kind = &v
}

// GetRemoteService returns the RemoteService field value if set, zero value otherwise.
func (o *TraceSpanDetail) GetRemoteService() string {
	if o == nil || IsNil(o.RemoteService) {
		var ret string
		return ret
	}
	return *o.RemoteService
}

// GetRemoteServiceOk returns a tuple with the RemoteService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceSpanDetail) GetRemoteServiceOk() (*string, bool) {
	if o == nil || IsNil(o.RemoteService) {
		return nil, false
	}
	return o.RemoteService, true
}

// HasRemoteService returns a boolean if a field has been set.
func (o *TraceSpanDetail) HasRemoteService() bool {
	if o != nil && !IsNil(o.RemoteService) {
		return true
	}

	return false
}

// SetRemoteService gets a reference to the given string and assigns it to the RemoteService field.
func (o *TraceSpanDetail) SetRemoteService(v string) {
	o.RemoteService = &v
}

// GetRemoteServiceColor returns the RemoteServiceColor field value if set, zero value otherwise.
func (o *TraceSpanDetail) GetRemoteServiceColor() string {
	if o == nil || IsNil(o.RemoteServiceColor) {
		var ret string
		return ret
	}
	return *o.RemoteServiceColor
}

// GetRemoteServiceColorOk returns a tuple with the RemoteServiceColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceSpanDetail) GetRemoteServiceColorOk() (*string, bool) {
	if o == nil || IsNil(o.RemoteServiceColor) {
		return nil, false
	}
	return o.RemoteServiceColor, true
}

// HasRemoteServiceColor returns a boolean if a field has been set.
func (o *TraceSpanDetail) HasRemoteServiceColor() bool {
	if o != nil && !IsNil(o.RemoteServiceColor) {
		return true
	}

	return false
}

// SetRemoteServiceColor gets a reference to the given string and assigns it to the RemoteServiceColor field.
func (o *TraceSpanDetail) SetRemoteServiceColor(v string) {
	o.RemoteServiceColor = &v
}

// GetRemoteServiceType returns the RemoteServiceType field value if set, zero value otherwise.
func (o *TraceSpanDetail) GetRemoteServiceType() string {
	if o == nil || IsNil(o.RemoteServiceType) {
		var ret string
		return ret
	}
	return *o.RemoteServiceType
}

// GetRemoteServiceTypeOk returns a tuple with the RemoteServiceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceSpanDetail) GetRemoteServiceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RemoteServiceType) {
		return nil, false
	}
	return o.RemoteServiceType, true
}

// HasRemoteServiceType returns a boolean if a field has been set.
func (o *TraceSpanDetail) HasRemoteServiceType() bool {
	if o != nil && !IsNil(o.RemoteServiceType) {
		return true
	}

	return false
}

// SetRemoteServiceType gets a reference to the given string and assigns it to the RemoteServiceType field.
func (o *TraceSpanDetail) SetRemoteServiceType(v string) {
	o.RemoteServiceType = &v
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *TraceSpanDetail) GetInfo() TraceSpanInfo {
	if o == nil || IsNil(o.Info) {
		var ret TraceSpanInfo
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceSpanDetail) GetInfoOk() (*TraceSpanInfo, bool) {
	if o == nil || IsNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *TraceSpanDetail) HasInfo() bool {
	if o != nil && !IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given TraceSpanInfo and assigns it to the Info field.
func (o *TraceSpanDetail) SetInfo(v TraceSpanInfo) {
	o.Info = &v
}

// GetNumberOfLinks returns the NumberOfLinks field value if set, zero value otherwise.
func (o *TraceSpanDetail) GetNumberOfLinks() int32 {
	if o == nil || IsNil(o.NumberOfLinks) {
		var ret int32
		return ret
	}
	return *o.NumberOfLinks
}

// GetNumberOfLinksOk returns a tuple with the NumberOfLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceSpanDetail) GetNumberOfLinksOk() (*int32, bool) {
	if o == nil || IsNil(o.NumberOfLinks) {
		return nil, false
	}
	return o.NumberOfLinks, true
}

// HasNumberOfLinks returns a boolean if a field has been set.
func (o *TraceSpanDetail) HasNumberOfLinks() bool {
	if o != nil && !IsNil(o.NumberOfLinks) {
		return true
	}

	return false
}

// SetNumberOfLinks gets a reference to the given int32 and assigns it to the NumberOfLinks field.
func (o *TraceSpanDetail) SetNumberOfLinks(v int32) {
	o.NumberOfLinks = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *TraceSpanDetail) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceSpanDetail) GetErrorMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *TraceSpanDetail) HasErrorMessage() bool {
	if o != nil && !IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *TraceSpanDetail) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *TraceSpanDetail) GetFields() map[string]TracingValue {
	if o == nil || IsNil(o.Fields) {
		var ret map[string]TracingValue
		return ret
	}
	return *o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceSpanDetail) GetFieldsOk() (*map[string]TracingValue, bool) {
	if o == nil || IsNil(o.Fields) {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *TraceSpanDetail) HasFields() bool {
	if o != nil && !IsNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given map[string]TracingValue and assigns it to the Fields field.
func (o *TraceSpanDetail) SetFields(v map[string]TracingValue) {
	o.Fields = &v
}

// GetCriticalPathContribution returns the CriticalPathContribution field value if set, zero value otherwise.
func (o *TraceSpanDetail) GetCriticalPathContribution() TraceSpanCriticalPathContribution {
	if o == nil || IsNil(o.CriticalPathContribution) {
		var ret TraceSpanCriticalPathContribution
		return ret
	}
	return *o.CriticalPathContribution
}

// GetCriticalPathContributionOk returns a tuple with the CriticalPathContribution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceSpanDetail) GetCriticalPathContributionOk() (*TraceSpanCriticalPathContribution, bool) {
	if o == nil || IsNil(o.CriticalPathContribution) {
		return nil, false
	}
	return o.CriticalPathContribution, true
}

// HasCriticalPathContribution returns a boolean if a field has been set.
func (o *TraceSpanDetail) HasCriticalPathContribution() bool {
	if o != nil && !IsNil(o.CriticalPathContribution) {
		return true
	}

	return false
}

// SetCriticalPathContribution gets a reference to the given TraceSpanCriticalPathContribution and assigns it to the CriticalPathContribution field.
func (o *TraceSpanDetail) SetCriticalPathContribution(v TraceSpanCriticalPathContribution) {
	o.CriticalPathContribution = &v
}

// GetLogs returns the Logs field value if set, zero value otherwise.
func (o *TraceSpanDetail) GetLogs() []string {
	if o == nil || IsNil(o.Logs) {
		var ret []string
		return ret
	}
	return o.Logs
}

// GetLogsOk returns a tuple with the Logs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceSpanDetail) GetLogsOk() ([]string, bool) {
	if o == nil || IsNil(o.Logs) {
		return nil, false
	}
	return o.Logs, true
}

// HasLogs returns a boolean if a field has been set.
func (o *TraceSpanDetail) HasLogs() bool {
	if o != nil && !IsNil(o.Logs) {
		return true
	}

	return false
}

// SetLogs gets a reference to the given []string and assigns it to the Logs field.
func (o *TraceSpanDetail) SetLogs(v []string) {
	o.Logs = v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *TraceSpanDetail) GetEvents() []SpanEvent {
	if o == nil || IsNil(o.Events) {
		var ret []SpanEvent
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceSpanDetail) GetEventsOk() ([]SpanEvent, bool) {
	if o == nil || IsNil(o.Events) {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *TraceSpanDetail) HasEvents() bool {
	if o != nil && !IsNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []SpanEvent and assigns it to the Events field.
func (o *TraceSpanDetail) SetEvents(v []SpanEvent) {
	o.Events = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *TraceSpanDetail) GetLinks() []SpanLink {
	if o == nil || IsNil(o.Links) {
		var ret []SpanLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceSpanDetail) GetLinksOk() ([]SpanLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *TraceSpanDetail) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []SpanLink and assigns it to the Links field.
func (o *TraceSpanDetail) SetLinks(v []SpanLink) {
	o.Links = v
}

func (o TraceSpanDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TraceSpanDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.ParentId) {
		toSerialize["parentId"] = o.ParentId
	}
	toSerialize["operationName"] = o.OperationName
	if !IsNil(o.Resource) {
		toSerialize["resource"] = o.Resource
	}
	if !IsNil(o.Service) {
		toSerialize["service"] = o.Service
	}
	if !IsNil(o.ServiceColor) {
		toSerialize["serviceColor"] = o.ServiceColor
	}
	if !IsNil(o.ServiceType) {
		toSerialize["serviceType"] = o.ServiceType
	}
	toSerialize["duration"] = o.Duration
	toSerialize["startedAt"] = o.StartedAt
	toSerialize["status"] = o.Status
	if !IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if !IsNil(o.RemoteService) {
		toSerialize["remoteService"] = o.RemoteService
	}
	if !IsNil(o.RemoteServiceColor) {
		toSerialize["remoteServiceColor"] = o.RemoteServiceColor
	}
	if !IsNil(o.RemoteServiceType) {
		toSerialize["remoteServiceType"] = o.RemoteServiceType
	}
	if !IsNil(o.Info) {
		toSerialize["info"] = o.Info
	}
	if !IsNil(o.NumberOfLinks) {
		toSerialize["numberOfLinks"] = o.NumberOfLinks
	}
	if !IsNil(o.ErrorMessage) {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	if !IsNil(o.Fields) {
		toSerialize["fields"] = o.Fields
	}
	if !IsNil(o.CriticalPathContribution) {
		toSerialize["criticalPathContribution"] = o.CriticalPathContribution
	}
	if !IsNil(o.Logs) {
		toSerialize["logs"] = o.Logs
	}
	if !IsNil(o.Events) {
		toSerialize["events"] = o.Events
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

func (o *TraceSpanDetail) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"operationName",
		"duration",
		"startedAt",
		"status",
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

	varTraceSpanDetail := _TraceSpanDetail{}

	err = json.Unmarshal(bytes, &varTraceSpanDetail)

	if err != nil {
		return err
	}

	*o = TraceSpanDetail(varTraceSpanDetail)

	return err
}

type NullableTraceSpanDetail struct {
	value *TraceSpanDetail
	isSet bool
}

func (v NullableTraceSpanDetail) Get() *TraceSpanDetail {
	return v.value
}

func (v *NullableTraceSpanDetail) Set(val *TraceSpanDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableTraceSpanDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableTraceSpanDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTraceSpanDetail(val *TraceSpanDetail) *NullableTraceSpanDetail {
	return &NullableTraceSpanDetail{value: val, isSet: true}
}

func (v NullableTraceSpanDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTraceSpanDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


