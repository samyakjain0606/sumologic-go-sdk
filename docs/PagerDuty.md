# PagerDuty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionId** | **string** | The identifier of the connection. | 
**PayloadOverride** | Pointer to **string** | The override of the default JSON payload of the connection. Should be in JSON format. | [optional] 
**ResolutionPayloadOverride** | Pointer to **string** | The override of the resolution JSON payload of the connection. Should be in JSON format. | [optional] 

## Methods

### NewPagerDuty

`func NewPagerDuty(connectionId string, ) *PagerDuty`

NewPagerDuty instantiates a new PagerDuty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPagerDutyWithDefaults

`func NewPagerDutyWithDefaults() *PagerDuty`

NewPagerDutyWithDefaults instantiates a new PagerDuty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionId

`func (o *PagerDuty) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *PagerDuty) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *PagerDuty) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetPayloadOverride

`func (o *PagerDuty) GetPayloadOverride() string`

GetPayloadOverride returns the PayloadOverride field if non-nil, zero value otherwise.

### GetPayloadOverrideOk

`func (o *PagerDuty) GetPayloadOverrideOk() (*string, bool)`

GetPayloadOverrideOk returns a tuple with the PayloadOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadOverride

`func (o *PagerDuty) SetPayloadOverride(v string)`

SetPayloadOverride sets PayloadOverride field to given value.

### HasPayloadOverride

`func (o *PagerDuty) HasPayloadOverride() bool`

HasPayloadOverride returns a boolean if a field has been set.

### GetResolutionPayloadOverride

`func (o *PagerDuty) GetResolutionPayloadOverride() string`

GetResolutionPayloadOverride returns the ResolutionPayloadOverride field if non-nil, zero value otherwise.

### GetResolutionPayloadOverrideOk

`func (o *PagerDuty) GetResolutionPayloadOverrideOk() (*string, bool)`

GetResolutionPayloadOverrideOk returns a tuple with the ResolutionPayloadOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionPayloadOverride

`func (o *PagerDuty) SetResolutionPayloadOverride(v string)`

SetResolutionPayloadOverride sets ResolutionPayloadOverride field to given value.

### HasResolutionPayloadOverride

`func (o *PagerDuty) HasResolutionPayloadOverride() bool`

HasResolutionPayloadOverride returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


