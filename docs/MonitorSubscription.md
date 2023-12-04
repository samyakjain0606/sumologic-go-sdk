# MonitorSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetId** | Pointer to **string** | The id of the subscription target. It can be either a monitor or a folder id. | [optional] 

## Methods

### NewMonitorSubscription

`func NewMonitorSubscription() *MonitorSubscription`

NewMonitorSubscription instantiates a new MonitorSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorSubscriptionWithDefaults

`func NewMonitorSubscriptionWithDefaults() *MonitorSubscription`

NewMonitorSubscriptionWithDefaults instantiates a new MonitorSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetId

`func (o *MonitorSubscription) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *MonitorSubscription) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *MonitorSubscription) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *MonitorSubscription) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


