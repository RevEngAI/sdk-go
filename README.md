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
*AgentAPI* | [**CheckReportAnalysisTaskStatusV2AnalysesAnalysisIdAgentReportAnalysisStatusGet**](docs/AgentAPI.md#checkreportanalysistaskstatusv2analysesanalysisidagentreportanalysisstatusget) | **Get** /v2/analyses/{analysis_id}/agent/report-analysis/status | Check the status of a report analysis workflow
*AgentAPI* | [**CheckTriageTaskStatusV2AnalysesAnalysisIdAgentTriageStatusGet**](docs/AgentAPI.md#checktriagetaskstatusv2analysesanalysisidagenttriagestatusget) | **Get** /v2/analyses/{analysis_id}/agent/triage/status | Check the status of a triage analysis workflow
*AgentAPI* | [**CreateCapabilitiesTaskV2AnalysesAnalysisIdAgentCapabilitiesPost**](docs/AgentAPI.md#createcapabilitiestaskv2analysesanalysisidagentcapabilitiespost) | **Post** /v2/analyses/{analysis_id}/agent/capabilities | Queues a capabilities analysis workflow process
*AgentAPI* | [**CreateReportAnalysisTaskV2AnalysesAnalysisIdAgentReportAnalysisPost**](docs/AgentAPI.md#createreportanalysistaskv2analysesanalysisidagentreportanalysispost) | **Post** /v2/analyses/{analysis_id}/agent/report-analysis | Queues a combined report analysis workflow process
*AgentAPI* | [**CreateTriageTaskV2AnalysesAnalysisIdAgentTriagePost**](docs/AgentAPI.md#createtriagetaskv2analysesanalysisidagenttriagepost) | **Post** /v2/analyses/{analysis_id}/agent/triage | Queues a triage analysis workflow process
*AgentAPI* | [**GetCapabilitiesResultV2AnalysesAnalysisIdAgentCapabilitiesGet**](docs/AgentAPI.md#getcapabilitiesresultv2analysesanalysisidagentcapabilitiesget) | **Get** /v2/analyses/{analysis_id}/agent/capabilities | Get Capabilities Result
*AgentAPI* | [**GetReportAnalysisResultV2AnalysesAnalysisIdAgentReportAnalysisGet**](docs/AgentAPI.md#getreportanalysisresultv2analysesanalysisidagentreportanalysisget) | **Get** /v2/analyses/{analysis_id}/agent/report-analysis | Get Report Analysis Result
*AgentAPI* | [**GetTriageResultV2AnalysesAnalysisIdAgentTriageGet**](docs/AgentAPI.md#gettriageresultv2analysesanalysisidagenttriageget) | **Get** /v2/analyses/{analysis_id}/agent/triage | Get Triage Result
*AnalysesBulkActionsAPI* | [**BulkAddAnalysisTags**](docs/AnalysesBulkActionsAPI.md#bulkaddanalysistags) | **Patch** /v2/analyses/tags/add | Bulk Add Analysis Tags
*AnalysesBulkActionsAPI* | [**BulkDeleteAnalyses**](docs/AnalysesBulkActionsAPI.md#bulkdeleteanalyses) | **Patch** /v2/analyses/delete | Bulk Delete Analyses
*AnalysesCommentsAPI* | [**CreateAnalysisComment**](docs/AnalysesCommentsAPI.md#createanalysiscomment) | **Post** /v2/analyses/{analysis_id}/comments | Create a comment for this analysis
*AnalysesCommentsAPI* | [**DeleteAnalysisComment**](docs/AnalysesCommentsAPI.md#deleteanalysiscomment) | **Delete** /v2/analyses/{analysis_id}/comments/{comment_id} | Delete a comment
*AnalysesCommentsAPI* | [**GetAnalysisComments**](docs/AnalysesCommentsAPI.md#getanalysiscomments) | **Get** /v2/analyses/{analysis_id}/comments | Get comments for this analysis
*AnalysesCommentsAPI* | [**UpdateAnalysisComment**](docs/AnalysesCommentsAPI.md#updateanalysiscomment) | **Patch** /v2/analyses/{analysis_id}/comments/{comment_id} | Update a comment
*AnalysesCoreAPI* | [**CreateAnalysis**](docs/AnalysesCoreAPI.md#createanalysis) | **Post** /v2/analyses | Create Analysis
*AnalysesCoreAPI* | [**DeleteAnalysis**](docs/AnalysesCoreAPI.md#deleteanalysis) | **Delete** /v2/analyses/{analysis_id} | Delete Analysis
*AnalysesCoreAPI* | [**GetAnalysisBasicInfo**](docs/AnalysesCoreAPI.md#getanalysisbasicinfo) | **Get** /v2/analyses/{analysis_id}/basic | Gets basic analysis information
*AnalysesCoreAPI* | [**GetAnalysisFunctionMap**](docs/AnalysesCoreAPI.md#getanalysisfunctionmap) | **Get** /v2/analyses/{analysis_id}/func_maps | Get Analysis Function Map
*AnalysesCoreAPI* | [**GetAnalysisLogs**](docs/AnalysesCoreAPI.md#getanalysislogs) | **Get** /v2/analyses/{analysis_id}/logs | Gets the logs of an analysis
*AnalysesCoreAPI* | [**GetAnalysisParams**](docs/AnalysesCoreAPI.md#getanalysisparams) | **Get** /v2/analyses/{analysis_id}/params | Gets analysis param information
*AnalysesCoreAPI* | [**GetAnalysisQueuePosition**](docs/AnalysesCoreAPI.md#getanalysisqueueposition) | **Get** /v2/analyses/{analysis_id}/queue-position | Get the queue position of an analysis
*AnalysesCoreAPI* | [**GetAnalysisStatus**](docs/AnalysesCoreAPI.md#getanalysisstatus) | **Get** /v2/analyses/{analysis_id}/status | Gets the status of an analysis
*AnalysesCoreAPI* | [**InsertAnalysisLog**](docs/AnalysesCoreAPI.md#insertanalysislog) | **Post** /v2/analyses/{analysis_id}/logs | Insert a log entry for an analysis
*AnalysesCoreAPI* | [**ListAnalyses**](docs/AnalysesCoreAPI.md#listanalyses) | **Get** /v2/analyses/list | Gets the most recent analyses
*AnalysesCoreAPI* | [**LookupBinaryId**](docs/AnalysesCoreAPI.md#lookupbinaryid) | **Get** /v2/analyses/lookup/{binary_id} | Gets the analysis ID from binary ID
*AnalysesCoreAPI* | [**PutAnalysisStrings**](docs/AnalysesCoreAPI.md#putanalysisstrings) | **Put** /v2/analyses/{analysis_id}/strings | Add strings to the analysis
*AnalysesCoreAPI* | [**RequeueAnalysis**](docs/AnalysesCoreAPI.md#requeueanalysis) | **Post** /v2/analyses/{analysis_id}/requeue | Requeue Analysis
*AnalysesCoreAPI* | [**UpdateAnalysis**](docs/AnalysesCoreAPI.md#updateanalysis) | **Patch** /v2/analyses/{analysis_id} | Update Analysis
*AnalysesCoreAPI* | [**UpdateAnalysisTags**](docs/AnalysesCoreAPI.md#updateanalysistags) | **Patch** /v2/analyses/{analysis_id}/tags | Update Analysis Tags
*AnalysesCoreAPI* | [**UploadFile**](docs/AnalysesCoreAPI.md#uploadfile) | **Post** /v2/upload | Upload File
*AnalysesResultsMetadataAPI* | [**GetAnalysisFunctionsPaginated**](docs/AnalysesResultsMetadataAPI.md#getanalysisfunctionspaginated) | **Get** /v2/analyses/{analysis_id}/functions | Get functions from analysis
*AnalysesResultsMetadataAPI* | [**GetCapabilities**](docs/AnalysesResultsMetadataAPI.md#getcapabilities) | **Get** /v2/analyses/{analysis_id}/capabilities | Gets the capabilities from the analysis
*AnalysesResultsMetadataAPI* | [**GetFunctionsList**](docs/AnalysesResultsMetadataAPI.md#getfunctionslist) | **Get** /v2/analyses/{analysis_id}/functions/list | Gets functions from analysis
*AnalysesResultsMetadataAPI* | [**GetPdf**](docs/AnalysesResultsMetadataAPI.md#getpdf) | **Get** /v2/analyses/{analysis_id}/pdf | Gets the PDF found in the analysis
*AnalysesResultsMetadataAPI* | [**GetSbom**](docs/AnalysesResultsMetadataAPI.md#getsbom) | **Get** /v2/analyses/{analysis_id}/sbom | Gets the software-bill-of-materials (SBOM) found in the analysis
*AnalysesResultsMetadataAPI* | [**GetTags**](docs/AnalysesResultsMetadataAPI.md#gettags) | **Get** /v2/analyses/{analysis_id}/tags | Get function tags with maliciousness score
*AnalysesResultsMetadataAPI* | [**GetVulnerabilities**](docs/AnalysesResultsMetadataAPI.md#getvulnerabilities) | **Get** /v2/analyses/{analysis_id}/vulnerabilities | Gets the vulnerabilities found in the analysis
*AnalysesXRefsAPI* | [**GetXrefByVaddr**](docs/AnalysesXRefsAPI.md#getxrefbyvaddr) | **Get** /v2/analyses/{analysis_id}/xrefs/{vaddr} | [Beta] Look up xrefs by virtual address
*AuthenticationUsersAPI* | [**GetRequesterUserInfo**](docs/AuthenticationUsersAPI.md#getrequesteruserinfo) | **Get** /v2/users/me | Get the requesters user information
*AuthenticationUsersAPI* | [**GetUser**](docs/AuthenticationUsersAPI.md#getuser) | **Get** /v2/users/{user_id} | Get a user&#39;s public information
*AuthenticationUsersAPI* | [**GetUserActivity**](docs/AuthenticationUsersAPI.md#getuseractivity) | **Get** /v2/users/activity | Get auth user activity
*AuthenticationUsersAPI* | [**GetUserComments**](docs/AuthenticationUsersAPI.md#getusercomments) | **Get** /v2/users/me/comments | Get comments by user
*AuthenticationUsersAPI* | [**SubmitUserFeedback**](docs/AuthenticationUsersAPI.md#submituserfeedback) | **Post** /v2/users/feedback | Submit feedback about the application
*BinariesAPI* | [**DownloadZippedBinary**](docs/BinariesAPI.md#downloadzippedbinary) | **Get** /v2/binaries/{binary_id}/download-zipped | Downloads a zipped binary with password protection
*BinariesAPI* | [**GetBinaryAdditionalDetails**](docs/BinariesAPI.md#getbinaryadditionaldetails) | **Get** /v2/binaries/{binary_id}/additional-details | Gets the additional details of a binary
*BinariesAPI* | [**GetBinaryAdditionalDetailsStatus**](docs/BinariesAPI.md#getbinaryadditionaldetailsstatus) | **Get** /v2/binaries/{binary_id}/additional-details/status | Gets the status of the additional details task for a binary
*BinariesAPI* | [**GetBinaryDetails**](docs/BinariesAPI.md#getbinarydetails) | **Get** /v2/binaries/{binary_id}/details | Gets the details of a binary
*BinariesAPI* | [**GetBinaryDieInfo**](docs/BinariesAPI.md#getbinarydieinfo) | **Get** /v2/binaries/{binary_id}/die-info | Gets the die info of a binary
*BinariesAPI* | [**GetBinaryExternals**](docs/BinariesAPI.md#getbinaryexternals) | **Get** /v2/binaries/{binary_id}/externals | Gets the external details of a binary
*BinariesAPI* | [**GetBinaryRelatedStatus**](docs/BinariesAPI.md#getbinaryrelatedstatus) | **Get** /v2/binaries/{binary_id}/related/status | Gets the status of the unpack binary task for a binary
*BinariesAPI* | [**GetRelatedBinaries**](docs/BinariesAPI.md#getrelatedbinaries) | **Get** /v2/binaries/{binary_id}/related | Gets the related binaries of a binary.
*CollectionsAPI* | [**CreateCollection**](docs/CollectionsAPI.md#createcollection) | **Post** /v2/collections | Creates new collection information
*CollectionsAPI* | [**DeleteCollection**](docs/CollectionsAPI.md#deletecollection) | **Delete** /v2/collections/{collection_id} | Deletes a collection
*CollectionsAPI* | [**GetCollection**](docs/CollectionsAPI.md#getcollection) | **Get** /v2/collections/{collection_id} | Returns a collection
*CollectionsAPI* | [**ListCollections**](docs/CollectionsAPI.md#listcollections) | **Get** /v2/collections | Gets basic collections information
*CollectionsAPI* | [**UpdateCollection**](docs/CollectionsAPI.md#updatecollection) | **Patch** /v2/collections/{collection_id} | Updates a collection
*CollectionsAPI* | [**UpdateCollectionBinaries**](docs/CollectionsAPI.md#updatecollectionbinaries) | **Patch** /v2/collections/{collection_id}/binaries | Updates a collection binaries
*CollectionsAPI* | [**UpdateCollectionTags**](docs/CollectionsAPI.md#updatecollectiontags) | **Patch** /v2/collections/{collection_id}/tags | Updates a collection tags
*ConfigAPI* | [**GetConfig**](docs/ConfigAPI.md#getconfig) | **Get** /v2/config | Get Config
*ConversationsAPI* | [**CancelRun**](docs/ConversationsAPI.md#cancelrun) | **Post** /v2/conversations/{id}/cancel | Cancel an active run
*ConversationsAPI* | [**ConfirmTool**](docs/ConversationsAPI.md#confirmtool) | **Post** /v2/conversations/{id}/confirm | Approve or reject a pending tool confirmation
*ConversationsAPI* | [**CreateConversation**](docs/ConversationsAPI.md#createconversation) | **Post** /v2/conversations | Create a new conversation
*ConversationsAPI* | [**GetConversation**](docs/ConversationsAPI.md#getconversation) | **Get** /v2/conversations/{id} | Get a conversation with its events
*ConversationsAPI* | [**ListConversations**](docs/ConversationsAPI.md#listconversations) | **Get** /v2/conversations | List conversations for the authenticated user
*ConversationsAPI* | [**SendMessage**](docs/ConversationsAPI.md#sendmessage) | **Post** /v2/conversations/{id}/messages | Send a message and start an agentic run
*ConversationsAPI* | [**StreamEvents**](docs/ConversationsAPI.md#streamevents) | **Get** /v2/conversations/{id}/events | Stream conversation events (SSE)
*ExternalSourcesAPI* | [**CreateExternalTaskVt**](docs/ExternalSourcesAPI.md#createexternaltaskvt) | **Post** /v2/analysis/{analysis_id}/external/vt | Pulls data from VirusTotal
*ExternalSourcesAPI* | [**GetVtData**](docs/ExternalSourcesAPI.md#getvtdata) | **Get** /v2/analysis/{analysis_id}/external/vt | Get VirusTotal data
*ExternalSourcesAPI* | [**GetVtTaskStatus**](docs/ExternalSourcesAPI.md#getvttaskstatus) | **Get** /v2/analysis/{analysis_id}/external/vt/status | Check the status of VirusTotal data retrieval
*FirmwareAPI* | [**GetBinariesForFirmwareTask**](docs/FirmwareAPI.md#getbinariesforfirmwaretask) | **Get** /v2/firmware/get-binaries/{task_id} | Upload firmware for unpacking
*FirmwareAPI* | [**UploadFirmware**](docs/FirmwareAPI.md#uploadfirmware) | **Post** /v2/firmware | Upload firmware for unpacking
*FunctionsAIDecompilationAPI* | [**CreateAiDecompilation**](docs/FunctionsAIDecompilationAPI.md#createaidecompilation) | **Post** /v3/functions/{function_id}/ai-decompilation | Start AI decompilation
*FunctionsAIDecompilationAPI* | [**CreateAiDecompilationComment**](docs/FunctionsAIDecompilationAPI.md#createaidecompilationcomment) | **Post** /v2/functions/{function_id}/ai-decompilation/comments | Create a comment for this function
*FunctionsAIDecompilationAPI* | [**CreateAiDecompilationTask**](docs/FunctionsAIDecompilationAPI.md#createaidecompilationtask) | **Post** /v2/functions/{function_id}/ai-decompilation | Begins AI Decompilation Process
*FunctionsAIDecompilationAPI* | [**DeleteAiDecompilationComment**](docs/FunctionsAIDecompilationAPI.md#deleteaidecompilationcomment) | **Delete** /v2/functions/{function_id}/ai-decompilation/comments/{comment_id} | Delete a comment
*FunctionsAIDecompilationAPI* | [**GetAiDecompilation**](docs/FunctionsAIDecompilationAPI.md#getaidecompilation) | **Get** /v3/functions/{function_id}/ai-decompilation | Get AI decompilation result
*FunctionsAIDecompilationAPI* | [**GetAiDecompilationComments**](docs/FunctionsAIDecompilationAPI.md#getaidecompilationcomments) | **Get** /v2/functions/{function_id}/ai-decompilation/comments | Get comments for this function
*FunctionsAIDecompilationAPI* | [**GetAiDecompilationInlineComments**](docs/FunctionsAIDecompilationAPI.md#getaidecompilationinlinecomments) | **Get** /v3/functions/{function_id}/ai-decompilation/inline-comments | Get AI decompilation inline comments
*FunctionsAIDecompilationAPI* | [**GetAiDecompilationInlineCommentsStatus**](docs/FunctionsAIDecompilationAPI.md#getaidecompilationinlinecommentsstatus) | **Get** /v3/functions/{function_id}/ai-decompilation/inline-comments/status | Get inline comments generation workflow status
*FunctionsAIDecompilationAPI* | [**GetAiDecompilationRating**](docs/FunctionsAIDecompilationAPI.md#getaidecompilationrating) | **Get** /v2/functions/{function_id}/ai-decompilation/rating | Get rating for AI decompilation
*FunctionsAIDecompilationAPI* | [**GetAiDecompilationStatus**](docs/FunctionsAIDecompilationAPI.md#getaidecompilationstatus) | **Get** /v3/functions/{function_id}/ai-decompilation/status | Get AI decompilation workflow status
*FunctionsAIDecompilationAPI* | [**GetAiDecompilationSummary**](docs/FunctionsAIDecompilationAPI.md#getaidecompilationsummary) | **Get** /v3/functions/{function_id}/ai-decompilation/summary | Get AI decompilation summary
*FunctionsAIDecompilationAPI* | [**GetAiDecompilationSummaryStatus**](docs/FunctionsAIDecompilationAPI.md#getaidecompilationsummarystatus) | **Get** /v3/functions/{function_id}/ai-decompilation/summary/status | Get summary generation workflow status
*FunctionsAIDecompilationAPI* | [**GetAiDecompilationTaskResult**](docs/FunctionsAIDecompilationAPI.md#getaidecompilationtaskresult) | **Get** /v2/functions/{function_id}/ai-decompilation | Polls AI Decompilation Process
*FunctionsAIDecompilationAPI* | [**GetAiDecompilationTaskStatus**](docs/FunctionsAIDecompilationAPI.md#getaidecompilationtaskstatus) | **Get** /v2/functions/{function_id}/ai-decompilation/status | Check the status of a function ai decompilation
*FunctionsAIDecompilationAPI* | [**GetAiDecompilationTokenised**](docs/FunctionsAIDecompilationAPI.md#getaidecompilationtokenised) | **Get** /v3/functions/{function_id}/ai-decompilation/tokenised | Get tokenised AI decompilation with function mapping
*FunctionsAIDecompilationAPI* | [**RegenerateAiDecompilationInlineComments**](docs/FunctionsAIDecompilationAPI.md#regenerateaidecompilationinlinecomments) | **Post** /v3/functions/{function_id}/ai-decompilation/inline-comments | Regenerate AI decompilation inline comments
*FunctionsAIDecompilationAPI* | [**RegenerateAiDecompilationSummary**](docs/FunctionsAIDecompilationAPI.md#regenerateaidecompilationsummary) | **Post** /v3/functions/{function_id}/ai-decompilation/summary | Regenerate AI decompilation summary
*FunctionsAIDecompilationAPI* | [**UpdateAiDecompilationComment**](docs/FunctionsAIDecompilationAPI.md#updateaidecompilationcomment) | **Patch** /v2/functions/{function_id}/ai-decompilation/comments/{comment_id} | Update a comment
*FunctionsAIDecompilationAPI* | [**UpsertAiDecompilationOverrides**](docs/FunctionsAIDecompilationAPI.md#upsertaidecompilationoverrides) | **Patch** /v3/functions/{function_id}/ai-decompilation/overrides | Upsert variable/function name overrides
*FunctionsAIDecompilationAPI* | [**UpsertAiDecompilationRating**](docs/FunctionsAIDecompilationAPI.md#upsertaidecompilationrating) | **Patch** /v2/functions/{function_id}/ai-decompilation/rating | Upsert rating for AI decompilation
*FunctionsCoreAPI* | [**AiUnstrip**](docs/FunctionsCoreAPI.md#aiunstrip) | **Post** /v2/analyses/{analysis_id}/functions/ai-unstrip | Performs matching and auto-unstrip for an analysis and its functions
*FunctionsCoreAPI* | [**AnalysisFunctionMatching**](docs/FunctionsCoreAPI.md#analysisfunctionmatching) | **Post** /v2/analyses/{analysis_id}/functions/matches | Perform matching for the functions of an analysis
*FunctionsCoreAPI* | [**AutoUnstrip**](docs/FunctionsCoreAPI.md#autounstrip) | **Post** /v2/analyses/{analysis_id}/functions/auto-unstrip | Performs matching and auto-unstrip for an analysis and its functions
*FunctionsCoreAPI* | [**BatchFunctionMatching**](docs/FunctionsCoreAPI.md#batchfunctionmatching) | **Post** /v2/functions/matches | Perform function matching for an arbitrary batch of functions, binaries or collections
*FunctionsCoreAPI* | [**CancelAiUnstrip**](docs/FunctionsCoreAPI.md#cancelaiunstrip) | **Delete** /v2/analyses/{analysis_id}/functions/ai-unstrip/cancel | Cancels a running ai-unstrip
*FunctionsCoreAPI* | [**CancelAutoUnstrip**](docs/FunctionsCoreAPI.md#cancelautounstrip) | **Delete** /v2/analyses/{analysis_id}/functions/unstrip/cancel | Cancels a running auto-unstrip
*FunctionsCoreAPI* | [**GetAnalysisStrings**](docs/FunctionsCoreAPI.md#getanalysisstrings) | **Get** /v2/analyses/{analysis_id}/functions/strings | Get string information found in the Analysis
*FunctionsCoreAPI* | [**GetAnalysisStringsStatus**](docs/FunctionsCoreAPI.md#getanalysisstringsstatus) | **Get** /v2/analyses/{analysis_id}/functions/strings/status | Get string processing state for the Analysis
*FunctionsCoreAPI* | [**GetFunctionBlocks**](docs/FunctionsCoreAPI.md#getfunctionblocks) | **Get** /v2/functions/{function_id}/blocks | Get disassembly blocks related to the function
*FunctionsCoreAPI* | [**GetFunctionCalleesCallers**](docs/FunctionsCoreAPI.md#getfunctioncalleescallers) | **Get** /v2/functions/{function_id}/callees_callers | Get list of functions that call or are called by the specified function
*FunctionsCoreAPI* | [**GetFunctionCalleesCallersBulk**](docs/FunctionsCoreAPI.md#getfunctioncalleescallersbulk) | **Get** /v2/functions/callees_callers | Get list of functions that call or are called for a list of functions
*FunctionsCoreAPI* | [**GetFunctionCapabilities**](docs/FunctionsCoreAPI.md#getfunctioncapabilities) | **Get** /v2/functions/{function_id}/capabilities | Retrieve a functions capabilities
*FunctionsCoreAPI* | [**GetFunctionDetails**](docs/FunctionsCoreAPI.md#getfunctiondetails) | **Get** /v2/functions/{function_id} | Get function details
*FunctionsCoreAPI* | [**GetFunctionStrings**](docs/FunctionsCoreAPI.md#getfunctionstrings) | **Get** /v2/functions/{function_id}/strings | Get string information found in the function
*FunctionsDataTypesAPI* | [**GenerateFunctionDataTypesForAnalysis**](docs/FunctionsDataTypesAPI.md#generatefunctiondatatypesforanalysis) | **Post** /v2/analyses/{analysis_id}/functions/data_types | Generate Function Data Types
*FunctionsDataTypesAPI* | [**GenerateFunctionDataTypesForFunctions**](docs/FunctionsDataTypesAPI.md#generatefunctiondatatypesforfunctions) | **Post** /v2/functions/data_types | Generate Function Data Types for an arbitrary list of functions
*FunctionsDataTypesAPI* | [**GetFunctionDataTypes**](docs/FunctionsDataTypesAPI.md#getfunctiondatatypes) | **Get** /v2/analyses/{analysis_id}/functions/{function_id}/data_types | Get Function Data Types
*FunctionsDataTypesAPI* | [**ListFunctionDataTypesForAnalysis**](docs/FunctionsDataTypesAPI.md#listfunctiondatatypesforanalysis) | **Get** /v2/analyses/{analysis_id}/functions/data_types | List Function Data Types
*FunctionsDataTypesAPI* | [**ListFunctionDataTypesForFunctions**](docs/FunctionsDataTypesAPI.md#listfunctiondatatypesforfunctions) | **Get** /v2/functions/data_types | List Function Data Types
*FunctionsDataTypesAPI* | [**UpdateFunctionDataTypes**](docs/FunctionsDataTypesAPI.md#updatefunctiondatatypes) | **Put** /v2/analyses/{analysis_id}/functions/{function_id}/data_types | Update Function Data Types
*FunctionsRenamingHistoryAPI* | [**BatchRenameFunction**](docs/FunctionsRenamingHistoryAPI.md#batchrenamefunction) | **Post** /v2/functions/rename/batch | Batch Rename Functions
*FunctionsRenamingHistoryAPI* | [**BatchRenameFunctions**](docs/FunctionsRenamingHistoryAPI.md#batchrenamefunctions) | **Post** /v3/functions/rename | Batch rename functions
*FunctionsRenamingHistoryAPI* | [**GetFunctionHistory**](docs/FunctionsRenamingHistoryAPI.md#getfunctionhistory) | **Get** /v3/functions/{function_id}/history | Get function name history
*FunctionsRenamingHistoryAPI* | [**GetFunctionNameHistory**](docs/FunctionsRenamingHistoryAPI.md#getfunctionnamehistory) | **Get** /v2/functions/history/{function_id} | Get Function Name History
*FunctionsRenamingHistoryAPI* | [**RenameFunction**](docs/FunctionsRenamingHistoryAPI.md#renamefunction) | **Post** /v3/functions/{function_id}/rename | Rename a function
*FunctionsRenamingHistoryAPI* | [**RenameFunctionId**](docs/FunctionsRenamingHistoryAPI.md#renamefunctionid) | **Post** /v2/functions/rename/{function_id} | Rename Function
*FunctionsRenamingHistoryAPI* | [**RevertFunctionName**](docs/FunctionsRenamingHistoryAPI.md#revertfunctionname) | **Post** /v2/functions/history/{function_id}/{history_id} | Revert the function name
*FunctionsRenamingHistoryAPI* | [**RevertFunctionName_0**](docs/FunctionsRenamingHistoryAPI.md#revertfunctionname_0) | **Post** /v3/functions/{function_id}/history/{history_id}/revert | Revert function name
*ModelsAPI* | [**GetModels**](docs/ModelsAPI.md#getmodels) | **Get** /v2/models | Gets models
*ReportsAPI* | [**CreatePdfReport**](docs/ReportsAPI.md#createpdfreport) | **Post** /v3/analysis/{analysis_id}/pdf | Start PDF report generation
*ReportsAPI* | [**DownloadPdfReport**](docs/ReportsAPI.md#downloadpdfreport) | **Get** /v3/analysis/{analysis_id}/pdf/{task_id} | Download generated PDF report
*ReportsAPI* | [**GetPdfReportStatus**](docs/ReportsAPI.md#getpdfreportstatus) | **Get** /v3/analysis/{analysis_id}/pdf/{task_id}/status | Get PDF report workflow status
*SearchAPI* | [**SearchBinaries**](docs/SearchAPI.md#searchbinaries) | **Get** /v2/search/binaries | Binaries search
*SearchAPI* | [**SearchCollections**](docs/SearchAPI.md#searchcollections) | **Get** /v2/search/collections | Collections search
*SearchAPI* | [**SearchFunctions**](docs/SearchAPI.md#searchfunctions) | **Get** /v2/search/functions | Functions search
*SearchAPI* | [**SearchTags**](docs/SearchAPI.md#searchtags) | **Get** /v2/search/tags | Tags search


## Documentation For Models

 - [APIError](docs/APIError.md)
 - [AdditionalDetailsStatusResponse](docs/AdditionalDetailsStatusResponse.md)
 - [Addr](docs/Addr.md)
 - [AiDecompilationRating](docs/AiDecompilationRating.md)
 - [AiDecompilationTaskStatus](docs/AiDecompilationTaskStatus.md)
 - [AiUnstripRequest](docs/AiUnstripRequest.md)
 - [AnalysisAccessInfo](docs/AnalysisAccessInfo.md)
 - [AnalysisBulkAddTagsRequest](docs/AnalysisBulkAddTagsRequest.md)
 - [AnalysisBulkAddTagsResponse](docs/AnalysisBulkAddTagsResponse.md)
 - [AnalysisBulkAddTagsResponseItem](docs/AnalysisBulkAddTagsResponseItem.md)
 - [AnalysisConfig](docs/AnalysisConfig.md)
 - [AnalysisCreateRequest](docs/AnalysisCreateRequest.md)
 - [AnalysisCreateResponse](docs/AnalysisCreateResponse.md)
 - [AnalysisDetailResponse](docs/AnalysisDetailResponse.md)
 - [AnalysisFunctionMapping](docs/AnalysisFunctionMapping.md)
 - [AnalysisFunctionMatchingRequest](docs/AnalysisFunctionMatchingRequest.md)
 - [AnalysisFunctions](docs/AnalysisFunctions.md)
 - [AnalysisFunctionsList](docs/AnalysisFunctionsList.md)
 - [AnalysisRecord](docs/AnalysisRecord.md)
 - [AnalysisReport](docs/AnalysisReport.md)
 - [AnalysisScope](docs/AnalysisScope.md)
 - [AnalysisStringInput](docs/AnalysisStringInput.md)
 - [AnalysisStringsResponse](docs/AnalysisStringsResponse.md)
 - [AnalysisStringsStatusResponse](docs/AnalysisStringsStatusResponse.md)
 - [AnalysisTags](docs/AnalysisTags.md)
 - [AnalysisUpdateRequest](docs/AnalysisUpdateRequest.md)
 - [AnalysisUpdateTagsRequest](docs/AnalysisUpdateTagsRequest.md)
 - [AnalysisUpdateTagsResponse](docs/AnalysisUpdateTagsResponse.md)
 - [ApiCall](docs/ApiCall.md)
 - [AppApiRestV2AgentSchemaCapability](docs/AppApiRestV2AgentSchemaCapability.md)
 - [AppApiRestV2AnalysesEnumsOrderBy](docs/AppApiRestV2AnalysesEnumsOrderBy.md)
 - [AppApiRestV2CollectionsEnumsOrderBy](docs/AppApiRestV2CollectionsEnumsOrderBy.md)
 - [AppApiRestV2FunctionsResponsesFunction](docs/AppApiRestV2FunctionsResponsesFunction.md)
 - [AppApiRestV2FunctionsTypesFunction](docs/AppApiRestV2FunctionsTypesFunction.md)
 - [AppApiRestV2InfoTypesCapability](docs/AppApiRestV2InfoTypesCapability.md)
 - [Argument](docs/Argument.md)
 - [AutoRunAgents](docs/AutoRunAgents.md)
 - [AutoUnstripRequest](docs/AutoUnstripRequest.md)
 - [AutoUnstripResponse](docs/AutoUnstripResponse.md)
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
 - [BaseResponseFunctionDataTypes](docs/BaseResponseFunctionDataTypes.md)
 - [BaseResponseFunctionDataTypesList](docs/BaseResponseFunctionDataTypesList.md)
 - [BaseResponseFunctionSearchResponse](docs/BaseResponseFunctionSearchResponse.md)
 - [BaseResponseFunctionStringsResponse](docs/BaseResponseFunctionStringsResponse.md)
 - [BaseResponseFunctionTaskResponse](docs/BaseResponseFunctionTaskResponse.md)
 - [BaseResponseFunctionsDetailResponse](docs/BaseResponseFunctionsDetailResponse.md)
 - [BaseResponseGenerateFunctionDataTypes](docs/BaseResponseGenerateFunctionDataTypes.md)
 - [BaseResponseGenerationStatusList](docs/BaseResponseGenerationStatusList.md)
 - [BaseResponseGetAiDecompilationTask](docs/BaseResponseGetAiDecompilationTask.md)
 - [BaseResponseGetMeResponse](docs/BaseResponseGetMeResponse.md)
 - [BaseResponseGetPublicUserResponse](docs/BaseResponseGetPublicUserResponse.md)
 - [BaseResponseListCalleesCallerFunctionsResponse](docs/BaseResponseListCalleesCallerFunctionsResponse.md)
 - [BaseResponseListCollectionResults](docs/BaseResponseListCollectionResults.md)
 - [BaseResponseListCommentResponse](docs/BaseResponseListCommentResponse.md)
 - [BaseResponseListDieMatch](docs/BaseResponseListDieMatch.md)
 - [BaseResponseListFunctionNameHistory](docs/BaseResponseListFunctionNameHistory.md)
 - [BaseResponseListSBOM](docs/BaseResponseListSBOM.md)
 - [BaseResponseListUserActivityResponse](docs/BaseResponseListUserActivityResponse.md)
 - [BaseResponseLogs](docs/BaseResponseLogs.md)
 - [BaseResponseModelsResponse](docs/BaseResponseModelsResponse.md)
 - [BaseResponseParams](docs/BaseResponseParams.md)
 - [BaseResponseQueuedWorkflowTaskResponse](docs/BaseResponseQueuedWorkflowTaskResponse.md)
 - [BaseResponseRecent](docs/BaseResponseRecent.md)
 - [BaseResponseReportAnalysisResponse](docs/BaseResponseReportAnalysisResponse.md)
 - [BaseResponseStatus](docs/BaseResponseStatus.md)
 - [BaseResponseStr](docs/BaseResponseStr.md)
 - [BaseResponseTagSearchResponse](docs/BaseResponseTagSearchResponse.md)
 - [BaseResponseTaskResponse](docs/BaseResponseTaskResponse.md)
 - [BaseResponseTaskStatusResponse](docs/BaseResponseTaskStatusResponse.md)
 - [BaseResponseTriageReportResponse](docs/BaseResponseTriageReportResponse.md)
 - [BaseResponseUnionGetAiDecompilationRatingResponseNoneType](docs/BaseResponseUnionGetAiDecompilationRatingResponseNoneType.md)
 - [BaseResponseUploadResponse](docs/BaseResponseUploadResponse.md)
 - [BaseResponseVulnerabilities](docs/BaseResponseVulnerabilities.md)
 - [BaseResponseXrefResponse](docs/BaseResponseXrefResponse.md)
 - [Basic](docs/Basic.md)
 - [BatchRenameInputBody](docs/BatchRenameInputBody.md)
 - [BatchRenameItem](docs/BatchRenameItem.md)
 - [BatchRenameOutputBody](docs/BatchRenameOutputBody.md)
 - [BinariesRelatedStatusResponse](docs/BinariesRelatedStatusResponse.md)
 - [BinariesTaskStatus](docs/BinariesTaskStatus.md)
 - [BinaryAdditionalDetailsDataResponse](docs/BinaryAdditionalDetailsDataResponse.md)
 - [BinaryAdditionalResponse](docs/BinaryAdditionalResponse.md)
 - [BinaryConfig](docs/BinaryConfig.md)
 - [BinaryDetailsResponse](docs/BinaryDetailsResponse.md)
 - [BinaryExternalsResponse](docs/BinaryExternalsResponse.md)
 - [BinarySearchResponse](docs/BinarySearchResponse.md)
 - [BinarySearchResult](docs/BinarySearchResult.md)
 - [BinaryTaskStatus](docs/BinaryTaskStatus.md)
 - [BulkDeleteAnalysesRequest](docs/BulkDeleteAnalysesRequest.md)
 - [CalleeFunctionInfo](docs/CalleeFunctionInfo.md)
 - [CalleesCallerFunctionsResponse](docs/CalleesCallerFunctionsResponse.md)
 - [CallerFunctionInfo](docs/CallerFunctionInfo.md)
 - [Capabilities](docs/Capabilities.md)
 - [CapabilitiesAgentResponse](docs/CapabilitiesAgentResponse.md)
 - [ChildBinariesResponse](docs/ChildBinariesResponse.md)
 - [CodeSignatureModel](docs/CodeSignatureModel.md)
 - [CollectionBinariesUpdateRequest](docs/CollectionBinariesUpdateRequest.md)
 - [CollectionBinariesUpdateResponse](docs/CollectionBinariesUpdateResponse.md)
 - [CollectionBinaryResponse](docs/CollectionBinaryResponse.md)
 - [CollectionCreateRequest](docs/CollectionCreateRequest.md)
 - [CollectionListItem](docs/CollectionListItem.md)
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
 - [ConfigResponse](docs/ConfigResponse.md)
 - [ConfirmToolInputBody](docs/ConfirmToolInputBody.md)
 - [Connection](docs/Connection.md)
 - [Context](docs/Context.md)
 - [Conversation](docs/Conversation.md)
 - [ConversationContext](docs/ConversationContext.md)
 - [ConversationWithEvents](docs/ConversationWithEvents.md)
 - [CreateAIDecompOutputBody](docs/CreateAIDecompOutputBody.md)
 - [CreateConversationRequest](docs/CreateConversationRequest.md)
 - [Created](docs/Created.md)
 - [DecompilationCommentContext](docs/DecompilationCommentContext.md)
 - [DecompilationData](docs/DecompilationData.md)
 - [DieMatch](docs/DieMatch.md)
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
 - [EntrypointModel](docs/EntrypointModel.md)
 - [Enumeration](docs/Enumeration.md)
 - [ErrorBody](docs/ErrorBody.md)
 - [ErrorModel](docs/ErrorModel.md)
 - [Event](docs/Event.md)
 - [EventCONTEXTCOMPACTED](docs/EventCONTEXTCOMPACTED.md)
 - [EventRUNCANCELLED](docs/EventRUNCANCELLED.md)
 - [EventRUNERROR](docs/EventRUNERROR.md)
 - [EventRUNFINISHED](docs/EventRUNFINISHED.md)
 - [EventRUNSTARTED](docs/EventRUNSTARTED.md)
 - [EventSTEPFINISHED](docs/EventSTEPFINISHED.md)
 - [EventSTEPSTARTED](docs/EventSTEPSTARTED.md)
 - [EventTEXTMESSAGECONTENT](docs/EventTEXTMESSAGECONTENT.md)
 - [EventTEXTMESSAGEEND](docs/EventTEXTMESSAGEEND.md)
 - [EventTEXTMESSAGESTART](docs/EventTEXTMESSAGESTART.md)
 - [EventTITLEUPDATED](docs/EventTITLEUPDATED.md)
 - [EventTOOLCALLARGSDELTA](docs/EventTOOLCALLARGSDELTA.md)
 - [EventTOOLCALLEND](docs/EventTOOLCALLEND.md)
 - [EventTOOLCALLRESULT](docs/EventTOOLCALLRESULT.md)
 - [EventTOOLCALLSTART](docs/EventTOOLCALLSTART.md)
 - [EventTOOLCONFIRMATIONREQUIRED](docs/EventTOOLCONFIRMATIONREQUIRED.md)
 - [ExportModel](docs/ExportModel.md)
 - [ExternalResponse](docs/ExternalResponse.md)
 - [ExtractedURL](docs/ExtractedURL.md)
 - [FileActivityEntry](docs/FileActivityEntry.md)
 - [FileFormat](docs/FileFormat.md)
 - [FileHashes](docs/FileHashes.md)
 - [FileMetadata](docs/FileMetadata.md)
 - [Filters](docs/Filters.md)
 - [FuncDepsInner](docs/FuncDepsInner.md)
 - [FunctionBlockDestinationResponse](docs/FunctionBlockDestinationResponse.md)
 - [FunctionBlockResponse](docs/FunctionBlockResponse.md)
 - [FunctionBlocksResponse](docs/FunctionBlocksResponse.md)
 - [FunctionBoundary](docs/FunctionBoundary.md)
 - [FunctionCapabilityResponse](docs/FunctionCapabilityResponse.md)
 - [FunctionCommentCreateRequest](docs/FunctionCommentCreateRequest.md)
 - [FunctionDataTypes](docs/FunctionDataTypes.md)
 - [FunctionDataTypesList](docs/FunctionDataTypesList.md)
 - [FunctionDataTypesListItem](docs/FunctionDataTypesListItem.md)
 - [FunctionDataTypesParams](docs/FunctionDataTypesParams.md)
 - [FunctionDataTypesStatus](docs/FunctionDataTypesStatus.md)
 - [FunctionHeader](docs/FunctionHeader.md)
 - [FunctionInfoInput](docs/FunctionInfoInput.md)
 - [FunctionInfoOutput](docs/FunctionInfoOutput.md)
 - [FunctionListItem](docs/FunctionListItem.md)
 - [FunctionLocalVariableResponse](docs/FunctionLocalVariableResponse.md)
 - [FunctionMapping](docs/FunctionMapping.md)
 - [FunctionMappingFull](docs/FunctionMappingFull.md)
 - [FunctionMatch](docs/FunctionMatch.md)
 - [FunctionMatchingFilters](docs/FunctionMatchingFilters.md)
 - [FunctionMatchingRequest](docs/FunctionMatchingRequest.md)
 - [FunctionMatchingResponse](docs/FunctionMatchingResponse.md)
 - [FunctionNameHistory](docs/FunctionNameHistory.md)
 - [FunctionParamResponse](docs/FunctionParamResponse.md)
 - [FunctionRename](docs/FunctionRename.md)
 - [FunctionRenameMap](docs/FunctionRenameMap.md)
 - [FunctionSearchResponse](docs/FunctionSearchResponse.md)
 - [FunctionSearchResult](docs/FunctionSearchResult.md)
 - [FunctionSourceType](docs/FunctionSourceType.md)
 - [FunctionString](docs/FunctionString.md)
 - [FunctionStringsResponse](docs/FunctionStringsResponse.md)
 - [FunctionTaskResponse](docs/FunctionTaskResponse.md)
 - [FunctionTaskStatus](docs/FunctionTaskStatus.md)
 - [FunctionTypeInput](docs/FunctionTypeInput.md)
 - [FunctionTypeOutput](docs/FunctionTypeOutput.md)
 - [FunctionsDetailResponse](docs/FunctionsDetailResponse.md)
 - [FunctionsListRename](docs/FunctionsListRename.md)
 - [GenerateFunctionDataTypes](docs/GenerateFunctionDataTypes.md)
 - [GeneratePDFOutputBody](docs/GeneratePDFOutputBody.md)
 - [GenerationStatusList](docs/GenerationStatusList.md)
 - [GetAiDecompilationRatingResponse](docs/GetAiDecompilationRatingResponse.md)
 - [GetAiDecompilationTask](docs/GetAiDecompilationTask.md)
 - [GetMeResponse](docs/GetMeResponse.md)
 - [GetPublicUserResponse](docs/GetPublicUserResponse.md)
 - [GlobalVariable](docs/GlobalVariable.md)
 - [HistoryEntry](docs/HistoryEntry.md)
 - [HttpRequest](docs/HttpRequest.md)
 - [IOC](docs/IOC.md)
 - [ISA](docs/ISA.md)
 - [IconModel](docs/IconModel.md)
 - [ImportModel](docs/ImportModel.md)
 - [InlineComment](docs/InlineComment.md)
 - [InsertAnalysisLogRequest](docs/InsertAnalysisLogRequest.md)
 - [InverseFunctionMapItem](docs/InverseFunctionMapItem.md)
 - [InverseStringMapItem](docs/InverseStringMapItem.md)
 - [InverseValue](docs/InverseValue.md)
 - [ListCollectionResults](docs/ListCollectionResults.md)
 - [Logs](docs/Logs.md)
 - [MITRETechnique](docs/MITRETechnique.md)
 - [MatchedFunction](docs/MatchedFunction.md)
 - [MatchedFunctionSuggestion](docs/MatchedFunctionSuggestion.md)
 - [MemdumpEntry](docs/MemdumpEntry.md)
 - [MetaModel](docs/MetaModel.md)
 - [ModelName](docs/ModelName.md)
 - [ModelsResponse](docs/ModelsResponse.md)
 - [ModuleLoadEntry](docs/ModuleLoadEntry.md)
 - [MutexEntry](docs/MutexEntry.md)
 - [NameConfidence](docs/NameConfidence.md)
 - [NameSourceType](docs/NameSourceType.md)
 - [NetworkActivity](docs/NetworkActivity.md)
 - [NumericAddr](docs/NumericAddr.md)
 - [Order](docs/Order.md)
 - [PDBDebugModel](docs/PDBDebugModel.md)
 - [PEModel](docs/PEModel.md)
 - [PaginationModel](docs/PaginationModel.md)
 - [Params](docs/Params.md)
 - [Platform](docs/Platform.md)
 - [ProcessActivityEntry](docs/ProcessActivityEntry.md)
 - [ProcessMemdumps](docs/ProcessMemdumps.md)
 - [ProcessNode](docs/ProcessNode.md)
 - [ProcessTree](docs/ProcessTree.md)
 - [ProgressMessage](docs/ProgressMessage.md)
 - [PutAnalysisStringsRequest](docs/PutAnalysisStringsRequest.md)
 - [QueuePositionResponse](docs/QueuePositionResponse.md)
 - [QueuedWorkflowTaskResponse](docs/QueuedWorkflowTaskResponse.md)
 - [ReAnalysisForm](docs/ReAnalysisForm.md)
 - [Recent](docs/Recent.md)
 - [RegenerateOutputBody](docs/RegenerateOutputBody.md)
 - [RegenerateTarget](docs/RegenerateTarget.md)
 - [RegistryOperation](docs/RegistryOperation.md)
 - [RelativeBinaryResponse](docs/RelativeBinaryResponse.md)
 - [RenameInputBody](docs/RenameInputBody.md)
 - [RenameOutputBody](docs/RenameOutputBody.md)
 - [ReplacementValue](docs/ReplacementValue.md)
 - [ReportAnalysisResponse](docs/ReportAnalysisResponse.md)
 - [ReportEvent](docs/ReportEvent.md)
 - [ReportInfo](docs/ReportInfo.md)
 - [ReportOptions](docs/ReportOptions.md)
 - [RevertOutputBody](docs/RevertOutputBody.md)
 - [SBOM](docs/SBOM.md)
 - [SBOMPackage](docs/SBOMPackage.md)
 - [SandboxOptions](docs/SandboxOptions.md)
 - [SandboxStartMethod](docs/SandboxStartMethod.md)
 - [SandboxTimeout](docs/SandboxTimeout.md)
 - [ScheduledTaskEntry](docs/ScheduledTaskEntry.md)
 - [ScrapeThirdPartyConfig](docs/ScrapeThirdPartyConfig.md)
 - [SectionModel](docs/SectionModel.md)
 - [SecurityModel](docs/SecurityModel.md)
 - [SegmentInfo](docs/SegmentInfo.md)
 - [SendMessageRequest](docs/SendMessageRequest.md)
 - [ServerSentEventsInner](docs/ServerSentEventsInner.md)
 - [ServiceEntry](docs/ServiceEntry.md)
 - [SingleCodeCertificateModel](docs/SingleCodeCertificateModel.md)
 - [SingleCodeSignatureModel](docs/SingleCodeSignatureModel.md)
 - [SinglePDBEntryModel](docs/SinglePDBEntryModel.md)
 - [SingleSectionModel](docs/SingleSectionModel.md)
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
 - [SseEventToolCallResultData](docs/SseEventToolCallResultData.md)
 - [SseEventToolCallStartData](docs/SseEventToolCallStartData.md)
 - [SseEventToolConfirmationRequiredData](docs/SseEventToolConfirmationRequiredData.md)
 - [StackVariable](docs/StackVariable.md)
 - [StartupInfo](docs/StartupInfo.md)
 - [StatusInput](docs/StatusInput.md)
 - [StatusOutput](docs/StatusOutput.md)
 - [StatusResponse](docs/StatusResponse.md)
 - [StringFunctions](docs/StringFunctions.md)
 - [StringSource](docs/StringSource.md)
 - [Structure](docs/Structure.md)
 - [StructureMember](docs/StructureMember.md)
 - [SubmitUserFeedbackRequest](docs/SubmitUserFeedbackRequest.md)
 - [SummaryData](docs/SummaryData.md)
 - [Symbols](docs/Symbols.md)
 - [Tag](docs/Tag.md)
 - [TagItem](docs/TagItem.md)
 - [TagResponse](docs/TagResponse.md)
 - [TagSearchResponse](docs/TagSearchResponse.md)
 - [TagSearchResult](docs/TagSearchResult.md)
 - [TaskResponse](docs/TaskResponse.md)
 - [TaskStatus](docs/TaskStatus.md)
 - [TaskStatusResponse](docs/TaskStatusResponse.md)
 - [TimestampModel](docs/TimestampModel.md)
 - [TokenisedData](docs/TokenisedData.md)
 - [TriageFunctionResponse](docs/TriageFunctionResponse.md)
 - [TriageReportResponse](docs/TriageReportResponse.md)
 - [Ttp](docs/Ttp.md)
 - [TypeDefinition](docs/TypeDefinition.md)
 - [UpdateFunctionDataTypes](docs/UpdateFunctionDataTypes.md)
 - [UploadFileType](docs/UploadFileType.md)
 - [UploadResponse](docs/UploadResponse.md)
 - [UpsertAiDecomplationRatingRequest](docs/UpsertAiDecomplationRatingRequest.md)
 - [UpsertOverridesData](docs/UpsertOverridesData.md)
 - [UpsertOverridesInputBody](docs/UpsertOverridesInputBody.md)
 - [UserActivityResponse](docs/UserActivityResponse.md)
 - [Vulnerabilities](docs/Vulnerabilities.md)
 - [Vulnerability](docs/Vulnerability.md)
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
