# \RoleManagementAPI

All URIs are relative to *https://api.au.sumologic.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignRoleToUser**](RoleManagementAPI.md#AssignRoleToUser) | **Put** /v1/roles/{roleId}/users/{userId} | Assign a role to a user.
[**CreateRole**](RoleManagementAPI.md#CreateRole) | **Post** /v1/roles | Create a new role.
[**DeleteRole**](RoleManagementAPI.md#DeleteRole) | **Delete** /v1/roles/{id} | Delete a role.
[**GetRole**](RoleManagementAPI.md#GetRole) | **Get** /v1/roles/{id} | Get a role.
[**ListRoles**](RoleManagementAPI.md#ListRoles) | **Get** /v1/roles | Get a list of roles.
[**RemoveRoleFromUser**](RoleManagementAPI.md#RemoveRoleFromUser) | **Delete** /v1/roles/{roleId}/users/{userId} | Remove role from a user.
[**UpdateRole**](RoleManagementAPI.md#UpdateRole) | **Put** /v1/roles/{id} | Update a role.



## AssignRoleToUser

> RoleModel AssignRoleToUser(ctx, roleId, userId).Execute()

Assign a role to a user.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    roleId := "roleId_example" // string | Identifier of the role to assign.
    userId := "userId_example" // string | Identifier of the user to assign the role to.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleManagementAPI.AssignRoleToUser(context.Background(), roleId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleManagementAPI.AssignRoleToUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssignRoleToUser`: RoleModel
    fmt.Fprintf(os.Stdout, "Response from `RoleManagementAPI.AssignRoleToUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **string** | Identifier of the role to assign. | 
**userId** | **string** | Identifier of the user to assign the role to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignRoleToUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RoleModel**](RoleModel.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRole

> RoleModel CreateRole(ctx).CreateRoleDefinition(createRoleDefinition).Execute()

Create a new role.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    createRoleDefinition := *openapiclient.NewCreateRoleDefinition("DataAdmin") // CreateRoleDefinition | Information about the new role.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleManagementAPI.CreateRole(context.Background()).CreateRoleDefinition(createRoleDefinition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleManagementAPI.CreateRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRole`: RoleModel
    fmt.Fprintf(os.Stdout, "Response from `RoleManagementAPI.CreateRole`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRoleDefinition** | [**CreateRoleDefinition**](CreateRoleDefinition.md) | Information about the new role. | 

### Return type

[**RoleModel**](RoleModel.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRole

> DeleteRole(ctx, id).Execute()

Delete a role.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | Identifier of the role to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RoleManagementAPI.DeleteRole(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleManagementAPI.DeleteRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the role to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRole

> RoleModel GetRole(ctx, id).Execute()

Get a role.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | Identifier of the role to fetch.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleManagementAPI.GetRole(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleManagementAPI.GetRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRole`: RoleModel
    fmt.Fprintf(os.Stdout, "Response from `RoleManagementAPI.GetRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the role to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RoleModel**](RoleModel.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRoles

> ListRoleModelsResponse ListRoles(ctx).Limit(limit).Token(token).SortBy(sortBy).Name(name).Execute()

Get a list of roles.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    limit := int32(56) // int32 | Limit the number of roles returned in the response. The number of roles returned may be less than the `limit`. (optional) (default to 100)
    token := "token_example" // string | Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. `token` is set to null when no more pages are left. (optional)
    sortBy := "sortBy_example" // string | Sort the list of roles by the `name` field. (optional)
    name := "name_example" // string | Only return roles matching the given name. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleManagementAPI.ListRoles(context.Background()).Limit(limit).Token(token).SortBy(sortBy).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleManagementAPI.ListRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRoles`: ListRoleModelsResponse
    fmt.Fprintf(os.Stdout, "Response from `RoleManagementAPI.ListRoles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit the number of roles returned in the response. The number of roles returned may be less than the &#x60;limit&#x60;. | [default to 100]
 **token** | **string** | Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. &#x60;token&#x60; is set to null when no more pages are left. | 
 **sortBy** | **string** | Sort the list of roles by the &#x60;name&#x60; field. | 
 **name** | **string** | Only return roles matching the given name. | 

### Return type

[**ListRoleModelsResponse**](ListRoleModelsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveRoleFromUser

> RemoveRoleFromUser(ctx, roleId, userId).Execute()

Remove role from a user.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    roleId := "roleId_example" // string | Identifier of the role to delete.
    userId := "userId_example" // string | Identifier of the user to remove the role from.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RoleManagementAPI.RemoveRoleFromUser(context.Background(), roleId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleManagementAPI.RemoveRoleFromUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **string** | Identifier of the role to delete. | 
**userId** | **string** | Identifier of the user to remove the role from. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveRoleFromUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRole

> RoleModel UpdateRole(ctx, id).UpdateRoleDefinition(updateRoleDefinition).Execute()

Update a role.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | Identifier of the role to update.
    updateRoleDefinition := *openapiclient.NewUpdateRoleDefinition("DataAdmin", "Manage data of the org.", "!_sourceCategory=billing", []string{"Users_example"}, []string{"Capabilities_example"}) // UpdateRoleDefinition | Information to update about the role.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleManagementAPI.UpdateRole(context.Background(), id).UpdateRoleDefinition(updateRoleDefinition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleManagementAPI.UpdateRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRole`: RoleModel
    fmt.Fprintf(os.Stdout, "Response from `RoleManagementAPI.UpdateRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the role to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRoleDefinition** | [**UpdateRoleDefinition**](UpdateRoleDefinition.md) | Information to update about the role. | 

### Return type

[**RoleModel**](RoleModel.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

