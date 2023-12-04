# VisualDataAxes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**X** | [**[]VisualAxisData**](VisualAxisData.md) | The data of the primary x axis. | 
**Y** | [**[]VisualAxisData**](VisualAxisData.md) | The data of the primary y axis. | 
**X2** | Pointer to [**[]VisualAxisData**](VisualAxisData.md) | The data of the secondary x axis. | [optional] 
**Y2** | Pointer to [**[]VisualAxisData**](VisualAxisData.md) | The data of the secondary y axis. | [optional] 

## Methods

### NewVisualDataAxes

`func NewVisualDataAxes(x []VisualAxisData, y []VisualAxisData, ) *VisualDataAxes`

NewVisualDataAxes instantiates a new VisualDataAxes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVisualDataAxesWithDefaults

`func NewVisualDataAxesWithDefaults() *VisualDataAxes`

NewVisualDataAxesWithDefaults instantiates a new VisualDataAxes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetX

`func (o *VisualDataAxes) GetX() []VisualAxisData`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *VisualDataAxes) GetXOk() (*[]VisualAxisData, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *VisualDataAxes) SetX(v []VisualAxisData)`

SetX sets X field to given value.


### GetY

`func (o *VisualDataAxes) GetY() []VisualAxisData`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *VisualDataAxes) GetYOk() (*[]VisualAxisData, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *VisualDataAxes) SetY(v []VisualAxisData)`

SetY sets Y field to given value.


### GetX2

`func (o *VisualDataAxes) GetX2() []VisualAxisData`

GetX2 returns the X2 field if non-nil, zero value otherwise.

### GetX2Ok

`func (o *VisualDataAxes) GetX2Ok() (*[]VisualAxisData, bool)`

GetX2Ok returns a tuple with the X2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX2

`func (o *VisualDataAxes) SetX2(v []VisualAxisData)`

SetX2 sets X2 field to given value.

### HasX2

`func (o *VisualDataAxes) HasX2() bool`

HasX2 returns a boolean if a field has been set.

### GetY2

`func (o *VisualDataAxes) GetY2() []VisualAxisData`

GetY2 returns the Y2 field if non-nil, zero value otherwise.

### GetY2Ok

`func (o *VisualDataAxes) GetY2Ok() (*[]VisualAxisData, bool)`

GetY2Ok returns a tuple with the Y2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY2

`func (o *VisualDataAxes) SetY2(v []VisualAxisData)`

SetY2 sets Y2 field to given value.

### HasY2

`func (o *VisualDataAxes) HasY2() bool`

HasY2 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


