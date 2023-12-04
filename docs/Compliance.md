# Compliance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComplianceType** | **string** | Compliance Type (rolling or calendar) | 
**Target** | **float32** | Target percentage for the SLI over the compliance period. | 
**Timezone** | **string** | Time zone for the SLO compliance. Follow the format in the [IANA Time Zone Database](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones#List). | 

## Methods

### NewCompliance

`func NewCompliance(complianceType string, target float32, timezone string, ) *Compliance`

NewCompliance instantiates a new Compliance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComplianceWithDefaults

`func NewComplianceWithDefaults() *Compliance`

NewComplianceWithDefaults instantiates a new Compliance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComplianceType

`func (o *Compliance) GetComplianceType() string`

GetComplianceType returns the ComplianceType field if non-nil, zero value otherwise.

### GetComplianceTypeOk

`func (o *Compliance) GetComplianceTypeOk() (*string, bool)`

GetComplianceTypeOk returns a tuple with the ComplianceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceType

`func (o *Compliance) SetComplianceType(v string)`

SetComplianceType sets ComplianceType field to given value.


### GetTarget

`func (o *Compliance) GetTarget() float32`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *Compliance) GetTargetOk() (*float32, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *Compliance) SetTarget(v float32)`

SetTarget sets Target field to given value.


### GetTimezone

`func (o *Compliance) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *Compliance) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *Compliance) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


