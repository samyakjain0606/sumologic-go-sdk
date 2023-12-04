# \TracesAPI

All URIs are relative to *https://api.au.sumologic.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelTraceQuery**](TracesAPI.md#CancelTraceQuery) | **Delete** /v1/tracing/tracequery/{queryId} | Cancel a trace search query.
[**CreateTraceQuery**](TracesAPI.md#CreateTraceQuery) | **Post** /v1/tracing/tracequery | Run a trace search query asynchronously.
[**GetCriticalPath**](TracesAPI.md#GetCriticalPath) | **Get** /v1/tracing/traces/{traceId}/criticalPath | Get a critical path of a trace.
[**GetCriticalPathServiceBreakdown**](TracesAPI.md#GetCriticalPathServiceBreakdown) | **Get** /v1/tracing/traces/{traceId}/criticalPath/breakdown/service | Get a critical path service breakdown of a trace.
[**GetMetrics**](TracesAPI.md#GetMetrics) | **Get** /v1/tracing/metrics | Get trace search query metrics.
[**GetSpan**](TracesAPI.md#GetSpan) | **Get** /v1/tracing/traces/{traceId}/spans/{spanId} | Get span details.
[**GetSpanBillingInfo**](TracesAPI.md#GetSpanBillingInfo) | **Get** /v1/tracing/traces/{traceId}/spans/{spanId}/billingInfo | Get span billing details.
[**GetSpans**](TracesAPI.md#GetSpans) | **Get** /v1/tracing/traces/{traceId}/spans | Get a list of trace spans.
[**GetTrace**](TracesAPI.md#GetTrace) | **Get** /v1/tracing/traces/{traceId} | Get trace details.
[**GetTraceLightEvents**](TracesAPI.md#GetTraceLightEvents) | **Get** /v1/tracing/traces/{traceId}/traceEvents | Get a list of events (without their attributes) per span for a trace.
[**GetTraceQueryFieldValues**](TracesAPI.md#GetTraceQueryFieldValues) | **Get** /v1/tracing/tracequery/fields/{field}/values | Get trace search query filter field values.
[**GetTraceQueryFields**](TracesAPI.md#GetTraceQueryFields) | **Get** /v1/tracing/tracequery/fields | Get filter fields for trace search queries.
[**GetTraceQueryResult**](TracesAPI.md#GetTraceQueryResult) | **Get** /v1/tracing/tracequery/{queryId}/rows/{rowId}/traces | Get results of a trace search query.
[**GetTraceQueryStatus**](TracesAPI.md#GetTraceQueryStatus) | **Get** /v1/tracing/tracequery/{queryId}/status | Get a trace search query status.
[**TraceExists**](TracesAPI.md#TraceExists) | **Get** /v1/tracing/traces/{traceId}/exists | Check if the trace exists.



## CancelTraceQuery

> CancelTraceQuery(ctx, queryId).Execute()

Cancel a trace search query.



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
    queryId := "798a13dc1ceeb19a" // string | Identifier of the query to cancel.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TracesAPI.CancelTraceQuery(context.Background(), queryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TracesAPI.CancelTraceQuery``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCancelTraceQueryRequest struct via the builder pattern


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


## CreateTraceQuery

> CreateTraceQueryResponse CreateTraceQuery(ctx).AsyncTraceQueryRequest(asyncTraceQueryRequest).Execute()

Run a trace search query asynchronously.



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
    asyncTraceQueryRequest := *openapiclient.NewAsyncTraceQueryRequest([]openapiclient.AsyncTraceQueryRow{*openapiclient.NewAsyncTraceQueryRow(*openapiclient.NewTraceQueryExpression("Type_example"), "#A")}, *openapiclient.NewResolvableTimeRange("Type_example")) // AsyncTraceQueryRequest | Query parameters.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TracesAPI.CreateTraceQuery(context.Background()).AsyncTraceQueryRequest(asyncTraceQueryRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TracesAPI.CreateTraceQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTraceQuery`: CreateTraceQueryResponse
    fmt.Fprintf(os.Stdout, "Response from `TracesAPI.CreateTraceQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTraceQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asyncTraceQueryRequest** | [**AsyncTraceQueryRequest**](AsyncTraceQueryRequest.md) | Query parameters. | 

### Return type

[**CreateTraceQueryResponse**](CreateTraceQueryResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCriticalPath

> CriticalPathResponse GetCriticalPath(ctx, traceId).Limit(limit).Token(token).Execute()

Get a critical path of a trace.



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
    traceId := "traceId_example" // string | Identifier of the trace.
    limit := int32(56) // int32 | The maximum number of results to fetch. (optional) (default to 100)
    token := "token_example" // string | Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. `token` is set to null when no more pages are left. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TracesAPI.GetCriticalPath(context.Background(), traceId).Limit(limit).Token(token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TracesAPI.GetCriticalPath``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCriticalPath`: CriticalPathResponse
    fmt.Fprintf(os.Stdout, "Response from `TracesAPI.GetCriticalPath`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**traceId** | **string** | Identifier of the trace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCriticalPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | The maximum number of results to fetch. | [default to 100]
 **token** | **string** | Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. &#x60;token&#x60; is set to null when no more pages are left. | 

### Return type

[**CriticalPathResponse**](CriticalPathResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCriticalPathServiceBreakdown

> CriticalPathServiceBreakdownResponse GetCriticalPathServiceBreakdown(ctx, traceId).Execute()

Get a critical path service breakdown of a trace.



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
    traceId := "traceId_example" // string | Identifier of the trace.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TracesAPI.GetCriticalPathServiceBreakdown(context.Background(), traceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TracesAPI.GetCriticalPathServiceBreakdown``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCriticalPathServiceBreakdown`: CriticalPathServiceBreakdownResponse
    fmt.Fprintf(os.Stdout, "Response from `TracesAPI.GetCriticalPathServiceBreakdown`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**traceId** | **string** | Identifier of the trace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCriticalPathServiceBreakdownRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CriticalPathServiceBreakdownResponse**](CriticalPathServiceBreakdownResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetrics

> TraceMetricsResponse GetMetrics(ctx).Execute()

Get trace search query metrics.



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
    resp, r, err := apiClient.TracesAPI.GetMetrics(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TracesAPI.GetMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMetrics`: TraceMetricsResponse
    fmt.Fprintf(os.Stdout, "Response from `TracesAPI.GetMetrics`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetricsRequest struct via the builder pattern


### Return type

[**TraceMetricsResponse**](TraceMetricsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpan

> TraceSpanDetail GetSpan(ctx, traceId, spanId).Execute()

Get span details.



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
    traceId := "traceId_example" // string | Identifier of the trace the span belongs to.
    spanId := "spanId_example" // string | Identifier of the span to get the details.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TracesAPI.GetSpan(context.Background(), traceId, spanId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TracesAPI.GetSpan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSpan`: TraceSpanDetail
    fmt.Fprintf(os.Stdout, "Response from `TracesAPI.GetSpan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**traceId** | **string** | Identifier of the trace the span belongs to. | 
**spanId** | **string** | Identifier of the span to get the details. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TraceSpanDetail**](TraceSpanDetail.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpanBillingInfo

> TraceSpanBillingInfo GetSpanBillingInfo(ctx, traceId, spanId).Execute()

Get span billing details.



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
    traceId := "traceId_example" // string | Identifier of the trace the span belongs to.
    spanId := "spanId_example" // string | Identifier of the span to get the billing info.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TracesAPI.GetSpanBillingInfo(context.Background(), traceId, spanId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TracesAPI.GetSpanBillingInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSpanBillingInfo`: TraceSpanBillingInfo
    fmt.Fprintf(os.Stdout, "Response from `TracesAPI.GetSpanBillingInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**traceId** | **string** | Identifier of the trace the span belongs to. | 
**spanId** | **string** | Identifier of the span to get the billing info. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpanBillingInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TraceSpanBillingInfo**](TraceSpanBillingInfo.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpans

> TraceSpansResponse GetSpans(ctx, traceId).Limit(limit).Token(token).Execute()

Get a list of trace spans.



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
    traceId := "traceId_example" // string | Identifier of the trace to get the spans.
    limit := int32(56) // int32 | The maximum number of results to fetch. (optional) (default to 100)
    token := "token_example" // string | Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. `token` is set to null when no more pages are left. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TracesAPI.GetSpans(context.Background(), traceId).Limit(limit).Token(token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TracesAPI.GetSpans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSpans`: TraceSpansResponse
    fmt.Fprintf(os.Stdout, "Response from `TracesAPI.GetSpans`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**traceId** | **string** | Identifier of the trace to get the spans. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | The maximum number of results to fetch. | [default to 100]
 **token** | **string** | Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. &#x60;token&#x60; is set to null when no more pages are left. | 

### Return type

[**TraceSpansResponse**](TraceSpansResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrace

> TraceDetail GetTrace(ctx, traceId).Execute()

Get trace details.



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
    traceId := "traceId_example" // string | Identifier of the trace to get the details.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TracesAPI.GetTrace(context.Background(), traceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TracesAPI.GetTrace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTrace`: TraceDetail
    fmt.Fprintf(os.Stdout, "Response from `TracesAPI.GetTrace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**traceId** | **string** | Identifier of the trace to get the details. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTraceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TraceDetail**](TraceDetail.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTraceLightEvents

> TraceLightEventsResponse GetTraceLightEvents(ctx, traceId).Limit(limit).Token(token).Execute()

Get a list of events (without their attributes) per span for a trace.



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
    traceId := "695068749d21cd104222a95cabc4707c" // string | Identifier of the trace for which span events will be returned.
    limit := int32(56) // int32 | The maximum number of spans with events returned by a single query. (optional) (default to 100)
    token := "dlFXd0lhSkxzRjAwYnpVZkMrRmlhYnF4cGtNMWdnVEI" // string | Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. `token` is set to null when no more pages are left. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TracesAPI.GetTraceLightEvents(context.Background(), traceId).Limit(limit).Token(token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TracesAPI.GetTraceLightEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTraceLightEvents`: TraceLightEventsResponse
    fmt.Fprintf(os.Stdout, "Response from `TracesAPI.GetTraceLightEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**traceId** | **string** | Identifier of the trace for which span events will be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTraceLightEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | The maximum number of spans with events returned by a single query. | [default to 100]
 **token** | **string** | Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. &#x60;token&#x60; is set to null when no more pages are left. | 

### Return type

[**TraceLightEventsResponse**](TraceLightEventsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTraceQueryFieldValues

> TraceFieldValuesResponse GetTraceQueryFieldValues(ctx, field).Query(query).Limit(limit).Token(token).FieldType(fieldType).Execute()

Get trace search query filter field values.



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
    fieldType := "SpanEventAttribute" // string | Indicates the kind of a field. Possible values: `SpanAttribute`, `SpanEventAttribute`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TracesAPI.GetTraceQueryFieldValues(context.Background(), field).Query(query).Limit(limit).Token(token).FieldType(fieldType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TracesAPI.GetTraceQueryFieldValues``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTraceQueryFieldValues`: TraceFieldValuesResponse
    fmt.Fprintf(os.Stdout, "Response from `TracesAPI.GetTraceQueryFieldValues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**field** | **string** | Field identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTraceQueryFieldValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **query** | **string** | Search filter to apply on the values to be returned. Only values containing the search query term will be returned. | 
 **limit** | **int32** | The maximum number of results to fetch. | [default to 10]
 **token** | **string** | Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. &#x60;token&#x60; is set to null when no more pages are left. | 
 **fieldType** | **string** | Indicates the kind of a field. Possible values: &#x60;SpanAttribute&#x60;, &#x60;SpanEventAttribute&#x60;. | 

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


## GetTraceQueryFields

> TraceFieldsResponse GetTraceQueryFields(ctx).Execute()

Get filter fields for trace search queries.



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
    resp, r, err := apiClient.TracesAPI.GetTraceQueryFields(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TracesAPI.GetTraceQueryFields``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTraceQueryFields`: TraceFieldsResponse
    fmt.Fprintf(os.Stdout, "Response from `TracesAPI.GetTraceQueryFields`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTraceQueryFieldsRequest struct via the builder pattern


### Return type

[**TraceFieldsResponse**](TraceFieldsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTraceQueryResult

> TraceQueryResultResponse GetTraceQueryResult(ctx, queryId, rowId).Limit(limit).Token(token).Execute()

Get results of a trace search query.



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
    limit := int32(100) // int32 | Limit of the number of traces returned in the response. (optional) (default to 100)
    token := "dlFXd0lhSkxzRjAwYnpVZkMrRmlhYnF4cGtNMWdnVEI" // string | Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. `token` is set to null when no more pages are left. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TracesAPI.GetTraceQueryResult(context.Background(), queryId, rowId).Limit(limit).Token(token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TracesAPI.GetTraceQueryResult``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTraceQueryResult`: TraceQueryResultResponse
    fmt.Fprintf(os.Stdout, "Response from `TracesAPI.GetTraceQueryResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**queryId** | **string** | Identifier of the executed query. | 
**rowId** | **string** | Identifier of the query row. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTraceQueryResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | Limit of the number of traces returned in the response. | [default to 100]
 **token** | **string** | Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. &#x60;token&#x60; is set to null when no more pages are left. | 

### Return type

[**TraceQueryResultResponse**](TraceQueryResultResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTraceQueryStatus

> TraceQueryStatusResponse GetTraceQueryStatus(ctx, queryId).Execute()

Get a trace search query status.



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
    resp, r, err := apiClient.TracesAPI.GetTraceQueryStatus(context.Background(), queryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TracesAPI.GetTraceQueryStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTraceQueryStatus`: TraceQueryStatusResponse
    fmt.Fprintf(os.Stdout, "Response from `TracesAPI.GetTraceQueryStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**queryId** | **string** | Identifier of the executed query. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTraceQueryStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TraceQueryStatusResponse**](TraceQueryStatusResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TraceExists

> TraceExistsResponse TraceExists(ctx, traceId).Execute()

Check if the trace exists.



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
    traceId := "traceId_example" // string | Identifier of the trace to check.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TracesAPI.TraceExists(context.Background(), traceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TracesAPI.TraceExists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TraceExists`: TraceExistsResponse
    fmt.Fprintf(os.Stdout, "Response from `TracesAPI.TraceExists`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**traceId** | **string** | Identifier of the trace to check. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTraceExistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TraceExistsResponse**](TraceExistsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

