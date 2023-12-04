# OTCReceiverSpansRefusedTracker

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | Pointer to **string** | Event type. | [optional] 
**InstanceId** | Pointer to **string** | The collector instance ID, e.g. &#x60;974b444b-4b45-4f32-aa03-1dbf2a16826d&#x60;. | [optional] 
**InstanceAddress** | Pointer to **string** | The collector instance address, e.g. &#x60;172.16.1.14&#x60;. | [optional] 
**ReceiverId** | Pointer to **string** | The collector receiver ID, e.g. &#x60;otlphttp/2&#x60;. | [optional] 
**Count** | Pointer to **string** | The count of refused spans. | [optional] 

## Methods

### NewOTCReceiverSpansRefusedTracker

`func NewOTCReceiverSpansRefusedTracker() *OTCReceiverSpansRefusedTracker`

NewOTCReceiverSpansRefusedTracker instantiates a new OTCReceiverSpansRefusedTracker object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOTCReceiverSpansRefusedTrackerWithDefaults

`func NewOTCReceiverSpansRefusedTrackerWithDefaults() *OTCReceiverSpansRefusedTracker`

NewOTCReceiverSpansRefusedTrackerWithDefaults instantiates a new OTCReceiverSpansRefusedTracker object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *OTCReceiverSpansRefusedTracker) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *OTCReceiverSpansRefusedTracker) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *OTCReceiverSpansRefusedTracker) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *OTCReceiverSpansRefusedTracker) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetInstanceId

`func (o *OTCReceiverSpansRefusedTracker) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *OTCReceiverSpansRefusedTracker) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *OTCReceiverSpansRefusedTracker) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *OTCReceiverSpansRefusedTracker) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetInstanceAddress

`func (o *OTCReceiverSpansRefusedTracker) GetInstanceAddress() string`

GetInstanceAddress returns the InstanceAddress field if non-nil, zero value otherwise.

### GetInstanceAddressOk

`func (o *OTCReceiverSpansRefusedTracker) GetInstanceAddressOk() (*string, bool)`

GetInstanceAddressOk returns a tuple with the InstanceAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceAddress

`func (o *OTCReceiverSpansRefusedTracker) SetInstanceAddress(v string)`

SetInstanceAddress sets InstanceAddress field to given value.

### HasInstanceAddress

`func (o *OTCReceiverSpansRefusedTracker) HasInstanceAddress() bool`

HasInstanceAddress returns a boolean if a field has been set.

### GetReceiverId

`func (o *OTCReceiverSpansRefusedTracker) GetReceiverId() string`

GetReceiverId returns the ReceiverId field if non-nil, zero value otherwise.

### GetReceiverIdOk

`func (o *OTCReceiverSpansRefusedTracker) GetReceiverIdOk() (*string, bool)`

GetReceiverIdOk returns a tuple with the ReceiverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverId

`func (o *OTCReceiverSpansRefusedTracker) SetReceiverId(v string)`

SetReceiverId sets ReceiverId field to given value.

### HasReceiverId

`func (o *OTCReceiverSpansRefusedTracker) HasReceiverId() bool`

HasReceiverId returns a boolean if a field has been set.

### GetCount

`func (o *OTCReceiverSpansRefusedTracker) GetCount() string`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *OTCReceiverSpansRefusedTracker) GetCountOk() (*string, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *OTCReceiverSpansRefusedTracker) SetCount(v string)`

SetCount sets Count field to given value.

### HasCount

`func (o *OTCReceiverSpansRefusedTracker) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


