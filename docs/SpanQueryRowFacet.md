# SpanQueryRowFacet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the field facet. | 
**Cardinality** | **int32** | The number of unique values this field occured. | 
**DataType** | **string** | Data type of the field. | 
**InSchema** | Pointer to **bool** | Indicates whether the field is available in the span schema. | [optional] 
**ValueFrequency** | Pointer to **map[string]int64** | Map of field value frequencies. | [optional] 

## Methods

### NewSpanQueryRowFacet

`func NewSpanQueryRowFacet(name string, cardinality int32, dataType string, ) *SpanQueryRowFacet`

NewSpanQueryRowFacet instantiates a new SpanQueryRowFacet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpanQueryRowFacetWithDefaults

`func NewSpanQueryRowFacetWithDefaults() *SpanQueryRowFacet`

NewSpanQueryRowFacetWithDefaults instantiates a new SpanQueryRowFacet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SpanQueryRowFacet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SpanQueryRowFacet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SpanQueryRowFacet) SetName(v string)`

SetName sets Name field to given value.


### GetCardinality

`func (o *SpanQueryRowFacet) GetCardinality() int32`

GetCardinality returns the Cardinality field if non-nil, zero value otherwise.

### GetCardinalityOk

`func (o *SpanQueryRowFacet) GetCardinalityOk() (*int32, bool)`

GetCardinalityOk returns a tuple with the Cardinality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardinality

`func (o *SpanQueryRowFacet) SetCardinality(v int32)`

SetCardinality sets Cardinality field to given value.


### GetDataType

`func (o *SpanQueryRowFacet) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *SpanQueryRowFacet) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *SpanQueryRowFacet) SetDataType(v string)`

SetDataType sets DataType field to given value.


### GetInSchema

`func (o *SpanQueryRowFacet) GetInSchema() bool`

GetInSchema returns the InSchema field if non-nil, zero value otherwise.

### GetInSchemaOk

`func (o *SpanQueryRowFacet) GetInSchemaOk() (*bool, bool)`

GetInSchemaOk returns a tuple with the InSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInSchema

`func (o *SpanQueryRowFacet) SetInSchema(v bool)`

SetInSchema sets InSchema field to given value.

### HasInSchema

`func (o *SpanQueryRowFacet) HasInSchema() bool`

HasInSchema returns a boolean if a field has been set.

### GetValueFrequency

`func (o *SpanQueryRowFacet) GetValueFrequency() map[string]int64`

GetValueFrequency returns the ValueFrequency field if non-nil, zero value otherwise.

### GetValueFrequencyOk

`func (o *SpanQueryRowFacet) GetValueFrequencyOk() (*map[string]int64, bool)`

GetValueFrequencyOk returns a tuple with the ValueFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueFrequency

`func (o *SpanQueryRowFacet) SetValueFrequency(v map[string]int64)`

SetValueFrequency sets ValueFrequency field to given value.

### HasValueFrequency

`func (o *SpanQueryRowFacet) HasValueFrequency() bool`

HasValueFrequency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


