# EndpointDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Unique name of endpoint. | 
**Url** | **string** | Address of endpoint. | 
**InputSchema** | **string** | Schema of the input table to endpoint. | 
**OutputSchema** | **string** | Schema of the output table from endpoint. | 
**Headers** | **string** | HTTP headers for endpoint. | 

## Methods

### NewEndpointDefinition

`func NewEndpointDefinition(name string, url string, inputSchema string, outputSchema string, headers string, ) *EndpointDefinition`

NewEndpointDefinition instantiates a new EndpointDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointDefinitionWithDefaults

`func NewEndpointDefinitionWithDefaults() *EndpointDefinition`

NewEndpointDefinitionWithDefaults instantiates a new EndpointDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EndpointDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EndpointDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EndpointDefinition) SetName(v string)`

SetName sets Name field to given value.


### GetUrl

`func (o *EndpointDefinition) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *EndpointDefinition) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *EndpointDefinition) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetInputSchema

`func (o *EndpointDefinition) GetInputSchema() string`

GetInputSchema returns the InputSchema field if non-nil, zero value otherwise.

### GetInputSchemaOk

`func (o *EndpointDefinition) GetInputSchemaOk() (*string, bool)`

GetInputSchemaOk returns a tuple with the InputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputSchema

`func (o *EndpointDefinition) SetInputSchema(v string)`

SetInputSchema sets InputSchema field to given value.


### GetOutputSchema

`func (o *EndpointDefinition) GetOutputSchema() string`

GetOutputSchema returns the OutputSchema field if non-nil, zero value otherwise.

### GetOutputSchemaOk

`func (o *EndpointDefinition) GetOutputSchemaOk() (*string, bool)`

GetOutputSchemaOk returns a tuple with the OutputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputSchema

`func (o *EndpointDefinition) SetOutputSchema(v string)`

SetOutputSchema sets OutputSchema field to given value.


### GetHeaders

`func (o *EndpointDefinition) GetHeaders() string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *EndpointDefinition) GetHeadersOk() (*string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *EndpointDefinition) SetHeaders(v string)`

SetHeaders sets Headers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


