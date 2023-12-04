# UsageDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntitlementUsages** | [**[]EntitlementUsage**](EntitlementUsage.md) | An array of usages. | 
**StartDate** | **string** | Start date of the data usage. | 
**EndDate** | **string** | End date of the data usage. | 

## Methods

### NewUsageDetails

`func NewUsageDetails(entitlementUsages []EntitlementUsage, startDate string, endDate string, ) *UsageDetails`

NewUsageDetails instantiates a new UsageDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageDetailsWithDefaults

`func NewUsageDetailsWithDefaults() *UsageDetails`

NewUsageDetailsWithDefaults instantiates a new UsageDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntitlementUsages

`func (o *UsageDetails) GetEntitlementUsages() []EntitlementUsage`

GetEntitlementUsages returns the EntitlementUsages field if non-nil, zero value otherwise.

### GetEntitlementUsagesOk

`func (o *UsageDetails) GetEntitlementUsagesOk() (*[]EntitlementUsage, bool)`

GetEntitlementUsagesOk returns a tuple with the EntitlementUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementUsages

`func (o *UsageDetails) SetEntitlementUsages(v []EntitlementUsage)`

SetEntitlementUsages sets EntitlementUsages field to given value.


### GetStartDate

`func (o *UsageDetails) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *UsageDetails) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *UsageDetails) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *UsageDetails) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *UsageDetails) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *UsageDetails) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


