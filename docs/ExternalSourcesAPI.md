# ExternalSourcesAPI

All URIs are relative to *https://api.reveng.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V3GetVirustotalScanOperation**](ExternalSourcesAPI.md#V3GetVirustotalScanOperation) | **Get** /v3/operations/virustotal-scan/{binary_id} | Get a VirusTotal scan operation.
[**V3RunVirustotalScan**](ExternalSourcesAPI.md#V3RunVirustotalScan) | **Post** /v3/binaries/{binary_id}/virustotal-scan:run | Trigger a VirusTotal lookup for a binary.



## V3GetVirustotalScanOperation

> OperationVirusTotalScanMetadataVirusTotalScanResult V3GetVirustotalScanOperation(ctx, binaryId).Execute()

Get a VirusTotal scan operation.



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
	resp, r, err := apiClient.ExternalSourcesAPI.V3GetVirustotalScanOperation(context.Background(), binaryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalSourcesAPI.V3GetVirustotalScanOperation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3GetVirustotalScanOperation`: OperationVirusTotalScanMetadataVirusTotalScanResult
	fmt.Fprintf(os.Stdout, "Response from `ExternalSourcesAPI.V3GetVirustotalScanOperation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**binaryId** | **int64** | Binary ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3GetVirustotalScanOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationVirusTotalScanMetadataVirusTotalScanResult**](OperationVirusTotalScanMetadataVirusTotalScanResult.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3RunVirustotalScan

> OperationVirusTotalScanMetadataVirusTotalScanResult V3RunVirustotalScan(ctx, binaryId).Execute()

Trigger a VirusTotal lookup for a binary.



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
	resp, r, err := apiClient.ExternalSourcesAPI.V3RunVirustotalScan(context.Background(), binaryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalSourcesAPI.V3RunVirustotalScan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3RunVirustotalScan`: OperationVirusTotalScanMetadataVirusTotalScanResult
	fmt.Fprintf(os.Stdout, "Response from `ExternalSourcesAPI.V3RunVirustotalScan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**binaryId** | **int64** | Binary ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3RunVirustotalScanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationVirusTotalScanMetadataVirusTotalScanResult**](OperationVirusTotalScanMetadataVirusTotalScanResult.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

