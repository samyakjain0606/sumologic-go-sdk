# \AppManagementAPI

All URIs are relative to *https://api.au.sumologic.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApp**](AppManagementAPI.md#GetApp) | **Get** /v1/apps/{uuid} | Get an app by UUID.
[**GetAsyncInstallStatus**](AppManagementAPI.md#GetAsyncInstallStatus) | **Get** /v1/apps/install/{jobId}/status | App install job status.
[**InstallApp**](AppManagementAPI.md#InstallApp) | **Post** /v1/apps/{uuid}/install | Install an app by UUID.
[**ListApps**](AppManagementAPI.md#ListApps) | **Get** /v1/apps | List available apps.



## GetApp

> App GetApp(ctx, uuid).Execute()

Get an app by UUID.



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
    uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The identifier of the app to retrieve.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppManagementAPI.GetApp(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppManagementAPI.GetApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApp`: App
    fmt.Fprintf(os.Stdout, "Response from `AppManagementAPI.GetApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | The identifier of the app to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**App**](App.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAsyncInstallStatus

> AsyncJobStatus GetAsyncInstallStatus(ctx, jobId).Execute()

App install job status.



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
    jobId := "jobId_example" // string | The identifier of the asynchronous install job.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppManagementAPI.GetAsyncInstallStatus(context.Background(), jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppManagementAPI.GetAsyncInstallStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAsyncInstallStatus`: AsyncJobStatus
    fmt.Fprintf(os.Stdout, "Response from `AppManagementAPI.GetAsyncInstallStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | The identifier of the asynchronous install job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAsyncInstallStatusRequest struct via the builder pattern


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


## InstallApp

> BeginAsyncJobResponse InstallApp(ctx, uuid).AppInstallRequest(appInstallRequest).Execute()

Install an app by UUID.



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
    uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the app to install.
    appInstallRequest := *openapiclient.NewAppInstallRequest("Sumo Logic Configuration App", "Sumo Logic Configuration App to configure collectors and data sources", "00000000000001C8") // AppInstallRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppManagementAPI.InstallApp(context.Background(), uuid).AppInstallRequest(appInstallRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppManagementAPI.InstallApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstallApp`: BeginAsyncJobResponse
    fmt.Fprintf(os.Stdout, "Response from `AppManagementAPI.InstallApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | UUID of the app to install. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstallAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appInstallRequest** | [**AppInstallRequest**](AppInstallRequest.md) |  | 

### Return type

[**BeginAsyncJobResponse**](BeginAsyncJobResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApps

> ListAppsResult ListApps(ctx).Execute()

List available apps.



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
    resp, r, err := apiClient.AppManagementAPI.ListApps(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppManagementAPI.ListApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApps`: ListAppsResult
    fmt.Fprintf(os.Stdout, "Response from `AppManagementAPI.ListApps`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAppsRequest struct via the builder pattern


### Return type

[**ListAppsResult**](ListAppsResult.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

