# AnalysesCoreAPI

All URIs are relative to *https://api.reveng.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddUserStringToAnalysis**](AnalysesCoreAPI.md#AddUserStringToAnalysis) | **Post** /v3/analyses/{analysis_id}/user-provided-strings | Add a user-provided string to an analysis.
[**CreateAnalysis**](AnalysesCoreAPI.md#CreateAnalysis) | **Post** /v2/analyses | Create Analysis
[**DeleteAnalysis**](AnalysesCoreAPI.md#DeleteAnalysis) | **Delete** /v2/analyses/{analysis_id} | Delete Analysis
[**GetAnalysisBasicInfo**](AnalysesCoreAPI.md#GetAnalysisBasicInfo) | **Get** /v2/analyses/{analysis_id}/basic | Gets basic analysis information
[**GetAnalysisBasicInfo_0**](AnalysesCoreAPI.md#GetAnalysisBasicInfo_0) | **Get** /v3/analyses/{analysis_id}/basic | Get basic analysis information
[**GetAnalysisBytes**](AnalysesCoreAPI.md#GetAnalysisBytes) | **Get** /v3/analyses/{analysis_id}/bytes | Get the bytes of a binary
[**GetAnalysisFunctionMap**](AnalysesCoreAPI.md#GetAnalysisFunctionMap) | **Get** /v2/analyses/{analysis_id}/func_maps | Get Analysis Function Map
[**GetAnalysisFunctionMatches**](AnalysesCoreAPI.md#GetAnalysisFunctionMatches) | **Get** /v3/analyses/{analysis_id}/functions/matches | Get function-matching results for an analysis
[**GetAnalysisFunctionMatchingStatus**](AnalysesCoreAPI.md#GetAnalysisFunctionMatchingStatus) | **Get** /v3/analyses/{analysis_id}/functions/matches/status | Get function-matching status for an analysis
[**GetAnalysisLogs**](AnalysesCoreAPI.md#GetAnalysisLogs) | **Get** /v2/analyses/{analysis_id}/logs | Gets the logs of an analysis
[**GetAnalysisParams**](AnalysesCoreAPI.md#GetAnalysisParams) | **Get** /v2/analyses/{analysis_id}/params | Gets analysis param information
[**GetAnalysisStatus**](AnalysesCoreAPI.md#GetAnalysisStatus) | **Get** /v2/analyses/{analysis_id}/status | Gets the status of an analysis
[**GetDynamicExecutionReport**](AnalysesCoreAPI.md#GetDynamicExecutionReport) | **Get** /v2/analyses/{analysis_id}/dynamic-execution/report | Get dynamic execution report
[**GetDynamicExecutionStatus**](AnalysesCoreAPI.md#GetDynamicExecutionStatus) | **Get** /v2/analyses/{analysis_id}/dynamic-execution/status | Get dynamic execution status
[**InsertAnalysisLog**](AnalysesCoreAPI.md#InsertAnalysisLog) | **Post** /v2/analyses/{analysis_id}/logs | Insert a log entry for an analysis
[**ListAnalyses**](AnalysesCoreAPI.md#ListAnalyses) | **Get** /v2/analyses/list | Gets the most recent analyses
[**LookupBinaryId**](AnalysesCoreAPI.md#LookupBinaryId) | **Get** /v2/analyses/lookup/{binary_id} | Gets the analysis ID from binary ID
[**PutAnalysisStrings**](AnalysesCoreAPI.md#PutAnalysisStrings) | **Put** /v2/analyses/{analysis_id}/strings | Add strings to the analysis
[**RequeueAnalysis**](AnalysesCoreAPI.md#RequeueAnalysis) | **Post** /v2/analyses/{analysis_id}/requeue | Requeue Analysis
[**StartAnalysisFunctionMatching**](AnalysesCoreAPI.md#StartAnalysisFunctionMatching) | **Post** /v3/analyses/{analysis_id}/functions/matches | Start function matching for an analysis
[**UpdateAnalysis**](AnalysesCoreAPI.md#UpdateAnalysis) | **Patch** /v2/analyses/{analysis_id} | Update Analysis
[**UpdateAnalysisTags**](AnalysesCoreAPI.md#UpdateAnalysisTags) | **Patch** /v2/analyses/{analysis_id}/tags | Update Analysis Tags
[**UploadFile**](AnalysesCoreAPI.md#UploadFile) | **Post** /v2/upload | Upload File
[**V3CreateAnalysis**](AnalysesCoreAPI.md#V3CreateAnalysis) | **Post** /v3/analyses | Create an analysis
[**V3DeleteAnalysis**](AnalysesCoreAPI.md#V3DeleteAnalysis) | **Delete** /v3/analyses/{analysis_id} | Delete an analysis.
[**V3DownloadBinaryExport**](AnalysesCoreAPI.md#V3DownloadBinaryExport) | **Get** /v3/analyses/{analysis_id}/binary-export | Download a binary export
[**V3GetAnalysis**](AnalysesCoreAPI.md#V3GetAnalysis) | **Get** /v3/analyses/{analysis_id} | Get an analysis.
[**V3GetAnalysisAutoUnstripStatus**](AnalysesCoreAPI.md#V3GetAnalysisAutoUnstripStatus) | **Get** /v3/analyses/{analysis_id}/auto-unstrip/status | Get the auto-unstrip status for an analysis.
[**V3GetAnalysisFunctionsProgress**](AnalysesCoreAPI.md#V3GetAnalysisFunctionsProgress) | **Get** /v3/analyses/{analysis_id}/progress/functions | Get function embedding progress for an analysis.
[**V3GetAnalysisLogs**](AnalysesCoreAPI.md#V3GetAnalysisLogs) | **Get** /v3/analyses/{analysis_id}/logs | Get the Analysis log
[**V3GetAnalysisOperation**](AnalysesCoreAPI.md#V3GetAnalysisOperation) | **Get** /v3/operations/analyses/{analysis_id} | Get an Analysis-creation operation
[**V3GetAnalysisStrings**](AnalysesCoreAPI.md#V3GetAnalysisStrings) | **Get** /v3/analyses/{analysis_id}/functions/strings | List strings for an analysis.
[**V3GetAnalysisStringsStatus**](AnalysesCoreAPI.md#V3GetAnalysisStringsStatus) | **Get** /v3/analyses/{analysis_id}/functions/strings/status | Get the string-extraction status for an analysis.
[**V3GetBinaryExportOperation**](AnalysesCoreAPI.md#V3GetBinaryExportOperation) | **Get** /v3/operations/binary-export/{task_id} | Get a binary export operation
[**V3ListAnalyses**](AnalysesCoreAPI.md#V3ListAnalyses) | **Get** /v3/analyses | List analyses
[**V3ListExampleAnalyses**](AnalysesCoreAPI.md#V3ListExampleAnalyses) | **Get** /v3/analyses/examples | List example analyses
[**V3LookupAnalysisByBinaryId**](AnalysesCoreAPI.md#V3LookupAnalysisByBinaryId) | **Get** /v3/analyses/lookup/{binary_id} | Look up the most recent analysis for a binary.
[**V3QueueBinaryExport**](AnalysesCoreAPI.md#V3QueueBinaryExport) | **Post** /v3/analyses/{analysis_id}/binary-export | Queue a binary export
[**V3SearchTags**](AnalysesCoreAPI.md#V3SearchTags) | **Get** /v3/tags | Search tags
[**V3UpdateAnalysis**](AnalysesCoreAPI.md#V3UpdateAnalysis) | **Patch** /v3/analyses/{analysis_id} | Update an analysis.
[**V3UpdateAnalysisTags**](AnalysesCoreAPI.md#V3UpdateAnalysisTags) | **Patch** /v3/analyses/{analysis_id}/tags | Replace an analysis&#39; tags.
[**V3UpgradeAnalysisModel**](AnalysesCoreAPI.md#V3UpgradeAnalysisModel) | **Post** /v3/analyses/{analysis_id}/upgrade-model | Re-analyse on the latest model



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
	revengai "github.com/RevEngAI/sdk-go/v4"
)

func main() {
	analysisId := int64(789) // int64 | Analysis ID
	addUserStringInputBody := *revengai.NewAddUserStringInputBody("String_example", int64(123)) // AddUserStringInputBody | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysesCoreAPI.AddUserStringToAnalysis(context.Background(), analysisId).AddUserStringInputBody(addUserStringInputBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysesCoreAPI.AddUserStringToAnalysis``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddUserStringToAnalysis`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AnalysesCoreAPI.AddUserStringToAnalysis`: %v\n", resp)
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

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAnalysis

> BaseResponseAnalysisCreateResponse CreateAnalysis(ctx).AnalysisCreateRequest(analysisCreateRequest).XRevEngApplication(xRevEngApplication).Execute()

Create Analysis



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
	analysisCreateRequest := *revengai.NewAnalysisCreateRequest("Filename_example", "Sha256Hash_example") // AnalysisCreateRequest | 
	xRevEngApplication := "xRevEngApplication_example" // string |  (optional)

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysesCoreAPI.CreateAnalysis(context.Background()).AnalysisCreateRequest(analysisCreateRequest).XRevEngApplication(xRevEngApplication).Execute()
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
 **xRevEngApplication** | **string** |  | 

### Return type

[**BaseResponseAnalysisCreateResponse**](BaseResponseAnalysisCreateResponse.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAnalysis

> BaseResponseDict DeleteAnalysis(ctx, analysisId).Execute()

Delete Analysis



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
	analysisId := int32(56) // int32 | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysesCoreAPI.DeleteAnalysis(context.Background(), analysisId).Execute()
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


### Return type

[**BaseResponseDict**](BaseResponseDict.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

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
	revengai "github.com/RevEngAI/sdk-go/v4"
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

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAnalysisBasicInfo_0

> AnalysisBasicInfoOutputBody GetAnalysisBasicInfo_0(ctx, analysisId).Execute()

Get basic analysis information



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
	analysisId := int64(789) // int64 | Analysis ID

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysesCoreAPI.GetAnalysisBasicInfo_0(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysesCoreAPI.GetAnalysisBasicInfo_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAnalysisBasicInfo_0`: AnalysisBasicInfoOutputBody
	fmt.Fprintf(os.Stdout, "Response from `AnalysesCoreAPI.GetAnalysisBasicInfo_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int64** | Analysis ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAnalysisBasicInfo_1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AnalysisBasicInfoOutputBody**](AnalysisBasicInfoOutputBody.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAnalysisBytes

> GetAnalysisBytes(ctx, analysisId).Page(page).Execute()

Get the bytes of a binary



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
	analysisId := int64(789) // int64 | Analysis ID
	page := int64(789) // int64 | 64kb page of binary data (optional)

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	r, err := apiClient.AnalysesCoreAPI.GetAnalysisBytes(context.Background(), analysisId).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysesCoreAPI.GetAnalysisBytes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int64** | Analysis ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAnalysisBytesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int64** | 64kb page of binary data | 

### Return type

 (empty response body)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAnalysisFunctionMap

> BaseResponseAnalysisFunctionMapping GetAnalysisFunctionMap(ctx, analysisId).Execute()

Get Analysis Function Map



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
	analysisId := int32(56) // int32 | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysesCoreAPI.GetAnalysisFunctionMap(context.Background(), analysisId).Execute()
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


### Return type

[**BaseResponseAnalysisFunctionMapping**](BaseResponseAnalysisFunctionMapping.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAnalysisFunctionMatches

> GetMatchesOutputBody GetAnalysisFunctionMatches(ctx, analysisId).MatchId(matchId).Execute()

Get function-matching results for an analysis



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
	analysisId := int64(789) // int64 | Analysis ID
	matchId := "matchId_example" // string | Opaque token from a start-matching response. When supplied, returns that specific run instead of the latest. (optional)

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysesCoreAPI.GetAnalysisFunctionMatches(context.Background(), analysisId).MatchId(matchId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysesCoreAPI.GetAnalysisFunctionMatches``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAnalysisFunctionMatches`: GetMatchesOutputBody
	fmt.Fprintf(os.Stdout, "Response from `AnalysesCoreAPI.GetAnalysisFunctionMatches`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int64** | Analysis ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAnalysisFunctionMatchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **matchId** | **string** | Opaque token from a start-matching response. When supplied, returns that specific run instead of the latest. | 

### Return type

[**GetMatchesOutputBody**](GetMatchesOutputBody.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAnalysisFunctionMatchingStatus

> GetMatchesStatusOutputBody GetAnalysisFunctionMatchingStatus(ctx, analysisId).MatchId(matchId).Execute()

Get function-matching status for an analysis



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
	analysisId := int64(789) // int64 | Analysis ID
	matchId := "matchId_example" // string | Opaque token from a start-matching response. When supplied, returns that specific run instead of the latest. (optional)

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysesCoreAPI.GetAnalysisFunctionMatchingStatus(context.Background(), analysisId).MatchId(matchId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysesCoreAPI.GetAnalysisFunctionMatchingStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAnalysisFunctionMatchingStatus`: GetMatchesStatusOutputBody
	fmt.Fprintf(os.Stdout, "Response from `AnalysesCoreAPI.GetAnalysisFunctionMatchingStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int64** | Analysis ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAnalysisFunctionMatchingStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **matchId** | **string** | Opaque token from a start-matching response. When supplied, returns that specific run instead of the latest. | 

### Return type

[**GetMatchesStatusOutputBody**](GetMatchesStatusOutputBody.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAnalysisLogs

> BaseResponseLogs GetAnalysisLogs(ctx, analysisId).Execute()

Gets the logs of an analysis



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
	analysisId := int32(56) // int32 | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysesCoreAPI.GetAnalysisLogs(context.Background(), analysisId).Execute()
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


### Return type

[**BaseResponseLogs**](BaseResponseLogs.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAnalysisParams

> BaseResponseParams GetAnalysisParams(ctx, analysisId).Execute()

Gets analysis param information



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
	analysisId := int32(56) // int32 | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysesCoreAPI.GetAnalysisParams(context.Background(), analysisId).Execute()
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


### Return type

[**BaseResponseParams**](BaseResponseParams.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

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
	revengai "github.com/RevEngAI/sdk-go/v4"
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

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDynamicExecutionReport

> AnalysisReport GetDynamicExecutionReport(ctx, analysisId).Execute()

Get dynamic execution report



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
	analysisId := int64(789) // int64 | Analysis ID

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysesCoreAPI.GetDynamicExecutionReport(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysesCoreAPI.GetDynamicExecutionReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDynamicExecutionReport`: AnalysisReport
	fmt.Fprintf(os.Stdout, "Response from `AnalysesCoreAPI.GetDynamicExecutionReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int64** | Analysis ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDynamicExecutionReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AnalysisReport**](AnalysisReport.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDynamicExecutionStatus

> DynamicExecutionStatusResponse GetDynamicExecutionStatus(ctx, analysisId).Execute()

Get dynamic execution status



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
	analysisId := int64(789) // int64 | Analysis ID

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysesCoreAPI.GetDynamicExecutionStatus(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysesCoreAPI.GetDynamicExecutionStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDynamicExecutionStatus`: DynamicExecutionStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `AnalysesCoreAPI.GetDynamicExecutionStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int64** | Analysis ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDynamicExecutionStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DynamicExecutionStatusResponse**](DynamicExecutionStatusResponse.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InsertAnalysisLog

> BaseResponse InsertAnalysisLog(ctx, analysisId).InsertAnalysisLogRequest(insertAnalysisLogRequest).Execute()

Insert a log entry for an analysis



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
	analysisId := int32(56) // int32 | 
	insertAnalysisLogRequest := *revengai.NewInsertAnalysisLogRequest("Log_example") // InsertAnalysisLogRequest | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysesCoreAPI.InsertAnalysisLog(context.Background(), analysisId).InsertAnalysisLogRequest(insertAnalysisLogRequest).Execute()
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
	revengai "github.com/RevEngAI/sdk-go/v4"
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

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

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
	revengai "github.com/RevEngAI/sdk-go/v4"
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

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

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
	revengai "github.com/RevEngAI/sdk-go/v4"
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

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequeueAnalysis

> BaseResponseCreated RequeueAnalysis(ctx, analysisId).ReAnalysisForm(reAnalysisForm).XRevEngApplication(xRevEngApplication).Execute()

Requeue Analysis



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
	analysisId := int32(56) // int32 | 
	reAnalysisForm := *revengai.NewReAnalysisForm() // ReAnalysisForm | 
	xRevEngApplication := "xRevEngApplication_example" // string |  (optional)

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysesCoreAPI.RequeueAnalysis(context.Background(), analysisId).ReAnalysisForm(reAnalysisForm).XRevEngApplication(xRevEngApplication).Execute()
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
 **xRevEngApplication** | **string** |  | 

### Return type

[**BaseResponseCreated**](BaseResponseCreated.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartAnalysisFunctionMatching

> StartMatchingOutputBody StartAnalysisFunctionMatching(ctx, analysisId).StartMatchingForAnalysisInputBody(startMatchingForAnalysisInputBody).Execute()

Start function matching for an analysis



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
	analysisId := int64(789) // int64 | Analysis ID
	startMatchingForAnalysisInputBody := *revengai.NewStartMatchingForAnalysisInputBody() // StartMatchingForAnalysisInputBody | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysesCoreAPI.StartAnalysisFunctionMatching(context.Background(), analysisId).StartMatchingForAnalysisInputBody(startMatchingForAnalysisInputBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysesCoreAPI.StartAnalysisFunctionMatching``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartAnalysisFunctionMatching`: StartMatchingOutputBody
	fmt.Fprintf(os.Stdout, "Response from `AnalysesCoreAPI.StartAnalysisFunctionMatching`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int64** | Analysis ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartAnalysisFunctionMatchingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startMatchingForAnalysisInputBody** | [**StartMatchingForAnalysisInputBody**](StartMatchingForAnalysisInputBody.md) |  | 

### Return type

[**StartMatchingOutputBody**](StartMatchingOutputBody.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAnalysis

> BaseResponseAnalysisDetailResponse UpdateAnalysis(ctx, analysisId).AnalysisUpdateRequest(analysisUpdateRequest).Execute()

Update Analysis



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
	analysisId := int32(56) // int32 | 
	analysisUpdateRequest := *revengai.NewAnalysisUpdateRequest() // AnalysisUpdateRequest | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysesCoreAPI.UpdateAnalysis(context.Background(), analysisId).AnalysisUpdateRequest(analysisUpdateRequest).Execute()
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

### Return type

[**BaseResponseAnalysisDetailResponse**](BaseResponseAnalysisDetailResponse.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAnalysisTags

> BaseResponseAnalysisUpdateTagsResponse UpdateAnalysisTags(ctx, analysisId).AnalysisUpdateTagsRequest(analysisUpdateTagsRequest).Execute()

Update Analysis Tags



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
	analysisId := int32(56) // int32 | 
	analysisUpdateTagsRequest := *revengai.NewAnalysisUpdateTagsRequest([]string{"Tags_example"}) // AnalysisUpdateTagsRequest | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysesCoreAPI.UpdateAnalysisTags(context.Background(), analysisId).AnalysisUpdateTagsRequest(analysisUpdateTagsRequest).Execute()
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

### Return type

[**BaseResponseAnalysisUpdateTagsResponse**](BaseResponseAnalysisUpdateTagsResponse.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadFile

> BaseResponseUploadResponse UploadFile(ctx).UploadFileType(uploadFileType).File(file).PackedPassword(packedPassword).ForceOverwrite(forceOverwrite).Execute()

Upload File

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
	uploadFileType := revengai.UploadFileType("BINARY") // UploadFileType | 
	file := os.NewFile(1234, "some_file") // *os.File | 
	packedPassword := "packedPassword_example" // string |  (optional)
	forceOverwrite := true // bool |  (optional) (default to false)

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysesCoreAPI.UploadFile(context.Background()).UploadFileType(uploadFileType).File(file).PackedPassword(packedPassword).ForceOverwrite(forceOverwrite).Execute()
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
 **file** | ***os.File** |  | 
 **packedPassword** | **string** |  | 
 **forceOverwrite** | **bool** |  | [default to false]

### Return type

[**BaseResponseUploadResponse**](BaseResponseUploadResponse.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3CreateAnalysis

> OperationCreateMetadataCreateResult V3CreateAnalysis(ctx).CreateRequest(createRequest).XRevEngApplication(xRevEngApplication).Execute()

Create an analysis



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
	createRequest := *revengai.NewCreateRequest("Filename_example", "Sha256Hash_example") // CreateRequest | 
	xRevEngApplication := "xRevEngApplication_example" // string | Identifies the calling RevEng application. Recorded on the Analysis log. (optional)

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysesCoreAPI.V3CreateAnalysis(context.Background()).CreateRequest(createRequest).XRevEngApplication(xRevEngApplication).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysesCoreAPI.V3CreateAnalysis``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3CreateAnalysis`: OperationCreateMetadataCreateResult
	fmt.Fprintf(os.Stdout, "Response from `AnalysesCoreAPI.V3CreateAnalysis`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3CreateAnalysisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRequest** | [**CreateRequest**](CreateRequest.md) |  | 
 **xRevEngApplication** | **string** | Identifies the calling RevEng application. Recorded on the Analysis log. | 

### Return type

[**OperationCreateMetadataCreateResult**](OperationCreateMetadataCreateResult.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3DeleteAnalysis

> V3DeleteAnalysis(ctx, analysisId).Execute()

Delete an analysis.



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
	analysisId := int64(789) // int64 | Analysis ID

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	r, err := apiClient.AnalysesCoreAPI.V3DeleteAnalysis(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysesCoreAPI.V3DeleteAnalysis``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int64** | Analysis ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3DeleteAnalysisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3DownloadBinaryExport

> V3DownloadBinaryExport(ctx, analysisId).TaskId(taskId).Execute()

Download a binary export



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
	analysisId := int64(789) // int64 | Analysis ID
	taskId := "taskId_example" // string | Task ID returned by queueing the export (optional)

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	r, err := apiClient.AnalysesCoreAPI.V3DownloadBinaryExport(context.Background(), analysisId).TaskId(taskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysesCoreAPI.V3DownloadBinaryExport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int64** | Analysis ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3DownloadBinaryExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **taskId** | **string** | Task ID returned by queueing the export | 

### Return type

 (empty response body)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3GetAnalysis

> AnalysisDetailOutputBody V3GetAnalysis(ctx, analysisId).Execute()

Get an analysis.



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
	analysisId := int64(789) // int64 | Analysis ID

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysesCoreAPI.V3GetAnalysis(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysesCoreAPI.V3GetAnalysis``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3GetAnalysis`: AnalysisDetailOutputBody
	fmt.Fprintf(os.Stdout, "Response from `AnalysesCoreAPI.V3GetAnalysis`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int64** | Analysis ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3GetAnalysisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AnalysisDetailOutputBody**](AnalysisDetailOutputBody.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3GetAnalysisAutoUnstripStatus

> AutoUnstripStatusOutputBody V3GetAnalysisAutoUnstripStatus(ctx, analysisId).Execute()

Get the auto-unstrip status for an analysis.



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
	analysisId := int64(789) // int64 | Analysis ID

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysesCoreAPI.V3GetAnalysisAutoUnstripStatus(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysesCoreAPI.V3GetAnalysisAutoUnstripStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3GetAnalysisAutoUnstripStatus`: AutoUnstripStatusOutputBody
	fmt.Fprintf(os.Stdout, "Response from `AnalysesCoreAPI.V3GetAnalysisAutoUnstripStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int64** | Analysis ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3GetAnalysisAutoUnstripStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AutoUnstripStatusOutputBody**](AutoUnstripStatusOutputBody.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3GetAnalysisFunctionsProgress

> FunctionsProgressOutputBody V3GetAnalysisFunctionsProgress(ctx, analysisId).Execute()

Get function embedding progress for an analysis.



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
	analysisId := int64(789) // int64 | Analysis ID

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysesCoreAPI.V3GetAnalysisFunctionsProgress(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysesCoreAPI.V3GetAnalysisFunctionsProgress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3GetAnalysisFunctionsProgress`: FunctionsProgressOutputBody
	fmt.Fprintf(os.Stdout, "Response from `AnalysesCoreAPI.V3GetAnalysisFunctionsProgress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int64** | Analysis ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3GetAnalysisFunctionsProgressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FunctionsProgressOutputBody**](FunctionsProgressOutputBody.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3GetAnalysisLogs

> GetAnalysisLogsOutputBody V3GetAnalysisLogs(ctx, analysisId).Execute()

Get the Analysis log



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
	analysisId := int64(789) // int64 | Analysis ID

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysesCoreAPI.V3GetAnalysisLogs(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysesCoreAPI.V3GetAnalysisLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3GetAnalysisLogs`: GetAnalysisLogsOutputBody
	fmt.Fprintf(os.Stdout, "Response from `AnalysesCoreAPI.V3GetAnalysisLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int64** | Analysis ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3GetAnalysisLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetAnalysisLogsOutputBody**](GetAnalysisLogsOutputBody.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3GetAnalysisOperation

> OperationCreateMetadataCreateResult V3GetAnalysisOperation(ctx, analysisId).Execute()

Get an Analysis-creation operation



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
	analysisId := int64(789) // int64 | Analysis ID

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysesCoreAPI.V3GetAnalysisOperation(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysesCoreAPI.V3GetAnalysisOperation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3GetAnalysisOperation`: OperationCreateMetadataCreateResult
	fmt.Fprintf(os.Stdout, "Response from `AnalysesCoreAPI.V3GetAnalysisOperation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int64** | Analysis ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3GetAnalysisOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationCreateMetadataCreateResult**](OperationCreateMetadataCreateResult.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3GetAnalysisStrings

> ListAnalysisStringsOutputBody V3GetAnalysisStrings(ctx, analysisId).Page(page).PageSize(pageSize).Search(search).SearchOperator(searchOperator).FunctionSearch(functionSearch).OrderBy(orderBy).SortOrder(sortOrder).Execute()

List strings for an analysis.



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
	analysisId := int64(789) // int64 | Analysis ID
	page := int64(789) // int64 | Page number (1-indexed). (optional) (default to 1)
	pageSize := int64(789) // int64 | Number of results per page. (optional) (default to 100)
	search := "search_example" // string | Filter by string value (case-insensitive substring match). (optional)
	searchOperator := "searchOperator_example" // string | How the search term matches string values. (optional) (default to "CONTAINS")
	functionSearch := "functionSearch_example" // string | Filter by function name (case-insensitive substring match). (optional)
	orderBy := "orderBy_example" // string | Field to order results by. (optional) (default to "value")
	sortOrder := "sortOrder_example" // string | Sort direction. (optional) (default to "ASC")

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysesCoreAPI.V3GetAnalysisStrings(context.Background(), analysisId).Page(page).PageSize(pageSize).Search(search).SearchOperator(searchOperator).FunctionSearch(functionSearch).OrderBy(orderBy).SortOrder(sortOrder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysesCoreAPI.V3GetAnalysisStrings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3GetAnalysisStrings`: ListAnalysisStringsOutputBody
	fmt.Fprintf(os.Stdout, "Response from `AnalysesCoreAPI.V3GetAnalysisStrings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int64** | Analysis ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3GetAnalysisStringsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int64** | Page number (1-indexed). | [default to 1]
 **pageSize** | **int64** | Number of results per page. | [default to 100]
 **search** | **string** | Filter by string value (case-insensitive substring match). | 
 **searchOperator** | **string** | How the search term matches string values. | [default to &quot;CONTAINS&quot;]
 **functionSearch** | **string** | Filter by function name (case-insensitive substring match). | 
 **orderBy** | **string** | Field to order results by. | [default to &quot;value&quot;]
 **sortOrder** | **string** | Sort direction. | [default to &quot;ASC&quot;]

### Return type

[**ListAnalysisStringsOutputBody**](ListAnalysisStringsOutputBody.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3GetAnalysisStringsStatus

> GetAnalysisStringsStatusOutputBody V3GetAnalysisStringsStatus(ctx, analysisId).Execute()

Get the string-extraction status for an analysis.



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
	analysisId := int64(789) // int64 | Analysis ID

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysesCoreAPI.V3GetAnalysisStringsStatus(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysesCoreAPI.V3GetAnalysisStringsStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3GetAnalysisStringsStatus`: GetAnalysisStringsStatusOutputBody
	fmt.Fprintf(os.Stdout, "Response from `AnalysesCoreAPI.V3GetAnalysisStringsStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int64** | Analysis ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3GetAnalysisStringsStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetAnalysisStringsStatusOutputBody**](GetAnalysisStringsStatusOutputBody.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3GetBinaryExportOperation

> OperationBinaryExportMetadataBinaryExportResult V3GetBinaryExportOperation(ctx, taskId).Execute()

Get a binary export operation



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
	taskId := "taskId_example" // string | Task ID returned by queueing the export

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysesCoreAPI.V3GetBinaryExportOperation(context.Background(), taskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysesCoreAPI.V3GetBinaryExportOperation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3GetBinaryExportOperation`: OperationBinaryExportMetadataBinaryExportResult
	fmt.Fprintf(os.Stdout, "Response from `AnalysesCoreAPI.V3GetBinaryExportOperation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** | Task ID returned by queueing the export | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3GetBinaryExportOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationBinaryExportMetadataBinaryExportResult**](OperationBinaryExportMetadataBinaryExportResult.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ListAnalyses

> ListAnalysesOutputBody V3ListAnalyses(ctx).SearchTerm(searchTerm).AnalysisScope(analysisScope).Status(status).ModelName(modelName).Usernames(usernames).Sha256Hash(sha256Hash).BinaryId(binaryId).Platform(platform).Architecture(architecture).PageSize(pageSize).NextPageToken(nextPageToken).OrderBy(orderBy).Order(order).Execute()

List analyses



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
	searchTerm := "searchTerm_example" // string |  (optional)
	analysisScope := []string{"AnalysisScope_example"} // []string | Leave empty to search your own, your team's and all public analyses (optional)
	status := []string{"Status_example"} // []string |  (optional)
	modelName := []*string{"Inner_example"} // []*string |  (optional)
	usernames := []*string{"Inner_example"} // []*string |  (optional)
	sha256Hash := "sha256Hash_example" // string |  (optional)
	binaryId := int64(789) // int64 | Restrict to analyses of this binary. A binary can carry more than one analysis; they are returned newest first under the default sort (optional)
	platform := []string{"Platform_example"} // []string | Restrict to binaries running on one of these operating-system platforms. Matches the uploader's override when they set one, the detected platform otherwise; a binary with neither is never matched. Leave empty for no filter (optional)
	architecture := []string{"Architecture_example"} // []string | Restrict to binaries built for one of these instruction-set architectures. Resolved the same way as platform. Leave empty for no filter (optional)
	pageSize := int64(789) // int64 |  (optional) (default to 20)
	nextPageToken := "nextPageToken_example" // string | Forward-pagination cursor from a prior response. When set, order_by/order are taken from the token (the sort cannot change mid-pagination). (optional)
	orderBy := "orderBy_example" // string |  (optional) (default to "created")
	order := "order_example" // string |  (optional) (default to "DESC")

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysesCoreAPI.V3ListAnalyses(context.Background()).SearchTerm(searchTerm).AnalysisScope(analysisScope).Status(status).ModelName(modelName).Usernames(usernames).Sha256Hash(sha256Hash).BinaryId(binaryId).Platform(platform).Architecture(architecture).PageSize(pageSize).NextPageToken(nextPageToken).OrderBy(orderBy).Order(order).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysesCoreAPI.V3ListAnalyses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3ListAnalyses`: ListAnalysesOutputBody
	fmt.Fprintf(os.Stdout, "Response from `AnalysesCoreAPI.V3ListAnalyses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3ListAnalysesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchTerm** | **string** |  | 
 **analysisScope** | **[]string** | Leave empty to search your own, your team&#39;s and all public analyses | 
 **status** | **[]string** |  | 
 **modelName** | **[]string** |  | 
 **usernames** | **[]string** |  | 
 **sha256Hash** | **string** |  | 
 **binaryId** | **int64** | Restrict to analyses of this binary. A binary can carry more than one analysis; they are returned newest first under the default sort | 
 **platform** | **[]string** | Restrict to binaries running on one of these operating-system platforms. Matches the uploader&#39;s override when they set one, the detected platform otherwise; a binary with neither is never matched. Leave empty for no filter | 
 **architecture** | **[]string** | Restrict to binaries built for one of these instruction-set architectures. Resolved the same way as platform. Leave empty for no filter | 
 **pageSize** | **int64** |  | [default to 20]
 **nextPageToken** | **string** | Forward-pagination cursor from a prior response. When set, order_by/order are taken from the token (the sort cannot change mid-pagination). | 
 **orderBy** | **string** |  | [default to &quot;created&quot;]
 **order** | **string** |  | [default to &quot;DESC&quot;]

### Return type

[**ListAnalysesOutputBody**](ListAnalysesOutputBody.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3ListExampleAnalyses

> ListExampleAnalysesOutputBody V3ListExampleAnalyses(ctx).Execute()

List example analyses



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
	resp, r, err := apiClient.AnalysesCoreAPI.V3ListExampleAnalyses(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysesCoreAPI.V3ListExampleAnalyses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3ListExampleAnalyses`: ListExampleAnalysesOutputBody
	fmt.Fprintf(os.Stdout, "Response from `AnalysesCoreAPI.V3ListExampleAnalyses`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV3ListExampleAnalysesRequest struct via the builder pattern


### Return type

[**ListExampleAnalysesOutputBody**](ListExampleAnalysesOutputBody.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3LookupAnalysisByBinaryId

> LookupAnalysisByBinaryIDOutputBody V3LookupAnalysisByBinaryId(ctx, binaryId).Execute()

Look up the most recent analysis for a binary.



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
	binaryId := int64(789) // int64 | Binary ID

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysesCoreAPI.V3LookupAnalysisByBinaryId(context.Background(), binaryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysesCoreAPI.V3LookupAnalysisByBinaryId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3LookupAnalysisByBinaryId`: LookupAnalysisByBinaryIDOutputBody
	fmt.Fprintf(os.Stdout, "Response from `AnalysesCoreAPI.V3LookupAnalysisByBinaryId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**binaryId** | **int64** | Binary ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3LookupAnalysisByBinaryIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LookupAnalysisByBinaryIDOutputBody**](LookupAnalysisByBinaryIDOutputBody.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3QueueBinaryExport

> OperationBinaryExportMetadataBinaryExportResult V3QueueBinaryExport(ctx, analysisId).Execute()

Queue a binary export



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
	analysisId := int64(789) // int64 | Analysis ID

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysesCoreAPI.V3QueueBinaryExport(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysesCoreAPI.V3QueueBinaryExport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3QueueBinaryExport`: OperationBinaryExportMetadataBinaryExportResult
	fmt.Fprintf(os.Stdout, "Response from `AnalysesCoreAPI.V3QueueBinaryExport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int64** | Analysis ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3QueueBinaryExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationBinaryExportMetadataBinaryExportResult**](OperationBinaryExportMetadataBinaryExportResult.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3SearchTags

> SearchTagsOutputBody V3SearchTags(ctx).PartialName(partialName).Limit(limit).Offset(offset).Execute()

Search tags



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
	partialName := "partialName_example" // string | Partial or full tag name to search for, at least 3 characters (optional)
	limit := int64(789) // int64 | Maximum results to return (optional) (default to 10)
	offset := int64(789) // int64 | Number of results to skip (optional) (default to 0)

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysesCoreAPI.V3SearchTags(context.Background()).PartialName(partialName).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysesCoreAPI.V3SearchTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3SearchTags`: SearchTagsOutputBody
	fmt.Fprintf(os.Stdout, "Response from `AnalysesCoreAPI.V3SearchTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3SearchTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **partialName** | **string** | Partial or full tag name to search for, at least 3 characters | 
 **limit** | **int64** | Maximum results to return | [default to 10]
 **offset** | **int64** | Number of results to skip | [default to 0]

### Return type

[**SearchTagsOutputBody**](SearchTagsOutputBody.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3UpdateAnalysis

> AnalysisDetailOutputBody V3UpdateAnalysis(ctx, analysisId).UpdateAnalysisInputBody(updateAnalysisInputBody).Execute()

Update an analysis.



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
	analysisId := int64(789) // int64 | Analysis ID
	updateAnalysisInputBody := *revengai.NewUpdateAnalysisInputBody() // UpdateAnalysisInputBody | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysesCoreAPI.V3UpdateAnalysis(context.Background(), analysisId).UpdateAnalysisInputBody(updateAnalysisInputBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysesCoreAPI.V3UpdateAnalysis``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3UpdateAnalysis`: AnalysisDetailOutputBody
	fmt.Fprintf(os.Stdout, "Response from `AnalysesCoreAPI.V3UpdateAnalysis`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int64** | Analysis ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3UpdateAnalysisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAnalysisInputBody** | [**UpdateAnalysisInputBody**](UpdateAnalysisInputBody.md) |  | 

### Return type

[**AnalysisDetailOutputBody**](AnalysisDetailOutputBody.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3UpdateAnalysisTags

> AnalysisTagsOutputBody V3UpdateAnalysisTags(ctx, analysisId).UpdateTagsInputBody(updateTagsInputBody).Execute()

Replace an analysis' tags.



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
	analysisId := int64(789) // int64 | Analysis ID
	updateTagsInputBody := *revengai.NewUpdateTagsInputBody([]string{"Tags_example"}) // UpdateTagsInputBody | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysesCoreAPI.V3UpdateAnalysisTags(context.Background(), analysisId).UpdateTagsInputBody(updateTagsInputBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysesCoreAPI.V3UpdateAnalysisTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3UpdateAnalysisTags`: AnalysisTagsOutputBody
	fmt.Fprintf(os.Stdout, "Response from `AnalysesCoreAPI.V3UpdateAnalysisTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int64** | Analysis ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3UpdateAnalysisTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTagsInputBody** | [**UpdateTagsInputBody**](UpdateTagsInputBody.md) |  | 

### Return type

[**AnalysisTagsOutputBody**](AnalysisTagsOutputBody.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3UpgradeAnalysisModel

> UpgradeAnalysisModelOutputBody V3UpgradeAnalysisModel(ctx, analysisId).Execute()

Re-analyse on the latest model



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
	analysisId := int64(789) // int64 | Analysis ID

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysesCoreAPI.V3UpgradeAnalysisModel(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysesCoreAPI.V3UpgradeAnalysisModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3UpgradeAnalysisModel`: UpgradeAnalysisModelOutputBody
	fmt.Fprintf(os.Stdout, "Response from `AnalysesCoreAPI.V3UpgradeAnalysisModel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int64** | Analysis ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3UpgradeAnalysisModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UpgradeAnalysisModelOutputBody**](UpgradeAnalysisModelOutputBody.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

