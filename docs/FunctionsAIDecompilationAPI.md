# FunctionsAIDecompilationAPI

All URIs are relative to *https://api.reveng.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAiDecompilationComment**](FunctionsAIDecompilationAPI.md#CreateAiDecompilationComment) | **Post** /v2/functions/{function_id}/ai-decompilation/comments | Create a comment for this function
[**CreateAiDecompilationTask**](FunctionsAIDecompilationAPI.md#CreateAiDecompilationTask) | **Post** /v2/functions/{function_id}/ai-decompilation | Begins AI Decompilation Process
[**DeleteAiDecompilationComment**](FunctionsAIDecompilationAPI.md#DeleteAiDecompilationComment) | **Delete** /v2/functions/{function_id}/ai-decompilation/comments/{comment_id} | Delete a comment
[**GetAiDecompilationComments**](FunctionsAIDecompilationAPI.md#GetAiDecompilationComments) | **Get** /v2/functions/{function_id}/ai-decompilation/comments | Get comments for this function
[**GetAiDecompilationRating**](FunctionsAIDecompilationAPI.md#GetAiDecompilationRating) | **Get** /v2/functions/{function_id}/ai-decompilation/rating | Get rating for AI decompilation
[**GetAiDecompilationTaskResult**](FunctionsAIDecompilationAPI.md#GetAiDecompilationTaskResult) | **Get** /v2/functions/{function_id}/ai-decompilation | Polls AI Decompilation Process
[**GetAiDecompilationTaskStatus**](FunctionsAIDecompilationAPI.md#GetAiDecompilationTaskStatus) | **Get** /v2/functions/{function_id}/ai-decompilation/status | Check the status of a function ai decompilation
[**UpdateAiDecompilationComment**](FunctionsAIDecompilationAPI.md#UpdateAiDecompilationComment) | **Patch** /v2/functions/{function_id}/ai-decompilation/comments/{comment_id} | Update a comment
[**UpsertAiDecompilationRating**](FunctionsAIDecompilationAPI.md#UpsertAiDecompilationRating) | **Patch** /v2/functions/{function_id}/ai-decompilation/rating | Upsert rating for AI decompilation



## CreateAiDecompilationComment

> BaseResponseCommentResponse CreateAiDecompilationComment(ctx, functionId).FunctionCommentCreateRequest(functionCommentCreateRequest).Execute()

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
	resp, r, err := apiClient.FunctionsAIDecompilationAPI.CreateAiDecompilationComment(context.Background(), functionId).FunctionCommentCreateRequest(functionCommentCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAIDecompilationAPI.CreateAiDecompilationComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAiDecompilationComment`: BaseResponseCommentResponse
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAIDecompilationAPI.CreateAiDecompilationComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAiDecompilationCommentRequest struct via the builder pattern


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


## CreateAiDecompilationTask

> BaseResponse CreateAiDecompilationTask(ctx, functionId).Execute()

Begins AI Decompilation Process



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
	functionId := int64(789) // int64 | The ID of the function for which we are creating the decompilation task

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsAIDecompilationAPI.CreateAiDecompilationTask(context.Background(), functionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAIDecompilationAPI.CreateAiDecompilationTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAiDecompilationTask`: BaseResponse
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAIDecompilationAPI.CreateAiDecompilationTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **int64** | The ID of the function for which we are creating the decompilation task | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAiDecompilationTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BaseResponse**](BaseResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAiDecompilationComment

> BaseResponseBool DeleteAiDecompilationComment(ctx, commentId, functionId).Execute()

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
	resp, r, err := apiClient.FunctionsAIDecompilationAPI.DeleteAiDecompilationComment(context.Background(), commentId, functionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAIDecompilationAPI.DeleteAiDecompilationComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAiDecompilationComment`: BaseResponseBool
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAIDecompilationAPI.DeleteAiDecompilationComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **int32** |  | 
**functionId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAiDecompilationCommentRequest struct via the builder pattern


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


## GetAiDecompilationComments

> BaseResponseListCommentResponse GetAiDecompilationComments(ctx, functionId).Execute()

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
	resp, r, err := apiClient.FunctionsAIDecompilationAPI.GetAiDecompilationComments(context.Background(), functionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAIDecompilationAPI.GetAiDecompilationComments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiDecompilationComments`: BaseResponseListCommentResponse
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAIDecompilationAPI.GetAiDecompilationComments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiDecompilationCommentsRequest struct via the builder pattern


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


## GetAiDecompilationRating

> BaseResponseGetAiDecompilationRatingResponse GetAiDecompilationRating(ctx, functionId).Execute()

Get rating for AI decompilation

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
	functionId := int64(789) // int64 | The ID of the function for which to get the rating

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsAIDecompilationAPI.GetAiDecompilationRating(context.Background(), functionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAIDecompilationAPI.GetAiDecompilationRating``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiDecompilationRating`: BaseResponseGetAiDecompilationRatingResponse
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAIDecompilationAPI.GetAiDecompilationRating`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **int64** | The ID of the function for which to get the rating | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiDecompilationRatingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BaseResponseGetAiDecompilationRatingResponse**](BaseResponseGetAiDecompilationRatingResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiDecompilationTaskResult

> BaseResponseGetAiDecompilationTask GetAiDecompilationTaskResult(ctx, functionId).Summarise(summarise).GenerateInlineComments(generateInlineComments).Execute()

Polls AI Decompilation Process



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
	functionId := int64(789) // int64 | The ID of the function being decompiled
	summarise := true // bool | Generate a summary for the decompilation (optional) (default to true)
	generateInlineComments := true // bool | Generate inline comments for the decompilation (only works if summarise is enabled) (optional) (default to true)

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsAIDecompilationAPI.GetAiDecompilationTaskResult(context.Background(), functionId).Summarise(summarise).GenerateInlineComments(generateInlineComments).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAIDecompilationAPI.GetAiDecompilationTaskResult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiDecompilationTaskResult`: BaseResponseGetAiDecompilationTask
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAIDecompilationAPI.GetAiDecompilationTaskResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **int64** | The ID of the function being decompiled | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiDecompilationTaskResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **summarise** | **bool** | Generate a summary for the decompilation | [default to true]
 **generateInlineComments** | **bool** | Generate inline comments for the decompilation (only works if summarise is enabled) | [default to true]

### Return type

[**BaseResponseGetAiDecompilationTask**](BaseResponseGetAiDecompilationTask.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiDecompilationTaskStatus

> BaseResponseFunctionTaskResponse GetAiDecompilationTaskStatus(ctx, functionId).Execute()

Check the status of a function ai decompilation

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
	functionId := int64(789) // int64 | The ID of the function being checked

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsAIDecompilationAPI.GetAiDecompilationTaskStatus(context.Background(), functionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAIDecompilationAPI.GetAiDecompilationTaskStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiDecompilationTaskStatus`: BaseResponseFunctionTaskResponse
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAIDecompilationAPI.GetAiDecompilationTaskStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **int64** | The ID of the function being checked | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiDecompilationTaskStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BaseResponseFunctionTaskResponse**](BaseResponseFunctionTaskResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAiDecompilationComment

> BaseResponseCommentResponse UpdateAiDecompilationComment(ctx, commentId, functionId).CommentUpdateRequest(commentUpdateRequest).Execute()

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
	resp, r, err := apiClient.FunctionsAIDecompilationAPI.UpdateAiDecompilationComment(context.Background(), commentId, functionId).CommentUpdateRequest(commentUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAIDecompilationAPI.UpdateAiDecompilationComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAiDecompilationComment`: BaseResponseCommentResponse
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAIDecompilationAPI.UpdateAiDecompilationComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **int32** |  | 
**functionId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAiDecompilationCommentRequest struct via the builder pattern


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


## UpsertAiDecompilationRating

> BaseResponse UpsertAiDecompilationRating(ctx, functionId).UpsertAiDecomplationRatingRequest(upsertAiDecomplationRatingRequest).Execute()

Upsert rating for AI decompilation

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
	functionId := int64(789) // int64 | The ID of the function being rated
	upsertAiDecomplationRatingRequest := *revengai.NewUpsertAiDecomplationRatingRequest(revengai.AiDecompilationRating("POSITIVE"), "Reason_example") // UpsertAiDecomplationRatingRequest | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsAIDecompilationAPI.UpsertAiDecompilationRating(context.Background(), functionId).UpsertAiDecomplationRatingRequest(upsertAiDecomplationRatingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAIDecompilationAPI.UpsertAiDecompilationRating``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpsertAiDecompilationRating`: BaseResponse
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAIDecompilationAPI.UpsertAiDecompilationRating`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **int64** | The ID of the function being rated | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpsertAiDecompilationRatingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upsertAiDecomplationRatingRequest** | [**UpsertAiDecomplationRatingRequest**](UpsertAiDecomplationRatingRequest.md) |  | 

### Return type

[**BaseResponse**](BaseResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

