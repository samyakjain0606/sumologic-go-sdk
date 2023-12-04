# ProductGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductGroupName** | **string** | Name of the Product group:  | 
**ProductGroupId** | Pointer to **string** | Id of the Product group | [optional] 
**ProductVariables** | [**[]ProductVariable**](ProductVariable.md) | Different product variables of the product group | 
**ProvisioningSupported** | Pointer to **bool** | Is provisioning supported on this Product Group. This is applicable for product variables which are not enabled by default. | [optional] 
**Description** | Pointer to **string** | Description about the Product group | [optional] 
**LearnMoreLink** | Pointer to **string** | Link to learn more about the Product group | [optional] 

## Methods

### NewProductGroup

`func NewProductGroup(productGroupName string, productVariables []ProductVariable, ) *ProductGroup`

NewProductGroup instantiates a new ProductGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductGroupWithDefaults

`func NewProductGroupWithDefaults() *ProductGroup`

NewProductGroupWithDefaults instantiates a new ProductGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductGroupName

`func (o *ProductGroup) GetProductGroupName() string`

GetProductGroupName returns the ProductGroupName field if non-nil, zero value otherwise.

### GetProductGroupNameOk

`func (o *ProductGroup) GetProductGroupNameOk() (*string, bool)`

GetProductGroupNameOk returns a tuple with the ProductGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductGroupName

`func (o *ProductGroup) SetProductGroupName(v string)`

SetProductGroupName sets ProductGroupName field to given value.


### GetProductGroupId

`func (o *ProductGroup) GetProductGroupId() string`

GetProductGroupId returns the ProductGroupId field if non-nil, zero value otherwise.

### GetProductGroupIdOk

`func (o *ProductGroup) GetProductGroupIdOk() (*string, bool)`

GetProductGroupIdOk returns a tuple with the ProductGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductGroupId

`func (o *ProductGroup) SetProductGroupId(v string)`

SetProductGroupId sets ProductGroupId field to given value.

### HasProductGroupId

`func (o *ProductGroup) HasProductGroupId() bool`

HasProductGroupId returns a boolean if a field has been set.

### GetProductVariables

`func (o *ProductGroup) GetProductVariables() []ProductVariable`

GetProductVariables returns the ProductVariables field if non-nil, zero value otherwise.

### GetProductVariablesOk

`func (o *ProductGroup) GetProductVariablesOk() (*[]ProductVariable, bool)`

GetProductVariablesOk returns a tuple with the ProductVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductVariables

`func (o *ProductGroup) SetProductVariables(v []ProductVariable)`

SetProductVariables sets ProductVariables field to given value.


### GetProvisioningSupported

`func (o *ProductGroup) GetProvisioningSupported() bool`

GetProvisioningSupported returns the ProvisioningSupported field if non-nil, zero value otherwise.

### GetProvisioningSupportedOk

`func (o *ProductGroup) GetProvisioningSupportedOk() (*bool, bool)`

GetProvisioningSupportedOk returns a tuple with the ProvisioningSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningSupported

`func (o *ProductGroup) SetProvisioningSupported(v bool)`

SetProvisioningSupported sets ProvisioningSupported field to given value.

### HasProvisioningSupported

`func (o *ProductGroup) HasProvisioningSupported() bool`

HasProvisioningSupported returns a boolean if a field has been set.

### GetDescription

`func (o *ProductGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProductGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProductGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProductGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLearnMoreLink

`func (o *ProductGroup) GetLearnMoreLink() string`

GetLearnMoreLink returns the LearnMoreLink field if non-nil, zero value otherwise.

### GetLearnMoreLinkOk

`func (o *ProductGroup) GetLearnMoreLinkOk() (*string, bool)`

GetLearnMoreLinkOk returns a tuple with the LearnMoreLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLearnMoreLink

`func (o *ProductGroup) SetLearnMoreLink(v string)`

SetLearnMoreLink sets LearnMoreLink field to given value.

### HasLearnMoreLink

`func (o *ProductGroup) HasLearnMoreLink() bool`

HasLearnMoreLink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


