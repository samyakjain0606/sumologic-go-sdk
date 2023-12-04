# VisualPointData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**X** | Pointer to **float64** | Value that represents a point on the x axis. | [optional] 
**Y** | **string** | Value that represents a point on the y axis. | 
**IsFilled** | Pointer to **bool** | Whether the field is interpolated or extrapolated - not derived from underlying data. | [optional] [default to false]
**XAxisValues** | Pointer to **map[string]string** | Values that represents a point on the x axis. | [optional] [default to {}]
**OutlierData** | Pointer to [**VisualOutlierData**](VisualOutlierData.md) |  | [optional] 

## Methods

### NewVisualPointData

`func NewVisualPointData(y string, ) *VisualPointData`

NewVisualPointData instantiates a new VisualPointData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVisualPointDataWithDefaults

`func NewVisualPointDataWithDefaults() *VisualPointData`

NewVisualPointDataWithDefaults instantiates a new VisualPointData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetX

`func (o *VisualPointData) GetX() float64`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *VisualPointData) GetXOk() (*float64, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *VisualPointData) SetX(v float64)`

SetX sets X field to given value.

### HasX

`func (o *VisualPointData) HasX() bool`

HasX returns a boolean if a field has been set.

### GetY

`func (o *VisualPointData) GetY() string`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *VisualPointData) GetYOk() (*string, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *VisualPointData) SetY(v string)`

SetY sets Y field to given value.


### GetIsFilled

`func (o *VisualPointData) GetIsFilled() bool`

GetIsFilled returns the IsFilled field if non-nil, zero value otherwise.

### GetIsFilledOk

`func (o *VisualPointData) GetIsFilledOk() (*bool, bool)`

GetIsFilledOk returns a tuple with the IsFilled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFilled

`func (o *VisualPointData) SetIsFilled(v bool)`

SetIsFilled sets IsFilled field to given value.

### HasIsFilled

`func (o *VisualPointData) HasIsFilled() bool`

HasIsFilled returns a boolean if a field has been set.

### GetXAxisValues

`func (o *VisualPointData) GetXAxisValues() map[string]string`

GetXAxisValues returns the XAxisValues field if non-nil, zero value otherwise.

### GetXAxisValuesOk

`func (o *VisualPointData) GetXAxisValuesOk() (*map[string]string, bool)`

GetXAxisValuesOk returns a tuple with the XAxisValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXAxisValues

`func (o *VisualPointData) SetXAxisValues(v map[string]string)`

SetXAxisValues sets XAxisValues field to given value.

### HasXAxisValues

`func (o *VisualPointData) HasXAxisValues() bool`

HasXAxisValues returns a boolean if a field has been set.

### GetOutlierData

`func (o *VisualPointData) GetOutlierData() VisualOutlierData`

GetOutlierData returns the OutlierData field if non-nil, zero value otherwise.

### GetOutlierDataOk

`func (o *VisualPointData) GetOutlierDataOk() (*VisualOutlierData, bool)`

GetOutlierDataOk returns a tuple with the OutlierData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutlierData

`func (o *VisualPointData) SetOutlierData(v VisualOutlierData)`

SetOutlierData sets OutlierData field to given value.

### HasOutlierData

`func (o *VisualPointData) HasOutlierData() bool`

HasOutlierData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


