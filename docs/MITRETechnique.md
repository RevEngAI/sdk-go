# MITRETechnique

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartAddr** | **string** | Starting address of the technique | 
**EndAddr** | **string** | Ending address of the technique | 
**FunctionAddr** | **string** | Function address where the technique is found | 
**TechniqueId** | **string** | MITRE technique identifier | 
**TechniqueName** | **string** | Name of the MITRE technique | 
**Description** | **string** | Description of the technique | 
**FunctionId** | **int32** | Unique identifier of the function containing the technique | 
**FunctionName** | **string** | Name of the function containing the technique | 
**TechniqueUrl** | **string** | URL to the MITRE ATT&amp;CK technique page | 
**TechniqueDescription** | **string** | Full description of the MITRE technique from ATT&amp;CK | 

## Methods

### NewMITRETechnique

`func NewMITRETechnique(startAddr string, endAddr string, functionAddr string, techniqueId string, techniqueName string, description string, functionId int32, functionName string, techniqueUrl string, techniqueDescription string, ) *MITRETechnique`

NewMITRETechnique instantiates a new MITRETechnique object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMITRETechniqueWithDefaults

`func NewMITRETechniqueWithDefaults() *MITRETechnique`

NewMITRETechniqueWithDefaults instantiates a new MITRETechnique object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartAddr

`func (o *MITRETechnique) GetStartAddr() string`

GetStartAddr returns the StartAddr field if non-nil, zero value otherwise.

### GetStartAddrOk

`func (o *MITRETechnique) GetStartAddrOk() (*string, bool)`

GetStartAddrOk returns a tuple with the StartAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAddr

`func (o *MITRETechnique) SetStartAddr(v string)`

SetStartAddr sets StartAddr field to given value.


### GetEndAddr

`func (o *MITRETechnique) GetEndAddr() string`

GetEndAddr returns the EndAddr field if non-nil, zero value otherwise.

### GetEndAddrOk

`func (o *MITRETechnique) GetEndAddrOk() (*string, bool)`

GetEndAddrOk returns a tuple with the EndAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAddr

`func (o *MITRETechnique) SetEndAddr(v string)`

SetEndAddr sets EndAddr field to given value.


### GetFunctionAddr

`func (o *MITRETechnique) GetFunctionAddr() string`

GetFunctionAddr returns the FunctionAddr field if non-nil, zero value otherwise.

### GetFunctionAddrOk

`func (o *MITRETechnique) GetFunctionAddrOk() (*string, bool)`

GetFunctionAddrOk returns a tuple with the FunctionAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionAddr

`func (o *MITRETechnique) SetFunctionAddr(v string)`

SetFunctionAddr sets FunctionAddr field to given value.


### GetTechniqueId

`func (o *MITRETechnique) GetTechniqueId() string`

GetTechniqueId returns the TechniqueId field if non-nil, zero value otherwise.

### GetTechniqueIdOk

`func (o *MITRETechnique) GetTechniqueIdOk() (*string, bool)`

GetTechniqueIdOk returns a tuple with the TechniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechniqueId

`func (o *MITRETechnique) SetTechniqueId(v string)`

SetTechniqueId sets TechniqueId field to given value.


### GetTechniqueName

`func (o *MITRETechnique) GetTechniqueName() string`

GetTechniqueName returns the TechniqueName field if non-nil, zero value otherwise.

### GetTechniqueNameOk

`func (o *MITRETechnique) GetTechniqueNameOk() (*string, bool)`

GetTechniqueNameOk returns a tuple with the TechniqueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechniqueName

`func (o *MITRETechnique) SetTechniqueName(v string)`

SetTechniqueName sets TechniqueName field to given value.


### GetDescription

`func (o *MITRETechnique) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MITRETechnique) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MITRETechnique) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetFunctionId

`func (o *MITRETechnique) GetFunctionId() int32`

GetFunctionId returns the FunctionId field if non-nil, zero value otherwise.

### GetFunctionIdOk

`func (o *MITRETechnique) GetFunctionIdOk() (*int32, bool)`

GetFunctionIdOk returns a tuple with the FunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionId

`func (o *MITRETechnique) SetFunctionId(v int32)`

SetFunctionId sets FunctionId field to given value.


### GetFunctionName

`func (o *MITRETechnique) GetFunctionName() string`

GetFunctionName returns the FunctionName field if non-nil, zero value otherwise.

### GetFunctionNameOk

`func (o *MITRETechnique) GetFunctionNameOk() (*string, bool)`

GetFunctionNameOk returns a tuple with the FunctionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionName

`func (o *MITRETechnique) SetFunctionName(v string)`

SetFunctionName sets FunctionName field to given value.


### GetTechniqueUrl

`func (o *MITRETechnique) GetTechniqueUrl() string`

GetTechniqueUrl returns the TechniqueUrl field if non-nil, zero value otherwise.

### GetTechniqueUrlOk

`func (o *MITRETechnique) GetTechniqueUrlOk() (*string, bool)`

GetTechniqueUrlOk returns a tuple with the TechniqueUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechniqueUrl

`func (o *MITRETechnique) SetTechniqueUrl(v string)`

SetTechniqueUrl sets TechniqueUrl field to given value.


### GetTechniqueDescription

`func (o *MITRETechnique) GetTechniqueDescription() string`

GetTechniqueDescription returns the TechniqueDescription field if non-nil, zero value otherwise.

### GetTechniqueDescriptionOk

`func (o *MITRETechnique) GetTechniqueDescriptionOk() (*string, bool)`

GetTechniqueDescriptionOk returns a tuple with the TechniqueDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechniqueDescription

`func (o *MITRETechnique) SetTechniqueDescription(v string)`

SetTechniqueDescription sets TechniqueDescription field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


