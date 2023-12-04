# CpcQueryRowStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RowId** | **string** | A unique identifier of the query. | 
**Buckets** | [**[]CpcQueryBucketStatus**](CpcQueryBucketStatus.md) | A list of CPC query statuses on a per time bucket basis.  Each status corresponds to the status of calculating aggregated CPC data from a sample of traces  matching search criteria falling within a specific time slice. | 
**Status** | **string** | Status of the query. Possible values: &#x60;Processing&#x60;, &#x60;Finished&#x60;, &#x60;Error&#x60;, &#x60;Canceled&#x60;. | 

## Methods

### NewCpcQueryRowStatus

`func NewCpcQueryRowStatus(rowId string, buckets []CpcQueryBucketStatus, status string, ) *CpcQueryRowStatus`

NewCpcQueryRowStatus instantiates a new CpcQueryRowStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCpcQueryRowStatusWithDefaults

`func NewCpcQueryRowStatusWithDefaults() *CpcQueryRowStatus`

NewCpcQueryRowStatusWithDefaults instantiates a new CpcQueryRowStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRowId

`func (o *CpcQueryRowStatus) GetRowId() string`

GetRowId returns the RowId field if non-nil, zero value otherwise.

### GetRowIdOk

`func (o *CpcQueryRowStatus) GetRowIdOk() (*string, bool)`

GetRowIdOk returns a tuple with the RowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowId

`func (o *CpcQueryRowStatus) SetRowId(v string)`

SetRowId sets RowId field to given value.


### GetBuckets

`func (o *CpcQueryRowStatus) GetBuckets() []CpcQueryBucketStatus`

GetBuckets returns the Buckets field if non-nil, zero value otherwise.

### GetBucketsOk

`func (o *CpcQueryRowStatus) GetBucketsOk() (*[]CpcQueryBucketStatus, bool)`

GetBucketsOk returns a tuple with the Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuckets

`func (o *CpcQueryRowStatus) SetBuckets(v []CpcQueryBucketStatus)`

SetBuckets sets Buckets field to given value.


### GetStatus

`func (o *CpcQueryRowStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CpcQueryRowStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CpcQueryRowStatus) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


