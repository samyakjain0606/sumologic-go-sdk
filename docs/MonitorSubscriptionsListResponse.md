# MonitorSubscriptionsListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subscriptions** | [**[]MonitorSubscription**](MonitorSubscription.md) | List of existing subscriptions. | 
**Exhaustive** | **bool** | If true, the list contains all existing subscriptions. | 

## Methods

### NewMonitorSubscriptionsListResponse

`func NewMonitorSubscriptionsListResponse(subscriptions []MonitorSubscription, exhaustive bool, ) *MonitorSubscriptionsListResponse`

NewMonitorSubscriptionsListResponse instantiates a new MonitorSubscriptionsListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorSubscriptionsListResponseWithDefaults

`func NewMonitorSubscriptionsListResponseWithDefaults() *MonitorSubscriptionsListResponse`

NewMonitorSubscriptionsListResponseWithDefaults instantiates a new MonitorSubscriptionsListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptions

`func (o *MonitorSubscriptionsListResponse) GetSubscriptions() []MonitorSubscription`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *MonitorSubscriptionsListResponse) GetSubscriptionsOk() (*[]MonitorSubscription, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *MonitorSubscriptionsListResponse) SetSubscriptions(v []MonitorSubscription)`

SetSubscriptions sets Subscriptions field to given value.


### GetExhaustive

`func (o *MonitorSubscriptionsListResponse) GetExhaustive() bool`

GetExhaustive returns the Exhaustive field if non-nil, zero value otherwise.

### GetExhaustiveOk

`func (o *MonitorSubscriptionsListResponse) GetExhaustiveOk() (*bool, bool)`

GetExhaustiveOk returns a tuple with the Exhaustive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExhaustive

`func (o *MonitorSubscriptionsListResponse) SetExhaustive(v bool)`

SetExhaustive sets Exhaustive field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


