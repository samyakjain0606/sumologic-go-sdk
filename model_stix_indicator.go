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

// checks if the StixIndicator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StixIndicator{}

// StixIndicator struct for StixIndicator
type StixIndicator struct {
	// The type property identifies the type of STIX Object.
	Type string `json:"type"`
	// The STIX version
	SpecVersion string `json:"spec_version"`
	// The ID of the indicator
	Id string `json:"id"`
	// The time from which this Indicator is considered a valid indicator of the behaviors it is related or represents.
	Created time.Time `json:"created"`
	// The time from which this Indicator is considered a valid indicator of the behaviors it is related or represents.
	Modified time.Time `json:"modified"`
	// Identifier of type identity
	CreatedByRef *string `json:"created_by_ref,omitempty"`
	// The revoked property is only used by STIX Objects that support versioning and indicates whether the object has been revoked.
	Revoked *bool `json:"revoked,omitempty"`
	// The labels property specifies a set of terms used to describe this object. The terms are user-defined or trust-group defined and their meaning is outside the scope of this specification and MAY be ignored.
	Labels []string `json:"labels,omitempty"`
	// Confidence that the creator has in the correctness of their data, where 100 is highest
	Confidence *int32 `json:"confidence,omitempty"`
	// The lang property identifies the language of the text content in this object. When present, it MUST be a language code conformant to [RFC5646]. If the property is not present, then the language of the content is en (English)
	Lang *string `json:"lang,omitempty"`
	// A list of external references which refer to non-STIX information. This property MAY be used to provide one or more Vulnerability identifiers, such as a CVE ID
	ExternalReferences []ExternalReference `json:"external_references,omitempty"`
	// The object_marking_refs property specifies a list of id properties of marking-definition objects that apply to this object.
	ObjectMarkingRefs []string `json:"object_marking_refs,omitempty"`
	// The granular_markings property specifies a list of granular markings applied to this object
	GranularMarkings []GranularMarkingType `json:"granular_markings,omitempty"`
	// Specifies any extensions of the object, as a dictionary
	Extensions *map[string]Extension `json:"extensions,omitempty"`
	// The name of the object
	Name *string `json:"name,omitempty"`
	// A human readable description
	Description *string `json:"description,omitempty"`
	// A set of categorizations for this indicator.
	IndicatorTypes []string `json:"indicator_types,omitempty"`
	// The detection pattern for this Indicator expressed as a STIX patter.
	Pattern string `json:"pattern"`
	// The type of pattern
	PatternType string `json:"pattern_type"`
	// The version of the pattern language that is used for the data in the pattern property which MUST match the type of pattern data included in the pattern property.
	PatternVersion *string `json:"pattern_version,omitempty"`
	// The time from which this Indicator is considered a valid indicator of the behaviors it is related or represents.
	ValidFrom time.Time `json:"valid_from"`
	// The time at which this Indicator should no longer be considered a valid indicator of the behaviors it is related to or represents.
	ValidUntil *time.Time `json:"valid_until,omitempty"`
	// The list of Kill Chain Phases for which this Attack Pattern is used
	KillChainPhases []KillChainPhase `json:"kill_chain_phases,omitempty"`
}

type _StixIndicator StixIndicator

// NewStixIndicator instantiates a new StixIndicator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStixIndicator(type_ string, specVersion string, id string, created time.Time, modified time.Time, pattern string, patternType string, validFrom time.Time) *StixIndicator {
	this := StixIndicator{}
	this.Type = type_
	this.SpecVersion = specVersion
	this.Id = id
	this.Created = created
	this.Modified = modified
	this.Pattern = pattern
	this.PatternType = patternType
	this.ValidFrom = validFrom
	return &this
}

// NewStixIndicatorWithDefaults instantiates a new StixIndicator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStixIndicatorWithDefaults() *StixIndicator {
	this := StixIndicator{}
	return &this
}

// GetType returns the Type field value
func (o *StixIndicator) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *StixIndicator) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *StixIndicator) SetType(v string) {
	o.Type = v
}

// GetSpecVersion returns the SpecVersion field value
func (o *StixIndicator) GetSpecVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SpecVersion
}

// GetSpecVersionOk returns a tuple with the SpecVersion field value
// and a boolean to check if the value has been set.
func (o *StixIndicator) GetSpecVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpecVersion, true
}

// SetSpecVersion sets field value
func (o *StixIndicator) SetSpecVersion(v string) {
	o.SpecVersion = v
}

// GetId returns the Id field value
func (o *StixIndicator) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *StixIndicator) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *StixIndicator) SetId(v string) {
	o.Id = v
}

// GetCreated returns the Created field value
func (o *StixIndicator) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *StixIndicator) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *StixIndicator) SetCreated(v time.Time) {
	o.Created = v
}

// GetModified returns the Modified field value
func (o *StixIndicator) GetModified() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value
// and a boolean to check if the value has been set.
func (o *StixIndicator) GetModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Modified, true
}

// SetModified sets field value
func (o *StixIndicator) SetModified(v time.Time) {
	o.Modified = v
}

// GetCreatedByRef returns the CreatedByRef field value if set, zero value otherwise.
func (o *StixIndicator) GetCreatedByRef() string {
	if o == nil || IsNil(o.CreatedByRef) {
		var ret string
		return ret
	}
	return *o.CreatedByRef
}

// GetCreatedByRefOk returns a tuple with the CreatedByRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StixIndicator) GetCreatedByRefOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedByRef) {
		return nil, false
	}
	return o.CreatedByRef, true
}

// HasCreatedByRef returns a boolean if a field has been set.
func (o *StixIndicator) HasCreatedByRef() bool {
	if o != nil && !IsNil(o.CreatedByRef) {
		return true
	}

	return false
}

// SetCreatedByRef gets a reference to the given string and assigns it to the CreatedByRef field.
func (o *StixIndicator) SetCreatedByRef(v string) {
	o.CreatedByRef = &v
}

// GetRevoked returns the Revoked field value if set, zero value otherwise.
func (o *StixIndicator) GetRevoked() bool {
	if o == nil || IsNil(o.Revoked) {
		var ret bool
		return ret
	}
	return *o.Revoked
}

// GetRevokedOk returns a tuple with the Revoked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StixIndicator) GetRevokedOk() (*bool, bool) {
	if o == nil || IsNil(o.Revoked) {
		return nil, false
	}
	return o.Revoked, true
}

// HasRevoked returns a boolean if a field has been set.
func (o *StixIndicator) HasRevoked() bool {
	if o != nil && !IsNil(o.Revoked) {
		return true
	}

	return false
}

// SetRevoked gets a reference to the given bool and assigns it to the Revoked field.
func (o *StixIndicator) SetRevoked(v bool) {
	o.Revoked = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *StixIndicator) GetLabels() []string {
	if o == nil || IsNil(o.Labels) {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StixIndicator) GetLabelsOk() ([]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *StixIndicator) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *StixIndicator) SetLabels(v []string) {
	o.Labels = v
}

// GetConfidence returns the Confidence field value if set, zero value otherwise.
func (o *StixIndicator) GetConfidence() int32 {
	if o == nil || IsNil(o.Confidence) {
		var ret int32
		return ret
	}
	return *o.Confidence
}

// GetConfidenceOk returns a tuple with the Confidence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StixIndicator) GetConfidenceOk() (*int32, bool) {
	if o == nil || IsNil(o.Confidence) {
		return nil, false
	}
	return o.Confidence, true
}

// HasConfidence returns a boolean if a field has been set.
func (o *StixIndicator) HasConfidence() bool {
	if o != nil && !IsNil(o.Confidence) {
		return true
	}

	return false
}

// SetConfidence gets a reference to the given int32 and assigns it to the Confidence field.
func (o *StixIndicator) SetConfidence(v int32) {
	o.Confidence = &v
}

// GetLang returns the Lang field value if set, zero value otherwise.
func (o *StixIndicator) GetLang() string {
	if o == nil || IsNil(o.Lang) {
		var ret string
		return ret
	}
	return *o.Lang
}

// GetLangOk returns a tuple with the Lang field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StixIndicator) GetLangOk() (*string, bool) {
	if o == nil || IsNil(o.Lang) {
		return nil, false
	}
	return o.Lang, true
}

// HasLang returns a boolean if a field has been set.
func (o *StixIndicator) HasLang() bool {
	if o != nil && !IsNil(o.Lang) {
		return true
	}

	return false
}

// SetLang gets a reference to the given string and assigns it to the Lang field.
func (o *StixIndicator) SetLang(v string) {
	o.Lang = &v
}

// GetExternalReferences returns the ExternalReferences field value if set, zero value otherwise.
func (o *StixIndicator) GetExternalReferences() []ExternalReference {
	if o == nil || IsNil(o.ExternalReferences) {
		var ret []ExternalReference
		return ret
	}
	return o.ExternalReferences
}

// GetExternalReferencesOk returns a tuple with the ExternalReferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StixIndicator) GetExternalReferencesOk() ([]ExternalReference, bool) {
	if o == nil || IsNil(o.ExternalReferences) {
		return nil, false
	}
	return o.ExternalReferences, true
}

// HasExternalReferences returns a boolean if a field has been set.
func (o *StixIndicator) HasExternalReferences() bool {
	if o != nil && !IsNil(o.ExternalReferences) {
		return true
	}

	return false
}

// SetExternalReferences gets a reference to the given []ExternalReference and assigns it to the ExternalReferences field.
func (o *StixIndicator) SetExternalReferences(v []ExternalReference) {
	o.ExternalReferences = v
}

// GetObjectMarkingRefs returns the ObjectMarkingRefs field value if set, zero value otherwise.
func (o *StixIndicator) GetObjectMarkingRefs() []string {
	if o == nil || IsNil(o.ObjectMarkingRefs) {
		var ret []string
		return ret
	}
	return o.ObjectMarkingRefs
}

// GetObjectMarkingRefsOk returns a tuple with the ObjectMarkingRefs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StixIndicator) GetObjectMarkingRefsOk() ([]string, bool) {
	if o == nil || IsNil(o.ObjectMarkingRefs) {
		return nil, false
	}
	return o.ObjectMarkingRefs, true
}

// HasObjectMarkingRefs returns a boolean if a field has been set.
func (o *StixIndicator) HasObjectMarkingRefs() bool {
	if o != nil && !IsNil(o.ObjectMarkingRefs) {
		return true
	}

	return false
}

// SetObjectMarkingRefs gets a reference to the given []string and assigns it to the ObjectMarkingRefs field.
func (o *StixIndicator) SetObjectMarkingRefs(v []string) {
	o.ObjectMarkingRefs = v
}

// GetGranularMarkings returns the GranularMarkings field value if set, zero value otherwise.
func (o *StixIndicator) GetGranularMarkings() []GranularMarkingType {
	if o == nil || IsNil(o.GranularMarkings) {
		var ret []GranularMarkingType
		return ret
	}
	return o.GranularMarkings
}

// GetGranularMarkingsOk returns a tuple with the GranularMarkings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StixIndicator) GetGranularMarkingsOk() ([]GranularMarkingType, bool) {
	if o == nil || IsNil(o.GranularMarkings) {
		return nil, false
	}
	return o.GranularMarkings, true
}

// HasGranularMarkings returns a boolean if a field has been set.
func (o *StixIndicator) HasGranularMarkings() bool {
	if o != nil && !IsNil(o.GranularMarkings) {
		return true
	}

	return false
}

// SetGranularMarkings gets a reference to the given []GranularMarkingType and assigns it to the GranularMarkings field.
func (o *StixIndicator) SetGranularMarkings(v []GranularMarkingType) {
	o.GranularMarkings = v
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *StixIndicator) GetExtensions() map[string]Extension {
	if o == nil || IsNil(o.Extensions) {
		var ret map[string]Extension
		return ret
	}
	return *o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StixIndicator) GetExtensionsOk() (*map[string]Extension, bool) {
	if o == nil || IsNil(o.Extensions) {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *StixIndicator) HasExtensions() bool {
	if o != nil && !IsNil(o.Extensions) {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given map[string]Extension and assigns it to the Extensions field.
func (o *StixIndicator) SetExtensions(v map[string]Extension) {
	o.Extensions = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StixIndicator) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StixIndicator) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StixIndicator) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StixIndicator) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *StixIndicator) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StixIndicator) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *StixIndicator) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *StixIndicator) SetDescription(v string) {
	o.Description = &v
}

// GetIndicatorTypes returns the IndicatorTypes field value if set, zero value otherwise.
func (o *StixIndicator) GetIndicatorTypes() []string {
	if o == nil || IsNil(o.IndicatorTypes) {
		var ret []string
		return ret
	}
	return o.IndicatorTypes
}

// GetIndicatorTypesOk returns a tuple with the IndicatorTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StixIndicator) GetIndicatorTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.IndicatorTypes) {
		return nil, false
	}
	return o.IndicatorTypes, true
}

// HasIndicatorTypes returns a boolean if a field has been set.
func (o *StixIndicator) HasIndicatorTypes() bool {
	if o != nil && !IsNil(o.IndicatorTypes) {
		return true
	}

	return false
}

// SetIndicatorTypes gets a reference to the given []string and assigns it to the IndicatorTypes field.
func (o *StixIndicator) SetIndicatorTypes(v []string) {
	o.IndicatorTypes = v
}

// GetPattern returns the Pattern field value
func (o *StixIndicator) GetPattern() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pattern
}

// GetPatternOk returns a tuple with the Pattern field value
// and a boolean to check if the value has been set.
func (o *StixIndicator) GetPatternOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pattern, true
}

// SetPattern sets field value
func (o *StixIndicator) SetPattern(v string) {
	o.Pattern = v
}

// GetPatternType returns the PatternType field value
func (o *StixIndicator) GetPatternType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PatternType
}

// GetPatternTypeOk returns a tuple with the PatternType field value
// and a boolean to check if the value has been set.
func (o *StixIndicator) GetPatternTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PatternType, true
}

// SetPatternType sets field value
func (o *StixIndicator) SetPatternType(v string) {
	o.PatternType = v
}

// GetPatternVersion returns the PatternVersion field value if set, zero value otherwise.
func (o *StixIndicator) GetPatternVersion() string {
	if o == nil || IsNil(o.PatternVersion) {
		var ret string
		return ret
	}
	return *o.PatternVersion
}

// GetPatternVersionOk returns a tuple with the PatternVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StixIndicator) GetPatternVersionOk() (*string, bool) {
	if o == nil || IsNil(o.PatternVersion) {
		return nil, false
	}
	return o.PatternVersion, true
}

// HasPatternVersion returns a boolean if a field has been set.
func (o *StixIndicator) HasPatternVersion() bool {
	if o != nil && !IsNil(o.PatternVersion) {
		return true
	}

	return false
}

// SetPatternVersion gets a reference to the given string and assigns it to the PatternVersion field.
func (o *StixIndicator) SetPatternVersion(v string) {
	o.PatternVersion = &v
}

// GetValidFrom returns the ValidFrom field value
func (o *StixIndicator) GetValidFrom() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ValidFrom
}

// GetValidFromOk returns a tuple with the ValidFrom field value
// and a boolean to check if the value has been set.
func (o *StixIndicator) GetValidFromOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValidFrom, true
}

// SetValidFrom sets field value
func (o *StixIndicator) SetValidFrom(v time.Time) {
	o.ValidFrom = v
}

// GetValidUntil returns the ValidUntil field value if set, zero value otherwise.
func (o *StixIndicator) GetValidUntil() time.Time {
	if o == nil || IsNil(o.ValidUntil) {
		var ret time.Time
		return ret
	}
	return *o.ValidUntil
}

// GetValidUntilOk returns a tuple with the ValidUntil field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StixIndicator) GetValidUntilOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ValidUntil) {
		return nil, false
	}
	return o.ValidUntil, true
}

// HasValidUntil returns a boolean if a field has been set.
func (o *StixIndicator) HasValidUntil() bool {
	if o != nil && !IsNil(o.ValidUntil) {
		return true
	}

	return false
}

// SetValidUntil gets a reference to the given time.Time and assigns it to the ValidUntil field.
func (o *StixIndicator) SetValidUntil(v time.Time) {
	o.ValidUntil = &v
}

// GetKillChainPhases returns the KillChainPhases field value if set, zero value otherwise.
func (o *StixIndicator) GetKillChainPhases() []KillChainPhase {
	if o == nil || IsNil(o.KillChainPhases) {
		var ret []KillChainPhase
		return ret
	}
	return o.KillChainPhases
}

// GetKillChainPhasesOk returns a tuple with the KillChainPhases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StixIndicator) GetKillChainPhasesOk() ([]KillChainPhase, bool) {
	if o == nil || IsNil(o.KillChainPhases) {
		return nil, false
	}
	return o.KillChainPhases, true
}

// HasKillChainPhases returns a boolean if a field has been set.
func (o *StixIndicator) HasKillChainPhases() bool {
	if o != nil && !IsNil(o.KillChainPhases) {
		return true
	}

	return false
}

// SetKillChainPhases gets a reference to the given []KillChainPhase and assigns it to the KillChainPhases field.
func (o *StixIndicator) SetKillChainPhases(v []KillChainPhase) {
	o.KillChainPhases = v
}

func (o StixIndicator) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StixIndicator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["spec_version"] = o.SpecVersion
	toSerialize["id"] = o.Id
	toSerialize["created"] = o.Created
	toSerialize["modified"] = o.Modified
	if !IsNil(o.CreatedByRef) {
		toSerialize["created_by_ref"] = o.CreatedByRef
	}
	if !IsNil(o.Revoked) {
		toSerialize["revoked"] = o.Revoked
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.Confidence) {
		toSerialize["confidence"] = o.Confidence
	}
	if !IsNil(o.Lang) {
		toSerialize["lang"] = o.Lang
	}
	if !IsNil(o.ExternalReferences) {
		toSerialize["external_references"] = o.ExternalReferences
	}
	if !IsNil(o.ObjectMarkingRefs) {
		toSerialize["object_marking_refs"] = o.ObjectMarkingRefs
	}
	if !IsNil(o.GranularMarkings) {
		toSerialize["granular_markings"] = o.GranularMarkings
	}
	if !IsNil(o.Extensions) {
		toSerialize["extensions"] = o.Extensions
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.IndicatorTypes) {
		toSerialize["indicator_types"] = o.IndicatorTypes
	}
	toSerialize["pattern"] = o.Pattern
	toSerialize["pattern_type"] = o.PatternType
	if !IsNil(o.PatternVersion) {
		toSerialize["pattern_version"] = o.PatternVersion
	}
	toSerialize["valid_from"] = o.ValidFrom
	if !IsNil(o.ValidUntil) {
		toSerialize["valid_until"] = o.ValidUntil
	}
	if !IsNil(o.KillChainPhases) {
		toSerialize["kill_chain_phases"] = o.KillChainPhases
	}
	return toSerialize, nil
}

func (o *StixIndicator) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"spec_version",
		"id",
		"created",
		"modified",
		"pattern",
		"pattern_type",
		"valid_from",
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

	varStixIndicator := _StixIndicator{}

	err = json.Unmarshal(bytes, &varStixIndicator)

	if err != nil {
		return err
	}

	*o = StixIndicator(varStixIndicator)

	return err
}

type NullableStixIndicator struct {
	value *StixIndicator
	isSet bool
}

func (v NullableStixIndicator) Get() *StixIndicator {
	return v.value
}

func (v *NullableStixIndicator) Set(val *StixIndicator) {
	v.value = val
	v.isSet = true
}

func (v NullableStixIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullableStixIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStixIndicator(val *StixIndicator) *NullableStixIndicator {
	return &NullableStixIndicator{value: val, isSet: true}
}

func (v NullableStixIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStixIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


