# DatastoreSourceStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | **string** | The indicator source | 
**DiskSize** | **int64** | Disk utilization in bytes estimate for the indicator source | 
**IndicatorCount** | **int64** | Number of indicators for the indicator source | 

## Methods

### NewDatastoreSourceStatusResponse

`func NewDatastoreSourceStatusResponse(source string, diskSize int64, indicatorCount int64, ) *DatastoreSourceStatusResponse`

NewDatastoreSourceStatusResponse instantiates a new DatastoreSourceStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatastoreSourceStatusResponseWithDefaults

`func NewDatastoreSourceStatusResponseWithDefaults() *DatastoreSourceStatusResponse`

NewDatastoreSourceStatusResponseWithDefaults instantiates a new DatastoreSourceStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *DatastoreSourceStatusResponse) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *DatastoreSourceStatusResponse) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *DatastoreSourceStatusResponse) SetSource(v string)`

SetSource sets Source field to given value.


### GetDiskSize

`func (o *DatastoreSourceStatusResponse) GetDiskSize() int64`

GetDiskSize returns the DiskSize field if non-nil, zero value otherwise.

### GetDiskSizeOk

`func (o *DatastoreSourceStatusResponse) GetDiskSizeOk() (*int64, bool)`

GetDiskSizeOk returns a tuple with the DiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSize

`func (o *DatastoreSourceStatusResponse) SetDiskSize(v int64)`

SetDiskSize sets DiskSize field to given value.


### GetIndicatorCount

`func (o *DatastoreSourceStatusResponse) GetIndicatorCount() int64`

GetIndicatorCount returns the IndicatorCount field if non-nil, zero value otherwise.

### GetIndicatorCountOk

`func (o *DatastoreSourceStatusResponse) GetIndicatorCountOk() (*int64, bool)`

GetIndicatorCountOk returns a tuple with the IndicatorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicatorCount

`func (o *DatastoreSourceStatusResponse) SetIndicatorCount(v int64)`

SetIndicatorCount sets IndicatorCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


