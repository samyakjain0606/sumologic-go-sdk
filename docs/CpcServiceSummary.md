# CpcServiceSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Service** | **string** | The name of the service. | 
**Color** | **string** | The color hex code assigned to the service. | 
**CpcSummary** | [**CpcSummary**](CpcSummary.md) |  | 

## Methods

### NewCpcServiceSummary

`func NewCpcServiceSummary(service string, color string, cpcSummary CpcSummary, ) *CpcServiceSummary`

NewCpcServiceSummary instantiates a new CpcServiceSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCpcServiceSummaryWithDefaults

`func NewCpcServiceSummaryWithDefaults() *CpcServiceSummary`

NewCpcServiceSummaryWithDefaults instantiates a new CpcServiceSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetService

`func (o *CpcServiceSummary) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *CpcServiceSummary) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *CpcServiceSummary) SetService(v string)`

SetService sets Service field to given value.


### GetColor

`func (o *CpcServiceSummary) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *CpcServiceSummary) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *CpcServiceSummary) SetColor(v string)`

SetColor sets Color field to given value.


### GetCpcSummary

`func (o *CpcServiceSummary) GetCpcSummary() CpcSummary`

GetCpcSummary returns the CpcSummary field if non-nil, zero value otherwise.

### GetCpcSummaryOk

`func (o *CpcServiceSummary) GetCpcSummaryOk() (*CpcSummary, bool)`

GetCpcSummaryOk returns a tuple with the CpcSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpcSummary

`func (o *CpcServiceSummary) SetCpcSummary(v CpcSummary)`

SetCpcSummary sets CpcSummary field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


