# ScheduleDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timezone** | **string** | Time zone for the schedule per [IANA Time Zone Database](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones#List). | 
**StartDate** | **string** | Schedule start date in the format of &#x60;yyyy-mm-dd&#x60; | 
**StartTime** | **string** | Schedule start time in the format of &#x60;hh:mm&#x60; | 
**Duration** | **int32** | Duration of the muting in minutes | 
**Rrule** | Pointer to **string** | RRule (Recurrence Rule) | [optional] 
**IsForm** | Pointer to **bool** | A flag identifying if the RRule is created or modified through Form UI | [optional] 

## Methods

### NewScheduleDefinition

`func NewScheduleDefinition(timezone string, startDate string, startTime string, duration int32, ) *ScheduleDefinition`

NewScheduleDefinition instantiates a new ScheduleDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleDefinitionWithDefaults

`func NewScheduleDefinitionWithDefaults() *ScheduleDefinition`

NewScheduleDefinitionWithDefaults instantiates a new ScheduleDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimezone

`func (o *ScheduleDefinition) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ScheduleDefinition) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ScheduleDefinition) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetStartDate

`func (o *ScheduleDefinition) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ScheduleDefinition) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ScheduleDefinition) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetStartTime

`func (o *ScheduleDefinition) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ScheduleDefinition) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ScheduleDefinition) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.


### GetDuration

`func (o *ScheduleDefinition) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ScheduleDefinition) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ScheduleDefinition) SetDuration(v int32)`

SetDuration sets Duration field to given value.


### GetRrule

`func (o *ScheduleDefinition) GetRrule() string`

GetRrule returns the Rrule field if non-nil, zero value otherwise.

### GetRruleOk

`func (o *ScheduleDefinition) GetRruleOk() (*string, bool)`

GetRruleOk returns a tuple with the Rrule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrule

`func (o *ScheduleDefinition) SetRrule(v string)`

SetRrule sets Rrule field to given value.

### HasRrule

`func (o *ScheduleDefinition) HasRrule() bool`

HasRrule returns a boolean if a field has been set.

### GetIsForm

`func (o *ScheduleDefinition) GetIsForm() bool`

GetIsForm returns the IsForm field if non-nil, zero value otherwise.

### GetIsFormOk

`func (o *ScheduleDefinition) GetIsFormOk() (*bool, bool)`

GetIsFormOk returns a tuple with the IsForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsForm

`func (o *ScheduleDefinition) SetIsForm(v bool)`

SetIsForm sets IsForm field to given value.

### HasIsForm

`func (o *ScheduleDefinition) HasIsForm() bool`

HasIsForm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


