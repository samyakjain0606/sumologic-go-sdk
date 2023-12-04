# \OrgsManagementAPI

All URIs are relative to *https://api.au.sumologic.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetChildUsages**](OrgsManagementAPI.md#GetChildUsages) | **Post** /v1/organizations/usages | Get usages for child orgs.



## GetChildUsages

> ChildUsageDetailsResponse GetChildUsages(ctx).ChildUsageDetailsRequest(childUsageDetailsRequest).Execute()

Get usages for child orgs.



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
    childUsageDetailsRequest := *openapiclient.NewChildUsageDetailsRequest() // ChildUsageDetailsRequest | Details for the usages to be fetched. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsManagementAPI.GetChildUsages(context.Background()).ChildUsageDetailsRequest(childUsageDetailsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsManagementAPI.GetChildUsages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetChildUsages`: ChildUsageDetailsResponse
    fmt.Fprintf(os.Stdout, "Response from `OrgsManagementAPI.GetChildUsages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetChildUsagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **childUsageDetailsRequest** | [**ChildUsageDetailsRequest**](ChildUsageDetailsRequest.md) | Details for the usages to be fetched. | 

### Return type

[**ChildUsageDetailsResponse**](ChildUsageDetailsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

