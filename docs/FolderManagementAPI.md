# \FolderManagementAPI

All URIs are relative to *https://api.au.sumologic.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFolder**](FolderManagementAPI.md#CreateFolder) | **Post** /v2/content/folders | Create a new folder.
[**GetAdminRecommendedFolderAsync**](FolderManagementAPI.md#GetAdminRecommendedFolderAsync) | **Get** /v2/content/folders/adminRecommended | Schedule Admin Recommended folder job
[**GetAdminRecommendedFolderAsyncResult**](FolderManagementAPI.md#GetAdminRecommendedFolderAsyncResult) | **Get** /v2/content/folders/adminRecommended/{jobId}/result | Get Admin Recommended folder job result
[**GetAdminRecommendedFolderAsyncStatus**](FolderManagementAPI.md#GetAdminRecommendedFolderAsyncStatus) | **Get** /v2/content/folders/adminRecommended/{jobId}/status | Get Admin Recommended folder job status
[**GetFolder**](FolderManagementAPI.md#GetFolder) | **Get** /v2/content/folders/{id} | Get a folder.
[**GetGlobalFolderAsync**](FolderManagementAPI.md#GetGlobalFolderAsync) | **Get** /v2/content/folders/global | Schedule Global View job
[**GetGlobalFolderAsyncResult**](FolderManagementAPI.md#GetGlobalFolderAsyncResult) | **Get** /v2/content/folders/global/{jobId}/result | Get Global View job result
[**GetGlobalFolderAsyncStatus**](FolderManagementAPI.md#GetGlobalFolderAsyncStatus) | **Get** /v2/content/folders/global/{jobId}/status | Get Global View job status
[**GetPersonalFolder**](FolderManagementAPI.md#GetPersonalFolder) | **Get** /v2/content/folders/personal | Get personal folder.
[**UpdateFolder**](FolderManagementAPI.md#UpdateFolder) | **Put** /v2/content/folders/{id} | Update a folder.



## CreateFolder

> Folder CreateFolder(ctx).FolderDefinition(folderDefinition).IsAdminMode(isAdminMode).Execute()

Create a new folder.



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
    folderDefinition := *openapiclient.NewFolderDefinition("SampleFolder", "ParentId_example") // FolderDefinition | Information about the new folder.
    isAdminMode := "isAdminMode_example" // string | Set this to \"true\" if you want to perform the request as a Content Administrator. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FolderManagementAPI.CreateFolder(context.Background()).FolderDefinition(folderDefinition).IsAdminMode(isAdminMode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FolderManagementAPI.CreateFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFolder`: Folder
    fmt.Fprintf(os.Stdout, "Response from `FolderManagementAPI.CreateFolder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **folderDefinition** | [**FolderDefinition**](FolderDefinition.md) | Information about the new folder. | 
 **isAdminMode** | **string** | Set this to \&quot;true\&quot; if you want to perform the request as a Content Administrator. | 

### Return type

[**Folder**](Folder.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdminRecommendedFolderAsync

> BeginAsyncJobResponse GetAdminRecommendedFolderAsync(ctx).IsAdminMode(isAdminMode).Execute()

Schedule Admin Recommended folder job



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
    isAdminMode := "isAdminMode_example" // string | Set this to \"true\" if you want to perform the request as a Content Administrator. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FolderManagementAPI.GetAdminRecommendedFolderAsync(context.Background()).IsAdminMode(isAdminMode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FolderManagementAPI.GetAdminRecommendedFolderAsync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAdminRecommendedFolderAsync`: BeginAsyncJobResponse
    fmt.Fprintf(os.Stdout, "Response from `FolderManagementAPI.GetAdminRecommendedFolderAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAdminRecommendedFolderAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isAdminMode** | **string** | Set this to \&quot;true\&quot; if you want to perform the request as a Content Administrator. | 

### Return type

[**BeginAsyncJobResponse**](BeginAsyncJobResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdminRecommendedFolderAsyncResult

> Folder GetAdminRecommendedFolderAsyncResult(ctx, jobId).Execute()

Get Admin Recommended folder job result



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
    jobId := "jobId_example" // string | The identifier of the asynchronous Admin Recommended folder job.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FolderManagementAPI.GetAdminRecommendedFolderAsyncResult(context.Background(), jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FolderManagementAPI.GetAdminRecommendedFolderAsyncResult``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAdminRecommendedFolderAsyncResult`: Folder
    fmt.Fprintf(os.Stdout, "Response from `FolderManagementAPI.GetAdminRecommendedFolderAsyncResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | The identifier of the asynchronous Admin Recommended folder job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdminRecommendedFolderAsyncResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Folder**](Folder.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdminRecommendedFolderAsyncStatus

> AsyncJobStatus GetAdminRecommendedFolderAsyncStatus(ctx, jobId).Execute()

Get Admin Recommended folder job status



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
    jobId := "jobId_example" // string | The identifier of the asynchronous Admin Recommended folder job.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FolderManagementAPI.GetAdminRecommendedFolderAsyncStatus(context.Background(), jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FolderManagementAPI.GetAdminRecommendedFolderAsyncStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAdminRecommendedFolderAsyncStatus`: AsyncJobStatus
    fmt.Fprintf(os.Stdout, "Response from `FolderManagementAPI.GetAdminRecommendedFolderAsyncStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | The identifier of the asynchronous Admin Recommended folder job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdminRecommendedFolderAsyncStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AsyncJobStatus**](AsyncJobStatus.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFolder

> Folder GetFolder(ctx, id).IsAdminMode(isAdminMode).Execute()

Get a folder.



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
    id := "id_example" // string | Identifier of the folder to fetch.
    isAdminMode := "isAdminMode_example" // string | Set this to \"true\" if you want to perform the request as a Content Administrator. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FolderManagementAPI.GetFolder(context.Background(), id).IsAdminMode(isAdminMode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FolderManagementAPI.GetFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFolder`: Folder
    fmt.Fprintf(os.Stdout, "Response from `FolderManagementAPI.GetFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the folder to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isAdminMode** | **string** | Set this to \&quot;true\&quot; if you want to perform the request as a Content Administrator. | 

### Return type

[**Folder**](Folder.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGlobalFolderAsync

> BeginAsyncJobResponse GetGlobalFolderAsync(ctx).IsAdminMode(isAdminMode).Execute()

Schedule Global View job



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
    isAdminMode := "isAdminMode_example" // string | Set this to \"true\" if you want to perform the request as a Content Administrator. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FolderManagementAPI.GetGlobalFolderAsync(context.Background()).IsAdminMode(isAdminMode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FolderManagementAPI.GetGlobalFolderAsync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGlobalFolderAsync`: BeginAsyncJobResponse
    fmt.Fprintf(os.Stdout, "Response from `FolderManagementAPI.GetGlobalFolderAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGlobalFolderAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isAdminMode** | **string** | Set this to \&quot;true\&quot; if you want to perform the request as a Content Administrator. | 

### Return type

[**BeginAsyncJobResponse**](BeginAsyncJobResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGlobalFolderAsyncResult

> ContentList GetGlobalFolderAsyncResult(ctx, jobId).Execute()

Get Global View job result



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
    jobId := "jobId_example" // string | The identifier of the asynchronous Global View job.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FolderManagementAPI.GetGlobalFolderAsyncResult(context.Background(), jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FolderManagementAPI.GetGlobalFolderAsyncResult``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGlobalFolderAsyncResult`: ContentList
    fmt.Fprintf(os.Stdout, "Response from `FolderManagementAPI.GetGlobalFolderAsyncResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | The identifier of the asynchronous Global View job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGlobalFolderAsyncResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ContentList**](ContentList.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGlobalFolderAsyncStatus

> AsyncJobStatus GetGlobalFolderAsyncStatus(ctx, jobId).Execute()

Get Global View job status



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
    jobId := "jobId_example" // string | The identifier of the asynchronous Global View job.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FolderManagementAPI.GetGlobalFolderAsyncStatus(context.Background(), jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FolderManagementAPI.GetGlobalFolderAsyncStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGlobalFolderAsyncStatus`: AsyncJobStatus
    fmt.Fprintf(os.Stdout, "Response from `FolderManagementAPI.GetGlobalFolderAsyncStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | The identifier of the asynchronous Global View job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGlobalFolderAsyncStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AsyncJobStatus**](AsyncJobStatus.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPersonalFolder

> Folder GetPersonalFolder(ctx).Execute()

Get personal folder.



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
    resp, r, err := apiClient.FolderManagementAPI.GetPersonalFolder(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FolderManagementAPI.GetPersonalFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPersonalFolder`: Folder
    fmt.Fprintf(os.Stdout, "Response from `FolderManagementAPI.GetPersonalFolder`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPersonalFolderRequest struct via the builder pattern


### Return type

[**Folder**](Folder.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFolder

> Folder UpdateFolder(ctx, id).UpdateFolderRequest(updateFolderRequest).IsAdminMode(isAdminMode).Execute()

Update a folder.



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
    id := "id_example" // string | Identifier of the folder to update.
    updateFolderRequest := *openapiclient.NewUpdateFolderRequest("SampleFolder") // UpdateFolderRequest | Information to update about the folder.
    isAdminMode := "isAdminMode_example" // string | Set this to \"true\" if you want to perform the request as a Content Administrator. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FolderManagementAPI.UpdateFolder(context.Background(), id).UpdateFolderRequest(updateFolderRequest).IsAdminMode(isAdminMode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FolderManagementAPI.UpdateFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFolder`: Folder
    fmt.Fprintf(os.Stdout, "Response from `FolderManagementAPI.UpdateFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the folder to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateFolderRequest** | [**UpdateFolderRequest**](UpdateFolderRequest.md) | Information to update about the folder. | 
 **isAdminMode** | **string** | Set this to \&quot;true\&quot; if you want to perform the request as a Content Administrator. | 

### Return type

[**Folder**](Folder.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

