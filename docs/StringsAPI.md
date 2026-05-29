# StringsAPI

All URIs are relative to *https://api.reveng.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddUserStringToAnalysis**](StringsAPI.md#AddUserStringToAnalysis) | **Post** /v3/analyses/{analysis_id}/user-provided-strings | Add a user-provided string to an analysis.
[**AddUserStringToFunction**](StringsAPI.md#AddUserStringToFunction) | **Post** /v3/functions/{function_id}/user-provided-strings | Add a user-provided string to a function.



## AddUserStringToAnalysis

> map[string]interface{} AddUserStringToAnalysis(ctx, analysisId).AddUserStringInputBody(addUserStringInputBody).Execute()

Add a user-provided string to an analysis.



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
	analysisId := int64(789) // int64 | Analysis ID
	addUserStringInputBody := *revengai.NewAddUserStringInputBody("String_example", int64(123)) // AddUserStringInputBody | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.StringsAPI.AddUserStringToAnalysis(context.Background(), analysisId).AddUserStringInputBody(addUserStringInputBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StringsAPI.AddUserStringToAnalysis``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddUserStringToAnalysis`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `StringsAPI.AddUserStringToAnalysis`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int64** | Analysis ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddUserStringToAnalysisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addUserStringInputBody** | [**AddUserStringInputBody**](AddUserStringInputBody.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddUserStringToFunction

> map[string]interface{} AddUserStringToFunction(ctx, functionId).AddUserStringToFunctionInputBody(addUserStringToFunctionInputBody).Execute()

Add a user-provided string to a function.



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
	functionId := int64(789) // int64 | Function ID
	addUserStringToFunctionInputBody := *revengai.NewAddUserStringToFunctionInputBody("String_example", int64(123)) // AddUserStringToFunctionInputBody | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.StringsAPI.AddUserStringToFunction(context.Background(), functionId).AddUserStringToFunctionInputBody(addUserStringToFunctionInputBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StringsAPI.AddUserStringToFunction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddUserStringToFunction`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `StringsAPI.AddUserStringToFunction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **int64** | Function ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddUserStringToFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addUserStringToFunctionInputBody** | [**AddUserStringToFunctionInputBody**](AddUserStringToFunctionInputBody.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

