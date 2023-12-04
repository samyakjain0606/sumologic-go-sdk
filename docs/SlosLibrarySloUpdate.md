# SlosLibrarySloUpdate

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

### NewSlosLibrarySloUpdate

`func NewSlosLibrarySloUpdate(signalType string, compliance Compliance, indicator Sli, ) *SlosLibrarySloUpdate`

NewSlosLibrarySloUpdate instantiates a new SlosLibrarySloUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlosLibrarySloUpdateWithDefaults

`func NewSlosLibrarySloUpdateWithDefaults() *SlosLibrarySloUpdate`

NewSlosLibrarySloUpdateWithDefaults instantiates a new SlosLibrarySloUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignalType

`func (o *SlosLibrarySloUpdate) GetSignalType() string`

GetSignalType returns the SignalType field if non-nil, zero value otherwise.

### GetSignalTypeOk

`func (o *SlosLibrarySloUpdate) GetSignalTypeOk() (*string, bool)`

GetSignalTypeOk returns a tuple with the SignalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalType

`func (o *SlosLibrarySloUpdate) SetSignalType(v string)`

SetSignalType sets SignalType field to given value.


### GetCompliance

`func (o *SlosLibrarySloUpdate) GetCompliance() Compliance`

GetCompliance returns the Compliance field if non-nil, zero value otherwise.

### GetComplianceOk

`func (o *SlosLibrarySloUpdate) GetComplianceOk() (*Compliance, bool)`

GetComplianceOk returns a tuple with the Compliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompliance

`func (o *SlosLibrarySloUpdate) SetCompliance(v Compliance)`

SetCompliance sets Compliance field to given value.


### GetIndicator

`func (o *SlosLibrarySloUpdate) GetIndicator() Sli`

GetIndicator returns the Indicator field if non-nil, zero value otherwise.

### GetIndicatorOk

`func (o *SlosLibrarySloUpdate) GetIndicatorOk() (*Sli, bool)`

GetIndicatorOk returns a tuple with the Indicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicator

`func (o *SlosLibrarySloUpdate) SetIndicator(v Sli)`

SetIndicator sets Indicator field to given value.


### GetService

`func (o *SlosLibrarySloUpdate) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *SlosLibrarySloUpdate) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *SlosLibrarySloUpdate) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *SlosLibrarySloUpdate) HasService() bool`

HasService returns a boolean if a field has been set.

### GetApplication

`func (o *SlosLibrarySloUpdate) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *SlosLibrarySloUpdate) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *SlosLibrarySloUpdate) SetApplication(v string)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *SlosLibrarySloUpdate) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetTags

`func (o *SlosLibrarySloUpdate) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SlosLibrarySloUpdate) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SlosLibrarySloUpdate) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SlosLibrarySloUpdate) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


