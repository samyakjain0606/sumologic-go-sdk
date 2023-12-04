# Monitor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MonitorTriggers** | [**[]MonitorTrigger**](MonitorTrigger.md) | Monitors over which the SLO is defined. | 

## Methods

### NewMonitor

`func NewMonitor(monitorTriggers []MonitorTrigger, ) *Monitor`

NewMonitor instantiates a new Monitor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorWithDefaults

`func NewMonitorWithDefaults() *Monitor`

NewMonitorWithDefaults instantiates a new Monitor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonitorTriggers

`func (o *Monitor) GetMonitorTriggers() []MonitorTrigger`

GetMonitorTriggers returns the MonitorTriggers field if non-nil, zero value otherwise.

### GetMonitorTriggersOk

`func (o *Monitor) GetMonitorTriggersOk() (*[]MonitorTrigger, bool)`

GetMonitorTriggersOk returns a tuple with the MonitorTriggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorTriggers

`func (o *Monitor) SetMonitorTriggers(v []MonitorTrigger)`

SetMonitorTriggers sets MonitorTriggers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


