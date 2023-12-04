# CpcQueryResultRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketIds** | **[]string** | A list of the identifiers of CPC query buckets for which aggregated Critical Path Contribution data should be fetched. | 

## Methods

### NewCpcQueryResultRequest

`func NewCpcQueryResultRequest(bucketIds []string, ) *CpcQueryResultRequest`

NewCpcQueryResultRequest instantiates a new CpcQueryResultRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCpcQueryResultRequestWithDefaults

`func NewCpcQueryResultRequestWithDefaults() *CpcQueryResultRequest`

NewCpcQueryResultRequestWithDefaults instantiates a new CpcQueryResultRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketIds

`func (o *CpcQueryResultRequest) GetBucketIds() []string`

GetBucketIds returns the BucketIds field if non-nil, zero value otherwise.

### GetBucketIdsOk

`func (o *CpcQueryResultRequest) GetBucketIdsOk() (*[]string, bool)`

GetBucketIdsOk returns a tuple with the BucketIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketIds

`func (o *CpcQueryResultRequest) SetBucketIds(v []string)`

SetBucketIds sets BucketIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


