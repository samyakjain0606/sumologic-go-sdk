# CpcQueryResultResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Buckets** | [**[]CpcQueryBucketResult**](CpcQueryBucketResult.md) | A list of CPC query results on a per time bucket basis.  Each bucket result corresponds to the aggregated CPC data from a sample of traces matching search criteria falling within a specific time slice. | 

## Methods

### NewCpcQueryResultResponse

`func NewCpcQueryResultResponse(buckets []CpcQueryBucketResult, ) *CpcQueryResultResponse`

NewCpcQueryResultResponse instantiates a new CpcQueryResultResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCpcQueryResultResponseWithDefaults

`func NewCpcQueryResultResponseWithDefaults() *CpcQueryResultResponse`

NewCpcQueryResultResponseWithDefaults instantiates a new CpcQueryResultResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuckets

`func (o *CpcQueryResultResponse) GetBuckets() []CpcQueryBucketResult`

GetBuckets returns the Buckets field if non-nil, zero value otherwise.

### GetBucketsOk

`func (o *CpcQueryResultResponse) GetBucketsOk() (*[]CpcQueryBucketResult, bool)`

GetBucketsOk returns a tuple with the Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuckets

`func (o *CpcQueryResultResponse) SetBuckets(v []CpcQueryBucketResult)`

SetBuckets sets Buckets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


