# \BinariesAPI

All URIs are relative to *https://api.reveng.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadZippedBinary**](BinariesAPI.md#DownloadZippedBinary) | **Get** /v2/binaries/{binary_id}/download-zipped | Downloads a zipped binary with password protection
[**GetBinaryAdditionalDetails**](BinariesAPI.md#GetBinaryAdditionalDetails) | **Get** /v2/binaries/{binary_id}/additional-details | Gets the additional details of a binary
[**GetBinaryAdditionalDetailsStatus**](BinariesAPI.md#GetBinaryAdditionalDetailsStatus) | **Get** /v2/binaries/{binary_id}/additional-details/status | Gets the status of the additional details task for a binary
[**GetBinaryDetails**](BinariesAPI.md#GetBinaryDetails) | **Get** /v2/binaries/{binary_id}/details | Gets the details of a binary
[**GetBinaryDieInfo**](BinariesAPI.md#GetBinaryDieInfo) | **Get** /v2/binaries/{binary_id}/die-info | Gets the die info of a binary
[**GetBinaryExternals**](BinariesAPI.md#GetBinaryExternals) | **Get** /v2/binaries/{binary_id}/externals | Gets the external details of a binary
[**GetBinaryRelatedStatus**](BinariesAPI.md#GetBinaryRelatedStatus) | **Get** /v2/binaries/{binary_id}/related/status | Gets the status of the unpack binary task for a binary
[**GetRelatedBinaries**](BinariesAPI.md#GetRelatedBinaries) | **Get** /v2/binaries/{binary_id}/related | Gets the related binaries of a binary.



## DownloadZippedBinary

> *os.File DownloadZippedBinary(ctx, binaryId).Execute()

Downloads a zipped binary with password protection

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/RevEngAI/sdk-go"
)

func main() {
	binaryId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BinariesAPI.DownloadZippedBinary(context.Background(), binaryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BinariesAPI.DownloadZippedBinary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadZippedBinary`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `BinariesAPI.DownloadZippedBinary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**binaryId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadZippedBinaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/zip, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBinaryAdditionalDetails

> BaseResponseBinaryAdditionalResponse GetBinaryAdditionalDetails(ctx, binaryId).Execute()

Gets the additional details of a binary

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/RevEngAI/sdk-go"
)

func main() {
	binaryId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BinariesAPI.GetBinaryAdditionalDetails(context.Background(), binaryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BinariesAPI.GetBinaryAdditionalDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBinaryAdditionalDetails`: BaseResponseBinaryAdditionalResponse
	fmt.Fprintf(os.Stdout, "Response from `BinariesAPI.GetBinaryAdditionalDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**binaryId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBinaryAdditionalDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BaseResponseBinaryAdditionalResponse**](BaseResponseBinaryAdditionalResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBinaryAdditionalDetailsStatus

> BaseResponseAdditionalDetailsStatusResponse GetBinaryAdditionalDetailsStatus(ctx, binaryId).Execute()

Gets the status of the additional details task for a binary

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/RevEngAI/sdk-go"
)

func main() {
	binaryId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BinariesAPI.GetBinaryAdditionalDetailsStatus(context.Background(), binaryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BinariesAPI.GetBinaryAdditionalDetailsStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBinaryAdditionalDetailsStatus`: BaseResponseAdditionalDetailsStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `BinariesAPI.GetBinaryAdditionalDetailsStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**binaryId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBinaryAdditionalDetailsStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BaseResponseAdditionalDetailsStatusResponse**](BaseResponseAdditionalDetailsStatusResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBinaryDetails

> BaseResponseBinaryDetailsResponse GetBinaryDetails(ctx, binaryId).Execute()

Gets the details of a binary

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/RevEngAI/sdk-go"
)

func main() {
	binaryId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BinariesAPI.GetBinaryDetails(context.Background(), binaryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BinariesAPI.GetBinaryDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBinaryDetails`: BaseResponseBinaryDetailsResponse
	fmt.Fprintf(os.Stdout, "Response from `BinariesAPI.GetBinaryDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**binaryId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBinaryDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BaseResponseBinaryDetailsResponse**](BaseResponseBinaryDetailsResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBinaryDieInfo

> BaseResponseListDieMatch GetBinaryDieInfo(ctx, binaryId).Execute()

Gets the die info of a binary

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/RevEngAI/sdk-go"
)

func main() {
	binaryId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BinariesAPI.GetBinaryDieInfo(context.Background(), binaryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BinariesAPI.GetBinaryDieInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBinaryDieInfo`: BaseResponseListDieMatch
	fmt.Fprintf(os.Stdout, "Response from `BinariesAPI.GetBinaryDieInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**binaryId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBinaryDieInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BaseResponseListDieMatch**](BaseResponseListDieMatch.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBinaryExternals

> BaseResponseBinaryExternalsResponse GetBinaryExternals(ctx, binaryId).Execute()

Gets the external details of a binary

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/RevEngAI/sdk-go"
)

func main() {
	binaryId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BinariesAPI.GetBinaryExternals(context.Background(), binaryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BinariesAPI.GetBinaryExternals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBinaryExternals`: BaseResponseBinaryExternalsResponse
	fmt.Fprintf(os.Stdout, "Response from `BinariesAPI.GetBinaryExternals`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**binaryId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBinaryExternalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BaseResponseBinaryExternalsResponse**](BaseResponseBinaryExternalsResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBinaryRelatedStatus

> BaseResponseBinariesRelatedStatusResponse GetBinaryRelatedStatus(ctx, binaryId).Execute()

Gets the status of the unpack binary task for a binary

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/RevEngAI/sdk-go"
)

func main() {
	binaryId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BinariesAPI.GetBinaryRelatedStatus(context.Background(), binaryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BinariesAPI.GetBinaryRelatedStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBinaryRelatedStatus`: BaseResponseBinariesRelatedStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `BinariesAPI.GetBinaryRelatedStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**binaryId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBinaryRelatedStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BaseResponseBinariesRelatedStatusResponse**](BaseResponseBinariesRelatedStatusResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRelatedBinaries

> BaseResponseChildBinariesResponse GetRelatedBinaries(ctx, binaryId).Execute()

Gets the related binaries of a binary.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/RevEngAI/sdk-go"
)

func main() {
	binaryId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BinariesAPI.GetRelatedBinaries(context.Background(), binaryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BinariesAPI.GetRelatedBinaries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRelatedBinaries`: BaseResponseChildBinariesResponse
	fmt.Fprintf(os.Stdout, "Response from `BinariesAPI.GetRelatedBinaries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**binaryId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRelatedBinariesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BaseResponseChildBinariesResponse**](BaseResponseChildBinariesResponse.md)

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

