# UploadNormalizedIndicatorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Indicators** | [**[]NormalizedIndicator**](NormalizedIndicator.md) | The list of normalized threat intel indicators to upload. | 

## Methods

### NewUploadNormalizedIndicatorRequest

`func NewUploadNormalizedIndicatorRequest(indicators []NormalizedIndicator, ) *UploadNormalizedIndicatorRequest`

NewUploadNormalizedIndicatorRequest instantiates a new UploadNormalizedIndicatorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadNormalizedIndicatorRequestWithDefaults

`func NewUploadNormalizedIndicatorRequestWithDefaults() *UploadNormalizedIndicatorRequest`

NewUploadNormalizedIndicatorRequestWithDefaults instantiates a new UploadNormalizedIndicatorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndicators

`func (o *UploadNormalizedIndicatorRequest) GetIndicators() []NormalizedIndicator`

GetIndicators returns the Indicators field if non-nil, zero value otherwise.

### GetIndicatorsOk

`func (o *UploadNormalizedIndicatorRequest) GetIndicatorsOk() (*[]NormalizedIndicator, bool)`

GetIndicatorsOk returns a tuple with the Indicators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicators

`func (o *UploadNormalizedIndicatorRequest) SetIndicators(v []NormalizedIndicator)`

SetIndicators sets Indicators field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


