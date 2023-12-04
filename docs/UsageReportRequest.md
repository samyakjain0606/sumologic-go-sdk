# UsageReportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDate** | Pointer to **string** | Start date, without the time, of the usage data to fetch. If no value is provided startDate is used as the start of the subscription. The start date cannot be before the start of the subscription. | [optional] 
**EndDate** | Pointer to **string** | End date, without the time, of usage data to fetch. If no value is provided endDate is used as the end of the subscription. The end date cannot be after the end of the subscription. | [optional] 
**GroupBy** | Pointer to **string** | Perform a groupBy operation on the usage details. If no value is provided data is grouped by &#x60;Day&#x60; - &#x60;day&#x60;: Aggregate the data by day - &#x60;week&#x60;: Aggregate the data by week. Week starts at Monday and ends at sunday night. - &#x60;month&#x60;: Aggregate the data by calendar month. | [optional] [default to "day"]
**ReportType** | Pointer to **string** | Specifies the type of report to be exported. Available types are &#x60;standard&#x60; and &#x60;detailed&#x60;. An additional &#x60;childDetailed&#x60; type is available for Sumo Orgs parents. Detailed report will have raw consumption along with the credits breakdown. If no value is provided Standard reports will be exported. | [optional] [default to "standard"]
**IncludeDeploymentCharge** | Pointer to **bool** | Deployment charges will be applied to the returned usages csv if this is set to true and the organization  is a part of Sumo Organizations as a child organization. | [optional] [default to false]

## Methods

### NewUsageReportRequest

`func NewUsageReportRequest() *UsageReportRequest`

NewUsageReportRequest instantiates a new UsageReportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageReportRequestWithDefaults

`func NewUsageReportRequestWithDefaults() *UsageReportRequest`

NewUsageReportRequestWithDefaults instantiates a new UsageReportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartDate

`func (o *UsageReportRequest) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *UsageReportRequest) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *UsageReportRequest) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *UsageReportRequest) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *UsageReportRequest) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *UsageReportRequest) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *UsageReportRequest) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *UsageReportRequest) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetGroupBy

`func (o *UsageReportRequest) GetGroupBy() string`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *UsageReportRequest) GetGroupByOk() (*string, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *UsageReportRequest) SetGroupBy(v string)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *UsageReportRequest) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetReportType

`func (o *UsageReportRequest) GetReportType() string`

GetReportType returns the ReportType field if non-nil, zero value otherwise.

### GetReportTypeOk

`func (o *UsageReportRequest) GetReportTypeOk() (*string, bool)`

GetReportTypeOk returns a tuple with the ReportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportType

`func (o *UsageReportRequest) SetReportType(v string)`

SetReportType sets ReportType field to given value.

### HasReportType

`func (o *UsageReportRequest) HasReportType() bool`

HasReportType returns a boolean if a field has been set.

### GetIncludeDeploymentCharge

`func (o *UsageReportRequest) GetIncludeDeploymentCharge() bool`

GetIncludeDeploymentCharge returns the IncludeDeploymentCharge field if non-nil, zero value otherwise.

### GetIncludeDeploymentChargeOk

`func (o *UsageReportRequest) GetIncludeDeploymentChargeOk() (*bool, bool)`

GetIncludeDeploymentChargeOk returns a tuple with the IncludeDeploymentCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeDeploymentCharge

`func (o *UsageReportRequest) SetIncludeDeploymentCharge(v bool)`

SetIncludeDeploymentCharge sets IncludeDeploymentCharge field to given value.

### HasIncludeDeploymentCharge

`func (o *UsageReportRequest) HasIncludeDeploymentCharge() bool`

HasIncludeDeploymentCharge returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


