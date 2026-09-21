# Technique

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | How the technique appears in this binary | 
**EndAddr** | **string** | End address of the containing function, hex-encoded | 
**FunctionAddr** | **string** | Address of the containing function, hex-encoded | 
**FunctionId** | **int64** | ID of the containing function | 
**FunctionName** | **string** | Name of the containing function | 
**StartAddr** | **string** | Start address of the containing function, hex-encoded | 
**TechniqueDescription** | **string** | Full ATT&amp;CK description of the technique | 
**TechniqueId** | **string** | MITRE ATT&amp;CK technique ID | 
**TechniqueName** | **string** | MITRE ATT&amp;CK technique name | 
**TechniqueUrl** | **string** | Link to the technique in the ATT&amp;CK catalogue | 

## Methods

### NewTechnique

`func NewTechnique(description string, endAddr string, functionAddr string, functionId int64, functionName string, startAddr string, techniqueDescription string, techniqueId string, techniqueName string, techniqueUrl string, ) *Technique`

NewTechnique instantiates a new Technique object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTechniqueWithDefaults

`func NewTechniqueWithDefaults() *Technique`

NewTechniqueWithDefaults instantiates a new Technique object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *Technique) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Technique) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Technique) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEndAddr

`func (o *Technique) GetEndAddr() string`

GetEndAddr returns the EndAddr field if non-nil, zero value otherwise.

### GetEndAddrOk

`func (o *Technique) GetEndAddrOk() (*string, bool)`

GetEndAddrOk returns a tuple with the EndAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAddr

`func (o *Technique) SetEndAddr(v string)`

SetEndAddr sets EndAddr field to given value.


### GetFunctionAddr

`func (o *Technique) GetFunctionAddr() string`

GetFunctionAddr returns the FunctionAddr field if non-nil, zero value otherwise.

### GetFunctionAddrOk

`func (o *Technique) GetFunctionAddrOk() (*string, bool)`

GetFunctionAddrOk returns a tuple with the FunctionAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionAddr

`func (o *Technique) SetFunctionAddr(v string)`

SetFunctionAddr sets FunctionAddr field to given value.


### GetFunctionId

`func (o *Technique) GetFunctionId() int64`

GetFunctionId returns the FunctionId field if non-nil, zero value otherwise.

### GetFunctionIdOk

`func (o *Technique) GetFunctionIdOk() (*int64, bool)`

GetFunctionIdOk returns a tuple with the FunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionId

`func (o *Technique) SetFunctionId(v int64)`

SetFunctionId sets FunctionId field to given value.


### GetFunctionName

`func (o *Technique) GetFunctionName() string`

GetFunctionName returns the FunctionName field if non-nil, zero value otherwise.

### GetFunctionNameOk

`func (o *Technique) GetFunctionNameOk() (*string, bool)`

GetFunctionNameOk returns a tuple with the FunctionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionName

`func (o *Technique) SetFunctionName(v string)`

SetFunctionName sets FunctionName field to given value.


### GetStartAddr

`func (o *Technique) GetStartAddr() string`

GetStartAddr returns the StartAddr field if non-nil, zero value otherwise.

### GetStartAddrOk

`func (o *Technique) GetStartAddrOk() (*string, bool)`

GetStartAddrOk returns a tuple with the StartAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAddr

`func (o *Technique) SetStartAddr(v string)`

SetStartAddr sets StartAddr field to given value.


### GetTechniqueDescription

`func (o *Technique) GetTechniqueDescription() string`

GetTechniqueDescription returns the TechniqueDescription field if non-nil, zero value otherwise.

### GetTechniqueDescriptionOk

`func (o *Technique) GetTechniqueDescriptionOk() (*string, bool)`

GetTechniqueDescriptionOk returns a tuple with the TechniqueDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechniqueDescription

`func (o *Technique) SetTechniqueDescription(v string)`

SetTechniqueDescription sets TechniqueDescription field to given value.


### GetTechniqueId

`func (o *Technique) GetTechniqueId() string`

GetTechniqueId returns the TechniqueId field if non-nil, zero value otherwise.

### GetTechniqueIdOk

`func (o *Technique) GetTechniqueIdOk() (*string, bool)`

GetTechniqueIdOk returns a tuple with the TechniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechniqueId

`func (o *Technique) SetTechniqueId(v string)`

SetTechniqueId sets TechniqueId field to given value.


### GetTechniqueName

`func (o *Technique) GetTechniqueName() string`

GetTechniqueName returns the TechniqueName field if non-nil, zero value otherwise.

### GetTechniqueNameOk

`func (o *Technique) GetTechniqueNameOk() (*string, bool)`

GetTechniqueNameOk returns a tuple with the TechniqueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechniqueName

`func (o *Technique) SetTechniqueName(v string)`

SetTechniqueName sets TechniqueName field to given value.


### GetTechniqueUrl

`func (o *Technique) GetTechniqueUrl() string`

GetTechniqueUrl returns the TechniqueUrl field if non-nil, zero value otherwise.

### GetTechniqueUrlOk

`func (o *Technique) GetTechniqueUrlOk() (*string, bool)`

GetTechniqueUrlOk returns a tuple with the TechniqueUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechniqueUrl

`func (o *Technique) SetTechniqueUrl(v string)`

SetTechniqueUrl sets TechniqueUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


