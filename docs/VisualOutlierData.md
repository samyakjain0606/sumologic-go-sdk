# VisualOutlierData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Baseline** | **float64** | The estimated value of the data point. | 
**Unit** | **float64** | The variation in the estimated value of the data point. | 
**LowerBound** | **float64** | The lower bound of the outlier band | 
**UpperBound** | **float64** | The upper bound of the outlier band | 
**IsOutlier** | **bool** | Indicates if the data point is outlier or not. | 

## Methods

### NewVisualOutlierData

`func NewVisualOutlierData(baseline float64, unit float64, lowerBound float64, upperBound float64, isOutlier bool, ) *VisualOutlierData`

NewVisualOutlierData instantiates a new VisualOutlierData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVisualOutlierDataWithDefaults

`func NewVisualOutlierDataWithDefaults() *VisualOutlierData`

NewVisualOutlierDataWithDefaults instantiates a new VisualOutlierData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseline

`func (o *VisualOutlierData) GetBaseline() float64`

GetBaseline returns the Baseline field if non-nil, zero value otherwise.

### GetBaselineOk

`func (o *VisualOutlierData) GetBaselineOk() (*float64, bool)`

GetBaselineOk returns a tuple with the Baseline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseline

`func (o *VisualOutlierData) SetBaseline(v float64)`

SetBaseline sets Baseline field to given value.


### GetUnit

`func (o *VisualOutlierData) GetUnit() float64`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *VisualOutlierData) GetUnitOk() (*float64, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *VisualOutlierData) SetUnit(v float64)`

SetUnit sets Unit field to given value.


### GetLowerBound

`func (o *VisualOutlierData) GetLowerBound() float64`

GetLowerBound returns the LowerBound field if non-nil, zero value otherwise.

### GetLowerBoundOk

`func (o *VisualOutlierData) GetLowerBoundOk() (*float64, bool)`

GetLowerBoundOk returns a tuple with the LowerBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowerBound

`func (o *VisualOutlierData) SetLowerBound(v float64)`

SetLowerBound sets LowerBound field to given value.


### GetUpperBound

`func (o *VisualOutlierData) GetUpperBound() float64`

GetUpperBound returns the UpperBound field if non-nil, zero value otherwise.

### GetUpperBoundOk

`func (o *VisualOutlierData) GetUpperBoundOk() (*float64, bool)`

GetUpperBoundOk returns a tuple with the UpperBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpperBound

`func (o *VisualOutlierData) SetUpperBound(v float64)`

SetUpperBound sets UpperBound field to given value.


### GetIsOutlier

`func (o *VisualOutlierData) GetIsOutlier() bool`

GetIsOutlier returns the IsOutlier field if non-nil, zero value otherwise.

### GetIsOutlierOk

`func (o *VisualOutlierData) GetIsOutlierOk() (*bool, bool)`

GetIsOutlierOk returns a tuple with the IsOutlier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOutlier

`func (o *VisualOutlierData) SetIsOutlier(v bool)`

SetIsOutlier sets IsOutlier field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


