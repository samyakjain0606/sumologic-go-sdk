# CompliancePeriodRef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComplianceRefType** | **string** | Type of reference to the compliance period. Must be &#x60;Relative&#x60;. | 
**RelativeShift** | Pointer to **int32** | Relative shift of compliance period from the latest/current compliance period. | [optional] 

## Methods

### NewCompliancePeriodRef

`func NewCompliancePeriodRef(complianceRefType string, ) *CompliancePeriodRef`

NewCompliancePeriodRef instantiates a new CompliancePeriodRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompliancePeriodRefWithDefaults

`func NewCompliancePeriodRefWithDefaults() *CompliancePeriodRef`

NewCompliancePeriodRefWithDefaults instantiates a new CompliancePeriodRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComplianceRefType

`func (o *CompliancePeriodRef) GetComplianceRefType() string`

GetComplianceRefType returns the ComplianceRefType field if non-nil, zero value otherwise.

### GetComplianceRefTypeOk

`func (o *CompliancePeriodRef) GetComplianceRefTypeOk() (*string, bool)`

GetComplianceRefTypeOk returns a tuple with the ComplianceRefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceRefType

`func (o *CompliancePeriodRef) SetComplianceRefType(v string)`

SetComplianceRefType sets ComplianceRefType field to given value.


### GetRelativeShift

`func (o *CompliancePeriodRef) GetRelativeShift() int32`

GetRelativeShift returns the RelativeShift field if non-nil, zero value otherwise.

### GetRelativeShiftOk

`func (o *CompliancePeriodRef) GetRelativeShiftOk() (*int32, bool)`

GetRelativeShiftOk returns a tuple with the RelativeShift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativeShift

`func (o *CompliancePeriodRef) SetRelativeShift(v int32)`

SetRelativeShift sets RelativeShift field to given value.

### HasRelativeShift

`func (o *CompliancePeriodRef) HasRelativeShift() bool`

HasRelativeShift returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


