# \MonitorsLibraryManagementAPI

All URIs are relative to *https://api.au.sumologic.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DisableMonitorByIds**](MonitorsLibraryManagementAPI.md#DisableMonitorByIds) | **Put** /v1/monitors/disable | Disable monitors.
[**GetMonitorUsageInfo**](MonitorsLibraryManagementAPI.md#GetMonitorUsageInfo) | **Get** /v1/monitors/usageInfo | Usage info of monitors.
[**GetMonitorsFullPath**](MonitorsLibraryManagementAPI.md#GetMonitorsFullPath) | **Get** /v1/monitors/{id}/path | Get the path of a monitor or folder.
[**GetMonitorsLibraryRoot**](MonitorsLibraryManagementAPI.md#GetMonitorsLibraryRoot) | **Get** /v1/monitors/root | Get the root monitors folder.
[**MonitorsCopy**](MonitorsLibraryManagementAPI.md#MonitorsCopy) | **Post** /v1/monitors/{id}/copy | Copy a monitor or folder.
[**MonitorsCreate**](MonitorsLibraryManagementAPI.md#MonitorsCreate) | **Post** /v1/monitors | Create a monitor or folder. 
[**MonitorsDeleteById**](MonitorsLibraryManagementAPI.md#MonitorsDeleteById) | **Delete** /v1/monitors/{id} | Delete a monitor or folder. 
[**MonitorsDeleteByIds**](MonitorsLibraryManagementAPI.md#MonitorsDeleteByIds) | **Delete** /v1/monitors | Bulk delete a monitor or folder. 
[**MonitorsExportItem**](MonitorsLibraryManagementAPI.md#MonitorsExportItem) | **Get** /v1/monitors/{id}/export | Export a monitor or folder.
[**MonitorsGetByPath**](MonitorsLibraryManagementAPI.md#MonitorsGetByPath) | **Get** /v1/monitors/path | Read a monitor or folder by its path.
[**MonitorsImportItem**](MonitorsLibraryManagementAPI.md#MonitorsImportItem) | **Post** /v1/monitors/{parentId}/import | Import a monitor or folder.
[**MonitorsMove**](MonitorsLibraryManagementAPI.md#MonitorsMove) | **Post** /v1/monitors/{id}/move | Move a monitor or folder.
[**MonitorsReadById**](MonitorsLibraryManagementAPI.md#MonitorsReadById) | **Get** /v1/monitors/{id} | Get a monitor or folder.
[**MonitorsReadByIds**](MonitorsLibraryManagementAPI.md#MonitorsReadByIds) | **Get** /v1/monitors | Bulk read a monitor or folder.
[**MonitorsReadPermissionSummariesByIdGroupBySubjects**](MonitorsLibraryManagementAPI.md#MonitorsReadPermissionSummariesByIdGroupBySubjects) | **Get** /v1/monitors/{id}/permissionSummariesBySubjects | List permission summaries for a monitor or folder. 
[**MonitorsReadPermissionsById**](MonitorsLibraryManagementAPI.md#MonitorsReadPermissionsById) | **Get** /v1/monitors/{id}/permissions | List explicit permissions on monitor or folder. 
[**MonitorsRevokePermissions**](MonitorsLibraryManagementAPI.md#MonitorsRevokePermissions) | **Put** /v1/monitors/permissions/revoke | Revoke all permissions on monitor or folder. 
[**MonitorsSearch**](MonitorsLibraryManagementAPI.md#MonitorsSearch) | **Get** /v1/monitors/search | Search for a monitor or folder.
[**MonitorsSetPermissions**](MonitorsLibraryManagementAPI.md#MonitorsSetPermissions) | **Put** /v1/monitors/permissions/set | Set permissions on monitor or folder. 
[**MonitorsUpdateById**](MonitorsLibraryManagementAPI.md#MonitorsUpdateById) | **Put** /v1/monitors/{id} | Update a monitor or folder. 



## DisableMonitorByIds

> DisableMonitorResponse DisableMonitorByIds(ctx).Ids(ids).Execute()

Disable monitors.



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
    ids := []string{"Inner_example"} // []string | A comma-separated list of identifiers.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorsLibraryManagementAPI.DisableMonitorByIds(context.Background()).Ids(ids).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorsLibraryManagementAPI.DisableMonitorByIds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DisableMonitorByIds`: DisableMonitorResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitorsLibraryManagementAPI.DisableMonitorByIds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDisableMonitorByIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **[]string** | A comma-separated list of identifiers. | 

### Return type

[**DisableMonitorResponse**](DisableMonitorResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMonitorUsageInfo

> []MonitorUsage GetMonitorUsageInfo(ctx).Execute()

Usage info of monitors.



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorsLibraryManagementAPI.GetMonitorUsageInfo(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorsLibraryManagementAPI.GetMonitorUsageInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMonitorUsageInfo`: []MonitorUsage
    fmt.Fprintf(os.Stdout, "Response from `MonitorsLibraryManagementAPI.GetMonitorUsageInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMonitorUsageInfoRequest struct via the builder pattern


### Return type

[**[]MonitorUsage**](MonitorUsage.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMonitorsFullPath

> Path GetMonitorsFullPath(ctx, id).Execute()

Get the path of a monitor or folder.



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
    id := "id_example" // string | Identifier of the monitor or folder.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorsLibraryManagementAPI.GetMonitorsFullPath(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorsLibraryManagementAPI.GetMonitorsFullPath``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMonitorsFullPath`: Path
    fmt.Fprintf(os.Stdout, "Response from `MonitorsLibraryManagementAPI.GetMonitorsFullPath`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the monitor or folder. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMonitorsFullPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Path**](Path.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMonitorsLibraryRoot

> MonitorsLibraryFolderResponse GetMonitorsLibraryRoot(ctx).Execute()

Get the root monitors folder.



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorsLibraryManagementAPI.GetMonitorsLibraryRoot(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorsLibraryManagementAPI.GetMonitorsLibraryRoot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMonitorsLibraryRoot`: MonitorsLibraryFolderResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitorsLibraryManagementAPI.GetMonitorsLibraryRoot`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMonitorsLibraryRootRequest struct via the builder pattern


### Return type

[**MonitorsLibraryFolderResponse**](MonitorsLibraryFolderResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitorsCopy

> MonitorsLibraryBaseResponse MonitorsCopy(ctx, id).ContentCopyParams(contentCopyParams).Execute()

Copy a monitor or folder.



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
    id := "id_example" // string | Identifier of the monitor or folder to copy.
    contentCopyParams := *openapiclient.NewContentCopyParams("ParentId_example") // ContentCopyParams | Fields include:   1) Identifier of the parent folder to copy to.   2) Optionally provide a new name.   3) Optionally provide a new description.   4) Optionally set to true if you want to copy and preserve the locked status. Requires `LockMonitors` capability.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorsLibraryManagementAPI.MonitorsCopy(context.Background(), id).ContentCopyParams(contentCopyParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorsLibraryManagementAPI.MonitorsCopy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MonitorsCopy`: MonitorsLibraryBaseResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitorsLibraryManagementAPI.MonitorsCopy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the monitor or folder to copy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMonitorsCopyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentCopyParams** | [**ContentCopyParams**](ContentCopyParams.md) | Fields include:   1) Identifier of the parent folder to copy to.   2) Optionally provide a new name.   3) Optionally provide a new description.   4) Optionally set to true if you want to copy and preserve the locked status. Requires &#x60;LockMonitors&#x60; capability. | 

### Return type

[**MonitorsLibraryBaseResponse**](MonitorsLibraryBaseResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitorsCreate

> MonitorsLibraryBaseResponse MonitorsCreate(ctx).ParentId(parentId).MonitorsLibraryBase(monitorsLibraryBase).Execute()

Create a monitor or folder. 



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
    parentId := "parentId_example" // string | Identifier of the parent folder in which to create the monitor or folder.
    monitorsLibraryBase := *openapiclient.NewMonitorsLibraryBase("Name_example", "Type_example") // MonitorsLibraryBase | The monitor or folder to create.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorsLibraryManagementAPI.MonitorsCreate(context.Background()).ParentId(parentId).MonitorsLibraryBase(monitorsLibraryBase).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorsLibraryManagementAPI.MonitorsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MonitorsCreate`: MonitorsLibraryBaseResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitorsLibraryManagementAPI.MonitorsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMonitorsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **parentId** | **string** | Identifier of the parent folder in which to create the monitor or folder. | 
 **monitorsLibraryBase** | [**MonitorsLibraryBase**](MonitorsLibraryBase.md) | The monitor or folder to create. | 

### Return type

[**MonitorsLibraryBaseResponse**](MonitorsLibraryBaseResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitorsDeleteById

> MonitorsDeleteById(ctx, id).Execute()

Delete a monitor or folder. 



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
    id := "id_example" // string | Identifier of the monitor or folder to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MonitorsLibraryManagementAPI.MonitorsDeleteById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorsLibraryManagementAPI.MonitorsDeleteById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the monitor or folder to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMonitorsDeleteByIdRequest struct via the builder pattern


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


## MonitorsDeleteByIds

> map[string]MonitorsLibraryBaseResponse MonitorsDeleteByIds(ctx).Ids(ids).Execute()

Bulk delete a monitor or folder. 



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
    ids := []string{"Inner_example"} // []string | A comma-separated list of identifiers.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorsLibraryManagementAPI.MonitorsDeleteByIds(context.Background()).Ids(ids).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorsLibraryManagementAPI.MonitorsDeleteByIds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MonitorsDeleteByIds`: map[string]MonitorsLibraryBaseResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitorsLibraryManagementAPI.MonitorsDeleteByIds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMonitorsDeleteByIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **[]string** | A comma-separated list of identifiers. | 

### Return type

[**map[string]MonitorsLibraryBaseResponse**](MonitorsLibraryBaseResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitorsExportItem

> MonitorsLibraryBaseExport MonitorsExportItem(ctx, id).Execute()

Export a monitor or folder.



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
    id := "id_example" // string | Identifier of the monitor or folder to export.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorsLibraryManagementAPI.MonitorsExportItem(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorsLibraryManagementAPI.MonitorsExportItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MonitorsExportItem`: MonitorsLibraryBaseExport
    fmt.Fprintf(os.Stdout, "Response from `MonitorsLibraryManagementAPI.MonitorsExportItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the monitor or folder to export. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMonitorsExportItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MonitorsLibraryBaseExport**](MonitorsLibraryBaseExport.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitorsGetByPath

> MonitorsLibraryBaseResponse MonitorsGetByPath(ctx).Path(path).Execute()

Read a monitor or folder by its path.



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
    path := "path_example" // string | The path of the monitor or folder.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorsLibraryManagementAPI.MonitorsGetByPath(context.Background()).Path(path).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorsLibraryManagementAPI.MonitorsGetByPath``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MonitorsGetByPath`: MonitorsLibraryBaseResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitorsLibraryManagementAPI.MonitorsGetByPath`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMonitorsGetByPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | The path of the monitor or folder. | 

### Return type

[**MonitorsLibraryBaseResponse**](MonitorsLibraryBaseResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitorsImportItem

> MonitorsLibraryBaseResponse MonitorsImportItem(ctx, parentId).MonitorsLibraryBaseExport(monitorsLibraryBaseExport).Execute()

Import a monitor or folder.



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
    parentId := "parentId_example" // string | Identifier of the parent folder in which to import the monitor or folder.
    monitorsLibraryBaseExport := *openapiclient.NewMonitorsLibraryBaseExport("Name_example", "Type_example") // MonitorsLibraryBaseExport | The monitor or folder to be imported.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorsLibraryManagementAPI.MonitorsImportItem(context.Background(), parentId).MonitorsLibraryBaseExport(monitorsLibraryBaseExport).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorsLibraryManagementAPI.MonitorsImportItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MonitorsImportItem`: MonitorsLibraryBaseResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitorsLibraryManagementAPI.MonitorsImportItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **string** | Identifier of the parent folder in which to import the monitor or folder. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMonitorsImportItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **monitorsLibraryBaseExport** | [**MonitorsLibraryBaseExport**](MonitorsLibraryBaseExport.md) | The monitor or folder to be imported. | 

### Return type

[**MonitorsLibraryBaseResponse**](MonitorsLibraryBaseResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitorsMove

> MonitorsLibraryBaseResponse MonitorsMove(ctx, id).ParentId(parentId).Execute()

Move a monitor or folder.



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
    id := "id_example" // string | Identifier of the monitor or folder to move.
    parentId := "parentId_example" // string | Identifier of the parent folder to move the monitor or folder to.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorsLibraryManagementAPI.MonitorsMove(context.Background(), id).ParentId(parentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorsLibraryManagementAPI.MonitorsMove``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MonitorsMove`: MonitorsLibraryBaseResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitorsLibraryManagementAPI.MonitorsMove`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the monitor or folder to move. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMonitorsMoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **parentId** | **string** | Identifier of the parent folder to move the monitor or folder to. | 

### Return type

[**MonitorsLibraryBaseResponse**](MonitorsLibraryBaseResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitorsReadById

> MonitorsLibraryBaseResponse MonitorsReadById(ctx, id).Execute()

Get a monitor or folder.



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
    id := "id_example" // string | Identifier of the monitor or folder to read.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorsLibraryManagementAPI.MonitorsReadById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorsLibraryManagementAPI.MonitorsReadById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MonitorsReadById`: MonitorsLibraryBaseResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitorsLibraryManagementAPI.MonitorsReadById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the monitor or folder to read. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMonitorsReadByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MonitorsLibraryBaseResponse**](MonitorsLibraryBaseResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitorsReadByIds

> map[string]MonitorsLibraryBaseResponse MonitorsReadByIds(ctx).Ids(ids).SkipChildren(skipChildren).Execute()

Bulk read a monitor or folder.



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
    ids := []string{"Inner_example"} // []string | A comma-separated list of identifiers.
    skipChildren := true // bool | a boolean parameter to control skipping fetching children of requested folder(s) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorsLibraryManagementAPI.MonitorsReadByIds(context.Background()).Ids(ids).SkipChildren(skipChildren).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorsLibraryManagementAPI.MonitorsReadByIds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MonitorsReadByIds`: map[string]MonitorsLibraryBaseResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitorsLibraryManagementAPI.MonitorsReadByIds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMonitorsReadByIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **[]string** | A comma-separated list of identifiers. | 
 **skipChildren** | **bool** | a boolean parameter to control skipping fetching children of requested folder(s) | 

### Return type

[**map[string]MonitorsLibraryBaseResponse**](MonitorsLibraryBaseResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitorsReadPermissionSummariesByIdGroupBySubjects

> PermissionSummariesBySubjects MonitorsReadPermissionSummariesByIdGroupBySubjects(ctx, id).Execute()

List permission summaries for a monitor or folder. 



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
    id := "id_example" // string | Identifier of the monitor or folder to list permissions.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorsLibraryManagementAPI.MonitorsReadPermissionSummariesByIdGroupBySubjects(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorsLibraryManagementAPI.MonitorsReadPermissionSummariesByIdGroupBySubjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MonitorsReadPermissionSummariesByIdGroupBySubjects`: PermissionSummariesBySubjects
    fmt.Fprintf(os.Stdout, "Response from `MonitorsLibraryManagementAPI.MonitorsReadPermissionSummariesByIdGroupBySubjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the monitor or folder to list permissions. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMonitorsReadPermissionSummariesByIdGroupBySubjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PermissionSummariesBySubjects**](PermissionSummariesBySubjects.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitorsReadPermissionsById

> ListPermissionsResponse MonitorsReadPermissionsById(ctx, id).Execute()

List explicit permissions on monitor or folder. 



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
    id := "id_example" // string | Identifier of the monitor or folder to list permissions.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorsLibraryManagementAPI.MonitorsReadPermissionsById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorsLibraryManagementAPI.MonitorsReadPermissionsById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MonitorsReadPermissionsById`: ListPermissionsResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitorsLibraryManagementAPI.MonitorsReadPermissionsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the monitor or folder to list permissions. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMonitorsReadPermissionsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListPermissionsResponse**](ListPermissionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitorsRevokePermissions

> MonitorsRevokePermissions(ctx).PermissionIdentifiers(permissionIdentifiers).Execute()

Revoke all permissions on monitor or folder. 



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
    permissionIdentifiers := *openapiclient.NewPermissionIdentifiers([]openapiclient.PermissionIdentifier{*openapiclient.NewPermissionIdentifier("role", "0000000006743FDA", "0000000006743FE2")}) // PermissionIdentifiers | The identifiers of the permissions statements to revoke.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MonitorsLibraryManagementAPI.MonitorsRevokePermissions(context.Background()).PermissionIdentifiers(permissionIdentifiers).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorsLibraryManagementAPI.MonitorsRevokePermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMonitorsRevokePermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **permissionIdentifiers** | [**PermissionIdentifiers**](PermissionIdentifiers.md) | The identifiers of the permissions statements to revoke. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitorsSearch

> []MonitorsLibraryItemWithPath MonitorsSearch(ctx).Query(query).Limit(limit).Offset(offset).SkipChildren(skipChildren).Execute()

Search for a monitor or folder.



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
    query := "createdBy:000000000000968B Test" // string | The search query to find monitor or folder. Below is the list of different filters with examples:   - **createdBy** : Filter by the user's identifier who created the content. Example: `createdBy:000000000000968B`.   - **createdBefore** : Filter by the content objects created before the given timestamp(in milliseconds). Example: `createdBefore:1457997222`.   - **createdAfter** : Filter by the content objects created after the given timestamp(in milliseconds). Example: `createdAfter:1457997111`.   - **modifiedBefore** : Filter by the content objects modified before the given timestamp(in milliseconds). Example: `modifiedBefore:1457997222`.   - **modifiedAfter** : Filter by the content objects modified after the given timestamp(in milliseconds). Example: `modifiedAfter:1457997111`.   - **type** : Filter by the type of the content object. Example: `type:folder`.   - **monitorStatus** : Filter by the status of the monitor: Normal, Critical, Warning, MissingData, Disabled, AllTriggered. Example: `monitorStatus:Normal`.  You can also use multiple filters in one query. For example to search for all content objects created by user with identifier 000000000000968B with creation timestamp after 1457997222 containing the text Test, the query would look like:    `createdBy:000000000000968B createdAfter:1457997222 Test`
    limit := int32(10) // int32 | Maximum number of items you want in the response. (optional) (default to 1000)
    offset := int32(5) // int32 | The position or row from where to start the search operation. (optional) (default to 0)
    skipChildren := true // bool | a boolean parameter to control skipping fetching children of requested folder(s) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorsLibraryManagementAPI.MonitorsSearch(context.Background()).Query(query).Limit(limit).Offset(offset).SkipChildren(skipChildren).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorsLibraryManagementAPI.MonitorsSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MonitorsSearch`: []MonitorsLibraryItemWithPath
    fmt.Fprintf(os.Stdout, "Response from `MonitorsLibraryManagementAPI.MonitorsSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMonitorsSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | The search query to find monitor or folder. Below is the list of different filters with examples:   - **createdBy** : Filter by the user&#39;s identifier who created the content. Example: &#x60;createdBy:000000000000968B&#x60;.   - **createdBefore** : Filter by the content objects created before the given timestamp(in milliseconds). Example: &#x60;createdBefore:1457997222&#x60;.   - **createdAfter** : Filter by the content objects created after the given timestamp(in milliseconds). Example: &#x60;createdAfter:1457997111&#x60;.   - **modifiedBefore** : Filter by the content objects modified before the given timestamp(in milliseconds). Example: &#x60;modifiedBefore:1457997222&#x60;.   - **modifiedAfter** : Filter by the content objects modified after the given timestamp(in milliseconds). Example: &#x60;modifiedAfter:1457997111&#x60;.   - **type** : Filter by the type of the content object. Example: &#x60;type:folder&#x60;.   - **monitorStatus** : Filter by the status of the monitor: Normal, Critical, Warning, MissingData, Disabled, AllTriggered. Example: &#x60;monitorStatus:Normal&#x60;.  You can also use multiple filters in one query. For example to search for all content objects created by user with identifier 000000000000968B with creation timestamp after 1457997222 containing the text Test, the query would look like:    &#x60;createdBy:000000000000968B createdAfter:1457997222 Test&#x60; | 
 **limit** | **int32** | Maximum number of items you want in the response. | [default to 1000]
 **offset** | **int32** | The position or row from where to start the search operation. | [default to 0]
 **skipChildren** | **bool** | a boolean parameter to control skipping fetching children of requested folder(s) | 

### Return type

[**[]MonitorsLibraryItemWithPath**](MonitorsLibraryItemWithPath.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitorsSetPermissions

> PermissionStatements MonitorsSetPermissions(ctx).PermissionStatementDefinitions(permissionStatementDefinitions).Execute()

Set permissions on monitor or folder. 



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
    permissionStatementDefinitions := *openapiclient.NewPermissionStatementDefinitions([]openapiclient.PermissionStatementDefinition{*openapiclient.NewPermissionStatementDefinition([]string{"Permissions_example"}, "role", "0000000006743FDA", "0000000006743FE2")}) // PermissionStatementDefinitions | The permission statement definitions to set.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorsLibraryManagementAPI.MonitorsSetPermissions(context.Background()).PermissionStatementDefinitions(permissionStatementDefinitions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorsLibraryManagementAPI.MonitorsSetPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MonitorsSetPermissions`: PermissionStatements
    fmt.Fprintf(os.Stdout, "Response from `MonitorsLibraryManagementAPI.MonitorsSetPermissions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMonitorsSetPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **permissionStatementDefinitions** | [**PermissionStatementDefinitions**](PermissionStatementDefinitions.md) | The permission statement definitions to set. | 

### Return type

[**PermissionStatements**](PermissionStatements.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitorsUpdateById

> MonitorsLibraryBaseResponse MonitorsUpdateById(ctx, id).MonitorsLibraryBaseUpdate(monitorsLibraryBaseUpdate).Execute()

Update a monitor or folder. 



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
    id := "id_example" // string | Identifier of the monitor or folder to update.
    monitorsLibraryBaseUpdate := *openapiclient.NewMonitorsLibraryBaseUpdate("Name_example", int64(123), "Type_example") // MonitorsLibraryBaseUpdate | The monitor or folder to update. The content version must match its latest version number in the monitors library. If the version does not match it will not be updated.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitorsLibraryManagementAPI.MonitorsUpdateById(context.Background(), id).MonitorsLibraryBaseUpdate(monitorsLibraryBaseUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorsLibraryManagementAPI.MonitorsUpdateById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MonitorsUpdateById`: MonitorsLibraryBaseResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitorsLibraryManagementAPI.MonitorsUpdateById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the monitor or folder to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMonitorsUpdateByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **monitorsLibraryBaseUpdate** | [**MonitorsLibraryBaseUpdate**](MonitorsLibraryBaseUpdate.md) | The monitor or folder to update. The content version must match its latest version number in the monitors library. If the version does not match it will not be updated. | 

### Return type

[**MonitorsLibraryBaseResponse**](MonitorsLibraryBaseResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

