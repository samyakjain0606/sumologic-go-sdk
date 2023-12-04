# SpansLimitItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Direction** | **string** | Describes whether the results should be sorted in an ascending or a descending order. | 
**LimitValue** | **int32** | The number of aggregated results returned, e.g. if 10 is requested, then only the first 10 aggregated results are returned.  | 

## Methods

### NewSpansLimitItem

`func NewSpansLimitItem(direction string, limitValue int32, ) *SpansLimitItem`

NewSpansLimitItem instantiates a new SpansLimitItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpansLimitItemWithDefaults

`func NewSpansLimitItemWithDefaults() *SpansLimitItem`

NewSpansLimitItemWithDefaults instantiates a new SpansLimitItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirection

`func (o *SpansLimitItem) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *SpansLimitItem) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *SpansLimitItem) SetDirection(v string)`

SetDirection sets Direction field to given value.


### GetLimitValue

`func (o *SpansLimitItem) GetLimitValue() int32`

GetLimitValue returns the LimitValue field if non-nil, zero value otherwise.

### GetLimitValueOk

`func (o *SpansLimitItem) GetLimitValueOk() (*int32, bool)`

GetLimitValueOk returns a tuple with the LimitValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitValue

`func (o *SpansLimitItem) SetLimitValue(v int32)`

SetLimitValue sets LimitValue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


