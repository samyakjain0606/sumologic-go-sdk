# BurnRate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BurnRateThreshold** | **float64** | The error budget depletion percentage. | 
**TimeRange** | **string** | The relative time range for measuring error budget depletion. | 

## Methods

### NewBurnRate

`func NewBurnRate(burnRateThreshold float64, timeRange string, ) *BurnRate`

NewBurnRate instantiates a new BurnRate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBurnRateWithDefaults

`func NewBurnRateWithDefaults() *BurnRate`

NewBurnRateWithDefaults instantiates a new BurnRate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBurnRateThreshold

`func (o *BurnRate) GetBurnRateThreshold() float64`

GetBurnRateThreshold returns the BurnRateThreshold field if non-nil, zero value otherwise.

### GetBurnRateThresholdOk

`func (o *BurnRate) GetBurnRateThresholdOk() (*float64, bool)`

GetBurnRateThresholdOk returns a tuple with the BurnRateThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurnRateThreshold

`func (o *BurnRate) SetBurnRateThreshold(v float64)`

SetBurnRateThreshold sets BurnRateThreshold field to given value.


### GetTimeRange

`func (o *BurnRate) GetTimeRange() string`

GetTimeRange returns the TimeRange field if non-nil, zero value otherwise.

### GetTimeRangeOk

`func (o *BurnRate) GetTimeRangeOk() (*string, bool)`

GetTimeRangeOk returns a tuple with the TimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRange

`func (o *BurnRate) SetTimeRange(v string)`

SetTimeRange sets TimeRange field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


