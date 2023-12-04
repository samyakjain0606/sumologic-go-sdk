# \ConnectionManagementAPI

All URIs are relative to *https://api.au.sumologic.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConnection**](ConnectionManagementAPI.md#CreateConnection) | **Post** /v1/connections | Create a new connection.
[**DeleteConnection**](ConnectionManagementAPI.md#DeleteConnection) | **Delete** /v1/connections/{id} | Delete a connection.
[**GetConnection**](ConnectionManagementAPI.md#GetConnection) | **Get** /v1/connections/{id} | Get a connection.
[**GetIncidentTemplates**](ConnectionManagementAPI.md#GetIncidentTemplates) | **Post** /v1/connections/incidentTemplates | Get incident templates for CloudSOAR connections.
[**ListConnections**](ConnectionManagementAPI.md#ListConnections) | **Get** /v1/connections | Get a list of connections.
[**TestConnection**](ConnectionManagementAPI.md#TestConnection) | **Post** /v1/connections/test | Test a new connection url.
[**UpdateConnection**](ConnectionManagementAPI.md#UpdateConnection) | **Put** /v1/connections/{id} | Update a connection.



## CreateConnection

> Connection CreateConnection(ctx).ConnectionDefinition(connectionDefinition).Execute()

Create a new connection.



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
    connectionDefinition := *openapiclient.NewConnectionDefinition("Type_example", "Name_example") // ConnectionDefinition | Information about the new connection.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConnectionManagementAPI.CreateConnection(context.Background()).ConnectionDefinition(connectionDefinition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionManagementAPI.CreateConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateConnection`: Connection
    fmt.Fprintf(os.Stdout, "Response from `ConnectionManagementAPI.CreateConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **connectionDefinition** | [**ConnectionDefinition**](ConnectionDefinition.md) | Information about the new connection. | 

### Return type

[**Connection**](Connection.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConnection

> DeleteConnection(ctx, id).Type_(type_).Execute()

Delete a connection.



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
    id := "id_example" // string | Identifier of the connection to delete.
    type_ := "type__example" // string | Type of connection to delete. Valid values are `WebhookConnection`, `ServiceNowConnection`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ConnectionManagementAPI.DeleteConnection(context.Background(), id).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionManagementAPI.DeleteConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the connection to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | Type of connection to delete. Valid values are &#x60;WebhookConnection&#x60;, &#x60;ServiceNowConnection&#x60;. | 

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


## GetConnection

> Connection GetConnection(ctx, id).Type_(type_).Execute()

Get a connection.



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
    id := "id_example" // string | Identifier of connection to return.
    type_ := "type__example" // string | Type of connection to return. Valid values are `WebhookConnection`, `ServiceNowConnection`. (optional) (default to "WebhookConnection")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConnectionManagementAPI.GetConnection(context.Background(), id).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionManagementAPI.GetConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConnection`: Connection
    fmt.Fprintf(os.Stdout, "Response from `ConnectionManagementAPI.GetConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of connection to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | Type of connection to return. Valid values are &#x60;WebhookConnection&#x60;, &#x60;ServiceNowConnection&#x60;. | [default to &quot;WebhookConnection&quot;]

### Return type

[**Connection**](Connection.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIncidentTemplates

> GetIncidentTemplatesResponse GetIncidentTemplates(ctx).GetIncidentTemplatesRequest(getIncidentTemplatesRequest).Execute()

Get incident templates for CloudSOAR connections.



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
    getIncidentTemplatesRequest := *openapiclient.NewGetIncidentTemplatesRequest() // GetIncidentTemplatesRequest | Information about the new connection. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConnectionManagementAPI.GetIncidentTemplates(context.Background()).GetIncidentTemplatesRequest(getIncidentTemplatesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionManagementAPI.GetIncidentTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIncidentTemplates`: GetIncidentTemplatesResponse
    fmt.Fprintf(os.Stdout, "Response from `ConnectionManagementAPI.GetIncidentTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIncidentTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getIncidentTemplatesRequest** | [**GetIncidentTemplatesRequest**](GetIncidentTemplatesRequest.md) | Information about the new connection. | 

### Return type

[**GetIncidentTemplatesResponse**](GetIncidentTemplatesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConnections

> ListConnectionsResponse ListConnections(ctx).Limit(limit).Token(token).Execute()

Get a list of connections.



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
    limit := int32(56) // int32 | Limit the number of connections returned in the response. The number of connections returned may be less than the `limit`. (optional) (default to 100)
    token := "token_example" // string | Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. `token` is set to null when no more pages are left. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConnectionManagementAPI.ListConnections(context.Background()).Limit(limit).Token(token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionManagementAPI.ListConnections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConnections`: ListConnectionsResponse
    fmt.Fprintf(os.Stdout, "Response from `ConnectionManagementAPI.ListConnections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit the number of connections returned in the response. The number of connections returned may be less than the &#x60;limit&#x60;. | [default to 100]
 **token** | **string** | Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. &#x60;token&#x60; is set to null when no more pages are left. | 

### Return type

[**ListConnectionsResponse**](ListConnectionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestConnection

> TestConnectionResponse TestConnection(ctx).ConnectionDefinition(connectionDefinition).Functionalities(functionalities).ConnectionId(connectionId).Execute()

Test a new connection url.



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
    connectionDefinition := *openapiclient.NewConnectionDefinition("Type_example", "Name_example") // ConnectionDefinition | Information about the new connection.
    functionalities := []string{"Inner_example"} // []string | A comma-separated functionalities of webhook payload to test. Acceptable values: `alert`, `resolution`. (optional) (default to ["alert"])
    connectionId := "0000000000123ABC" // string | Unique identifier of an existing connection to test. It should be provided when the request body of an existing connection contains masked authorization headers. If not provided, the authorization headers will not be correctly unmasked, and the test may fail due to unauthorized access. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConnectionManagementAPI.TestConnection(context.Background()).ConnectionDefinition(connectionDefinition).Functionalities(functionalities).ConnectionId(connectionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionManagementAPI.TestConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestConnection`: TestConnectionResponse
    fmt.Fprintf(os.Stdout, "Response from `ConnectionManagementAPI.TestConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **connectionDefinition** | [**ConnectionDefinition**](ConnectionDefinition.md) | Information about the new connection. | 
 **functionalities** | **[]string** | A comma-separated functionalities of webhook payload to test. Acceptable values: &#x60;alert&#x60;, &#x60;resolution&#x60;. | [default to [&quot;alert&quot;]]
 **connectionId** | **string** | Unique identifier of an existing connection to test. It should be provided when the request body of an existing connection contains masked authorization headers. If not provided, the authorization headers will not be correctly unmasked, and the test may fail due to unauthorized access. | 

### Return type

[**TestConnectionResponse**](TestConnectionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConnection

> Connection UpdateConnection(ctx, id).ConnectionDefinition(connectionDefinition).Execute()

Update a connection.



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
    id := "id_example" // string | Identifier of the connection to update.
    connectionDefinition := *openapiclient.NewConnectionDefinition("Type_example", "Name_example") // ConnectionDefinition | Information to update about the connection.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConnectionManagementAPI.UpdateConnection(context.Background(), id).ConnectionDefinition(connectionDefinition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionManagementAPI.UpdateConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateConnection`: Connection
    fmt.Fprintf(os.Stdout, "Response from `ConnectionManagementAPI.UpdateConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the connection to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **connectionDefinition** | [**ConnectionDefinition**](ConnectionDefinition.md) | Information to update about the connection. | 

### Return type

[**Connection**](Connection.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

