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

// checks if the Extension type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Extension{}

// Extension struct for Extension
type Extension struct {
	// The type property identifies the type of object
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
	CreatedByRef string `json:"created_by_ref"`
	// The revoked property is only used by STIX Objects that support versioning and indicates whether the object has been revoked.
	Revoked *bool `json:"revoked,omitempty"`
	// The labels property specifies a set of terms used to describe this object. The terms are user-defined or trust-group defined and their meaning is outside the scope of this specification and MAY be ignored.
	Labels []string `json:"labels,omitempty"`
	// A list of external references which refer to non-STIX information. This property MAY be used to provide one or more Vulnerability identifiers, such as a CVE ID
	ExternalReferences []ExternalReference `json:"external_references,omitempty"`
	// The object_marking_refs property specifies a list of id properties of marking-definition objects that apply to this object.
	ObjectMarkingRefs []string `json:"object_marking_refs,omitempty"`
	// The granular_markings property specifies a list of granular markings applied to this object
	GranularMarkings []GranularMarkingType `json:"granular_markings,omitempty"`
	// The name of the object
	Name string `json:"name"`
	// A human readable description
	Description *string `json:"description,omitempty"`
	// The normative definition of the extension, either as a URL or as plain text explaining the definition
	Schema string `json:"schema"`
	// The version of this extension
	Version string `json:"version"`
	// This property specifies one or more extension types contained within this extension
	ExtensionTypes []string `json:"extension_types"`
	// This property contains the list of new property names that are added to an object by an extension
	ExtensionProperties []string `json:"extension_properties,omitempty"`
}

type _Extension Extension

// NewExtension instantiates a new Extension object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtension(type_ string, specVersion string, id string, created time.Time, modified time.Time, createdByRef string, name string, schema string, version string, extensionTypes []string) *Extension {
	this := Extension{}
	this.Type = type_
	this.SpecVersion = specVersion
	this.Id = id
	this.Created = created
	this.Modified = modified
	this.CreatedByRef = createdByRef
	this.Name = name
	this.Schema = schema
	this.Version = version
	this.ExtensionTypes = extensionTypes
	return &this
}

// NewExtensionWithDefaults instantiates a new Extension object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtensionWithDefaults() *Extension {
	this := Extension{}
	return &this
}

// GetType returns the Type field value
func (o *Extension) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Extension) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Extension) SetType(v string) {
	o.Type = v
}

// GetSpecVersion returns the SpecVersion field value
func (o *Extension) GetSpecVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SpecVersion
}

// GetSpecVersionOk returns a tuple with the SpecVersion field value
// and a boolean to check if the value has been set.
func (o *Extension) GetSpecVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpecVersion, true
}

// SetSpecVersion sets field value
func (o *Extension) SetSpecVersion(v string) {
	o.SpecVersion = v
}

// GetId returns the Id field value
func (o *Extension) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Extension) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Extension) SetId(v string) {
	o.Id = v
}

// GetCreated returns the Created field value
func (o *Extension) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *Extension) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *Extension) SetCreated(v time.Time) {
	o.Created = v
}

// GetModified returns the Modified field value
func (o *Extension) GetModified() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value
// and a boolean to check if the value has been set.
func (o *Extension) GetModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Modified, true
}

// SetModified sets field value
func (o *Extension) SetModified(v time.Time) {
	o.Modified = v
}

// GetCreatedByRef returns the CreatedByRef field value
func (o *Extension) GetCreatedByRef() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedByRef
}

// GetCreatedByRefOk returns a tuple with the CreatedByRef field value
// and a boolean to check if the value has been set.
func (o *Extension) GetCreatedByRefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedByRef, true
}

// SetCreatedByRef sets field value
func (o *Extension) SetCreatedByRef(v string) {
	o.CreatedByRef = v
}

// GetRevoked returns the Revoked field value if set, zero value otherwise.
func (o *Extension) GetRevoked() bool {
	if o == nil || IsNil(o.Revoked) {
		var ret bool
		return ret
	}
	return *o.Revoked
}

// GetRevokedOk returns a tuple with the Revoked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Extension) GetRevokedOk() (*bool, bool) {
	if o == nil || IsNil(o.Revoked) {
		return nil, false
	}
	return o.Revoked, true
}

// HasRevoked returns a boolean if a field has been set.
func (o *Extension) HasRevoked() bool {
	if o != nil && !IsNil(o.Revoked) {
		return true
	}

	return false
}

// SetRevoked gets a reference to the given bool and assigns it to the Revoked field.
func (o *Extension) SetRevoked(v bool) {
	o.Revoked = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *Extension) GetLabels() []string {
	if o == nil || IsNil(o.Labels) {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Extension) GetLabelsOk() ([]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *Extension) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *Extension) SetLabels(v []string) {
	o.Labels = v
}

// GetExternalReferences returns the ExternalReferences field value if set, zero value otherwise.
func (o *Extension) GetExternalReferences() []ExternalReference {
	if o == nil || IsNil(o.ExternalReferences) {
		var ret []ExternalReference
		return ret
	}
	return o.ExternalReferences
}

// GetExternalReferencesOk returns a tuple with the ExternalReferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Extension) GetExternalReferencesOk() ([]ExternalReference, bool) {
	if o == nil || IsNil(o.ExternalReferences) {
		return nil, false
	}
	return o.ExternalReferences, true
}

// HasExternalReferences returns a boolean if a field has been set.
func (o *Extension) HasExternalReferences() bool {
	if o != nil && !IsNil(o.ExternalReferences) {
		return true
	}

	return false
}

// SetExternalReferences gets a reference to the given []ExternalReference and assigns it to the ExternalReferences field.
func (o *Extension) SetExternalReferences(v []ExternalReference) {
	o.ExternalReferences = v
}

// GetObjectMarkingRefs returns the ObjectMarkingRefs field value if set, zero value otherwise.
func (o *Extension) GetObjectMarkingRefs() []string {
	if o == nil || IsNil(o.ObjectMarkingRefs) {
		var ret []string
		return ret
	}
	return o.ObjectMarkingRefs
}

// GetObjectMarkingRefsOk returns a tuple with the ObjectMarkingRefs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Extension) GetObjectMarkingRefsOk() ([]string, bool) {
	if o == nil || IsNil(o.ObjectMarkingRefs) {
		return nil, false
	}
	return o.ObjectMarkingRefs, true
}

// HasObjectMarkingRefs returns a boolean if a field has been set.
func (o *Extension) HasObjectMarkingRefs() bool {
	if o != nil && !IsNil(o.ObjectMarkingRefs) {
		return true
	}

	return false
}

// SetObjectMarkingRefs gets a reference to the given []string and assigns it to the ObjectMarkingRefs field.
func (o *Extension) SetObjectMarkingRefs(v []string) {
	o.ObjectMarkingRefs = v
}

// GetGranularMarkings returns the GranularMarkings field value if set, zero value otherwise.
func (o *Extension) GetGranularMarkings() []GranularMarkingType {
	if o == nil || IsNil(o.GranularMarkings) {
		var ret []GranularMarkingType
		return ret
	}
	return o.GranularMarkings
}

// GetGranularMarkingsOk returns a tuple with the GranularMarkings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Extension) GetGranularMarkingsOk() ([]GranularMarkingType, bool) {
	if o == nil || IsNil(o.GranularMarkings) {
		return nil, false
	}
	return o.GranularMarkings, true
}

// HasGranularMarkings returns a boolean if a field has been set.
func (o *Extension) HasGranularMarkings() bool {
	if o != nil && !IsNil(o.GranularMarkings) {
		return true
	}

	return false
}

// SetGranularMarkings gets a reference to the given []GranularMarkingType and assigns it to the GranularMarkings field.
func (o *Extension) SetGranularMarkings(v []GranularMarkingType) {
	o.GranularMarkings = v
}

// GetName returns the Name field value
func (o *Extension) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Extension) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Extension) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Extension) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Extension) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Extension) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Extension) SetDescription(v string) {
	o.Description = &v
}

// GetSchema returns the Schema field value
func (o *Extension) GetSchema() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value
// and a boolean to check if the value has been set.
func (o *Extension) GetSchemaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Schema, true
}

// SetSchema sets field value
func (o *Extension) SetSchema(v string) {
	o.Schema = v
}

// GetVersion returns the Version field value
func (o *Extension) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *Extension) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *Extension) SetVersion(v string) {
	o.Version = v
}

// GetExtensionTypes returns the ExtensionTypes field value
func (o *Extension) GetExtensionTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ExtensionTypes
}

// GetExtensionTypesOk returns a tuple with the ExtensionTypes field value
// and a boolean to check if the value has been set.
func (o *Extension) GetExtensionTypesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExtensionTypes, true
}

// SetExtensionTypes sets field value
func (o *Extension) SetExtensionTypes(v []string) {
	o.ExtensionTypes = v
}

// GetExtensionProperties returns the ExtensionProperties field value if set, zero value otherwise.
func (o *Extension) GetExtensionProperties() []string {
	if o == nil || IsNil(o.ExtensionProperties) {
		var ret []string
		return ret
	}
	return o.ExtensionProperties
}

// GetExtensionPropertiesOk returns a tuple with the ExtensionProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Extension) GetExtensionPropertiesOk() ([]string, bool) {
	if o == nil || IsNil(o.ExtensionProperties) {
		return nil, false
	}
	return o.ExtensionProperties, true
}

// HasExtensionProperties returns a boolean if a field has been set.
func (o *Extension) HasExtensionProperties() bool {
	if o != nil && !IsNil(o.ExtensionProperties) {
		return true
	}

	return false
}

// SetExtensionProperties gets a reference to the given []string and assigns it to the ExtensionProperties field.
func (o *Extension) SetExtensionProperties(v []string) {
	o.ExtensionProperties = v
}

func (o Extension) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Extension) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["spec_version"] = o.SpecVersion
	toSerialize["id"] = o.Id
	toSerialize["created"] = o.Created
	toSerialize["modified"] = o.Modified
	toSerialize["created_by_ref"] = o.CreatedByRef
	if !IsNil(o.Revoked) {
		toSerialize["revoked"] = o.Revoked
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
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
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["schema"] = o.Schema
	toSerialize["version"] = o.Version
	toSerialize["extension_types"] = o.ExtensionTypes
	if !IsNil(o.ExtensionProperties) {
		toSerialize["extension_properties"] = o.ExtensionProperties
	}
	return toSerialize, nil
}

func (o *Extension) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"spec_version",
		"id",
		"created",
		"modified",
		"created_by_ref",
		"name",
		"schema",
		"version",
		"extension_types",
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

	varExtension := _Extension{}

	err = json.Unmarshal(bytes, &varExtension)

	if err != nil {
		return err
	}

	*o = Extension(varExtension)

	return err
}

type NullableExtension struct {
	value *Extension
	isSet bool
}

func (v NullableExtension) Get() *Extension {
	return v.value
}

func (v *NullableExtension) Set(val *Extension) {
	v.value = val
	v.isSet = true
}

func (v NullableExtension) IsSet() bool {
	return v.isSet
}

func (v *NullableExtension) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtension(val *Extension) *NullableExtension {
	return &NullableExtension{value: val, isSet: true}
}

func (v NullableExtension) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtension) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


