# ChildUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCreditsUsed** | **float64** | Total Credits used by the child org. | 
**UsagePercentage** | Pointer to **float64** | Percentage of used credits from the allocated credits. | [optional] 
**ForecastPercentage** | Pointer to **float64** | Forecasted percentage of credits will be used in the given time period. | [optional] 
**UsagePercentChangeWoW** | Pointer to **float64** | Week over week usage percentage for the subscription period. | [optional] 
**UsagePercentChange** | Pointer to **float64** | Percentage of usage change over the given time period. | [optional] 

## Methods

### NewChildUsage

`func NewChildUsage(totalCreditsUsed float64, ) *ChildUsage`

NewChildUsage instantiates a new ChildUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChildUsageWithDefaults

`func NewChildUsageWithDefaults() *ChildUsage`

NewChildUsageWithDefaults instantiates a new ChildUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCreditsUsed

`func (o *ChildUsage) GetTotalCreditsUsed() float64`

GetTotalCreditsUsed returns the TotalCreditsUsed field if non-nil, zero value otherwise.

### GetTotalCreditsUsedOk

`func (o *ChildUsage) GetTotalCreditsUsedOk() (*float64, bool)`

GetTotalCreditsUsedOk returns a tuple with the TotalCreditsUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCreditsUsed

`func (o *ChildUsage) SetTotalCreditsUsed(v float64)`

SetTotalCreditsUsed sets TotalCreditsUsed field to given value.


### GetUsagePercentage

`func (o *ChildUsage) GetUsagePercentage() float64`

GetUsagePercentage returns the UsagePercentage field if non-nil, zero value otherwise.

### GetUsagePercentageOk

`func (o *ChildUsage) GetUsagePercentageOk() (*float64, bool)`

GetUsagePercentageOk returns a tuple with the UsagePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsagePercentage

`func (o *ChildUsage) SetUsagePercentage(v float64)`

SetUsagePercentage sets UsagePercentage field to given value.

### HasUsagePercentage

`func (o *ChildUsage) HasUsagePercentage() bool`

HasUsagePercentage returns a boolean if a field has been set.

### GetForecastPercentage

`func (o *ChildUsage) GetForecastPercentage() float64`

GetForecastPercentage returns the ForecastPercentage field if non-nil, zero value otherwise.

### GetForecastPercentageOk

`func (o *ChildUsage) GetForecastPercentageOk() (*float64, bool)`

GetForecastPercentageOk returns a tuple with the ForecastPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForecastPercentage

`func (o *ChildUsage) SetForecastPercentage(v float64)`

SetForecastPercentage sets ForecastPercentage field to given value.

### HasForecastPercentage

`func (o *ChildUsage) HasForecastPercentage() bool`

HasForecastPercentage returns a boolean if a field has been set.

### GetUsagePercentChangeWoW

`func (o *ChildUsage) GetUsagePercentChangeWoW() float64`

GetUsagePercentChangeWoW returns the UsagePercentChangeWoW field if non-nil, zero value otherwise.

### GetUsagePercentChangeWoWOk

`func (o *ChildUsage) GetUsagePercentChangeWoWOk() (*float64, bool)`

GetUsagePercentChangeWoWOk returns a tuple with the UsagePercentChangeWoW field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsagePercentChangeWoW

`func (o *ChildUsage) SetUsagePercentChangeWoW(v float64)`

SetUsagePercentChangeWoW sets UsagePercentChangeWoW field to given value.

### HasUsagePercentChangeWoW

`func (o *ChildUsage) HasUsagePercentChangeWoW() bool`

HasUsagePercentChangeWoW returns a boolean if a field has been set.

### GetUsagePercentChange

`func (o *ChildUsage) GetUsagePercentChange() float64`

GetUsagePercentChange returns the UsagePercentChange field if non-nil, zero value otherwise.

### GetUsagePercentChangeOk

`func (o *ChildUsage) GetUsagePercentChangeOk() (*float64, bool)`

GetUsagePercentChangeOk returns a tuple with the UsagePercentChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsagePercentChange

`func (o *ChildUsage) SetUsagePercentChange(v float64)`

SetUsagePercentChange sets UsagePercentChange field to given value.

### HasUsagePercentChange

`func (o *ChildUsage) HasUsagePercentChange() bool`

HasUsagePercentChange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


