# ListPermissionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PermissionStatements** | [**[]PermissionStatement**](PermissionStatement.md) | A list of permission statements. | 

## Methods

### NewListPermissionsResponse

`func NewListPermissionsResponse(permissionStatements []PermissionStatement, ) *ListPermissionsResponse`

NewListPermissionsResponse instantiates a new ListPermissionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPermissionsResponseWithDefaults

`func NewListPermissionsResponseWithDefaults() *ListPermissionsResponse`

NewListPermissionsResponseWithDefaults instantiates a new ListPermissionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissionStatements

`func (o *ListPermissionsResponse) GetPermissionStatements() []PermissionStatement`

GetPermissionStatements returns the PermissionStatements field if non-nil, zero value otherwise.

### GetPermissionStatementsOk

`func (o *ListPermissionsResponse) GetPermissionStatementsOk() (*[]PermissionStatement, bool)`

GetPermissionStatementsOk returns a tuple with the PermissionStatements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionStatements

`func (o *ListPermissionsResponse) SetPermissionStatements(v []PermissionStatement)`

SetPermissionStatements sets PermissionStatements field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


