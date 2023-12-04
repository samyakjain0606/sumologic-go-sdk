# EventContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventContextType** | **string** | Context for which correlated events are to be fetched. | 

## Methods

### NewEventContext

`func NewEventContext(eventContextType string, ) *EventContext`

NewEventContext instantiates a new EventContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventContextWithDefaults

`func NewEventContextWithDefaults() *EventContext`

NewEventContextWithDefaults instantiates a new EventContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventContextType

`func (o *EventContext) GetEventContextType() string`

GetEventContextType returns the EventContextType field if non-nil, zero value otherwise.

### GetEventContextTypeOk

`func (o *EventContext) GetEventContextTypeOk() (*string, bool)`

GetEventContextTypeOk returns a tuple with the EventContextType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventContextType

`func (o *EventContext) SetEventContextType(v string)`

SetEventContextType sets EventContextType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


