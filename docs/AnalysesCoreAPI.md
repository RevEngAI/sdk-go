# AnalysesCoreAPI

All URIs are relative to *https://api.reveng.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAnalysis**](AnalysesCoreAPI.md#CreateAnalysis) | **Post** /v2/analyses | Create Analysis
[**DeleteAnalysis**](AnalysesCoreAPI.md#DeleteAnalysis) | **Delete** /v2/analyses/{analysis_id} | Delete Analysis
[**GetAnalysisBasicInfo**](AnalysesCoreAPI.md#GetAnalysisBasicInfo) | **Get** /v2/analyses/{analysis_id}/basic | Gets basic analysis information
[**GetAnalysisFunctionMap**](AnalysesCoreAPI.md#GetAnalysisFunctionMap) | **Get** /v2/analyses/{analysis_id}/func_maps | Get Analysis Function Map
[**GetAnalysisLogs**](AnalysesCoreAPI.md#GetAnalysisLogs) | **Get** /v2/analyses/{analysis_id}/logs | Gets the logs of an analysis
[**GetAnalysisParams**](AnalysesCoreAPI.md#GetAnalysisParams) | **Get** /v2/analyses/{analysis_id}/params | Gets analysis param information
[**GetAnalysisStatus**](AnalysesCoreAPI.md#GetAnalysisStatus) | **Get** /v2/analyses/{analysis_id}/status | Gets the status of an analysis
[**InsertAnalysisLog**](AnalysesCoreAPI.md#InsertAnalysisLog) | **Post** /v2/analyses/{analysis_id}/logs | Insert a log entry for an analysis
[**ListAnalyses**](AnalysesCoreAPI.md#ListAnalyses) | **Get** /v2/analyses/list | Gets the most recent analyses
[**LookupBinaryId**](AnalysesCoreAPI.md#LookupBinaryId) | **Get** /v2/analyses/lookup/{binary_id} | Gets the analysis ID from binary ID
[**PutAnalysisStrings**](AnalysesCoreAPI.md#PutAnalysisStrings) | **Put** /v2/analyses/{analysis_id}/strings | Add strings to the analysis
[**RequeueAnalysis**](AnalysesCoreAPI.md#RequeueAnalysis) | **Post** /v2/analyses/{analysis_id}/requeue | Requeue Analysis
[**UpdateAnalysis**](AnalysesCoreAPI.md#UpdateAnalysis) | **Patch** /v2/analyses/{analysis_id} | Update Analysis
[**UpdateAnalysisTags**](AnalysesCoreAPI.md#UpdateAnalysisTags) | **Patch** /v2/analyses/{analysis_id}/tags | Update Analysis Tags
[**UploadFile**](AnalysesCoreAPI.md#UploadFile) | **Post** /v2/upload | Upload File



## CreateAnalysis

> BaseResponseAnalysisCreateResponse CreateAnalysis(ctx).AnalysisCreateRequest(analysisCreateRequest).EndpointUrl(endpointUrl).LocalCacheDir(localCacheDir).LocalCacheMaxSizeMb(localCacheMaxSizeMb).CustomerSamplesBucket(customerSamplesBucket).FirmwareSamplesBucket(firmwareSamplesBucket).MaxRetryAttempts(maxRetryAttempts).XRevEngApplication(xRevEngApplication).Execute()

Create Analysis



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
	analysisCreateRequest := *revengai.NewAnalysisCreateRequest("Filename_example", "Sha256Hash_example") // AnalysisCreateRequest | 
	endpointUrl := "endpointUrl_example" // string |  (optional)
	localCacheDir := "localCacheDir_example" // string |  (optional)
	localCacheMaxSizeMb := int32(56) // int32 |  (optional)
	customerSamplesBucket := "customerSamplesBucket_example" // string |  (optional)
	firmwareSamplesBucket := "firmwareSamplesBucket_example" // string |  (optional)
	maxRetryAttempts := int32(56) // int32 |  (optional) (default to 5)
	xRevEngApplication := "xRevEngApplication_example" // string |  (optional)

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysesCoreAPI.CreateAnalysis(context.Background()).AnalysisCreateRequest(analysisCreateRequest).EndpointUrl(endpointUrl).LocalCacheDir(localCacheDir).LocalCacheMaxSizeMb(localCacheMaxSizeMb).CustomerSamplesBucket(customerSamplesBucket).FirmwareSamplesBucket(firmwareSamplesBucket).MaxRetryAttempts(maxRetryAttempts).XRevEngApplication(xRevEngApplication).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysesCoreAPI.CreateAnalysis``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAnalysis`: BaseResponseAnalysisCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `AnalysesCoreAPI.CreateAnalysis`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAnalysisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **analysisCreateRequest** | [**AnalysisCreateRequest**](AnalysisCreateRequest.md) |  | 
 **endpointUrl** | **string** |  | 
 **localCacheDir** | **string** |  | 
 **localCacheMaxSizeMb** | **int32** |  | 
 **customerSamplesBucket** | **string** |  | 
 **firmwareSamplesBucket** | **string** |  | 
 **maxRetryAttempts** | **int32** |  | [default to 5]
 **xRevEngApplication** | **string** |  | 

### Return type

[**BaseResponseAnalysisCreateResponse**](BaseResponseAnalysisCreateResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAnalysis

> BaseResponseDict DeleteAnalysis(ctx, analysisId).EndpointUrl(endpointUrl).LocalCacheDir(localCacheDir).LocalCacheMaxSizeMb(localCacheMaxSizeMb).CustomerSamplesBucket(customerSamplesBucket).FirmwareSamplesBucket(firmwareSamplesBucket).MaxRetryAttempts(maxRetryAttempts).Execute()

Delete Analysis



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
	analysisId := int32(56) // int32 | 
	endpointUrl := "endpointUrl_example" // string |  (optional)
	localCacheDir := "localCacheDir_example" // string |  (optional)
	localCacheMaxSizeMb := int32(56) // int32 |  (optional)
	customerSamplesBucket := "customerSamplesBucket_example" // string |  (optional)
	firmwareSamplesBucket := "firmwareSamplesBucket_example" // string |  (optional)
	maxRetryAttempts := int32(56) // int32 |  (optional) (default to 5)

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysesCoreAPI.DeleteAnalysis(context.Background(), analysisId).EndpointUrl(endpointUrl).LocalCacheDir(localCacheDir).LocalCacheMaxSizeMb(localCacheMaxSizeMb).CustomerSamplesBucket(customerSamplesBucket).FirmwareSamplesBucket(firmwareSamplesBucket).MaxRetryAttempts(maxRetryAttempts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysesCoreAPI.DeleteAnalysis``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAnalysis`: BaseResponseDict
	fmt.Fprintf(os.Stdout, "Response from `AnalysesCoreAPI.DeleteAnalysis`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAnalysisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **endpointUrl** | **string** |  | 
 **localCacheDir** | **string** |  | 
 **localCacheMaxSizeMb** | **int32** |  | 
 **customerSamplesBucket** | **string** |  | 
 **firmwareSamplesBucket** | **string** |  | 
 **maxRetryAttempts** | **int32** |  | [default to 5]

### Return type

[**BaseResponseDict**](BaseResponseDict.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAnalysisBasicInfo

> BaseResponseBasic GetAnalysisBasicInfo(ctx, analysisId).Execute()

Gets basic analysis information



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
	analysisId := int32(56) // int32 | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysesCoreAPI.GetAnalysisBasicInfo(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysesCoreAPI.GetAnalysisBasicInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAnalysisBasicInfo`: BaseResponseBasic
	fmt.Fprintf(os.Stdout, "Response from `AnalysesCoreAPI.GetAnalysisBasicInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAnalysisBasicInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BaseResponseBasic**](BaseResponseBasic.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAnalysisFunctionMap

> BaseResponseAnalysisFunctionMapping GetAnalysisFunctionMap(ctx, analysisId).EndpointUrl(endpointUrl).LocalCacheDir(localCacheDir).LocalCacheMaxSizeMb(localCacheMaxSizeMb).CustomerSamplesBucket(customerSamplesBucket).FirmwareSamplesBucket(firmwareSamplesBucket).MaxRetryAttempts(maxRetryAttempts).Execute()

Get Analysis Function Map



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
	analysisId := int32(56) // int32 | 
	endpointUrl := "endpointUrl_example" // string |  (optional)
	localCacheDir := "localCacheDir_example" // string |  (optional)
	localCacheMaxSizeMb := int32(56) // int32 |  (optional)
	customerSamplesBucket := "customerSamplesBucket_example" // string |  (optional)
	firmwareSamplesBucket := "firmwareSamplesBucket_example" // string |  (optional)
	maxRetryAttempts := int32(56) // int32 |  (optional) (default to 5)

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysesCoreAPI.GetAnalysisFunctionMap(context.Background(), analysisId).EndpointUrl(endpointUrl).LocalCacheDir(localCacheDir).LocalCacheMaxSizeMb(localCacheMaxSizeMb).CustomerSamplesBucket(customerSamplesBucket).FirmwareSamplesBucket(firmwareSamplesBucket).MaxRetryAttempts(maxRetryAttempts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysesCoreAPI.GetAnalysisFunctionMap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAnalysisFunctionMap`: BaseResponseAnalysisFunctionMapping
	fmt.Fprintf(os.Stdout, "Response from `AnalysesCoreAPI.GetAnalysisFunctionMap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAnalysisFunctionMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **endpointUrl** | **string** |  | 
 **localCacheDir** | **string** |  | 
 **localCacheMaxSizeMb** | **int32** |  | 
 **customerSamplesBucket** | **string** |  | 
 **firmwareSamplesBucket** | **string** |  | 
 **maxRetryAttempts** | **int32** |  | [default to 5]

### Return type

[**BaseResponseAnalysisFunctionMapping**](BaseResponseAnalysisFunctionMapping.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAnalysisLogs

> BaseResponseLogs GetAnalysisLogs(ctx, analysisId).EndpointUrl(endpointUrl).LocalCacheDir(localCacheDir).LocalCacheMaxSizeMb(localCacheMaxSizeMb).CustomerSamplesBucket(customerSamplesBucket).FirmwareSamplesBucket(firmwareSamplesBucket).MaxRetryAttempts(maxRetryAttempts).Execute()

Gets the logs of an analysis



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
	analysisId := int32(56) // int32 | 
	endpointUrl := "endpointUrl_example" // string |  (optional)
	localCacheDir := "localCacheDir_example" // string |  (optional)
	localCacheMaxSizeMb := int32(56) // int32 |  (optional)
	customerSamplesBucket := "customerSamplesBucket_example" // string |  (optional)
	firmwareSamplesBucket := "firmwareSamplesBucket_example" // string |  (optional)
	maxRetryAttempts := int32(56) // int32 |  (optional) (default to 5)

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysesCoreAPI.GetAnalysisLogs(context.Background(), analysisId).EndpointUrl(endpointUrl).LocalCacheDir(localCacheDir).LocalCacheMaxSizeMb(localCacheMaxSizeMb).CustomerSamplesBucket(customerSamplesBucket).FirmwareSamplesBucket(firmwareSamplesBucket).MaxRetryAttempts(maxRetryAttempts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysesCoreAPI.GetAnalysisLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAnalysisLogs`: BaseResponseLogs
	fmt.Fprintf(os.Stdout, "Response from `AnalysesCoreAPI.GetAnalysisLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAnalysisLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **endpointUrl** | **string** |  | 
 **localCacheDir** | **string** |  | 
 **localCacheMaxSizeMb** | **int32** |  | 
 **customerSamplesBucket** | **string** |  | 
 **firmwareSamplesBucket** | **string** |  | 
 **maxRetryAttempts** | **int32** |  | [default to 5]

### Return type

[**BaseResponseLogs**](BaseResponseLogs.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAnalysisParams

> BaseResponseParams GetAnalysisParams(ctx, analysisId).EndpointUrl(endpointUrl).LocalCacheDir(localCacheDir).LocalCacheMaxSizeMb(localCacheMaxSizeMb).CustomerSamplesBucket(customerSamplesBucket).FirmwareSamplesBucket(firmwareSamplesBucket).MaxRetryAttempts(maxRetryAttempts).Execute()

Gets analysis param information



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
	analysisId := int32(56) // int32 | 
	endpointUrl := "endpointUrl_example" // string |  (optional)
	localCacheDir := "localCacheDir_example" // string |  (optional)
	localCacheMaxSizeMb := int32(56) // int32 |  (optional)
	customerSamplesBucket := "customerSamplesBucket_example" // string |  (optional)
	firmwareSamplesBucket := "firmwareSamplesBucket_example" // string |  (optional)
	maxRetryAttempts := int32(56) // int32 |  (optional) (default to 5)

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysesCoreAPI.GetAnalysisParams(context.Background(), analysisId).EndpointUrl(endpointUrl).LocalCacheDir(localCacheDir).LocalCacheMaxSizeMb(localCacheMaxSizeMb).CustomerSamplesBucket(customerSamplesBucket).FirmwareSamplesBucket(firmwareSamplesBucket).MaxRetryAttempts(maxRetryAttempts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysesCoreAPI.GetAnalysisParams``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAnalysisParams`: BaseResponseParams
	fmt.Fprintf(os.Stdout, "Response from `AnalysesCoreAPI.GetAnalysisParams`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAnalysisParamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **endpointUrl** | **string** |  | 
 **localCacheDir** | **string** |  | 
 **localCacheMaxSizeMb** | **int32** |  | 
 **customerSamplesBucket** | **string** |  | 
 **firmwareSamplesBucket** | **string** |  | 
 **maxRetryAttempts** | **int32** |  | [default to 5]

### Return type

[**BaseResponseParams**](BaseResponseParams.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAnalysisStatus

> BaseResponseStatus GetAnalysisStatus(ctx, analysisId).Execute()

Gets the status of an analysis



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
	analysisId := int32(56) // int32 | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysesCoreAPI.GetAnalysisStatus(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysesCoreAPI.GetAnalysisStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAnalysisStatus`: BaseResponseStatus
	fmt.Fprintf(os.Stdout, "Response from `AnalysesCoreAPI.GetAnalysisStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAnalysisStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BaseResponseStatus**](BaseResponseStatus.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InsertAnalysisLog

> BaseResponse InsertAnalysisLog(ctx, analysisId).InsertAnalysisLogRequest(insertAnalysisLogRequest).EndpointUrl(endpointUrl).LocalCacheDir(localCacheDir).LocalCacheMaxSizeMb(localCacheMaxSizeMb).CustomerSamplesBucket(customerSamplesBucket).FirmwareSamplesBucket(firmwareSamplesBucket).MaxRetryAttempts(maxRetryAttempts).Execute()

Insert a log entry for an analysis



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
	analysisId := int32(56) // int32 | 
	insertAnalysisLogRequest := *revengai.NewInsertAnalysisLogRequest("Log_example") // InsertAnalysisLogRequest | 
	endpointUrl := "endpointUrl_example" // string |  (optional)
	localCacheDir := "localCacheDir_example" // string |  (optional)
	localCacheMaxSizeMb := int32(56) // int32 |  (optional)
	customerSamplesBucket := "customerSamplesBucket_example" // string |  (optional)
	firmwareSamplesBucket := "firmwareSamplesBucket_example" // string |  (optional)
	maxRetryAttempts := int32(56) // int32 |  (optional) (default to 5)

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysesCoreAPI.InsertAnalysisLog(context.Background(), analysisId).InsertAnalysisLogRequest(insertAnalysisLogRequest).EndpointUrl(endpointUrl).LocalCacheDir(localCacheDir).LocalCacheMaxSizeMb(localCacheMaxSizeMb).CustomerSamplesBucket(customerSamplesBucket).FirmwareSamplesBucket(firmwareSamplesBucket).MaxRetryAttempts(maxRetryAttempts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysesCoreAPI.InsertAnalysisLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InsertAnalysisLog`: BaseResponse
	fmt.Fprintf(os.Stdout, "Response from `AnalysesCoreAPI.InsertAnalysisLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInsertAnalysisLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **insertAnalysisLogRequest** | [**InsertAnalysisLogRequest**](InsertAnalysisLogRequest.md) |  | 
 **endpointUrl** | **string** |  | 
 **localCacheDir** | **string** |  | 
 **localCacheMaxSizeMb** | **int32** |  | 
 **customerSamplesBucket** | **string** |  | 
 **firmwareSamplesBucket** | **string** |  | 
 **maxRetryAttempts** | **int32** |  | [default to 5]

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


## ListAnalyses

> BaseResponseRecent ListAnalyses(ctx).SearchTerm(searchTerm).Workspace(workspace).Status(status).ModelName(modelName).DynamicExecutionStatus(dynamicExecutionStatus).Usernames(usernames).Sha256Hash(sha256Hash).Limit(limit).Offset(offset).OrderBy(orderBy).Order(order).Execute()

Gets the most recent analyses



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
	searchTerm := "searchTerm_example" // string |  (optional) (default to "")
	workspace := []revengai.Workspace{revengai.Workspace("personal")} // []Workspace | The workspace to be viewed (optional) (default to {"personal"})
	status := []revengai.StatusInput{revengai.Status-Input("Uploaded")} // []StatusInput | The status of the analysis (optional) (default to {"All"})
	modelName := []revengai.ModelName{revengai.ModelName("binnet-0.7-x86-64-windows")} // []ModelName | Show analysis belonging to the model (optional)
	dynamicExecutionStatus := revengai.DynamicExecutionStatus("PENDING") // DynamicExecutionStatus | Show analysis that have a dynamic execution with the given status (optional)
	usernames := []*string{"Inner_example"} // []*string | Show analysis belonging to the user (optional) (default to {})
	sha256Hash := "sha256Hash_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 20)
	offset := int32(56) // int32 |  (optional) (default to 0)
	orderBy := revengai.app__api__rest__v2__analyses__enums__OrderBy("created") // AppApiRestV2AnalysesEnumsOrderBy |  (optional) (default to "created")
	order := revengai.Order("ASC") // Order |  (optional) (default to "DESC")

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysesCoreAPI.ListAnalyses(context.Background()).SearchTerm(searchTerm).Workspace(workspace).Status(status).ModelName(modelName).DynamicExecutionStatus(dynamicExecutionStatus).Usernames(usernames).Sha256Hash(sha256Hash).Limit(limit).Offset(offset).OrderBy(orderBy).Order(order).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysesCoreAPI.ListAnalyses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAnalyses`: BaseResponseRecent
	fmt.Fprintf(os.Stdout, "Response from `AnalysesCoreAPI.ListAnalyses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAnalysesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchTerm** | **string** |  | [default to &quot;&quot;]
 **workspace** | [**[]Workspace**](Workspace.md) | The workspace to be viewed | [default to {&quot;personal&quot;}]
 **status** | [**[]StatusInput**](StatusInput.md) | The status of the analysis | [default to {&quot;All&quot;}]
 **modelName** | [**[]ModelName**](ModelName.md) | Show analysis belonging to the model | 
 **dynamicExecutionStatus** | [**DynamicExecutionStatus**](DynamicExecutionStatus.md) | Show analysis that have a dynamic execution with the given status | 
 **usernames** | **[]string** | Show analysis belonging to the user | [default to {}]
 **sha256Hash** | **string** |  | 
 **limit** | **int32** |  | [default to 20]
 **offset** | **int32** |  | [default to 0]
 **orderBy** | [**AppApiRestV2AnalysesEnumsOrderBy**](AppApiRestV2AnalysesEnumsOrderBy.md) |  | [default to &quot;created&quot;]
 **order** | [**Order**](Order.md) |  | [default to &quot;DESC&quot;]

### Return type

[**BaseResponseRecent**](BaseResponseRecent.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LookupBinaryId

> interface{} LookupBinaryId(ctx, binaryId).Execute()

Gets the analysis ID from binary ID



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
	binaryId := int32(56) // int32 | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysesCoreAPI.LookupBinaryId(context.Background(), binaryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysesCoreAPI.LookupBinaryId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LookupBinaryId`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AnalysesCoreAPI.LookupBinaryId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**binaryId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLookupBinaryIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**interface{}**

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAnalysisStrings

> BaseResponse PutAnalysisStrings(ctx, analysisId).PutAnalysisStringsRequest(putAnalysisStringsRequest).Execute()

Add strings to the analysis



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
	analysisId := int32(56) // int32 | 
	putAnalysisStringsRequest := *revengai.NewPutAnalysisStringsRequest([]revengai.AnalysisStringInput{*revengai.NewAnalysisStringInput("Value_example", int32(123), revengai.StringSource("SYSTEM"))}) // PutAnalysisStringsRequest | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysesCoreAPI.PutAnalysisStrings(context.Background(), analysisId).PutAnalysisStringsRequest(putAnalysisStringsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysesCoreAPI.PutAnalysisStrings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutAnalysisStrings`: BaseResponse
	fmt.Fprintf(os.Stdout, "Response from `AnalysesCoreAPI.PutAnalysisStrings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutAnalysisStringsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putAnalysisStringsRequest** | [**PutAnalysisStringsRequest**](PutAnalysisStringsRequest.md) |  | 

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


## RequeueAnalysis

> BaseResponseCreated RequeueAnalysis(ctx, analysisId).ReAnalysisForm(reAnalysisForm).EndpointUrl(endpointUrl).LocalCacheDir(localCacheDir).LocalCacheMaxSizeMb(localCacheMaxSizeMb).CustomerSamplesBucket(customerSamplesBucket).FirmwareSamplesBucket(firmwareSamplesBucket).MaxRetryAttempts(maxRetryAttempts).XRevEngApplication(xRevEngApplication).Execute()

Requeue Analysis



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
	analysisId := int32(56) // int32 | 
	reAnalysisForm := *revengai.NewReAnalysisForm() // ReAnalysisForm | 
	endpointUrl := "endpointUrl_example" // string |  (optional)
	localCacheDir := "localCacheDir_example" // string |  (optional)
	localCacheMaxSizeMb := int32(56) // int32 |  (optional)
	customerSamplesBucket := "customerSamplesBucket_example" // string |  (optional)
	firmwareSamplesBucket := "firmwareSamplesBucket_example" // string |  (optional)
	maxRetryAttempts := int32(56) // int32 |  (optional) (default to 5)
	xRevEngApplication := "xRevEngApplication_example" // string |  (optional)

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysesCoreAPI.RequeueAnalysis(context.Background(), analysisId).ReAnalysisForm(reAnalysisForm).EndpointUrl(endpointUrl).LocalCacheDir(localCacheDir).LocalCacheMaxSizeMb(localCacheMaxSizeMb).CustomerSamplesBucket(customerSamplesBucket).FirmwareSamplesBucket(firmwareSamplesBucket).MaxRetryAttempts(maxRetryAttempts).XRevEngApplication(xRevEngApplication).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysesCoreAPI.RequeueAnalysis``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequeueAnalysis`: BaseResponseCreated
	fmt.Fprintf(os.Stdout, "Response from `AnalysesCoreAPI.RequeueAnalysis`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequeueAnalysisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reAnalysisForm** | [**ReAnalysisForm**](ReAnalysisForm.md) |  | 
 **endpointUrl** | **string** |  | 
 **localCacheDir** | **string** |  | 
 **localCacheMaxSizeMb** | **int32** |  | 
 **customerSamplesBucket** | **string** |  | 
 **firmwareSamplesBucket** | **string** |  | 
 **maxRetryAttempts** | **int32** |  | [default to 5]
 **xRevEngApplication** | **string** |  | 

### Return type

[**BaseResponseCreated**](BaseResponseCreated.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAnalysis

> BaseResponseAnalysisDetailResponse UpdateAnalysis(ctx, analysisId).AnalysisUpdateRequest(analysisUpdateRequest).EndpointUrl(endpointUrl).LocalCacheDir(localCacheDir).LocalCacheMaxSizeMb(localCacheMaxSizeMb).CustomerSamplesBucket(customerSamplesBucket).FirmwareSamplesBucket(firmwareSamplesBucket).MaxRetryAttempts(maxRetryAttempts).Execute()

Update Analysis



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
	analysisId := int32(56) // int32 | 
	analysisUpdateRequest := *revengai.NewAnalysisUpdateRequest() // AnalysisUpdateRequest | 
	endpointUrl := "endpointUrl_example" // string |  (optional)
	localCacheDir := "localCacheDir_example" // string |  (optional)
	localCacheMaxSizeMb := int32(56) // int32 |  (optional)
	customerSamplesBucket := "customerSamplesBucket_example" // string |  (optional)
	firmwareSamplesBucket := "firmwareSamplesBucket_example" // string |  (optional)
	maxRetryAttempts := int32(56) // int32 |  (optional) (default to 5)

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysesCoreAPI.UpdateAnalysis(context.Background(), analysisId).AnalysisUpdateRequest(analysisUpdateRequest).EndpointUrl(endpointUrl).LocalCacheDir(localCacheDir).LocalCacheMaxSizeMb(localCacheMaxSizeMb).CustomerSamplesBucket(customerSamplesBucket).FirmwareSamplesBucket(firmwareSamplesBucket).MaxRetryAttempts(maxRetryAttempts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysesCoreAPI.UpdateAnalysis``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAnalysis`: BaseResponseAnalysisDetailResponse
	fmt.Fprintf(os.Stdout, "Response from `AnalysesCoreAPI.UpdateAnalysis`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAnalysisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **analysisUpdateRequest** | [**AnalysisUpdateRequest**](AnalysisUpdateRequest.md) |  | 
 **endpointUrl** | **string** |  | 
 **localCacheDir** | **string** |  | 
 **localCacheMaxSizeMb** | **int32** |  | 
 **customerSamplesBucket** | **string** |  | 
 **firmwareSamplesBucket** | **string** |  | 
 **maxRetryAttempts** | **int32** |  | [default to 5]

### Return type

[**BaseResponseAnalysisDetailResponse**](BaseResponseAnalysisDetailResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAnalysisTags

> BaseResponseAnalysisUpdateTagsResponse UpdateAnalysisTags(ctx, analysisId).AnalysisUpdateTagsRequest(analysisUpdateTagsRequest).EndpointUrl(endpointUrl).LocalCacheDir(localCacheDir).LocalCacheMaxSizeMb(localCacheMaxSizeMb).CustomerSamplesBucket(customerSamplesBucket).FirmwareSamplesBucket(firmwareSamplesBucket).MaxRetryAttempts(maxRetryAttempts).Execute()

Update Analysis Tags



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
	analysisId := int32(56) // int32 | 
	analysisUpdateTagsRequest := *revengai.NewAnalysisUpdateTagsRequest([]string{"Tags_example"}) // AnalysisUpdateTagsRequest | 
	endpointUrl := "endpointUrl_example" // string |  (optional)
	localCacheDir := "localCacheDir_example" // string |  (optional)
	localCacheMaxSizeMb := int32(56) // int32 |  (optional)
	customerSamplesBucket := "customerSamplesBucket_example" // string |  (optional)
	firmwareSamplesBucket := "firmwareSamplesBucket_example" // string |  (optional)
	maxRetryAttempts := int32(56) // int32 |  (optional) (default to 5)

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysesCoreAPI.UpdateAnalysisTags(context.Background(), analysisId).AnalysisUpdateTagsRequest(analysisUpdateTagsRequest).EndpointUrl(endpointUrl).LocalCacheDir(localCacheDir).LocalCacheMaxSizeMb(localCacheMaxSizeMb).CustomerSamplesBucket(customerSamplesBucket).FirmwareSamplesBucket(firmwareSamplesBucket).MaxRetryAttempts(maxRetryAttempts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysesCoreAPI.UpdateAnalysisTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAnalysisTags`: BaseResponseAnalysisUpdateTagsResponse
	fmt.Fprintf(os.Stdout, "Response from `AnalysesCoreAPI.UpdateAnalysisTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAnalysisTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **analysisUpdateTagsRequest** | [**AnalysisUpdateTagsRequest**](AnalysisUpdateTagsRequest.md) |  | 
 **endpointUrl** | **string** |  | 
 **localCacheDir** | **string** |  | 
 **localCacheMaxSizeMb** | **int32** |  | 
 **customerSamplesBucket** | **string** |  | 
 **firmwareSamplesBucket** | **string** |  | 
 **maxRetryAttempts** | **int32** |  | [default to 5]

### Return type

[**BaseResponseAnalysisUpdateTagsResponse**](BaseResponseAnalysisUpdateTagsResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadFile

> BaseResponseUploadResponse UploadFile(ctx).UploadFileType(uploadFileType).File(file).PackedPassword(packedPassword).EndpointUrl(endpointUrl).LocalCacheDir(localCacheDir).LocalCacheMaxSizeMb(localCacheMaxSizeMb).CustomerSamplesBucket(customerSamplesBucket).FirmwareSamplesBucket(firmwareSamplesBucket).MaxRetryAttempts(maxRetryAttempts).ForceOverwrite(forceOverwrite).Execute()

Upload File

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
	uploadFileType := revengai.UploadFileType("BINARY") // UploadFileType | 
	file := "file_example" // string | 
	packedPassword := "packedPassword_example" // string |  (optional)
	endpointUrl := "endpointUrl_example" // string |  (optional)
	localCacheDir := "localCacheDir_example" // string |  (optional)
	localCacheMaxSizeMb := int32(56) // int32 |  (optional)
	customerSamplesBucket := "customerSamplesBucket_example" // string |  (optional)
	firmwareSamplesBucket := "firmwareSamplesBucket_example" // string |  (optional)
	maxRetryAttempts := int32(56) // int32 |  (optional) (default to 5)
	forceOverwrite := true // bool |  (optional) (default to false)

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysesCoreAPI.UploadFile(context.Background()).UploadFileType(uploadFileType).File(file).PackedPassword(packedPassword).EndpointUrl(endpointUrl).LocalCacheDir(localCacheDir).LocalCacheMaxSizeMb(localCacheMaxSizeMb).CustomerSamplesBucket(customerSamplesBucket).FirmwareSamplesBucket(firmwareSamplesBucket).MaxRetryAttempts(maxRetryAttempts).ForceOverwrite(forceOverwrite).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysesCoreAPI.UploadFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadFile`: BaseResponseUploadResponse
	fmt.Fprintf(os.Stdout, "Response from `AnalysesCoreAPI.UploadFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uploadFileType** | [**UploadFileType**](UploadFileType.md) |  | 
 **file** | **string** |  | 
 **packedPassword** | **string** |  | 
 **endpointUrl** | **string** |  | 
 **localCacheDir** | **string** |  | 
 **localCacheMaxSizeMb** | **int32** |  | 
 **customerSamplesBucket** | **string** |  | 
 **firmwareSamplesBucket** | **string** |  | 
 **maxRetryAttempts** | **int32** |  | [default to 5]
 **forceOverwrite** | **bool** |  | [default to false]

### Return type

[**BaseResponseUploadResponse**](BaseResponseUploadResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

