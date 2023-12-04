# \MetricsSearchesManagementAPI

All URIs are relative to *https://api.au.sumologic.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMetricsSearch**](MetricsSearchesManagementAPI.md#CreateMetricsSearch) | **Post** /v1/metricsSearches | Save a metrics search.
[**DeleteMetricsSearch**](MetricsSearchesManagementAPI.md#DeleteMetricsSearch) | **Delete** /v1/metricsSearches/{id} | Deletes a metrics search.
[**GetMetricsSearch**](MetricsSearchesManagementAPI.md#GetMetricsSearch) | **Get** /v1/metricsSearches/{id} | Get a metrics search.
[**UpdateMetricsSearch**](MetricsSearchesManagementAPI.md#UpdateMetricsSearch) | **Put** /v1/metricsSearches/{id} | Updates a metrics search.



## CreateMetricsSearch

> MetricsSearchInstance CreateMetricsSearch(ctx).SaveMetricsSearchRequest(saveMetricsSearchRequest).Execute()

Save a metrics search.



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
    saveMetricsSearchRequest := *openapiclient.NewSaveMetricsSearchRequest("Short title", "Long and detailed description", *openapiclient.NewResolvableTimeRange("Type_example"), []openapiclient.MetricsSearchQuery{*openapiclient.NewMetricsSearchQuery("A", "my_metric | avg")}, "000000000000001A") // SaveMetricsSearchRequest | The definition of the metrics search.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetricsSearchesManagementAPI.CreateMetricsSearch(context.Background()).SaveMetricsSearchRequest(saveMetricsSearchRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsSearchesManagementAPI.CreateMetricsSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMetricsSearch`: MetricsSearchInstance
    fmt.Fprintf(os.Stdout, "Response from `MetricsSearchesManagementAPI.CreateMetricsSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMetricsSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **saveMetricsSearchRequest** | [**SaveMetricsSearchRequest**](SaveMetricsSearchRequest.md) | The definition of the metrics search. | 

### Return type

[**MetricsSearchInstance**](MetricsSearchInstance.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMetricsSearch

> DeleteMetricsSearch(ctx, id).Execute()

Deletes a metrics search.



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
    id := "id_example" // string | Identifier of the metrics search.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MetricsSearchesManagementAPI.DeleteMetricsSearch(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsSearchesManagementAPI.DeleteMetricsSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the metrics search. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMetricsSearchRequest struct via the builder pattern


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


## GetMetricsSearch

> MetricsSearchInstance GetMetricsSearch(ctx, id).Execute()

Get a metrics search.



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
    id := "id_example" // string | Identifier of the metrics search.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetricsSearchesManagementAPI.GetMetricsSearch(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsSearchesManagementAPI.GetMetricsSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMetricsSearch`: MetricsSearchInstance
    fmt.Fprintf(os.Stdout, "Response from `MetricsSearchesManagementAPI.GetMetricsSearch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the metrics search. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetricsSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MetricsSearchInstance**](MetricsSearchInstance.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMetricsSearch

> MetricsSearchInstance UpdateMetricsSearch(ctx, id).MetricsSearchV1(metricsSearchV1).Execute()

Updates a metrics search.



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
    id := "id_example" // string | Identifier of the metrics search.
    metricsSearchV1 := *openapiclient.NewMetricsSearchV1("Short title", "Long and detailed description", *openapiclient.NewResolvableTimeRange("Type_example"), []openapiclient.MetricsSearchQuery{*openapiclient.NewMetricsSearchQuery("A", "my_metric | avg")}) // MetricsSearchV1 | An updated metrics search definition.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetricsSearchesManagementAPI.UpdateMetricsSearch(context.Background(), id).MetricsSearchV1(metricsSearchV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsSearchesManagementAPI.UpdateMetricsSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMetricsSearch`: MetricsSearchInstance
    fmt.Fprintf(os.Stdout, "Response from `MetricsSearchesManagementAPI.UpdateMetricsSearch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the metrics search. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMetricsSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **metricsSearchV1** | [**MetricsSearchV1**](MetricsSearchV1.md) | An updated metrics search definition. | 

### Return type

[**MetricsSearchInstance**](MetricsSearchInstance.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

