# \DashboardManagementAPI

All URIs are relative to *https://api.au.sumologic.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDashboard**](DashboardManagementAPI.md#CreateDashboard) | **Post** /v2/dashboards | Create a new dashboard.
[**DeleteDashboard**](DashboardManagementAPI.md#DeleteDashboard) | **Delete** /v2/dashboards/{id} | Delete a dashboard.
[**GenerateDashboardReport**](DashboardManagementAPI.md#GenerateDashboardReport) | **Post** /v2/dashboards/reportJobs | Start a report job
[**GetAsyncReportGenerationResult**](DashboardManagementAPI.md#GetAsyncReportGenerationResult) | **Get** /v2/dashboards/reportJobs/{jobId}/result | Get report generation job result
[**GetAsyncReportGenerationStatus**](DashboardManagementAPI.md#GetAsyncReportGenerationStatus) | **Get** /v2/dashboards/reportJobs/{jobId}/status | Get report generation job status
[**GetDashboard**](DashboardManagementAPI.md#GetDashboard) | **Get** /v2/dashboards/{id} | Get a dashboard.
[**GetDashboardMigrationResult**](DashboardManagementAPI.md#GetDashboardMigrationResult) | **Get** /v2/dashboards/migrate/{jobId}/result | Get dashboard migration result.
[**GetDashboardMigrationStatus**](DashboardManagementAPI.md#GetDashboardMigrationStatus) | **Get** /v2/dashboards/migrate/{jobId}/status | Get dashboard migration status.
[**ListDashboards**](DashboardManagementAPI.md#ListDashboards) | **Get** /v2/dashboards | List all dashboards.
[**MigrateReportToDashboard**](DashboardManagementAPI.md#MigrateReportToDashboard) | **Post** /v2/dashboards/migrate | Migrate Legacy Dashboards to Dashboards(New)
[**PreviewMigrateReportToDashboard**](DashboardManagementAPI.md#PreviewMigrateReportToDashboard) | **Post** /v2/dashboards/migrate/preview | Preview of Migrating Legacy Dashboards to Dashboards(New)
[**UpdateDashboard**](DashboardManagementAPI.md#UpdateDashboard) | **Put** /v2/dashboards/{id} | Update a dashboard.



## CreateDashboard

> Dashboard CreateDashboard(ctx).DashboardRequest(dashboardRequest).Execute()

Create a new dashboard.



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
    dashboardRequest := *openapiclient.NewDashboardRequest("Kubernetes Dashboard", *openapiclient.NewResolvableTimeRange("Type_example")) // DashboardRequest | Information to create the new dashboard.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DashboardManagementAPI.CreateDashboard(context.Background()).DashboardRequest(dashboardRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardManagementAPI.CreateDashboard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDashboard`: Dashboard
    fmt.Fprintf(os.Stdout, "Response from `DashboardManagementAPI.CreateDashboard`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dashboardRequest** | [**DashboardRequest**](DashboardRequest.md) | Information to create the new dashboard. | 

### Return type

[**Dashboard**](Dashboard.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDashboard

> DeleteDashboard(ctx, id).Execute()

Delete a dashboard.



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
    id := "id_example" // string | Identifier of the dashboard to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DashboardManagementAPI.DeleteDashboard(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardManagementAPI.DeleteDashboard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the dashboard to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDashboardRequest struct via the builder pattern


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


## GenerateDashboardReport

> BeginAsyncJobResponse GenerateDashboardReport(ctx).GenerateReportRequest(generateReportRequest).Execute()

Start a report job



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
    generateReportRequest := *openapiclient.NewGenerateReportRequest(*openapiclient.NewReportAction("DirectDownloadReportAction"), "Pdf", "America/Los_Angeles", *openapiclient.NewTemplate("DashboardTemplate")) // GenerateReportRequest | Request for a report.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DashboardManagementAPI.GenerateDashboardReport(context.Background()).GenerateReportRequest(generateReportRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardManagementAPI.GenerateDashboardReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateDashboardReport`: BeginAsyncJobResponse
    fmt.Fprintf(os.Stdout, "Response from `DashboardManagementAPI.GenerateDashboardReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateDashboardReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **generateReportRequest** | [**GenerateReportRequest**](GenerateReportRequest.md) | Request for a report. | 

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


## GetAsyncReportGenerationResult

> *os.File GetAsyncReportGenerationResult(ctx, jobId).Execute()

Get report generation job result



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
    jobId := "jobId_example" // string | The identifier of the asynchronous report generation job.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DashboardManagementAPI.GetAsyncReportGenerationResult(context.Background(), jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardManagementAPI.GetAsyncReportGenerationResult``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAsyncReportGenerationResult`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `DashboardManagementAPI.GetAsyncReportGenerationResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | The identifier of the asynchronous report generation job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAsyncReportGenerationResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/pdf, image/png, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAsyncReportGenerationStatus

> AsyncJobStatus GetAsyncReportGenerationStatus(ctx, jobId).Execute()

Get report generation job status



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
    jobId := "jobId_example" // string | The identifier of the asynchronous report generation job.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DashboardManagementAPI.GetAsyncReportGenerationStatus(context.Background(), jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardManagementAPI.GetAsyncReportGenerationStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAsyncReportGenerationStatus`: AsyncJobStatus
    fmt.Fprintf(os.Stdout, "Response from `DashboardManagementAPI.GetAsyncReportGenerationStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | The identifier of the asynchronous report generation job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAsyncReportGenerationStatusRequest struct via the builder pattern


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


## GetDashboard

> Dashboard GetDashboard(ctx, id).Execute()

Get a dashboard.



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
    id := "id_example" // string | UUID of the dashboard to return.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DashboardManagementAPI.GetDashboard(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardManagementAPI.GetDashboard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDashboard`: Dashboard
    fmt.Fprintf(os.Stdout, "Response from `DashboardManagementAPI.GetDashboard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | UUID of the dashboard to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Dashboard**](Dashboard.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDashboardMigrationResult

> DashboardMigrationResult GetDashboardMigrationResult(ctx, jobId).Execute()

Get dashboard migration result.



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
    jobId := "jobId_example" // string | The identifier of the asynchronous Dashboard Migration job.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DashboardManagementAPI.GetDashboardMigrationResult(context.Background(), jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardManagementAPI.GetDashboardMigrationResult``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDashboardMigrationResult`: DashboardMigrationResult
    fmt.Fprintf(os.Stdout, "Response from `DashboardManagementAPI.GetDashboardMigrationResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | The identifier of the asynchronous Dashboard Migration job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDashboardMigrationResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DashboardMigrationResult**](DashboardMigrationResult.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDashboardMigrationStatus

> AsyncJobStatus GetDashboardMigrationStatus(ctx, jobId).Execute()

Get dashboard migration status.



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
    jobId := "jobId_example" // string | The identifier of the asynchronous Dashboard Migration job.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DashboardManagementAPI.GetDashboardMigrationStatus(context.Background(), jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardManagementAPI.GetDashboardMigrationStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDashboardMigrationStatus`: AsyncJobStatus
    fmt.Fprintf(os.Stdout, "Response from `DashboardManagementAPI.GetDashboardMigrationStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | The identifier of the asynchronous Dashboard Migration job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDashboardMigrationStatusRequest struct via the builder pattern


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


## ListDashboards

> PaginatedDashboards ListDashboards(ctx).Limit(limit).Token(token).Mode(mode).Execute()

List all dashboards.



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
    limit := int32(50) // int32 | Limit the number of dashboard returned in the response. The number of dashboards returned may be less than the `limit`. (optional) (default to 50)
    token := "GDCiRv4vebF3UWFJQ1kySXBOR3Bzh69GR0RyWm9vCtc" // string | Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. `token` is set to null when no more pages are left. (optional)
    mode := "createdByUser" // string | whether to list all viewable dashboards under the folders (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DashboardManagementAPI.ListDashboards(context.Background()).Limit(limit).Token(token).Mode(mode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardManagementAPI.ListDashboards``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDashboards`: PaginatedDashboards
    fmt.Fprintf(os.Stdout, "Response from `DashboardManagementAPI.ListDashboards`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDashboardsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit the number of dashboard returned in the response. The number of dashboards returned may be less than the &#x60;limit&#x60;. | [default to 50]
 **token** | **string** | Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. &#x60;token&#x60; is set to null when no more pages are left. | 
 **mode** | **string** | whether to list all viewable dashboards under the folders | 

### Return type

[**PaginatedDashboards**](PaginatedDashboards.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MigrateReportToDashboard

> BeginAsyncJobResponseV2 MigrateReportToDashboard(ctx).DashboardMigrationRequest(dashboardMigrationRequest).Execute()

Migrate Legacy Dashboards to Dashboards(New)



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
    dashboardMigrationRequest := *openapiclient.NewDashboardMigrationRequest([]string{"00000000000001C8"}) // DashboardMigrationRequest | List of legacy dashboard content identifiers.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DashboardManagementAPI.MigrateReportToDashboard(context.Background()).DashboardMigrationRequest(dashboardMigrationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardManagementAPI.MigrateReportToDashboard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MigrateReportToDashboard`: BeginAsyncJobResponseV2
    fmt.Fprintf(os.Stdout, "Response from `DashboardManagementAPI.MigrateReportToDashboard`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMigrateReportToDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dashboardMigrationRequest** | [**DashboardMigrationRequest**](DashboardMigrationRequest.md) | List of legacy dashboard content identifiers. | 

### Return type

[**BeginAsyncJobResponseV2**](BeginAsyncJobResponseV2.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PreviewMigrateReportToDashboard

> MigrationPreviewResponse PreviewMigrateReportToDashboard(ctx).DashboardMigrationRequest(dashboardMigrationRequest).Execute()

Preview of Migrating Legacy Dashboards to Dashboards(New)



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
    dashboardMigrationRequest := *openapiclient.NewDashboardMigrationRequest([]string{"00000000000001C8"}) // DashboardMigrationRequest | List of content identifiers. Can be folders or classic dashboard.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DashboardManagementAPI.PreviewMigrateReportToDashboard(context.Background()).DashboardMigrationRequest(dashboardMigrationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardManagementAPI.PreviewMigrateReportToDashboard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PreviewMigrateReportToDashboard`: MigrationPreviewResponse
    fmt.Fprintf(os.Stdout, "Response from `DashboardManagementAPI.PreviewMigrateReportToDashboard`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPreviewMigrateReportToDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dashboardMigrationRequest** | [**DashboardMigrationRequest**](DashboardMigrationRequest.md) | List of content identifiers. Can be folders or classic dashboard. | 

### Return type

[**MigrationPreviewResponse**](MigrationPreviewResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDashboard

> Dashboard UpdateDashboard(ctx, id).DashboardRequest(dashboardRequest).Execute()

Update a dashboard.



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
    id := "id_example" // string | Identifier of the dashboard to update.
    dashboardRequest := *openapiclient.NewDashboardRequest("Kubernetes Dashboard", *openapiclient.NewResolvableTimeRange("Type_example")) // DashboardRequest | Information to update on the dashboard.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DashboardManagementAPI.UpdateDashboard(context.Background(), id).DashboardRequest(dashboardRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardManagementAPI.UpdateDashboard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDashboard`: Dashboard
    fmt.Fprintf(os.Stdout, "Response from `DashboardManagementAPI.UpdateDashboard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the dashboard to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dashboardRequest** | [**DashboardRequest**](DashboardRequest.md) | Information to update on the dashboard. | 

### Return type

[**Dashboard**](Dashboard.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

