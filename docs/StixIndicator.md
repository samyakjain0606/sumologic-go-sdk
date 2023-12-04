# StixIndicator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type property identifies the type of STIX Object. | 
**SpecVersion** | **string** | The STIX version | 
**Id** | **string** | The ID of the indicator | 
**Created** | **time.Time** | The time from which this Indicator is considered a valid indicator of the behaviors it is related or represents. | 
**Modified** | **time.Time** | The time from which this Indicator is considered a valid indicator of the behaviors it is related or represents. | 
**CreatedByRef** | Pointer to **string** | Identifier of type identity | [optional] 
**Revoked** | Pointer to **bool** | The revoked property is only used by STIX Objects that support versioning and indicates whether the object has been revoked. | [optional] 
**Labels** | Pointer to **[]string** | The labels property specifies a set of terms used to describe this object. The terms are user-defined or trust-group defined and their meaning is outside the scope of this specification and MAY be ignored. | [optional] 
**Confidence** | Pointer to **int32** | Confidence that the creator has in the correctness of their data, where 100 is highest | [optional] 
**Lang** | Pointer to **string** | The lang property identifies the language of the text content in this object. When present, it MUST be a language code conformant to [RFC5646]. If the property is not present, then the language of the content is en (English) | [optional] 
**ExternalReferences** | Pointer to [**[]ExternalReference**](ExternalReference.md) | A list of external references which refer to non-STIX information. This property MAY be used to provide one or more Vulnerability identifiers, such as a CVE ID | [optional] 
**ObjectMarkingRefs** | Pointer to **[]string** | The object_marking_refs property specifies a list of id properties of marking-definition objects that apply to this object. | [optional] 
**GranularMarkings** | Pointer to [**[]GranularMarkingType**](GranularMarkingType.md) | The granular_markings property specifies a list of granular markings applied to this object | [optional] 
**Extensions** | Pointer to [**map[string]Extension**](Extension.md) | Specifies any extensions of the object, as a dictionary | [optional] 
**Name** | Pointer to **string** | The name of the object | [optional] 
**Description** | Pointer to **string** | A human readable description | [optional] 
**IndicatorTypes** | Pointer to **[]string** | A set of categorizations for this indicator. | [optional] 
**Pattern** | **string** | The detection pattern for this Indicator expressed as a STIX patter. | 
**PatternType** | **string** | The type of pattern | 
**PatternVersion** | Pointer to **string** | The version of the pattern language that is used for the data in the pattern property which MUST match the type of pattern data included in the pattern property. | [optional] 
**ValidFrom** | **time.Time** | The time from which this Indicator is considered a valid indicator of the behaviors it is related or represents. | 
**ValidUntil** | Pointer to **time.Time** | The time at which this Indicator should no longer be considered a valid indicator of the behaviors it is related to or represents. | [optional] 
**KillChainPhases** | Pointer to [**[]KillChainPhase**](KillChainPhase.md) | The list of Kill Chain Phases for which this Attack Pattern is used | [optional] 

## Methods

### NewStixIndicator

`func NewStixIndicator(type_ string, specVersion string, id string, created time.Time, modified time.Time, pattern string, patternType string, validFrom time.Time, ) *StixIndicator`

NewStixIndicator instantiates a new StixIndicator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStixIndicatorWithDefaults

`func NewStixIndicatorWithDefaults() *StixIndicator`

NewStixIndicatorWithDefaults instantiates a new StixIndicator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *StixIndicator) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StixIndicator) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StixIndicator) SetType(v string)`

SetType sets Type field to given value.


### GetSpecVersion

`func (o *StixIndicator) GetSpecVersion() string`

GetSpecVersion returns the SpecVersion field if non-nil, zero value otherwise.

### GetSpecVersionOk

`func (o *StixIndicator) GetSpecVersionOk() (*string, bool)`

GetSpecVersionOk returns a tuple with the SpecVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecVersion

`func (o *StixIndicator) SetSpecVersion(v string)`

SetSpecVersion sets SpecVersion field to given value.


### GetId

`func (o *StixIndicator) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StixIndicator) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StixIndicator) SetId(v string)`

SetId sets Id field to given value.


### GetCreated

`func (o *StixIndicator) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *StixIndicator) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *StixIndicator) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetModified

`func (o *StixIndicator) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *StixIndicator) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *StixIndicator) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetCreatedByRef

`func (o *StixIndicator) GetCreatedByRef() string`

GetCreatedByRef returns the CreatedByRef field if non-nil, zero value otherwise.

### GetCreatedByRefOk

`func (o *StixIndicator) GetCreatedByRefOk() (*string, bool)`

GetCreatedByRefOk returns a tuple with the CreatedByRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByRef

`func (o *StixIndicator) SetCreatedByRef(v string)`

SetCreatedByRef sets CreatedByRef field to given value.

### HasCreatedByRef

`func (o *StixIndicator) HasCreatedByRef() bool`

HasCreatedByRef returns a boolean if a field has been set.

### GetRevoked

`func (o *StixIndicator) GetRevoked() bool`

GetRevoked returns the Revoked field if non-nil, zero value otherwise.

### GetRevokedOk

`func (o *StixIndicator) GetRevokedOk() (*bool, bool)`

GetRevokedOk returns a tuple with the Revoked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevoked

`func (o *StixIndicator) SetRevoked(v bool)`

SetRevoked sets Revoked field to given value.

### HasRevoked

`func (o *StixIndicator) HasRevoked() bool`

HasRevoked returns a boolean if a field has been set.

### GetLabels

`func (o *StixIndicator) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *StixIndicator) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *StixIndicator) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *StixIndicator) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetConfidence

`func (o *StixIndicator) GetConfidence() int32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *StixIndicator) GetConfidenceOk() (*int32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *StixIndicator) SetConfidence(v int32)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *StixIndicator) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetLang

`func (o *StixIndicator) GetLang() string`

GetLang returns the Lang field if non-nil, zero value otherwise.

### GetLangOk

`func (o *StixIndicator) GetLangOk() (*string, bool)`

GetLangOk returns a tuple with the Lang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLang

`func (o *StixIndicator) SetLang(v string)`

SetLang sets Lang field to given value.

### HasLang

`func (o *StixIndicator) HasLang() bool`

HasLang returns a boolean if a field has been set.

### GetExternalReferences

`func (o *StixIndicator) GetExternalReferences() []ExternalReference`

GetExternalReferences returns the ExternalReferences field if non-nil, zero value otherwise.

### GetExternalReferencesOk

`func (o *StixIndicator) GetExternalReferencesOk() (*[]ExternalReference, bool)`

GetExternalReferencesOk returns a tuple with the ExternalReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalReferences

`func (o *StixIndicator) SetExternalReferences(v []ExternalReference)`

SetExternalReferences sets ExternalReferences field to given value.

### HasExternalReferences

`func (o *StixIndicator) HasExternalReferences() bool`

HasExternalReferences returns a boolean if a field has been set.

### GetObjectMarkingRefs

`func (o *StixIndicator) GetObjectMarkingRefs() []string`

GetObjectMarkingRefs returns the ObjectMarkingRefs field if non-nil, zero value otherwise.

### GetObjectMarkingRefsOk

`func (o *StixIndicator) GetObjectMarkingRefsOk() (*[]string, bool)`

GetObjectMarkingRefsOk returns a tuple with the ObjectMarkingRefs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectMarkingRefs

`func (o *StixIndicator) SetObjectMarkingRefs(v []string)`

SetObjectMarkingRefs sets ObjectMarkingRefs field to given value.

### HasObjectMarkingRefs

`func (o *StixIndicator) HasObjectMarkingRefs() bool`

HasObjectMarkingRefs returns a boolean if a field has been set.

### GetGranularMarkings

`func (o *StixIndicator) GetGranularMarkings() []GranularMarkingType`

GetGranularMarkings returns the GranularMarkings field if non-nil, zero value otherwise.

### GetGranularMarkingsOk

`func (o *StixIndicator) GetGranularMarkingsOk() (*[]GranularMarkingType, bool)`

GetGranularMarkingsOk returns a tuple with the GranularMarkings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularMarkings

`func (o *StixIndicator) SetGranularMarkings(v []GranularMarkingType)`

SetGranularMarkings sets GranularMarkings field to given value.

### HasGranularMarkings

`func (o *StixIndicator) HasGranularMarkings() bool`

HasGranularMarkings returns a boolean if a field has been set.

### GetExtensions

`func (o *StixIndicator) GetExtensions() map[string]Extension`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *StixIndicator) GetExtensionsOk() (*map[string]Extension, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *StixIndicator) SetExtensions(v map[string]Extension)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *StixIndicator) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetName

`func (o *StixIndicator) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StixIndicator) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StixIndicator) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StixIndicator) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *StixIndicator) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StixIndicator) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StixIndicator) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StixIndicator) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIndicatorTypes

`func (o *StixIndicator) GetIndicatorTypes() []string`

GetIndicatorTypes returns the IndicatorTypes field if non-nil, zero value otherwise.

### GetIndicatorTypesOk

`func (o *StixIndicator) GetIndicatorTypesOk() (*[]string, bool)`

GetIndicatorTypesOk returns a tuple with the IndicatorTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicatorTypes

`func (o *StixIndicator) SetIndicatorTypes(v []string)`

SetIndicatorTypes sets IndicatorTypes field to given value.

### HasIndicatorTypes

`func (o *StixIndicator) HasIndicatorTypes() bool`

HasIndicatorTypes returns a boolean if a field has been set.

### GetPattern

`func (o *StixIndicator) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *StixIndicator) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *StixIndicator) SetPattern(v string)`

SetPattern sets Pattern field to given value.


### GetPatternType

`func (o *StixIndicator) GetPatternType() string`

GetPatternType returns the PatternType field if non-nil, zero value otherwise.

### GetPatternTypeOk

`func (o *StixIndicator) GetPatternTypeOk() (*string, bool)`

GetPatternTypeOk returns a tuple with the PatternType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatternType

`func (o *StixIndicator) SetPatternType(v string)`

SetPatternType sets PatternType field to given value.


### GetPatternVersion

`func (o *StixIndicator) GetPatternVersion() string`

GetPatternVersion returns the PatternVersion field if non-nil, zero value otherwise.

### GetPatternVersionOk

`func (o *StixIndicator) GetPatternVersionOk() (*string, bool)`

GetPatternVersionOk returns a tuple with the PatternVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatternVersion

`func (o *StixIndicator) SetPatternVersion(v string)`

SetPatternVersion sets PatternVersion field to given value.

### HasPatternVersion

`func (o *StixIndicator) HasPatternVersion() bool`

HasPatternVersion returns a boolean if a field has been set.

### GetValidFrom

`func (o *StixIndicator) GetValidFrom() time.Time`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *StixIndicator) GetValidFromOk() (*time.Time, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *StixIndicator) SetValidFrom(v time.Time)`

SetValidFrom sets ValidFrom field to given value.


### GetValidUntil

`func (o *StixIndicator) GetValidUntil() time.Time`

GetValidUntil returns the ValidUntil field if non-nil, zero value otherwise.

### GetValidUntilOk

`func (o *StixIndicator) GetValidUntilOk() (*time.Time, bool)`

GetValidUntilOk returns a tuple with the ValidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUntil

`func (o *StixIndicator) SetValidUntil(v time.Time)`

SetValidUntil sets ValidUntil field to given value.

### HasValidUntil

`func (o *StixIndicator) HasValidUntil() bool`

HasValidUntil returns a boolean if a field has been set.

### GetKillChainPhases

`func (o *StixIndicator) GetKillChainPhases() []KillChainPhase`

GetKillChainPhases returns the KillChainPhases field if non-nil, zero value otherwise.

### GetKillChainPhasesOk

`func (o *StixIndicator) GetKillChainPhasesOk() (*[]KillChainPhase, bool)`

GetKillChainPhasesOk returns a tuple with the KillChainPhases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKillChainPhases

`func (o *StixIndicator) SetKillChainPhases(v []KillChainPhase)`

SetKillChainPhases sets KillChainPhases field to given value.

### HasKillChainPhases

`func (o *StixIndicator) HasKillChainPhases() bool`

HasKillChainPhases returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


