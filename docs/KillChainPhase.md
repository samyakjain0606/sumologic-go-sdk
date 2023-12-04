# KillChainPhase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KillChainName** | **string** | The name of the kill chain. The value of this property SHOULD be all lowercase and SHOULD use hyphens instead of spaces or underscores as word separators | 
**PhaseName** | Pointer to **string** | The name of the phase in the kill chain. The value of this property SHOULD be all lowercase and SHOULD use hyphens instead of spaces or underscores as word separators | [optional] 

## Methods

### NewKillChainPhase

`func NewKillChainPhase(killChainName string, ) *KillChainPhase`

NewKillChainPhase instantiates a new KillChainPhase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKillChainPhaseWithDefaults

`func NewKillChainPhaseWithDefaults() *KillChainPhase`

NewKillChainPhaseWithDefaults instantiates a new KillChainPhase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKillChainName

`func (o *KillChainPhase) GetKillChainName() string`

GetKillChainName returns the KillChainName field if non-nil, zero value otherwise.

### GetKillChainNameOk

`func (o *KillChainPhase) GetKillChainNameOk() (*string, bool)`

GetKillChainNameOk returns a tuple with the KillChainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKillChainName

`func (o *KillChainPhase) SetKillChainName(v string)`

SetKillChainName sets KillChainName field to given value.


### GetPhaseName

`func (o *KillChainPhase) GetPhaseName() string`

GetPhaseName returns the PhaseName field if non-nil, zero value otherwise.

### GetPhaseNameOk

`func (o *KillChainPhase) GetPhaseNameOk() (*string, bool)`

GetPhaseNameOk returns a tuple with the PhaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhaseName

`func (o *KillChainPhase) SetPhaseName(v string)`

SetPhaseName sets PhaseName field to given value.

### HasPhaseName

`func (o *KillChainPhase) HasPhaseName() bool`

HasPhaseName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


