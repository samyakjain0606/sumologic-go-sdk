# BucketValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketValueType** | **string** | Bucket value type of the object model. | 
**TraceCount** | **int64** | The number of traces per bucket. | 

## Methods

### NewBucketValue

`func NewBucketValue(bucketValueType string, traceCount int64, ) *BucketValue`

NewBucketValue instantiates a new BucketValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBucketValueWithDefaults

`func NewBucketValueWithDefaults() *BucketValue`

NewBucketValueWithDefaults instantiates a new BucketValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketValueType

`func (o *BucketValue) GetBucketValueType() string`

GetBucketValueType returns the BucketValueType field if non-nil, zero value otherwise.

### GetBucketValueTypeOk

`func (o *BucketValue) GetBucketValueTypeOk() (*string, bool)`

GetBucketValueTypeOk returns a tuple with the BucketValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketValueType

`func (o *BucketValue) SetBucketValueType(v string)`

SetBucketValueType sets BucketValueType field to given value.


### GetTraceCount

`func (o *BucketValue) GetTraceCount() int64`

GetTraceCount returns the TraceCount field if non-nil, zero value otherwise.

### GetTraceCountOk

`func (o *BucketValue) GetTraceCountOk() (*int64, bool)`

GetTraceCountOk returns a tuple with the TraceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceCount

`func (o *BucketValue) SetTraceCount(v int64)`

SetTraceCount sets TraceCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


