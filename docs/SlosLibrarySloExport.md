# SlosLibrarySloExport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SignalType** | **string** | Type of SLI Signal (latency, error, throughput, availability or other). | 
**Compliance** | [**Compliance**](Compliance.md) |  | 
**Indicator** | [**Sli**](Sli.md) |  | 
**Service** | Pointer to **string** | Name of the service. | [optional] 
**Application** | Pointer to **string** | Name of the application. | [optional] 
**Tags** | Pointer to **map[string]string** | Tags to be associated with the SLO. | [optional] 

## Methods

### NewSlosLibrarySloExport

`func NewSlosLibrarySloExport(signalType string, compliance Compliance, indicator Sli, ) *SlosLibrarySloExport`

NewSlosLibrarySloExport instantiates a new SlosLibrarySloExport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlosLibrarySloExportWithDefaults

`func NewSlosLibrarySloExportWithDefaults() *SlosLibrarySloExport`

NewSlosLibrarySloExportWithDefaults instantiates a new SlosLibrarySloExport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignalType

`func (o *SlosLibrarySloExport) GetSignalType() string`

GetSignalType returns the SignalType field if non-nil, zero value otherwise.

### GetSignalTypeOk

`func (o *SlosLibrarySloExport) GetSignalTypeOk() (*string, bool)`

GetSignalTypeOk returns a tuple with the SignalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalType

`func (o *SlosLibrarySloExport) SetSignalType(v string)`

SetSignalType sets SignalType field to given value.


### GetCompliance

`func (o *SlosLibrarySloExport) GetCompliance() Compliance`

GetCompliance returns the Compliance field if non-nil, zero value otherwise.

### GetComplianceOk

`func (o *SlosLibrarySloExport) GetComplianceOk() (*Compliance, bool)`

GetComplianceOk returns a tuple with the Compliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompliance

`func (o *SlosLibrarySloExport) SetCompliance(v Compliance)`

SetCompliance sets Compliance field to given value.


### GetIndicator

`func (o *SlosLibrarySloExport) GetIndicator() Sli`

GetIndicator returns the Indicator field if non-nil, zero value otherwise.

### GetIndicatorOk

`func (o *SlosLibrarySloExport) GetIndicatorOk() (*Sli, bool)`

GetIndicatorOk returns a tuple with the Indicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicator

`func (o *SlosLibrarySloExport) SetIndicator(v Sli)`

SetIndicator sets Indicator field to given value.


### GetService

`func (o *SlosLibrarySloExport) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *SlosLibrarySloExport) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *SlosLibrarySloExport) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *SlosLibrarySloExport) HasService() bool`

HasService returns a boolean if a field has been set.

### GetApplication

`func (o *SlosLibrarySloExport) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *SlosLibrarySloExport) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *SlosLibrarySloExport) SetApplication(v string)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *SlosLibrarySloExport) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetTags

`func (o *SlosLibrarySloExport) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SlosLibrarySloExport) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SlosLibrarySloExport) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SlosLibrarySloExport) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


