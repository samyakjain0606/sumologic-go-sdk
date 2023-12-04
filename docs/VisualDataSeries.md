# VisualDataSeries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryId** | **string** | The id of the query. | 
**Name** | **string** | The meaning of &#39;name&#39; depends on the series type.   - For results of type &#39;timeseries&#39;, it is the value of the &#39;metric&#39; key.   - For results of type &#39;nontimeseries&#39;, it is the name of one of the fields that is not part of &#39;xAxisKeys&#39;.   - For results of type &#39;table&#39;, it is the comma-separated string of names of all fields.  | 
**DataPoints** | [**[]VisualPointData**](VisualPointData.md) | A list of data points in the visual series. | 
**AggregateInfo** | Pointer to [**VisualAggregateData**](VisualAggregateData.md) |  | [optional] 
**MetaData** | Pointer to [**VisualMetaData**](VisualMetaData.md) |  | [optional] 
**SeriesType** | Pointer to **string** | Type of the visual series. | [optional] 
**XAxisKeys** | Pointer to **[]string** | Keys that will be plotted as a point on the x axis. | [optional] 
**ValueType** | Pointer to **string** | Value that represents if the series values are String or Double | [optional] 
**Source** | Pointer to **string** | Source of the visual series. | [optional] 
**XAxisKeyTypes** | Pointer to **map[string]string** | Keys that will be plotted as a point on the x axis and their data type | [optional] [default to {}]
**QueryInfo** | Pointer to [**MetricsQueryResultInfo**](MetricsQueryResultInfo.md) |  | [optional] 

## Methods

### NewVisualDataSeries

`func NewVisualDataSeries(queryId string, name string, dataPoints []VisualPointData, ) *VisualDataSeries`

NewVisualDataSeries instantiates a new VisualDataSeries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVisualDataSeriesWithDefaults

`func NewVisualDataSeriesWithDefaults() *VisualDataSeries`

NewVisualDataSeriesWithDefaults instantiates a new VisualDataSeries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueryId

`func (o *VisualDataSeries) GetQueryId() string`

GetQueryId returns the QueryId field if non-nil, zero value otherwise.

### GetQueryIdOk

`func (o *VisualDataSeries) GetQueryIdOk() (*string, bool)`

GetQueryIdOk returns a tuple with the QueryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryId

`func (o *VisualDataSeries) SetQueryId(v string)`

SetQueryId sets QueryId field to given value.


### GetName

`func (o *VisualDataSeries) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VisualDataSeries) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VisualDataSeries) SetName(v string)`

SetName sets Name field to given value.


### GetDataPoints

`func (o *VisualDataSeries) GetDataPoints() []VisualPointData`

GetDataPoints returns the DataPoints field if non-nil, zero value otherwise.

### GetDataPointsOk

`func (o *VisualDataSeries) GetDataPointsOk() (*[]VisualPointData, bool)`

GetDataPointsOk returns a tuple with the DataPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataPoints

`func (o *VisualDataSeries) SetDataPoints(v []VisualPointData)`

SetDataPoints sets DataPoints field to given value.


### GetAggregateInfo

`func (o *VisualDataSeries) GetAggregateInfo() VisualAggregateData`

GetAggregateInfo returns the AggregateInfo field if non-nil, zero value otherwise.

### GetAggregateInfoOk

`func (o *VisualDataSeries) GetAggregateInfoOk() (*VisualAggregateData, bool)`

GetAggregateInfoOk returns a tuple with the AggregateInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregateInfo

`func (o *VisualDataSeries) SetAggregateInfo(v VisualAggregateData)`

SetAggregateInfo sets AggregateInfo field to given value.

### HasAggregateInfo

`func (o *VisualDataSeries) HasAggregateInfo() bool`

HasAggregateInfo returns a boolean if a field has been set.

### GetMetaData

`func (o *VisualDataSeries) GetMetaData() VisualMetaData`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *VisualDataSeries) GetMetaDataOk() (*VisualMetaData, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *VisualDataSeries) SetMetaData(v VisualMetaData)`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *VisualDataSeries) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.

### GetSeriesType

`func (o *VisualDataSeries) GetSeriesType() string`

GetSeriesType returns the SeriesType field if non-nil, zero value otherwise.

### GetSeriesTypeOk

`func (o *VisualDataSeries) GetSeriesTypeOk() (*string, bool)`

GetSeriesTypeOk returns a tuple with the SeriesType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesType

`func (o *VisualDataSeries) SetSeriesType(v string)`

SetSeriesType sets SeriesType field to given value.

### HasSeriesType

`func (o *VisualDataSeries) HasSeriesType() bool`

HasSeriesType returns a boolean if a field has been set.

### GetXAxisKeys

`func (o *VisualDataSeries) GetXAxisKeys() []string`

GetXAxisKeys returns the XAxisKeys field if non-nil, zero value otherwise.

### GetXAxisKeysOk

`func (o *VisualDataSeries) GetXAxisKeysOk() (*[]string, bool)`

GetXAxisKeysOk returns a tuple with the XAxisKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXAxisKeys

`func (o *VisualDataSeries) SetXAxisKeys(v []string)`

SetXAxisKeys sets XAxisKeys field to given value.

### HasXAxisKeys

`func (o *VisualDataSeries) HasXAxisKeys() bool`

HasXAxisKeys returns a boolean if a field has been set.

### GetValueType

`func (o *VisualDataSeries) GetValueType() string`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *VisualDataSeries) GetValueTypeOk() (*string, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *VisualDataSeries) SetValueType(v string)`

SetValueType sets ValueType field to given value.

### HasValueType

`func (o *VisualDataSeries) HasValueType() bool`

HasValueType returns a boolean if a field has been set.

### GetSource

`func (o *VisualDataSeries) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *VisualDataSeries) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *VisualDataSeries) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *VisualDataSeries) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetXAxisKeyTypes

`func (o *VisualDataSeries) GetXAxisKeyTypes() map[string]string`

GetXAxisKeyTypes returns the XAxisKeyTypes field if non-nil, zero value otherwise.

### GetXAxisKeyTypesOk

`func (o *VisualDataSeries) GetXAxisKeyTypesOk() (*map[string]string, bool)`

GetXAxisKeyTypesOk returns a tuple with the XAxisKeyTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXAxisKeyTypes

`func (o *VisualDataSeries) SetXAxisKeyTypes(v map[string]string)`

SetXAxisKeyTypes sets XAxisKeyTypes field to given value.

### HasXAxisKeyTypes

`func (o *VisualDataSeries) HasXAxisKeyTypes() bool`

HasXAxisKeyTypes returns a boolean if a field has been set.

### GetQueryInfo

`func (o *VisualDataSeries) GetQueryInfo() MetricsQueryResultInfo`

GetQueryInfo returns the QueryInfo field if non-nil, zero value otherwise.

### GetQueryInfoOk

`func (o *VisualDataSeries) GetQueryInfoOk() (*MetricsQueryResultInfo, bool)`

GetQueryInfoOk returns a tuple with the QueryInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryInfo

`func (o *VisualDataSeries) SetQueryInfo(v MetricsQueryResultInfo)`

SetQueryInfo sets QueryInfo field to given value.

### HasQueryInfo

`func (o *VisualDataSeries) HasQueryInfo() bool`

HasQueryInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


