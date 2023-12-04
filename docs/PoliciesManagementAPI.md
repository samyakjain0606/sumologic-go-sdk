# \PoliciesManagementAPI

All URIs are relative to *https://api.au.sumologic.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAuditPolicy**](PoliciesManagementAPI.md#GetAuditPolicy) | **Get** /v1/policies/audit | Get Audit policy.
[**GetDataAccessLevelPolicy**](PoliciesManagementAPI.md#GetDataAccessLevelPolicy) | **Get** /v1/policies/dataAccessLevel | Get Data Access Level policy.
[**GetMaxUserSessionTimeoutPolicy**](PoliciesManagementAPI.md#GetMaxUserSessionTimeoutPolicy) | **Get** /v1/policies/maxUserSessionTimeout | Get Max User Session Timeout policy.
[**GetSearchAuditPolicy**](PoliciesManagementAPI.md#GetSearchAuditPolicy) | **Get** /v1/policies/searchAudit | Get Search Audit policy.
[**GetShareDashboardsOutsideOrganizationPolicy**](PoliciesManagementAPI.md#GetShareDashboardsOutsideOrganizationPolicy) | **Get** /v1/policies/shareDashboardsOutsideOrganization | Get Share Dashboards Outside Organization policy.
[**GetUserConcurrentSessionsLimitPolicy**](PoliciesManagementAPI.md#GetUserConcurrentSessionsLimitPolicy) | **Get** /v1/policies/userConcurrentSessionsLimit | Get User Concurrent Sessions Limit policy.
[**SetAuditPolicy**](PoliciesManagementAPI.md#SetAuditPolicy) | **Put** /v1/policies/audit | Set Audit policy.
[**SetDataAccessLevelPolicy**](PoliciesManagementAPI.md#SetDataAccessLevelPolicy) | **Put** /v1/policies/dataAccessLevel | Set Data Access Level policy.
[**SetMaxUserSessionTimeoutPolicy**](PoliciesManagementAPI.md#SetMaxUserSessionTimeoutPolicy) | **Put** /v1/policies/maxUserSessionTimeout | Set Max User Session Timeout policy.
[**SetSearchAuditPolicy**](PoliciesManagementAPI.md#SetSearchAuditPolicy) | **Put** /v1/policies/searchAudit | Set Search Audit policy.
[**SetShareDashboardsOutsideOrganizationPolicy**](PoliciesManagementAPI.md#SetShareDashboardsOutsideOrganizationPolicy) | **Put** /v1/policies/shareDashboardsOutsideOrganization | Set Share Dashboards Outside Organization policy.
[**SetUserConcurrentSessionsLimitPolicy**](PoliciesManagementAPI.md#SetUserConcurrentSessionsLimitPolicy) | **Put** /v1/policies/userConcurrentSessionsLimit | Set User Concurrent Sessions Limit policy.



## GetAuditPolicy

> AuditPolicy GetAuditPolicy(ctx).Execute()

Get Audit policy.



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
    resp, r, err := apiClient.PoliciesManagementAPI.GetAuditPolicy(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesManagementAPI.GetAuditPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuditPolicy`: AuditPolicy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesManagementAPI.GetAuditPolicy`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuditPolicyRequest struct via the builder pattern


### Return type

[**AuditPolicy**](AuditPolicy.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataAccessLevelPolicy

> DataAccessLevelPolicy GetDataAccessLevelPolicy(ctx).Execute()

Get Data Access Level policy.



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
    resp, r, err := apiClient.PoliciesManagementAPI.GetDataAccessLevelPolicy(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesManagementAPI.GetDataAccessLevelPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDataAccessLevelPolicy`: DataAccessLevelPolicy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesManagementAPI.GetDataAccessLevelPolicy`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataAccessLevelPolicyRequest struct via the builder pattern


### Return type

[**DataAccessLevelPolicy**](DataAccessLevelPolicy.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMaxUserSessionTimeoutPolicy

> MaxUserSessionTimeoutPolicy GetMaxUserSessionTimeoutPolicy(ctx).Execute()

Get Max User Session Timeout policy.



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
    resp, r, err := apiClient.PoliciesManagementAPI.GetMaxUserSessionTimeoutPolicy(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesManagementAPI.GetMaxUserSessionTimeoutPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMaxUserSessionTimeoutPolicy`: MaxUserSessionTimeoutPolicy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesManagementAPI.GetMaxUserSessionTimeoutPolicy`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMaxUserSessionTimeoutPolicyRequest struct via the builder pattern


### Return type

[**MaxUserSessionTimeoutPolicy**](MaxUserSessionTimeoutPolicy.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSearchAuditPolicy

> SearchAuditPolicy GetSearchAuditPolicy(ctx).Execute()

Get Search Audit policy.



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
    resp, r, err := apiClient.PoliciesManagementAPI.GetSearchAuditPolicy(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesManagementAPI.GetSearchAuditPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSearchAuditPolicy`: SearchAuditPolicy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesManagementAPI.GetSearchAuditPolicy`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSearchAuditPolicyRequest struct via the builder pattern


### Return type

[**SearchAuditPolicy**](SearchAuditPolicy.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetShareDashboardsOutsideOrganizationPolicy

> ShareDashboardsOutsideOrganizationPolicy GetShareDashboardsOutsideOrganizationPolicy(ctx).Execute()

Get Share Dashboards Outside Organization policy.



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
    resp, r, err := apiClient.PoliciesManagementAPI.GetShareDashboardsOutsideOrganizationPolicy(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesManagementAPI.GetShareDashboardsOutsideOrganizationPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetShareDashboardsOutsideOrganizationPolicy`: ShareDashboardsOutsideOrganizationPolicy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesManagementAPI.GetShareDashboardsOutsideOrganizationPolicy`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetShareDashboardsOutsideOrganizationPolicyRequest struct via the builder pattern


### Return type

[**ShareDashboardsOutsideOrganizationPolicy**](ShareDashboardsOutsideOrganizationPolicy.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserConcurrentSessionsLimitPolicy

> UserConcurrentSessionsLimitPolicy GetUserConcurrentSessionsLimitPolicy(ctx).Execute()

Get User Concurrent Sessions Limit policy.



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
    resp, r, err := apiClient.PoliciesManagementAPI.GetUserConcurrentSessionsLimitPolicy(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesManagementAPI.GetUserConcurrentSessionsLimitPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserConcurrentSessionsLimitPolicy`: UserConcurrentSessionsLimitPolicy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesManagementAPI.GetUserConcurrentSessionsLimitPolicy`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserConcurrentSessionsLimitPolicyRequest struct via the builder pattern


### Return type

[**UserConcurrentSessionsLimitPolicy**](UserConcurrentSessionsLimitPolicy.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetAuditPolicy

> AuditPolicy SetAuditPolicy(ctx).AuditPolicy(auditPolicy).Execute()

Set Audit policy.



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
    auditPolicy := *openapiclient.NewAuditPolicy(true) // AuditPolicy | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PoliciesManagementAPI.SetAuditPolicy(context.Background()).AuditPolicy(auditPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesManagementAPI.SetAuditPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetAuditPolicy`: AuditPolicy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesManagementAPI.SetAuditPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetAuditPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **auditPolicy** | [**AuditPolicy**](AuditPolicy.md) |  | 

### Return type

[**AuditPolicy**](AuditPolicy.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetDataAccessLevelPolicy

> DataAccessLevelPolicy SetDataAccessLevelPolicy(ctx).DataAccessLevelPolicy(dataAccessLevelPolicy).Execute()

Set Data Access Level policy.



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
    dataAccessLevelPolicy := *openapiclient.NewDataAccessLevelPolicy(true) // DataAccessLevelPolicy | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PoliciesManagementAPI.SetDataAccessLevelPolicy(context.Background()).DataAccessLevelPolicy(dataAccessLevelPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesManagementAPI.SetDataAccessLevelPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetDataAccessLevelPolicy`: DataAccessLevelPolicy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesManagementAPI.SetDataAccessLevelPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetDataAccessLevelPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataAccessLevelPolicy** | [**DataAccessLevelPolicy**](DataAccessLevelPolicy.md) |  | 

### Return type

[**DataAccessLevelPolicy**](DataAccessLevelPolicy.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetMaxUserSessionTimeoutPolicy

> MaxUserSessionTimeoutPolicy SetMaxUserSessionTimeoutPolicy(ctx).MaxUserSessionTimeoutPolicy(maxUserSessionTimeoutPolicy).Execute()

Set Max User Session Timeout policy.



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
    maxUserSessionTimeoutPolicy := *openapiclient.NewMaxUserSessionTimeoutPolicy("1d") // MaxUserSessionTimeoutPolicy | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PoliciesManagementAPI.SetMaxUserSessionTimeoutPolicy(context.Background()).MaxUserSessionTimeoutPolicy(maxUserSessionTimeoutPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesManagementAPI.SetMaxUserSessionTimeoutPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetMaxUserSessionTimeoutPolicy`: MaxUserSessionTimeoutPolicy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesManagementAPI.SetMaxUserSessionTimeoutPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetMaxUserSessionTimeoutPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **maxUserSessionTimeoutPolicy** | [**MaxUserSessionTimeoutPolicy**](MaxUserSessionTimeoutPolicy.md) |  | 

### Return type

[**MaxUserSessionTimeoutPolicy**](MaxUserSessionTimeoutPolicy.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetSearchAuditPolicy

> SearchAuditPolicy SetSearchAuditPolicy(ctx).SearchAuditPolicy(searchAuditPolicy).Execute()

Set Search Audit policy.



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
    searchAuditPolicy := *openapiclient.NewSearchAuditPolicy(true) // SearchAuditPolicy | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PoliciesManagementAPI.SetSearchAuditPolicy(context.Background()).SearchAuditPolicy(searchAuditPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesManagementAPI.SetSearchAuditPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetSearchAuditPolicy`: SearchAuditPolicy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesManagementAPI.SetSearchAuditPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetSearchAuditPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchAuditPolicy** | [**SearchAuditPolicy**](SearchAuditPolicy.md) |  | 

### Return type

[**SearchAuditPolicy**](SearchAuditPolicy.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetShareDashboardsOutsideOrganizationPolicy

> ShareDashboardsOutsideOrganizationPolicy SetShareDashboardsOutsideOrganizationPolicy(ctx).ShareDashboardsOutsideOrganizationPolicy(shareDashboardsOutsideOrganizationPolicy).Execute()

Set Share Dashboards Outside Organization policy.



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
    shareDashboardsOutsideOrganizationPolicy := *openapiclient.NewShareDashboardsOutsideOrganizationPolicy(true) // ShareDashboardsOutsideOrganizationPolicy | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PoliciesManagementAPI.SetShareDashboardsOutsideOrganizationPolicy(context.Background()).ShareDashboardsOutsideOrganizationPolicy(shareDashboardsOutsideOrganizationPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesManagementAPI.SetShareDashboardsOutsideOrganizationPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetShareDashboardsOutsideOrganizationPolicy`: ShareDashboardsOutsideOrganizationPolicy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesManagementAPI.SetShareDashboardsOutsideOrganizationPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetShareDashboardsOutsideOrganizationPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **shareDashboardsOutsideOrganizationPolicy** | [**ShareDashboardsOutsideOrganizationPolicy**](ShareDashboardsOutsideOrganizationPolicy.md) |  | 

### Return type

[**ShareDashboardsOutsideOrganizationPolicy**](ShareDashboardsOutsideOrganizationPolicy.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetUserConcurrentSessionsLimitPolicy

> UserConcurrentSessionsLimitPolicy SetUserConcurrentSessionsLimitPolicy(ctx).UserConcurrentSessionsLimitPolicy(userConcurrentSessionsLimitPolicy).Execute()

Set User Concurrent Sessions Limit policy.



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
    userConcurrentSessionsLimitPolicy := *openapiclient.NewUserConcurrentSessionsLimitPolicy(true) // UserConcurrentSessionsLimitPolicy | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PoliciesManagementAPI.SetUserConcurrentSessionsLimitPolicy(context.Background()).UserConcurrentSessionsLimitPolicy(userConcurrentSessionsLimitPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesManagementAPI.SetUserConcurrentSessionsLimitPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetUserConcurrentSessionsLimitPolicy`: UserConcurrentSessionsLimitPolicy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesManagementAPI.SetUserConcurrentSessionsLimitPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetUserConcurrentSessionsLimitPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userConcurrentSessionsLimitPolicy** | [**UserConcurrentSessionsLimitPolicy**](UserConcurrentSessionsLimitPolicy.md) |  | 

### Return type

[**UserConcurrentSessionsLimitPolicy**](UserConcurrentSessionsLimitPolicy.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

