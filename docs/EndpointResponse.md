# EndpointResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Identifier of endpoint. | 
**Name** | **string** | Unique name of endpoint. | 
**Url** | **string** | Address of endpoint. | 
**InputSchema** | **string** | Schema of the input table to endpoint. | 
**OutputSchema** | **string** | Schema of the output table from endpoint. | 

## Methods

### NewEndpointResponse

`func NewEndpointResponse(id string, name string, url string, inputSchema string, outputSchema string, ) *EndpointResponse`

NewEndpointResponse instantiates a new EndpointResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointResponseWithDefaults

`func NewEndpointResponseWithDefaults() *EndpointResponse`

NewEndpointResponseWithDefaults instantiates a new EndpointResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EndpointResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EndpointResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EndpointResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *EndpointResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EndpointResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EndpointResponse) SetName(v string)`

SetName sets Name field to given value.


### GetUrl

`func (o *EndpointResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *EndpointResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *EndpointResponse) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetInputSchema

`func (o *EndpointResponse) GetInputSchema() string`

GetInputSchema returns the InputSchema field if non-nil, zero value otherwise.

### GetInputSchemaOk

`func (o *EndpointResponse) GetInputSchemaOk() (*string, bool)`

GetInputSchemaOk returns a tuple with the InputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputSchema

`func (o *EndpointResponse) SetInputSchema(v string)`

SetInputSchema sets InputSchema field to given value.


### GetOutputSchema

`func (o *EndpointResponse) GetOutputSchema() string`

GetOutputSchema returns the OutputSchema field if non-nil, zero value otherwise.

### GetOutputSchemaOk

`func (o *EndpointResponse) GetOutputSchemaOk() (*string, bool)`

GetOutputSchemaOk returns a tuple with the OutputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputSchema

`func (o *EndpointResponse) SetOutputSchema(v string)`

SetOutputSchema sets OutputSchema field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


