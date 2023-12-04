# SloBurnRateCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BurnRateThreshold** | Pointer to **float64** | The error budget depletion percentage. | [optional] 
**TimeRange** | Pointer to **string** | The relative time range for measuring error budget depletion. | [optional] 

## Methods

### NewSloBurnRateCondition

`func NewSloBurnRateCondition() *SloBurnRateCondition`

NewSloBurnRateCondition instantiates a new SloBurnRateCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSloBurnRateConditionWithDefaults

`func NewSloBurnRateConditionWithDefaults() *SloBurnRateCondition`

NewSloBurnRateConditionWithDefaults instantiates a new SloBurnRateCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBurnRateThreshold

`func (o *SloBurnRateCondition) GetBurnRateThreshold() float64`

GetBurnRateThreshold returns the BurnRateThreshold field if non-nil, zero value otherwise.

### GetBurnRateThresholdOk

`func (o *SloBurnRateCondition) GetBurnRateThresholdOk() (*float64, bool)`

GetBurnRateThresholdOk returns a tuple with the BurnRateThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurnRateThreshold

`func (o *SloBurnRateCondition) SetBurnRateThreshold(v float64)`

SetBurnRateThreshold sets BurnRateThreshold field to given value.

### HasBurnRateThreshold

`func (o *SloBurnRateCondition) HasBurnRateThreshold() bool`

HasBurnRateThreshold returns a boolean if a field has been set.

### GetTimeRange

`func (o *SloBurnRateCondition) GetTimeRange() string`

GetTimeRange returns the TimeRange field if non-nil, zero value otherwise.

### GetTimeRangeOk

`func (o *SloBurnRateCondition) GetTimeRangeOk() (*string, bool)`

GetTimeRangeOk returns a tuple with the TimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRange

`func (o *SloBurnRateCondition) SetTimeRange(v string)`

SetTimeRange sets TimeRange field to given value.

### HasTimeRange

`func (o *SloBurnRateCondition) HasTimeRange() bool`

HasTimeRange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


