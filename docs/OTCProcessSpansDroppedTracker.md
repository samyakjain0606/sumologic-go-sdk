# OTCProcessSpansDroppedTracker

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | Pointer to **string** | Event type. | [optional] 
**InstanceId** | Pointer to **string** | The collector instance ID, e.g. &#x60;974b444b-4b45-4f32-aa03-1dbf2a16826d&#x60;. | [optional] 
**InstanceAddress** | Pointer to **string** | The collector instance address, e.g. &#x60;172.16.1.14&#x60;. | [optional] 
**ProcessorId** | Pointer to **string** | The collector processor ID, e.g. &#x60;cascading_filter&#x60;. | [optional] 
**Count** | Pointer to **string** | The count of dropped spans. | [optional] 

## Methods

### NewOTCProcessSpansDroppedTracker

`func NewOTCProcessSpansDroppedTracker() *OTCProcessSpansDroppedTracker`

NewOTCProcessSpansDroppedTracker instantiates a new OTCProcessSpansDroppedTracker object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOTCProcessSpansDroppedTrackerWithDefaults

`func NewOTCProcessSpansDroppedTrackerWithDefaults() *OTCProcessSpansDroppedTracker`

NewOTCProcessSpansDroppedTrackerWithDefaults instantiates a new OTCProcessSpansDroppedTracker object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *OTCProcessSpansDroppedTracker) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *OTCProcessSpansDroppedTracker) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *OTCProcessSpansDroppedTracker) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *OTCProcessSpansDroppedTracker) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetInstanceId

`func (o *OTCProcessSpansDroppedTracker) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *OTCProcessSpansDroppedTracker) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *OTCProcessSpansDroppedTracker) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *OTCProcessSpansDroppedTracker) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetInstanceAddress

`func (o *OTCProcessSpansDroppedTracker) GetInstanceAddress() string`

GetInstanceAddress returns the InstanceAddress field if non-nil, zero value otherwise.

### GetInstanceAddressOk

`func (o *OTCProcessSpansDroppedTracker) GetInstanceAddressOk() (*string, bool)`

GetInstanceAddressOk returns a tuple with the InstanceAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceAddress

`func (o *OTCProcessSpansDroppedTracker) SetInstanceAddress(v string)`

SetInstanceAddress sets InstanceAddress field to given value.

### HasInstanceAddress

`func (o *OTCProcessSpansDroppedTracker) HasInstanceAddress() bool`

HasInstanceAddress returns a boolean if a field has been set.

### GetProcessorId

`func (o *OTCProcessSpansDroppedTracker) GetProcessorId() string`

GetProcessorId returns the ProcessorId field if non-nil, zero value otherwise.

### GetProcessorIdOk

`func (o *OTCProcessSpansDroppedTracker) GetProcessorIdOk() (*string, bool)`

GetProcessorIdOk returns a tuple with the ProcessorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorId

`func (o *OTCProcessSpansDroppedTracker) SetProcessorId(v string)`

SetProcessorId sets ProcessorId field to given value.

### HasProcessorId

`func (o *OTCProcessSpansDroppedTracker) HasProcessorId() bool`

HasProcessorId returns a boolean if a field has been set.

### GetCount

`func (o *OTCProcessSpansDroppedTracker) GetCount() string`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *OTCProcessSpansDroppedTracker) GetCountOk() (*string, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *OTCProcessSpansDroppedTracker) SetCount(v string)`

SetCount sets Count field to given value.

### HasCount

`func (o *OTCProcessSpansDroppedTracker) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


