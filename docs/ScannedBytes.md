# ScannedBytes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Infrequent** | Pointer to **int64** | The total number of scanned bytes from infrequent tier data for the query in bytes. | [optional] 
**Continuous** | Pointer to **int64** | The total number of scanned bytes from continuous tier data for the query in bytes. | [optional] 
**Frequent** | Pointer to **int64** | The total number of scanned bytes from frequent tier data for the query in bytes. | [optional] 
**Security** | Pointer to **int64** | The total number of scanned bytes from security tier data for the query in bytes. | [optional] 
**Tracing** | Pointer to **int64** | The total number of scanned bytes from tracing tier data for the query in bytes. | [optional] 
**Upfront** | Pointer to **int64** | The total number of scanned bytes from upfront tier data for the query in bytes. | [optional] 
**Metered** | Pointer to **int64** | The total number of scanned bytes from metered tier data for the query in bytes. | [optional] 
**Rce** | Pointer to **int64** | The total number of scanned bytes from rce tier data for the query in bytes. | [optional] 
**Flex** | Pointer to **int64** | The total number of scanned bytes from flex tier data for the query in bytes. | [optional] 

## Methods

### NewScannedBytes

`func NewScannedBytes() *ScannedBytes`

NewScannedBytes instantiates a new ScannedBytes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScannedBytesWithDefaults

`func NewScannedBytesWithDefaults() *ScannedBytes`

NewScannedBytesWithDefaults instantiates a new ScannedBytes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInfrequent

`func (o *ScannedBytes) GetInfrequent() int64`

GetInfrequent returns the Infrequent field if non-nil, zero value otherwise.

### GetInfrequentOk

`func (o *ScannedBytes) GetInfrequentOk() (*int64, bool)`

GetInfrequentOk returns a tuple with the Infrequent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrequent

`func (o *ScannedBytes) SetInfrequent(v int64)`

SetInfrequent sets Infrequent field to given value.

### HasInfrequent

`func (o *ScannedBytes) HasInfrequent() bool`

HasInfrequent returns a boolean if a field has been set.

### GetContinuous

`func (o *ScannedBytes) GetContinuous() int64`

GetContinuous returns the Continuous field if non-nil, zero value otherwise.

### GetContinuousOk

`func (o *ScannedBytes) GetContinuousOk() (*int64, bool)`

GetContinuousOk returns a tuple with the Continuous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuous

`func (o *ScannedBytes) SetContinuous(v int64)`

SetContinuous sets Continuous field to given value.

### HasContinuous

`func (o *ScannedBytes) HasContinuous() bool`

HasContinuous returns a boolean if a field has been set.

### GetFrequent

`func (o *ScannedBytes) GetFrequent() int64`

GetFrequent returns the Frequent field if non-nil, zero value otherwise.

### GetFrequentOk

`func (o *ScannedBytes) GetFrequentOk() (*int64, bool)`

GetFrequentOk returns a tuple with the Frequent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequent

`func (o *ScannedBytes) SetFrequent(v int64)`

SetFrequent sets Frequent field to given value.

### HasFrequent

`func (o *ScannedBytes) HasFrequent() bool`

HasFrequent returns a boolean if a field has been set.

### GetSecurity

`func (o *ScannedBytes) GetSecurity() int64`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *ScannedBytes) GetSecurityOk() (*int64, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *ScannedBytes) SetSecurity(v int64)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *ScannedBytes) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetTracing

`func (o *ScannedBytes) GetTracing() int64`

GetTracing returns the Tracing field if non-nil, zero value otherwise.

### GetTracingOk

`func (o *ScannedBytes) GetTracingOk() (*int64, bool)`

GetTracingOk returns a tuple with the Tracing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTracing

`func (o *ScannedBytes) SetTracing(v int64)`

SetTracing sets Tracing field to given value.

### HasTracing

`func (o *ScannedBytes) HasTracing() bool`

HasTracing returns a boolean if a field has been set.

### GetUpfront

`func (o *ScannedBytes) GetUpfront() int64`

GetUpfront returns the Upfront field if non-nil, zero value otherwise.

### GetUpfrontOk

`func (o *ScannedBytes) GetUpfrontOk() (*int64, bool)`

GetUpfrontOk returns a tuple with the Upfront field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpfront

`func (o *ScannedBytes) SetUpfront(v int64)`

SetUpfront sets Upfront field to given value.

### HasUpfront

`func (o *ScannedBytes) HasUpfront() bool`

HasUpfront returns a boolean if a field has been set.

### GetMetered

`func (o *ScannedBytes) GetMetered() int64`

GetMetered returns the Metered field if non-nil, zero value otherwise.

### GetMeteredOk

`func (o *ScannedBytes) GetMeteredOk() (*int64, bool)`

GetMeteredOk returns a tuple with the Metered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetered

`func (o *ScannedBytes) SetMetered(v int64)`

SetMetered sets Metered field to given value.

### HasMetered

`func (o *ScannedBytes) HasMetered() bool`

HasMetered returns a boolean if a field has been set.

### GetRce

`func (o *ScannedBytes) GetRce() int64`

GetRce returns the Rce field if non-nil, zero value otherwise.

### GetRceOk

`func (o *ScannedBytes) GetRceOk() (*int64, bool)`

GetRceOk returns a tuple with the Rce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRce

`func (o *ScannedBytes) SetRce(v int64)`

SetRce sets Rce field to given value.

### HasRce

`func (o *ScannedBytes) HasRce() bool`

HasRce returns a boolean if a field has been set.

### GetFlex

`func (o *ScannedBytes) GetFlex() int64`

GetFlex returns the Flex field if non-nil, zero value otherwise.

### GetFlexOk

`func (o *ScannedBytes) GetFlexOk() (*int64, bool)`

GetFlexOk returns a tuple with the Flex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlex

`func (o *ScannedBytes) SetFlex(v int64)`

SetFlex sets Flex field to given value.

### HasFlex

`func (o *ScannedBytes) HasFlex() bool`

HasFlex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


