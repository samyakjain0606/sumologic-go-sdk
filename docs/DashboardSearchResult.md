# DashboardSearchResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**DashboardSearchStatus**](DashboardSearchStatus.md) |  | 
**Axes** | [**VisualDataAxes**](VisualDataAxes.md) |  | 
**Series** | [**[]VisualDataSeries**](VisualDataSeries.md) | The series returned from a search. | 
**Errors** | Pointer to [**[]ErrorDescription**](ErrorDescription.md) | Errors returned by backend. | [optional] 
**TimeRange** | Pointer to [**BeginBoundedTimeRange**](BeginBoundedTimeRange.md) |  | [optional] 
**RequestToken** | Pointer to **string** | A user-generated string to uniquely identify the search request. This field can be safely ignored if you don&#39;t intend to identify a search request. | [optional] 
**FieldOrdering** | Pointer to **[]string** | The expected ordering of the column fields in tabular format. If null or empty, the ordering is unknown or indeterminate.  | [optional] 
**InfrequentScannedBytes** | Pointer to **float32** | The total number of scanned bytes from infrequent tier data for the query in bytes. | [optional] 
**ScannedBytes** | Pointer to [**ScannedBytes**](ScannedBytes.md) |  | [optional] 
**BackfillPercent** | Pointer to **float32** | The backfill percentage of a continuous query. | [optional] 

## Methods

### NewDashboardSearchResult

`func NewDashboardSearchResult(status DashboardSearchStatus, axes VisualDataAxes, series []VisualDataSeries, ) *DashboardSearchResult`

NewDashboardSearchResult instantiates a new DashboardSearchResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardSearchResultWithDefaults

`func NewDashboardSearchResultWithDefaults() *DashboardSearchResult`

NewDashboardSearchResultWithDefaults instantiates a new DashboardSearchResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *DashboardSearchResult) GetStatus() DashboardSearchStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DashboardSearchResult) GetStatusOk() (*DashboardSearchStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DashboardSearchResult) SetStatus(v DashboardSearchStatus)`

SetStatus sets Status field to given value.


### GetAxes

`func (o *DashboardSearchResult) GetAxes() VisualDataAxes`

GetAxes returns the Axes field if non-nil, zero value otherwise.

### GetAxesOk

`func (o *DashboardSearchResult) GetAxesOk() (*VisualDataAxes, bool)`

GetAxesOk returns a tuple with the Axes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAxes

`func (o *DashboardSearchResult) SetAxes(v VisualDataAxes)`

SetAxes sets Axes field to given value.


### GetSeries

`func (o *DashboardSearchResult) GetSeries() []VisualDataSeries`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *DashboardSearchResult) GetSeriesOk() (*[]VisualDataSeries, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *DashboardSearchResult) SetSeries(v []VisualDataSeries)`

SetSeries sets Series field to given value.


### GetErrors

`func (o *DashboardSearchResult) GetErrors() []ErrorDescription`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *DashboardSearchResult) GetErrorsOk() (*[]ErrorDescription, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *DashboardSearchResult) SetErrors(v []ErrorDescription)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *DashboardSearchResult) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetTimeRange

`func (o *DashboardSearchResult) GetTimeRange() BeginBoundedTimeRange`

GetTimeRange returns the TimeRange field if non-nil, zero value otherwise.

### GetTimeRangeOk

`func (o *DashboardSearchResult) GetTimeRangeOk() (*BeginBoundedTimeRange, bool)`

GetTimeRangeOk returns a tuple with the TimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRange

`func (o *DashboardSearchResult) SetTimeRange(v BeginBoundedTimeRange)`

SetTimeRange sets TimeRange field to given value.

### HasTimeRange

`func (o *DashboardSearchResult) HasTimeRange() bool`

HasTimeRange returns a boolean if a field has been set.

### GetRequestToken

`func (o *DashboardSearchResult) GetRequestToken() string`

GetRequestToken returns the RequestToken field if non-nil, zero value otherwise.

### GetRequestTokenOk

`func (o *DashboardSearchResult) GetRequestTokenOk() (*string, bool)`

GetRequestTokenOk returns a tuple with the RequestToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestToken

`func (o *DashboardSearchResult) SetRequestToken(v string)`

SetRequestToken sets RequestToken field to given value.

### HasRequestToken

`func (o *DashboardSearchResult) HasRequestToken() bool`

HasRequestToken returns a boolean if a field has been set.

### GetFieldOrdering

`func (o *DashboardSearchResult) GetFieldOrdering() []string`

GetFieldOrdering returns the FieldOrdering field if non-nil, zero value otherwise.

### GetFieldOrderingOk

`func (o *DashboardSearchResult) GetFieldOrderingOk() (*[]string, bool)`

GetFieldOrderingOk returns a tuple with the FieldOrdering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldOrdering

`func (o *DashboardSearchResult) SetFieldOrdering(v []string)`

SetFieldOrdering sets FieldOrdering field to given value.

### HasFieldOrdering

`func (o *DashboardSearchResult) HasFieldOrdering() bool`

HasFieldOrdering returns a boolean if a field has been set.

### GetInfrequentScannedBytes

`func (o *DashboardSearchResult) GetInfrequentScannedBytes() float32`

GetInfrequentScannedBytes returns the InfrequentScannedBytes field if non-nil, zero value otherwise.

### GetInfrequentScannedBytesOk

`func (o *DashboardSearchResult) GetInfrequentScannedBytesOk() (*float32, bool)`

GetInfrequentScannedBytesOk returns a tuple with the InfrequentScannedBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrequentScannedBytes

`func (o *DashboardSearchResult) SetInfrequentScannedBytes(v float32)`

SetInfrequentScannedBytes sets InfrequentScannedBytes field to given value.

### HasInfrequentScannedBytes

`func (o *DashboardSearchResult) HasInfrequentScannedBytes() bool`

HasInfrequentScannedBytes returns a boolean if a field has been set.

### GetScannedBytes

`func (o *DashboardSearchResult) GetScannedBytes() ScannedBytes`

GetScannedBytes returns the ScannedBytes field if non-nil, zero value otherwise.

### GetScannedBytesOk

`func (o *DashboardSearchResult) GetScannedBytesOk() (*ScannedBytes, bool)`

GetScannedBytesOk returns a tuple with the ScannedBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScannedBytes

`func (o *DashboardSearchResult) SetScannedBytes(v ScannedBytes)`

SetScannedBytes sets ScannedBytes field to given value.

### HasScannedBytes

`func (o *DashboardSearchResult) HasScannedBytes() bool`

HasScannedBytes returns a boolean if a field has been set.

### GetBackfillPercent

`func (o *DashboardSearchResult) GetBackfillPercent() float32`

GetBackfillPercent returns the BackfillPercent field if non-nil, zero value otherwise.

### GetBackfillPercentOk

`func (o *DashboardSearchResult) GetBackfillPercentOk() (*float32, bool)`

GetBackfillPercentOk returns a tuple with the BackfillPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackfillPercent

`func (o *DashboardSearchResult) SetBackfillPercent(v float32)`

SetBackfillPercent sets BackfillPercent field to given value.

### HasBackfillPercent

`func (o *DashboardSearchResult) HasBackfillPercent() bool`

HasBackfillPercent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


