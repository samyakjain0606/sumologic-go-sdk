# CorrelatedEvents

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CorrelationFinished** | **bool** | Flag indicating correlation completion. | 
**Events** | [**[]CorrelatedEvent**](CorrelatedEvent.md) | List of events. | 

## Methods

### NewCorrelatedEvents

`func NewCorrelatedEvents(correlationFinished bool, events []CorrelatedEvent, ) *CorrelatedEvents`

NewCorrelatedEvents instantiates a new CorrelatedEvents object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCorrelatedEventsWithDefaults

`func NewCorrelatedEventsWithDefaults() *CorrelatedEvents`

NewCorrelatedEventsWithDefaults instantiates a new CorrelatedEvents object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCorrelationFinished

`func (o *CorrelatedEvents) GetCorrelationFinished() bool`

GetCorrelationFinished returns the CorrelationFinished field if non-nil, zero value otherwise.

### GetCorrelationFinishedOk

`func (o *CorrelatedEvents) GetCorrelationFinishedOk() (*bool, bool)`

GetCorrelationFinishedOk returns a tuple with the CorrelationFinished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationFinished

`func (o *CorrelatedEvents) SetCorrelationFinished(v bool)`

SetCorrelationFinished sets CorrelationFinished field to given value.


### GetEvents

`func (o *CorrelatedEvents) GetEvents() []CorrelatedEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *CorrelatedEvents) GetEventsOk() (*[]CorrelatedEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *CorrelatedEvents) SetEvents(v []CorrelatedEvent)`

SetEvents sets Events field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


