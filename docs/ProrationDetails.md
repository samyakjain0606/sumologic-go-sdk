# ProrationDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemainingDays** | **int32** | Remaining days in the billing cycle for which the new plan is prorated. | 
**ProratedCredits** | **int32** | Total prorated credits that get added to the bucket based on the remaining billing period. | 
**ProratedCost** | **float64** | Cost of the total prorated credits. | 

## Methods

### NewProrationDetails

`func NewProrationDetails(remainingDays int32, proratedCredits int32, proratedCost float64, ) *ProrationDetails`

NewProrationDetails instantiates a new ProrationDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProrationDetailsWithDefaults

`func NewProrationDetailsWithDefaults() *ProrationDetails`

NewProrationDetailsWithDefaults instantiates a new ProrationDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemainingDays

`func (o *ProrationDetails) GetRemainingDays() int32`

GetRemainingDays returns the RemainingDays field if non-nil, zero value otherwise.

### GetRemainingDaysOk

`func (o *ProrationDetails) GetRemainingDaysOk() (*int32, bool)`

GetRemainingDaysOk returns a tuple with the RemainingDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingDays

`func (o *ProrationDetails) SetRemainingDays(v int32)`

SetRemainingDays sets RemainingDays field to given value.


### GetProratedCredits

`func (o *ProrationDetails) GetProratedCredits() int32`

GetProratedCredits returns the ProratedCredits field if non-nil, zero value otherwise.

### GetProratedCreditsOk

`func (o *ProrationDetails) GetProratedCreditsOk() (*int32, bool)`

GetProratedCreditsOk returns a tuple with the ProratedCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProratedCredits

`func (o *ProrationDetails) SetProratedCredits(v int32)`

SetProratedCredits sets ProratedCredits field to given value.


### GetProratedCost

`func (o *ProrationDetails) GetProratedCost() float64`

GetProratedCost returns the ProratedCost field if non-nil, zero value otherwise.

### GetProratedCostOk

`func (o *ProrationDetails) GetProratedCostOk() (*float64, bool)`

GetProratedCostOk returns a tuple with the ProratedCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProratedCost

`func (o *ProrationDetails) SetProratedCost(v float64)`

SetProratedCost sets ProratedCost field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


