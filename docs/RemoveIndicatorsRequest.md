# RemoveIndicatorsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | **string** | The source of the indicator ID to match against | 
**IndicatorIds** | **[]string** | The list of indicator IDs to match against | 

## Methods

### NewRemoveIndicatorsRequest

`func NewRemoveIndicatorsRequest(source string, indicatorIds []string, ) *RemoveIndicatorsRequest`

NewRemoveIndicatorsRequest instantiates a new RemoveIndicatorsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoveIndicatorsRequestWithDefaults

`func NewRemoveIndicatorsRequestWithDefaults() *RemoveIndicatorsRequest`

NewRemoveIndicatorsRequestWithDefaults instantiates a new RemoveIndicatorsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *RemoveIndicatorsRequest) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *RemoveIndicatorsRequest) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *RemoveIndicatorsRequest) SetSource(v string)`

SetSource sets Source field to given value.


### GetIndicatorIds

`func (o *RemoveIndicatorsRequest) GetIndicatorIds() []string`

GetIndicatorIds returns the IndicatorIds field if non-nil, zero value otherwise.

### GetIndicatorIdsOk

`func (o *RemoveIndicatorsRequest) GetIndicatorIdsOk() (*[]string, bool)`

GetIndicatorIdsOk returns a tuple with the IndicatorIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicatorIds

`func (o *RemoveIndicatorsRequest) SetIndicatorIds(v []string)`

SetIndicatorIds sets IndicatorIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


