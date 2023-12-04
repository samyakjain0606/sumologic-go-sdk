# CorrelatedEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the event. | 
**EventType** | **string** | The type of event. | 
**EventSource** | **string** | The source of the event. | 
**StartTimeMs** | **int64** | The start time of the event. | 
**MetadataJson** | Pointer to **string** | JSON string containing metadata of the event. | [optional] 

## Methods

### NewCorrelatedEvent

`func NewCorrelatedEvent(name string, eventType string, eventSource string, startTimeMs int64, ) *CorrelatedEvent`

NewCorrelatedEvent instantiates a new CorrelatedEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCorrelatedEventWithDefaults

`func NewCorrelatedEventWithDefaults() *CorrelatedEvent`

NewCorrelatedEventWithDefaults instantiates a new CorrelatedEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CorrelatedEvent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CorrelatedEvent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CorrelatedEvent) SetName(v string)`

SetName sets Name field to given value.


### GetEventType

`func (o *CorrelatedEvent) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *CorrelatedEvent) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *CorrelatedEvent) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetEventSource

`func (o *CorrelatedEvent) GetEventSource() string`

GetEventSource returns the EventSource field if non-nil, zero value otherwise.

### GetEventSourceOk

`func (o *CorrelatedEvent) GetEventSourceOk() (*string, bool)`

GetEventSourceOk returns a tuple with the EventSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSource

`func (o *CorrelatedEvent) SetEventSource(v string)`

SetEventSource sets EventSource field to given value.


### GetStartTimeMs

`func (o *CorrelatedEvent) GetStartTimeMs() int64`

GetStartTimeMs returns the StartTimeMs field if non-nil, zero value otherwise.

### GetStartTimeMsOk

`func (o *CorrelatedEvent) GetStartTimeMsOk() (*int64, bool)`

GetStartTimeMsOk returns a tuple with the StartTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeMs

`func (o *CorrelatedEvent) SetStartTimeMs(v int64)`

SetStartTimeMs sets StartTimeMs field to given value.


### GetMetadataJson

`func (o *CorrelatedEvent) GetMetadataJson() string`

GetMetadataJson returns the MetadataJson field if non-nil, zero value otherwise.

### GetMetadataJsonOk

`func (o *CorrelatedEvent) GetMetadataJsonOk() (*string, bool)`

GetMetadataJsonOk returns a tuple with the MetadataJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataJson

`func (o *CorrelatedEvent) SetMetadataJson(v string)`

SetMetadataJson sets MetadataJson field to given value.

### HasMetadataJson

`func (o *CorrelatedEvent) HasMetadataJson() bool`

HasMetadataJson returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


