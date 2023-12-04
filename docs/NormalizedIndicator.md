# NormalizedIndicator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the indicator | 
**Indicator** | **string** | Value of the indicator | 
**Type** | **string** | Type of indicator | 
**Source** | **string** | User-provided text to identify the source of the indicator | 
**Imported** | **time.Time** | When this indicator was first loaded into Sumo. Timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format. | 
**Updated** | Pointer to **time.Time** | When this indicator was most recently updated in Sumo. Timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format. | [optional] 
**ValidFrom** | **time.Time** | Beginning time this indicator is valid. Timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format. | 
**ValidUntil** | Pointer to **time.Time** | Time at which this indicator expires. If not set, the indicator never expires. . Timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format. | [optional] 
**Confidence** | **int32** | Confidence that the creator has in the correctness of their data, where 100 is highest | 
**ThreatType** | **string** | Type of indicator ( https://docs.oasis-open.org/cti/stix/v2.1/os/stix-v2.1-os.html#_cvhfwe3t9vuo ) | 
**Fields** | Pointer to **map[string]string** | Flattened fields from the original indicator object (e.g. flattened STIX fields) | [optional] 

## Methods

### NewNormalizedIndicator

`func NewNormalizedIndicator(id string, indicator string, type_ string, source string, imported time.Time, validFrom time.Time, confidence int32, threatType string, ) *NormalizedIndicator`

NewNormalizedIndicator instantiates a new NormalizedIndicator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNormalizedIndicatorWithDefaults

`func NewNormalizedIndicatorWithDefaults() *NormalizedIndicator`

NewNormalizedIndicatorWithDefaults instantiates a new NormalizedIndicator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NormalizedIndicator) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NormalizedIndicator) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NormalizedIndicator) SetId(v string)`

SetId sets Id field to given value.


### GetIndicator

`func (o *NormalizedIndicator) GetIndicator() string`

GetIndicator returns the Indicator field if non-nil, zero value otherwise.

### GetIndicatorOk

`func (o *NormalizedIndicator) GetIndicatorOk() (*string, bool)`

GetIndicatorOk returns a tuple with the Indicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicator

`func (o *NormalizedIndicator) SetIndicator(v string)`

SetIndicator sets Indicator field to given value.


### GetType

`func (o *NormalizedIndicator) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NormalizedIndicator) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NormalizedIndicator) SetType(v string)`

SetType sets Type field to given value.


### GetSource

`func (o *NormalizedIndicator) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *NormalizedIndicator) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *NormalizedIndicator) SetSource(v string)`

SetSource sets Source field to given value.


### GetImported

`func (o *NormalizedIndicator) GetImported() time.Time`

GetImported returns the Imported field if non-nil, zero value otherwise.

### GetImportedOk

`func (o *NormalizedIndicator) GetImportedOk() (*time.Time, bool)`

GetImportedOk returns a tuple with the Imported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImported

`func (o *NormalizedIndicator) SetImported(v time.Time)`

SetImported sets Imported field to given value.


### GetUpdated

`func (o *NormalizedIndicator) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *NormalizedIndicator) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *NormalizedIndicator) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *NormalizedIndicator) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetValidFrom

`func (o *NormalizedIndicator) GetValidFrom() time.Time`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *NormalizedIndicator) GetValidFromOk() (*time.Time, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *NormalizedIndicator) SetValidFrom(v time.Time)`

SetValidFrom sets ValidFrom field to given value.


### GetValidUntil

`func (o *NormalizedIndicator) GetValidUntil() time.Time`

GetValidUntil returns the ValidUntil field if non-nil, zero value otherwise.

### GetValidUntilOk

`func (o *NormalizedIndicator) GetValidUntilOk() (*time.Time, bool)`

GetValidUntilOk returns a tuple with the ValidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUntil

`func (o *NormalizedIndicator) SetValidUntil(v time.Time)`

SetValidUntil sets ValidUntil field to given value.

### HasValidUntil

`func (o *NormalizedIndicator) HasValidUntil() bool`

HasValidUntil returns a boolean if a field has been set.

### GetConfidence

`func (o *NormalizedIndicator) GetConfidence() int32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *NormalizedIndicator) GetConfidenceOk() (*int32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *NormalizedIndicator) SetConfidence(v int32)`

SetConfidence sets Confidence field to given value.


### GetThreatType

`func (o *NormalizedIndicator) GetThreatType() string`

GetThreatType returns the ThreatType field if non-nil, zero value otherwise.

### GetThreatTypeOk

`func (o *NormalizedIndicator) GetThreatTypeOk() (*string, bool)`

GetThreatTypeOk returns a tuple with the ThreatType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatType

`func (o *NormalizedIndicator) SetThreatType(v string)`

SetThreatType sets ThreatType field to given value.


### GetFields

`func (o *NormalizedIndicator) GetFields() map[string]string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *NormalizedIndicator) GetFieldsOk() (*map[string]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *NormalizedIndicator) SetFields(v map[string]string)`

SetFields sets Fields field to given value.

### HasFields

`func (o *NormalizedIndicator) HasFields() bool`

HasFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


