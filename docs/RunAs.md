# RunAs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RunAsId** | **string** | The runAsId indicates the context in which monitors will run. If not provided, then it will run in the context of the monitor author. | 

## Methods

### NewRunAs

`func NewRunAs(runAsId string, ) *RunAs`

NewRunAs instantiates a new RunAs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunAsWithDefaults

`func NewRunAsWithDefaults() *RunAs`

NewRunAsWithDefaults instantiates a new RunAs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRunAsId

`func (o *RunAs) GetRunAsId() string`

GetRunAsId returns the RunAsId field if non-nil, zero value otherwise.

### GetRunAsIdOk

`func (o *RunAs) GetRunAsIdOk() (*string, bool)`

GetRunAsIdOk returns a tuple with the RunAsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunAsId

`func (o *RunAs) SetRunAsId(v string)`

SetRunAsId sets RunAsId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


