# \MutingSchedulesLibraryManagementAPI

All URIs are relative to *https://api.au.sumologic.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMutingSchedulesFullPath**](MutingSchedulesLibraryManagementAPI.md#GetMutingSchedulesFullPath) | **Get** /v1/mutingSchedules/{id}/path | Get the path of a mutingschedule or folder.
[**GetMutingSchedulesLibraryRoot**](MutingSchedulesLibraryManagementAPI.md#GetMutingSchedulesLibraryRoot) | **Get** /v1/mutingSchedules/root | Get the root mutingSchedules folder.
[**MutingSchedulesCopy**](MutingSchedulesLibraryManagementAPI.md#MutingSchedulesCopy) | **Post** /v1/mutingSchedules/{id}/copy | Copy a mutingschedule or folder.
[**MutingSchedulesCreate**](MutingSchedulesLibraryManagementAPI.md#MutingSchedulesCreate) | **Post** /v1/mutingSchedules | Create a mutingschedule or folder. 
[**MutingSchedulesDeleteById**](MutingSchedulesLibraryManagementAPI.md#MutingSchedulesDeleteById) | **Delete** /v1/mutingSchedules/{id} | Delete a mutingschedule or folder. 
[**MutingSchedulesDeleteByIds**](MutingSchedulesLibraryManagementAPI.md#MutingSchedulesDeleteByIds) | **Delete** /v1/mutingSchedules | Bulk delete a mutingschedule or folder. 
[**MutingSchedulesExportItem**](MutingSchedulesLibraryManagementAPI.md#MutingSchedulesExportItem) | **Get** /v1/mutingSchedules/{id}/export | Export a mutingschedule or folder.
[**MutingSchedulesImportItem**](MutingSchedulesLibraryManagementAPI.md#MutingSchedulesImportItem) | **Post** /v1/mutingSchedules/{parentId}/import | Import a mutingschedule or folder.
[**MutingSchedulesReadById**](MutingSchedulesLibraryManagementAPI.md#MutingSchedulesReadById) | **Get** /v1/mutingSchedules/{id} | Get a mutingschedule or folder.
[**MutingSchedulesReadByIds**](MutingSchedulesLibraryManagementAPI.md#MutingSchedulesReadByIds) | **Get** /v1/mutingSchedules | Bulk read a mutingschedule or folder.
[**MutingSchedulesSearch**](MutingSchedulesLibraryManagementAPI.md#MutingSchedulesSearch) | **Get** /v1/mutingSchedules/search | Search for a mutingschedule or folder.
[**MutingSchedulesUpdateById**](MutingSchedulesLibraryManagementAPI.md#MutingSchedulesUpdateById) | **Put** /v1/mutingSchedules/{id} | Update a mutingschedule or folder. 



## GetMutingSchedulesFullPath

> Path GetMutingSchedulesFullPath(ctx, id).Execute()

Get the path of a mutingschedule or folder.



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
    id := "id_example" // string | Identifier of the mutingschedule or folder.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MutingSchedulesLibraryManagementAPI.GetMutingSchedulesFullPath(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MutingSchedulesLibraryManagementAPI.GetMutingSchedulesFullPath``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMutingSchedulesFullPath`: Path
    fmt.Fprintf(os.Stdout, "Response from `MutingSchedulesLibraryManagementAPI.GetMutingSchedulesFullPath`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the mutingschedule or folder. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMutingSchedulesFullPathRequest struct via the builder pattern


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


## GetMutingSchedulesLibraryRoot

> MutingSchedulesLibraryFolderResponse GetMutingSchedulesLibraryRoot(ctx).Execute()

Get the root mutingSchedules folder.



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
    resp, r, err := apiClient.MutingSchedulesLibraryManagementAPI.GetMutingSchedulesLibraryRoot(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MutingSchedulesLibraryManagementAPI.GetMutingSchedulesLibraryRoot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMutingSchedulesLibraryRoot`: MutingSchedulesLibraryFolderResponse
    fmt.Fprintf(os.Stdout, "Response from `MutingSchedulesLibraryManagementAPI.GetMutingSchedulesLibraryRoot`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMutingSchedulesLibraryRootRequest struct via the builder pattern


### Return type

[**MutingSchedulesLibraryFolderResponse**](MutingSchedulesLibraryFolderResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MutingSchedulesCopy

> MutingSchedulesLibraryBaseResponse MutingSchedulesCopy(ctx, id).ContentCopyParams(contentCopyParams).Execute()

Copy a mutingschedule or folder.



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
    id := "id_example" // string | Identifier of the mutingschedule or folder to copy.
    contentCopyParams := *openapiclient.NewContentCopyParams("ParentId_example") // ContentCopyParams | Fields include:   1) Identifier of the parent folder to copy to.   2) Optionally provide a new name.   3) Optionally provide a new description.   4) Optionally set to true if you want to copy and preserve the locked status. Requires `LockMutingSchedules` capability.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MutingSchedulesLibraryManagementAPI.MutingSchedulesCopy(context.Background(), id).ContentCopyParams(contentCopyParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MutingSchedulesLibraryManagementAPI.MutingSchedulesCopy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MutingSchedulesCopy`: MutingSchedulesLibraryBaseResponse
    fmt.Fprintf(os.Stdout, "Response from `MutingSchedulesLibraryManagementAPI.MutingSchedulesCopy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the mutingschedule or folder to copy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMutingSchedulesCopyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentCopyParams** | [**ContentCopyParams**](ContentCopyParams.md) | Fields include:   1) Identifier of the parent folder to copy to.   2) Optionally provide a new name.   3) Optionally provide a new description.   4) Optionally set to true if you want to copy and preserve the locked status. Requires &#x60;LockMutingSchedules&#x60; capability. | 

### Return type

[**MutingSchedulesLibraryBaseResponse**](MutingSchedulesLibraryBaseResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MutingSchedulesCreate

> MutingSchedulesLibraryBaseResponse MutingSchedulesCreate(ctx).ParentId(parentId).MutingSchedulesLibraryBase(mutingSchedulesLibraryBase).Execute()

Create a mutingschedule or folder. 



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
    parentId := "parentId_example" // string | Identifier of the parent folder in which to create the mutingschedule or folder.
    mutingSchedulesLibraryBase := *openapiclient.NewMutingSchedulesLibraryBase("Name_example", "Type_example") // MutingSchedulesLibraryBase | The mutingschedule or folder to create.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MutingSchedulesLibraryManagementAPI.MutingSchedulesCreate(context.Background()).ParentId(parentId).MutingSchedulesLibraryBase(mutingSchedulesLibraryBase).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MutingSchedulesLibraryManagementAPI.MutingSchedulesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MutingSchedulesCreate`: MutingSchedulesLibraryBaseResponse
    fmt.Fprintf(os.Stdout, "Response from `MutingSchedulesLibraryManagementAPI.MutingSchedulesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMutingSchedulesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **parentId** | **string** | Identifier of the parent folder in which to create the mutingschedule or folder. | 
 **mutingSchedulesLibraryBase** | [**MutingSchedulesLibraryBase**](MutingSchedulesLibraryBase.md) | The mutingschedule or folder to create. | 

### Return type

[**MutingSchedulesLibraryBaseResponse**](MutingSchedulesLibraryBaseResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MutingSchedulesDeleteById

> MutingSchedulesDeleteById(ctx, id).Execute()

Delete a mutingschedule or folder. 



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
    id := "id_example" // string | Identifier of the mutingschedule or folder to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MutingSchedulesLibraryManagementAPI.MutingSchedulesDeleteById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MutingSchedulesLibraryManagementAPI.MutingSchedulesDeleteById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the mutingschedule or folder to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMutingSchedulesDeleteByIdRequest struct via the builder pattern


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


## MutingSchedulesDeleteByIds

> map[string]MutingSchedulesLibraryBaseResponse MutingSchedulesDeleteByIds(ctx).Ids(ids).Execute()

Bulk delete a mutingschedule or folder. 



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
    resp, r, err := apiClient.MutingSchedulesLibraryManagementAPI.MutingSchedulesDeleteByIds(context.Background()).Ids(ids).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MutingSchedulesLibraryManagementAPI.MutingSchedulesDeleteByIds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MutingSchedulesDeleteByIds`: map[string]MutingSchedulesLibraryBaseResponse
    fmt.Fprintf(os.Stdout, "Response from `MutingSchedulesLibraryManagementAPI.MutingSchedulesDeleteByIds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMutingSchedulesDeleteByIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **[]string** | A comma-separated list of identifiers. | 

### Return type

[**map[string]MutingSchedulesLibraryBaseResponse**](MutingSchedulesLibraryBaseResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MutingSchedulesExportItem

> MutingSchedulesLibraryBaseExport MutingSchedulesExportItem(ctx, id).Execute()

Export a mutingschedule or folder.



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
    id := "id_example" // string | Identifier of the mutingschedule or folder to export.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MutingSchedulesLibraryManagementAPI.MutingSchedulesExportItem(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MutingSchedulesLibraryManagementAPI.MutingSchedulesExportItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MutingSchedulesExportItem`: MutingSchedulesLibraryBaseExport
    fmt.Fprintf(os.Stdout, "Response from `MutingSchedulesLibraryManagementAPI.MutingSchedulesExportItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the mutingschedule or folder to export. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMutingSchedulesExportItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MutingSchedulesLibraryBaseExport**](MutingSchedulesLibraryBaseExport.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MutingSchedulesImportItem

> MutingSchedulesLibraryBaseResponse MutingSchedulesImportItem(ctx, parentId).MutingSchedulesLibraryBaseExport(mutingSchedulesLibraryBaseExport).Execute()

Import a mutingschedule or folder.



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
    parentId := "parentId_example" // string | Identifier of the parent folder in which to import the mutingschedule or folder.
    mutingSchedulesLibraryBaseExport := *openapiclient.NewMutingSchedulesLibraryBaseExport("Name_example", "Type_example") // MutingSchedulesLibraryBaseExport | The mutingschedule or folder to be imported.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MutingSchedulesLibraryManagementAPI.MutingSchedulesImportItem(context.Background(), parentId).MutingSchedulesLibraryBaseExport(mutingSchedulesLibraryBaseExport).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MutingSchedulesLibraryManagementAPI.MutingSchedulesImportItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MutingSchedulesImportItem`: MutingSchedulesLibraryBaseResponse
    fmt.Fprintf(os.Stdout, "Response from `MutingSchedulesLibraryManagementAPI.MutingSchedulesImportItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | **string** | Identifier of the parent folder in which to import the mutingschedule or folder. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMutingSchedulesImportItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mutingSchedulesLibraryBaseExport** | [**MutingSchedulesLibraryBaseExport**](MutingSchedulesLibraryBaseExport.md) | The mutingschedule or folder to be imported. | 

### Return type

[**MutingSchedulesLibraryBaseResponse**](MutingSchedulesLibraryBaseResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MutingSchedulesReadById

> MutingSchedulesLibraryBaseResponse MutingSchedulesReadById(ctx, id).Execute()

Get a mutingschedule or folder.



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
    id := "id_example" // string | Identifier of the mutingschedule or folder to read.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MutingSchedulesLibraryManagementAPI.MutingSchedulesReadById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MutingSchedulesLibraryManagementAPI.MutingSchedulesReadById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MutingSchedulesReadById`: MutingSchedulesLibraryBaseResponse
    fmt.Fprintf(os.Stdout, "Response from `MutingSchedulesLibraryManagementAPI.MutingSchedulesReadById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the mutingschedule or folder to read. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMutingSchedulesReadByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MutingSchedulesLibraryBaseResponse**](MutingSchedulesLibraryBaseResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MutingSchedulesReadByIds

> map[string]MutingSchedulesLibraryBaseResponse MutingSchedulesReadByIds(ctx).Ids(ids).SkipChildren(skipChildren).Execute()

Bulk read a mutingschedule or folder.



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
    resp, r, err := apiClient.MutingSchedulesLibraryManagementAPI.MutingSchedulesReadByIds(context.Background()).Ids(ids).SkipChildren(skipChildren).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MutingSchedulesLibraryManagementAPI.MutingSchedulesReadByIds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MutingSchedulesReadByIds`: map[string]MutingSchedulesLibraryBaseResponse
    fmt.Fprintf(os.Stdout, "Response from `MutingSchedulesLibraryManagementAPI.MutingSchedulesReadByIds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMutingSchedulesReadByIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **[]string** | A comma-separated list of identifiers. | 
 **skipChildren** | **bool** | a boolean parameter to control skipping fetching children of requested folder(s) | 

### Return type

[**map[string]MutingSchedulesLibraryBaseResponse**](MutingSchedulesLibraryBaseResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MutingSchedulesSearch

> []MutingSchedulesLibraryItemWithPath MutingSchedulesSearch(ctx).Query(query).Limit(limit).Offset(offset).SkipChildren(skipChildren).Execute()

Search for a mutingschedule or folder.



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
    query := "createdBy:000000000000968B Test" // string | The search query to find mutingschedule or folder. Below is the list of different filters with examples:   - **createdBy** : Filter by the user's identifier who created the content. Example: `createdBy:000000000000968B`.   - **createdBefore** : Filter by the content objects created before the given timestamp(in milliseconds). Example: `createdBefore:1457997222`.   - **createdAfter** : Filter by the content objects created after the given timestamp(in milliseconds). Example: `createdAfter:1457997111`.   - **modifiedBefore** : Filter by the content objects modified before the given timestamp(in milliseconds). Example: `modifiedBefore:1457997222`.   - **modifiedAfter** : Filter by the content objects modified after the given timestamp(in milliseconds). Example: `modifiedAfter:1457997111`.   - **type** : Filter by the type of the content object. Example: `type:folder`.  You can also use multiple filters in one query. For example to search for all content objects created by user with identifier 000000000000968B with creation timestamp after 1457997222 containing the text Test, the query would look like:    `createdBy:000000000000968B createdAfter:1457997222 Test`
    limit := int32(10) // int32 | Maximum number of items you want in the response. (optional) (default to 1000)
    offset := int32(5) // int32 | The position or row from where to start the search operation. (optional) (default to 0)
    skipChildren := true // bool | a boolean parameter to control skipping fetching children of requested folder(s) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MutingSchedulesLibraryManagementAPI.MutingSchedulesSearch(context.Background()).Query(query).Limit(limit).Offset(offset).SkipChildren(skipChildren).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MutingSchedulesLibraryManagementAPI.MutingSchedulesSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MutingSchedulesSearch`: []MutingSchedulesLibraryItemWithPath
    fmt.Fprintf(os.Stdout, "Response from `MutingSchedulesLibraryManagementAPI.MutingSchedulesSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMutingSchedulesSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | The search query to find mutingschedule or folder. Below is the list of different filters with examples:   - **createdBy** : Filter by the user&#39;s identifier who created the content. Example: &#x60;createdBy:000000000000968B&#x60;.   - **createdBefore** : Filter by the content objects created before the given timestamp(in milliseconds). Example: &#x60;createdBefore:1457997222&#x60;.   - **createdAfter** : Filter by the content objects created after the given timestamp(in milliseconds). Example: &#x60;createdAfter:1457997111&#x60;.   - **modifiedBefore** : Filter by the content objects modified before the given timestamp(in milliseconds). Example: &#x60;modifiedBefore:1457997222&#x60;.   - **modifiedAfter** : Filter by the content objects modified after the given timestamp(in milliseconds). Example: &#x60;modifiedAfter:1457997111&#x60;.   - **type** : Filter by the type of the content object. Example: &#x60;type:folder&#x60;.  You can also use multiple filters in one query. For example to search for all content objects created by user with identifier 000000000000968B with creation timestamp after 1457997222 containing the text Test, the query would look like:    &#x60;createdBy:000000000000968B createdAfter:1457997222 Test&#x60; | 
 **limit** | **int32** | Maximum number of items you want in the response. | [default to 1000]
 **offset** | **int32** | The position or row from where to start the search operation. | [default to 0]
 **skipChildren** | **bool** | a boolean parameter to control skipping fetching children of requested folder(s) | 

### Return type

[**[]MutingSchedulesLibraryItemWithPath**](MutingSchedulesLibraryItemWithPath.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MutingSchedulesUpdateById

> MutingSchedulesLibraryBaseResponse MutingSchedulesUpdateById(ctx, id).MutingSchedulesLibraryBaseUpdate(mutingSchedulesLibraryBaseUpdate).Execute()

Update a mutingschedule or folder. 



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
    id := "id_example" // string | Identifier of the mutingschedule or folder to update.
    mutingSchedulesLibraryBaseUpdate := *openapiclient.NewMutingSchedulesLibraryBaseUpdate("Name_example", int64(123), "Type_example") // MutingSchedulesLibraryBaseUpdate | The mutingschedule or folder to update. The content version must match its latest version number in the mutingSchedules library. If the version does not match it will not be updated.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MutingSchedulesLibraryManagementAPI.MutingSchedulesUpdateById(context.Background(), id).MutingSchedulesLibraryBaseUpdate(mutingSchedulesLibraryBaseUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MutingSchedulesLibraryManagementAPI.MutingSchedulesUpdateById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MutingSchedulesUpdateById`: MutingSchedulesLibraryBaseResponse
    fmt.Fprintf(os.Stdout, "Response from `MutingSchedulesLibraryManagementAPI.MutingSchedulesUpdateById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the mutingschedule or folder to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMutingSchedulesUpdateByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mutingSchedulesLibraryBaseUpdate** | [**MutingSchedulesLibraryBaseUpdate**](MutingSchedulesLibraryBaseUpdate.md) | The mutingschedule or folder to update. The content version must match its latest version number in the mutingSchedules library. If the version does not match it will not be updated. | 

### Return type

[**MutingSchedulesLibraryBaseResponse**](MutingSchedulesLibraryBaseResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

