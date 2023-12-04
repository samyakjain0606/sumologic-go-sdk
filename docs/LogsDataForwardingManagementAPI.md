# \LogsDataForwardingManagementAPI

All URIs are relative to *https://api.au.sumologic.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDataForwardingBucket**](LogsDataForwardingManagementAPI.md#CreateDataForwardingBucket) | **Post** /v1/logsDataForwarding/destinations | Create an S3 data forwarding destination.
[**CreateDataForwardingRule**](LogsDataForwardingManagementAPI.md#CreateDataForwardingRule) | **Post** /v1/logsDataForwarding/rules | Create an S3 data forwarding rule.
[**DeleteDataForwardingBucket**](LogsDataForwardingManagementAPI.md#DeleteDataForwardingBucket) | **Delete** /v1/logsDataForwarding/destinations/{id} | Delete an S3 data forwarding destination.
[**DeleteDataForwardingRule**](LogsDataForwardingManagementAPI.md#DeleteDataForwardingRule) | **Delete** /v1/logsDataForwarding/rules/{indexId} | Delete an S3 data forwarding rule by its index.
[**GetDataForwardingBuckets**](LogsDataForwardingManagementAPI.md#GetDataForwardingBuckets) | **Get** /v1/logsDataForwarding/destinations | Get Amazon S3 data forwarding destinations.
[**GetDataForwardingDestination**](LogsDataForwardingManagementAPI.md#GetDataForwardingDestination) | **Get** /v1/logsDataForwarding/destinations/{id} | Get an S3 data forwarding destination.
[**GetDataForwardingRule**](LogsDataForwardingManagementAPI.md#GetDataForwardingRule) | **Get** /v1/logsDataForwarding/rules/{indexId} | Get an S3 data forwarding rule by its index.
[**GetRulesAndBuckets**](LogsDataForwardingManagementAPI.md#GetRulesAndBuckets) | **Get** /v1/logsDataForwarding/rules | Get all S3 data forwarding rules.
[**UpdateDataForwardingBucket**](LogsDataForwardingManagementAPI.md#UpdateDataForwardingBucket) | **Put** /v1/logsDataForwarding/destinations/{id} | Update an S3 data forwarding destination.
[**UpdateDataForwardingRule**](LogsDataForwardingManagementAPI.md#UpdateDataForwardingRule) | **Put** /v1/logsDataForwarding/rules/{indexId} | Update an S3 data forwarding rule by its index.



## CreateDataForwardingBucket

> BucketDefinition CreateDataForwardingBucket(ctx).CreateBucketDefinition(createBucketDefinition).Execute()

Create an S3 data forwarding destination.



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
    createBucketDefinition := *openapiclient.NewCreateBucketDefinition("df-destination", "RoleBased", "df-bucket") // CreateBucketDefinition | Parameters to create new S3 data forwarding destination.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogsDataForwardingManagementAPI.CreateDataForwardingBucket(context.Background()).CreateBucketDefinition(createBucketDefinition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsDataForwardingManagementAPI.CreateDataForwardingBucket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDataForwardingBucket`: BucketDefinition
    fmt.Fprintf(os.Stdout, "Response from `LogsDataForwardingManagementAPI.CreateDataForwardingBucket`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDataForwardingBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createBucketDefinition** | [**CreateBucketDefinition**](CreateBucketDefinition.md) | Parameters to create new S3 data forwarding destination. | 

### Return type

[**BucketDefinition**](BucketDefinition.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDataForwardingRule

> DataForwardingRule CreateDataForwardingRule(ctx).CreateDataForwardingRule(createDataForwardingRule).Execute()

Create an S3 data forwarding rule.



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
    createDataForwardingRule := *openapiclient.NewCreateDataForwardingRule("1", "1") // CreateDataForwardingRule | Parameters to create the new S3 data forwarding rule.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogsDataForwardingManagementAPI.CreateDataForwardingRule(context.Background()).CreateDataForwardingRule(createDataForwardingRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsDataForwardingManagementAPI.CreateDataForwardingRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDataForwardingRule`: DataForwardingRule
    fmt.Fprintf(os.Stdout, "Response from `LogsDataForwardingManagementAPI.CreateDataForwardingRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDataForwardingRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDataForwardingRule** | [**CreateDataForwardingRule**](CreateDataForwardingRule.md) | Parameters to create the new S3 data forwarding rule. | 

### Return type

[**DataForwardingRule**](DataForwardingRule.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDataForwardingBucket

> DeleteDataForwardingBucket(ctx, id).Execute()

Delete an S3 data forwarding destination.



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
    id := "1" // string | Identifier of the data forwarding destination to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.LogsDataForwardingManagementAPI.DeleteDataForwardingBucket(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsDataForwardingManagementAPI.DeleteDataForwardingBucket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the data forwarding destination to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDataForwardingBucketRequest struct via the builder pattern


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


## DeleteDataForwardingRule

> DeleteDataForwardingRule(ctx, indexId).Execute()

Delete an S3 data forwarding rule by its index.



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
    indexId := "1" // string | The `id` of the Partition or Scheduled View with the data forwarding rule to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.LogsDataForwardingManagementAPI.DeleteDataForwardingRule(context.Background(), indexId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsDataForwardingManagementAPI.DeleteDataForwardingRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**indexId** | **string** | The &#x60;id&#x60; of the Partition or Scheduled View with the data forwarding rule to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDataForwardingRuleRequest struct via the builder pattern


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


## GetDataForwardingBuckets

> GetDataForwardingDestinations GetDataForwardingBuckets(ctx).Limit(limit).Token(token).Execute()

Get Amazon S3 data forwarding destinations.



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
    limit := int32(56) // int32 | Limit the number of data forwarding destinations returned in the response. The number of data forwarding destinations returned may be less than the `limit`. (optional) (default to 10)
    token := "token_example" // string | Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. `token` is set to null when no more pages are left. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogsDataForwardingManagementAPI.GetDataForwardingBuckets(context.Background()).Limit(limit).Token(token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsDataForwardingManagementAPI.GetDataForwardingBuckets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDataForwardingBuckets`: GetDataForwardingDestinations
    fmt.Fprintf(os.Stdout, "Response from `LogsDataForwardingManagementAPI.GetDataForwardingBuckets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDataForwardingBucketsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit the number of data forwarding destinations returned in the response. The number of data forwarding destinations returned may be less than the &#x60;limit&#x60;. | [default to 10]
 **token** | **string** | Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. &#x60;token&#x60; is set to null when no more pages are left. | 

### Return type

[**GetDataForwardingDestinations**](GetDataForwardingDestinations.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataForwardingDestination

> BucketDefinition GetDataForwardingDestination(ctx, id).Execute()

Get an S3 data forwarding destination.



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
    id := "1" // string | Identifier of the S3 data forwarding destination to return.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogsDataForwardingManagementAPI.GetDataForwardingDestination(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsDataForwardingManagementAPI.GetDataForwardingDestination``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDataForwardingDestination`: BucketDefinition
    fmt.Fprintf(os.Stdout, "Response from `LogsDataForwardingManagementAPI.GetDataForwardingDestination`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the S3 data forwarding destination to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataForwardingDestinationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BucketDefinition**](BucketDefinition.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataForwardingRule

> RuleAndBucketDetail GetDataForwardingRule(ctx, indexId).Execute()

Get an S3 data forwarding rule by its index.



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
    indexId := "1" // string | The `id` of the Partition or Scheduled View.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogsDataForwardingManagementAPI.GetDataForwardingRule(context.Background(), indexId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsDataForwardingManagementAPI.GetDataForwardingRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDataForwardingRule`: RuleAndBucketDetail
    fmt.Fprintf(os.Stdout, "Response from `LogsDataForwardingManagementAPI.GetDataForwardingRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**indexId** | **string** | The &#x60;id&#x60; of the Partition or Scheduled View. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataForwardingRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RuleAndBucketDetail**](RuleAndBucketDetail.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRulesAndBuckets

> GetRulesAndBucketsResult GetRulesAndBuckets(ctx).Limit(limit).Token(token).Execute()

Get all S3 data forwarding rules.



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
    limit := int32(56) // int32 | Limit the number of data forwarding rules returned in the response. The number of data forwarding rules returned may be less than the `limit`. (optional) (default to 10)
    token := "token_example" // string | Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. `token` is set to null when no more pages are left. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogsDataForwardingManagementAPI.GetRulesAndBuckets(context.Background()).Limit(limit).Token(token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsDataForwardingManagementAPI.GetRulesAndBuckets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRulesAndBuckets`: GetRulesAndBucketsResult
    fmt.Fprintf(os.Stdout, "Response from `LogsDataForwardingManagementAPI.GetRulesAndBuckets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRulesAndBucketsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit the number of data forwarding rules returned in the response. The number of data forwarding rules returned may be less than the &#x60;limit&#x60;. | [default to 10]
 **token** | **string** | Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. &#x60;token&#x60; is set to null when no more pages are left. | 

### Return type

[**GetRulesAndBucketsResult**](GetRulesAndBucketsResult.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDataForwardingBucket

> BucketDefinition UpdateDataForwardingBucket(ctx, id).UpdateBucketDefinition(updateBucketDefinition).Execute()

Update an S3 data forwarding destination.



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
    id := "1" // string | Identifier of the data forwarding destination to update.
    updateBucketDefinition := *openapiclient.NewUpdateBucketDefinition("RoleBased") // UpdateBucketDefinition | Object with the updated parameters.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogsDataForwardingManagementAPI.UpdateDataForwardingBucket(context.Background(), id).UpdateBucketDefinition(updateBucketDefinition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsDataForwardingManagementAPI.UpdateDataForwardingBucket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDataForwardingBucket`: BucketDefinition
    fmt.Fprintf(os.Stdout, "Response from `LogsDataForwardingManagementAPI.UpdateDataForwardingBucket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the data forwarding destination to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDataForwardingBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateBucketDefinition** | [**UpdateBucketDefinition**](UpdateBucketDefinition.md) | Object with the updated parameters. | 

### Return type

[**BucketDefinition**](BucketDefinition.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDataForwardingRule

> DataForwardingRule UpdateDataForwardingRule(ctx, indexId).UpdateDataForwardingRule(updateDataForwardingRule).Execute()

Update an S3 data forwarding rule by its index.



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
    indexId := "1" // string | The `id` of the Partition or Scheduled View with the data forwarding rule to update.
    updateDataForwardingRule := *openapiclient.NewUpdateDataForwardingRule() // UpdateDataForwardingRule | Parameters of an S3 data forwarding rule.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogsDataForwardingManagementAPI.UpdateDataForwardingRule(context.Background(), indexId).UpdateDataForwardingRule(updateDataForwardingRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsDataForwardingManagementAPI.UpdateDataForwardingRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDataForwardingRule`: DataForwardingRule
    fmt.Fprintf(os.Stdout, "Response from `LogsDataForwardingManagementAPI.UpdateDataForwardingRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**indexId** | **string** | The &#x60;id&#x60; of the Partition or Scheduled View with the data forwarding rule to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDataForwardingRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDataForwardingRule** | [**UpdateDataForwardingRule**](UpdateDataForwardingRule.md) | Parameters of an S3 data forwarding rule. | 

### Return type

[**DataForwardingRule**](DataForwardingRule.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

