# PendingUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedOn** | **string** | The date on which the update request was created. | 
**Plan** | [**CurrentPlan**](CurrentPlan.md) |  | 

## Methods

### NewPendingUpdateRequest

`func NewPendingUpdateRequest(createdOn string, plan CurrentPlan, ) *PendingUpdateRequest`

NewPendingUpdateRequest instantiates a new PendingUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPendingUpdateRequestWithDefaults

`func NewPendingUpdateRequestWithDefaults() *PendingUpdateRequest`

NewPendingUpdateRequestWithDefaults instantiates a new PendingUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedOn

`func (o *PendingUpdateRequest) GetCreatedOn() string`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *PendingUpdateRequest) GetCreatedOnOk() (*string, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *PendingUpdateRequest) SetCreatedOn(v string)`

SetCreatedOn sets CreatedOn field to given value.


### GetPlan

`func (o *PendingUpdateRequest) GetPlan() CurrentPlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *PendingUpdateRequest) GetPlanOk() (*CurrentPlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *PendingUpdateRequest) SetPlan(v CurrentPlan)`

SetPlan sets Plan field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


