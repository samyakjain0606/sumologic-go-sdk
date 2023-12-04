# SlosLibrarySloResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SignalType** | **string** | Type of SLI Signal (latency, error, throughput, availability or other). | 
**Compliance** | [**Compliance**](Compliance.md) |  | 
**Indicator** | [**Sli**](Sli.md) |  | 
**Service** | Pointer to **string** | Name of the service. | [optional] 
**Application** | Pointer to **string** | Name of the application. | [optional] 
**Tags** | Pointer to **map[string]string** | Tags to be associated with the SLO. | [optional] 
**SloVersion** | Pointer to **int64** | Current SLO Version. This is incremented on every change of a critical field of the SLO (i.e, SLI or Compliance period timezone), that requires a recompute of the SLI values over the compliance period. | [optional] 

## Methods

### NewSlosLibrarySloResponse

`func NewSlosLibrarySloResponse(signalType string, compliance Compliance, indicator Sli, ) *SlosLibrarySloResponse`

NewSlosLibrarySloResponse instantiates a new SlosLibrarySloResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlosLibrarySloResponseWithDefaults

`func NewSlosLibrarySloResponseWithDefaults() *SlosLibrarySloResponse`

NewSlosLibrarySloResponseWithDefaults instantiates a new SlosLibrarySloResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignalType

`func (o *SlosLibrarySloResponse) GetSignalType() string`

GetSignalType returns the SignalType field if non-nil, zero value otherwise.

### GetSignalTypeOk

`func (o *SlosLibrarySloResponse) GetSignalTypeOk() (*string, bool)`

GetSignalTypeOk returns a tuple with the SignalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalType

`func (o *SlosLibrarySloResponse) SetSignalType(v string)`

SetSignalType sets SignalType field to given value.


### GetCompliance

`func (o *SlosLibrarySloResponse) GetCompliance() Compliance`

GetCompliance returns the Compliance field if non-nil, zero value otherwise.

### GetComplianceOk

`func (o *SlosLibrarySloResponse) GetComplianceOk() (*Compliance, bool)`

GetComplianceOk returns a tuple with the Compliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompliance

`func (o *SlosLibrarySloResponse) SetCompliance(v Compliance)`

SetCompliance sets Compliance field to given value.


### GetIndicator

`func (o *SlosLibrarySloResponse) GetIndicator() Sli`

GetIndicator returns the Indicator field if non-nil, zero value otherwise.

### GetIndicatorOk

`func (o *SlosLibrarySloResponse) GetIndicatorOk() (*Sli, bool)`

GetIndicatorOk returns a tuple with the Indicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicator

`func (o *SlosLibrarySloResponse) SetIndicator(v Sli)`

SetIndicator sets Indicator field to given value.


### GetService

`func (o *SlosLibrarySloResponse) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *SlosLibrarySloResponse) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *SlosLibrarySloResponse) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *SlosLibrarySloResponse) HasService() bool`

HasService returns a boolean if a field has been set.

### GetApplication

`func (o *SlosLibrarySloResponse) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *SlosLibrarySloResponse) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *SlosLibrarySloResponse) SetApplication(v string)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *SlosLibrarySloResponse) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetTags

`func (o *SlosLibrarySloResponse) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SlosLibrarySloResponse) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SlosLibrarySloResponse) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SlosLibrarySloResponse) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetSloVersion

`func (o *SlosLibrarySloResponse) GetSloVersion() int64`

GetSloVersion returns the SloVersion field if non-nil, zero value otherwise.

### GetSloVersionOk

`func (o *SlosLibrarySloResponse) GetSloVersionOk() (*int64, bool)`

GetSloVersionOk returns a tuple with the SloVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSloVersion

`func (o *SlosLibrarySloResponse) SetSloVersion(v int64)`

SetSloVersion sets SloVersion field to given value.

### HasSloVersion

`func (o *SlosLibrarySloResponse) HasSloVersion() bool`

HasSloVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


