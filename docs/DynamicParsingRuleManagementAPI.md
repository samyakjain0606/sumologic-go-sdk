# \DynamicParsingRuleManagementAPI

All URIs are relative to *https://api.au.sumologic.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDynamicParsingRule**](DynamicParsingRuleManagementAPI.md#CreateDynamicParsingRule) | **Post** /v1/dynamicParsingRules | Create a new dynamic parsing rule.
[**DeleteDynamicParsingRule**](DynamicParsingRuleManagementAPI.md#DeleteDynamicParsingRule) | **Delete** /v1/dynamicParsingRules/{id} | Delete a dynamic parsing rule.
[**GetDynamicParsingRule**](DynamicParsingRuleManagementAPI.md#GetDynamicParsingRule) | **Get** /v1/dynamicParsingRules/{id} | Get a dynamic parsing rule.
[**ListDynamicParsingRules**](DynamicParsingRuleManagementAPI.md#ListDynamicParsingRules) | **Get** /v1/dynamicParsingRules | Get a list of dynamic parsing rules.
[**UpdateDynamicParsingRule**](DynamicParsingRuleManagementAPI.md#UpdateDynamicParsingRule) | **Put** /v1/dynamicParsingRules/{id} | Update a dynamic parsing rule.



## CreateDynamicParsingRule

> DynamicRule CreateDynamicParsingRule(ctx).DynamicRuleDefinition(dynamicRuleDefinition).Execute()

Create a new dynamic parsing rule.



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
    dynamicRuleDefinition := *openapiclient.NewDynamicRuleDefinition("DynamicParsingRule123", "_sourceHost=127.0.0.1", false) // DynamicRuleDefinition | Information about the new dynamic parsing rule.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DynamicParsingRuleManagementAPI.CreateDynamicParsingRule(context.Background()).DynamicRuleDefinition(dynamicRuleDefinition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DynamicParsingRuleManagementAPI.CreateDynamicParsingRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDynamicParsingRule`: DynamicRule
    fmt.Fprintf(os.Stdout, "Response from `DynamicParsingRuleManagementAPI.CreateDynamicParsingRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDynamicParsingRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dynamicRuleDefinition** | [**DynamicRuleDefinition**](DynamicRuleDefinition.md) | Information about the new dynamic parsing rule. | 

### Return type

[**DynamicRule**](DynamicRule.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDynamicParsingRule

> DeleteDynamicParsingRule(ctx, id).Execute()

Delete a dynamic parsing rule.



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
    id := "0000000001C41EE4" // string | Identifier of the dynamic parsing rule to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DynamicParsingRuleManagementAPI.DeleteDynamicParsingRule(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DynamicParsingRuleManagementAPI.DeleteDynamicParsingRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the dynamic parsing rule to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDynamicParsingRuleRequest struct via the builder pattern


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


## GetDynamicParsingRule

> DynamicRule GetDynamicParsingRule(ctx, id).Execute()

Get a dynamic parsing rule.



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
    id := "0000000001C41EE4" // string | Identifier of dynamic parsing rule to return.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DynamicParsingRuleManagementAPI.GetDynamicParsingRule(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DynamicParsingRuleManagementAPI.GetDynamicParsingRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDynamicParsingRule`: DynamicRule
    fmt.Fprintf(os.Stdout, "Response from `DynamicParsingRuleManagementAPI.GetDynamicParsingRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of dynamic parsing rule to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDynamicParsingRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DynamicRule**](DynamicRule.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDynamicParsingRules

> ListDynamicRulesResponse ListDynamicParsingRules(ctx).Limit(limit).Token(token).Execute()

Get a list of dynamic parsing rules.



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
    limit := int32(10) // int32 | Limit the number of dynamic parsing rules returned in the response. The number of dynamic parsing rules returned may be less than the `limit`. (optional) (default to 100)
    token := "0000000001C51FF7" // string | Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DynamicParsingRuleManagementAPI.ListDynamicParsingRules(context.Background()).Limit(limit).Token(token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DynamicParsingRuleManagementAPI.ListDynamicParsingRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDynamicParsingRules`: ListDynamicRulesResponse
    fmt.Fprintf(os.Stdout, "Response from `DynamicParsingRuleManagementAPI.ListDynamicParsingRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDynamicParsingRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit the number of dynamic parsing rules returned in the response. The number of dynamic parsing rules returned may be less than the &#x60;limit&#x60;. | [default to 100]
 **token** | **string** | Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. | 

### Return type

[**ListDynamicRulesResponse**](ListDynamicRulesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDynamicParsingRule

> DynamicRule UpdateDynamicParsingRule(ctx, id).DynamicRuleDefinition(dynamicRuleDefinition).Execute()

Update a dynamic parsing rule.



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
    id := "0000000001C41EE4" // string | Identifier of the dynamic parsing rule to update.
    dynamicRuleDefinition := *openapiclient.NewDynamicRuleDefinition("DynamicParsingRule123", "_sourceHost=127.0.0.1", false) // DynamicRuleDefinition | Information to update about the dynamic parsing rule.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DynamicParsingRuleManagementAPI.UpdateDynamicParsingRule(context.Background(), id).DynamicRuleDefinition(dynamicRuleDefinition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DynamicParsingRuleManagementAPI.UpdateDynamicParsingRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDynamicParsingRule`: DynamicRule
    fmt.Fprintf(os.Stdout, "Response from `DynamicParsingRuleManagementAPI.UpdateDynamicParsingRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the dynamic parsing rule to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDynamicParsingRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dynamicRuleDefinition** | [**DynamicRuleDefinition**](DynamicRuleDefinition.md) | Information to update about the dynamic parsing rule. | 

### Return type

[**DynamicRule**](DynamicRule.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

