# SpanQueryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryRows** | [**[]SpanQueryRow**](SpanQueryRow.md) | A list of span analytics queries. | 
**TimeRange** | [**ResolvableTimeRange**](ResolvableTimeRange.md) |  | 
**TimeZone** | Pointer to **string** | Time zone for the query time ranges. Follow the format in the [IANA Time Zone Database](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones#List). | [optional] [default to "UTC"]

## Methods

### NewSpanQueryRequest

`func NewSpanQueryRequest(queryRows []SpanQueryRow, timeRange ResolvableTimeRange, ) *SpanQueryRequest`

NewSpanQueryRequest instantiates a new SpanQueryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpanQueryRequestWithDefaults

`func NewSpanQueryRequestWithDefaults() *SpanQueryRequest`

NewSpanQueryRequestWithDefaults instantiates a new SpanQueryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueryRows

`func (o *SpanQueryRequest) GetQueryRows() []SpanQueryRow`

GetQueryRows returns the QueryRows field if non-nil, zero value otherwise.

### GetQueryRowsOk

`func (o *SpanQueryRequest) GetQueryRowsOk() (*[]SpanQueryRow, bool)`

GetQueryRowsOk returns a tuple with the QueryRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryRows

`func (o *SpanQueryRequest) SetQueryRows(v []SpanQueryRow)`

SetQueryRows sets QueryRows field to given value.


### GetTimeRange

`func (o *SpanQueryRequest) GetTimeRange() ResolvableTimeRange`

GetTimeRange returns the TimeRange field if non-nil, zero value otherwise.

### GetTimeRangeOk

`func (o *SpanQueryRequest) GetTimeRangeOk() (*ResolvableTimeRange, bool)`

GetTimeRangeOk returns a tuple with the TimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRange

`func (o *SpanQueryRequest) SetTimeRange(v ResolvableTimeRange)`

SetTimeRange sets TimeRange field to given value.


### GetTimeZone

`func (o *SpanQueryRequest) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *SpanQueryRequest) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *SpanQueryRequest) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *SpanQueryRequest) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


