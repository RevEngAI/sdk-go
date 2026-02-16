# \FirmwareAPI

All URIs are relative to *https://api.reveng.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBinariesForFirmwareTask**](FirmwareAPI.md#GetBinariesForFirmwareTask) | **Get** /v2/firmware/get-binaries/{task_id} | Upload firmware for unpacking
[**UploadFirmware**](FirmwareAPI.md#UploadFirmware) | **Post** /v2/firmware | Upload firmware for unpacking



## GetBinariesForFirmwareTask

> interface{} GetBinariesForFirmwareTask(ctx, taskId).Execute()

Upload firmware for unpacking



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
	taskId := "taskId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirmwareAPI.GetBinariesForFirmwareTask(context.Background(), taskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirmwareAPI.GetBinariesForFirmwareTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBinariesForFirmwareTask`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `FirmwareAPI.GetBinariesForFirmwareTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBinariesForFirmwareTaskRequest struct via the builder pattern


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


## UploadFirmware

> interface{} UploadFirmware(ctx).File(file).Password(password).Execute()

Upload firmware for unpacking



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
	file := os.NewFile(1234, "some_file") // *os.File | 
	password := "password_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirmwareAPI.UploadFirmware(context.Background()).File(file).Password(password).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirmwareAPI.UploadFirmware``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadFirmware`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `FirmwareAPI.UploadFirmware`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadFirmwareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** |  | 
 **password** | **string** |  | 

### Return type

**interface{}**

### Authorization

[APIKey](../README.md#APIKey)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

