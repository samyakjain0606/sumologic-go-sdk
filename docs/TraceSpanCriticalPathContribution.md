# TraceSpanCriticalPathContribution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | **int64** | Overall processing time in nanoseconds consumed by this span in the critical path of its trace (a sum of the duration times of this span&#39;s critical path segments). | 
**Fraction** | **float64** | The total fraction (value between 0.0 and 1.0) of the trace duration time consumed by this span in the critical path of its trace. | 

## Methods

### NewTraceSpanCriticalPathContribution

`func NewTraceSpanCriticalPathContribution(duration int64, fraction float64, ) *TraceSpanCriticalPathContribution`

NewTraceSpanCriticalPathContribution instantiates a new TraceSpanCriticalPathContribution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTraceSpanCriticalPathContributionWithDefaults

`func NewTraceSpanCriticalPathContributionWithDefaults() *TraceSpanCriticalPathContribution`

NewTraceSpanCriticalPathContributionWithDefaults instantiates a new TraceSpanCriticalPathContribution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *TraceSpanCriticalPathContribution) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *TraceSpanCriticalPathContribution) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *TraceSpanCriticalPathContribution) SetDuration(v int64)`

SetDuration sets Duration field to given value.


### GetFraction

`func (o *TraceSpanCriticalPathContribution) GetFraction() float64`

GetFraction returns the Fraction field if non-nil, zero value otherwise.

### GetFractionOk

`func (o *TraceSpanCriticalPathContribution) GetFractionOk() (*float64, bool)`

GetFractionOk returns a tuple with the Fraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFraction

`func (o *TraceSpanCriticalPathContribution) SetFraction(v float64)`

SetFraction sets Fraction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


