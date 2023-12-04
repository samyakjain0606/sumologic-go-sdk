# QueryBasedSli

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryType** | **string** | Type of Raw Data Queries for SLI (Logs/Metrics). | 
**Queries** | [**[]SliQueryGroup**](SliQueryGroup.md) | Queries for defining SLI. | 

## Methods

### NewQueryBasedSli

`func NewQueryBasedSli(queryType string, queries []SliQueryGroup, ) *QueryBasedSli`

NewQueryBasedSli instantiates a new QueryBasedSli object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryBasedSliWithDefaults

`func NewQueryBasedSliWithDefaults() *QueryBasedSli`

NewQueryBasedSliWithDefaults instantiates a new QueryBasedSli object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueryType

`func (o *QueryBasedSli) GetQueryType() string`

GetQueryType returns the QueryType field if non-nil, zero value otherwise.

### GetQueryTypeOk

`func (o *QueryBasedSli) GetQueryTypeOk() (*string, bool)`

GetQueryTypeOk returns a tuple with the QueryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryType

`func (o *QueryBasedSli) SetQueryType(v string)`

SetQueryType sets QueryType field to given value.


### GetQueries

`func (o *QueryBasedSli) GetQueries() []SliQueryGroup`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *QueryBasedSli) GetQueriesOk() (*[]SliQueryGroup, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *QueryBasedSli) SetQueries(v []SliQueryGroup)`

SetQueries sets Queries field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


