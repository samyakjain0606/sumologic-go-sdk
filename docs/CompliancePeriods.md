# CompliancePeriods

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timezone** | **string** | Time zone for the compliance periods as per the [IANA Time Zone Database](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones#List). | 
**Periods** | [**[]CompliancePeriodProgress**](CompliancePeriodProgress.md) | List of CompliancePeriodProgress. | 

## Methods

### NewCompliancePeriods

`func NewCompliancePeriods(timezone string, periods []CompliancePeriodProgress, ) *CompliancePeriods`

NewCompliancePeriods instantiates a new CompliancePeriods object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompliancePeriodsWithDefaults

`func NewCompliancePeriodsWithDefaults() *CompliancePeriods`

NewCompliancePeriodsWithDefaults instantiates a new CompliancePeriods object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimezone

`func (o *CompliancePeriods) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *CompliancePeriods) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *CompliancePeriods) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetPeriods

`func (o *CompliancePeriods) GetPeriods() []CompliancePeriodProgress`

GetPeriods returns the Periods field if non-nil, zero value otherwise.

### GetPeriodsOk

`func (o *CompliancePeriods) GetPeriodsOk() (*[]CompliancePeriodProgress, bool)`

GetPeriodsOk returns a tuple with the Periods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriods

`func (o *CompliancePeriods) SetPeriods(v []CompliancePeriodProgress)`

SetPeriods sets Periods field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


