# \SpanAnalyticsAPI

All URIs are relative to *https://api.au.sumologic.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelSpanQuery**](SpanAnalyticsAPI.md#CancelSpanQuery) | **Delete** /v1/tracing/spanquery/{queryId} | Cancel a span analytics query.
[**CreateSpanQuery**](SpanAnalyticsAPI.md#CreateSpanQuery) | **Post** /v1/tracing/spanquery | Run a span analytics query asynchronously.
[**GetSpanQueryAggregates**](SpanAnalyticsAPI.md#GetSpanQueryAggregates) | **Get** /v1/tracing/spanquery/{queryId}/aggregates | Get span analytics query aggregated results.
[**GetSpanQueryFacets**](SpanAnalyticsAPI.md#GetSpanQueryFacets) | **Get** /v1/tracing/spanquery/{queryId}/rows/{rowId}/facets | Get a list of facets of a span analytics query.
[**GetSpanQueryFieldValues**](SpanAnalyticsAPI.md#GetSpanQueryFieldValues) | **Get** /v1/tracing/spanquery/fields/{field}/values | Get span analytics query filter field values.
[**GetSpanQueryFields**](SpanAnalyticsAPI.md#GetSpanQueryFields) | **Get** /v1/tracing/spanquery/fields | Get filter fields for span analytics queries.
[**GetSpanQueryResult**](SpanAnalyticsAPI.md#GetSpanQueryResult) | **Get** /v1/tracing/spanquery/{queryId}/rows/{rowId}/spans | Get results of a span analytics query.
[**GetSpanQueryStatus**](SpanAnalyticsAPI.md#GetSpanQueryStatus) | **Get** /v1/tracing/spanquery/{queryId}/status | Get a span analytics query status.
[**PauseSpanQuery**](SpanAnalyticsAPI.md#PauseSpanQuery) | **Put** /v1/tracing/spanquery/{queryId}/pause | Pause a span analytics query.
[**ResumeSpanQuery**](SpanAnalyticsAPI.md#ResumeSpanQuery) | **Put** /v1/tracing/spanquery/{queryId}/resume | Resume a span analytics query.



## CancelSpanQuery

> CancelSpanQuery(ctx, queryId).Execute()

Cancel a span analytics query.



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
    queryId := "195038749d21ad109242c95cbbc8709d" // string | Identifier of the query to cancel.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SpanAnalyticsAPI.CancelSpanQuery(context.Background(), queryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpanAnalyticsAPI.CancelSpanQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**queryId** | **string** | Identifier of the query to cancel. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelSpanQueryRequest struct via the builder pattern


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


## CreateSpanQuery

> SpanQueryResponse CreateSpanQuery(ctx).SpanQueryRequest(spanQueryRequest).Execute()

Run a span analytics query asynchronously.



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
    spanQueryRequest := *openapiclient.NewSpanQueryRequest([]openapiclient.SpanQueryRow{*openapiclient.NewSpanQueryRow("QueryString_example", "A")}, *openapiclient.NewResolvableTimeRange("Type_example")) // SpanQueryRequest | Query parameters.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpanAnalyticsAPI.CreateSpanQuery(context.Background()).SpanQueryRequest(spanQueryRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpanAnalyticsAPI.CreateSpanQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSpanQuery`: SpanQueryResponse
    fmt.Fprintf(os.Stdout, "Response from `SpanAnalyticsAPI.CreateSpanQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSpanQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **spanQueryRequest** | [**SpanQueryRequest**](SpanQueryRequest.md) | Query parameters. | 

### Return type

[**SpanQueryResponse**](SpanQueryResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpanQueryAggregates

> SpanQueryAggregateResponse GetSpanQueryAggregates(ctx, queryId).Execute()

Get span analytics query aggregated results.



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
    queryId := "195038749d21ad109242c95cbbc8709d" // string | Identifier of the executed query.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpanAnalyticsAPI.GetSpanQueryAggregates(context.Background(), queryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpanAnalyticsAPI.GetSpanQueryAggregates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSpanQueryAggregates`: SpanQueryAggregateResponse
    fmt.Fprintf(os.Stdout, "Response from `SpanAnalyticsAPI.GetSpanQueryAggregates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**queryId** | **string** | Identifier of the executed query. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpanQueryAggregatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SpanQueryAggregateResponse**](SpanQueryAggregateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpanQueryFacets

> SpanQueryResultFacetsResponse GetSpanQueryFacets(ctx, queryId, rowId).Execute()

Get a list of facets of a span analytics query.



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
    queryId := "195038749d21ad109242c95cbbc8709d" // string | Identifier of the executed query.
    rowId := "A" // string | Identifier of the query row.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpanAnalyticsAPI.GetSpanQueryFacets(context.Background(), queryId, rowId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpanAnalyticsAPI.GetSpanQueryFacets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSpanQueryFacets`: SpanQueryResultFacetsResponse
    fmt.Fprintf(os.Stdout, "Response from `SpanAnalyticsAPI.GetSpanQueryFacets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**queryId** | **string** | Identifier of the executed query. | 
**rowId** | **string** | Identifier of the query row. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpanQueryFacetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SpanQueryResultFacetsResponse**](SpanQueryResultFacetsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpanQueryFieldValues

> TraceFieldValuesResponse GetSpanQueryFieldValues(ctx, field).Query(query).Limit(limit).Token(token).Execute()

Get span analytics query filter field values.



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
    field := "field_example" // string | Field identifier.
    query := "query_example" // string | Search filter to apply on the values to be returned. Only values containing the search query term will be returned. (optional)
    limit := int32(56) // int32 | The maximum number of results to fetch. (optional) (default to 10)
    token := "token_example" // string | Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. `token` is set to null when no more pages are left. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpanAnalyticsAPI.GetSpanQueryFieldValues(context.Background(), field).Query(query).Limit(limit).Token(token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpanAnalyticsAPI.GetSpanQueryFieldValues``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSpanQueryFieldValues`: TraceFieldValuesResponse
    fmt.Fprintf(os.Stdout, "Response from `SpanAnalyticsAPI.GetSpanQueryFieldValues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**field** | **string** | Field identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpanQueryFieldValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **query** | **string** | Search filter to apply on the values to be returned. Only values containing the search query term will be returned. | 
 **limit** | **int32** | The maximum number of results to fetch. | [default to 10]
 **token** | **string** | Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. &#x60;token&#x60; is set to null when no more pages are left. | 

### Return type

[**TraceFieldValuesResponse**](TraceFieldValuesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpanQueryFields

> SpanQueryFieldsResponse GetSpanQueryFields(ctx).Execute()

Get filter fields for span analytics queries.



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
    resp, r, err := apiClient.SpanAnalyticsAPI.GetSpanQueryFields(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpanAnalyticsAPI.GetSpanQueryFields``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSpanQueryFields`: SpanQueryFieldsResponse
    fmt.Fprintf(os.Stdout, "Response from `SpanAnalyticsAPI.GetSpanQueryFields`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpanQueryFieldsRequest struct via the builder pattern


### Return type

[**SpanQueryFieldsResponse**](SpanQueryFieldsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpanQueryResult

> SpanQueryResultSpansResponse GetSpanQueryResult(ctx, queryId, rowId).Limit(limit).Token(token).Execute()

Get results of a span analytics query.



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
    queryId := "195038749d21ad109242c95cbbc8709d" // string | Identifier of the executed query.
    rowId := "A" // string | Identifier of the query row.
    limit := int32(100) // int32 | Limit of the number of spans returned in the response. (optional) (default to 100)
    token := "dlFXd0lhSkxzRjAwYnpVZkMrRmlhYnF4cGtNMWdnVEI" // string | Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. `token` is set to null when no more pages are left. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpanAnalyticsAPI.GetSpanQueryResult(context.Background(), queryId, rowId).Limit(limit).Token(token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpanAnalyticsAPI.GetSpanQueryResult``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSpanQueryResult`: SpanQueryResultSpansResponse
    fmt.Fprintf(os.Stdout, "Response from `SpanAnalyticsAPI.GetSpanQueryResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**queryId** | **string** | Identifier of the executed query. | 
**rowId** | **string** | Identifier of the query row. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpanQueryResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | Limit of the number of spans returned in the response. | [default to 100]
 **token** | **string** | Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. &#x60;token&#x60; is set to null when no more pages are left. | 

### Return type

[**SpanQueryResultSpansResponse**](SpanQueryResultSpansResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpanQueryStatus

> SpanQueryStatusResponse GetSpanQueryStatus(ctx, queryId).Execute()

Get a span analytics query status.



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
    queryId := "195038749d21ad109242c95cbbc8709d" // string | Identifier of the executed query.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SpanAnalyticsAPI.GetSpanQueryStatus(context.Background(), queryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpanAnalyticsAPI.GetSpanQueryStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSpanQueryStatus`: SpanQueryStatusResponse
    fmt.Fprintf(os.Stdout, "Response from `SpanAnalyticsAPI.GetSpanQueryStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**queryId** | **string** | Identifier of the executed query. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpanQueryStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SpanQueryStatusResponse**](SpanQueryStatusResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PauseSpanQuery

> PauseSpanQuery(ctx, queryId).Execute()

Pause a span analytics query.



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
    queryId := "195038749d21ad109242c95cbbc8709d" // string | Identifier of the query to pause.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SpanAnalyticsAPI.PauseSpanQuery(context.Background(), queryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpanAnalyticsAPI.PauseSpanQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**queryId** | **string** | Identifier of the query to pause. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPauseSpanQueryRequest struct via the builder pattern


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


## ResumeSpanQuery

> ResumeSpanQuery(ctx, queryId).Execute()

Resume a span analytics query.



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
    queryId := "195038749d21ad109242c95cbbc8709d" // string | Identifier of the query to resume.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SpanAnalyticsAPI.ResumeSpanQuery(context.Background(), queryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SpanAnalyticsAPI.ResumeSpanQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**queryId** | **string** | Identifier of the query to resume. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResumeSpanQueryRequest struct via the builder pattern


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

