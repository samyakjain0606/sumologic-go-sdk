# UploadCsvIndicatorsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Imported** | **time.Time** | When this indicator was first loaded into Sumo. Timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format. | 
**Csv** | **string** | The CSV data containing indicators. | 

## Methods

### NewUploadCsvIndicatorsRequest

`func NewUploadCsvIndicatorsRequest(imported time.Time, csv string, ) *UploadCsvIndicatorsRequest`

NewUploadCsvIndicatorsRequest instantiates a new UploadCsvIndicatorsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadCsvIndicatorsRequestWithDefaults

`func NewUploadCsvIndicatorsRequestWithDefaults() *UploadCsvIndicatorsRequest`

NewUploadCsvIndicatorsRequestWithDefaults instantiates a new UploadCsvIndicatorsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImported

`func (o *UploadCsvIndicatorsRequest) GetImported() time.Time`

GetImported returns the Imported field if non-nil, zero value otherwise.

### GetImportedOk

`func (o *UploadCsvIndicatorsRequest) GetImportedOk() (*time.Time, bool)`

GetImportedOk returns a tuple with the Imported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImported

`func (o *UploadCsvIndicatorsRequest) SetImported(v time.Time)`

SetImported sets Imported field to given value.


### GetCsv

`func (o *UploadCsvIndicatorsRequest) GetCsv() string`

GetCsv returns the Csv field if non-nil, zero value otherwise.

### GetCsvOk

`func (o *UploadCsvIndicatorsRequest) GetCsvOk() (*string, bool)`

GetCsvOk returns a tuple with the Csv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsv

`func (o *UploadCsvIndicatorsRequest) SetCsv(v string)`

SetCsv sets Csv field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


