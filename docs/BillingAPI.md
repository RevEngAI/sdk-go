# BillingAPI

All URIs are relative to *https://api.reveng.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V3GetCreditEstimates**](BillingAPI.md#V3GetCreditEstimates) | **Get** /v3/credit-estimates | Get credit estimates.



## V3GetCreditEstimates

> GetCreditEstimatesOutputBody V3GetCreditEstimates(ctx).Execute()

Get credit estimates.



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
	resp, r, err := apiClient.BillingAPI.V3GetCreditEstimates(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.V3GetCreditEstimates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V3GetCreditEstimates`: GetCreditEstimatesOutputBody
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.V3GetCreditEstimates`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV3GetCreditEstimatesRequest struct via the builder pattern


### Return type

[**GetCreditEstimatesOutputBody**](GetCreditEstimatesOutputBody.md)

### Authorization

[APIKey](../README.md#APIKey), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

