# BinariesAPI

All URIs are relative to *https://api.reveng.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadZippedBinary**](BinariesAPI.md#DownloadZippedBinary) | **Get** /v2/binaries/{binary_id}/download-zipped | Downloads a zipped binary with password protection
[**GetBinaryAdditionalDetails**](BinariesAPI.md#GetBinaryAdditionalDetails) | **Get** /v2/binaries/{binary_id}/additional-details | Gets the additional details of a binary
[**GetBinaryAdditionalDetailsStatus**](BinariesAPI.md#GetBinaryAdditionalDetailsStatus) | **Get** /v2/binaries/{binary_id}/additional-details/status | Gets the status of the additional details task for a binary
[**GetBinaryAdditionalDetailsStatus_0**](BinariesAPI.md#GetBinaryAdditionalDetailsStatus_0) | **Get** /v3/binaries/{binary_id}/additional-details/status | Get the additional-details extraction status for a binary.
[**GetBinaryAdditionalDetails_0**](BinariesAPI.md#GetBinaryAdditionalDetails_0) | **Get** /v3/binaries/{binary_id}/additional-details | Get additional details for a binary.
[**GetBinaryDetails**](BinariesAPI.md#GetBinaryDetails) | **Get** /v2/binaries/{binary_id}/details | Gets the details of a binary
[**GetBinaryDieInfo**](BinariesAPI.md#GetBinaryDieInfo) | **Get** /v2/binaries/{binary_id}/die-info | Gets the die info of a binary
[**GetBinaryExternals**](BinariesAPI.md#GetBinaryExternals) | **Get** /v2/binaries/{binary_id}/externals | Gets the external details of a binary
[**GetBinaryRelatedStatus**](BinariesAPI.md#GetBinaryRelatedStatus) | **Get** /v2/binaries/{binary_id}/related/status | Gets the status of the unpack binary task for a binary
[**GetRelatedBinaries**](BinariesAPI.md#GetRelatedBinaries) | **Get** /v2/binaries/{binary_id}/related | Gets the related binaries of a binary.
[**V3GetBinaryDieInfo**](BinariesAPI.md#V3GetBinaryDieInfo) | **Get** /v3/binaries/{binary_id}/die-info | Get Detect It Easy matches for a binary.
[**V3GetBinaryRelated**](BinariesAPI.md#V3GetBinaryRelated) | **Get** /v3/binaries/{binary_id}/related | Get the binaries related to this one by unpacking.
[**V3GetBinaryRelatedStatus**](BinariesAPI.md#V3GetBinaryRelatedStatus) | **Get** /v3/binaries/{binary_id}/related/status | Get the archive-unpacking status for a binary.
[**V3UploadFile**](BinariesAPI.md#V3UploadFile) | **Post** /v3/upload | Upload a file.



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
	revengai "github.com/RevEngAI/sdk-go/v4"
)

func main() {
	binaryId := int32(56) // int32 | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
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

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

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
	revengai "github.com/RevEngAI/sdk-go/v4"
)

func main() {
	binaryId := int32(56) // int32 | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
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

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

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
	revengai "github.com/RevEngAI/sdk-go/v4"
)

func main() {
	binaryId := int32(56) // int32 | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
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

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBinaryAdditionalDetailsStatus_0

> GetAdditionalDetailsStatusOutputBody GetBinaryAdditionalDetailsStatus_0(ctx, binaryId).Execute()

Get the additional-details extraction status for a binary.



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
	resp, r, err := apiClient.BinariesAPI.GetBinaryAdditionalDetailsStatus_0(context.Background(), binaryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BinariesAPI.GetBinaryAdditionalDetailsStatus_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBinaryAdditionalDetailsStatus_0`: GetAdditionalDetailsStatusOutputBody
	fmt.Fprintf(os.Stdout, "Response from `BinariesAPI.GetBinaryAdditionalDetailsStatus_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**binaryId** | **int64** | Binary ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBinaryAdditionalDetailsStatus_1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetAdditionalDetailsStatusOutputBody**](GetAdditionalDetailsStatusOutputBody.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBinaryAdditionalDetails_0

> GetAdditionalDetailsOutputBody GetBinaryAdditionalDetails_0(ctx, binaryId).Execute()

Get additional details for a binary.



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
	resp, r, err := apiClient.BinariesAPI.GetBinaryAdditionalDetails_0(context.Background(), binaryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BinariesAPI.GetBinaryAdditionalDetails_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBinaryAdditionalDetails_0`: GetAdditionalDetailsOutputBody
	fmt.Fprintf(os.Stdout, "Response from `BinariesAPI.GetBinaryAdditionalDetails_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**binaryId** | **int64** | Binary ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBinaryAdditionalDetails_2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetAdditionalDetailsOutputBody**](GetAdditionalDetailsOutputBody.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

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
	revengai "github.com/RevEngAI/sdk-go/v4"
)

func main() {
	binaryId := int32(56) // int32 | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
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

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

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
	revengai "github.com/RevEngAI/sdk-go/v4"
)

func main() {
	binaryId := int32(56) // int32 | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
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

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

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
	revengai "github.com/RevEngAI/sdk-go/v4"
)

func main() {
	binaryId := int32(56) // int32 | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
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

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

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
	revengai "github.com/RevEngAI/sdk-go/v4"
)

func main() {
	binaryId := int32(56) // int32 | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
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

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

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
	revengai "github.com/RevEngAI/sdk-go/v4"
)

func main() {
	binaryId := int32(56) // int32 | 

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
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

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3GetBinaryDieInfo

> GetDieInfoOutputBody V3GetBinaryDieInfo(ctx, binaryId).Execute()

Get Detect It Easy matches for a binary.



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
	resp, r, err := apiClient.BinariesAPI.V3GetBinaryDieInfo(context.Background(), binaryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BinariesAPI.V3GetBinaryDieInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3GetBinaryDieInfo`: GetDieInfoOutputBody
	fmt.Fprintf(os.Stdout, "Response from `BinariesAPI.V3GetBinaryDieInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**binaryId** | **int64** | Binary ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3GetBinaryDieInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetDieInfoOutputBody**](GetDieInfoOutputBody.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3GetBinaryRelated

> GetRelatedBinariesOutputBody V3GetBinaryRelated(ctx, binaryId).Execute()

Get the binaries related to this one by unpacking.



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
	resp, r, err := apiClient.BinariesAPI.V3GetBinaryRelated(context.Background(), binaryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BinariesAPI.V3GetBinaryRelated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3GetBinaryRelated`: GetRelatedBinariesOutputBody
	fmt.Fprintf(os.Stdout, "Response from `BinariesAPI.V3GetBinaryRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**binaryId** | **int64** | Binary ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3GetBinaryRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetRelatedBinariesOutputBody**](GetRelatedBinariesOutputBody.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3GetBinaryRelatedStatus

> GetRelatedStatusOutputBody V3GetBinaryRelatedStatus(ctx, binaryId).Execute()

Get the archive-unpacking status for a binary.



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
	resp, r, err := apiClient.BinariesAPI.V3GetBinaryRelatedStatus(context.Background(), binaryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BinariesAPI.V3GetBinaryRelatedStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3GetBinaryRelatedStatus`: GetRelatedStatusOutputBody
	fmt.Fprintf(os.Stdout, "Response from `BinariesAPI.V3GetBinaryRelatedStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**binaryId** | **int64** | Binary ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3GetBinaryRelatedStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetRelatedStatusOutputBody**](GetRelatedStatusOutputBody.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3UploadFile

> UploadOutputBody V3UploadFile(ctx).File(file).UploadFileType(uploadFileType).ForceOverwrite(forceOverwrite).Execute()

Upload a file.



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
	file := os.NewFile(1234, "some_file") // *os.File | The file's raw bytes.
	uploadFileType := "uploadFileType_example" // string | The kind of file being uploaded.
	forceOverwrite := true // bool | Re-upload and overwrite even if a file with this hash already exists. (optional)

	configuration := revengai.NewConfiguration()
	apiClient := revengai.NewAPIClient(configuration)
	resp, r, err := apiClient.BinariesAPI.V3UploadFile(context.Background()).File(file).UploadFileType(uploadFileType).ForceOverwrite(forceOverwrite).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BinariesAPI.V3UploadFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3UploadFile`: UploadOutputBody
	fmt.Fprintf(os.Stdout, "Response from `BinariesAPI.V3UploadFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3UploadFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** | The file&#39;s raw bytes. | 
 **uploadFileType** | **string** | The kind of file being uploaded. | 
 **forceOverwrite** | **bool** | Re-upload and overwrite even if a file with this hash already exists. | 

### Return type

[**UploadOutputBody**](UploadOutputBody.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

