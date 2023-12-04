# CpcQueryBucketResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketId** | **string** | A unique identifier of a time bucket. | 
**StartTimestamp** | **time.Time** | A start of the time bucket in the [ISO 8601 / RFC3339](https://tools.ietf.org/html/rfc3339) format. | 
**Length** | **int64** | The length of a time bucket expressed in milliseconds. | 
**TotalNumOfTraces** | **int64** | The total number of traces matching the search criteria based on which the CPC data is aggregated. | 
**AvgTraceDuration** | **float64** | The average duration in nanoseconds of the traces matching the search criteria  based on which the CPC data is aggregated. | 
**PerServiceCpcSummaries** | [**[]CpcServiceSummary**](CpcServiceSummary.md) | The summary of aggregated Critical Path Contribution data on a per service basis.  Each element of the array corresponds to a summary for a specific service. | 
**OtherServicesCpcSummary** | [**CpcSummary**](CpcSummary.md) |  | 
**IdleTimeCpcSummary** | [**CpcSummary**](CpcSummary.md) |  | 

## Methods

### NewCpcQueryBucketResult

`func NewCpcQueryBucketResult(bucketId string, startTimestamp time.Time, length int64, totalNumOfTraces int64, avgTraceDuration float64, perServiceCpcSummaries []CpcServiceSummary, otherServicesCpcSummary CpcSummary, idleTimeCpcSummary CpcSummary, ) *CpcQueryBucketResult`

NewCpcQueryBucketResult instantiates a new CpcQueryBucketResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCpcQueryBucketResultWithDefaults

`func NewCpcQueryBucketResultWithDefaults() *CpcQueryBucketResult`

NewCpcQueryBucketResultWithDefaults instantiates a new CpcQueryBucketResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketId

`func (o *CpcQueryBucketResult) GetBucketId() string`

GetBucketId returns the BucketId field if non-nil, zero value otherwise.

### GetBucketIdOk

`func (o *CpcQueryBucketResult) GetBucketIdOk() (*string, bool)`

GetBucketIdOk returns a tuple with the BucketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketId

`func (o *CpcQueryBucketResult) SetBucketId(v string)`

SetBucketId sets BucketId field to given value.


### GetStartTimestamp

`func (o *CpcQueryBucketResult) GetStartTimestamp() time.Time`

GetStartTimestamp returns the StartTimestamp field if non-nil, zero value otherwise.

### GetStartTimestampOk

`func (o *CpcQueryBucketResult) GetStartTimestampOk() (*time.Time, bool)`

GetStartTimestampOk returns a tuple with the StartTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimestamp

`func (o *CpcQueryBucketResult) SetStartTimestamp(v time.Time)`

SetStartTimestamp sets StartTimestamp field to given value.


### GetLength

`func (o *CpcQueryBucketResult) GetLength() int64`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *CpcQueryBucketResult) GetLengthOk() (*int64, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *CpcQueryBucketResult) SetLength(v int64)`

SetLength sets Length field to given value.


### GetTotalNumOfTraces

`func (o *CpcQueryBucketResult) GetTotalNumOfTraces() int64`

GetTotalNumOfTraces returns the TotalNumOfTraces field if non-nil, zero value otherwise.

### GetTotalNumOfTracesOk

`func (o *CpcQueryBucketResult) GetTotalNumOfTracesOk() (*int64, bool)`

GetTotalNumOfTracesOk returns a tuple with the TotalNumOfTraces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNumOfTraces

`func (o *CpcQueryBucketResult) SetTotalNumOfTraces(v int64)`

SetTotalNumOfTraces sets TotalNumOfTraces field to given value.


### GetAvgTraceDuration

`func (o *CpcQueryBucketResult) GetAvgTraceDuration() float64`

GetAvgTraceDuration returns the AvgTraceDuration field if non-nil, zero value otherwise.

### GetAvgTraceDurationOk

`func (o *CpcQueryBucketResult) GetAvgTraceDurationOk() (*float64, bool)`

GetAvgTraceDurationOk returns a tuple with the AvgTraceDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgTraceDuration

`func (o *CpcQueryBucketResult) SetAvgTraceDuration(v float64)`

SetAvgTraceDuration sets AvgTraceDuration field to given value.


### GetPerServiceCpcSummaries

`func (o *CpcQueryBucketResult) GetPerServiceCpcSummaries() []CpcServiceSummary`

GetPerServiceCpcSummaries returns the PerServiceCpcSummaries field if non-nil, zero value otherwise.

### GetPerServiceCpcSummariesOk

`func (o *CpcQueryBucketResult) GetPerServiceCpcSummariesOk() (*[]CpcServiceSummary, bool)`

GetPerServiceCpcSummariesOk returns a tuple with the PerServiceCpcSummaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerServiceCpcSummaries

`func (o *CpcQueryBucketResult) SetPerServiceCpcSummaries(v []CpcServiceSummary)`

SetPerServiceCpcSummaries sets PerServiceCpcSummaries field to given value.


### GetOtherServicesCpcSummary

`func (o *CpcQueryBucketResult) GetOtherServicesCpcSummary() CpcSummary`

GetOtherServicesCpcSummary returns the OtherServicesCpcSummary field if non-nil, zero value otherwise.

### GetOtherServicesCpcSummaryOk

`func (o *CpcQueryBucketResult) GetOtherServicesCpcSummaryOk() (*CpcSummary, bool)`

GetOtherServicesCpcSummaryOk returns a tuple with the OtherServicesCpcSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherServicesCpcSummary

`func (o *CpcQueryBucketResult) SetOtherServicesCpcSummary(v CpcSummary)`

SetOtherServicesCpcSummary sets OtherServicesCpcSummary field to given value.


### GetIdleTimeCpcSummary

`func (o *CpcQueryBucketResult) GetIdleTimeCpcSummary() CpcSummary`

GetIdleTimeCpcSummary returns the IdleTimeCpcSummary field if non-nil, zero value otherwise.

### GetIdleTimeCpcSummaryOk

`func (o *CpcQueryBucketResult) GetIdleTimeCpcSummaryOk() (*CpcSummary, bool)`

GetIdleTimeCpcSummaryOk returns a tuple with the IdleTimeCpcSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleTimeCpcSummary

`func (o *CpcQueryBucketResult) SetIdleTimeCpcSummary(v CpcSummary)`

SetIdleTimeCpcSummary sets IdleTimeCpcSummary field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


