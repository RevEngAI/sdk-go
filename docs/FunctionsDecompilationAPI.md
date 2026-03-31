# FunctionsDecompilationAPI

All URIs are relative to *https://api.reveng.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDecompilationComment**](FunctionsDecompilationAPI.md#CreateDecompilationComment) | **Post** /v2/functions/{function_id}/decompilation/comments | Create a comment for this function
[**DeleteDecompilationComment**](FunctionsDecompilationAPI.md#DeleteDecompilationComment) | **Delete** /v2/functions/{function_id}/decompilation/comments/{comment_id} | Delete a comment
[**GetDecompilationComments**](FunctionsDecompilationAPI.md#GetDecompilationComments) | **Get** /v2/functions/{function_id}/decompilation/comments | Get comments for this function
[**UpdateDecompilationComment**](FunctionsDecompilationAPI.md#UpdateDecompilationComment) | **Patch** /v2/functions/{function_id}/decompilation/comments/{comment_id} | Update a comment



## CreateDecompilationComment

> BaseResponseCommentResponse CreateDecompilationComment(ctx, functionId).FunctionCommentCreateRequest(functionCommentCreateRequest).EndpointUrl(endpointUrl).LocalCacheDir(localCacheDir).LocalCacheMaxSizeMb(localCacheMaxSizeMb).CustomerSamplesBucket(customerSamplesBucket).FirmwareSamplesBucket(firmwareSamplesBucket).MaxRetryAttempts(maxRetryAttempts).Execute()

Create a comment for this function



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
	functionId := int32(56) // int32 | 
	functionCommentCreateRequest := *revengai.NewFunctionCommentCreateRequest("Content_example") // FunctionCommentCreateRequest | 
	endpointUrl := "endpointUrl_example" // string |  (optional)
	localCacheDir := "localCacheDir_example" // string |  (optional)
	localCacheMaxSizeMb := int32(56) // int32 |  (optional)
	customerSamplesBucket := "customerSamplesBucket_example" // string |  (optional)
	firmwareSamplesBucket := "firmwareSamplesBucket_example" // string |  (optional)
	maxRetryAttempts := int32(56) // int32 |  (optional) (default to 5)

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsDecompilationAPI.CreateDecompilationComment(context.Background(), functionId).FunctionCommentCreateRequest(functionCommentCreateRequest).EndpointUrl(endpointUrl).LocalCacheDir(localCacheDir).LocalCacheMaxSizeMb(localCacheMaxSizeMb).CustomerSamplesBucket(customerSamplesBucket).FirmwareSamplesBucket(firmwareSamplesBucket).MaxRetryAttempts(maxRetryAttempts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsDecompilationAPI.CreateDecompilationComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDecompilationComment`: BaseResponseCommentResponse
	fmt.Fprintf(os.Stdout, "Response from `FunctionsDecompilationAPI.CreateDecompilationComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDecompilationCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **functionCommentCreateRequest** | [**FunctionCommentCreateRequest**](FunctionCommentCreateRequest.md) |  | 
 **endpointUrl** | **string** |  | 
 **localCacheDir** | **string** |  | 
 **localCacheMaxSizeMb** | **int32** |  | 
 **customerSamplesBucket** | **string** |  | 
 **firmwareSamplesBucket** | **string** |  | 
 **maxRetryAttempts** | **int32** |  | [default to 5]

### Return type

[**BaseResponseCommentResponse**](BaseResponseCommentResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDecompilationComment

> BaseResponseBool DeleteDecompilationComment(ctx, commentId, functionId).EndpointUrl(endpointUrl).LocalCacheDir(localCacheDir).LocalCacheMaxSizeMb(localCacheMaxSizeMb).CustomerSamplesBucket(customerSamplesBucket).FirmwareSamplesBucket(firmwareSamplesBucket).MaxRetryAttempts(maxRetryAttempts).Execute()

Delete a comment



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
	commentId := int32(56) // int32 | 
	functionId := int32(56) // int32 | 
	endpointUrl := "endpointUrl_example" // string |  (optional)
	localCacheDir := "localCacheDir_example" // string |  (optional)
	localCacheMaxSizeMb := int32(56) // int32 |  (optional)
	customerSamplesBucket := "customerSamplesBucket_example" // string |  (optional)
	firmwareSamplesBucket := "firmwareSamplesBucket_example" // string |  (optional)
	maxRetryAttempts := int32(56) // int32 |  (optional) (default to 5)

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsDecompilationAPI.DeleteDecompilationComment(context.Background(), commentId, functionId).EndpointUrl(endpointUrl).LocalCacheDir(localCacheDir).LocalCacheMaxSizeMb(localCacheMaxSizeMb).CustomerSamplesBucket(customerSamplesBucket).FirmwareSamplesBucket(firmwareSamplesBucket).MaxRetryAttempts(maxRetryAttempts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsDecompilationAPI.DeleteDecompilationComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDecompilationComment`: BaseResponseBool
	fmt.Fprintf(os.Stdout, "Response from `FunctionsDecompilationAPI.DeleteDecompilationComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **int32** |  | 
**functionId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDecompilationCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **endpointUrl** | **string** |  | 
 **localCacheDir** | **string** |  | 
 **localCacheMaxSizeMb** | **int32** |  | 
 **customerSamplesBucket** | **string** |  | 
 **firmwareSamplesBucket** | **string** |  | 
 **maxRetryAttempts** | **int32** |  | [default to 5]

### Return type

[**BaseResponseBool**](BaseResponseBool.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDecompilationComments

> BaseResponseListCommentResponse GetDecompilationComments(ctx, functionId).EndpointUrl(endpointUrl).LocalCacheDir(localCacheDir).LocalCacheMaxSizeMb(localCacheMaxSizeMb).CustomerSamplesBucket(customerSamplesBucket).FirmwareSamplesBucket(firmwareSamplesBucket).MaxRetryAttempts(maxRetryAttempts).Execute()

Get comments for this function



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
	functionId := int32(56) // int32 | 
	endpointUrl := "endpointUrl_example" // string |  (optional)
	localCacheDir := "localCacheDir_example" // string |  (optional)
	localCacheMaxSizeMb := int32(56) // int32 |  (optional)
	customerSamplesBucket := "customerSamplesBucket_example" // string |  (optional)
	firmwareSamplesBucket := "firmwareSamplesBucket_example" // string |  (optional)
	maxRetryAttempts := int32(56) // int32 |  (optional) (default to 5)

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsDecompilationAPI.GetDecompilationComments(context.Background(), functionId).EndpointUrl(endpointUrl).LocalCacheDir(localCacheDir).LocalCacheMaxSizeMb(localCacheMaxSizeMb).CustomerSamplesBucket(customerSamplesBucket).FirmwareSamplesBucket(firmwareSamplesBucket).MaxRetryAttempts(maxRetryAttempts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsDecompilationAPI.GetDecompilationComments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDecompilationComments`: BaseResponseListCommentResponse
	fmt.Fprintf(os.Stdout, "Response from `FunctionsDecompilationAPI.GetDecompilationComments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDecompilationCommentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **endpointUrl** | **string** |  | 
 **localCacheDir** | **string** |  | 
 **localCacheMaxSizeMb** | **int32** |  | 
 **customerSamplesBucket** | **string** |  | 
 **firmwareSamplesBucket** | **string** |  | 
 **maxRetryAttempts** | **int32** |  | [default to 5]

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


## UpdateDecompilationComment

> BaseResponseCommentResponse UpdateDecompilationComment(ctx, commentId, functionId).CommentUpdateRequest(commentUpdateRequest).EndpointUrl(endpointUrl).LocalCacheDir(localCacheDir).LocalCacheMaxSizeMb(localCacheMaxSizeMb).CustomerSamplesBucket(customerSamplesBucket).FirmwareSamplesBucket(firmwareSamplesBucket).MaxRetryAttempts(maxRetryAttempts).Execute()

Update a comment



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
	commentId := int32(56) // int32 | 
	functionId := int32(56) // int32 | 
	commentUpdateRequest := *revengai.NewCommentUpdateRequest("Content_example") // CommentUpdateRequest | 
	endpointUrl := "endpointUrl_example" // string |  (optional)
	localCacheDir := "localCacheDir_example" // string |  (optional)
	localCacheMaxSizeMb := int32(56) // int32 |  (optional)
	customerSamplesBucket := "customerSamplesBucket_example" // string |  (optional)
	firmwareSamplesBucket := "firmwareSamplesBucket_example" // string |  (optional)
	maxRetryAttempts := int32(56) // int32 |  (optional) (default to 5)

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsDecompilationAPI.UpdateDecompilationComment(context.Background(), commentId, functionId).CommentUpdateRequest(commentUpdateRequest).EndpointUrl(endpointUrl).LocalCacheDir(localCacheDir).LocalCacheMaxSizeMb(localCacheMaxSizeMb).CustomerSamplesBucket(customerSamplesBucket).FirmwareSamplesBucket(firmwareSamplesBucket).MaxRetryAttempts(maxRetryAttempts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsDecompilationAPI.UpdateDecompilationComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDecompilationComment`: BaseResponseCommentResponse
	fmt.Fprintf(os.Stdout, "Response from `FunctionsDecompilationAPI.UpdateDecompilationComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **int32** |  | 
**functionId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDecompilationCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **commentUpdateRequest** | [**CommentUpdateRequest**](CommentUpdateRequest.md) |  | 
 **endpointUrl** | **string** |  | 
 **localCacheDir** | **string** |  | 
 **localCacheMaxSizeMb** | **int32** |  | 
 **customerSamplesBucket** | **string** |  | 
 **firmwareSamplesBucket** | **string** |  | 
 **maxRetryAttempts** | **int32** |  | [default to 5]

### Return type

[**BaseResponseCommentResponse**](BaseResponseCommentResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

