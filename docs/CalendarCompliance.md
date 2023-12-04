# CalendarCompliance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WindowType** | **string** | Type of Calendar Window (week/month/quarter). | 
**StartFrom** | Pointer to **string** | Start of the calendar window. For week, it would be the day of the week (for e.g Sunday, Monday etc). For month, it will always be the first day of the month (therefore not required to specify for monthly compliance). For quarter, it would be the first month of the quarter (for e.g January, February etc.) | [optional] 

## Methods

### NewCalendarCompliance

`func NewCalendarCompliance(windowType string, ) *CalendarCompliance`

NewCalendarCompliance instantiates a new CalendarCompliance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalendarComplianceWithDefaults

`func NewCalendarComplianceWithDefaults() *CalendarCompliance`

NewCalendarComplianceWithDefaults instantiates a new CalendarCompliance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWindowType

`func (o *CalendarCompliance) GetWindowType() string`

GetWindowType returns the WindowType field if non-nil, zero value otherwise.

### GetWindowTypeOk

`func (o *CalendarCompliance) GetWindowTypeOk() (*string, bool)`

GetWindowTypeOk returns a tuple with the WindowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowType

`func (o *CalendarCompliance) SetWindowType(v string)`

SetWindowType sets WindowType field to given value.


### GetStartFrom

`func (o *CalendarCompliance) GetStartFrom() string`

GetStartFrom returns the StartFrom field if non-nil, zero value otherwise.

### GetStartFromOk

`func (o *CalendarCompliance) GetStartFromOk() (*string, bool)`

GetStartFromOk returns a tuple with the StartFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartFrom

`func (o *CalendarCompliance) SetStartFrom(v string)`

SetStartFrom sets StartFrom field to given value.

### HasStartFrom

`func (o *CalendarCompliance) HasStartFrom() bool`

HasStartFrom returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


