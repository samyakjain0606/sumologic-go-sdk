# AnomalyDataPointsCountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RelativeTimeRange** | **string** | The relative time range of the monitor. Valid values of time ranges are &#x60;-5m&#x60;, &#x60;-10m&#x60;, &#x60;-15m&#x60;, &#x60;-30m&#x60;, &#x60;-1h&#x60;, &#x60;-3h&#x60;, &#x60;-6h&#x60;, &#x60;-12h&#x60;, or &#x60;-24h&#x60;. | 
**Queries** | [**[]MonitorQuery**](MonitorQuery.md) | All queries from the monitor. | 

## Methods

### NewAnomalyDataPointsCountRequest

`func NewAnomalyDataPointsCountRequest(relativeTimeRange string, queries []MonitorQuery, ) *AnomalyDataPointsCountRequest`

NewAnomalyDataPointsCountRequest instantiates a new AnomalyDataPointsCountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnomalyDataPointsCountRequestWithDefaults

`func NewAnomalyDataPointsCountRequestWithDefaults() *AnomalyDataPointsCountRequest`

NewAnomalyDataPointsCountRequestWithDefaults instantiates a new AnomalyDataPointsCountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRelativeTimeRange

`func (o *AnomalyDataPointsCountRequest) GetRelativeTimeRange() string`

GetRelativeTimeRange returns the RelativeTimeRange field if non-nil, zero value otherwise.

### GetRelativeTimeRangeOk

`func (o *AnomalyDataPointsCountRequest) GetRelativeTimeRangeOk() (*string, bool)`

GetRelativeTimeRangeOk returns a tuple with the RelativeTimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativeTimeRange

`func (o *AnomalyDataPointsCountRequest) SetRelativeTimeRange(v string)`

SetRelativeTimeRange sets RelativeTimeRange field to given value.


### GetQueries

`func (o *AnomalyDataPointsCountRequest) GetQueries() []MonitorQuery`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *AnomalyDataPointsCountRequest) GetQueriesOk() (*[]MonitorQuery, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *AnomalyDataPointsCountRequest) SetQueries(v []MonitorQuery)`

SetQueries sets Queries field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


