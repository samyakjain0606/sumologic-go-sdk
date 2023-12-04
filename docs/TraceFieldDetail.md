# TraceFieldDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** | Filter field name. | 
**FieldType** | **string** | Indicates the kind of a field. Possible values: &#x60;SpanAttribute&#x60;, &#x60;SpanEventAttribute&#x60;. | [default to "SpanAttribute"]
**ValueListing** | Pointer to **bool** | Indicates whether values for this field can be listed. | [optional] 
**Description** | Pointer to **string** | Short description of the field. | [optional] 
**Type** | **string** | The type the values of this field will have. Possible values: &#x60;DoubleTracingValue&#x60;, &#x60;IntegerTracingValue&#x60;, &#x60;StringTracingValue&#x60;, &#x60;DateTimeTracingValue&#x60;. | 
**NoValuesReason** | Pointer to [**NoTraceFieldValuesReason**](NoTraceFieldValuesReason.md) |  | [optional] 

## Methods

### NewTraceFieldDetail

`func NewTraceFieldDetail(field string, fieldType string, type_ string, ) *TraceFieldDetail`

NewTraceFieldDetail instantiates a new TraceFieldDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTraceFieldDetailWithDefaults

`func NewTraceFieldDetailWithDefaults() *TraceFieldDetail`

NewTraceFieldDetailWithDefaults instantiates a new TraceFieldDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *TraceFieldDetail) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *TraceFieldDetail) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *TraceFieldDetail) SetField(v string)`

SetField sets Field field to given value.


### GetFieldType

`func (o *TraceFieldDetail) GetFieldType() string`

GetFieldType returns the FieldType field if non-nil, zero value otherwise.

### GetFieldTypeOk

`func (o *TraceFieldDetail) GetFieldTypeOk() (*string, bool)`

GetFieldTypeOk returns a tuple with the FieldType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldType

`func (o *TraceFieldDetail) SetFieldType(v string)`

SetFieldType sets FieldType field to given value.


### GetValueListing

`func (o *TraceFieldDetail) GetValueListing() bool`

GetValueListing returns the ValueListing field if non-nil, zero value otherwise.

### GetValueListingOk

`func (o *TraceFieldDetail) GetValueListingOk() (*bool, bool)`

GetValueListingOk returns a tuple with the ValueListing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueListing

`func (o *TraceFieldDetail) SetValueListing(v bool)`

SetValueListing sets ValueListing field to given value.

### HasValueListing

`func (o *TraceFieldDetail) HasValueListing() bool`

HasValueListing returns a boolean if a field has been set.

### GetDescription

`func (o *TraceFieldDetail) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TraceFieldDetail) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TraceFieldDetail) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TraceFieldDetail) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *TraceFieldDetail) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TraceFieldDetail) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TraceFieldDetail) SetType(v string)`

SetType sets Type field to given value.


### GetNoValuesReason

`func (o *TraceFieldDetail) GetNoValuesReason() NoTraceFieldValuesReason`

GetNoValuesReason returns the NoValuesReason field if non-nil, zero value otherwise.

### GetNoValuesReasonOk

`func (o *TraceFieldDetail) GetNoValuesReasonOk() (*NoTraceFieldValuesReason, bool)`

GetNoValuesReasonOk returns a tuple with the NoValuesReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoValuesReason

`func (o *TraceFieldDetail) SetNoValuesReason(v NoTraceFieldValuesReason)`

SetNoValuesReason sets NoValuesReason field to given value.

### HasNoValuesReason

`func (o *TraceFieldDetail) HasNoValuesReason() bool`

HasNoValuesReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


