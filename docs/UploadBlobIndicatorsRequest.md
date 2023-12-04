# UploadBlobIndicatorsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | **string** | User-provided text to identify the source of the indicator | 
**Imported** | **time.Time** | When this indicator was first loaded into Sumo. Timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format. | 
**Format** | **string** | The input format. | 
**Blob** | **string** | The blob containing indicators. | 

## Methods

### NewUploadBlobIndicatorsRequest

`func NewUploadBlobIndicatorsRequest(source string, imported time.Time, format string, blob string, ) *UploadBlobIndicatorsRequest`

NewUploadBlobIndicatorsRequest instantiates a new UploadBlobIndicatorsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadBlobIndicatorsRequestWithDefaults

`func NewUploadBlobIndicatorsRequestWithDefaults() *UploadBlobIndicatorsRequest`

NewUploadBlobIndicatorsRequestWithDefaults instantiates a new UploadBlobIndicatorsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *UploadBlobIndicatorsRequest) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *UploadBlobIndicatorsRequest) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *UploadBlobIndicatorsRequest) SetSource(v string)`

SetSource sets Source field to given value.


### GetImported

`func (o *UploadBlobIndicatorsRequest) GetImported() time.Time`

GetImported returns the Imported field if non-nil, zero value otherwise.

### GetImportedOk

`func (o *UploadBlobIndicatorsRequest) GetImportedOk() (*time.Time, bool)`

GetImportedOk returns a tuple with the Imported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImported

`func (o *UploadBlobIndicatorsRequest) SetImported(v time.Time)`

SetImported sets Imported field to given value.


### GetFormat

`func (o *UploadBlobIndicatorsRequest) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *UploadBlobIndicatorsRequest) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *UploadBlobIndicatorsRequest) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetBlob

`func (o *UploadBlobIndicatorsRequest) GetBlob() string`

GetBlob returns the Blob field if non-nil, zero value otherwise.

### GetBlobOk

`func (o *UploadBlobIndicatorsRequest) GetBlobOk() (*string, bool)`

GetBlobOk returns a tuple with the Blob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlob

`func (o *UploadBlobIndicatorsRequest) SetBlob(v string)`

SetBlob sets Blob field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


