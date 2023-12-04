# SpanQueryFieldDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** | Filter field name. | 
**FieldType** | **string** | Indicates the kind of a field. Possible values: &#x60;SpanAttribute&#x60;, &#x60;SpanEventAttribute&#x60;. | [default to "SpanAttribute"]
**ValueListing** | Pointer to **bool** | Indicates whether values for this field can be listed. | [optional] 
**Description** | Pointer to **string** | Short description of the field. | [optional] 
**Type** | **string** | The type the values of this field will have. Possible values: &#x60;DoubleTracingValue&#x60;, &#x60;IntegerTracingValue&#x60;, &#x60;StringTracingValue&#x60;, &#x60;DateTimeTracingValue&#x60;. | 
**NoValuesReason** | Pointer to [**NoTraceFieldValuesReason**](NoTraceFieldValuesReason.md) |  | [optional] 
**InSchema** | **bool** | Indicates whether the field is available in the schema. | 

## Methods

### NewSpanQueryFieldDetail

`func NewSpanQueryFieldDetail(field string, fieldType string, type_ string, inSchema bool, ) *SpanQueryFieldDetail`

NewSpanQueryFieldDetail instantiates a new SpanQueryFieldDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpanQueryFieldDetailWithDefaults

`func NewSpanQueryFieldDetailWithDefaults() *SpanQueryFieldDetail`

NewSpanQueryFieldDetailWithDefaults instantiates a new SpanQueryFieldDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *SpanQueryFieldDetail) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *SpanQueryFieldDetail) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *SpanQueryFieldDetail) SetField(v string)`

SetField sets Field field to given value.


### GetFieldType

`func (o *SpanQueryFieldDetail) GetFieldType() string`

GetFieldType returns the FieldType field if non-nil, zero value otherwise.

### GetFieldTypeOk

`func (o *SpanQueryFieldDetail) GetFieldTypeOk() (*string, bool)`

GetFieldTypeOk returns a tuple with the FieldType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldType

`func (o *SpanQueryFieldDetail) SetFieldType(v string)`

SetFieldType sets FieldType field to given value.


### GetValueListing

`func (o *SpanQueryFieldDetail) GetValueListing() bool`

GetValueListing returns the ValueListing field if non-nil, zero value otherwise.

### GetValueListingOk

`func (o *SpanQueryFieldDetail) GetValueListingOk() (*bool, bool)`

GetValueListingOk returns a tuple with the ValueListing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueListing

`func (o *SpanQueryFieldDetail) SetValueListing(v bool)`

SetValueListing sets ValueListing field to given value.

### HasValueListing

`func (o *SpanQueryFieldDetail) HasValueListing() bool`

HasValueListing returns a boolean if a field has been set.

### GetDescription

`func (o *SpanQueryFieldDetail) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SpanQueryFieldDetail) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SpanQueryFieldDetail) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SpanQueryFieldDetail) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *SpanQueryFieldDetail) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SpanQueryFieldDetail) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SpanQueryFieldDetail) SetType(v string)`

SetType sets Type field to given value.


### GetNoValuesReason

`func (o *SpanQueryFieldDetail) GetNoValuesReason() NoTraceFieldValuesReason`

GetNoValuesReason returns the NoValuesReason field if non-nil, zero value otherwise.

### GetNoValuesReasonOk

`func (o *SpanQueryFieldDetail) GetNoValuesReasonOk() (*NoTraceFieldValuesReason, bool)`

GetNoValuesReasonOk returns a tuple with the NoValuesReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoValuesReason

`func (o *SpanQueryFieldDetail) SetNoValuesReason(v NoTraceFieldValuesReason)`

SetNoValuesReason sets NoValuesReason field to given value.

### HasNoValuesReason

`func (o *SpanQueryFieldDetail) HasNoValuesReason() bool`

HasNoValuesReason returns a boolean if a field has been set.

### GetInSchema

`func (o *SpanQueryFieldDetail) GetInSchema() bool`

GetInSchema returns the InSchema field if non-nil, zero value otherwise.

### GetInSchemaOk

`func (o *SpanQueryFieldDetail) GetInSchemaOk() (*bool, bool)`

GetInSchemaOk returns a tuple with the InSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInSchema

`func (o *SpanQueryFieldDetail) SetInSchema(v bool)`

SetInSchema sets InSchema field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


