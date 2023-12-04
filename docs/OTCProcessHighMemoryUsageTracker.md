# OTCProcessHighMemoryUsageTracker

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | Pointer to **string** | Event type. | [optional] 
**InstanceId** | Pointer to **string** | The collector instance ID, e.g. &#x60;974b444b-4b45-4f32-aa03-1dbf2a16826d&#x60;. | [optional] 
**InstanceAddress** | Pointer to **string** | The collector instance address, e.g. &#x60;172.16.1.14&#x60;. | [optional] 
**MemoryUsage** | Pointer to **string** | The collector memory usage in bytes, e.g. &#x60;142606592&#x60; | [optional] 
**MemoryLimit** | Pointer to **string** | The collector memory limit (if set) in bytes, e.g. &#x60;4000000000&#x60; | [optional] 

## Methods

### NewOTCProcessHighMemoryUsageTracker

`func NewOTCProcessHighMemoryUsageTracker() *OTCProcessHighMemoryUsageTracker`

NewOTCProcessHighMemoryUsageTracker instantiates a new OTCProcessHighMemoryUsageTracker object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOTCProcessHighMemoryUsageTrackerWithDefaults

`func NewOTCProcessHighMemoryUsageTrackerWithDefaults() *OTCProcessHighMemoryUsageTracker`

NewOTCProcessHighMemoryUsageTrackerWithDefaults instantiates a new OTCProcessHighMemoryUsageTracker object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *OTCProcessHighMemoryUsageTracker) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *OTCProcessHighMemoryUsageTracker) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *OTCProcessHighMemoryUsageTracker) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *OTCProcessHighMemoryUsageTracker) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetInstanceId

`func (o *OTCProcessHighMemoryUsageTracker) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *OTCProcessHighMemoryUsageTracker) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *OTCProcessHighMemoryUsageTracker) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *OTCProcessHighMemoryUsageTracker) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetInstanceAddress

`func (o *OTCProcessHighMemoryUsageTracker) GetInstanceAddress() string`

GetInstanceAddress returns the InstanceAddress field if non-nil, zero value otherwise.

### GetInstanceAddressOk

`func (o *OTCProcessHighMemoryUsageTracker) GetInstanceAddressOk() (*string, bool)`

GetInstanceAddressOk returns a tuple with the InstanceAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceAddress

`func (o *OTCProcessHighMemoryUsageTracker) SetInstanceAddress(v string)`

SetInstanceAddress sets InstanceAddress field to given value.

### HasInstanceAddress

`func (o *OTCProcessHighMemoryUsageTracker) HasInstanceAddress() bool`

HasInstanceAddress returns a boolean if a field has been set.

### GetMemoryUsage

`func (o *OTCProcessHighMemoryUsageTracker) GetMemoryUsage() string`

GetMemoryUsage returns the MemoryUsage field if non-nil, zero value otherwise.

### GetMemoryUsageOk

`func (o *OTCProcessHighMemoryUsageTracker) GetMemoryUsageOk() (*string, bool)`

GetMemoryUsageOk returns a tuple with the MemoryUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryUsage

`func (o *OTCProcessHighMemoryUsageTracker) SetMemoryUsage(v string)`

SetMemoryUsage sets MemoryUsage field to given value.

### HasMemoryUsage

`func (o *OTCProcessHighMemoryUsageTracker) HasMemoryUsage() bool`

HasMemoryUsage returns a boolean if a field has been set.

### GetMemoryLimit

`func (o *OTCProcessHighMemoryUsageTracker) GetMemoryLimit() string`

GetMemoryLimit returns the MemoryLimit field if non-nil, zero value otherwise.

### GetMemoryLimitOk

`func (o *OTCProcessHighMemoryUsageTracker) GetMemoryLimitOk() (*string, bool)`

GetMemoryLimitOk returns a tuple with the MemoryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryLimit

`func (o *OTCProcessHighMemoryUsageTracker) SetMemoryLimit(v string)`

SetMemoryLimit sets MemoryLimit field to given value.

### HasMemoryLimit

`func (o *OTCProcessHighMemoryUsageTracker) HasMemoryLimit() bool`

HasMemoryLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


