# NextInstancesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timezone** | **string** | Time zone for the schedule per [IANA Time Zone Database](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones#List). | 
**StartDate** | **string** | Schedule start date in the format of &#x60;yyyy-mm-dd&#x60; | 
**StartTime** | **string** | Schedule start time in the format of &#x60;hh:mm&#x60; | 
**Rrule** | **string** | RRule (Recurrence Rule) | 

## Methods

### NewNextInstancesRequest

`func NewNextInstancesRequest(timezone string, startDate string, startTime string, rrule string, ) *NextInstancesRequest`

NewNextInstancesRequest instantiates a new NextInstancesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNextInstancesRequestWithDefaults

`func NewNextInstancesRequestWithDefaults() *NextInstancesRequest`

NewNextInstancesRequestWithDefaults instantiates a new NextInstancesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimezone

`func (o *NextInstancesRequest) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *NextInstancesRequest) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *NextInstancesRequest) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetStartDate

`func (o *NextInstancesRequest) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *NextInstancesRequest) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *NextInstancesRequest) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetStartTime

`func (o *NextInstancesRequest) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *NextInstancesRequest) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *NextInstancesRequest) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.


### GetRrule

`func (o *NextInstancesRequest) GetRrule() string`

GetRrule returns the Rrule field if non-nil, zero value otherwise.

### GetRruleOk

`func (o *NextInstancesRequest) GetRruleOk() (*string, bool)`

GetRruleOk returns a tuple with the Rrule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrule

`func (o *NextInstancesRequest) SetRrule(v string)`

SetRrule sets Rrule field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


