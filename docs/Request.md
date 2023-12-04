# Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryType** | **string** | Type of Raw Data Queries for SLI (Logs/Metrics). | 
**Queries** | [**[]SliQueryGroup**](SliQueryGroup.md) | Queries for defining SLI. | 
**Threshold** | Pointer to **float32** | Compared against threshold query&#39;s raw data points to determine success. | [optional] 
**Op** | Pointer to **string** | Comparison function with threshold (LessThan/GreaterThan/LessThanOrEqual/GreaterThanOrEqual). | [optional] 

## Methods

### NewRequest

`func NewRequest(queryType string, queries []SliQueryGroup, ) *Request`

NewRequest instantiates a new Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestWithDefaults

`func NewRequestWithDefaults() *Request`

NewRequestWithDefaults instantiates a new Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueryType

`func (o *Request) GetQueryType() string`

GetQueryType returns the QueryType field if non-nil, zero value otherwise.

### GetQueryTypeOk

`func (o *Request) GetQueryTypeOk() (*string, bool)`

GetQueryTypeOk returns a tuple with the QueryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryType

`func (o *Request) SetQueryType(v string)`

SetQueryType sets QueryType field to given value.


### GetQueries

`func (o *Request) GetQueries() []SliQueryGroup`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *Request) GetQueriesOk() (*[]SliQueryGroup, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *Request) SetQueries(v []SliQueryGroup)`

SetQueries sets Queries field to given value.


### GetThreshold

`func (o *Request) GetThreshold() float32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *Request) GetThresholdOk() (*float32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *Request) SetThreshold(v float32)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *Request) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetOp

`func (o *Request) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *Request) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *Request) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *Request) HasOp() bool`

HasOp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


