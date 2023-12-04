# DatastoreStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiskSize** | **int64** | Total DB size in terms of disk bytes | 
**IndicatorCount** | **int64** | Total number of indicators in the DB | 
**SourceStatus** | [**[]DatastoreSourceStatusResponse**](DatastoreSourceStatusResponse.md) | A list of sources and their individual DB sizes and indicator counts | 

## Methods

### NewDatastoreStatusResponse

`func NewDatastoreStatusResponse(diskSize int64, indicatorCount int64, sourceStatus []DatastoreSourceStatusResponse, ) *DatastoreStatusResponse`

NewDatastoreStatusResponse instantiates a new DatastoreStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatastoreStatusResponseWithDefaults

`func NewDatastoreStatusResponseWithDefaults() *DatastoreStatusResponse`

NewDatastoreStatusResponseWithDefaults instantiates a new DatastoreStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiskSize

`func (o *DatastoreStatusResponse) GetDiskSize() int64`

GetDiskSize returns the DiskSize field if non-nil, zero value otherwise.

### GetDiskSizeOk

`func (o *DatastoreStatusResponse) GetDiskSizeOk() (*int64, bool)`

GetDiskSizeOk returns a tuple with the DiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSize

`func (o *DatastoreStatusResponse) SetDiskSize(v int64)`

SetDiskSize sets DiskSize field to given value.


### GetIndicatorCount

`func (o *DatastoreStatusResponse) GetIndicatorCount() int64`

GetIndicatorCount returns the IndicatorCount field if non-nil, zero value otherwise.

### GetIndicatorCountOk

`func (o *DatastoreStatusResponse) GetIndicatorCountOk() (*int64, bool)`

GetIndicatorCountOk returns a tuple with the IndicatorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicatorCount

`func (o *DatastoreStatusResponse) SetIndicatorCount(v int64)`

SetIndicatorCount sets IndicatorCount field to given value.


### GetSourceStatus

`func (o *DatastoreStatusResponse) GetSourceStatus() []DatastoreSourceStatusResponse`

GetSourceStatus returns the SourceStatus field if non-nil, zero value otherwise.

### GetSourceStatusOk

`func (o *DatastoreStatusResponse) GetSourceStatusOk() (*[]DatastoreSourceStatusResponse, bool)`

GetSourceStatusOk returns a tuple with the SourceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceStatus

`func (o *DatastoreStatusResponse) SetSourceStatus(v []DatastoreSourceStatusResponse)`

SetSourceStatus sets SourceStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


