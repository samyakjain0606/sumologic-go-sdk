# \PasswordPolicyAPI

All URIs are relative to *https://api.au.sumologic.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPasswordPolicy**](PasswordPolicyAPI.md#GetPasswordPolicy) | **Get** /v1/passwordPolicy | Get the current password policy.
[**SetPasswordPolicy**](PasswordPolicyAPI.md#SetPasswordPolicy) | **Put** /v1/passwordPolicy | Update password policy.



## GetPasswordPolicy

> PasswordPolicy GetPasswordPolicy(ctx).Execute()

Get the current password policy.



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
    resp, r, err := apiClient.PasswordPolicyAPI.GetPasswordPolicy(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordPolicyAPI.GetPasswordPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPasswordPolicy`: PasswordPolicy
    fmt.Fprintf(os.Stdout, "Response from `PasswordPolicyAPI.GetPasswordPolicy`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPasswordPolicyRequest struct via the builder pattern


### Return type

[**PasswordPolicy**](PasswordPolicy.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetPasswordPolicy

> PasswordPolicy SetPasswordPolicy(ctx).PasswordPolicy(passwordPolicy).Execute()

Update password policy.



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
    passwordPolicy := *openapiclient.NewPasswordPolicy() // PasswordPolicy | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PasswordPolicyAPI.SetPasswordPolicy(context.Background()).PasswordPolicy(passwordPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordPolicyAPI.SetPasswordPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetPasswordPolicy`: PasswordPolicy
    fmt.Fprintf(os.Stdout, "Response from `PasswordPolicyAPI.SetPasswordPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetPasswordPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **passwordPolicy** | [**PasswordPolicy**](PasswordPolicy.md) |  | 

### Return type

[**PasswordPolicy**](PasswordPolicy.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

