# AuthenticationUsersAPI

All URIs are relative to *https://api.reveng.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRequesterUserInfo**](AuthenticationUsersAPI.md#GetRequesterUserInfo) | **Get** /v2/users/me | Get the requesters user information
[**GetUser**](AuthenticationUsersAPI.md#GetUser) | **Get** /v2/users/{user_id} | Get a user&#39;s public information
[**GetUserActivity**](AuthenticationUsersAPI.md#GetUserActivity) | **Get** /v2/users/activity | Get auth user activity
[**GetUserComments**](AuthenticationUsersAPI.md#GetUserComments) | **Get** /v2/users/me/comments | Get comments by user
[**LoginUser**](AuthenticationUsersAPI.md#LoginUser) | **Post** /v2/auth/login | Authenticate a user



## GetRequesterUserInfo

> BaseResponseGetMeResponse GetRequesterUserInfo(ctx).Execute()

Get the requesters user information

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
	resp, r, err := apiClient.AuthenticationUsersAPI.GetRequesterUserInfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationUsersAPI.GetRequesterUserInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRequesterUserInfo`: BaseResponseGetMeResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationUsersAPI.GetRequesterUserInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequesterUserInfoRequest struct via the builder pattern


### Return type

[**BaseResponseGetMeResponse**](BaseResponseGetMeResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUser

> BaseResponseGetPublicUserResponse GetUser(ctx, userId).Execute()

Get a user's public information

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
	userId := int32(56) // int32 | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationUsersAPI.GetUser(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationUsersAPI.GetUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUser`: BaseResponseGetPublicUserResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationUsersAPI.GetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BaseResponseGetPublicUserResponse**](BaseResponseGetPublicUserResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserActivity

> BaseResponseListUserActivityResponse GetUserActivity(ctx).Execute()

Get auth user activity

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
	resp, r, err := apiClient.AuthenticationUsersAPI.GetUserActivity(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationUsersAPI.GetUserActivity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserActivity`: BaseResponseListUserActivityResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationUsersAPI.GetUserActivity`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserActivityRequest struct via the builder pattern


### Return type

[**BaseResponseListUserActivityResponse**](BaseResponseListUserActivityResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserComments

> BaseResponseListCommentResponse GetUserComments(ctx).Execute()

Get comments by user



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
	resp, r, err := apiClient.AuthenticationUsersAPI.GetUserComments(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationUsersAPI.GetUserComments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserComments`: BaseResponseListCommentResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationUsersAPI.GetUserComments`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserCommentsRequest struct via the builder pattern


### Return type

[**BaseResponseListCommentResponse**](BaseResponseListCommentResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LoginUser

> BaseResponseLoginResponse LoginUser(ctx).LoginRequest(loginRequest).Execute()

Authenticate a user



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
	loginRequest := *revengai.NewLoginRequest("Username_example", "Password_example") // LoginRequest | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationUsersAPI.LoginUser(context.Background()).LoginRequest(loginRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationUsersAPI.LoginUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LoginUser`: BaseResponseLoginResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationUsersAPI.LoginUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLoginUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **loginRequest** | [**LoginRequest**](LoginRequest.md) |  | 

### Return type

[**BaseResponseLoginResponse**](BaseResponseLoginResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

