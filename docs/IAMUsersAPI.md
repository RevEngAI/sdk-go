# IAMUsersAPI

All URIs are relative to *https://api.reveng.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMe**](IAMUsersAPI.md#GetMe) | **Get** /v2/iam/me | Get current user
[**GetMyPermissions**](IAMUsersAPI.md#GetMyPermissions) | **Get** /v2/iam/me/permissions | Get current user permissions



## GetMe

> User GetMe(ctx).Execute()

Get current user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	revengai "github.com/RevEngAI/sdk-go/v3"
)

func main() {

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.IAMUsersAPI.GetMe(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IAMUsersAPI.GetMe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMe`: User
	fmt.Fprintf(os.Stdout, "Response from `IAMUsersAPI.GetMe`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMeRequest struct via the builder pattern


### Return type

[**User**](User.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMyPermissions

> Permissions GetMyPermissions(ctx).Execute()

Get current user permissions



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	revengai "github.com/RevEngAI/sdk-go/v3"
)

func main() {

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.IAMUsersAPI.GetMyPermissions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IAMUsersAPI.GetMyPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyPermissions`: Permissions
	fmt.Fprintf(os.Stdout, "Response from `IAMUsersAPI.GetMyPermissions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMyPermissionsRequest struct via the builder pattern


### Return type

[**Permissions**](Permissions.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

