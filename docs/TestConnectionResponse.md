# TestConnectionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusCode** | **int32** | Status code of the response of the connection test. | 
**ResponseContent** | **string** | Content of the response of the connection test. | 
**AlertStatusCode** | Pointer to **int32** | Status code of the response of alert payload test. | [optional] 
**AlertResponseContent** | Pointer to **string** | Content of the response of alert payload test. | [optional] 
**ResolutionStatusCode** | Pointer to **int32** | Status code of the response of resolution payload test. | [optional] 
**ResolutionResponseContent** | Pointer to **string** | Content of the response of resolution payload test. | [optional] 

## Methods

### NewTestConnectionResponse

`func NewTestConnectionResponse(statusCode int32, responseContent string, ) *TestConnectionResponse`

NewTestConnectionResponse instantiates a new TestConnectionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestConnectionResponseWithDefaults

`func NewTestConnectionResponseWithDefaults() *TestConnectionResponse`

NewTestConnectionResponseWithDefaults instantiates a new TestConnectionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusCode

`func (o *TestConnectionResponse) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *TestConnectionResponse) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *TestConnectionResponse) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.


### GetResponseContent

`func (o *TestConnectionResponse) GetResponseContent() string`

GetResponseContent returns the ResponseContent field if non-nil, zero value otherwise.

### GetResponseContentOk

`func (o *TestConnectionResponse) GetResponseContentOk() (*string, bool)`

GetResponseContentOk returns a tuple with the ResponseContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContent

`func (o *TestConnectionResponse) SetResponseContent(v string)`

SetResponseContent sets ResponseContent field to given value.


### GetAlertStatusCode

`func (o *TestConnectionResponse) GetAlertStatusCode() int32`

GetAlertStatusCode returns the AlertStatusCode field if non-nil, zero value otherwise.

### GetAlertStatusCodeOk

`func (o *TestConnectionResponse) GetAlertStatusCodeOk() (*int32, bool)`

GetAlertStatusCodeOk returns a tuple with the AlertStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertStatusCode

`func (o *TestConnectionResponse) SetAlertStatusCode(v int32)`

SetAlertStatusCode sets AlertStatusCode field to given value.

### HasAlertStatusCode

`func (o *TestConnectionResponse) HasAlertStatusCode() bool`

HasAlertStatusCode returns a boolean if a field has been set.

### GetAlertResponseContent

`func (o *TestConnectionResponse) GetAlertResponseContent() string`

GetAlertResponseContent returns the AlertResponseContent field if non-nil, zero value otherwise.

### GetAlertResponseContentOk

`func (o *TestConnectionResponse) GetAlertResponseContentOk() (*string, bool)`

GetAlertResponseContentOk returns a tuple with the AlertResponseContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertResponseContent

`func (o *TestConnectionResponse) SetAlertResponseContent(v string)`

SetAlertResponseContent sets AlertResponseContent field to given value.

### HasAlertResponseContent

`func (o *TestConnectionResponse) HasAlertResponseContent() bool`

HasAlertResponseContent returns a boolean if a field has been set.

### GetResolutionStatusCode

`func (o *TestConnectionResponse) GetResolutionStatusCode() int32`

GetResolutionStatusCode returns the ResolutionStatusCode field if non-nil, zero value otherwise.

### GetResolutionStatusCodeOk

`func (o *TestConnectionResponse) GetResolutionStatusCodeOk() (*int32, bool)`

GetResolutionStatusCodeOk returns a tuple with the ResolutionStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionStatusCode

`func (o *TestConnectionResponse) SetResolutionStatusCode(v int32)`

SetResolutionStatusCode sets ResolutionStatusCode field to given value.

### HasResolutionStatusCode

`func (o *TestConnectionResponse) HasResolutionStatusCode() bool`

HasResolutionStatusCode returns a boolean if a field has been set.

### GetResolutionResponseContent

`func (o *TestConnectionResponse) GetResolutionResponseContent() string`

GetResolutionResponseContent returns the ResolutionResponseContent field if non-nil, zero value otherwise.

### GetResolutionResponseContentOk

`func (o *TestConnectionResponse) GetResolutionResponseContentOk() (*string, bool)`

GetResolutionResponseContentOk returns a tuple with the ResolutionResponseContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionResponseContent

`func (o *TestConnectionResponse) SetResolutionResponseContent(v string)`

SetResolutionResponseContent sets ResolutionResponseContent field to given value.

### HasResolutionResponseContent

`func (o *TestConnectionResponse) HasResolutionResponseContent() bool`

HasResolutionResponseContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


