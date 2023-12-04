# ExternalReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceName** | **string** | The name of the source that the external-reference is defined within | 
**Description** | Pointer to **string** | A human readable description | [optional] 
**Url** | Pointer to **string** | A URL reference to an external resource | [optional] 
**Hashes** | Pointer to **map[string]string** | Specifies a dictionary of hashes for the contents of the url | [optional] 
**ExternalId** | Pointer to **string** | An identifier for the external reference content | [optional] 

## Methods

### NewExternalReference

`func NewExternalReference(sourceName string, ) *ExternalReference`

NewExternalReference instantiates a new ExternalReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalReferenceWithDefaults

`func NewExternalReferenceWithDefaults() *ExternalReference`

NewExternalReferenceWithDefaults instantiates a new ExternalReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceName

`func (o *ExternalReference) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *ExternalReference) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *ExternalReference) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.


### GetDescription

`func (o *ExternalReference) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExternalReference) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExternalReference) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExternalReference) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUrl

`func (o *ExternalReference) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ExternalReference) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ExternalReference) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ExternalReference) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetHashes

`func (o *ExternalReference) GetHashes() map[string]string`

GetHashes returns the Hashes field if non-nil, zero value otherwise.

### GetHashesOk

`func (o *ExternalReference) GetHashesOk() (*map[string]string, bool)`

GetHashesOk returns a tuple with the Hashes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashes

`func (o *ExternalReference) SetHashes(v map[string]string)`

SetHashes sets Hashes field to given value.

### HasHashes

`func (o *ExternalReference) HasHashes() bool`

HasHashes returns a boolean if a field has been set.

### GetExternalId

`func (o *ExternalReference) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ExternalReference) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ExternalReference) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ExternalReference) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


