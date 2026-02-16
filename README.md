# RevEng.AI Go SDK

This is the Go SDK for the RevEng.AI API.

To use the SDK you will first need to obtain an API key from [https://reveng.ai](https://reveng.ai/register).

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import sdk "github.com/RevEngAI/sdk-go"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sdk.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), sdk.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sdk.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), sdk.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sdk.ContextOperationServerIndices` and `sdk.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), sdk.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sdk.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.reveng.ai*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
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
*AnalysesCoreAPI* | [**GetAnalysisStatus**](docs/AnalysesCoreAPI.md#getanalysisstatus) | **Get** /v2/analyses/{analysis_id}/status | Gets the status of an analysis
*AnalysesCoreAPI* | [**InsertAnalysisLog**](docs/AnalysesCoreAPI.md#insertanalysislog) | **Post** /v2/analyses/{analysis_id}/logs | Insert a log entry for an analysis
*AnalysesCoreAPI* | [**ListAnalyses**](docs/AnalysesCoreAPI.md#listanalyses) | **Get** /v2/analyses/list | Gets the most recent analyses
*AnalysesCoreAPI* | [**LookupBinaryId**](docs/AnalysesCoreAPI.md#lookupbinaryid) | **Get** /v2/analyses/lookup/{binary_id} | Gets the analysis ID from binary ID
*AnalysesCoreAPI* | [**RequeueAnalysis**](docs/AnalysesCoreAPI.md#requeueanalysis) | **Post** /v2/analyses/{analysis_id}/requeue | Requeue Analysis
*AnalysesCoreAPI* | [**UpdateAnalysis**](docs/AnalysesCoreAPI.md#updateanalysis) | **Patch** /v2/analyses/{analysis_id} | Update Analysis
*AnalysesCoreAPI* | [**UpdateAnalysisTags**](docs/AnalysesCoreAPI.md#updateanalysistags) | **Patch** /v2/analyses/{analysis_id}/tags | Update Analysis Tags
*AnalysesCoreAPI* | [**UploadFile**](docs/AnalysesCoreAPI.md#uploadfile) | **Post** /v2/upload | Upload File
*AnalysesDynamicExecutionAPI* | [**GetDynamicExecutionStatus**](docs/AnalysesDynamicExecutionAPI.md#getdynamicexecutionstatus) | **Get** /v2/analyses/{analysis_id}/dynamic-execution/status | Get the status of a dynamic execution task
*AnalysesDynamicExecutionAPI* | [**GetNetworkOverview**](docs/AnalysesDynamicExecutionAPI.md#getnetworkoverview) | **Get** /v2/analyses/{analysis_id}/dynamic-execution/network-overview | Get the dynamic execution results for network overview
*AnalysesDynamicExecutionAPI* | [**GetProcessDump**](docs/AnalysesDynamicExecutionAPI.md#getprocessdump) | **Get** /v2/analyses/{analysis_id}/dynamic-execution/process-dumps/{dump_name} | Get the dynamic execution results for a specific process dump
*AnalysesDynamicExecutionAPI* | [**GetProcessDumps**](docs/AnalysesDynamicExecutionAPI.md#getprocessdumps) | **Get** /v2/analyses/{analysis_id}/dynamic-execution/process-dumps | Get the dynamic execution results for process dumps
*AnalysesDynamicExecutionAPI* | [**GetProcessRegistry**](docs/AnalysesDynamicExecutionAPI.md#getprocessregistry) | **Get** /v2/analyses/{analysis_id}/dynamic-execution/process-registry | Get the dynamic execution results for process registry
*AnalysesDynamicExecutionAPI* | [**GetProcessTree**](docs/AnalysesDynamicExecutionAPI.md#getprocesstree) | **Get** /v2/analyses/{analysis_id}/dynamic-execution/process-tree | Get the dynamic execution results for process tree
*AnalysesDynamicExecutionAPI* | [**GetTtps**](docs/AnalysesDynamicExecutionAPI.md#getttps) | **Get** /v2/analyses/{analysis_id}/dynamic-execution/ttps | Get the dynamic execution results for ttps
*AnalysesResultsMetadataAPI* | [**GetAnalysisFunctionsPaginated**](docs/AnalysesResultsMetadataAPI.md#getanalysisfunctionspaginated) | **Get** /v2/analyses/{analysis_id}/functions | Get functions from analysis
*AnalysesResultsMetadataAPI* | [**GetCapabilities**](docs/AnalysesResultsMetadataAPI.md#getcapabilities) | **Get** /v2/analyses/{analysis_id}/capabilities | Gets the capabilities from the analysis
*AnalysesResultsMetadataAPI* | [**GetCommunities**](docs/AnalysesResultsMetadataAPI.md#getcommunities) | **Get** /v2/analyses/{analysis_id}/communities | Gets the communities found in the analysis
*AnalysesResultsMetadataAPI* | [**GetFunctionsList**](docs/AnalysesResultsMetadataAPI.md#getfunctionslist) | **Get** /v2/analyses/{analysis_id}/functions/list | Gets functions from analysis
*AnalysesResultsMetadataAPI* | [**GetPdf**](docs/AnalysesResultsMetadataAPI.md#getpdf) | **Get** /v2/analyses/{analysis_id}/pdf | Gets the PDF found in the analysis
*AnalysesResultsMetadataAPI* | [**GetSbom**](docs/AnalysesResultsMetadataAPI.md#getsbom) | **Get** /v2/analyses/{analysis_id}/sbom | Gets the software-bill-of-materials (SBOM) found in the analysis
*AnalysesResultsMetadataAPI* | [**GetTags**](docs/AnalysesResultsMetadataAPI.md#gettags) | **Get** /v2/analyses/{analysis_id}/tags | Get function tags with maliciousness score
*AnalysesResultsMetadataAPI* | [**GetVulnerabilities**](docs/AnalysesResultsMetadataAPI.md#getvulnerabilities) | **Get** /v2/analyses/{analysis_id}/vulnerabilities | Gets the vulnerabilities found in the analysis
*AnalysesSecurityChecksAPI* | [**CreateScurityChecksTask**](docs/AnalysesSecurityChecksAPI.md#createscuritycheckstask) | **Post** /v2/analyses/{analysis_id}/security-checks | Queues a security check process
*AnalysesSecurityChecksAPI* | [**GetSecurityChecks**](docs/AnalysesSecurityChecksAPI.md#getsecuritychecks) | **Get** /v2/analyses/{analysis_id}/security-checks | Get Security Checks
*AnalysesSecurityChecksAPI* | [**GetSecurityChecksTaskStatus**](docs/AnalysesSecurityChecksAPI.md#getsecuritycheckstaskstatus) | **Get** /v2/analyses/{analysis_id}/security-checks/status | Check the status of a security check process
*AuthenticationUsersAPI* | [**GetRequesterUserInfo**](docs/AuthenticationUsersAPI.md#getrequesteruserinfo) | **Get** /v2/users/me | Get the requesters user information
*AuthenticationUsersAPI* | [**GetUser**](docs/AuthenticationUsersAPI.md#getuser) | **Get** /v2/users/{user_id} | Get a user&#39;s public information
*AuthenticationUsersAPI* | [**GetUserActivity**](docs/AuthenticationUsersAPI.md#getuseractivity) | **Get** /v2/users/activity | Get auth user activity
*AuthenticationUsersAPI* | [**GetUserComments**](docs/AuthenticationUsersAPI.md#getusercomments) | **Get** /v2/users/me/comments | Get comments by user
*AuthenticationUsersAPI* | [**LoginUser**](docs/AuthenticationUsersAPI.md#loginuser) | **Post** /v2/auth/login | Authenticate a user
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
*ExternalSourcesAPI* | [**CreateExternalTaskVt**](docs/ExternalSourcesAPI.md#createexternaltaskvt) | **Post** /v2/analysis/{analysis_id}/external/vt | Pulls data from VirusTotal
*ExternalSourcesAPI* | [**GetVtData**](docs/ExternalSourcesAPI.md#getvtdata) | **Get** /v2/analysis/{analysis_id}/external/vt | Get VirusTotal data
*ExternalSourcesAPI* | [**GetVtTaskStatus**](docs/ExternalSourcesAPI.md#getvttaskstatus) | **Get** /v2/analysis/{analysis_id}/external/vt/status | Check the status of VirusTotal data retrieval
*FirmwareAPI* | [**GetBinariesForFirmwareTask**](docs/FirmwareAPI.md#getbinariesforfirmwaretask) | **Get** /v2/firmware/get-binaries/{task_id} | Upload firmware for unpacking
*FirmwareAPI* | [**UploadFirmware**](docs/FirmwareAPI.md#uploadfirmware) | **Post** /v2/firmware | Upload firmware for unpacking
*FunctionsAIDecompilationAPI* | [**CreateAiDecompilationComment**](docs/FunctionsAIDecompilationAPI.md#createaidecompilationcomment) | **Post** /v2/functions/{function_id}/ai-decompilation/comments | Create a comment for this function
*FunctionsAIDecompilationAPI* | [**CreateAiDecompilationTask**](docs/FunctionsAIDecompilationAPI.md#createaidecompilationtask) | **Post** /v2/functions/{function_id}/ai-decompilation | Begins AI Decompilation Process
*FunctionsAIDecompilationAPI* | [**DeleteAiDecompilationComment**](docs/FunctionsAIDecompilationAPI.md#deleteaidecompilationcomment) | **Delete** /v2/functions/{function_id}/ai-decompilation/comments/{comment_id} | Delete a comment
*FunctionsAIDecompilationAPI* | [**GetAiDecompilationComments**](docs/FunctionsAIDecompilationAPI.md#getaidecompilationcomments) | **Get** /v2/functions/{function_id}/ai-decompilation/comments | Get comments for this function
*FunctionsAIDecompilationAPI* | [**GetAiDecompilationRating**](docs/FunctionsAIDecompilationAPI.md#getaidecompilationrating) | **Get** /v2/functions/{function_id}/ai-decompilation/rating | Get rating for AI decompilation
*FunctionsAIDecompilationAPI* | [**GetAiDecompilationTaskResult**](docs/FunctionsAIDecompilationAPI.md#getaidecompilationtaskresult) | **Get** /v2/functions/{function_id}/ai-decompilation | Polls AI Decompilation Process
*FunctionsAIDecompilationAPI* | [**GetAiDecompilationTaskStatus**](docs/FunctionsAIDecompilationAPI.md#getaidecompilationtaskstatus) | **Get** /v2/functions/{function_id}/ai-decompilation/status | Check the status of a function ai decompilation
*FunctionsAIDecompilationAPI* | [**UpdateAiDecompilationComment**](docs/FunctionsAIDecompilationAPI.md#updateaidecompilationcomment) | **Patch** /v2/functions/{function_id}/ai-decompilation/comments/{comment_id} | Update a comment
*FunctionsAIDecompilationAPI* | [**UpsertAiDecompilationRating**](docs/FunctionsAIDecompilationAPI.md#upsertaidecompilationrating) | **Patch** /v2/functions/{function_id}/ai-decompilation/rating | Upsert rating for AI decompilation
*FunctionsBlockCommentsAPI* | [**GenerateBlockCommentsForBlockInFunction**](docs/FunctionsBlockCommentsAPI.md#generateblockcommentsforblockinfunction) | **Post** /v2/functions/{function_id}/block-comments/single | Generate block comments for a specific block in a function
*FunctionsBlockCommentsAPI* | [**GenerateBlockCommentsForFunction**](docs/FunctionsBlockCommentsAPI.md#generateblockcommentsforfunction) | **Post** /v2/functions/{function_id}/block-comments | Generate block comments for a function
*FunctionsBlockCommentsAPI* | [**GenerateOverviewCommentForFunction**](docs/FunctionsBlockCommentsAPI.md#generateoverviewcommentforfunction) | **Post** /v2/functions/{function_id}/block-comments/overview | Generate overview comment for a function
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
*FunctionsCoreAPI* | [**GetFunctionCapabilities**](docs/FunctionsCoreAPI.md#getfunctioncapabilities) | **Get** /v2/functions/{function_id}/capabilities | Retrieve a functions capabilities
*FunctionsCoreAPI* | [**GetFunctionDetails**](docs/FunctionsCoreAPI.md#getfunctiondetails) | **Get** /v2/functions/{function_id} | Get function details
*FunctionsCoreAPI* | [**GetFunctionStrings**](docs/FunctionsCoreAPI.md#getfunctionstrings) | **Get** /v2/functions/{function_id}/strings | Get string information found in the function
*FunctionsDataTypesAPI* | [**GenerateFunctionDataTypesForAnalysis**](docs/FunctionsDataTypesAPI.md#generatefunctiondatatypesforanalysis) | **Post** /v2/analyses/{analysis_id}/functions/data_types | Generate Function Data Types
*FunctionsDataTypesAPI* | [**GenerateFunctionDataTypesForFunctions**](docs/FunctionsDataTypesAPI.md#generatefunctiondatatypesforfunctions) | **Post** /v2/functions/data_types | Generate Function Data Types for an arbitrary list of functions
*FunctionsDataTypesAPI* | [**GetFunctionDataTypes**](docs/FunctionsDataTypesAPI.md#getfunctiondatatypes) | **Get** /v2/analyses/{analysis_id}/functions/{function_id}/data_types | Get Function Data Types
*FunctionsDataTypesAPI* | [**ListFunctionDataTypesForAnalysis**](docs/FunctionsDataTypesAPI.md#listfunctiondatatypesforanalysis) | **Get** /v2/analyses/{analysis_id}/functions/data_types | List Function Data Types
*FunctionsDataTypesAPI* | [**ListFunctionDataTypesForFunctions**](docs/FunctionsDataTypesAPI.md#listfunctiondatatypesforfunctions) | **Get** /v2/functions/data_types | List Function Data Types
*FunctionsDataTypesAPI* | [**UpdateFunctionDataTypes**](docs/FunctionsDataTypesAPI.md#updatefunctiondatatypes) | **Put** /v2/analyses/{analysis_id}/functions/{function_id}/data_types | Update Function Data Types
*FunctionsDecompilationAPI* | [**CreateDecompilationComment**](docs/FunctionsDecompilationAPI.md#createdecompilationcomment) | **Post** /v2/functions/{function_id}/decompilation/comments | Create a comment for this function
*FunctionsDecompilationAPI* | [**DeleteDecompilationComment**](docs/FunctionsDecompilationAPI.md#deletedecompilationcomment) | **Delete** /v2/functions/{function_id}/decompilation/comments/{comment_id} | Delete a comment
*FunctionsDecompilationAPI* | [**GetDecompilationComments**](docs/FunctionsDecompilationAPI.md#getdecompilationcomments) | **Get** /v2/functions/{function_id}/decompilation/comments | Get comments for this function
*FunctionsDecompilationAPI* | [**UpdateDecompilationComment**](docs/FunctionsDecompilationAPI.md#updatedecompilationcomment) | **Patch** /v2/functions/{function_id}/decompilation/comments/{comment_id} | Update a comment
*FunctionsRenamingHistoryAPI* | [**BatchRenameFunction**](docs/FunctionsRenamingHistoryAPI.md#batchrenamefunction) | **Post** /v2/functions/rename/batch | Batch Rename Functions
*FunctionsRenamingHistoryAPI* | [**GetFunctionNameHistory**](docs/FunctionsRenamingHistoryAPI.md#getfunctionnamehistory) | **Get** /v2/functions/history/{function_id} | Get Function Name History
*FunctionsRenamingHistoryAPI* | [**RenameFunctionId**](docs/FunctionsRenamingHistoryAPI.md#renamefunctionid) | **Post** /v2/functions/rename/{function_id} | Rename Function
*FunctionsRenamingHistoryAPI* | [**RevertFunctionName**](docs/FunctionsRenamingHistoryAPI.md#revertfunctionname) | **Post** /v2/functions/history/{function_id}/{history_id} | Revert the function name
*ModelsAPI* | [**GetModels**](docs/ModelsAPI.md#getmodels) | **Get** /v2/models | Gets models
*SearchAPI* | [**SearchBinaries**](docs/SearchAPI.md#searchbinaries) | **Get** /v2/search/binaries | Binaries search
*SearchAPI* | [**SearchCollections**](docs/SearchAPI.md#searchcollections) | **Get** /v2/search/collections | Collections search
*SearchAPI* | [**SearchFunctions**](docs/SearchAPI.md#searchfunctions) | **Get** /v2/search/functions | Functions search
*SearchAPI* | [**SearchTags**](docs/SearchAPI.md#searchtags) | **Get** /v2/search/tags | Tags search


## Documentation For Models

 - [AdditionalDetailsStatusResponse](docs/AdditionalDetailsStatusResponse.md)
 - [Addr](docs/Addr.md)
 - [AiDecompilationRating](docs/AiDecompilationRating.md)
 - [AiUnstripRequest](docs/AiUnstripRequest.md)
 - [AnalysisAccessInfo](docs/AnalysisAccessInfo.md)
 - [AnalysisConfig](docs/AnalysisConfig.md)
 - [AnalysisCreateRequest](docs/AnalysisCreateRequest.md)
 - [AnalysisCreateResponse](docs/AnalysisCreateResponse.md)
 - [AnalysisDetailResponse](docs/AnalysisDetailResponse.md)
 - [AnalysisFunctionMapping](docs/AnalysisFunctionMapping.md)
 - [AnalysisFunctionMatchingRequest](docs/AnalysisFunctionMatchingRequest.md)
 - [AnalysisFunctions](docs/AnalysisFunctions.md)
 - [AnalysisFunctionsList](docs/AnalysisFunctionsList.md)
 - [AnalysisRecord](docs/AnalysisRecord.md)
 - [AnalysisScope](docs/AnalysisScope.md)
 - [AnalysisStringsResponse](docs/AnalysisStringsResponse.md)
 - [AnalysisStringsStatusResponse](docs/AnalysisStringsStatusResponse.md)
 - [AnalysisTags](docs/AnalysisTags.md)
 - [AnalysisUpdateRequest](docs/AnalysisUpdateRequest.md)
 - [AnalysisUpdateTagsRequest](docs/AnalysisUpdateTagsRequest.md)
 - [AnalysisUpdateTagsResponse](docs/AnalysisUpdateTagsResponse.md)
 - [AppApiRestV2AnalysesEnumsDynamicExecutionStatus](docs/AppApiRestV2AnalysesEnumsDynamicExecutionStatus.md)
 - [AppApiRestV2AnalysesEnumsOrderBy](docs/AppApiRestV2AnalysesEnumsOrderBy.md)
 - [AppApiRestV2CollectionsEnumsOrderBy](docs/AppApiRestV2CollectionsEnumsOrderBy.md)
 - [AppApiRestV2FunctionsResponsesFunction](docs/AppApiRestV2FunctionsResponsesFunction.md)
 - [AppApiRestV2FunctionsTypesFunction](docs/AppApiRestV2FunctionsTypesFunction.md)
 - [AppServicesDynamicExecutionSchemasDynamicExecutionStatus](docs/AppServicesDynamicExecutionSchemasDynamicExecutionStatus.md)
 - [Argument](docs/Argument.md)
 - [AutoUnstripRequest](docs/AutoUnstripRequest.md)
 - [AutoUnstripResponse](docs/AutoUnstripResponse.md)
 - [BaseResponse](docs/BaseResponse.md)
 - [BaseResponseAdditionalDetailsStatusResponse](docs/BaseResponseAdditionalDetailsStatusResponse.md)
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
 - [BaseResponseBlockCommentsGenerationForFunctionResponse](docs/BaseResponseBlockCommentsGenerationForFunctionResponse.md)
 - [BaseResponseBlockCommentsOverviewGenerationResponse](docs/BaseResponseBlockCommentsOverviewGenerationResponse.md)
 - [BaseResponseBool](docs/BaseResponseBool.md)
 - [BaseResponseCalleesCallerFunctionsResponse](docs/BaseResponseCalleesCallerFunctionsResponse.md)
 - [BaseResponseCapabilities](docs/BaseResponseCapabilities.md)
 - [BaseResponseCheckSecurityChecksTaskResponse](docs/BaseResponseCheckSecurityChecksTaskResponse.md)
 - [BaseResponseChildBinariesResponse](docs/BaseResponseChildBinariesResponse.md)
 - [BaseResponseCollectionBinariesUpdateResponse](docs/BaseResponseCollectionBinariesUpdateResponse.md)
 - [BaseResponseCollectionResponse](docs/BaseResponseCollectionResponse.md)
 - [BaseResponseCollectionSearchResponse](docs/BaseResponseCollectionSearchResponse.md)
 - [BaseResponseCollectionTagsUpdateResponse](docs/BaseResponseCollectionTagsUpdateResponse.md)
 - [BaseResponseCommentResponse](docs/BaseResponseCommentResponse.md)
 - [BaseResponseCommunities](docs/BaseResponseCommunities.md)
 - [BaseResponseConfigResponse](docs/BaseResponseConfigResponse.md)
 - [BaseResponseCreated](docs/BaseResponseCreated.md)
 - [BaseResponseDict](docs/BaseResponseDict.md)
 - [BaseResponseDynamicExecutionStatus](docs/BaseResponseDynamicExecutionStatus.md)
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
 - [BaseResponseGetAiDecompilationRatingResponse](docs/BaseResponseGetAiDecompilationRatingResponse.md)
 - [BaseResponseGetAiDecompilationTask](docs/BaseResponseGetAiDecompilationTask.md)
 - [BaseResponseGetMeResponse](docs/BaseResponseGetMeResponse.md)
 - [BaseResponseGetPublicUserResponse](docs/BaseResponseGetPublicUserResponse.md)
 - [BaseResponseListCollectionResults](docs/BaseResponseListCollectionResults.md)
 - [BaseResponseListCommentResponse](docs/BaseResponseListCommentResponse.md)
 - [BaseResponseListDieMatch](docs/BaseResponseListDieMatch.md)
 - [BaseResponseListFunctionNameHistory](docs/BaseResponseListFunctionNameHistory.md)
 - [BaseResponseListSBOM](docs/BaseResponseListSBOM.md)
 - [BaseResponseListUserActivityResponse](docs/BaseResponseListUserActivityResponse.md)
 - [BaseResponseLoginResponse](docs/BaseResponseLoginResponse.md)
 - [BaseResponseLogs](docs/BaseResponseLogs.md)
 - [BaseResponseModelsResponse](docs/BaseResponseModelsResponse.md)
 - [BaseResponseNetworkOverviewResponse](docs/BaseResponseNetworkOverviewResponse.md)
 - [BaseResponseParams](docs/BaseResponseParams.md)
 - [BaseResponseProcessDumps](docs/BaseResponseProcessDumps.md)
 - [BaseResponseProcessRegistry](docs/BaseResponseProcessRegistry.md)
 - [BaseResponseProcessTree](docs/BaseResponseProcessTree.md)
 - [BaseResponseQueuedSecurityChecksTaskResponse](docs/BaseResponseQueuedSecurityChecksTaskResponse.md)
 - [BaseResponseRecent](docs/BaseResponseRecent.md)
 - [BaseResponseSecurityChecksResponse](docs/BaseResponseSecurityChecksResponse.md)
 - [BaseResponseStatus](docs/BaseResponseStatus.md)
 - [BaseResponseStr](docs/BaseResponseStr.md)
 - [BaseResponseTTPS](docs/BaseResponseTTPS.md)
 - [BaseResponseTagSearchResponse](docs/BaseResponseTagSearchResponse.md)
 - [BaseResponseTaskResponse](docs/BaseResponseTaskResponse.md)
 - [BaseResponseUploadResponse](docs/BaseResponseUploadResponse.md)
 - [BaseResponseVulnerabilities](docs/BaseResponseVulnerabilities.md)
 - [Basic](docs/Basic.md)
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
 - [Block](docs/Block.md)
 - [BlockCommentsGenerationForFunctionResponse](docs/BlockCommentsGenerationForFunctionResponse.md)
 - [CalleeFunctionInfo](docs/CalleeFunctionInfo.md)
 - [CalleesCallerFunctionsResponse](docs/CalleesCallerFunctionsResponse.md)
 - [CallerFunctionInfo](docs/CallerFunctionInfo.md)
 - [Capabilities](docs/Capabilities.md)
 - [Capability](docs/Capability.md)
 - [CheckSecurityChecksTaskResponse](docs/CheckSecurityChecksTaskResponse.md)
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
 - [Communities](docs/Communities.md)
 - [CommunityMatchPercentages](docs/CommunityMatchPercentages.md)
 - [ConfidenceType](docs/ConfidenceType.md)
 - [ConfigResponse](docs/ConfigResponse.md)
 - [Context](docs/Context.md)
 - [Created](docs/Created.md)
 - [DecompilationCommentContext](docs/DecompilationCommentContext.md)
 - [DieMatch](docs/DieMatch.md)
 - [DynamicExecutionStatusInput](docs/DynamicExecutionStatusInput.md)
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
 - [ErrorModel](docs/ErrorModel.md)
 - [ExportModel](docs/ExportModel.md)
 - [ExternalResponse](docs/ExternalResponse.md)
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
 - [GenerationStatusList](docs/GenerationStatusList.md)
 - [GetAiDecompilationRatingResponse](docs/GetAiDecompilationRatingResponse.md)
 - [GetAiDecompilationTask](docs/GetAiDecompilationTask.md)
 - [GetMeResponse](docs/GetMeResponse.md)
 - [GetPublicUserResponse](docs/GetPublicUserResponse.md)
 - [GlobalVariable](docs/GlobalVariable.md)
 - [ISA](docs/ISA.md)
 - [IconModel](docs/IconModel.md)
 - [ImportModel](docs/ImportModel.md)
 - [InsertAnalysisLogRequest](docs/InsertAnalysisLogRequest.md)
 - [InverseFunctionMapItem](docs/InverseFunctionMapItem.md)
 - [InverseStringMapItem](docs/InverseStringMapItem.md)
 - [InverseValue](docs/InverseValue.md)
 - [ListCollectionResults](docs/ListCollectionResults.md)
 - [LoginRequest](docs/LoginRequest.md)
 - [LoginResponse](docs/LoginResponse.md)
 - [Logs](docs/Logs.md)
 - [MatchedFunction](docs/MatchedFunction.md)
 - [MatchedFunctionSuggestion](docs/MatchedFunctionSuggestion.md)
 - [MetaModel](docs/MetaModel.md)
 - [ModelName](docs/ModelName.md)
 - [ModelsResponse](docs/ModelsResponse.md)
 - [NameConfidence](docs/NameConfidence.md)
 - [NameSourceType](docs/NameSourceType.md)
 - [NetworkOverviewDns](docs/NetworkOverviewDns.md)
 - [NetworkOverviewDnsAnswer](docs/NetworkOverviewDnsAnswer.md)
 - [NetworkOverviewMetadata](docs/NetworkOverviewMetadata.md)
 - [NetworkOverviewResponse](docs/NetworkOverviewResponse.md)
 - [Order](docs/Order.md)
 - [PDBDebugModel](docs/PDBDebugModel.md)
 - [PEModel](docs/PEModel.md)
 - [PaginationModel](docs/PaginationModel.md)
 - [Params](docs/Params.md)
 - [Platform](docs/Platform.md)
 - [Process](docs/Process.md)
 - [ProcessDump](docs/ProcessDump.md)
 - [ProcessDumpMetadata](docs/ProcessDumpMetadata.md)
 - [ProcessDumps](docs/ProcessDumps.md)
 - [ProcessDumpsData](docs/ProcessDumpsData.md)
 - [ProcessRegistry](docs/ProcessRegistry.md)
 - [ProcessTree](docs/ProcessTree.md)
 - [QueuedSecurityChecksTaskResponse](docs/QueuedSecurityChecksTaskResponse.md)
 - [ReAnalysisForm](docs/ReAnalysisForm.md)
 - [Recent](docs/Recent.md)
 - [Registry](docs/Registry.md)
 - [RelativeBinaryResponse](docs/RelativeBinaryResponse.md)
 - [SBOM](docs/SBOM.md)
 - [SBOMPackage](docs/SBOMPackage.md)
 - [SandboxOptions](docs/SandboxOptions.md)
 - [ScrapeThirdPartyConfig](docs/ScrapeThirdPartyConfig.md)
 - [SectionModel](docs/SectionModel.md)
 - [SecurityChecksResponse](docs/SecurityChecksResponse.md)
 - [SecurityChecksResult](docs/SecurityChecksResult.md)
 - [SecurityModel](docs/SecurityModel.md)
 - [SeverityType](docs/SeverityType.md)
 - [SingleCodeCertificateModel](docs/SingleCodeCertificateModel.md)
 - [SingleCodeSignatureModel](docs/SingleCodeSignatureModel.md)
 - [SinglePDBEntryModel](docs/SinglePDBEntryModel.md)
 - [SingleSectionModel](docs/SingleSectionModel.md)
 - [StackVariable](docs/StackVariable.md)
 - [StatusInput](docs/StatusInput.md)
 - [StatusOutput](docs/StatusOutput.md)
 - [StringFunctions](docs/StringFunctions.md)
 - [Structure](docs/Structure.md)
 - [StructureMember](docs/StructureMember.md)
 - [Symbols](docs/Symbols.md)
 - [TTPS](docs/TTPS.md)
 - [TTPSAttack](docs/TTPSAttack.md)
 - [TTPSData](docs/TTPSData.md)
 - [TTPSElement](docs/TTPSElement.md)
 - [TTPSOccurance](docs/TTPSOccurance.md)
 - [Tag](docs/Tag.md)
 - [TagItem](docs/TagItem.md)
 - [TagResponse](docs/TagResponse.md)
 - [TagSearchResponse](docs/TagSearchResponse.md)
 - [TagSearchResult](docs/TagSearchResult.md)
 - [TaskResponse](docs/TaskResponse.md)
 - [TaskStatus](docs/TaskStatus.md)
 - [TimestampModel](docs/TimestampModel.md)
 - [TypeDefinition](docs/TypeDefinition.md)
 - [UpdateFunctionDataTypes](docs/UpdateFunctionDataTypes.md)
 - [UploadFileType](docs/UploadFileType.md)
 - [UploadResponse](docs/UploadResponse.md)
 - [UpsertAiDecomplationRatingRequest](docs/UpsertAiDecomplationRatingRequest.md)
 - [UserActivityResponse](docs/UserActivityResponse.md)
 - [Vulnerabilities](docs/Vulnerabilities.md)
 - [Vulnerability](docs/Vulnerability.md)
 - [VulnerabilityType](docs/VulnerabilityType.md)
 - [Workspace](docs/Workspace.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### APIKey

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: APIKey and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		sdk.ContextAPIKeys,
		map[string]sdk.APIKey{
			"APIKey": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```


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

## Author



