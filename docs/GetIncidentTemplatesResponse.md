# GetIncidentTemplatesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Templates** | [**[]IncidentTemplate**](IncidentTemplate.md) | List of incident templates. | 

## Methods

### NewGetIncidentTemplatesResponse

`func NewGetIncidentTemplatesResponse(templates []IncidentTemplate, ) *GetIncidentTemplatesResponse`

NewGetIncidentTemplatesResponse instantiates a new GetIncidentTemplatesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIncidentTemplatesResponseWithDefaults

`func NewGetIncidentTemplatesResponseWithDefaults() *GetIncidentTemplatesResponse`

NewGetIncidentTemplatesResponseWithDefaults instantiates a new GetIncidentTemplatesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplates

`func (o *GetIncidentTemplatesResponse) GetTemplates() []IncidentTemplate`

GetTemplates returns the Templates field if non-nil, zero value otherwise.

### GetTemplatesOk

`func (o *GetIncidentTemplatesResponse) GetTemplatesOk() (*[]IncidentTemplate, bool)`

GetTemplatesOk returns a tuple with the Templates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplates

`func (o *GetIncidentTemplatesResponse) SetTemplates(v []IncidentTemplate)`

SetTemplates sets Templates field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


