# UpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductId** | **string** | Unique identifier of the product in current plan. Valid values are: 1. &#x60;Free&#x60; 2. &#x60;Trial&#x60; 3. &#x60;Essentials&#x60; 4. &#x60;EnterpriseOps&#x60; 5. &#x60;EnterpriseSec&#x60; 6. &#x60;EnterpriseSuite&#x60;  | 
**BillingFrequency** | **string** | Identifier for the plans billing term. Valid values are:  1. Monthly  2. Annually  | 
**Baselines** | [**SelfServiceCreditsBaselines**](SelfServiceCreditsBaselines.md) |  | 
**Immediate** | Pointer to **bool** | true in case the subscription baselines need to be updated immediately(only for monthly customers who are staying on the monthly plan) | [optional] [default to false]

## Methods

### NewUpdateRequest

`func NewUpdateRequest(productId string, billingFrequency string, baselines SelfServiceCreditsBaselines, ) *UpdateRequest`

NewUpdateRequest instantiates a new UpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRequestWithDefaults

`func NewUpdateRequestWithDefaults() *UpdateRequest`

NewUpdateRequestWithDefaults instantiates a new UpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductId

`func (o *UpdateRequest) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *UpdateRequest) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *UpdateRequest) SetProductId(v string)`

SetProductId sets ProductId field to given value.


### GetBillingFrequency

`func (o *UpdateRequest) GetBillingFrequency() string`

GetBillingFrequency returns the BillingFrequency field if non-nil, zero value otherwise.

### GetBillingFrequencyOk

`func (o *UpdateRequest) GetBillingFrequencyOk() (*string, bool)`

GetBillingFrequencyOk returns a tuple with the BillingFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingFrequency

`func (o *UpdateRequest) SetBillingFrequency(v string)`

SetBillingFrequency sets BillingFrequency field to given value.


### GetBaselines

`func (o *UpdateRequest) GetBaselines() SelfServiceCreditsBaselines`

GetBaselines returns the Baselines field if non-nil, zero value otherwise.

### GetBaselinesOk

`func (o *UpdateRequest) GetBaselinesOk() (*SelfServiceCreditsBaselines, bool)`

GetBaselinesOk returns a tuple with the Baselines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselines

`func (o *UpdateRequest) SetBaselines(v SelfServiceCreditsBaselines)`

SetBaselines sets Baselines field to given value.


### GetImmediate

`func (o *UpdateRequest) GetImmediate() bool`

GetImmediate returns the Immediate field if non-nil, zero value otherwise.

### GetImmediateOk

`func (o *UpdateRequest) GetImmediateOk() (*bool, bool)`

GetImmediateOk returns a tuple with the Immediate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmediate

`func (o *UpdateRequest) SetImmediate(v bool)`

SetImmediate sets Immediate field to given value.

### HasImmediate

`func (o *UpdateRequest) HasImmediate() bool`

HasImmediate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


