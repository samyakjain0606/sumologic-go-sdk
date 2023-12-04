# ListRoleModelsResponseV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]GetRoleDefinitionV2**](GetRoleDefinitionV2.md) | List of roles. | 
**Next** | Pointer to **string** | Next continuation token. | [optional] 

## Methods

### NewListRoleModelsResponseV2

`func NewListRoleModelsResponseV2(data []GetRoleDefinitionV2, ) *ListRoleModelsResponseV2`

NewListRoleModelsResponseV2 instantiates a new ListRoleModelsResponseV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListRoleModelsResponseV2WithDefaults

`func NewListRoleModelsResponseV2WithDefaults() *ListRoleModelsResponseV2`

NewListRoleModelsResponseV2WithDefaults instantiates a new ListRoleModelsResponseV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListRoleModelsResponseV2) GetData() []GetRoleDefinitionV2`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListRoleModelsResponseV2) GetDataOk() (*[]GetRoleDefinitionV2, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListRoleModelsResponseV2) SetData(v []GetRoleDefinitionV2)`

SetData sets Data field to given value.


### GetNext

`func (o *ListRoleModelsResponseV2) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *ListRoleModelsResponseV2) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *ListRoleModelsResponseV2) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *ListRoleModelsResponseV2) HasNext() bool`

HasNext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


