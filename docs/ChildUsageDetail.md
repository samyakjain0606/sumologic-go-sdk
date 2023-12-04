# ChildUsageDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | Status of the child org. | 
**OrgName** | Pointer to **string** | Name of the child org. | [optional] 
**OrgId** | **string** | The unique identifier of an organization. It consists of the deployment ID and the hexadecimal account ID separated by a dash &#x60;-&#x60; character. | 
**AllocatedCredits** | Pointer to **float64** | Denotes the total number of credits provisioned for the child organization to use. | [optional] 
**Usages** | [**ChildUsage**](ChildUsage.md) |  | 

## Methods

### NewChildUsageDetail

`func NewChildUsageDetail(status string, orgId string, usages ChildUsage, ) *ChildUsageDetail`

NewChildUsageDetail instantiates a new ChildUsageDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChildUsageDetailWithDefaults

`func NewChildUsageDetailWithDefaults() *ChildUsageDetail`

NewChildUsageDetailWithDefaults instantiates a new ChildUsageDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ChildUsageDetail) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ChildUsageDetail) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ChildUsageDetail) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetOrgName

`func (o *ChildUsageDetail) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *ChildUsageDetail) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *ChildUsageDetail) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *ChildUsageDetail) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetOrgId

`func (o *ChildUsageDetail) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ChildUsageDetail) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ChildUsageDetail) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetAllocatedCredits

`func (o *ChildUsageDetail) GetAllocatedCredits() float64`

GetAllocatedCredits returns the AllocatedCredits field if non-nil, zero value otherwise.

### GetAllocatedCreditsOk

`func (o *ChildUsageDetail) GetAllocatedCreditsOk() (*float64, bool)`

GetAllocatedCreditsOk returns a tuple with the AllocatedCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedCredits

`func (o *ChildUsageDetail) SetAllocatedCredits(v float64)`

SetAllocatedCredits sets AllocatedCredits field to given value.

### HasAllocatedCredits

`func (o *ChildUsageDetail) HasAllocatedCredits() bool`

HasAllocatedCredits returns a boolean if a field has been set.

### GetUsages

`func (o *ChildUsageDetail) GetUsages() ChildUsage`

GetUsages returns the Usages field if non-nil, zero value otherwise.

### GetUsagesOk

`func (o *ChildUsageDetail) GetUsagesOk() (*ChildUsage, bool)`

GetUsagesOk returns a tuple with the Usages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsages

`func (o *ChildUsageDetail) SetUsages(v ChildUsage)`

SetUsages sets Usages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


