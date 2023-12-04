# CpcQueryBucketStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketId** | **string** | A unique identifier of the bucket. | 
**Status** | **string** | Status of the query. Possible values: &#x60;Processing&#x60;, &#x60;Finished&#x60;, &#x60;Error&#x60;, &#x60;Canceled&#x60;. | 
**StatusMessage** | Pointer to **string** | Descriptive message of the status. | [optional] 

## Methods

### NewCpcQueryBucketStatus

`func NewCpcQueryBucketStatus(bucketId string, status string, ) *CpcQueryBucketStatus`

NewCpcQueryBucketStatus instantiates a new CpcQueryBucketStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCpcQueryBucketStatusWithDefaults

`func NewCpcQueryBucketStatusWithDefaults() *CpcQueryBucketStatus`

NewCpcQueryBucketStatusWithDefaults instantiates a new CpcQueryBucketStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketId

`func (o *CpcQueryBucketStatus) GetBucketId() string`

GetBucketId returns the BucketId field if non-nil, zero value otherwise.

### GetBucketIdOk

`func (o *CpcQueryBucketStatus) GetBucketIdOk() (*string, bool)`

GetBucketIdOk returns a tuple with the BucketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketId

`func (o *CpcQueryBucketStatus) SetBucketId(v string)`

SetBucketId sets BucketId field to given value.


### GetStatus

`func (o *CpcQueryBucketStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CpcQueryBucketStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CpcQueryBucketStatus) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStatusMessage

`func (o *CpcQueryBucketStatus) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *CpcQueryBucketStatus) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *CpcQueryBucketStatus) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *CpcQueryBucketStatus) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


