package main

import (
	"context"
	"fmt"
	"os"

	sdk "github.com/RevEngAI/sdk-go/v3"
)

func main() {
	cfg := sdk.NewConfiguration()
	cfg.Servers = sdk.ServerConfigurations{
		{URL: "http://localhost:10000", Description: "Local"},
	}

	client := sdk.NewAPIClient(cfg)

	ctx := context.WithValue(context.Background(), sdk.ContextAPIKeys, map[string]sdk.APIKey{
		"APIKey": {Key: "FAKE_KEY_NOT_IN_USE1"},
	})

	analysisId := int32(8861687)

	result, httpResp, err := client.FunctionsRenamingHistoryAPI.GetFunctionNameHistory(ctx, analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Exception when calling AnalysesCoreAPI#GetAnalysisLogs\n")
		if httpResp != nil {
			fmt.Fprintf(os.Stderr, "Status code: %d\n", httpResp.StatusCode)
		}
		if apiErr, ok := err.(*sdk.GenericOpenAPIError); ok {
			fmt.Fprintf(os.Stderr, "Reason: %s\n", string(apiErr.Body()))
		}
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if result.HasData() {
		data := result.GetData()
		fmt.Println(data)
	} else {
		fmt.Println("No logs available")
	}
}
