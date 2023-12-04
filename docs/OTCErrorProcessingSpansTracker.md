# OTCErrorProcessingSpansTracker

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | Pointer to **string** | Event type. | [optional] 
**InstanceId** | Pointer to **string** | The collector instance ID, e.g. &#x60;974b444b-4b45-4f32-aa03-1dbf2a16826d&#x60;. | [optional] 
**InstanceAddress** | Pointer to **string** | The collector instance address, e.g. &#x60;172.16.1.14&#x60;. | [optional] 
**ProcessorId** | Pointer to **string** | The collector processor ID, e.g. &#x60;cascading_filter&#x60;. | [optional] 
**Message** | Pointer to **string** | The error message. | [optional] 

## Methods

### NewOTCErrorProcessingSpansTracker

`func NewOTCErrorProcessingSpansTracker() *OTCErrorProcessingSpansTracker`

NewOTCErrorProcessingSpansTracker instantiates a new OTCErrorProcessingSpansTracker object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOTCErrorProcessingSpansTrackerWithDefaults

`func NewOTCErrorProcessingSpansTrackerWithDefaults() *OTCErrorProcessingSpansTracker`

NewOTCErrorProcessingSpansTrackerWithDefaults instantiates a new OTCErrorProcessingSpansTracker object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *OTCErrorProcessingSpansTracker) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *OTCErrorProcessingSpansTracker) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *OTCErrorProcessingSpansTracker) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *OTCErrorProcessingSpansTracker) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetInstanceId

`func (o *OTCErrorProcessingSpansTracker) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *OTCErrorProcessingSpansTracker) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *OTCErrorProcessingSpansTracker) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *OTCErrorProcessingSpansTracker) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetInstanceAddress

`func (o *OTCErrorProcessingSpansTracker) GetInstanceAddress() string`

GetInstanceAddress returns the InstanceAddress field if non-nil, zero value otherwise.

### GetInstanceAddressOk

`func (o *OTCErrorProcessingSpansTracker) GetInstanceAddressOk() (*string, bool)`

GetInstanceAddressOk returns a tuple with the InstanceAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceAddress

`func (o *OTCErrorProcessingSpansTracker) SetInstanceAddress(v string)`

SetInstanceAddress sets InstanceAddress field to given value.

### HasInstanceAddress

`func (o *OTCErrorProcessingSpansTracker) HasInstanceAddress() bool`

HasInstanceAddress returns a boolean if a field has been set.

### GetProcessorId

`func (o *OTCErrorProcessingSpansTracker) GetProcessorId() string`

GetProcessorId returns the ProcessorId field if non-nil, zero value otherwise.

### GetProcessorIdOk

`func (o *OTCErrorProcessingSpansTracker) GetProcessorIdOk() (*string, bool)`

GetProcessorIdOk returns a tuple with the ProcessorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorId

`func (o *OTCErrorProcessingSpansTracker) SetProcessorId(v string)`

SetProcessorId sets ProcessorId field to given value.

### HasProcessorId

`func (o *OTCErrorProcessingSpansTracker) HasProcessorId() bool`

HasProcessorId returns a boolean if a field has been set.

### GetMessage

`func (o *OTCErrorProcessingSpansTracker) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *OTCErrorProcessingSpansTracker) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *OTCErrorProcessingSpansTracker) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *OTCErrorProcessingSpansTracker) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


