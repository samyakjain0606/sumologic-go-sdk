# Extension

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type property identifies the type of object | 
**SpecVersion** | **string** | The STIX version | 
**Id** | **string** | The ID of the indicator | 
**Created** | **time.Time** | The time from which this Indicator is considered a valid indicator of the behaviors it is related or represents. | 
**Modified** | **time.Time** | The time from which this Indicator is considered a valid indicator of the behaviors it is related or represents. | 
**CreatedByRef** | **string** | Identifier of type identity | 
**Revoked** | Pointer to **bool** | The revoked property is only used by STIX Objects that support versioning and indicates whether the object has been revoked. | [optional] 
**Labels** | Pointer to **[]string** | The labels property specifies a set of terms used to describe this object. The terms are user-defined or trust-group defined and their meaning is outside the scope of this specification and MAY be ignored. | [optional] 
**ExternalReferences** | Pointer to [**[]ExternalReference**](ExternalReference.md) | A list of external references which refer to non-STIX information. This property MAY be used to provide one or more Vulnerability identifiers, such as a CVE ID | [optional] 
**ObjectMarkingRefs** | Pointer to **[]string** | The object_marking_refs property specifies a list of id properties of marking-definition objects that apply to this object. | [optional] 
**GranularMarkings** | Pointer to [**[]GranularMarkingType**](GranularMarkingType.md) | The granular_markings property specifies a list of granular markings applied to this object | [optional] 
**Name** | **string** | The name of the object | 
**Description** | Pointer to **string** | A human readable description | [optional] 
**Schema** | **string** | The normative definition of the extension, either as a URL or as plain text explaining the definition | 
**Version** | **string** | The version of this extension | 
**ExtensionTypes** | **[]string** | This property specifies one or more extension types contained within this extension | 
**ExtensionProperties** | Pointer to **[]string** | This property contains the list of new property names that are added to an object by an extension | [optional] 

## Methods

### NewExtension

`func NewExtension(type_ string, specVersion string, id string, created time.Time, modified time.Time, createdByRef string, name string, schema string, version string, extensionTypes []string, ) *Extension`

NewExtension instantiates a new Extension object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionWithDefaults

`func NewExtensionWithDefaults() *Extension`

NewExtensionWithDefaults instantiates a new Extension object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Extension) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Extension) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Extension) SetType(v string)`

SetType sets Type field to given value.


### GetSpecVersion

`func (o *Extension) GetSpecVersion() string`

GetSpecVersion returns the SpecVersion field if non-nil, zero value otherwise.

### GetSpecVersionOk

`func (o *Extension) GetSpecVersionOk() (*string, bool)`

GetSpecVersionOk returns a tuple with the SpecVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecVersion

`func (o *Extension) SetSpecVersion(v string)`

SetSpecVersion sets SpecVersion field to given value.


### GetId

`func (o *Extension) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Extension) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Extension) SetId(v string)`

SetId sets Id field to given value.


### GetCreated

`func (o *Extension) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Extension) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Extension) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetModified

`func (o *Extension) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *Extension) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *Extension) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetCreatedByRef

`func (o *Extension) GetCreatedByRef() string`

GetCreatedByRef returns the CreatedByRef field if non-nil, zero value otherwise.

### GetCreatedByRefOk

`func (o *Extension) GetCreatedByRefOk() (*string, bool)`

GetCreatedByRefOk returns a tuple with the CreatedByRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByRef

`func (o *Extension) SetCreatedByRef(v string)`

SetCreatedByRef sets CreatedByRef field to given value.


### GetRevoked

`func (o *Extension) GetRevoked() bool`

GetRevoked returns the Revoked field if non-nil, zero value otherwise.

### GetRevokedOk

`func (o *Extension) GetRevokedOk() (*bool, bool)`

GetRevokedOk returns a tuple with the Revoked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevoked

`func (o *Extension) SetRevoked(v bool)`

SetRevoked sets Revoked field to given value.

### HasRevoked

`func (o *Extension) HasRevoked() bool`

HasRevoked returns a boolean if a field has been set.

### GetLabels

`func (o *Extension) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *Extension) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *Extension) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *Extension) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetExternalReferences

`func (o *Extension) GetExternalReferences() []ExternalReference`

GetExternalReferences returns the ExternalReferences field if non-nil, zero value otherwise.

### GetExternalReferencesOk

`func (o *Extension) GetExternalReferencesOk() (*[]ExternalReference, bool)`

GetExternalReferencesOk returns a tuple with the ExternalReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalReferences

`func (o *Extension) SetExternalReferences(v []ExternalReference)`

SetExternalReferences sets ExternalReferences field to given value.

### HasExternalReferences

`func (o *Extension) HasExternalReferences() bool`

HasExternalReferences returns a boolean if a field has been set.

### GetObjectMarkingRefs

`func (o *Extension) GetObjectMarkingRefs() []string`

GetObjectMarkingRefs returns the ObjectMarkingRefs field if non-nil, zero value otherwise.

### GetObjectMarkingRefsOk

`func (o *Extension) GetObjectMarkingRefsOk() (*[]string, bool)`

GetObjectMarkingRefsOk returns a tuple with the ObjectMarkingRefs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectMarkingRefs

`func (o *Extension) SetObjectMarkingRefs(v []string)`

SetObjectMarkingRefs sets ObjectMarkingRefs field to given value.

### HasObjectMarkingRefs

`func (o *Extension) HasObjectMarkingRefs() bool`

HasObjectMarkingRefs returns a boolean if a field has been set.

### GetGranularMarkings

`func (o *Extension) GetGranularMarkings() []GranularMarkingType`

GetGranularMarkings returns the GranularMarkings field if non-nil, zero value otherwise.

### GetGranularMarkingsOk

`func (o *Extension) GetGranularMarkingsOk() (*[]GranularMarkingType, bool)`

GetGranularMarkingsOk returns a tuple with the GranularMarkings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularMarkings

`func (o *Extension) SetGranularMarkings(v []GranularMarkingType)`

SetGranularMarkings sets GranularMarkings field to given value.

### HasGranularMarkings

`func (o *Extension) HasGranularMarkings() bool`

HasGranularMarkings returns a boolean if a field has been set.

### GetName

`func (o *Extension) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Extension) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Extension) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Extension) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Extension) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Extension) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Extension) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSchema

`func (o *Extension) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *Extension) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *Extension) SetSchema(v string)`

SetSchema sets Schema field to given value.


### GetVersion

`func (o *Extension) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Extension) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Extension) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetExtensionTypes

`func (o *Extension) GetExtensionTypes() []string`

GetExtensionTypes returns the ExtensionTypes field if non-nil, zero value otherwise.

### GetExtensionTypesOk

`func (o *Extension) GetExtensionTypesOk() (*[]string, bool)`

GetExtensionTypesOk returns a tuple with the ExtensionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionTypes

`func (o *Extension) SetExtensionTypes(v []string)`

SetExtensionTypes sets ExtensionTypes field to given value.


### GetExtensionProperties

`func (o *Extension) GetExtensionProperties() []string`

GetExtensionProperties returns the ExtensionProperties field if non-nil, zero value otherwise.

### GetExtensionPropertiesOk

`func (o *Extension) GetExtensionPropertiesOk() (*[]string, bool)`

GetExtensionPropertiesOk returns a tuple with the ExtensionProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionProperties

`func (o *Extension) SetExtensionProperties(v []string)`

SetExtensionProperties sets ExtensionProperties field to given value.

### HasExtensionProperties

`func (o *Extension) HasExtensionProperties() bool`

HasExtensionProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


