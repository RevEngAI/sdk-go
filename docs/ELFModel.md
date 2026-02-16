# ELFModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileType** | **string** |  | 
**Architecture** | **string** |  | 
**Endianness** | **string** |  | 
**EntryPoint** | **int32** |  | 
**EntryPointBytes** | **string** |  | 
**ImportHash** | **string** |  | 
**ExportHash** | **string** |  | 
**BuildId** | **string** |  | 
**Security** | [**ELFSecurity**](ELFSecurity.md) |  | 
**Sections** | [**[]ELFSection**](ELFSection.md) |  | 
**Segments** | [**[]ELFSegment**](ELFSegment.md) |  | 
**Symbols** | [**[]ELFSymbol**](ELFSymbol.md) |  | 
**DynamicSymbols** | [**[]ELFSymbol**](ELFSymbol.md) |  | 
**Relocations** | [**[]ELFRelocation**](ELFRelocation.md) |  | 
**Imports** | [**ELFImportModel**](ELFImportModel.md) |  | 
**ExportedFunctions** | **[]string** |  | 
**DynamicEntries** | [**[]ElfDynamicEntry**](ElfDynamicEntry.md) |  | 
**Notes** | **[]map[string]interface{}** |  | 
**DebugInfo** | **map[string]interface{}** |  | 
**VersionInfo** | **map[string]interface{}** |  | 

## Methods

### NewELFModel

`func NewELFModel(fileType string, architecture string, endianness string, entryPoint int32, entryPointBytes string, importHash string, exportHash string, buildId string, security ELFSecurity, sections []ELFSection, segments []ELFSegment, symbols []ELFSymbol, dynamicSymbols []ELFSymbol, relocations []ELFRelocation, imports ELFImportModel, exportedFunctions []string, dynamicEntries []ElfDynamicEntry, notes []map[string]interface{}, debugInfo map[string]interface{}, versionInfo map[string]interface{}, ) *ELFModel`

NewELFModel instantiates a new ELFModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewELFModelWithDefaults

`func NewELFModelWithDefaults() *ELFModel`

NewELFModelWithDefaults instantiates a new ELFModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileType

`func (o *ELFModel) GetFileType() string`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *ELFModel) GetFileTypeOk() (*string, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *ELFModel) SetFileType(v string)`

SetFileType sets FileType field to given value.


### GetArchitecture

`func (o *ELFModel) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *ELFModel) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *ELFModel) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.


### GetEndianness

`func (o *ELFModel) GetEndianness() string`

GetEndianness returns the Endianness field if non-nil, zero value otherwise.

### GetEndiannessOk

`func (o *ELFModel) GetEndiannessOk() (*string, bool)`

GetEndiannessOk returns a tuple with the Endianness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndianness

`func (o *ELFModel) SetEndianness(v string)`

SetEndianness sets Endianness field to given value.


### GetEntryPoint

`func (o *ELFModel) GetEntryPoint() int32`

GetEntryPoint returns the EntryPoint field if non-nil, zero value otherwise.

### GetEntryPointOk

`func (o *ELFModel) GetEntryPointOk() (*int32, bool)`

GetEntryPointOk returns a tuple with the EntryPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryPoint

`func (o *ELFModel) SetEntryPoint(v int32)`

SetEntryPoint sets EntryPoint field to given value.


### GetEntryPointBytes

`func (o *ELFModel) GetEntryPointBytes() string`

GetEntryPointBytes returns the EntryPointBytes field if non-nil, zero value otherwise.

### GetEntryPointBytesOk

`func (o *ELFModel) GetEntryPointBytesOk() (*string, bool)`

GetEntryPointBytesOk returns a tuple with the EntryPointBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryPointBytes

`func (o *ELFModel) SetEntryPointBytes(v string)`

SetEntryPointBytes sets EntryPointBytes field to given value.


### GetImportHash

`func (o *ELFModel) GetImportHash() string`

GetImportHash returns the ImportHash field if non-nil, zero value otherwise.

### GetImportHashOk

`func (o *ELFModel) GetImportHashOk() (*string, bool)`

GetImportHashOk returns a tuple with the ImportHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportHash

`func (o *ELFModel) SetImportHash(v string)`

SetImportHash sets ImportHash field to given value.


### GetExportHash

`func (o *ELFModel) GetExportHash() string`

GetExportHash returns the ExportHash field if non-nil, zero value otherwise.

### GetExportHashOk

`func (o *ELFModel) GetExportHashOk() (*string, bool)`

GetExportHashOk returns a tuple with the ExportHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportHash

`func (o *ELFModel) SetExportHash(v string)`

SetExportHash sets ExportHash field to given value.


### GetBuildId

`func (o *ELFModel) GetBuildId() string`

GetBuildId returns the BuildId field if non-nil, zero value otherwise.

### GetBuildIdOk

`func (o *ELFModel) GetBuildIdOk() (*string, bool)`

GetBuildIdOk returns a tuple with the BuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildId

`func (o *ELFModel) SetBuildId(v string)`

SetBuildId sets BuildId field to given value.


### GetSecurity

`func (o *ELFModel) GetSecurity() ELFSecurity`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *ELFModel) GetSecurityOk() (*ELFSecurity, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *ELFModel) SetSecurity(v ELFSecurity)`

SetSecurity sets Security field to given value.


### GetSections

`func (o *ELFModel) GetSections() []ELFSection`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *ELFModel) GetSectionsOk() (*[]ELFSection, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *ELFModel) SetSections(v []ELFSection)`

SetSections sets Sections field to given value.


### GetSegments

`func (o *ELFModel) GetSegments() []ELFSegment`

GetSegments returns the Segments field if non-nil, zero value otherwise.

### GetSegmentsOk

`func (o *ELFModel) GetSegmentsOk() (*[]ELFSegment, bool)`

GetSegmentsOk returns a tuple with the Segments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegments

`func (o *ELFModel) SetSegments(v []ELFSegment)`

SetSegments sets Segments field to given value.


### GetSymbols

`func (o *ELFModel) GetSymbols() []ELFSymbol`

GetSymbols returns the Symbols field if non-nil, zero value otherwise.

### GetSymbolsOk

`func (o *ELFModel) GetSymbolsOk() (*[]ELFSymbol, bool)`

GetSymbolsOk returns a tuple with the Symbols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbols

`func (o *ELFModel) SetSymbols(v []ELFSymbol)`

SetSymbols sets Symbols field to given value.


### GetDynamicSymbols

`func (o *ELFModel) GetDynamicSymbols() []ELFSymbol`

GetDynamicSymbols returns the DynamicSymbols field if non-nil, zero value otherwise.

### GetDynamicSymbolsOk

`func (o *ELFModel) GetDynamicSymbolsOk() (*[]ELFSymbol, bool)`

GetDynamicSymbolsOk returns a tuple with the DynamicSymbols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicSymbols

`func (o *ELFModel) SetDynamicSymbols(v []ELFSymbol)`

SetDynamicSymbols sets DynamicSymbols field to given value.


### GetRelocations

`func (o *ELFModel) GetRelocations() []ELFRelocation`

GetRelocations returns the Relocations field if non-nil, zero value otherwise.

### GetRelocationsOk

`func (o *ELFModel) GetRelocationsOk() (*[]ELFRelocation, bool)`

GetRelocationsOk returns a tuple with the Relocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelocations

`func (o *ELFModel) SetRelocations(v []ELFRelocation)`

SetRelocations sets Relocations field to given value.


### GetImports

`func (o *ELFModel) GetImports() ELFImportModel`

GetImports returns the Imports field if non-nil, zero value otherwise.

### GetImportsOk

`func (o *ELFModel) GetImportsOk() (*ELFImportModel, bool)`

GetImportsOk returns a tuple with the Imports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImports

`func (o *ELFModel) SetImports(v ELFImportModel)`

SetImports sets Imports field to given value.


### GetExportedFunctions

`func (o *ELFModel) GetExportedFunctions() []string`

GetExportedFunctions returns the ExportedFunctions field if non-nil, zero value otherwise.

### GetExportedFunctionsOk

`func (o *ELFModel) GetExportedFunctionsOk() (*[]string, bool)`

GetExportedFunctionsOk returns a tuple with the ExportedFunctions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportedFunctions

`func (o *ELFModel) SetExportedFunctions(v []string)`

SetExportedFunctions sets ExportedFunctions field to given value.


### GetDynamicEntries

`func (o *ELFModel) GetDynamicEntries() []ElfDynamicEntry`

GetDynamicEntries returns the DynamicEntries field if non-nil, zero value otherwise.

### GetDynamicEntriesOk

`func (o *ELFModel) GetDynamicEntriesOk() (*[]ElfDynamicEntry, bool)`

GetDynamicEntriesOk returns a tuple with the DynamicEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicEntries

`func (o *ELFModel) SetDynamicEntries(v []ElfDynamicEntry)`

SetDynamicEntries sets DynamicEntries field to given value.


### GetNotes

`func (o *ELFModel) GetNotes() []map[string]interface{}`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ELFModel) GetNotesOk() (*[]map[string]interface{}, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ELFModel) SetNotes(v []map[string]interface{})`

SetNotes sets Notes field to given value.


### GetDebugInfo

`func (o *ELFModel) GetDebugInfo() map[string]interface{}`

GetDebugInfo returns the DebugInfo field if non-nil, zero value otherwise.

### GetDebugInfoOk

`func (o *ELFModel) GetDebugInfoOk() (*map[string]interface{}, bool)`

GetDebugInfoOk returns a tuple with the DebugInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugInfo

`func (o *ELFModel) SetDebugInfo(v map[string]interface{})`

SetDebugInfo sets DebugInfo field to given value.


### GetVersionInfo

`func (o *ELFModel) GetVersionInfo() map[string]interface{}`

GetVersionInfo returns the VersionInfo field if non-nil, zero value otherwise.

### GetVersionInfoOk

`func (o *ELFModel) GetVersionInfoOk() (*map[string]interface{}, bool)`

GetVersionInfoOk returns a tuple with the VersionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionInfo

`func (o *ELFModel) SetVersionInfo(v map[string]interface{})`

SetVersionInfo sets VersionInfo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


