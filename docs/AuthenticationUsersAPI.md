# AuthenticationUsersAPI

All URIs are relative to *https://api.reveng.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUser**](AuthenticationUsersAPI.md#GetUser) | **Get** /v2/users/{user_id} | Get a user&#39;s public information
[**GetUserActivity**](AuthenticationUsersAPI.md#GetUserActivity) | **Get** /v2/users/activity | Get auth user activity
[**SubmitUserFeedback**](AuthenticationUsersAPI.md#SubmitUserFeedback) | **Post** /v2/users/feedback | Submit feedback about the application
[**V3GetUser**](AuthenticationUsersAPI.md#V3GetUser) | **Get** /v3/users/{user_id} | Get a user&#39;s public information
[**V3GetUserActivity**](AuthenticationUsersAPI.md#V3GetUserActivity) | **Get** /v3/users/activity | Get the caller&#39;s activity feed
[**V3SubmitUserFeedback**](AuthenticationUsersAPI.md#V3SubmitUserFeedback) | **Post** /v3/users/feedback | Submit feedback



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
	revengai "github.com/RevEngAI/sdk-go/v4"
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

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

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
	revengai "github.com/RevEngAI/sdk-go/v4"
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

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitUserFeedback

> BaseResponse SubmitUserFeedback(ctx).SubmitUserFeedbackRequest(submitUserFeedbackRequest).Execute()

Submit feedback about the application



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	revengai "github.com/RevEngAI/sdk-go/v4"
)

func main() {
	submitUserFeedbackRequest := *revengai.NewSubmitUserFeedbackRequest("CurrentRoute_example", "Feedback_example") // SubmitUserFeedbackRequest | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationUsersAPI.SubmitUserFeedback(context.Background()).SubmitUserFeedbackRequest(submitUserFeedbackRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationUsersAPI.SubmitUserFeedback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubmitUserFeedback`: BaseResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationUsersAPI.SubmitUserFeedback`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitUserFeedbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **submitUserFeedbackRequest** | [**SubmitUserFeedbackRequest**](SubmitUserFeedbackRequest.md) |  | 

### Return type

[**BaseResponse**](BaseResponse.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3GetUser

> GetPublicUserOutputBody V3GetUser(ctx, userId).Execute()

Get a user's public information



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	revengai "github.com/RevEngAI/sdk-go/v4"
)

func main() {
	userId := int64(789) // int64 | User ID

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationUsersAPI.V3GetUser(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationUsersAPI.V3GetUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3GetUser`: GetPublicUserOutputBody
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationUsersAPI.V3GetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3GetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetPublicUserOutputBody**](GetPublicUserOutputBody.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3GetUserActivity

> GetUserActivityOutputBody V3GetUserActivity(ctx).Execute()

Get the caller's activity feed



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	revengai "github.com/RevEngAI/sdk-go/v4"
)

func main() {

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationUsersAPI.V3GetUserActivity(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationUsersAPI.V3GetUserActivity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3GetUserActivity`: GetUserActivityOutputBody
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationUsersAPI.V3GetUserActivity`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV3GetUserActivityRequest struct via the builder pattern


### Return type

[**GetUserActivityOutputBody**](GetUserActivityOutputBody.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3SubmitUserFeedback

> SubmitFeedbackOutputBody V3SubmitUserFeedback(ctx).SubmitFeedbackBody(submitFeedbackBody).Execute()

Submit feedback



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	revengai "github.com/RevEngAI/sdk-go/v4"
)

func main() {
	submitFeedbackBody := *revengai.NewSubmitFeedbackBody("CurrentRoute_example", "Feedback_example") // SubmitFeedbackBody | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationUsersAPI.V3SubmitUserFeedback(context.Background()).SubmitFeedbackBody(submitFeedbackBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationUsersAPI.V3SubmitUserFeedback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3SubmitUserFeedback`: SubmitFeedbackOutputBody
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationUsersAPI.V3SubmitUserFeedback`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3SubmitUserFeedbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **submitFeedbackBody** | [**SubmitFeedbackBody**](SubmitFeedbackBody.md) |  | 

### Return type

[**SubmitFeedbackOutputBody**](SubmitFeedbackOutputBody.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

