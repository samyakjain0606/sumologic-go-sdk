# \ServiceMapAPI

All URIs are relative to *https://api.au.sumologic.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetServiceMap**](ServiceMapAPI.md#GetServiceMap) | **Get** /v1/tracing/serviceMap | Get a service map.



## GetServiceMap

> ServiceMapResponse GetServiceMap(ctx).Execute()

Get a service map.



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
    resp, r, err := apiClient.ServiceMapAPI.GetServiceMap(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceMapAPI.GetServiceMap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceMap`: ServiceMapResponse
    fmt.Fprintf(os.Stdout, "Response from `ServiceMapAPI.GetServiceMap`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceMapRequest struct via the builder pattern


### Return type

[**ServiceMapResponse**](ServiceMapResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

