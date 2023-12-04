# MetricsQueryResultInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RowId** | Pointer to **string** | Metrics Query row id. | [optional] 
**ResultContext** | Pointer to [**MetricsQueryResultContext**](MetricsQueryResultContext.md) |  | [optional] 

## Methods

### NewMetricsQueryResultInfo

`func NewMetricsQueryResultInfo() *MetricsQueryResultInfo`

NewMetricsQueryResultInfo instantiates a new MetricsQueryResultInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsQueryResultInfoWithDefaults

`func NewMetricsQueryResultInfoWithDefaults() *MetricsQueryResultInfo`

NewMetricsQueryResultInfoWithDefaults instantiates a new MetricsQueryResultInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRowId

`func (o *MetricsQueryResultInfo) GetRowId() string`

GetRowId returns the RowId field if non-nil, zero value otherwise.

### GetRowIdOk

`func (o *MetricsQueryResultInfo) GetRowIdOk() (*string, bool)`

GetRowIdOk returns a tuple with the RowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowId

`func (o *MetricsQueryResultInfo) SetRowId(v string)`

SetRowId sets RowId field to given value.

### HasRowId

`func (o *MetricsQueryResultInfo) HasRowId() bool`

HasRowId returns a boolean if a field has been set.

### GetResultContext

`func (o *MetricsQueryResultInfo) GetResultContext() MetricsQueryResultContext`

GetResultContext returns the ResultContext field if non-nil, zero value otherwise.

### GetResultContextOk

`func (o *MetricsQueryResultInfo) GetResultContextOk() (*MetricsQueryResultContext, bool)`

GetResultContextOk returns a tuple with the ResultContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultContext

`func (o *MetricsQueryResultInfo) SetResultContext(v MetricsQueryResultContext)`

SetResultContext sets ResultContext field to given value.

### HasResultContext

`func (o *MetricsQueryResultInfo) HasResultContext() bool`

HasResultContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


