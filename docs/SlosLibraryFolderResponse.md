# SlosLibraryFolderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Permissions** | **[]string** | Aggregated permission summary for the calling user. If detailed permission statements are required, please call list permissions endpoint. | 
**Children** | [**[]SlosLibraryBaseResponse**](SlosLibraryBaseResponse.md) | Children of the folder. NOTE: Permissions field will not be filled (empty list) for children. | 

## Methods

### NewSlosLibraryFolderResponse

`func NewSlosLibraryFolderResponse(permissions []string, children []SlosLibraryBaseResponse, ) *SlosLibraryFolderResponse`

NewSlosLibraryFolderResponse instantiates a new SlosLibraryFolderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlosLibraryFolderResponseWithDefaults

`func NewSlosLibraryFolderResponseWithDefaults() *SlosLibraryFolderResponse`

NewSlosLibraryFolderResponseWithDefaults instantiates a new SlosLibraryFolderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissions

`func (o *SlosLibraryFolderResponse) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *SlosLibraryFolderResponse) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *SlosLibraryFolderResponse) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.


### GetChildren

`func (o *SlosLibraryFolderResponse) GetChildren() []SlosLibraryBaseResponse`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *SlosLibraryFolderResponse) GetChildrenOk() (*[]SlosLibraryBaseResponse, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *SlosLibraryFolderResponse) SetChildren(v []SlosLibraryBaseResponse)`

SetChildren sets Children field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


