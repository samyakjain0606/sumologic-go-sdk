# DisableUnusedAccessKeysPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnusedAccessKeysDisableAfterInDays** | **int32** | The number of days it will take for an unused access key to automatically disable. Setting it to 0 (never) means that the accessKeys will not be disabled automatically. | 

## Methods

### NewDisableUnusedAccessKeysPolicy

`func NewDisableUnusedAccessKeysPolicy(unusedAccessKeysDisableAfterInDays int32, ) *DisableUnusedAccessKeysPolicy`

NewDisableUnusedAccessKeysPolicy instantiates a new DisableUnusedAccessKeysPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisableUnusedAccessKeysPolicyWithDefaults

`func NewDisableUnusedAccessKeysPolicyWithDefaults() *DisableUnusedAccessKeysPolicy`

NewDisableUnusedAccessKeysPolicyWithDefaults instantiates a new DisableUnusedAccessKeysPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnusedAccessKeysDisableAfterInDays

`func (o *DisableUnusedAccessKeysPolicy) GetUnusedAccessKeysDisableAfterInDays() int32`

GetUnusedAccessKeysDisableAfterInDays returns the UnusedAccessKeysDisableAfterInDays field if non-nil, zero value otherwise.

### GetUnusedAccessKeysDisableAfterInDaysOk

`func (o *DisableUnusedAccessKeysPolicy) GetUnusedAccessKeysDisableAfterInDaysOk() (*int32, bool)`

GetUnusedAccessKeysDisableAfterInDaysOk returns a tuple with the UnusedAccessKeysDisableAfterInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnusedAccessKeysDisableAfterInDays

`func (o *DisableUnusedAccessKeysPolicy) SetUnusedAccessKeysDisableAfterInDays(v int32)`

SetUnusedAccessKeysDisableAfterInDays sets UnusedAccessKeysDisableAfterInDays field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


