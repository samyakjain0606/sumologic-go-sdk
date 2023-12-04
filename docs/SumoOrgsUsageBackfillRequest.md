# SumoOrgsUsageBackfillRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | **int64** | the customer ID of a mam org | 
**From** | **int64** | epoch millis of date from which usage is to be copied | 
**To** | **int64** | epoch millis of date upto which usage is to be copied | 

## Methods

### NewSumoOrgsUsageBackfillRequest

`func NewSumoOrgsUsageBackfillRequest(customerId int64, from int64, to int64, ) *SumoOrgsUsageBackfillRequest`

NewSumoOrgsUsageBackfillRequest instantiates a new SumoOrgsUsageBackfillRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSumoOrgsUsageBackfillRequestWithDefaults

`func NewSumoOrgsUsageBackfillRequestWithDefaults() *SumoOrgsUsageBackfillRequest`

NewSumoOrgsUsageBackfillRequestWithDefaults instantiates a new SumoOrgsUsageBackfillRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *SumoOrgsUsageBackfillRequest) GetCustomerId() int64`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *SumoOrgsUsageBackfillRequest) GetCustomerIdOk() (*int64, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *SumoOrgsUsageBackfillRequest) SetCustomerId(v int64)`

SetCustomerId sets CustomerId field to given value.


### GetFrom

`func (o *SumoOrgsUsageBackfillRequest) GetFrom() int64`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *SumoOrgsUsageBackfillRequest) GetFromOk() (*int64, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *SumoOrgsUsageBackfillRequest) SetFrom(v int64)`

SetFrom sets From field to given value.


### GetTo

`func (o *SumoOrgsUsageBackfillRequest) GetTo() int64`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *SumoOrgsUsageBackfillRequest) GetToOk() (*int64, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *SumoOrgsUsageBackfillRequest) SetTo(v int64)`

SetTo sets To field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


