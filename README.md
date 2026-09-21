# RevEng.AI Go SDK

This is the Go SDK for the RevEng.AI API.

To use the SDK you will first need to obtain an API key from [https://reveng.ai](https://reveng.ai/register).

## Installation

Once you have the API key you can install the SDK using:
```shell
go get github.com/RevEngAI/sdk-go
```

## Usage

The following is an example of how to use the SDK to get the logs of an analysis:

```go
package main

import (
	"context"
	"fmt"
	"os"

	sdk "github.com/RevEngAI/sdk-go"
)

func main() {
	// Create a new configuration with default settings
	cfg := sdk.NewConfiguration()

	// Create a new API client
	client := sdk.NewAPIClient(cfg)

	// Configure API key authorization via context
	ctx := context.WithValue(context.Background(), sdk.ContextAPIKeys, map[string]sdk.APIKey{
		"APIKey": {Key: os.Getenv("API_KEY")},
	})

	analysisId := int32(715320)

	// Call GetAnalysisLogs on the AnalysesCoreAPI service
	result, httpResp, err := client.AnalysesCoreAPI.GetAnalysisLogs(ctx, analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Exception when calling AnalysesCoreAPI#GetAnalysisLogs\n")
		fmt.Fprintf(os.Stderr, "Status code: %d\n", httpResp.StatusCode)

		// Check if the error is a GenericOpenAPIError for more details
		if apiErr, ok := err.(*sdk.GenericOpenAPIError); ok {
			fmt.Fprintf(os.Stderr, "Reason: %s\n", string(apiErr.Body()))
		}

		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Unwrap the response: BaseResponseLogs -> Logs (Data) -> string
	if result.HasData() {
		data := result.GetData()
		fmt.Println(data.GetLogs())
	} else {
		fmt.Println("No logs available")
	}
}
```

## Documentation for API Endpoints

All URIs are relative to *https://api.reveng.ai*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AgentAPI* | [**CheckCapabilitiesTaskStatusV2AnalysesAnalysisIdAgentCapabilitiesStatusGet**](docs/AgentAPI.md#checkcapabilitiestaskstatusv2analysesanalysisidagentcapabilitiesstatusget) | **Get** /v2/analyses/{analysis_id}/agent/capabilities/status | Check the status of a capabilities analysis workflow
*AgentAPI* | [**CheckProtocolsTaskStatusV2AnalysesAnalysisIdAgentProtocolsStatusGet**](docs/AgentAPI.md#checkprotocolstaskstatusv2analysesanalysisidagentprotocolsstatusget) | **Get** /v2/analyses/{analysis_id}/agent/protocols/status | Check the status of a protocols discovery workflow
*AgentAPI* | [**CheckRemediationTaskStatusV2AnalysesAnalysisIdAgentRemediationStatusGet**](docs/AgentAPI.md#checkremediationtaskstatusv2analysesanalysisidagentremediationstatusget) | **Get** /v2/analyses/{analysis_id}/agent/remediation/status | Check the status of a remediation analysis workflow
*AgentAPI* | [**CheckReportAnalysisTaskStatusV2AnalysesAnalysisIdAgentReportAnalysisStatusGet**](docs/AgentAPI.md#checkreportanalysistaskstatusv2analysesanalysisidagentreportanalysisstatusget) | **Get** /v2/analyses/{analysis_id}/agent/report-analysis/status | Check the status of a report analysis workflow
*AgentAPI* | [**CheckSecretsTaskStatusV2AnalysesAnalysisIdAgentSecretsStatusGet**](docs/AgentAPI.md#checksecretstaskstatusv2analysesanalysisidagentsecretsstatusget) | **Get** /v2/analyses/{analysis_id}/agent/secrets/status | Check the status of a secrets discovery workflow
*AgentAPI* | [**CheckTriageTaskStatusV2AnalysesAnalysisIdAgentTriageStatusGet**](docs/AgentAPI.md#checktriagetaskstatusv2analysesanalysisidagenttriagestatusget) | **Get** /v2/analyses/{analysis_id}/agent/triage/status | Check the status of a triage analysis workflow
*AgentAPI* | [**CreateCapabilitiesTaskV2AnalysesAnalysisIdAgentCapabilitiesPost**](docs/AgentAPI.md#createcapabilitiestaskv2analysesanalysisidagentcapabilitiespost) | **Post** /v2/analyses/{analysis_id}/agent/capabilities | Queues a capabilities analysis workflow process
*AgentAPI* | [**CreateProtocolsTaskV2AnalysesAnalysisIdAgentProtocolsPost**](docs/AgentAPI.md#createprotocolstaskv2analysesanalysisidagentprotocolspost) | **Post** /v2/analyses/{analysis_id}/agent/protocols | Queues a protocols discovery workflow process
*AgentAPI* | [**CreateRemediationTaskV2AnalysesAnalysisIdAgentRemediationPost**](docs/AgentAPI.md#createremediationtaskv2analysesanalysisidagentremediationpost) | **Post** /v2/analyses/{analysis_id}/agent/remediation | Queues a remediation analysis workflow process
*AgentAPI* | [**CreateReportAnalysisTaskV2AnalysesAnalysisIdAgentReportAnalysisPost**](docs/AgentAPI.md#createreportanalysistaskv2analysesanalysisidagentreportanalysispost) | **Post** /v2/analyses/{analysis_id}/agent/report-analysis | Queues a combined report analysis workflow process
*AgentAPI* | [**CreateSecretsTaskV2AnalysesAnalysisIdAgentSecretsPost**](docs/AgentAPI.md#createsecretstaskv2analysesanalysisidagentsecretspost) | **Post** /v2/analyses/{analysis_id}/agent/secrets | Queues a secrets discovery workflow process
*AgentAPI* | [**CreateTriageTaskV2AnalysesAnalysisIdAgentTriagePost**](docs/AgentAPI.md#createtriagetaskv2analysesanalysisidagenttriagepost) | **Post** /v2/analyses/{analysis_id}/agent/triage | Queues a triage analysis workflow process
*AgentAPI* | [**GetCapabilitiesResultV2AnalysesAnalysisIdAgentCapabilitiesGet**](docs/AgentAPI.md#getcapabilitiesresultv2analysesanalysisidagentcapabilitiesget) | **Get** /v2/analyses/{analysis_id}/agent/capabilities | Get Capabilities Result
*AgentAPI* | [**GetProtocolsResultV2AnalysesAnalysisIdAgentProtocolsGet**](docs/AgentAPI.md#getprotocolsresultv2analysesanalysisidagentprotocolsget) | **Get** /v2/analyses/{analysis_id}/agent/protocols | Get Protocols Result
*AgentAPI* | [**GetRemediationResultV2AnalysesAnalysisIdAgentRemediationGet**](docs/AgentAPI.md#getremediationresultv2analysesanalysisidagentremediationget) | **Get** /v2/analyses/{analysis_id}/agent/remediation | Get Remediation Result
*AgentAPI* | [**GetReportAnalysisResultV2AnalysesAnalysisIdAgentReportAnalysisGet**](docs/AgentAPI.md#getreportanalysisresultv2analysesanalysisidagentreportanalysisget) | **Get** /v2/analyses/{analysis_id}/agent/report-analysis | Get Report Analysis Result
*AgentAPI* | [**GetSecretsResultV2AnalysesAnalysisIdAgentSecretsGet**](docs/AgentAPI.md#getsecretsresultv2analysesanalysisidagentsecretsget) | **Get** /v2/analyses/{analysis_id}/agent/secrets | Get Secrets Result
*AgentAPI* | [**GetTriageResultV2AnalysesAnalysisIdAgentTriageGet**](docs/AgentAPI.md#gettriageresultv2analysesanalysisidagenttriageget) | **Get** /v2/analyses/{analysis_id}/agent/triage | Get Triage Result
*AgentAPI* | [**V3CancelRenameUnnamedFunctions**](docs/AgentAPI.md#v3cancelrenameunnamedfunctions) | **Post** /v3/analyses/{analysis_id}/agents/rename-unnamed-functions/cancel | Cancel the rename-unnamed-functions agent.
*AgentAPI* | [**V3CancelSecurityScanOperation**](docs/AgentAPI.md#v3cancelsecurityscanoperation) | **Post** /v3/operations/security-scan/{analysis_id}:cancel | Cancel a security-scan operation.
*AgentAPI* | [**V3GetBinaryAgentFeedback**](docs/AgentAPI.md#v3getbinaryagentfeedback) | **Get** /v3/analyses/{analysis_id}/agents/{agent}/feedback | Get the caller&#39;s feedback on an agent&#39;s output.
*AgentAPI* | [**V3GetCapabilitiesOperation**](docs/AgentAPI.md#v3getcapabilitiesoperation) | **Get** /v3/operations/capabilities/{analysis_id} | Get a capabilities operation.
*AgentAPI* | [**V3GetCryptoExplainOperation**](docs/AgentAPI.md#v3getcryptoexplainoperation) | **Get** /v3/operations/crypto-explain/{function_id} | Get a crypto-explain operation.
*AgentAPI* | [**V3GetCryptoScanOperation**](docs/AgentAPI.md#v3getcryptoscanoperation) | **Get** /v3/operations/crypto-scan/{analysis_id} | Get a crypto-scan operation.
*AgentAPI* | [**V3GetExecutionExplainOperation**](docs/AgentAPI.md#v3getexecutionexplainoperation) | **Get** /v3/operations/execution-explain/{function_id} | Get an execution-explain operation.
*AgentAPI* | [**V3GetExecutionScanOperation**](docs/AgentAPI.md#v3getexecutionscanoperation) | **Get** /v3/operations/execution-scan/{analysis_id} | Get an execution-scan operation.
*AgentAPI* | [**V3GetFilesystemAnalyseOperation**](docs/AgentAPI.md#v3getfilesystemanalyseoperation) | **Get** /v3/operations/filesystem-analyse/{function_id} | Get a filesystem-analyse operation.
*AgentAPI* | [**V3GetFilesystemScanOperation**](docs/AgentAPI.md#v3getfilesystemscanoperation) | **Get** /v3/operations/filesystem-scan/{analysis_id} | Get a filesystem-scan operation.
*AgentAPI* | [**V3GetNetworkingExplainOperation**](docs/AgentAPI.md#v3getnetworkingexplainoperation) | **Get** /v3/operations/networking-explain/{function_id} | Get a networking-explain operation.
*AgentAPI* | [**V3GetNetworkingScanOperation**](docs/AgentAPI.md#v3getnetworkingscanoperation) | **Get** /v3/operations/networking-scan/{analysis_id} | Get a networking-scan operation.
*AgentAPI* | [**V3GetProtocolsOperation**](docs/AgentAPI.md#v3getprotocolsoperation) | **Get** /v3/operations/protocols/{analysis_id} | Get a protocols operation.
*AgentAPI* | [**V3GetRemediationOperation**](docs/AgentAPI.md#v3getremediationoperation) | **Get** /v3/operations/remediation/{analysis_id} | Get a remediation operation.
*AgentAPI* | [**V3GetRenameUnnamedFunctionsResult**](docs/AgentAPI.md#v3getrenameunnamedfunctionsresult) | **Get** /v3/analyses/{analysis_id}/agents/rename-unnamed-functions | Get rename-unnamed-functions agent result.
*AgentAPI* | [**V3GetRenameUnnamedFunctionsStatus**](docs/AgentAPI.md#v3getrenameunnamedfunctionsstatus) | **Get** /v3/analyses/{analysis_id}/agents/rename-unnamed-functions/status | Get rename-unnamed-functions agent status.
*AgentAPI* | [**V3GetReportAnalysisOperation**](docs/AgentAPI.md#v3getreportanalysisoperation) | **Get** /v3/operations/report-analysis/{analysis_id} | Get a report-analysis operation.
*AgentAPI* | [**V3GetSecretsOperation**](docs/AgentAPI.md#v3getsecretsoperation) | **Get** /v3/operations/secrets/{analysis_id} | Get a secrets operation.
*AgentAPI* | [**V3GetSecurityScanOperation**](docs/AgentAPI.md#v3getsecurityscanoperation) | **Get** /v3/operations/security-scan/{analysis_id} | Get a security-scan operation.
*AgentAPI* | [**V3GetTriageOperation**](docs/AgentAPI.md#v3gettriageoperation) | **Get** /v3/operations/triage/{analysis_id} | Get a triage operation.
*AgentAPI* | [**V3RunCapabilities**](docs/AgentAPI.md#v3runcapabilities) | **Post** /v3/analyses/{analysis_id}/capabilities:run | Run the capabilities agent.
*AgentAPI* | [**V3RunCryptoExplain**](docs/AgentAPI.md#v3runcryptoexplain) | **Post** /v3/functions/{function_id}/crypto-explain:run | Run the crypto-explain agent.
*AgentAPI* | [**V3RunCryptoScan**](docs/AgentAPI.md#v3runcryptoscan) | **Post** /v3/analyses/{analysis_id}/crypto-scan:run | Run the crypto-scan agent.
*AgentAPI* | [**V3RunExecutionExplain**](docs/AgentAPI.md#v3runexecutionexplain) | **Post** /v3/functions/{function_id}/execution-explain:run | Run the execution-explain agent.
*AgentAPI* | [**V3RunExecutionScan**](docs/AgentAPI.md#v3runexecutionscan) | **Post** /v3/analyses/{analysis_id}/execution-scan:run | Run the execution-scan agent.
*AgentAPI* | [**V3RunFilesystemAnalyse**](docs/AgentAPI.md#v3runfilesystemanalyse) | **Post** /v3/functions/{function_id}/filesystem-analyse:run | Run the filesystem-analyse agent.
*AgentAPI* | [**V3RunFilesystemScan**](docs/AgentAPI.md#v3runfilesystemscan) | **Post** /v3/analyses/{analysis_id}/filesystem-scan:run | Run the filesystem-scan agent.
*AgentAPI* | [**V3RunNetworkingExplain**](docs/AgentAPI.md#v3runnetworkingexplain) | **Post** /v3/functions/{function_id}/networking-explain:run | Run the networking-explain agent.
*AgentAPI* | [**V3RunNetworkingScan**](docs/AgentAPI.md#v3runnetworkingscan) | **Post** /v3/analyses/{analysis_id}/networking-scan:run | Run the networking-scan agent.
*AgentAPI* | [**V3RunProtocols**](docs/AgentAPI.md#v3runprotocols) | **Post** /v3/analyses/{analysis_id}/protocols:run | Run the protocols agent.
*AgentAPI* | [**V3RunRemediation**](docs/AgentAPI.md#v3runremediation) | **Post** /v3/analyses/{analysis_id}/remediation:run | Run the remediation agent.
*AgentAPI* | [**V3RunReportAnalysis**](docs/AgentAPI.md#v3runreportanalysis) | **Post** /v3/analyses/{analysis_id}/report-analysis:run | Run the report-analysis agent.
*AgentAPI* | [**V3RunSecrets**](docs/AgentAPI.md#v3runsecrets) | **Post** /v3/analyses/{analysis_id}/secrets:run | Run the secrets agent.
*AgentAPI* | [**V3RunSecurityScan**](docs/AgentAPI.md#v3runsecurityscan) | **Post** /v3/analyses/{analysis_id}/security-scan:run | Run the security-scan agent.
*AgentAPI* | [**V3RunTriage**](docs/AgentAPI.md#v3runtriage) | **Post** /v3/analyses/{analysis_id}/triage:run | Run the triage agent.
*AgentAPI* | [**V3TriggerRenameUnnamedFunctions**](docs/AgentAPI.md#v3triggerrenameunnamedfunctions) | **Post** /v3/analyses/{analysis_id}/agents/rename-unnamed-functions | Run the rename-unnamed-functions agent.
*AgentAPI* | [**V3UpsertBinaryAgentFeedback**](docs/AgentAPI.md#v3upsertbinaryagentfeedback) | **Put** /v3/analyses/{analysis_id}/agents/{agent}/feedback | Record feedback on an agent&#39;s output.
*AnalysesBulkActionsAPI* | [**BulkAddAnalysisTags**](docs/AnalysesBulkActionsAPI.md#bulkaddanalysistags) | **Patch** /v2/analyses/tags/add | Bulk Add Analysis Tags
*AnalysesBulkActionsAPI* | [**BulkDeleteAnalyses**](docs/AnalysesBulkActionsAPI.md#bulkdeleteanalyses) | **Patch** /v2/analyses/delete | Bulk Delete Analyses
*AnalysesCommentsAPI* | [**CreateAnalysisComment**](docs/AnalysesCommentsAPI.md#createanalysiscomment) | **Post** /v2/analyses/{analysis_id}/comments | Create a comment for this analysis
*AnalysesCommentsAPI* | [**DeleteAnalysisComment**](docs/AnalysesCommentsAPI.md#deleteanalysiscomment) | **Delete** /v2/analyses/{analysis_id}/comments/{comment_id} | Delete a comment
*AnalysesCommentsAPI* | [**GetAnalysisComments**](docs/AnalysesCommentsAPI.md#getanalysiscomments) | **Get** /v2/analyses/{analysis_id}/comments | Get comments for this analysis
*AnalysesCommentsAPI* | [**UpdateAnalysisComment**](docs/AnalysesCommentsAPI.md#updateanalysiscomment) | **Patch** /v2/analyses/{analysis_id}/comments/{comment_id} | Update a comment
*AnalysesCoreAPI* | [**AddUserStringToAnalysis**](docs/AnalysesCoreAPI.md#adduserstringtoanalysis) | **Post** /v3/analyses/{analysis_id}/user-provided-strings | Add a user-provided string to an analysis.
*AnalysesCoreAPI* | [**CreateAnalysis**](docs/AnalysesCoreAPI.md#createanalysis) | **Post** /v2/analyses | Create Analysis
*AnalysesCoreAPI* | [**DeleteAnalysis**](docs/AnalysesCoreAPI.md#deleteanalysis) | **Delete** /v2/analyses/{analysis_id} | Delete Analysis
*AnalysesCoreAPI* | [**GetAnalysisBasicInfo**](docs/AnalysesCoreAPI.md#getanalysisbasicinfo) | **Get** /v2/analyses/{analysis_id}/basic | Gets basic analysis information
*AnalysesCoreAPI* | [**GetAnalysisBasicInfo_0**](docs/AnalysesCoreAPI.md#getanalysisbasicinfo_0) | **Get** /v3/analyses/{analysis_id}/basic | Get basic analysis information
*AnalysesCoreAPI* | [**GetAnalysisBytes**](docs/AnalysesCoreAPI.md#getanalysisbytes) | **Get** /v3/analyses/{analysis_id}/bytes | Get the bytes of a binary
*AnalysesCoreAPI* | [**GetAnalysisFunctionMap**](docs/AnalysesCoreAPI.md#getanalysisfunctionmap) | **Get** /v2/analyses/{analysis_id}/func_maps | Get Analysis Function Map
*AnalysesCoreAPI* | [**GetAnalysisFunctionMatches**](docs/AnalysesCoreAPI.md#getanalysisfunctionmatches) | **Get** /v3/analyses/{analysis_id}/functions/matches | Get function-matching results for an analysis
*AnalysesCoreAPI* | [**GetAnalysisFunctionMatchingStatus**](docs/AnalysesCoreAPI.md#getanalysisfunctionmatchingstatus) | **Get** /v3/analyses/{analysis_id}/functions/matches/status | Get function-matching status for an analysis
*AnalysesCoreAPI* | [**GetAnalysisLogs**](docs/AnalysesCoreAPI.md#getanalysislogs) | **Get** /v2/analyses/{analysis_id}/logs | Gets the logs of an analysis
*AnalysesCoreAPI* | [**GetAnalysisParams**](docs/AnalysesCoreAPI.md#getanalysisparams) | **Get** /v2/analyses/{analysis_id}/params | Gets analysis param information
*AnalysesCoreAPI* | [**GetAnalysisStatus**](docs/AnalysesCoreAPI.md#getanalysisstatus) | **Get** /v2/analyses/{analysis_id}/status | Gets the status of an analysis
*AnalysesCoreAPI* | [**GetDynamicExecutionReport**](docs/AnalysesCoreAPI.md#getdynamicexecutionreport) | **Get** /v2/analyses/{analysis_id}/dynamic-execution/report | Get dynamic execution report
*AnalysesCoreAPI* | [**GetDynamicExecutionStatus**](docs/AnalysesCoreAPI.md#getdynamicexecutionstatus) | **Get** /v2/analyses/{analysis_id}/dynamic-execution/status | Get dynamic execution status
*AnalysesCoreAPI* | [**InsertAnalysisLog**](docs/AnalysesCoreAPI.md#insertanalysislog) | **Post** /v2/analyses/{analysis_id}/logs | Insert a log entry for an analysis
*AnalysesCoreAPI* | [**ListAnalyses**](docs/AnalysesCoreAPI.md#listanalyses) | **Get** /v2/analyses/list | Gets the most recent analyses
*AnalysesCoreAPI* | [**LookupBinaryId**](docs/AnalysesCoreAPI.md#lookupbinaryid) | **Get** /v2/analyses/lookup/{binary_id} | Gets the analysis ID from binary ID
*AnalysesCoreAPI* | [**PutAnalysisStrings**](docs/AnalysesCoreAPI.md#putanalysisstrings) | **Put** /v2/analyses/{analysis_id}/strings | Add strings to the analysis
*AnalysesCoreAPI* | [**RequeueAnalysis**](docs/AnalysesCoreAPI.md#requeueanalysis) | **Post** /v2/analyses/{analysis_id}/requeue | Requeue Analysis
*AnalysesCoreAPI* | [**StartAnalysisFunctionMatching**](docs/AnalysesCoreAPI.md#startanalysisfunctionmatching) | **Post** /v3/analyses/{analysis_id}/functions/matches | Start function matching for an analysis
*AnalysesCoreAPI* | [**UpdateAnalysis**](docs/AnalysesCoreAPI.md#updateanalysis) | **Patch** /v2/analyses/{analysis_id} | Update Analysis
*AnalysesCoreAPI* | [**UpdateAnalysisTags**](docs/AnalysesCoreAPI.md#updateanalysistags) | **Patch** /v2/analyses/{analysis_id}/tags | Update Analysis Tags
*AnalysesCoreAPI* | [**UploadFile**](docs/AnalysesCoreAPI.md#uploadfile) | **Post** /v2/upload | Upload File
*AnalysesCoreAPI* | [**V3CreateAnalysis**](docs/AnalysesCoreAPI.md#v3createanalysis) | **Post** /v3/analyses | Create an analysis
*AnalysesCoreAPI* | [**V3GetAnalysisAutoUnstripStatus**](docs/AnalysesCoreAPI.md#v3getanalysisautounstripstatus) | **Get** /v3/analyses/{analysis_id}/auto-unstrip/status | Get the auto-unstrip status for an analysis.
*AnalysesCoreAPI* | [**V3GetAnalysisFunctionsProgress**](docs/AnalysesCoreAPI.md#v3getanalysisfunctionsprogress) | **Get** /v3/analyses/{analysis_id}/progress/functions | Get function embedding progress for an analysis.
*AnalysesCoreAPI* | [**V3GetAnalysisLogs**](docs/AnalysesCoreAPI.md#v3getanalysislogs) | **Get** /v3/analyses/{analysis_id}/logs | Get the Analysis log
*AnalysesCoreAPI* | [**V3GetAnalysisOperation**](docs/AnalysesCoreAPI.md#v3getanalysisoperation) | **Get** /v3/operations/analyses/{analysis_id} | Get an Analysis-creation operation
*AnalysesCoreAPI* | [**V3GetAnalysisStrings**](docs/AnalysesCoreAPI.md#v3getanalysisstrings) | **Get** /v3/analyses/{analysis_id}/functions/strings | List strings for an analysis.
*AnalysesCoreAPI* | [**V3GetAnalysisStringsStatus**](docs/AnalysesCoreAPI.md#v3getanalysisstringsstatus) | **Get** /v3/analyses/{analysis_id}/functions/strings/status | Get the string-extraction status for an analysis.
*AnalysesCoreAPI* | [**V3ListAnalyses**](docs/AnalysesCoreAPI.md#v3listanalyses) | **Get** /v3/analyses | List analyses
*AnalysesCoreAPI* | [**V3ListExampleAnalyses**](docs/AnalysesCoreAPI.md#v3listexampleanalyses) | **Get** /v3/analyses/examples | List example analyses
*AnalysesCoreAPI* | [**V3UpgradeAnalysisModel**](docs/AnalysesCoreAPI.md#v3upgradeanalysismodel) | **Post** /v3/analyses/{analysis_id}/upgrade-model | Re-analyse on the latest model
*AnalysesResultsMetadataAPI* | [**GetAnalysisFunctionsPaginated**](docs/AnalysesResultsMetadataAPI.md#getanalysisfunctionspaginated) | **Get** /v2/analyses/{analysis_id}/functions | Get functions from analysis
*AnalysesResultsMetadataAPI* | [**GetCapabilities**](docs/AnalysesResultsMetadataAPI.md#getcapabilities) | **Get** /v2/analyses/{analysis_id}/capabilities | Gets the capabilities from the analysis
*AnalysesResultsMetadataAPI* | [**GetFunctionsList**](docs/AnalysesResultsMetadataAPI.md#getfunctionslist) | **Get** /v2/analyses/{analysis_id}/functions/list | Gets functions from analysis
*AnalysesResultsMetadataAPI* | [**GetTags**](docs/AnalysesResultsMetadataAPI.md#gettags) | **Get** /v2/analyses/{analysis_id}/tags | Get function tags with maliciousness score
*AnalysesXRefsAPI* | [**GetXrefByVaddr**](docs/AnalysesXRefsAPI.md#getxrefbyvaddr) | **Get** /v2/analyses/{analysis_id}/xrefs/{vaddr} | [Beta] Look up xrefs by virtual address
*AuthenticationUsersAPI* | [**GetUser**](docs/AuthenticationUsersAPI.md#getuser) | **Get** /v2/users/{user_id} | Get a user&#39;s public information
*AuthenticationUsersAPI* | [**GetUserActivity**](docs/AuthenticationUsersAPI.md#getuseractivity) | **Get** /v2/users/activity | Get auth user activity
*AuthenticationUsersAPI* | [**SubmitUserFeedback**](docs/AuthenticationUsersAPI.md#submituserfeedback) | **Post** /v2/users/feedback | Submit feedback about the application
*BinariesAPI* | [**DownloadZippedBinary**](docs/BinariesAPI.md#downloadzippedbinary) | **Get** /v2/binaries/{binary_id}/download-zipped | Downloads a zipped binary with password protection
*BinariesAPI* | [**GetBinaryAdditionalDetails**](docs/BinariesAPI.md#getbinaryadditionaldetails) | **Get** /v2/binaries/{binary_id}/additional-details | Gets the additional details of a binary
*BinariesAPI* | [**GetBinaryAdditionalDetailsStatus**](docs/BinariesAPI.md#getbinaryadditionaldetailsstatus) | **Get** /v2/binaries/{binary_id}/additional-details/status | Gets the status of the additional details task for a binary
*BinariesAPI* | [**GetBinaryAdditionalDetailsStatus_0**](docs/BinariesAPI.md#getbinaryadditionaldetailsstatus_0) | **Get** /v3/binaries/{binary_id}/additional-details/status | Get the additional-details extraction status for a binary.
*BinariesAPI* | [**GetBinaryAdditionalDetails_0**](docs/BinariesAPI.md#getbinaryadditionaldetails_0) | **Get** /v3/binaries/{binary_id}/additional-details | Get additional details for a binary.
*BinariesAPI* | [**GetBinaryDetails**](docs/BinariesAPI.md#getbinarydetails) | **Get** /v2/binaries/{binary_id}/details | Gets the details of a binary
*BinariesAPI* | [**GetBinaryDieInfo**](docs/BinariesAPI.md#getbinarydieinfo) | **Get** /v2/binaries/{binary_id}/die-info | Gets the die info of a binary
*BinariesAPI* | [**GetBinaryExternals**](docs/BinariesAPI.md#getbinaryexternals) | **Get** /v2/binaries/{binary_id}/externals | Gets the external details of a binary
*BinariesAPI* | [**GetBinaryRelatedStatus**](docs/BinariesAPI.md#getbinaryrelatedstatus) | **Get** /v2/binaries/{binary_id}/related/status | Gets the status of the unpack binary task for a binary
*BinariesAPI* | [**GetRelatedBinaries**](docs/BinariesAPI.md#getrelatedbinaries) | **Get** /v2/binaries/{binary_id}/related | Gets the related binaries of a binary.
*BinariesAPI* | [**V3GetBinaryDieInfo**](docs/BinariesAPI.md#v3getbinarydieinfo) | **Get** /v3/binaries/{binary_id}/die-info | Get Detect It Easy matches for a binary.
*BinariesAPI* | [**V3GetBinaryRelated**](docs/BinariesAPI.md#v3getbinaryrelated) | **Get** /v3/binaries/{binary_id}/related | Get the binaries related to this one by unpacking.
*BinariesAPI* | [**V3GetBinaryRelatedStatus**](docs/BinariesAPI.md#v3getbinaryrelatedstatus) | **Get** /v3/binaries/{binary_id}/related/status | Get the archive-unpacking status for a binary.
*BinariesAPI* | [**V3UploadFile**](docs/BinariesAPI.md#v3uploadfile) | **Post** /v3/upload | Upload a file.
*CollectionsAPI* | [**CreateCollection**](docs/CollectionsAPI.md#createcollection) | **Post** /v2/collections | Creates new collection information
*CollectionsAPI* | [**DeleteCollection**](docs/CollectionsAPI.md#deletecollection) | **Delete** /v2/collections/{collection_id} | Deletes a collection
*CollectionsAPI* | [**GetCollection**](docs/CollectionsAPI.md#getcollection) | **Get** /v2/collections/{collection_id} | Returns a collection
*CollectionsAPI* | [**ListCollections**](docs/CollectionsAPI.md#listcollections) | **Get** /v2/collections | Gets basic collections information
*CollectionsAPI* | [**UpdateCollection**](docs/CollectionsAPI.md#updatecollection) | **Patch** /v2/collections/{collection_id} | Updates a collection
*CollectionsAPI* | [**UpdateCollectionBinaries**](docs/CollectionsAPI.md#updatecollectionbinaries) | **Patch** /v2/collections/{collection_id}/binaries | Updates a collection binaries
*CollectionsAPI* | [**UpdateCollectionTags**](docs/CollectionsAPI.md#updatecollectiontags) | **Patch** /v2/collections/{collection_id}/tags | Updates a collection tags
*CollectionsAPI* | [**V3AddCollectionBinaries**](docs/CollectionsAPI.md#v3addcollectionbinaries) | **Post** /v3/collections/{collection_id}/binaries | Add binaries to a collection.
*CollectionsAPI* | [**V3CreateCollection**](docs/CollectionsAPI.md#v3createcollection) | **Post** /v3/collections | Create a collection.
*CollectionsAPI* | [**V3DeleteCollection**](docs/CollectionsAPI.md#v3deletecollection) | **Delete** /v3/collections/{collection_id} | Delete a collection.
*CollectionsAPI* | [**V3GetCollection**](docs/CollectionsAPI.md#v3getcollection) | **Get** /v3/collections/{collection_id} | Get a collection.
*CollectionsAPI* | [**V3ListCollections**](docs/CollectionsAPI.md#v3listcollections) | **Get** /v3/collections | List collections.
*CollectionsAPI* | [**V3PatchCollection**](docs/CollectionsAPI.md#v3patchcollection) | **Patch** /v3/collections/{collection_id} | Update a collection.
*CollectionsAPI* | [**V3PatchCollectionBinaries**](docs/CollectionsAPI.md#v3patchcollectionbinaries) | **Patch** /v3/collections/{collection_id}/binaries | Replace the binaries in a collection.
*CollectionsAPI* | [**V3PatchCollectionTags**](docs/CollectionsAPI.md#v3patchcollectiontags) | **Patch** /v3/collections/{collection_id}/tags | Replace the tags on a collection.
*CollectionsAPI* | [**V3RemoveCollectionBinaries**](docs/CollectionsAPI.md#v3removecollectionbinaries) | **Delete** /v3/collections/{collection_id}/binaries | Remove binaries from a collection.
*ConfigAPI* | [**GetConfig**](docs/ConfigAPI.md#getconfig) | **Get** /v2/config | Get Config
*ConfigAPI* | [**V3GetConfig**](docs/ConfigAPI.md#v3getconfig) | **Get** /v3/config | Get client configuration.
*ConfigAPI* | [**V3GetModels**](docs/ConfigAPI.md#v3getmodels) | **Get** /v3/models | Get the models available for analysis.
*ConversationsAPI* | [**CancelRun**](docs/ConversationsAPI.md#cancelrun) | **Post** /v2/conversations/{id}/cancel | Cancel an active run
*ConversationsAPI* | [**ConfirmTool**](docs/ConversationsAPI.md#confirmtool) | **Post** /v2/conversations/{id}/confirm | Approve or reject a pending tool confirmation
*ConversationsAPI* | [**CreateConversation**](docs/ConversationsAPI.md#createconversation) | **Post** /v2/conversations | Create a new conversation
*ConversationsAPI* | [**GetConversation**](docs/ConversationsAPI.md#getconversation) | **Get** /v2/conversations/{id} | Get a conversation with its events
*ConversationsAPI* | [**ListConversations**](docs/ConversationsAPI.md#listconversations) | **Get** /v2/conversations | List conversations for the authenticated user
*ConversationsAPI* | [**SendMessage**](docs/ConversationsAPI.md#sendmessage) | **Post** /v2/conversations/{id}/messages | Send a message and start an agentic run
*ConversationsAPI* | [**StreamEvents**](docs/ConversationsAPI.md#streamevents) | **Get** /v2/conversations/{id}/events | Stream conversation events (SSE)
*DataTypesAPI* | [**V3CopyFunctionSignatures**](docs/DataTypesAPI.md#v3copyfunctionsignatures) | **Post** /v3/analyses/{analysis_id}/signatures/copy | Copy function signatures
*DataTypesAPI* | [**V3CreateAnalysisDataTypes**](docs/DataTypesAPI.md#v3createanalysisdatatypes) | **Post** /v3/analyses/{analysis_id}/data-types | Create an analysis&#39;s data types
*DataTypesAPI* | [**V3GetAnalysisDataType**](docs/DataTypesAPI.md#v3getanalysisdatatype) | **Get** /v3/analyses/{analysis_id}/data-types/{data_type_id} | Get one of an analysis&#39;s data types
*DataTypesAPI* | [**V3GetAnalysisDataTypeHistory**](docs/DataTypesAPI.md#v3getanalysisdatatypehistory) | **Get** /v3/analyses/{analysis_id}/data-types/{data_type_id}/history | Get a data type&#39;s edit history
*DataTypesAPI* | [**V3GetFunctionSignature**](docs/DataTypesAPI.md#v3getfunctionsignature) | **Get** /v3/analyses/{analysis_id}/functions/{function_id}/signature | Get a function&#39;s signature
*DataTypesAPI* | [**V3GetFunctionSignatureHistory**](docs/DataTypesAPI.md#v3getfunctionsignaturehistory) | **Get** /v3/analyses/{analysis_id}/functions/{function_id}/signature/history | Get a function signature&#39;s edit history
*DataTypesAPI* | [**V3ListAnalysisDataTypes**](docs/DataTypesAPI.md#v3listanalysisdatatypes) | **Get** /v3/analyses/{analysis_id}/data-types | List an analysis&#39;s data types
*DataTypesAPI* | [**V3ListDataTypeFunctions**](docs/DataTypesAPI.md#v3listdatatypefunctions) | **Get** /v3/analyses/{analysis_id}/data-types/{data_type_id}/functions | List the functions using a data type
*DataTypesAPI* | [**V3ListFunctionSignatures**](docs/DataTypesAPI.md#v3listfunctionsignatures) | **Get** /v3/functions/signatures | Get signatures for many functions
*DataTypesAPI* | [**V3UpdateAnalysisDataTypes**](docs/DataTypesAPI.md#v3updateanalysisdatatypes) | **Put** /v3/analyses/{analysis_id}/data-types | Update an analysis&#39;s data types
*DataTypesAPI* | [**V3UpdateFunctionSignature**](docs/DataTypesAPI.md#v3updatefunctionsignature) | **Put** /v3/analyses/{analysis_id}/functions/{function_id}/signature | Update a function&#39;s signature
*ExternalSourcesAPI* | [**CreateExternalTaskVt**](docs/ExternalSourcesAPI.md#createexternaltaskvt) | **Post** /v2/analysis/{analysis_id}/external/vt | Pulls data from VirusTotal
*ExternalSourcesAPI* | [**GetVtData**](docs/ExternalSourcesAPI.md#getvtdata) | **Get** /v2/analysis/{analysis_id}/external/vt | Get VirusTotal data
*ExternalSourcesAPI* | [**GetVtTaskStatus**](docs/ExternalSourcesAPI.md#getvttaskstatus) | **Get** /v2/analysis/{analysis_id}/external/vt/status | Check the status of VirusTotal data retrieval
*FunctionsAIDecompilationAPI* | [**CreateAiDecompilation**](docs/FunctionsAIDecompilationAPI.md#createaidecompilation) | **Post** /v3/functions/{function_id}/ai-decompilation | Start AI decompilation
*FunctionsAIDecompilationAPI* | [**DeleteAiDecompilationInlineComment**](docs/FunctionsAIDecompilationAPI.md#deleteaidecompilationinlinecomment) | **Delete** /v3/functions/{function_id}/ai-decompilation/inline-comments/{line} | Delete a single inline comment
*FunctionsAIDecompilationAPI* | [**GetAiDecompilation**](docs/FunctionsAIDecompilationAPI.md#getaidecompilation) | **Get** /v3/functions/{function_id}/ai-decompilation | Get AI decompilation result
*FunctionsAIDecompilationAPI* | [**GetAiDecompilationInlineComments**](docs/FunctionsAIDecompilationAPI.md#getaidecompilationinlinecomments) | **Get** /v3/functions/{function_id}/ai-decompilation/inline-comments | Get AI decompilation inline comments
*FunctionsAIDecompilationAPI* | [**GetAiDecompilationInlineCommentsStatus**](docs/FunctionsAIDecompilationAPI.md#getaidecompilationinlinecommentsstatus) | **Get** /v3/functions/{function_id}/ai-decompilation/inline-comments/status | Get inline comments generation workflow status
*FunctionsAIDecompilationAPI* | [**GetAiDecompilationRating**](docs/FunctionsAIDecompilationAPI.md#getaidecompilationrating) | **Get** /v2/functions/{function_id}/ai-decompilation/rating | Get rating for AI decompilation
*FunctionsAIDecompilationAPI* | [**GetAiDecompilationStatus**](docs/FunctionsAIDecompilationAPI.md#getaidecompilationstatus) | **Get** /v3/functions/{function_id}/ai-decompilation/status | Get AI decompilation workflow status
*FunctionsAIDecompilationAPI* | [**GetAiDecompilationSummary**](docs/FunctionsAIDecompilationAPI.md#getaidecompilationsummary) | **Get** /v3/functions/{function_id}/ai-decompilation/summary | Get AI decompilation summary
*FunctionsAIDecompilationAPI* | [**GetAiDecompilationSummaryStatus**](docs/FunctionsAIDecompilationAPI.md#getaidecompilationsummarystatus) | **Get** /v3/functions/{function_id}/ai-decompilation/summary/status | Get summary generation workflow status
*FunctionsAIDecompilationAPI* | [**PatchAiDecompilationInlineComment**](docs/FunctionsAIDecompilationAPI.md#patchaidecompilationinlinecomment) | **Patch** /v3/functions/{function_id}/ai-decompilation/inline-comments | Update a single inline comment
*FunctionsAIDecompilationAPI* | [**RegenerateAiDecompilationInlineComments**](docs/FunctionsAIDecompilationAPI.md#regenerateaidecompilationinlinecomments) | **Post** /v3/functions/{function_id}/ai-decompilation/inline-comments | Regenerate AI decompilation inline comments
*FunctionsAIDecompilationAPI* | [**RegenerateAiDecompilationSummary**](docs/FunctionsAIDecompilationAPI.md#regenerateaidecompilationsummary) | **Post** /v3/functions/{function_id}/ai-decompilation/summary | Regenerate AI decompilation summary
*FunctionsAIDecompilationAPI* | [**StreamAiDecompilation**](docs/FunctionsAIDecompilationAPI.md#streamaidecompilation) | **Get** /v3/functions/{function_id}/ai-decompilation/events | Stream live AI decompilation output (SSE)
*FunctionsAIDecompilationAPI* | [**UpsertAiDecompilationRating**](docs/FunctionsAIDecompilationAPI.md#upsertaidecompilationrating) | **Patch** /v2/functions/{function_id}/ai-decompilation/rating | Upsert rating for AI decompilation
*FunctionsAIDecompilationAPI* | [**V3GetAiDecompilationLineAttributions**](docs/FunctionsAIDecompilationAPI.md#v3getaidecompilationlineattributions) | **Get** /v3/functions/{function_id}/ai-decompilation/line-attributions | Get AI decompilation line attributions
*FunctionsAIDecompilationAPI* | [**V3GetAiDecompilationTokens**](docs/FunctionsAIDecompilationAPI.md#v3getaidecompilationtokens) | **Get** /v3/functions/{function_id}/ai-decompilation/tokens | Get AI decompilation tokens and user overrides
*FunctionsAIDecompilationAPI* | [**V3GetAiDecompilationTypeSuggestions**](docs/FunctionsAIDecompilationAPI.md#v3getaidecompilationtypesuggestions) | **Get** /v3/functions/{function_id}/ai-decompilation/type-suggestions | Get AI decompilation type suggestions
*FunctionsAIDecompilationAPI* | [**V3UpsertAiDecompilationOverrides**](docs/FunctionsAIDecompilationAPI.md#v3upsertaidecompilationoverrides) | **Patch** /v3/functions/{function_id}/ai-decompilation/overrides | Upsert variable/function name overrides
*FunctionsCoreAPI* | [**AddFunctionCallee**](docs/FunctionsCoreAPI.md#addfunctioncallee) | **Post** /v3/functions/{function_id}/callees | Add a callee to a function
*FunctionsCoreAPI* | [**AddUserStringToFunction**](docs/FunctionsCoreAPI.md#adduserstringtofunction) | **Post** /v3/functions/{function_id}/user-provided-strings | Add a user-provided string to a function.
*FunctionsCoreAPI* | [**GetAnalysisStrings**](docs/FunctionsCoreAPI.md#getanalysisstrings) | **Get** /v2/analyses/{analysis_id}/functions/strings | Get string information found in the Analysis
*FunctionsCoreAPI* | [**GetAnalysisStringsStatus**](docs/FunctionsCoreAPI.md#getanalysisstringsstatus) | **Get** /v2/analyses/{analysis_id}/functions/strings/status | Get string processing state for the Analysis
*FunctionsCoreAPI* | [**GetFunctionBlocks**](docs/FunctionsCoreAPI.md#getfunctionblocks) | **Get** /v2/functions/{function_id}/blocks | Get disassembly blocks related to the function
*FunctionsCoreAPI* | [**GetFunctionBlocks_0**](docs/FunctionsCoreAPI.md#getfunctionblocks_0) | **Get** /v3/functions/{function_id}/blocks | Get function disassembly
*FunctionsCoreAPI* | [**GetFunctionCalleesCallers**](docs/FunctionsCoreAPI.md#getfunctioncalleescallers) | **Get** /v2/functions/{function_id}/callees_callers | Get list of functions that call or are called by the specified function
*FunctionsCoreAPI* | [**GetFunctionCalleesCallersBulk**](docs/FunctionsCoreAPI.md#getfunctioncalleescallersbulk) | **Get** /v2/functions/callees_callers | Get list of functions that call or are called for a list of functions
*FunctionsCoreAPI* | [**GetFunctionCalleesCallers_0**](docs/FunctionsCoreAPI.md#getfunctioncalleescallers_0) | **Get** /v3/functions/{function_id}/callees-callers | Get callees and callers for a function
*FunctionsCoreAPI* | [**GetFunctionCapabilities**](docs/FunctionsCoreAPI.md#getfunctioncapabilities) | **Get** /v2/functions/{function_id}/capabilities | Retrieve a functions capabilities
*FunctionsCoreAPI* | [**GetFunctionCapabilities_0**](docs/FunctionsCoreAPI.md#getfunctioncapabilities_0) | **Get** /v3/functions/{function_id}/capabilities | Get capabilities for a function
*FunctionsCoreAPI* | [**GetFunctionDetails**](docs/FunctionsCoreAPI.md#getfunctiondetails) | **Get** /v2/functions/{function_id} | Get function details
*FunctionsCoreAPI* | [**GetFunctionDetails_0**](docs/FunctionsCoreAPI.md#getfunctiondetails_0) | **Get** /v3/functions/{function_id} | Get function details
*FunctionsCoreAPI* | [**GetFunctionIndirectCallSites**](docs/FunctionsCoreAPI.md#getfunctionindirectcallsites) | **Get** /v3/functions/{function_id}/indirect-call-sites | Get indirect call sites for a function
*FunctionsCoreAPI* | [**GetFunctionStrings**](docs/FunctionsCoreAPI.md#getfunctionstrings) | **Get** /v2/functions/{function_id}/strings | Get string information found in the function
*FunctionsCoreAPI* | [**GetFunctionStrings_0**](docs/FunctionsCoreAPI.md#getfunctionstrings_0) | **Get** /v3/functions/{function_id}/strings | List strings for a function.
*FunctionsCoreAPI* | [**GetFunctionsCalleesCallers**](docs/FunctionsCoreAPI.md#getfunctionscalleescallers) | **Get** /v3/functions/callees-callers | Get callees and callers for many functions
*FunctionsCoreAPI* | [**GetFunctionsMatches**](docs/FunctionsCoreAPI.md#getfunctionsmatches) | **Get** /v3/functions/matches | Get function-matching results for an explicit set of functions
*FunctionsCoreAPI* | [**GetFunctionsMatchingStatus**](docs/FunctionsCoreAPI.md#getfunctionsmatchingstatus) | **Get** /v3/functions/matches/status | Get function-matching status for an explicit set of functions
*FunctionsCoreAPI* | [**GetImportedFunction**](docs/FunctionsCoreAPI.md#getimportedfunction) | **Get** /v3/analyses/{analysis_id}/imported-functions/{imported_function_id} | Get an imported function with its callers
*FunctionsCoreAPI* | [**ListAnalysisFunctions**](docs/FunctionsCoreAPI.md#listanalysisfunctions) | **Get** /v3/analyses/{analysis_id}/functions | List functions in an analysis
*FunctionsCoreAPI* | [**ListImportedFunctions**](docs/FunctionsCoreAPI.md#listimportedfunctions) | **Get** /v3/analyses/{analysis_id}/imported-functions | List imported functions in an analysis
*FunctionsCoreAPI* | [**StartFunctionsMatching**](docs/FunctionsCoreAPI.md#startfunctionsmatching) | **Post** /v3/functions/matches | Start function matching for an explicit set of functions
*FunctionsCoreAPI* | [**V3CanonicalizeFunctionNames**](docs/FunctionsCoreAPI.md#v3canonicalizefunctionnames) | **Post** /v3/functions/canonical-names | Canonicalize a batch of function names
*FunctionsRenamingHistoryAPI* | [**BatchRenameFunction**](docs/FunctionsRenamingHistoryAPI.md#batchrenamefunction) | **Post** /v2/functions/rename/batch | Batch Rename Functions
*FunctionsRenamingHistoryAPI* | [**BatchRenameFunctions**](docs/FunctionsRenamingHistoryAPI.md#batchrenamefunctions) | **Post** /v3/functions/rename | Batch rename functions
*FunctionsRenamingHistoryAPI* | [**GetFunctionHistory**](docs/FunctionsRenamingHistoryAPI.md#getfunctionhistory) | **Get** /v3/functions/{function_id}/history | Get function name history
*FunctionsRenamingHistoryAPI* | [**GetFunctionNameHistory**](docs/FunctionsRenamingHistoryAPI.md#getfunctionnamehistory) | **Get** /v2/functions/history/{function_id} | Get Function Name History
*FunctionsRenamingHistoryAPI* | [**RenameFunction**](docs/FunctionsRenamingHistoryAPI.md#renamefunction) | **Post** /v3/functions/{function_id}/rename | Rename a function
*FunctionsRenamingHistoryAPI* | [**RenameFunctionId**](docs/FunctionsRenamingHistoryAPI.md#renamefunctionid) | **Post** /v2/functions/rename/{function_id} | Rename Function
*FunctionsRenamingHistoryAPI* | [**RevertFunctionName**](docs/FunctionsRenamingHistoryAPI.md#revertfunctionname) | **Post** /v2/functions/history/{function_id}/{history_id} | Revert the function name
*FunctionsRenamingHistoryAPI* | [**RevertFunctionName_0**](docs/FunctionsRenamingHistoryAPI.md#revertfunctionname_0) | **Post** /v3/functions/{function_id}/history/{history_id}/revert | Revert function name
*IAMUsersAPI* | [**GetMe**](docs/IAMUsersAPI.md#getme) | **Get** /v2/iam/me | Get current user
*IAMUsersAPI* | [**GetMyPermissions**](docs/IAMUsersAPI.md#getmypermissions) | **Get** /v2/iam/me/permissions | Get current user permissions
*ModelsAPI* | [**GetModels**](docs/ModelsAPI.md#getmodels) | **Get** /v2/models | Gets models
*ReportsAPI* | [**CreatePdfReport**](docs/ReportsAPI.md#createpdfreport) | **Post** /v3/analyses/{analysis_id}/pdf | Start PDF report generation
*ReportsAPI* | [**DownloadPdfReport**](docs/ReportsAPI.md#downloadpdfreport) | **Get** /v3/analyses/{analysis_id}/pdf | Download generated PDF report
*ReportsAPI* | [**GetPdfReportStatus**](docs/ReportsAPI.md#getpdfreportstatus) | **Get** /v3/analyses/{analysis_id}/pdf/status | Get PDF report workflow status
*SearchAPI* | [**SearchBinaries**](docs/SearchAPI.md#searchbinaries) | **Get** /v2/search/binaries | Binaries search
*SearchAPI* | [**SearchCollections**](docs/SearchAPI.md#searchcollections) | **Get** /v2/search/collections | Collections search
*SearchAPI* | [**SearchFunctions**](docs/SearchAPI.md#searchfunctions) | **Get** /v2/search/functions | Functions search
*SearchAPI* | [**SearchTags**](docs/SearchAPI.md#searchtags) | **Get** /v2/search/tags | Tags search


## Documentation For Models

 - [APIError](docs/APIError.md)
 - [AddCalleeInputBody](docs/AddCalleeInputBody.md)
 - [AddCollectionBinariesInputBody](docs/AddCollectionBinariesInputBody.md)
 - [AddIssuerDomainInputBody](docs/AddIssuerDomainInputBody.md)
 - [AddOwnerInputBody](docs/AddOwnerInputBody.md)
 - [AddTeamMemberInputBody](docs/AddTeamMemberInputBody.md)
 - [AddUserStringInputBody](docs/AddUserStringInputBody.md)
 - [AddUserStringToFunctionInputBody](docs/AddUserStringToFunctionInputBody.md)
 - [AdditionalDetailsStatusResponse](docs/AdditionalDetailsStatusResponse.md)
 - [AiDecompilationRating](docs/AiDecompilationRating.md)
 - [AnalysisAccessInfo](docs/AnalysisAccessInfo.md)
 - [AnalysisBasicInfoOutputBody](docs/AnalysisBasicInfoOutputBody.md)
 - [AnalysisBulkAddTagsRequest](docs/AnalysisBulkAddTagsRequest.md)
 - [AnalysisBulkAddTagsResponse](docs/AnalysisBulkAddTagsResponse.md)
 - [AnalysisBulkAddTagsResponseItem](docs/AnalysisBulkAddTagsResponseItem.md)
 - [AnalysisConfig](docs/AnalysisConfig.md)
 - [AnalysisConfigSnapshot](docs/AnalysisConfigSnapshot.md)
 - [AnalysisCreateRequest](docs/AnalysisCreateRequest.md)
 - [AnalysisCreateResponse](docs/AnalysisCreateResponse.md)
 - [AnalysisDataTypesGroup](docs/AnalysisDataTypesGroup.md)
 - [AnalysisDataTypesOutputBody](docs/AnalysisDataTypesOutputBody.md)
 - [AnalysisDetailResponse](docs/AnalysisDetailResponse.md)
 - [AnalysisFunctionEntry](docs/AnalysisFunctionEntry.md)
 - [AnalysisFunctionMapping](docs/AnalysisFunctionMapping.md)
 - [AnalysisFunctions](docs/AnalysisFunctions.md)
 - [AnalysisFunctionsList](docs/AnalysisFunctionsList.md)
 - [AnalysisLogEntry](docs/AnalysisLogEntry.md)
 - [AnalysisLogMessage](docs/AnalysisLogMessage.md)
 - [AnalysisLogs](docs/AnalysisLogs.md)
 - [AnalysisRecord](docs/AnalysisRecord.md)
 - [AnalysisRecordBody](docs/AnalysisRecordBody.md)
 - [AnalysisReport](docs/AnalysisReport.md)
 - [AnalysisRequirement](docs/AnalysisRequirement.md)
 - [AnalysisScope](docs/AnalysisScope.md)
 - [AnalysisStringFunction](docs/AnalysisStringFunction.md)
 - [AnalysisStringInput](docs/AnalysisStringInput.md)
 - [AnalysisStringItem](docs/AnalysisStringItem.md)
 - [AnalysisStringsResponse](docs/AnalysisStringsResponse.md)
 - [AnalysisStringsStatusResponse](docs/AnalysisStringsStatusResponse.md)
 - [AnalysisTagBody](docs/AnalysisTagBody.md)
 - [AnalysisTags](docs/AnalysisTags.md)
 - [AnalysisUpdateRequest](docs/AnalysisUpdateRequest.md)
 - [AnalysisUpdateTagsRequest](docs/AnalysisUpdateTagsRequest.md)
 - [AnalysisUpdateTagsResponse](docs/AnalysisUpdateTagsResponse.md)
 - [ApiCall](docs/ApiCall.md)
 - [ApiCombinationEvidence](docs/ApiCombinationEvidence.md)
 - [AppApiRestV2AgentSchemaCapability](docs/AppApiRestV2AgentSchemaCapability.md)
 - [AppApiRestV2AnalysesEnumsOrderBy](docs/AppApiRestV2AnalysesEnumsOrderBy.md)
 - [AppApiRestV2CollectionsEnumsOrderBy](docs/AppApiRestV2CollectionsEnumsOrderBy.md)
 - [AppApiRestV2FunctionsResponsesFunction](docs/AppApiRestV2FunctionsResponsesFunction.md)
 - [AppApiRestV2FunctionsTypesFunction](docs/AppApiRestV2FunctionsTypesFunction.md)
 - [AppApiRestV2InfoTypesCapability](docs/AppApiRestV2InfoTypesCapability.md)
 - [ArchiveContentEntry](docs/ArchiveContentEntry.md)
 - [ArrayDataType](docs/ArrayDataType.md)
 - [ArrayDefinition](docs/ArrayDefinition.md)
 - [Artifact](docs/Artifact.md)
 - [AttemptFailedEvent](docs/AttemptFailedEvent.md)
 - [AttemptStartedEvent](docs/AttemptStartedEvent.md)
 - [AutoRunAgents](docs/AutoRunAgents.md)
 - [AutoUnstripStatusOutputBody](docs/AutoUnstripStatusOutputBody.md)
 - [BaseDataType](docs/BaseDataType.md)
 - [BaseResponse](docs/BaseResponse.md)
 - [BaseResponseAdditionalDetailsStatusResponse](docs/BaseResponseAdditionalDetailsStatusResponse.md)
 - [BaseResponseAnalysisBulkAddTagsResponse](docs/BaseResponseAnalysisBulkAddTagsResponse.md)
 - [BaseResponseAnalysisCreateResponse](docs/BaseResponseAnalysisCreateResponse.md)
 - [BaseResponseAnalysisDetailResponse](docs/BaseResponseAnalysisDetailResponse.md)
 - [BaseResponseAnalysisFunctionMapping](docs/BaseResponseAnalysisFunctionMapping.md)
 - [BaseResponseAnalysisFunctions](docs/BaseResponseAnalysisFunctions.md)
 - [BaseResponseAnalysisFunctionsList](docs/BaseResponseAnalysisFunctionsList.md)
 - [BaseResponseAnalysisStringsResponse](docs/BaseResponseAnalysisStringsResponse.md)
 - [BaseResponseAnalysisStringsStatusResponse](docs/BaseResponseAnalysisStringsStatusResponse.md)
 - [BaseResponseAnalysisTags](docs/BaseResponseAnalysisTags.md)
 - [BaseResponseAnalysisUpdateTagsResponse](docs/BaseResponseAnalysisUpdateTagsResponse.md)
 - [BaseResponseBasic](docs/BaseResponseBasic.md)
 - [BaseResponseBinariesRelatedStatusResponse](docs/BaseResponseBinariesRelatedStatusResponse.md)
 - [BaseResponseBinaryAdditionalResponse](docs/BaseResponseBinaryAdditionalResponse.md)
 - [BaseResponseBinaryDetailsResponse](docs/BaseResponseBinaryDetailsResponse.md)
 - [BaseResponseBinaryExternalsResponse](docs/BaseResponseBinaryExternalsResponse.md)
 - [BaseResponseBinarySearchResponse](docs/BaseResponseBinarySearchResponse.md)
 - [BaseResponseBool](docs/BaseResponseBool.md)
 - [BaseResponseCalleesCallerFunctionsResponse](docs/BaseResponseCalleesCallerFunctionsResponse.md)
 - [BaseResponseCapabilities](docs/BaseResponseCapabilities.md)
 - [BaseResponseCapabilitiesAgentResponse](docs/BaseResponseCapabilitiesAgentResponse.md)
 - [BaseResponseChildBinariesResponse](docs/BaseResponseChildBinariesResponse.md)
 - [BaseResponseCollectionBinariesUpdateResponse](docs/BaseResponseCollectionBinariesUpdateResponse.md)
 - [BaseResponseCollectionResponse](docs/BaseResponseCollectionResponse.md)
 - [BaseResponseCollectionSearchResponse](docs/BaseResponseCollectionSearchResponse.md)
 - [BaseResponseCollectionTagsUpdateResponse](docs/BaseResponseCollectionTagsUpdateResponse.md)
 - [BaseResponseCommentResponse](docs/BaseResponseCommentResponse.md)
 - [BaseResponseConfigResponse](docs/BaseResponseConfigResponse.md)
 - [BaseResponseCreated](docs/BaseResponseCreated.md)
 - [BaseResponseDict](docs/BaseResponseDict.md)
 - [BaseResponseExternalResponse](docs/BaseResponseExternalResponse.md)
 - [BaseResponseFunctionBlocksResponse](docs/BaseResponseFunctionBlocksResponse.md)
 - [BaseResponseFunctionCapabilityResponse](docs/BaseResponseFunctionCapabilityResponse.md)
 - [BaseResponseFunctionSearchResponse](docs/BaseResponseFunctionSearchResponse.md)
 - [BaseResponseFunctionStringsResponse](docs/BaseResponseFunctionStringsResponse.md)
 - [BaseResponseFunctionsDetailResponse](docs/BaseResponseFunctionsDetailResponse.md)
 - [BaseResponseGetPublicUserResponse](docs/BaseResponseGetPublicUserResponse.md)
 - [BaseResponseListCalleesCallerFunctionsResponse](docs/BaseResponseListCalleesCallerFunctionsResponse.md)
 - [BaseResponseListCollectionResults](docs/BaseResponseListCollectionResults.md)
 - [BaseResponseListCommentResponse](docs/BaseResponseListCommentResponse.md)
 - [BaseResponseListDieMatch](docs/BaseResponseListDieMatch.md)
 - [BaseResponseListFunctionNameHistory](docs/BaseResponseListFunctionNameHistory.md)
 - [BaseResponseListUserActivityResponse](docs/BaseResponseListUserActivityResponse.md)
 - [BaseResponseLogs](docs/BaseResponseLogs.md)
 - [BaseResponseModelsResponse](docs/BaseResponseModelsResponse.md)
 - [BaseResponseParams](docs/BaseResponseParams.md)
 - [BaseResponseProtocolsAgentResponse](docs/BaseResponseProtocolsAgentResponse.md)
 - [BaseResponseQueuedWorkflowTaskResponse](docs/BaseResponseQueuedWorkflowTaskResponse.md)
 - [BaseResponseRecent](docs/BaseResponseRecent.md)
 - [BaseResponseRemediationAgentResponse](docs/BaseResponseRemediationAgentResponse.md)
 - [BaseResponseReportAnalysisResponse](docs/BaseResponseReportAnalysisResponse.md)
 - [BaseResponseSecretsAgentResponse](docs/BaseResponseSecretsAgentResponse.md)
 - [BaseResponseStatus](docs/BaseResponseStatus.md)
 - [BaseResponseStr](docs/BaseResponseStr.md)
 - [BaseResponseTagSearchResponse](docs/BaseResponseTagSearchResponse.md)
 - [BaseResponseTaskResponse](docs/BaseResponseTaskResponse.md)
 - [BaseResponseTaskStatusResponse](docs/BaseResponseTaskStatusResponse.md)
 - [BaseResponseTriageReportResponse](docs/BaseResponseTriageReportResponse.md)
 - [BaseResponseUnionGetAiDecompilationRatingResponseNoneType](docs/BaseResponseUnionGetAiDecompilationRatingResponseNoneType.md)
 - [BaseResponseUploadResponse](docs/BaseResponseUploadResponse.md)
 - [BaseResponseXrefResponse](docs/BaseResponseXrefResponse.md)
 - [Basic](docs/Basic.md)
 - [BatchBinaryMatchResult](docs/BatchBinaryMatchResult.md)
 - [BatchFunctionSignatureEntry](docs/BatchFunctionSignatureEntry.md)
 - [BatchMatchingOutputBody](docs/BatchMatchingOutputBody.md)
 - [BatchRenameInputBody](docs/BatchRenameInputBody.md)
 - [BatchRenameItem](docs/BatchRenameItem.md)
 - [BatchRenameOutputBody](docs/BatchRenameOutputBody.md)
 - [BinariesRelatedStatusResponse](docs/BinariesRelatedStatusResponse.md)
 - [BinariesTaskStatus](docs/BinariesTaskStatus.md)
 - [Binary](docs/Binary.md)
 - [BinaryAdditionalDetailsDataResponse](docs/BinaryAdditionalDetailsDataResponse.md)
 - [BinaryAdditionalResponse](docs/BinaryAdditionalResponse.md)
 - [BinaryConfig](docs/BinaryConfig.md)
 - [BinaryDetailsResponse](docs/BinaryDetailsResponse.md)
 - [BinaryExternalsResponse](docs/BinaryExternalsResponse.md)
 - [BinarySearchResponse](docs/BinarySearchResponse.md)
 - [BinarySearchResult](docs/BinarySearchResult.md)
 - [BinaryTaskStatus](docs/BinaryTaskStatus.md)
 - [BitfieldDataType](docs/BitfieldDataType.md)
 - [BulkCreateUserResult](docs/BulkCreateUserResult.md)
 - [BulkCreateUsersOutputBody](docs/BulkCreateUsersOutputBody.md)
 - [BulkDeleteAnalysesRequest](docs/BulkDeleteAnalysesRequest.md)
 - [BytesConstant](docs/BytesConstant.md)
 - [CallChain](docs/CallChain.md)
 - [CallChainEvidence](docs/CallChainEvidence.md)
 - [CallEdge](docs/CallEdge.md)
 - [CallEdgesOutputBody](docs/CallEdgesOutputBody.md)
 - [CalleeFunctionInfo](docs/CalleeFunctionInfo.md)
 - [CalleesCallerFunctionsResponse](docs/CalleesCallerFunctionsResponse.md)
 - [CallerFunctionInfo](docs/CallerFunctionInfo.md)
 - [CanonicalName](docs/CanonicalName.md)
 - [CanonicalizeNamesInputBody](docs/CanonicalizeNamesInputBody.md)
 - [CanonicalizeNamesOutputBody](docs/CanonicalizeNamesOutputBody.md)
 - [Capabilities](docs/Capabilities.md)
 - [CapabilitiesAgentResponse](docs/CapabilitiesAgentResponse.md)
 - [CapabilitiesOutputBody](docs/CapabilitiesOutputBody.md)
 - [CapabilitiesResult](docs/CapabilitiesResult.md)
 - [Capability](docs/Capability.md)
 - [CapabilityEntry](docs/CapabilityEntry.md)
 - [ChildBinariesResponse](docs/ChildBinariesResponse.md)
 - [CodeSignatureModel](docs/CodeSignatureModel.md)
 - [CollectionBinariesUpdateRequest](docs/CollectionBinariesUpdateRequest.md)
 - [CollectionBinariesUpdateResponse](docs/CollectionBinariesUpdateResponse.md)
 - [CollectionBinaryResponse](docs/CollectionBinaryResponse.md)
 - [CollectionCreateRequest](docs/CollectionCreateRequest.md)
 - [CollectionListItem](docs/CollectionListItem.md)
 - [CollectionListItemBody](docs/CollectionListItemBody.md)
 - [CollectionResponse](docs/CollectionResponse.md)
 - [CollectionResponseBinariesInner](docs/CollectionResponseBinariesInner.md)
 - [CollectionScope](docs/CollectionScope.md)
 - [CollectionSearchResponse](docs/CollectionSearchResponse.md)
 - [CollectionSearchResult](docs/CollectionSearchResult.md)
 - [CollectionTagsUpdateRequest](docs/CollectionTagsUpdateRequest.md)
 - [CollectionTagsUpdateResponse](docs/CollectionTagsUpdateResponse.md)
 - [CollectionUpdateRequest](docs/CollectionUpdateRequest.md)
 - [CommentBase](docs/CommentBase.md)
 - [CommentResponse](docs/CommentResponse.md)
 - [CommentUpdateRequest](docs/CommentUpdateRequest.md)
 - [CommentsData](docs/CommentsData.md)
 - [Config](docs/Config.md)
 - [ConfigResponse](docs/ConfigResponse.md)
 - [ConfirmToolInputBody](docs/ConfirmToolInputBody.md)
 - [Connection](docs/Connection.md)
 - [ConsoleOutputEntry](docs/ConsoleOutputEntry.md)
 - [Context](docs/Context.md)
 - [Conversation](docs/Conversation.md)
 - [ConversationContext](docs/ConversationContext.md)
 - [ConversationWithEvents](docs/ConversationWithEvents.md)
 - [CopyFunctionSignaturesInputBody](docs/CopyFunctionSignaturesInputBody.md)
 - [CopyFunctionSignaturesOutputBody](docs/CopyFunctionSignaturesOutputBody.md)
 - [CopySignatureItem](docs/CopySignatureItem.md)
 - [CreateAIDecompOutputBody](docs/CreateAIDecompOutputBody.md)
 - [CreateAnalysisDataTypesInputBody](docs/CreateAnalysisDataTypesInputBody.md)
 - [CreateArrayDataType](docs/CreateArrayDataType.md)
 - [CreateBaseDataType](docs/CreateBaseDataType.md)
 - [CreateBitfieldDataType](docs/CreateBitfieldDataType.md)
 - [CreateCheckoutSessionInputBody](docs/CreateCheckoutSessionInputBody.md)
 - [CreateCollectionInputBody](docs/CreateCollectionInputBody.md)
 - [CreateCollectionOutputBody](docs/CreateCollectionOutputBody.md)
 - [CreateConversationRequest](docs/CreateConversationRequest.md)
 - [CreateDataTypeEntry](docs/CreateDataTypeEntry.md)
 - [CreateEnumDataType](docs/CreateEnumDataType.md)
 - [CreateFunctionDataType](docs/CreateFunctionDataType.md)
 - [CreateGroupInputBody](docs/CreateGroupInputBody.md)
 - [CreateIdentityInputBody](docs/CreateIdentityInputBody.md)
 - [CreateIssuerInputBody](docs/CreateIssuerInputBody.md)
 - [CreateMetadata](docs/CreateMetadata.md)
 - [CreateOrganisationInputBody](docs/CreateOrganisationInputBody.md)
 - [CreatePointerDataType](docs/CreatePointerDataType.md)
 - [CreatePortalSessionInputBody](docs/CreatePortalSessionInputBody.md)
 - [CreateRequest](docs/CreateRequest.md)
 - [CreateResult](docs/CreateResult.md)
 - [CreateStructDataType](docs/CreateStructDataType.md)
 - [CreateTeamInputBody](docs/CreateTeamInputBody.md)
 - [CreateTypedefDataType](docs/CreateTypedefDataType.md)
 - [CreateUnionDataType](docs/CreateUnionDataType.md)
 - [CreateUnknownDataType](docs/CreateUnknownDataType.md)
 - [CreateUserInputBody](docs/CreateUserInputBody.md)
 - [Created](docs/Created.md)
 - [CryptoCall](docs/CryptoCall.md)
 - [CryptoDirectMatch](docs/CryptoDirectMatch.md)
 - [CryptoExplainMetadata](docs/CryptoExplainMetadata.md)
 - [CryptoExplainResult](docs/CryptoExplainResult.md)
 - [CryptoExplainedFunction](docs/CryptoExplainedFunction.md)
 - [CryptoFinding](docs/CryptoFinding.md)
 - [CryptoScanMetadata](docs/CryptoScanMetadata.md)
 - [CryptoScanResult](docs/CryptoScanResult.md)
 - [DataTypeEntry](docs/DataTypeEntry.md)
 - [DataTypeEnumValueEntry](docs/DataTypeEnumValueEntry.md)
 - [DataTypeFunctionEntry](docs/DataTypeFunctionEntry.md)
 - [DataTypeFunctionParameterEntry](docs/DataTypeFunctionParameterEntry.md)
 - [DataTypeMemberEntry](docs/DataTypeMemberEntry.md)
 - [DataTypeVersion](docs/DataTypeVersion.md)
 - [DecompFailedEvent](docs/DecompFailedEvent.md)
 - [DecompFinishedEvent](docs/DecompFinishedEvent.md)
 - [DecompilationCommentContext](docs/DecompilationCommentContext.md)
 - [DecompilationData](docs/DecompilationData.md)
 - [DecompilerSummary](docs/DecompilerSummary.md)
 - [DecompilerSummaryEvidence](docs/DecompilerSummaryEvidence.md)
 - [DieMatch](docs/DieMatch.md)
 - [DisassemblyOutputBody](docs/DisassemblyOutputBody.md)
 - [Display](docs/Display.md)
 - [DnsQuery](docs/DnsQuery.md)
 - [DrakvufFileMetadata](docs/DrakvufFileMetadata.md)
 - [DynamicExecutionStatus](docs/DynamicExecutionStatus.md)
 - [DynamicExecutionStatusResponse](docs/DynamicExecutionStatusResponse.md)
 - [ELFImportModel](docs/ELFImportModel.md)
 - [ELFModel](docs/ELFModel.md)
 - [ELFRelocation](docs/ELFRelocation.md)
 - [ELFSection](docs/ELFSection.md)
 - [ELFSecurity](docs/ELFSecurity.md)
 - [ELFSegment](docs/ELFSegment.md)
 - [ELFSymbol](docs/ELFSymbol.md)
 - [ElfDynamicEntry](docs/ElfDynamicEntry.md)
 - [Endianness](docs/Endianness.md)
 - [EntrypointModel](docs/EntrypointModel.md)
 - [EnumDataType](docs/EnumDataType.md)
 - [EnumDefinition](docs/EnumDefinition.md)
 - [ErrorBody](docs/ErrorBody.md)
 - [ErrorModel](docs/ErrorModel.md)
 - [Event](docs/Event.md)
 - [EventAttemptFailed](docs/EventAttemptFailed.md)
 - [EventAttemptStarted](docs/EventAttemptStarted.md)
 - [EventCONTEXTCOMPACTED](docs/EventCONTEXTCOMPACTED.md)
 - [EventDecompFailed](docs/EventDecompFailed.md)
 - [EventDecompFinished](docs/EventDecompFinished.md)
 - [EventNamesFinished](docs/EventNamesFinished.md)
 - [EventProse](docs/EventProse.md)
 - [EventRUNCANCELLED](docs/EventRUNCANCELLED.md)
 - [EventRUNERROR](docs/EventRUNERROR.md)
 - [EventRUNFINISHED](docs/EventRUNFINISHED.md)
 - [EventRUNSTARTED](docs/EventRUNSTARTED.md)
 - [EventRenameApplied](docs/EventRenameApplied.md)
 - [EventSTEPFINISHED](docs/EventSTEPFINISHED.md)
 - [EventSTEPSTARTED](docs/EventSTEPSTARTED.md)
 - [EventSourceDelta](docs/EventSourceDelta.md)
 - [EventSourceReset](docs/EventSourceReset.md)
 - [EventTEXTMESSAGECONTENT](docs/EventTEXTMESSAGECONTENT.md)
 - [EventTEXTMESSAGEEND](docs/EventTEXTMESSAGEEND.md)
 - [EventTEXTMESSAGESTART](docs/EventTEXTMESSAGESTART.md)
 - [EventTITLEUPDATED](docs/EventTITLEUPDATED.md)
 - [EventTOOLCALLARGSDELTA](docs/EventTOOLCALLARGSDELTA.md)
 - [EventTOOLCALLEND](docs/EventTOOLCALLEND.md)
 - [EventTOOLCALLPROGRESS](docs/EventTOOLCALLPROGRESS.md)
 - [EventTOOLCALLRESULT](docs/EventTOOLCALLRESULT.md)
 - [EventTOOLCALLSTART](docs/EventTOOLCALLSTART.md)
 - [EventTOOLCONFIRMATIONREQUIRED](docs/EventTOOLCONFIRMATIONREQUIRED.md)
 - [EventTypesSuggested](docs/EventTypesSuggested.md)
 - [EventWarning](docs/EventWarning.md)
 - [EvidenceEffect](docs/EvidenceEffect.md)
 - [EvidenceInner](docs/EvidenceInner.md)
 - [EvidenceStrength](docs/EvidenceStrength.md)
 - [Example](docs/Example.md)
 - [ExecutionCall](docs/ExecutionCall.md)
 - [ExecutionDirectMatch](docs/ExecutionDirectMatch.md)
 - [ExecutionExplainMetadata](docs/ExecutionExplainMetadata.md)
 - [ExecutionExplainResult](docs/ExecutionExplainResult.md)
 - [ExecutionExplainedFunction](docs/ExecutionExplainedFunction.md)
 - [ExecutionFinding](docs/ExecutionFinding.md)
 - [ExecutionScanMetadata](docs/ExecutionScanMetadata.md)
 - [ExecutionScanResult](docs/ExecutionScanResult.md)
 - [ExportModel](docs/ExportModel.md)
 - [ExternalResponse](docs/ExternalResponse.md)
 - [ExtractedBinary](docs/ExtractedBinary.md)
 - [ExtractedURL](docs/ExtractedURL.md)
 - [ExtractionFailure](docs/ExtractionFailure.md)
 - [FeedbackOutputBody](docs/FeedbackOutputBody.md)
 - [FileActivityEntry](docs/FileActivityEntry.md)
 - [FileFormat](docs/FileFormat.md)
 - [FileHashes](docs/FileHashes.md)
 - [FileMetadata](docs/FileMetadata.md)
 - [FilesystemAnalyseMetadata](docs/FilesystemAnalyseMetadata.md)
 - [FilesystemAnalyseResult](docs/FilesystemAnalyseResult.md)
 - [FilesystemCall](docs/FilesystemCall.md)
 - [FilesystemDirectMatch](docs/FilesystemDirectMatch.md)
 - [FilesystemExplainedFunction](docs/FilesystemExplainedFunction.md)
 - [FilesystemFinding](docs/FilesystemFinding.md)
 - [FilesystemScanMetadata](docs/FilesystemScanMetadata.md)
 - [FilesystemScanResult](docs/FilesystemScanResult.md)
 - [Filters](docs/Filters.md)
 - [Finding](docs/Finding.md)
 - [FormFile](docs/FormFile.md)
 - [FunctionBlockDestinationResponse](docs/FunctionBlockDestinationResponse.md)
 - [FunctionBlockResponse](docs/FunctionBlockResponse.md)
 - [FunctionBlocksResponse](docs/FunctionBlocksResponse.md)
 - [FunctionBoundary](docs/FunctionBoundary.md)
 - [FunctionCallEdges](docs/FunctionCallEdges.md)
 - [FunctionCapabilityResponse](docs/FunctionCapabilityResponse.md)
 - [FunctionDataType](docs/FunctionDataType.md)
 - [FunctionDetailsOutputBody](docs/FunctionDetailsOutputBody.md)
 - [FunctionListItem](docs/FunctionListItem.md)
 - [FunctionLocalVariableResponse](docs/FunctionLocalVariableResponse.md)
 - [FunctionMapping](docs/FunctionMapping.md)
 - [FunctionMatch](docs/FunctionMatch.md)
 - [FunctionNameHistory](docs/FunctionNameHistory.md)
 - [FunctionParamResponse](docs/FunctionParamResponse.md)
 - [FunctionRename](docs/FunctionRename.md)
 - [FunctionRenameMap](docs/FunctionRenameMap.md)
 - [FunctionSearchResponse](docs/FunctionSearchResponse.md)
 - [FunctionSearchResult](docs/FunctionSearchResult.md)
 - [FunctionSignatureBody](docs/FunctionSignatureBody.md)
 - [FunctionSignatureEntry](docs/FunctionSignatureEntry.md)
 - [FunctionSignatureVersion](docs/FunctionSignatureVersion.md)
 - [FunctionSimilarity](docs/FunctionSimilarity.md)
 - [FunctionSimilarityEvidence](docs/FunctionSimilarityEvidence.md)
 - [FunctionSourceType](docs/FunctionSourceType.md)
 - [FunctionString](docs/FunctionString.md)
 - [FunctionStringItem](docs/FunctionStringItem.md)
 - [FunctionStringsResponse](docs/FunctionStringsResponse.md)
 - [FunctionTypeDefinition](docs/FunctionTypeDefinition.md)
 - [FunctionsDetailResponse](docs/FunctionsDetailResponse.md)
 - [FunctionsListRename](docs/FunctionsListRename.md)
 - [FunctionsProgressOutputBody](docs/FunctionsProgressOutputBody.md)
 - [GeneratePDFOutputBody](docs/GeneratePDFOutputBody.md)
 - [GetAdditionalDetailsOutputBody](docs/GetAdditionalDetailsOutputBody.md)
 - [GetAdditionalDetailsStatusOutputBody](docs/GetAdditionalDetailsStatusOutputBody.md)
 - [GetAiDecompilationRatingResponse](docs/GetAiDecompilationRatingResponse.md)
 - [GetAnalysisLogsOutputBody](docs/GetAnalysisLogsOutputBody.md)
 - [GetAnalysisStringsStatusOutputBody](docs/GetAnalysisStringsStatusOutputBody.md)
 - [GetCollectionOutputBody](docs/GetCollectionOutputBody.md)
 - [GetConfigOutputBody](docs/GetConfigOutputBody.md)
 - [GetDataTypeHistoryBody](docs/GetDataTypeHistoryBody.md)
 - [GetDieInfoOutputBody](docs/GetDieInfoOutputBody.md)
 - [GetFunctionSignatureHistoryBody](docs/GetFunctionSignatureHistoryBody.md)
 - [GetMatchesOutputBody](docs/GetMatchesOutputBody.md)
 - [GetMatchesStatusOutputBody](docs/GetMatchesStatusOutputBody.md)
 - [GetModelsOutputBody](docs/GetModelsOutputBody.md)
 - [GetProductsOutputBody](docs/GetProductsOutputBody.md)
 - [GetPublicUserResponse](docs/GetPublicUserResponse.md)
 - [GetRelatedBinariesOutputBody](docs/GetRelatedBinariesOutputBody.md)
 - [GetRelatedStatusOutputBody](docs/GetRelatedStatusOutputBody.md)
 - [GetSubscriptionOutputBody](docs/GetSubscriptionOutputBody.md)
 - [GetTokensResponse](docs/GetTokensResponse.md)
 - [HardcodedSecretEvidence](docs/HardcodedSecretEvidence.md)
 - [HistoryActor](docs/HistoryActor.md)
 - [HistoryEntry](docs/HistoryEntry.md)
 - [HttpRequest](docs/HttpRequest.md)
 - [IOC](docs/IOC.md)
 - [ISA](docs/ISA.md)
 - [IconModel](docs/IconModel.md)
 - [ImportModel](docs/ImportModel.md)
 - [ImportedApi](docs/ImportedApi.md)
 - [ImportedApiCall](docs/ImportedApiCall.md)
 - [ImportedApiCallEvidence](docs/ImportedApiCallEvidence.md)
 - [ImportedFunctionCallerEntry](docs/ImportedFunctionCallerEntry.md)
 - [ImportedFunctionDetailOutputBody](docs/ImportedFunctionDetailOutputBody.md)
 - [ImportedFunctionEntry](docs/ImportedFunctionEntry.md)
 - [IndirectCallSite](docs/IndirectCallSite.md)
 - [IndirectCallSitesOutputBody](docs/IndirectCallSitesOutputBody.md)
 - [InlineComment](docs/InlineComment.md)
 - [InputBody](docs/InputBody.md)
 - [InsertAnalysisLogRequest](docs/InsertAnalysisLogRequest.md)
 - [InviteUserInputBody](docs/InviteUserInputBody.md)
 - [IssuerAllowedDomain](docs/IssuerAllowedDomain.md)
 - [KnownConstantEvidence](docs/KnownConstantEvidence.md)
 - [LineAttributionsData](docs/LineAttributionsData.md)
 - [ListAnalysesOutputBody](docs/ListAnalysesOutputBody.md)
 - [ListAnalysisDataTypesOutputBody](docs/ListAnalysisDataTypesOutputBody.md)
 - [ListAnalysisFunctionsOutputBody](docs/ListAnalysisFunctionsOutputBody.md)
 - [ListAnalysisStringsOutputBody](docs/ListAnalysisStringsOutputBody.md)
 - [ListArchiveContentsOutputBody](docs/ListArchiveContentsOutputBody.md)
 - [ListCollectionResults](docs/ListCollectionResults.md)
 - [ListCollectionsOutputBody](docs/ListCollectionsOutputBody.md)
 - [ListDataTypeFunctionsBody](docs/ListDataTypeFunctionsBody.md)
 - [ListExampleAnalysesOutputBody](docs/ListExampleAnalysesOutputBody.md)
 - [ListFunctionSignaturesOutputBody](docs/ListFunctionSignaturesOutputBody.md)
 - [ListFunctionStringsOutputBody](docs/ListFunctionStringsOutputBody.md)
 - [ListImportedFunctionsOutputBody](docs/ListImportedFunctionsOutputBody.md)
 - [ListTeamsOutputBody](docs/ListTeamsOutputBody.md)
 - [ListUsersOutputBody](docs/ListUsersOutputBody.md)
 - [LocationOutputBody](docs/LocationOutputBody.md)
 - [Logs](docs/Logs.md)
 - [MITRETechnique](docs/MITRETechnique.md)
 - [MatchFilters](docs/MatchFilters.md)
 - [MatchedFunction](docs/MatchedFunction.md)
 - [MemdumpEntry](docs/MemdumpEntry.md)
 - [MessageBody](docs/MessageBody.md)
 - [Meta](docs/Meta.md)
 - [MetaModel](docs/MetaModel.md)
 - [Metadata](docs/Metadata.md)
 - [ModelInterpretation](docs/ModelInterpretation.md)
 - [ModelInterpretationEvidence](docs/ModelInterpretationEvidence.md)
 - [ModelName](docs/ModelName.md)
 - [ModelsResponse](docs/ModelsResponse.md)
 - [ModuleLoadEntry](docs/ModuleLoadEntry.md)
 - [MutexEntry](docs/MutexEntry.md)
 - [NameConfidence](docs/NameConfidence.md)
 - [NameSourceType](docs/NameSourceType.md)
 - [NamesFinishedEvent](docs/NamesFinishedEvent.md)
 - [NetworkActivity](docs/NetworkActivity.md)
 - [NetworkingCall](docs/NetworkingCall.md)
 - [NetworkingDirectMatch](docs/NetworkingDirectMatch.md)
 - [NetworkingExplainMetadata](docs/NetworkingExplainMetadata.md)
 - [NetworkingExplainResult](docs/NetworkingExplainResult.md)
 - [NetworkingExplainedFunction](docs/NetworkingExplainedFunction.md)
 - [NetworkingFinding](docs/NetworkingFinding.md)
 - [NetworkingScanMetadata](docs/NetworkingScanMetadata.md)
 - [NetworkingScanResult](docs/NetworkingScanResult.md)
 - [OIDCCallbackInputBody](docs/OIDCCallbackInputBody.md)
 - [OperationCreateMetadataCreateResult](docs/OperationCreateMetadataCreateResult.md)
 - [OperationCryptoExplainMetadataCryptoExplainResult](docs/OperationCryptoExplainMetadataCryptoExplainResult.md)
 - [OperationCryptoScanMetadataCryptoScanResult](docs/OperationCryptoScanMetadataCryptoScanResult.md)
 - [OperationExecutionExplainMetadataExecutionExplainResult](docs/OperationExecutionExplainMetadataExecutionExplainResult.md)
 - [OperationExecutionScanMetadataExecutionScanResult](docs/OperationExecutionScanMetadataExecutionScanResult.md)
 - [OperationFilesystemAnalyseMetadataFilesystemAnalyseResult](docs/OperationFilesystemAnalyseMetadataFilesystemAnalyseResult.md)
 - [OperationFilesystemScanMetadataFilesystemScanResult](docs/OperationFilesystemScanMetadataFilesystemScanResult.md)
 - [OperationMetadataCapabilitiesResult](docs/OperationMetadataCapabilitiesResult.md)
 - [OperationMetadataRemediationResult](docs/OperationMetadataRemediationResult.md)
 - [OperationMetadataReportResult](docs/OperationMetadataReportResult.md)
 - [OperationMetadataThreatReportResult](docs/OperationMetadataThreatReportResult.md)
 - [OperationMetadataTriageResult](docs/OperationMetadataTriageResult.md)
 - [OperationNetworkingExplainMetadataNetworkingExplainResult](docs/OperationNetworkingExplainMetadataNetworkingExplainResult.md)
 - [OperationNetworkingScanMetadataNetworkingScanResult](docs/OperationNetworkingScanMetadataNetworkingScanResult.md)
 - [OperationSecurityScanMetadataSecurityScanResult](docs/OperationSecurityScanMetadataSecurityScanResult.md)
 - [OperationWorkflowProgressResultBody](docs/OperationWorkflowProgressResultBody.md)
 - [Order](docs/Order.md)
 - [Organisation](docs/Organisation.md)
 - [OrganisationGroup](docs/OrganisationGroup.md)
 - [OrganisationIssuer](docs/OrganisationIssuer.md)
 - [OrganisationOwner](docs/OrganisationOwner.md)
 - [PDBDebugModel](docs/PDBDebugModel.md)
 - [PEModel](docs/PEModel.md)
 - [PaginationModel](docs/PaginationModel.md)
 - [Params](docs/Params.md)
 - [PasswordResetInputBody](docs/PasswordResetInputBody.md)
 - [PatchCollectionBinariesInputBody](docs/PatchCollectionBinariesInputBody.md)
 - [PatchCollectionBinariesOutputBody](docs/PatchCollectionBinariesOutputBody.md)
 - [PatchCollectionInputBody](docs/PatchCollectionInputBody.md)
 - [PatchCollectionOutputBody](docs/PatchCollectionOutputBody.md)
 - [PatchCollectionTagsInputBody](docs/PatchCollectionTagsInputBody.md)
 - [PatchCollectionTagsOutputBody](docs/PatchCollectionTagsOutputBody.md)
 - [PatchCommentBody](docs/PatchCommentBody.md)
 - [PcapBodyInfo](docs/PcapBodyInfo.md)
 - [Permissions](docs/Permissions.md)
 - [Platform](docs/Platform.md)
 - [PointerDataType](docs/PointerDataType.md)
 - [PointerDefinition](docs/PointerDefinition.md)
 - [PriceOutput](docs/PriceOutput.md)
 - [PriceSummary](docs/PriceSummary.md)
 - [ProcessActivityEntry](docs/ProcessActivityEntry.md)
 - [ProcessMemdumps](docs/ProcessMemdumps.md)
 - [ProcessNode](docs/ProcessNode.md)
 - [ProcessTree](docs/ProcessTree.md)
 - [ProductOutput](docs/ProductOutput.md)
 - [ProductSummary](docs/ProductSummary.md)
 - [ProgressMessage](docs/ProgressMessage.md)
 - [ProseEvent](docs/ProseEvent.md)
 - [ProtocolsAgentResponse](docs/ProtocolsAgentResponse.md)
 - [PutAnalysisStringsRequest](docs/PutAnalysisStringsRequest.md)
 - [QueuedWorkflowTaskResponse](docs/QueuedWorkflowTaskResponse.md)
 - [ReAnalysisForm](docs/ReAnalysisForm.md)
 - [Recent](docs/Recent.md)
 - [ReferencedConstant](docs/ReferencedConstant.md)
 - [ReferencedConstantEvidence](docs/ReferencedConstantEvidence.md)
 - [RefreshBody](docs/RefreshBody.md)
 - [RegenerateOutputBody](docs/RegenerateOutputBody.md)
 - [RegisterUserInputBody](docs/RegisterUserInputBody.md)
 - [RegistryOperation](docs/RegistryOperation.md)
 - [RelatedBinary](docs/RelatedBinary.md)
 - [RelativeBinaryResponse](docs/RelativeBinaryResponse.md)
 - [RemediationAgentResponse](docs/RemediationAgentResponse.md)
 - [RemediationResult](docs/RemediationResult.md)
 - [RemoveCollectionBinariesInputBody](docs/RemoveCollectionBinariesInputBody.md)
 - [RenameAppliedEvent](docs/RenameAppliedEvent.md)
 - [RenameInputBody](docs/RenameInputBody.md)
 - [RenameOutputBody](docs/RenameOutputBody.md)
 - [RenameUnnamedFunctionsResult](docs/RenameUnnamedFunctionsResult.md)
 - [RenderedToken](docs/RenderedToken.md)
 - [ReportAnalysisResponse](docs/ReportAnalysisResponse.md)
 - [ReportEvent](docs/ReportEvent.md)
 - [ReportInfo](docs/ReportInfo.md)
 - [ReportOptions](docs/ReportOptions.md)
 - [ReportReachabilityStatus](docs/ReportReachabilityStatus.md)
 - [ReportResult](docs/ReportResult.md)
 - [ResolvedEntity](docs/ResolvedEntity.md)
 - [ResultBody](docs/ResultBody.md)
 - [RevokeBody](docs/RevokeBody.md)
 - [RuleKind](docs/RuleKind.md)
 - [SSOProvider](docs/SSOProvider.md)
 - [SSOProvidersOutputBody](docs/SSOProvidersOutputBody.md)
 - [SandboxConfig](docs/SandboxConfig.md)
 - [SandboxOptions](docs/SandboxOptions.md)
 - [SandboxStartMethod](docs/SandboxStartMethod.md)
 - [SandboxTimeout](docs/SandboxTimeout.md)
 - [ScheduledTaskEntry](docs/ScheduledTaskEntry.md)
 - [ScrapeThirdPartyConfig](docs/ScrapeThirdPartyConfig.md)
 - [ScreenshotEntry](docs/ScreenshotEntry.md)
 - [ScreenshotsIndex](docs/ScreenshotsIndex.md)
 - [SecretsAgentResponse](docs/SecretsAgentResponse.md)
 - [SectionModel](docs/SectionModel.md)
 - [SecurityFinding](docs/SecurityFinding.md)
 - [SecurityModel](docs/SecurityModel.md)
 - [SecurityScanMetadata](docs/SecurityScanMetadata.md)
 - [SecurityScanResult](docs/SecurityScanResult.md)
 - [SegmentInfo](docs/SegmentInfo.md)
 - [SendMessageRequest](docs/SendMessageRequest.md)
 - [ServerSentEventsInner](docs/ServerSentEventsInner.md)
 - [ServerSentEventsInner1](docs/ServerSentEventsInner1.md)
 - [ServiceEntry](docs/ServiceEntry.md)
 - [SessionOutputBody](docs/SessionOutputBody.md)
 - [SignatureParameterEntry](docs/SignatureParameterEntry.md)
 - [SignatureParameterInput](docs/SignatureParameterInput.md)
 - [SignatureStorageEntry](docs/SignatureStorageEntry.md)
 - [SignatureStorageInput](docs/SignatureStorageInput.md)
 - [SingleCodeCertificateModel](docs/SingleCodeCertificateModel.md)
 - [SingleCodeSignatureModel](docs/SingleCodeSignatureModel.md)
 - [SinglePDBEntryModel](docs/SinglePDBEntryModel.md)
 - [SingleSectionModel](docs/SingleSectionModel.md)
 - [SourceDeltaEvent](docs/SourceDeltaEvent.md)
 - [SourceResetEvent](docs/SourceResetEvent.md)
 - [SseEventContextCompactedData](docs/SseEventContextCompactedData.md)
 - [SseEventRunCancelledData](docs/SseEventRunCancelledData.md)
 - [SseEventRunErrorData](docs/SseEventRunErrorData.md)
 - [SseEventRunFinishedData](docs/SseEventRunFinishedData.md)
 - [SseEventRunStartedData](docs/SseEventRunStartedData.md)
 - [SseEventStepFinishedData](docs/SseEventStepFinishedData.md)
 - [SseEventStepStartedData](docs/SseEventStepStartedData.md)
 - [SseEventTextMessageContentData](docs/SseEventTextMessageContentData.md)
 - [SseEventTextMessageEndData](docs/SseEventTextMessageEndData.md)
 - [SseEventTextMessageStartData](docs/SseEventTextMessageStartData.md)
 - [SseEventTitleUpdatedData](docs/SseEventTitleUpdatedData.md)
 - [SseEventToolCallArgsDeltaData](docs/SseEventToolCallArgsDeltaData.md)
 - [SseEventToolCallEndData](docs/SseEventToolCallEndData.md)
 - [SseEventToolCallProgressData](docs/SseEventToolCallProgressData.md)
 - [SseEventToolCallResultData](docs/SseEventToolCallResultData.md)
 - [SseEventToolCallStartData](docs/SseEventToolCallStartData.md)
 - [SseEventToolConfirmationRequiredData](docs/SseEventToolConfirmationRequiredData.md)
 - [StartBatchMatchingInputBody](docs/StartBatchMatchingInputBody.md)
 - [StartMatchingForAnalysisInputBody](docs/StartMatchingForAnalysisInputBody.md)
 - [StartMatchingForFunctionsInputBody](docs/StartMatchingForFunctionsInputBody.md)
 - [StartMatchingOutputBody](docs/StartMatchingOutputBody.md)
 - [StartupInfo](docs/StartupInfo.md)
 - [Status](docs/Status.md)
 - [StatusBody](docs/StatusBody.md)
 - [StatusInput](docs/StatusInput.md)
 - [StatusOutput](docs/StatusOutput.md)
 - [StatusResponse](docs/StatusResponse.md)
 - [StringFunctions](docs/StringFunctions.md)
 - [StringSource](docs/StringSource.md)
 - [StructDataType](docs/StructDataType.md)
 - [StructDefinition](docs/StructDefinition.md)
 - [Subject](docs/Subject.md)
 - [SubjectAnyOf](docs/SubjectAnyOf.md)
 - [SubjectAnyOf1](docs/SubjectAnyOf1.md)
 - [SubjectAnyOf2](docs/SubjectAnyOf2.md)
 - [SubjectAnyOf3](docs/SubjectAnyOf3.md)
 - [SubmitFeedbackInputBody](docs/SubmitFeedbackInputBody.md)
 - [SubmitUserFeedbackRequest](docs/SubmitUserFeedbackRequest.md)
 - [SuggestedHole](docs/SuggestedHole.md)
 - [SuggestedMemberView](docs/SuggestedMemberView.md)
 - [SuggestedTypeView](docs/SuggestedTypeView.md)
 - [SummaryData](docs/SummaryData.md)
 - [SuspiciousString](docs/SuspiciousString.md)
 - [SuspiciousStringEvidence](docs/SuspiciousStringEvidence.md)
 - [Symbols](docs/Symbols.md)
 - [Tag](docs/Tag.md)
 - [TagItem](docs/TagItem.md)
 - [TagResponse](docs/TagResponse.md)
 - [TagSearchResponse](docs/TagSearchResponse.md)
 - [TagSearchResult](docs/TagSearchResult.md)
 - [TaskResponse](docs/TaskResponse.md)
 - [TaskStatus](docs/TaskStatus.md)
 - [TaskStatusResponse](docs/TaskStatusResponse.md)
 - [TcpCarvedFile](docs/TcpCarvedFile.md)
 - [Team](docs/Team.md)
 - [TeamMember](docs/TeamMember.md)
 - [Technique](docs/Technique.md)
 - [ThreatReportResult](docs/ThreatReportResult.md)
 - [TimestampModel](docs/TimestampModel.md)
 - [Token](docs/Token.md)
 - [TokenInputBody](docs/TokenInputBody.md)
 - [TokenResponse](docs/TokenResponse.md)
 - [TokenisedData](docs/TokenisedData.md)
 - [TriageFunction](docs/TriageFunction.md)
 - [TriageFunctionResponse](docs/TriageFunctionResponse.md)
 - [TriageReportResponse](docs/TriageReportResponse.md)
 - [TriageResult](docs/TriageResult.md)
 - [TriggerCryptoScanInputBody](docs/TriggerCryptoScanInputBody.md)
 - [TriggerDynamicExecutionInputBody](docs/TriggerDynamicExecutionInputBody.md)
 - [TriggerExecutionExplainInputBody](docs/TriggerExecutionExplainInputBody.md)
 - [TriggerExecutionScanInputBody](docs/TriggerExecutionScanInputBody.md)
 - [TriggerFilesystemAnalyseInputBody](docs/TriggerFilesystemAnalyseInputBody.md)
 - [TriggerFilesystemScanInputBody](docs/TriggerFilesystemScanInputBody.md)
 - [TriggerNetworkingExplainInputBody](docs/TriggerNetworkingExplainInputBody.md)
 - [TriggerNetworkingScanInputBody](docs/TriggerNetworkingScanInputBody.md)
 - [TriggerRenameUnnamedFunctionsInputBody](docs/TriggerRenameUnnamedFunctionsInputBody.md)
 - [TriggerSecurityScanInputBody](docs/TriggerSecurityScanInputBody.md)
 - [Ttp](docs/Ttp.md)
 - [TypeSuggestionsData](docs/TypeSuggestionsData.md)
 - [TypedefDataType](docs/TypedefDataType.md)
 - [TypedefDefinition](docs/TypedefDefinition.md)
 - [TypesSuggestedEvent](docs/TypesSuggestedEvent.md)
 - [UnionDataType](docs/UnionDataType.md)
 - [UnionDefinition](docs/UnionDefinition.md)
 - [UnknownDataType](docs/UnknownDataType.md)
 - [UpdateAnalysisDataTypesInputBody](docs/UpdateAnalysisDataTypesInputBody.md)
 - [UpdateArrayDataType](docs/UpdateArrayDataType.md)
 - [UpdateBaseDataType](docs/UpdateBaseDataType.md)
 - [UpdateBitfieldDataType](docs/UpdateBitfieldDataType.md)
 - [UpdateDataTypeEntry](docs/UpdateDataTypeEntry.md)
 - [UpdateEnumDataType](docs/UpdateEnumDataType.md)
 - [UpdateFunctionDataType](docs/UpdateFunctionDataType.md)
 - [UpdateFunctionSignatureInputBody](docs/UpdateFunctionSignatureInputBody.md)
 - [UpdateIssuerInputBody](docs/UpdateIssuerInputBody.md)
 - [UpdateOrganisationInputBody](docs/UpdateOrganisationInputBody.md)
 - [UpdatePasswordInputBody](docs/UpdatePasswordInputBody.md)
 - [UpdatePointerDataType](docs/UpdatePointerDataType.md)
 - [UpdateProfileInputBody](docs/UpdateProfileInputBody.md)
 - [UpdateStructDataType](docs/UpdateStructDataType.md)
 - [UpdateTeamInputBody](docs/UpdateTeamInputBody.md)
 - [UpdateTypedefDataType](docs/UpdateTypedefDataType.md)
 - [UpdateUnionDataType](docs/UpdateUnionDataType.md)
 - [UpdateUnknownDataType](docs/UpdateUnknownDataType.md)
 - [UpdateUserCreditsInputBody](docs/UpdateUserCreditsInputBody.md)
 - [UpdateUserInputBody](docs/UpdateUserInputBody.md)
 - [UpdateUserPasswordInputBody](docs/UpdateUserPasswordInputBody.md)
 - [UpgradeAnalysisModelOutputBody](docs/UpgradeAnalysisModelOutputBody.md)
 - [UploadFileType](docs/UploadFileType.md)
 - [UploadOutputBody](docs/UploadOutputBody.md)
 - [UploadResponse](docs/UploadResponse.md)
 - [UpsertAiDecomplationRatingRequest](docs/UpsertAiDecomplationRatingRequest.md)
 - [UpsertOverridesData](docs/UpsertOverridesData.md)
 - [UpsertOverridesInputBody](docs/UpsertOverridesInputBody.md)
 - [User](docs/User.md)
 - [UserActivityResponse](docs/UserActivityResponse.md)
 - [UserCredits](docs/UserCredits.md)
 - [UserIdentity](docs/UserIdentity.md)
 - [UserProfile](docs/UserProfile.md)
 - [WarningEvent](docs/WarningEvent.md)
 - [WorkflowProgress](docs/WorkflowProgress.md)
 - [Workspace](docs/Workspace.md)
 - [XrefFromResponse](docs/XrefFromResponse.md)
 - [XrefResponse](docs/XrefResponse.md)
 - [XrefToResponse](docs/XrefToResponse.md)


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`
