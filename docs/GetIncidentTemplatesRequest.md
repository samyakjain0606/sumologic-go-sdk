# GetIncidentTemplatesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | Optional CloudSOAR domain URL to use for the API call to get incident templates. | [optional] 
**AuthHeader** | Pointer to **string** | Optional CloudSOAR authorization header to use for the API call to get incident templates. | [optional] 
**ConnectionId** | Pointer to **string** | Optional connectionId to get incident templates for an existing CloudSOAR connection. If provided, the authHeader and url will be taken from the existing connection object. | [optional] 

## Methods

### NewGetIncidentTemplatesRequest

`func NewGetIncidentTemplatesRequest() *GetIncidentTemplatesRequest`

NewGetIncidentTemplatesRequest instantiates a new GetIncidentTemplatesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIncidentTemplatesRequestWithDefaults

`func NewGetIncidentTemplatesRequestWithDefaults() *GetIncidentTemplatesRequest`

NewGetIncidentTemplatesRequestWithDefaults instantiates a new GetIncidentTemplatesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *GetIncidentTemplatesRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GetIncidentTemplatesRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GetIncidentTemplatesRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *GetIncidentTemplatesRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetAuthHeader

`func (o *GetIncidentTemplatesRequest) GetAuthHeader() string`

GetAuthHeader returns the AuthHeader field if non-nil, zero value otherwise.

### GetAuthHeaderOk

`func (o *GetIncidentTemplatesRequest) GetAuthHeaderOk() (*string, bool)`

GetAuthHeaderOk returns a tuple with the AuthHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthHeader

`func (o *GetIncidentTemplatesRequest) SetAuthHeader(v string)`

SetAuthHeader sets AuthHeader field to given value.

### HasAuthHeader

`func (o *GetIncidentTemplatesRequest) HasAuthHeader() bool`

HasAuthHeader returns a boolean if a field has been set.

### GetConnectionId

`func (o *GetIncidentTemplatesRequest) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *GetIncidentTemplatesRequest) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *GetIncidentTemplatesRequest) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *GetIncidentTemplatesRequest) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


