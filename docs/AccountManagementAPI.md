# \AccountManagementAPI

All URIs are relative to *https://api.au.sumologic.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSubdomain**](AccountManagementAPI.md#CreateSubdomain) | **Post** /v1/account/subdomain | Create account subdomain.
[**DeletePendingUpdateRequest**](AccountManagementAPI.md#DeletePendingUpdateRequest) | **Delete** /v1/plan/pendingUpdateRequest | Delete the pending plan update request, if any.
[**DeleteSubdomain**](AccountManagementAPI.md#DeleteSubdomain) | **Delete** /v1/account/subdomain | Delete the configured subdomain.
[**ExportUsageReport**](AccountManagementAPI.md#ExportUsageReport) | **Post** /v1/account/usage/report | Export credits usage details as CSV.
[**GetAccountOwner**](AccountManagementAPI.md#GetAccountOwner) | **Get** /v1/account/accountOwner | Get the owner of an account.
[**GetPendingUpdateRequest**](AccountManagementAPI.md#GetPendingUpdateRequest) | **Get** /v1/plan/pendingUpdateRequest | Get the pending plan update request, if any.
[**GetStatus**](AccountManagementAPI.md#GetStatus) | **Get** /v1/account/status | Get overview of the account status.
[**GetStatusForReport**](AccountManagementAPI.md#GetStatusForReport) | **Get** /v1/account/usage/report/{jobId}/status | Get report generation status.
[**GetSubdomain**](AccountManagementAPI.md#GetSubdomain) | **Get** /v1/account/subdomain | Get the configured subdomain.
[**RecoverSubdomains**](AccountManagementAPI.md#RecoverSubdomains) | **Post** /v1/account/subdomain/recover | Recover subdomains for a user.
[**UpdateSubdomain**](AccountManagementAPI.md#UpdateSubdomain) | **Put** /v1/account/subdomain | Update account subdomain.



## CreateSubdomain

> SubdomainDefinitionResponse CreateSubdomain(ctx).ConfigureSubdomainRequest(configureSubdomainRequest).Execute()

Create account subdomain.



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
    configureSubdomainRequest := *openapiclient.NewConfigureSubdomainRequest("my-company") // ConfigureSubdomainRequest | The new subdomain.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountManagementAPI.CreateSubdomain(context.Background()).ConfigureSubdomainRequest(configureSubdomainRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountManagementAPI.CreateSubdomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSubdomain`: SubdomainDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountManagementAPI.CreateSubdomain`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSubdomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **configureSubdomainRequest** | [**ConfigureSubdomainRequest**](ConfigureSubdomainRequest.md) | The new subdomain. | 

### Return type

[**SubdomainDefinitionResponse**](SubdomainDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePendingUpdateRequest

> DeletePendingUpdateRequest(ctx).Execute()

Delete the pending plan update request, if any.



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
    r, err := apiClient.AccountManagementAPI.DeletePendingUpdateRequest(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountManagementAPI.DeletePendingUpdateRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePendingUpdateRequestRequest struct via the builder pattern


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


## DeleteSubdomain

> DeleteSubdomain(ctx).Execute()

Delete the configured subdomain.



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
    r, err := apiClient.AccountManagementAPI.DeleteSubdomain(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountManagementAPI.DeleteSubdomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSubdomainRequest struct via the builder pattern


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


## ExportUsageReport

> UsageReportResponse ExportUsageReport(ctx).UsageReportRequest(usageReportRequest).Execute()

Export credits usage details as CSV.



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
    usageReportRequest := *openapiclient.NewUsageReportRequest() // UsageReportRequest | Export Usage Report Request.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountManagementAPI.ExportUsageReport(context.Background()).UsageReportRequest(usageReportRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountManagementAPI.ExportUsageReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportUsageReport`: UsageReportResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountManagementAPI.ExportUsageReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportUsageReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **usageReportRequest** | [**UsageReportRequest**](UsageReportRequest.md) | Export Usage Report Request. | 

### Return type

[**UsageReportResponse**](UsageReportResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountOwner

> string GetAccountOwner(ctx).Execute()

Get the owner of an account.



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
    resp, r, err := apiClient.AccountManagementAPI.GetAccountOwner(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountManagementAPI.GetAccountOwner``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountOwner`: string
    fmt.Fprintf(os.Stdout, "Response from `AccountManagementAPI.GetAccountOwner`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountOwnerRequest struct via the builder pattern


### Return type

**string**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPendingUpdateRequest

> PendingUpdateRequest GetPendingUpdateRequest(ctx).Execute()

Get the pending plan update request, if any.



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
    resp, r, err := apiClient.AccountManagementAPI.GetPendingUpdateRequest(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountManagementAPI.GetPendingUpdateRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPendingUpdateRequest`: PendingUpdateRequest
    fmt.Fprintf(os.Stdout, "Response from `AccountManagementAPI.GetPendingUpdateRequest`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPendingUpdateRequestRequest struct via the builder pattern


### Return type

[**PendingUpdateRequest**](PendingUpdateRequest.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStatus

> AccountStatusResponse GetStatus(ctx).Execute()

Get overview of the account status.



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
    resp, r, err := apiClient.AccountManagementAPI.GetStatus(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountManagementAPI.GetStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatus`: AccountStatusResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountManagementAPI.GetStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatusRequest struct via the builder pattern


### Return type

[**AccountStatusResponse**](AccountStatusResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStatusForReport

> UsageReportStatusResponse GetStatusForReport(ctx, jobId).Execute()

Get report generation status.



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
    jobId := "jobId_example" // string | Job Id for the report to be exported.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountManagementAPI.GetStatusForReport(context.Background(), jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountManagementAPI.GetStatusForReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatusForReport`: UsageReportStatusResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountManagementAPI.GetStatusForReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Job Id for the report to be exported. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatusForReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UsageReportStatusResponse**](UsageReportStatusResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubdomain

> SubdomainDefinitionResponse GetSubdomain(ctx).Execute()

Get the configured subdomain.



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
    resp, r, err := apiClient.AccountManagementAPI.GetSubdomain(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountManagementAPI.GetSubdomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSubdomain`: SubdomainDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountManagementAPI.GetSubdomain`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubdomainRequest struct via the builder pattern


### Return type

[**SubdomainDefinitionResponse**](SubdomainDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecoverSubdomains

> RecoverSubdomains(ctx).Email(email).Execute()

Recover subdomains for a user.



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
    email := "email_example" // string | Email address of the user to get subdomain information.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AccountManagementAPI.RecoverSubdomains(context.Background()).Email(email).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountManagementAPI.RecoverSubdomains``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRecoverSubdomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | Email address of the user to get subdomain information. | 

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


## UpdateSubdomain

> SubdomainDefinitionResponse UpdateSubdomain(ctx).ConfigureSubdomainRequest(configureSubdomainRequest).Execute()

Update account subdomain.



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
    configureSubdomainRequest := *openapiclient.NewConfigureSubdomainRequest("my-company") // ConfigureSubdomainRequest | The new subdomain.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountManagementAPI.UpdateSubdomain(context.Background()).ConfigureSubdomainRequest(configureSubdomainRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountManagementAPI.UpdateSubdomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSubdomain`: SubdomainDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountManagementAPI.UpdateSubdomain`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSubdomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **configureSubdomainRequest** | [**ConfigureSubdomainRequest**](ConfigureSubdomainRequest.md) | The new subdomain. | 

### Return type

[**SubdomainDefinitionResponse**](SubdomainDefinitionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

