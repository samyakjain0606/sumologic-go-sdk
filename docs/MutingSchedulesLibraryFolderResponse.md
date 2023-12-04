# MutingSchedulesLibraryFolderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Permissions** | **[]string** | Aggregated permission summary for the calling user. If detailed permission statements are required, please call list permissions endpoint. | 
**Children** | [**[]MutingSchedulesLibraryBaseResponse**](MutingSchedulesLibraryBaseResponse.md) | Children of the folder. NOTE: Permissions field will not be filled (empty list) for children. | 

## Methods

### NewMutingSchedulesLibraryFolderResponse

`func NewMutingSchedulesLibraryFolderResponse(permissions []string, children []MutingSchedulesLibraryBaseResponse, ) *MutingSchedulesLibraryFolderResponse`

NewMutingSchedulesLibraryFolderResponse instantiates a new MutingSchedulesLibraryFolderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMutingSchedulesLibraryFolderResponseWithDefaults

`func NewMutingSchedulesLibraryFolderResponseWithDefaults() *MutingSchedulesLibraryFolderResponse`

NewMutingSchedulesLibraryFolderResponseWithDefaults instantiates a new MutingSchedulesLibraryFolderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissions

`func (o *MutingSchedulesLibraryFolderResponse) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *MutingSchedulesLibraryFolderResponse) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *MutingSchedulesLibraryFolderResponse) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.


### GetChildren

`func (o *MutingSchedulesLibraryFolderResponse) GetChildren() []MutingSchedulesLibraryBaseResponse`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *MutingSchedulesLibraryFolderResponse) GetChildrenOk() (*[]MutingSchedulesLibraryBaseResponse, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *MutingSchedulesLibraryFolderResponse) SetChildren(v []MutingSchedulesLibraryBaseResponse)`

SetChildren sets Children field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


