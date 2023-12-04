# ChildUsageDetailsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDate** | Pointer to **string** | Start date, without the time, of the usage data to fetch. | [optional] 
**EndDate** | Pointer to **string** | End date, without the time, of usage data to fetch. | [optional] 

## Methods

### NewChildUsageDetailsRequest

`func NewChildUsageDetailsRequest() *ChildUsageDetailsRequest`

NewChildUsageDetailsRequest instantiates a new ChildUsageDetailsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChildUsageDetailsRequestWithDefaults

`func NewChildUsageDetailsRequestWithDefaults() *ChildUsageDetailsRequest`

NewChildUsageDetailsRequestWithDefaults instantiates a new ChildUsageDetailsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartDate

`func (o *ChildUsageDetailsRequest) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ChildUsageDetailsRequest) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ChildUsageDetailsRequest) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ChildUsageDetailsRequest) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *ChildUsageDetailsRequest) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ChildUsageDetailsRequest) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ChildUsageDetailsRequest) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ChildUsageDetailsRequest) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


