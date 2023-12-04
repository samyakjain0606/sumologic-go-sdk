# \ExtractionRuleManagementAPI

All URIs are relative to *https://api.au.sumologic.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateExtractionRule**](ExtractionRuleManagementAPI.md#CreateExtractionRule) | **Post** /v1/extractionRules | Create a new field extraction rule.
[**DeleteExtractionRule**](ExtractionRuleManagementAPI.md#DeleteExtractionRule) | **Delete** /v1/extractionRules/{id} | Delete a field extraction rule.
[**GetExtractionRule**](ExtractionRuleManagementAPI.md#GetExtractionRule) | **Get** /v1/extractionRules/{id} | Get a field extraction rule.
[**ListExtractionRules**](ExtractionRuleManagementAPI.md#ListExtractionRules) | **Get** /v1/extractionRules | Get a list of field extraction rules.
[**UpdateExtractionRule**](ExtractionRuleManagementAPI.md#UpdateExtractionRule) | **Put** /v1/extractionRules/{id} | Update a field extraction rule.



## CreateExtractionRule

> ExtractionRule CreateExtractionRule(ctx).ExtractionRuleDefinition(extractionRuleDefinition).Execute()

Create a new field extraction rule.



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
    extractionRuleDefinition := *openapiclient.NewExtractionRuleDefinition("ExtractionRule123", "_sourceHost=127.0.0.1", "csv _raw extract 1 as f1") // ExtractionRuleDefinition | Information about the new field extraction rule.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExtractionRuleManagementAPI.CreateExtractionRule(context.Background()).ExtractionRuleDefinition(extractionRuleDefinition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExtractionRuleManagementAPI.CreateExtractionRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateExtractionRule`: ExtractionRule
    fmt.Fprintf(os.Stdout, "Response from `ExtractionRuleManagementAPI.CreateExtractionRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateExtractionRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **extractionRuleDefinition** | [**ExtractionRuleDefinition**](ExtractionRuleDefinition.md) | Information about the new field extraction rule. | 

### Return type

[**ExtractionRule**](ExtractionRule.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteExtractionRule

> DeleteExtractionRule(ctx, id).Execute()

Delete a field extraction rule.



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
    id := "id_example" // string | Identifier of the field extraction rule to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ExtractionRuleManagementAPI.DeleteExtractionRule(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExtractionRuleManagementAPI.DeleteExtractionRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the field extraction rule to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteExtractionRuleRequest struct via the builder pattern


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


## GetExtractionRule

> ExtractionRule GetExtractionRule(ctx, id).Execute()

Get a field extraction rule.



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
    id := "id_example" // string | Identifier of field extraction rule to return.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExtractionRuleManagementAPI.GetExtractionRule(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExtractionRuleManagementAPI.GetExtractionRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExtractionRule`: ExtractionRule
    fmt.Fprintf(os.Stdout, "Response from `ExtractionRuleManagementAPI.GetExtractionRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of field extraction rule to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExtractionRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExtractionRule**](ExtractionRule.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListExtractionRules

> ListExtractionRulesResponse ListExtractionRules(ctx).Limit(limit).Token(token).Execute()

Get a list of field extraction rules.



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
    limit := int32(56) // int32 | Limit the number of field extraction rules returned in the response. The number of field extraction rules returned may be less than the `limit`. (optional) (default to 100)
    token := "token_example" // string | Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExtractionRuleManagementAPI.ListExtractionRules(context.Background()).Limit(limit).Token(token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExtractionRuleManagementAPI.ListExtractionRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListExtractionRules`: ListExtractionRulesResponse
    fmt.Fprintf(os.Stdout, "Response from `ExtractionRuleManagementAPI.ListExtractionRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListExtractionRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit the number of field extraction rules returned in the response. The number of field extraction rules returned may be less than the &#x60;limit&#x60;. | [default to 100]
 **token** | **string** | Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. | 

### Return type

[**ListExtractionRulesResponse**](ListExtractionRulesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateExtractionRule

> ExtractionRule UpdateExtractionRule(ctx, id).UpdateExtractionRuleDefinition(updateExtractionRuleDefinition).Execute()

Update a field extraction rule.



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
    id := "id_example" // string | Identifier of the field extraction rule to update.
    updateExtractionRuleDefinition := *openapiclient.NewUpdateExtractionRuleDefinition("ExtractionRule123", "_sourceHost=127.0.0.1", "csv _raw extract 1 as f1", false) // UpdateExtractionRuleDefinition | Information to update about the field extraction rule.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExtractionRuleManagementAPI.UpdateExtractionRule(context.Background(), id).UpdateExtractionRuleDefinition(updateExtractionRuleDefinition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExtractionRuleManagementAPI.UpdateExtractionRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateExtractionRule`: ExtractionRule
    fmt.Fprintf(os.Stdout, "Response from `ExtractionRuleManagementAPI.UpdateExtractionRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the field extraction rule to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateExtractionRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateExtractionRuleDefinition** | [**UpdateExtractionRuleDefinition**](UpdateExtractionRuleDefinition.md) | Information to update about the field extraction rule. | 

### Return type

[**ExtractionRule**](ExtractionRule.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

