# CompliancePeriodProgress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RelativeReference** | **int32** | Relative shift of compliance period from the latest/current compliance period. | 
**StartTime** | **time.Time** | Start time of the compliance period. | 
**EndTime** | **time.Time** | End time of the compliance period. | 
**Progress** | **float64** | SLO data availability progress. | 
**IrrecoverableError** | **bool** | Whether a permanent error is encountered and no further progress is expected. | 

## Methods

### NewCompliancePeriodProgress

`func NewCompliancePeriodProgress(relativeReference int32, startTime time.Time, endTime time.Time, progress float64, irrecoverableError bool, ) *CompliancePeriodProgress`

NewCompliancePeriodProgress instantiates a new CompliancePeriodProgress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompliancePeriodProgressWithDefaults

`func NewCompliancePeriodProgressWithDefaults() *CompliancePeriodProgress`

NewCompliancePeriodProgressWithDefaults instantiates a new CompliancePeriodProgress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRelativeReference

`func (o *CompliancePeriodProgress) GetRelativeReference() int32`

GetRelativeReference returns the RelativeReference field if non-nil, zero value otherwise.

### GetRelativeReferenceOk

`func (o *CompliancePeriodProgress) GetRelativeReferenceOk() (*int32, bool)`

GetRelativeReferenceOk returns a tuple with the RelativeReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativeReference

`func (o *CompliancePeriodProgress) SetRelativeReference(v int32)`

SetRelativeReference sets RelativeReference field to given value.


### GetStartTime

`func (o *CompliancePeriodProgress) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *CompliancePeriodProgress) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *CompliancePeriodProgress) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.


### GetEndTime

`func (o *CompliancePeriodProgress) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *CompliancePeriodProgress) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *CompliancePeriodProgress) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.


### GetProgress

`func (o *CompliancePeriodProgress) GetProgress() float64`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *CompliancePeriodProgress) GetProgressOk() (*float64, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *CompliancePeriodProgress) SetProgress(v float64)`

SetProgress sets Progress field to given value.


### GetIrrecoverableError

`func (o *CompliancePeriodProgress) GetIrrecoverableError() bool`

GetIrrecoverableError returns the IrrecoverableError field if non-nil, zero value otherwise.

### GetIrrecoverableErrorOk

`func (o *CompliancePeriodProgress) GetIrrecoverableErrorOk() (*bool, bool)`

GetIrrecoverableErrorOk returns a tuple with the IrrecoverableError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIrrecoverableError

`func (o *CompliancePeriodProgress) SetIrrecoverableError(v bool)`

SetIrrecoverableError sets IrrecoverableError field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


