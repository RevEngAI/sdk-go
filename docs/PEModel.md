# PEModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Timestamps** | [**NullableTimestampModel**](TimestampModel.md) |  | 
**Architecture** | **string** |  | 
**Checksum** | **int32** |  | 
**ImageBase** | **int32** |  | 
**Security** | [**NullableSecurityModel**](SecurityModel.md) |  | 
**VersionInfo** | **map[string]interface{}** |  | 
**DebugInfo** | [**NullablePDBDebugModel**](PDBDebugModel.md) |  | 
**NumberOfResources** | **NullableInt32** |  | 
**EntryPoint** | [**NullableEntrypointModel**](EntrypointModel.md) |  | 
**Signature** | [**NullableCodeSignatureModel**](CodeSignatureModel.md) |  | 
**Dotnet** | **bool** |  | 
**DebugStripped** | **bool** |  | 
**ImportHash** | **string** |  | 
**ExportHash** | **string** |  | 
**RichHeaderHash** | **string** |  | 
**Sections** | [**NullableSectionModel**](SectionModel.md) |  | 
**Imports** | [**NullableImportModel**](ImportModel.md) |  | 
**Exports** | [**NullableExportModel**](ExportModel.md) |  | 
**IconData** | [**NullableIconModel**](IconModel.md) |  | 

## Methods

### NewPEModel

`func NewPEModel(type_ string, timestamps NullableTimestampModel, architecture string, checksum int32, imageBase int32, security NullableSecurityModel, versionInfo map[string]interface{}, debugInfo NullablePDBDebugModel, numberOfResources NullableInt32, entryPoint NullableEntrypointModel, signature NullableCodeSignatureModel, dotnet bool, debugStripped bool, importHash string, exportHash string, richHeaderHash string, sections NullableSectionModel, imports NullableImportModel, exports NullableExportModel, iconData NullableIconModel, ) *PEModel`

NewPEModel instantiates a new PEModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPEModelWithDefaults

`func NewPEModelWithDefaults() *PEModel`

NewPEModelWithDefaults instantiates a new PEModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PEModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PEModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PEModel) SetType(v string)`

SetType sets Type field to given value.


### GetTimestamps

`func (o *PEModel) GetTimestamps() TimestampModel`

GetTimestamps returns the Timestamps field if non-nil, zero value otherwise.

### GetTimestampsOk

`func (o *PEModel) GetTimestampsOk() (*TimestampModel, bool)`

GetTimestampsOk returns a tuple with the Timestamps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamps

`func (o *PEModel) SetTimestamps(v TimestampModel)`

SetTimestamps sets Timestamps field to given value.


### SetTimestampsNil

`func (o *PEModel) SetTimestampsNil(b bool)`

 SetTimestampsNil sets the value for Timestamps to be an explicit nil

### UnsetTimestamps
`func (o *PEModel) UnsetTimestamps()`

UnsetTimestamps ensures that no value is present for Timestamps, not even an explicit nil
### GetArchitecture

`func (o *PEModel) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *PEModel) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *PEModel) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.


### GetChecksum

`func (o *PEModel) GetChecksum() int32`

GetChecksum returns the Checksum field if non-nil, zero value otherwise.

### GetChecksumOk

`func (o *PEModel) GetChecksumOk() (*int32, bool)`

GetChecksumOk returns a tuple with the Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksum

`func (o *PEModel) SetChecksum(v int32)`

SetChecksum sets Checksum field to given value.


### GetImageBase

`func (o *PEModel) GetImageBase() int32`

GetImageBase returns the ImageBase field if non-nil, zero value otherwise.

### GetImageBaseOk

`func (o *PEModel) GetImageBaseOk() (*int32, bool)`

GetImageBaseOk returns a tuple with the ImageBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageBase

`func (o *PEModel) SetImageBase(v int32)`

SetImageBase sets ImageBase field to given value.


### GetSecurity

`func (o *PEModel) GetSecurity() SecurityModel`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *PEModel) GetSecurityOk() (*SecurityModel, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *PEModel) SetSecurity(v SecurityModel)`

SetSecurity sets Security field to given value.


### SetSecurityNil

`func (o *PEModel) SetSecurityNil(b bool)`

 SetSecurityNil sets the value for Security to be an explicit nil

### UnsetSecurity
`func (o *PEModel) UnsetSecurity()`

UnsetSecurity ensures that no value is present for Security, not even an explicit nil
### GetVersionInfo

`func (o *PEModel) GetVersionInfo() map[string]interface{}`

GetVersionInfo returns the VersionInfo field if non-nil, zero value otherwise.

### GetVersionInfoOk

`func (o *PEModel) GetVersionInfoOk() (*map[string]interface{}, bool)`

GetVersionInfoOk returns a tuple with the VersionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionInfo

`func (o *PEModel) SetVersionInfo(v map[string]interface{})`

SetVersionInfo sets VersionInfo field to given value.


### SetVersionInfoNil

`func (o *PEModel) SetVersionInfoNil(b bool)`

 SetVersionInfoNil sets the value for VersionInfo to be an explicit nil

### UnsetVersionInfo
`func (o *PEModel) UnsetVersionInfo()`

UnsetVersionInfo ensures that no value is present for VersionInfo, not even an explicit nil
### GetDebugInfo

`func (o *PEModel) GetDebugInfo() PDBDebugModel`

GetDebugInfo returns the DebugInfo field if non-nil, zero value otherwise.

### GetDebugInfoOk

`func (o *PEModel) GetDebugInfoOk() (*PDBDebugModel, bool)`

GetDebugInfoOk returns a tuple with the DebugInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugInfo

`func (o *PEModel) SetDebugInfo(v PDBDebugModel)`

SetDebugInfo sets DebugInfo field to given value.


### SetDebugInfoNil

`func (o *PEModel) SetDebugInfoNil(b bool)`

 SetDebugInfoNil sets the value for DebugInfo to be an explicit nil

### UnsetDebugInfo
`func (o *PEModel) UnsetDebugInfo()`

UnsetDebugInfo ensures that no value is present for DebugInfo, not even an explicit nil
### GetNumberOfResources

`func (o *PEModel) GetNumberOfResources() int32`

GetNumberOfResources returns the NumberOfResources field if non-nil, zero value otherwise.

### GetNumberOfResourcesOk

`func (o *PEModel) GetNumberOfResourcesOk() (*int32, bool)`

GetNumberOfResourcesOk returns a tuple with the NumberOfResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfResources

`func (o *PEModel) SetNumberOfResources(v int32)`

SetNumberOfResources sets NumberOfResources field to given value.


### SetNumberOfResourcesNil

`func (o *PEModel) SetNumberOfResourcesNil(b bool)`

 SetNumberOfResourcesNil sets the value for NumberOfResources to be an explicit nil

### UnsetNumberOfResources
`func (o *PEModel) UnsetNumberOfResources()`

UnsetNumberOfResources ensures that no value is present for NumberOfResources, not even an explicit nil
### GetEntryPoint

`func (o *PEModel) GetEntryPoint() EntrypointModel`

GetEntryPoint returns the EntryPoint field if non-nil, zero value otherwise.

### GetEntryPointOk

`func (o *PEModel) GetEntryPointOk() (*EntrypointModel, bool)`

GetEntryPointOk returns a tuple with the EntryPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryPoint

`func (o *PEModel) SetEntryPoint(v EntrypointModel)`

SetEntryPoint sets EntryPoint field to given value.


### SetEntryPointNil

`func (o *PEModel) SetEntryPointNil(b bool)`

 SetEntryPointNil sets the value for EntryPoint to be an explicit nil

### UnsetEntryPoint
`func (o *PEModel) UnsetEntryPoint()`

UnsetEntryPoint ensures that no value is present for EntryPoint, not even an explicit nil
### GetSignature

`func (o *PEModel) GetSignature() CodeSignatureModel`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *PEModel) GetSignatureOk() (*CodeSignatureModel, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *PEModel) SetSignature(v CodeSignatureModel)`

SetSignature sets Signature field to given value.


### SetSignatureNil

`func (o *PEModel) SetSignatureNil(b bool)`

 SetSignatureNil sets the value for Signature to be an explicit nil

### UnsetSignature
`func (o *PEModel) UnsetSignature()`

UnsetSignature ensures that no value is present for Signature, not even an explicit nil
### GetDotnet

`func (o *PEModel) GetDotnet() bool`

GetDotnet returns the Dotnet field if non-nil, zero value otherwise.

### GetDotnetOk

`func (o *PEModel) GetDotnetOk() (*bool, bool)`

GetDotnetOk returns a tuple with the Dotnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDotnet

`func (o *PEModel) SetDotnet(v bool)`

SetDotnet sets Dotnet field to given value.


### GetDebugStripped

`func (o *PEModel) GetDebugStripped() bool`

GetDebugStripped returns the DebugStripped field if non-nil, zero value otherwise.

### GetDebugStrippedOk

`func (o *PEModel) GetDebugStrippedOk() (*bool, bool)`

GetDebugStrippedOk returns a tuple with the DebugStripped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugStripped

`func (o *PEModel) SetDebugStripped(v bool)`

SetDebugStripped sets DebugStripped field to given value.


### GetImportHash

`func (o *PEModel) GetImportHash() string`

GetImportHash returns the ImportHash field if non-nil, zero value otherwise.

### GetImportHashOk

`func (o *PEModel) GetImportHashOk() (*string, bool)`

GetImportHashOk returns a tuple with the ImportHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportHash

`func (o *PEModel) SetImportHash(v string)`

SetImportHash sets ImportHash field to given value.


### GetExportHash

`func (o *PEModel) GetExportHash() string`

GetExportHash returns the ExportHash field if non-nil, zero value otherwise.

### GetExportHashOk

`func (o *PEModel) GetExportHashOk() (*string, bool)`

GetExportHashOk returns a tuple with the ExportHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportHash

`func (o *PEModel) SetExportHash(v string)`

SetExportHash sets ExportHash field to given value.


### GetRichHeaderHash

`func (o *PEModel) GetRichHeaderHash() string`

GetRichHeaderHash returns the RichHeaderHash field if non-nil, zero value otherwise.

### GetRichHeaderHashOk

`func (o *PEModel) GetRichHeaderHashOk() (*string, bool)`

GetRichHeaderHashOk returns a tuple with the RichHeaderHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRichHeaderHash

`func (o *PEModel) SetRichHeaderHash(v string)`

SetRichHeaderHash sets RichHeaderHash field to given value.


### GetSections

`func (o *PEModel) GetSections() SectionModel`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *PEModel) GetSectionsOk() (*SectionModel, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *PEModel) SetSections(v SectionModel)`

SetSections sets Sections field to given value.


### SetSectionsNil

`func (o *PEModel) SetSectionsNil(b bool)`

 SetSectionsNil sets the value for Sections to be an explicit nil

### UnsetSections
`func (o *PEModel) UnsetSections()`

UnsetSections ensures that no value is present for Sections, not even an explicit nil
### GetImports

`func (o *PEModel) GetImports() ImportModel`

GetImports returns the Imports field if non-nil, zero value otherwise.

### GetImportsOk

`func (o *PEModel) GetImportsOk() (*ImportModel, bool)`

GetImportsOk returns a tuple with the Imports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImports

`func (o *PEModel) SetImports(v ImportModel)`

SetImports sets Imports field to given value.


### SetImportsNil

`func (o *PEModel) SetImportsNil(b bool)`

 SetImportsNil sets the value for Imports to be an explicit nil

### UnsetImports
`func (o *PEModel) UnsetImports()`

UnsetImports ensures that no value is present for Imports, not even an explicit nil
### GetExports

`func (o *PEModel) GetExports() ExportModel`

GetExports returns the Exports field if non-nil, zero value otherwise.

### GetExportsOk

`func (o *PEModel) GetExportsOk() (*ExportModel, bool)`

GetExportsOk returns a tuple with the Exports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExports

`func (o *PEModel) SetExports(v ExportModel)`

SetExports sets Exports field to given value.


### SetExportsNil

`func (o *PEModel) SetExportsNil(b bool)`

 SetExportsNil sets the value for Exports to be an explicit nil

### UnsetExports
`func (o *PEModel) UnsetExports()`

UnsetExports ensures that no value is present for Exports, not even an explicit nil
### GetIconData

`func (o *PEModel) GetIconData() IconModel`

GetIconData returns the IconData field if non-nil, zero value otherwise.

### GetIconDataOk

`func (o *PEModel) GetIconDataOk() (*IconModel, bool)`

GetIconDataOk returns a tuple with the IconData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconData

`func (o *PEModel) SetIconData(v IconModel)`

SetIconData sets IconData field to given value.


### SetIconDataNil

`func (o *PEModel) SetIconDataNil(b bool)`

 SetIconDataNil sets the value for IconData to be an explicit nil

### UnsetIconData
`func (o *PEModel) UnsetIconData()`

UnsetIconData ensures that no value is present for IconData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


