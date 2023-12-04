# AggregationQueryResultResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Buckets** | [**[]AggregationQueryBucketResult**](AggregationQueryBucketResult.md) | A list of an aggregation query results on a per bucket basis.  Each bucket result corresponds to the number of trace query results falling into the bucket. | 

## Methods

### NewAggregationQueryResultResponse

`func NewAggregationQueryResultResponse(buckets []AggregationQueryBucketResult, ) *AggregationQueryResultResponse`

NewAggregationQueryResultResponse instantiates a new AggregationQueryResultResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregationQueryResultResponseWithDefaults

`func NewAggregationQueryResultResponseWithDefaults() *AggregationQueryResultResponse`

NewAggregationQueryResultResponseWithDefaults instantiates a new AggregationQueryResultResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuckets

`func (o *AggregationQueryResultResponse) GetBuckets() []AggregationQueryBucketResult`

GetBuckets returns the Buckets field if non-nil, zero value otherwise.

### GetBucketsOk

`func (o *AggregationQueryResultResponse) GetBucketsOk() (*[]AggregationQueryBucketResult, bool)`

GetBucketsOk returns a tuple with the Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuckets

`func (o *AggregationQueryResultResponse) SetBuckets(v []AggregationQueryBucketResult)`

SetBuckets sets Buckets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


