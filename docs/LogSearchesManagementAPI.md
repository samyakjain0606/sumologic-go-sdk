# \LogSearchesManagementAPI

All URIs are relative to *https://api.au.sumologic.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLogSearch**](LogSearchesManagementAPI.md#CreateLogSearch) | **Post** /v1/logSearches | Save a log search.
[**DeleteLogSearch**](LogSearchesManagementAPI.md#DeleteLogSearch) | **Delete** /v1/logSearches/{id} | Delete the saved log search.
[**GetLogSearch**](LogSearchesManagementAPI.md#GetLogSearch) | **Get** /v1/logSearches/{id} | Get the saved log search.
[**ListLogSearches**](LogSearchesManagementAPI.md#ListLogSearches) | **Get** /v1/logSearches | List all saved log searches.
[**UpdateLogSearch**](LogSearchesManagementAPI.md#UpdateLogSearch) | **Put** /v1/logSearches/{id} | Update the saved log Search.



## CreateLogSearch

> LogSearch CreateLogSearch(ctx).SaveLogSearchRequest(saveLogSearchRequest).Execute()

Save a log search.



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
    saveLogSearchRequest := *openapiclient.NewSaveLogSearchRequest("error {{sourceCategory}}| count by _sourceCategory", *openapiclient.NewResolvableTimeRange("Type_example"), "Short title", "000000000000001A") // SaveLogSearchRequest | The definition of the saved log search.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogSearchesManagementAPI.CreateLogSearch(context.Background()).SaveLogSearchRequest(saveLogSearchRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogSearchesManagementAPI.CreateLogSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLogSearch`: LogSearch
    fmt.Fprintf(os.Stdout, "Response from `LogSearchesManagementAPI.CreateLogSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLogSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **saveLogSearchRequest** | [**SaveLogSearchRequest**](SaveLogSearchRequest.md) | The definition of the saved log search. | 

### Return type

[**LogSearch**](LogSearch.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLogSearch

> DeleteLogSearch(ctx, id).Execute()

Delete the saved log search.



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
    id := "id_example" // string | Identifier of the saved log search.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.LogSearchesManagementAPI.DeleteLogSearch(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogSearchesManagementAPI.DeleteLogSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the saved log search. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLogSearchRequest struct via the builder pattern


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


## GetLogSearch

> LogSearch GetLogSearch(ctx, id).Execute()

Get the saved log search.



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
    id := "id_example" // string | Identifier of the saved log search.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogSearchesManagementAPI.GetLogSearch(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogSearchesManagementAPI.GetLogSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogSearch`: LogSearch
    fmt.Fprintf(os.Stdout, "Response from `LogSearchesManagementAPI.GetLogSearch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the saved log search. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LogSearch**](LogSearch.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLogSearches

> PaginatedLogSearches ListLogSearches(ctx).Limit(limit).Token(token).Execute()

List all saved log searches.



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
    limit := int32(50) // int32 | Limit the number of log searches returned in the response. The number of log searches returned may be less than the `limit`. (optional) (default to 50)
    token := "GDCiRv4vebF3UWFJQ1kySXBOR3Bzh69GR0RyWm9vCtc" // string | Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. `token` is set to null when no more pages are left. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogSearchesManagementAPI.ListLogSearches(context.Background()).Limit(limit).Token(token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogSearchesManagementAPI.ListLogSearches``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLogSearches`: PaginatedLogSearches
    fmt.Fprintf(os.Stdout, "Response from `LogSearchesManagementAPI.ListLogSearches`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLogSearchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit the number of log searches returned in the response. The number of log searches returned may be less than the &#x60;limit&#x60;. | [default to 50]
 **token** | **string** | Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. &#x60;token&#x60; is set to null when no more pages are left. | 

### Return type

[**PaginatedLogSearches**](PaginatedLogSearches.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLogSearch

> LogSearch UpdateLogSearch(ctx, id).LogSearchDefinition(logSearchDefinition).Execute()

Update the saved log Search.



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
    id := "id_example" // string | Identifier of the saved log search.
    logSearchDefinition := *openapiclient.NewLogSearchDefinition("error {{sourceCategory}}| count by _sourceCategory", *openapiclient.NewResolvableTimeRange("Type_example"), "Short title") // LogSearchDefinition | An updated saved log search definition.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogSearchesManagementAPI.UpdateLogSearch(context.Background(), id).LogSearchDefinition(logSearchDefinition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogSearchesManagementAPI.UpdateLogSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLogSearch`: LogSearch
    fmt.Fprintf(os.Stdout, "Response from `LogSearchesManagementAPI.UpdateLogSearch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the saved log search. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLogSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **logSearchDefinition** | [**LogSearchDefinition**](LogSearchDefinition.md) | An updated saved log search definition. | 

### Return type

[**LogSearch**](LogSearch.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

