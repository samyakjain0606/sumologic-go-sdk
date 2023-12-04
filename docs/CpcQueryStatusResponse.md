# CpcQueryStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryRows** | [**[]CpcQueryRowStatus**](CpcQueryRowStatus.md) | A list of statuses on a per query row basis. | 
**Status** | **string** | Status of the query. Possible values: &#x60;Processing&#x60;, &#x60;Finished&#x60;, &#x60;Error&#x60;, &#x60;Canceled&#x60;. | 

## Methods

### NewCpcQueryStatusResponse

`func NewCpcQueryStatusResponse(queryRows []CpcQueryRowStatus, status string, ) *CpcQueryStatusResponse`

NewCpcQueryStatusResponse instantiates a new CpcQueryStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCpcQueryStatusResponseWithDefaults

`func NewCpcQueryStatusResponseWithDefaults() *CpcQueryStatusResponse`

NewCpcQueryStatusResponseWithDefaults instantiates a new CpcQueryStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueryRows

`func (o *CpcQueryStatusResponse) GetQueryRows() []CpcQueryRowStatus`

GetQueryRows returns the QueryRows field if non-nil, zero value otherwise.

### GetQueryRowsOk

`func (o *CpcQueryStatusResponse) GetQueryRowsOk() (*[]CpcQueryRowStatus, bool)`

GetQueryRowsOk returns a tuple with the QueryRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryRows

`func (o *CpcQueryStatusResponse) SetQueryRows(v []CpcQueryRowStatus)`

SetQueryRows sets QueryRows field to given value.


### GetStatus

`func (o *CpcQueryStatusResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CpcQueryStatusResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CpcQueryStatusResponse) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


