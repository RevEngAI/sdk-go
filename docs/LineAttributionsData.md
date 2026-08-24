# LineAttributionsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisassemblyLineNumberToAiDecompilationLineNumbers** | **map[string][]int64** | Each disassembly line number mapped to the AI-decompilation line numbers it fed, e.g. {\&quot;12\&quot;: [3, 4, 6], \&quot;17\&quot;: [4]}. Both sides 0-based; many-to-many in both directions. Empty when no completed run has produced a correspondence, which is ordinary and not an error. | 

## Methods

### NewLineAttributionsData

`func NewLineAttributionsData(disassemblyLineNumberToAiDecompilationLineNumbers map[string][]int64, ) *LineAttributionsData`

NewLineAttributionsData instantiates a new LineAttributionsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLineAttributionsDataWithDefaults

`func NewLineAttributionsDataWithDefaults() *LineAttributionsData`

NewLineAttributionsDataWithDefaults instantiates a new LineAttributionsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisassemblyLineNumberToAiDecompilationLineNumbers

`func (o *LineAttributionsData) GetDisassemblyLineNumberToAiDecompilationLineNumbers() map[string][]int64`

GetDisassemblyLineNumberToAiDecompilationLineNumbers returns the DisassemblyLineNumberToAiDecompilationLineNumbers field if non-nil, zero value otherwise.

### GetDisassemblyLineNumberToAiDecompilationLineNumbersOk

`func (o *LineAttributionsData) GetDisassemblyLineNumberToAiDecompilationLineNumbersOk() (*map[string][]int64, bool)`

GetDisassemblyLineNumberToAiDecompilationLineNumbersOk returns a tuple with the DisassemblyLineNumberToAiDecompilationLineNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisassemblyLineNumberToAiDecompilationLineNumbers

`func (o *LineAttributionsData) SetDisassemblyLineNumberToAiDecompilationLineNumbers(v map[string][]int64)`

SetDisassemblyLineNumberToAiDecompilationLineNumbers sets DisassemblyLineNumberToAiDecompilationLineNumbers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


