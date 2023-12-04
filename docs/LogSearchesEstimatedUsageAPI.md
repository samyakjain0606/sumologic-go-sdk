# \LogSearchesEstimatedUsageAPI

All URIs are relative to *https://api.au.sumologic.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLogSearchEstimatedUsage**](LogSearchesEstimatedUsageAPI.md#GetLogSearchEstimatedUsage) | **Post** /v1/logSearches/estimatedUsage | Gets estimated usage details.
[**GetLogSearchEstimatedUsageByTier**](LogSearchesEstimatedUsageAPI.md#GetLogSearchEstimatedUsageByTier) | **Post** /v1/logSearches/estimatedUsageByTier | Gets Tier Wise estimated usage details.



## GetLogSearchEstimatedUsage

> LogSearchEstimatedUsageDefinition GetLogSearchEstimatedUsage(ctx).LogSearchEstimatedUsageRequest(logSearchEstimatedUsageRequest).Execute()

Gets estimated usage details.



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
    logSearchEstimatedUsageRequest := *openapiclient.NewLogSearchEstimatedUsageRequest("error {{sourceCategory}}| count by _sourceCategory", *openapiclient.NewResolvableTimeRange("Type_example"), "America/Los_Angeles") // LogSearchEstimatedUsageRequest | The definition of the log search estimated usage.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogSearchesEstimatedUsageAPI.GetLogSearchEstimatedUsage(context.Background()).LogSearchEstimatedUsageRequest(logSearchEstimatedUsageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogSearchesEstimatedUsageAPI.GetLogSearchEstimatedUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogSearchEstimatedUsage`: LogSearchEstimatedUsageDefinition
    fmt.Fprintf(os.Stdout, "Response from `LogSearchesEstimatedUsageAPI.GetLogSearchEstimatedUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLogSearchEstimatedUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **logSearchEstimatedUsageRequest** | [**LogSearchEstimatedUsageRequest**](LogSearchEstimatedUsageRequest.md) | The definition of the log search estimated usage. | 

### Return type

[**LogSearchEstimatedUsageDefinition**](LogSearchEstimatedUsageDefinition.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogSearchEstimatedUsageByTier

> LogSearchEstimatedUsageByTierDefinition GetLogSearchEstimatedUsageByTier(ctx).LogSearchEstimatedUsageRequestV2(logSearchEstimatedUsageRequestV2).Execute()

Gets Tier Wise estimated usage details.



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
    logSearchEstimatedUsageRequestV2 := *openapiclient.NewLogSearchEstimatedUsageRequestV2("error {{sourceCategory}}| count by _sourceCategory", *openapiclient.NewResolvableTimeRange("Type_example"), "America/Los_Angeles") // LogSearchEstimatedUsageRequestV2 | The definition of the log search estimated usage.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogSearchesEstimatedUsageAPI.GetLogSearchEstimatedUsageByTier(context.Background()).LogSearchEstimatedUsageRequestV2(logSearchEstimatedUsageRequestV2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogSearchesEstimatedUsageAPI.GetLogSearchEstimatedUsageByTier``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogSearchEstimatedUsageByTier`: LogSearchEstimatedUsageByTierDefinition
    fmt.Fprintf(os.Stdout, "Response from `LogSearchesEstimatedUsageAPI.GetLogSearchEstimatedUsageByTier`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLogSearchEstimatedUsageByTierRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **logSearchEstimatedUsageRequestV2** | [**LogSearchEstimatedUsageRequestV2**](LogSearchEstimatedUsageRequestV2.md) | The definition of the log search estimated usage. | 

### Return type

[**LogSearchEstimatedUsageByTierDefinition**](LogSearchEstimatedUsageByTierDefinition.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

