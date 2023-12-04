# SpanQueryAggregateDataSeries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryId** | **string** | The id of the query. | 
**Name** | **string** | The meaning of &#39;name&#39; depends on the series type.   - For results of type &#39;timeseries&#39;, it is the value of the x axis &#39;field&#39; key.   - For results of type &#39;nontimeseries&#39;, it is the name of one of the fields that is not part of &#39;xAxisKeys&#39;.   - For results of type &#39;table&#39;, it is the comma-separated string of names of all fields.  | 
**DataPoints** | [**[]SpanQueryAggregatePointData**](SpanQueryAggregatePointData.md) | A list of data points in the series. | 
**AggregateInfo** | Pointer to [**SpanQueryAggregateAggregateData**](SpanQueryAggregateAggregateData.md) |  | [optional] 
**MetaData** | Pointer to [**SpanQueryAggregateMetaData**](SpanQueryAggregateMetaData.md) |  | [optional] 
**SeriesType** | Pointer to **string** | Type of the visual series. | [optional] 
**XAxisKeys** | Pointer to **[]string** | Keys that will be plotted as a point on the x axis. | [optional] 
**ValueType** | Pointer to **string** | Type of the values in the series. | [optional] 

## Methods

### NewSpanQueryAggregateDataSeries

`func NewSpanQueryAggregateDataSeries(queryId string, name string, dataPoints []SpanQueryAggregatePointData, ) *SpanQueryAggregateDataSeries`

NewSpanQueryAggregateDataSeries instantiates a new SpanQueryAggregateDataSeries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpanQueryAggregateDataSeriesWithDefaults

`func NewSpanQueryAggregateDataSeriesWithDefaults() *SpanQueryAggregateDataSeries`

NewSpanQueryAggregateDataSeriesWithDefaults instantiates a new SpanQueryAggregateDataSeries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueryId

`func (o *SpanQueryAggregateDataSeries) GetQueryId() string`

GetQueryId returns the QueryId field if non-nil, zero value otherwise.

### GetQueryIdOk

`func (o *SpanQueryAggregateDataSeries) GetQueryIdOk() (*string, bool)`

GetQueryIdOk returns a tuple with the QueryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryId

`func (o *SpanQueryAggregateDataSeries) SetQueryId(v string)`

SetQueryId sets QueryId field to given value.


### GetName

`func (o *SpanQueryAggregateDataSeries) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SpanQueryAggregateDataSeries) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SpanQueryAggregateDataSeries) SetName(v string)`

SetName sets Name field to given value.


### GetDataPoints

`func (o *SpanQueryAggregateDataSeries) GetDataPoints() []SpanQueryAggregatePointData`

GetDataPoints returns the DataPoints field if non-nil, zero value otherwise.

### GetDataPointsOk

`func (o *SpanQueryAggregateDataSeries) GetDataPointsOk() (*[]SpanQueryAggregatePointData, bool)`

GetDataPointsOk returns a tuple with the DataPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataPoints

`func (o *SpanQueryAggregateDataSeries) SetDataPoints(v []SpanQueryAggregatePointData)`

SetDataPoints sets DataPoints field to given value.


### GetAggregateInfo

`func (o *SpanQueryAggregateDataSeries) GetAggregateInfo() SpanQueryAggregateAggregateData`

GetAggregateInfo returns the AggregateInfo field if non-nil, zero value otherwise.

### GetAggregateInfoOk

`func (o *SpanQueryAggregateDataSeries) GetAggregateInfoOk() (*SpanQueryAggregateAggregateData, bool)`

GetAggregateInfoOk returns a tuple with the AggregateInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregateInfo

`func (o *SpanQueryAggregateDataSeries) SetAggregateInfo(v SpanQueryAggregateAggregateData)`

SetAggregateInfo sets AggregateInfo field to given value.

### HasAggregateInfo

`func (o *SpanQueryAggregateDataSeries) HasAggregateInfo() bool`

HasAggregateInfo returns a boolean if a field has been set.

### GetMetaData

`func (o *SpanQueryAggregateDataSeries) GetMetaData() SpanQueryAggregateMetaData`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *SpanQueryAggregateDataSeries) GetMetaDataOk() (*SpanQueryAggregateMetaData, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *SpanQueryAggregateDataSeries) SetMetaData(v SpanQueryAggregateMetaData)`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *SpanQueryAggregateDataSeries) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.

### GetSeriesType

`func (o *SpanQueryAggregateDataSeries) GetSeriesType() string`

GetSeriesType returns the SeriesType field if non-nil, zero value otherwise.

### GetSeriesTypeOk

`func (o *SpanQueryAggregateDataSeries) GetSeriesTypeOk() (*string, bool)`

GetSeriesTypeOk returns a tuple with the SeriesType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesType

`func (o *SpanQueryAggregateDataSeries) SetSeriesType(v string)`

SetSeriesType sets SeriesType field to given value.

### HasSeriesType

`func (o *SpanQueryAggregateDataSeries) HasSeriesType() bool`

HasSeriesType returns a boolean if a field has been set.

### GetXAxisKeys

`func (o *SpanQueryAggregateDataSeries) GetXAxisKeys() []string`

GetXAxisKeys returns the XAxisKeys field if non-nil, zero value otherwise.

### GetXAxisKeysOk

`func (o *SpanQueryAggregateDataSeries) GetXAxisKeysOk() (*[]string, bool)`

GetXAxisKeysOk returns a tuple with the XAxisKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXAxisKeys

`func (o *SpanQueryAggregateDataSeries) SetXAxisKeys(v []string)`

SetXAxisKeys sets XAxisKeys field to given value.

### HasXAxisKeys

`func (o *SpanQueryAggregateDataSeries) HasXAxisKeys() bool`

HasXAxisKeys returns a boolean if a field has been set.

### GetValueType

`func (o *SpanQueryAggregateDataSeries) GetValueType() string`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *SpanQueryAggregateDataSeries) GetValueTypeOk() (*string, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *SpanQueryAggregateDataSeries) SetValueType(v string)`

SetValueType sets ValueType field to given value.

### HasValueType

`func (o *SpanQueryAggregateDataSeries) HasValueType() bool`

HasValueType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


