# SloUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SliType** | Pointer to **string** | The type of SLO usage info (Logs/Metrics/Monitor based). | [optional] 
**Usage** | Pointer to **int32** | Current number of active Logs/Metrics/Monitors SLOs. | [optional] 
**Limit** | Pointer to **int32** | The limit of active Logs/Metrics/Monitors SLOs. | [optional] 

## Methods

### NewSloUsage

`func NewSloUsage() *SloUsage`

NewSloUsage instantiates a new SloUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSloUsageWithDefaults

`func NewSloUsageWithDefaults() *SloUsage`

NewSloUsageWithDefaults instantiates a new SloUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSliType

`func (o *SloUsage) GetSliType() string`

GetSliType returns the SliType field if non-nil, zero value otherwise.

### GetSliTypeOk

`func (o *SloUsage) GetSliTypeOk() (*string, bool)`

GetSliTypeOk returns a tuple with the SliType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSliType

`func (o *SloUsage) SetSliType(v string)`

SetSliType sets SliType field to given value.

### HasSliType

`func (o *SloUsage) HasSliType() bool`

HasSliType returns a boolean if a field has been set.

### GetUsage

`func (o *SloUsage) GetUsage() int32`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *SloUsage) GetUsageOk() (*int32, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *SloUsage) SetUsage(v int32)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *SloUsage) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetLimit

`func (o *SloUsage) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *SloUsage) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *SloUsage) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *SloUsage) HasLimit() bool`

HasLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


