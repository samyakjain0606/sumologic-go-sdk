# \OtCollectorManagementExternalAPI

All URIs are relative to *https://api.au.sumologic.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteOTCollector**](OtCollectorManagementExternalAPI.md#DeleteOTCollector) | **Delete** /v1/otCollectors/{id} | Delete an OT Collector.
[**DeleteOfflineOTCollectors**](OtCollectorManagementExternalAPI.md#DeleteOfflineOTCollectors) | **Delete** /v1/otCollectors/offline | Delete all Offline OT Collectors
[**GetOTCollector**](OtCollectorManagementExternalAPI.md#GetOTCollector) | **Get** /v1/otCollectors/{id} | Get OT Collector by ID.
[**GetOTCollectorsByNames**](OtCollectorManagementExternalAPI.md#GetOTCollectorsByNames) | **Get** /v1/otCollectors/otCollectorsByName | Get OT Collectors by name.
[**GetOTCollectorsCount**](OtCollectorManagementExternalAPI.md#GetOTCollectorsCount) | **Get** /v1/otCollectors/totalCount | Get a count of OT Collectors.
[**GetPaginatedOTCollectors**](OtCollectorManagementExternalAPI.md#GetPaginatedOTCollectors) | **Post** /v1/otCollectors | Get paginated list of OT Collectors



## DeleteOTCollector

> DeleteOTCollector(ctx, id).Execute()

Delete an OT Collector.



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
    id := "id_example" // string | Identifier of the OT Collector to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OtCollectorManagementExternalAPI.DeleteOTCollector(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OtCollectorManagementExternalAPI.DeleteOTCollector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the OT Collector to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOTCollectorRequest struct via the builder pattern


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


## DeleteOfflineOTCollectors

> DeleteOfflineOTCollectors(ctx).Execute()

Delete all Offline OT Collectors



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
    r, err := apiClient.OtCollectorManagementExternalAPI.DeleteOfflineOTCollectors(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OtCollectorManagementExternalAPI.DeleteOfflineOTCollectors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOfflineOTCollectorsRequest struct via the builder pattern


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


## GetOTCollector

> OTCollector GetOTCollector(ctx, id).Execute()

Get OT Collector by ID.



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
    id := "id_example" // string | Identifier of the OT Collector to get.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OtCollectorManagementExternalAPI.GetOTCollector(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OtCollectorManagementExternalAPI.GetOTCollector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOTCollector`: OTCollector
    fmt.Fprintf(os.Stdout, "Response from `OtCollectorManagementExternalAPI.GetOTCollector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the OT Collector to get. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOTCollectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OTCollector**](OTCollector.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOTCollectorsByNames

> OTCollectorListResponse GetOTCollectorsByNames(ctx).Names(names).Execute()

Get OT Collectors by name.



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
    names := []string{"Inner_example"} // []string | A required parameter that accepts a list of names for which we need to collect all metadata.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OtCollectorManagementExternalAPI.GetOTCollectorsByNames(context.Background()).Names(names).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OtCollectorManagementExternalAPI.GetOTCollectorsByNames``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOTCollectorsByNames`: OTCollectorListResponse
    fmt.Fprintf(os.Stdout, "Response from `OtCollectorManagementExternalAPI.GetOTCollectorsByNames`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOTCollectorsByNamesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **names** | **[]string** | A required parameter that accepts a list of names for which we need to collect all metadata. | 

### Return type

[**OTCollectorListResponse**](OTCollectorListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOTCollectorsCount

> OTCollectorCountResponse GetOTCollectorsCount(ctx).Execute()

Get a count of OT Collectors.



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
    resp, r, err := apiClient.OtCollectorManagementExternalAPI.GetOTCollectorsCount(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OtCollectorManagementExternalAPI.GetOTCollectorsCount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOTCollectorsCount`: OTCollectorCountResponse
    fmt.Fprintf(os.Stdout, "Response from `OtCollectorManagementExternalAPI.GetOTCollectorsCount`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOTCollectorsCountRequest struct via the builder pattern


### Return type

[**OTCollectorCountResponse**](OTCollectorCountResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPaginatedOTCollectors

> PaginatedOTCollectorsResponse GetPaginatedOTCollectors(ctx).PaginatedOTCollectorsRequest(paginatedOTCollectorsRequest).Execute()

Get paginated list of OT Collectors



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
    paginatedOTCollectorsRequest := *openapiclient.NewPaginatedOTCollectorsRequest() // PaginatedOTCollectorsRequest | pagination request details

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OtCollectorManagementExternalAPI.GetPaginatedOTCollectors(context.Background()).PaginatedOTCollectorsRequest(paginatedOTCollectorsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OtCollectorManagementExternalAPI.GetPaginatedOTCollectors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPaginatedOTCollectors`: PaginatedOTCollectorsResponse
    fmt.Fprintf(os.Stdout, "Response from `OtCollectorManagementExternalAPI.GetPaginatedOTCollectors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPaginatedOTCollectorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paginatedOTCollectorsRequest** | [**PaginatedOTCollectorsRequest**](PaginatedOTCollectorsRequest.md) | pagination request details | 

### Return type

[**PaginatedOTCollectorsResponse**](PaginatedOTCollectorsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

