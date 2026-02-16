# AnalysesResultsMetadataAPI

All URIs are relative to *https://api.reveng.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAnalysisFunctionsPaginated**](AnalysesResultsMetadataAPI.md#GetAnalysisFunctionsPaginated) | **Get** /v2/analyses/{analysis_id}/functions | Get functions from analysis
[**GetCapabilities**](AnalysesResultsMetadataAPI.md#GetCapabilities) | **Get** /v2/analyses/{analysis_id}/capabilities | Gets the capabilities from the analysis
[**GetCommunities**](AnalysesResultsMetadataAPI.md#GetCommunities) | **Get** /v2/analyses/{analysis_id}/communities | Gets the communities found in the analysis
[**GetFunctionsList**](AnalysesResultsMetadataAPI.md#GetFunctionsList) | **Get** /v2/analyses/{analysis_id}/functions/list | Gets functions from analysis
[**GetPdf**](AnalysesResultsMetadataAPI.md#GetPdf) | **Get** /v2/analyses/{analysis_id}/pdf | Gets the PDF found in the analysis
[**GetSbom**](AnalysesResultsMetadataAPI.md#GetSbom) | **Get** /v2/analyses/{analysis_id}/sbom | Gets the software-bill-of-materials (SBOM) found in the analysis
[**GetTags**](AnalysesResultsMetadataAPI.md#GetTags) | **Get** /v2/analyses/{analysis_id}/tags | Get function tags with maliciousness score
[**GetVulnerabilities**](AnalysesResultsMetadataAPI.md#GetVulnerabilities) | **Get** /v2/analyses/{analysis_id}/vulnerabilities | Gets the vulnerabilities found in the analysis



## GetAnalysisFunctionsPaginated

> BaseResponseAnalysisFunctionsList GetAnalysisFunctionsPaginated(ctx, analysisId).Page(page).PageSize(pageSize).Execute()

Get functions from analysis



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	revengai "github.com/RevEngAI/sdk-go"
)

func main() {
	analysisId := int32(56) // int32 | 
	page := int32(56) // int32 | The page number to retrieve. (optional) (default to 1)
	pageSize := int32(56) // int32 | Number of items per page. (optional) (default to 1000)

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysesResultsMetadataAPI.GetAnalysisFunctionsPaginated(context.Background(), analysisId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysesResultsMetadataAPI.GetAnalysisFunctionsPaginated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAnalysisFunctionsPaginated`: BaseResponseAnalysisFunctionsList
	fmt.Fprintf(os.Stdout, "Response from `AnalysesResultsMetadataAPI.GetAnalysisFunctionsPaginated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAnalysisFunctionsPaginatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | The page number to retrieve. | [default to 1]
 **pageSize** | **int32** | Number of items per page. | [default to 1000]

### Return type

[**BaseResponseAnalysisFunctionsList**](BaseResponseAnalysisFunctionsList.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCapabilities

> BaseResponseCapabilities GetCapabilities(ctx, analysisId).Execute()

Gets the capabilities from the analysis

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	revengai "github.com/RevEngAI/sdk-go"
)

func main() {
	analysisId := int32(56) // int32 | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysesResultsMetadataAPI.GetCapabilities(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysesResultsMetadataAPI.GetCapabilities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCapabilities`: BaseResponseCapabilities
	fmt.Fprintf(os.Stdout, "Response from `AnalysesResultsMetadataAPI.GetCapabilities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCapabilitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BaseResponseCapabilities**](BaseResponseCapabilities.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommunities

> BaseResponseCommunities GetCommunities(ctx, analysisId).UserName(userName).Execute()

Gets the communities found in the analysis

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	revengai "github.com/RevEngAI/sdk-go"
)

func main() {
	analysisId := int32(56) // int32 | 
	userName := "userName_example" // string | The user name to limit communities to (optional)

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysesResultsMetadataAPI.GetCommunities(context.Background(), analysisId).UserName(userName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysesResultsMetadataAPI.GetCommunities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCommunities`: BaseResponseCommunities
	fmt.Fprintf(os.Stdout, "Response from `AnalysesResultsMetadataAPI.GetCommunities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommunitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userName** | **string** | The user name to limit communities to | 

### Return type

[**BaseResponseCommunities**](BaseResponseCommunities.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFunctionsList

> BaseResponseAnalysisFunctions GetFunctionsList(ctx, analysisId).SearchTerm(searchTerm).MinVAddr(minVAddr).MaxVAddr(maxVAddr).IncludeEmbeddings(includeEmbeddings).Page(page).PageSize(pageSize).Execute()

Gets functions from analysis



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	revengai "github.com/RevEngAI/sdk-go"
)

func main() {
	analysisId := int32(56) // int32 | 
	searchTerm := "searchTerm_example" // string |  (optional)
	minVAddr := int32(56) // int32 |  (optional)
	maxVAddr := int32(56) // int32 |  (optional)
	includeEmbeddings := true // bool |  (optional) (default to true)
	page := int32(56) // int32 | The page number to retrieve. (optional) (default to 1)
	pageSize := int32(56) // int32 | Number of items per page. (optional) (default to 1000)

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysesResultsMetadataAPI.GetFunctionsList(context.Background(), analysisId).SearchTerm(searchTerm).MinVAddr(minVAddr).MaxVAddr(maxVAddr).IncludeEmbeddings(includeEmbeddings).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysesResultsMetadataAPI.GetFunctionsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFunctionsList`: BaseResponseAnalysisFunctions
	fmt.Fprintf(os.Stdout, "Response from `AnalysesResultsMetadataAPI.GetFunctionsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFunctionsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **searchTerm** | **string** |  | 
 **minVAddr** | **int32** |  | 
 **maxVAddr** | **int32** |  | 
 **includeEmbeddings** | **bool** |  | [default to true]
 **page** | **int32** | The page number to retrieve. | [default to 1]
 **pageSize** | **int32** | Number of items per page. | [default to 1000]

### Return type

[**BaseResponseAnalysisFunctions**](BaseResponseAnalysisFunctions.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPdf

> interface{} GetPdf(ctx, analysisId).Execute()

Gets the PDF found in the analysis

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	revengai "github.com/RevEngAI/sdk-go"
)

func main() {
	analysisId := int32(56) // int32 | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysesResultsMetadataAPI.GetPdf(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysesResultsMetadataAPI.GetPdf``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPdf`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AnalysesResultsMetadataAPI.GetPdf`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPdfRequest struct via the builder pattern


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


## GetSbom

> BaseResponseListSBOM GetSbom(ctx, analysisId).Execute()

Gets the software-bill-of-materials (SBOM) found in the analysis

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	revengai "github.com/RevEngAI/sdk-go"
)

func main() {
	analysisId := int32(56) // int32 | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysesResultsMetadataAPI.GetSbom(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysesResultsMetadataAPI.GetSbom``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSbom`: BaseResponseListSBOM
	fmt.Fprintf(os.Stdout, "Response from `AnalysesResultsMetadataAPI.GetSbom`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSbomRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BaseResponseListSBOM**](BaseResponseListSBOM.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTags

> BaseResponseAnalysisTags GetTags(ctx, analysisId).Execute()

Get function tags with maliciousness score

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	revengai "github.com/RevEngAI/sdk-go"
)

func main() {
	analysisId := int32(56) // int32 | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysesResultsMetadataAPI.GetTags(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysesResultsMetadataAPI.GetTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTags`: BaseResponseAnalysisTags
	fmt.Fprintf(os.Stdout, "Response from `AnalysesResultsMetadataAPI.GetTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BaseResponseAnalysisTags**](BaseResponseAnalysisTags.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVulnerabilities

> BaseResponseVulnerabilities GetVulnerabilities(ctx, analysisId).Execute()

Gets the vulnerabilities found in the analysis

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	revengai "github.com/RevEngAI/sdk-go"
)

func main() {
	analysisId := int32(56) // int32 | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysesResultsMetadataAPI.GetVulnerabilities(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysesResultsMetadataAPI.GetVulnerabilities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVulnerabilities`: BaseResponseVulnerabilities
	fmt.Fprintf(os.Stdout, "Response from `AnalysesResultsMetadataAPI.GetVulnerabilities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVulnerabilitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BaseResponseVulnerabilities**](BaseResponseVulnerabilities.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

