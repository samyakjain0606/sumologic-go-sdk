# CpcSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumOfTraces** | **int64** | The total number of traces matching the search criteria for a given service  based on which the CPC data is aggregated. | 
**AvgPercentageInTrace** | **float64** | The total fraction (value between 0.0 and 1.0) of the trace duration time consumed  by a given service (or a group of services) in the critical path of analyzed traces. | 
**AvgTimeInTrace** | **float64** | The average time in nanoseconds spent by a given service (or a group of services) in the critical path of analyzed traces. | 
**TotalTimeTaken** | **int64** | The total time in nanoseconds spent by a given service (or a group of services) in the critical path  of analyzed traces. | 

## Methods

### NewCpcSummary

`func NewCpcSummary(numOfTraces int64, avgPercentageInTrace float64, avgTimeInTrace float64, totalTimeTaken int64, ) *CpcSummary`

NewCpcSummary instantiates a new CpcSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCpcSummaryWithDefaults

`func NewCpcSummaryWithDefaults() *CpcSummary`

NewCpcSummaryWithDefaults instantiates a new CpcSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumOfTraces

`func (o *CpcSummary) GetNumOfTraces() int64`

GetNumOfTraces returns the NumOfTraces field if non-nil, zero value otherwise.

### GetNumOfTracesOk

`func (o *CpcSummary) GetNumOfTracesOk() (*int64, bool)`

GetNumOfTracesOk returns a tuple with the NumOfTraces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOfTraces

`func (o *CpcSummary) SetNumOfTraces(v int64)`

SetNumOfTraces sets NumOfTraces field to given value.


### GetAvgPercentageInTrace

`func (o *CpcSummary) GetAvgPercentageInTrace() float64`

GetAvgPercentageInTrace returns the AvgPercentageInTrace field if non-nil, zero value otherwise.

### GetAvgPercentageInTraceOk

`func (o *CpcSummary) GetAvgPercentageInTraceOk() (*float64, bool)`

GetAvgPercentageInTraceOk returns a tuple with the AvgPercentageInTrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPercentageInTrace

`func (o *CpcSummary) SetAvgPercentageInTrace(v float64)`

SetAvgPercentageInTrace sets AvgPercentageInTrace field to given value.


### GetAvgTimeInTrace

`func (o *CpcSummary) GetAvgTimeInTrace() float64`

GetAvgTimeInTrace returns the AvgTimeInTrace field if non-nil, zero value otherwise.

### GetAvgTimeInTraceOk

`func (o *CpcSummary) GetAvgTimeInTraceOk() (*float64, bool)`

GetAvgTimeInTraceOk returns a tuple with the AvgTimeInTrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgTimeInTrace

`func (o *CpcSummary) SetAvgTimeInTrace(v float64)`

SetAvgTimeInTrace sets AvgTimeInTrace field to given value.


### GetTotalTimeTaken

`func (o *CpcSummary) GetTotalTimeTaken() int64`

GetTotalTimeTaken returns the TotalTimeTaken field if non-nil, zero value otherwise.

### GetTotalTimeTakenOk

`func (o *CpcSummary) GetTotalTimeTakenOk() (*int64, bool)`

GetTotalTimeTakenOk returns a tuple with the TotalTimeTaken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTimeTaken

`func (o *CpcSummary) SetTotalTimeTaken(v int64)`

SetTotalTimeTaken sets TotalTimeTaken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


