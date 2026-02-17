# FunctionsDecompilationAPI

All URIs are relative to *https://api.reveng.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDecompilationComment**](FunctionsDecompilationAPI.md#CreateDecompilationComment) | **Post** /v2/functions/{function_id}/decompilation/comments | Create a comment for this function
[**DeleteDecompilationComment**](FunctionsDecompilationAPI.md#DeleteDecompilationComment) | **Delete** /v2/functions/{function_id}/decompilation/comments/{comment_id} | Delete a comment
[**GetDecompilationComments**](FunctionsDecompilationAPI.md#GetDecompilationComments) | **Get** /v2/functions/{function_id}/decompilation/comments | Get comments for this function
[**UpdateDecompilationComment**](FunctionsDecompilationAPI.md#UpdateDecompilationComment) | **Patch** /v2/functions/{function_id}/decompilation/comments/{comment_id} | Update a comment



## CreateDecompilationComment

> BaseResponseCommentResponse CreateDecompilationComment(ctx, functionId).FunctionCommentCreateRequest(functionCommentCreateRequest).Execute()

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

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsDecompilationAPI.CreateDecompilationComment(context.Background(), functionId).FunctionCommentCreateRequest(functionCommentCreateRequest).Execute()
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

> BaseResponseBool DeleteDecompilationComment(ctx, commentId, functionId).Execute()

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

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsDecompilationAPI.DeleteDecompilationComment(context.Background(), commentId, functionId).Execute()
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

> BaseResponseListCommentResponse GetDecompilationComments(ctx, functionId).Execute()

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

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsDecompilationAPI.GetDecompilationComments(context.Background(), functionId).Execute()
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

> BaseResponseCommentResponse UpdateDecompilationComment(ctx, commentId, functionId).CommentUpdateRequest(commentUpdateRequest).Execute()

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

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsDecompilationAPI.UpdateDecompilationComment(context.Background(), commentId, functionId).CommentUpdateRequest(commentUpdateRequest).Execute()
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

