# Params

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DebugHash** | **NullableString** |  | 
**BinarySize** | **int32** | The size of the binary data | 
**Architecture** | **string** | The architecture of the binary data | 
**BinaryType** | **string** | The type of binary data | 
**BinaryFormat** | **string** | The format of the binary data | 
**BinaryDynamic** | **bool** | Whether the binary data is dynamic | 
**ModelName** | **string** | The name of the model | 

## Methods

### NewParams

`func NewParams(debugHash NullableString, binarySize int32, architecture string, binaryType string, binaryFormat string, binaryDynamic bool, modelName string, ) *Params`

NewParams instantiates a new Params object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParamsWithDefaults

`func NewParamsWithDefaults() *Params`

NewParamsWithDefaults instantiates a new Params object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDebugHash

`func (o *Params) GetDebugHash() string`

GetDebugHash returns the DebugHash field if non-nil, zero value otherwise.

### GetDebugHashOk

`func (o *Params) GetDebugHashOk() (*string, bool)`

GetDebugHashOk returns a tuple with the DebugHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugHash

`func (o *Params) SetDebugHash(v string)`

SetDebugHash sets DebugHash field to given value.


### SetDebugHashNil

`func (o *Params) SetDebugHashNil(b bool)`

 SetDebugHashNil sets the value for DebugHash to be an explicit nil

### UnsetDebugHash
`func (o *Params) UnsetDebugHash()`

UnsetDebugHash ensures that no value is present for DebugHash, not even an explicit nil
### GetBinarySize

`func (o *Params) GetBinarySize() int32`

GetBinarySize returns the BinarySize field if non-nil, zero value otherwise.

### GetBinarySizeOk

`func (o *Params) GetBinarySizeOk() (*int32, bool)`

GetBinarySizeOk returns a tuple with the BinarySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinarySize

`func (o *Params) SetBinarySize(v int32)`

SetBinarySize sets BinarySize field to given value.


### GetArchitecture

`func (o *Params) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *Params) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *Params) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.


### GetBinaryType

`func (o *Params) GetBinaryType() string`

GetBinaryType returns the BinaryType field if non-nil, zero value otherwise.

### GetBinaryTypeOk

`func (o *Params) GetBinaryTypeOk() (*string, bool)`

GetBinaryTypeOk returns a tuple with the BinaryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryType

`func (o *Params) SetBinaryType(v string)`

SetBinaryType sets BinaryType field to given value.


### GetBinaryFormat

`func (o *Params) GetBinaryFormat() string`

GetBinaryFormat returns the BinaryFormat field if non-nil, zero value otherwise.

### GetBinaryFormatOk

`func (o *Params) GetBinaryFormatOk() (*string, bool)`

GetBinaryFormatOk returns a tuple with the BinaryFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryFormat

`func (o *Params) SetBinaryFormat(v string)`

SetBinaryFormat sets BinaryFormat field to given value.


### GetBinaryDynamic

`func (o *Params) GetBinaryDynamic() bool`

GetBinaryDynamic returns the BinaryDynamic field if non-nil, zero value otherwise.

### GetBinaryDynamicOk

`func (o *Params) GetBinaryDynamicOk() (*bool, bool)`

GetBinaryDynamicOk returns a tuple with the BinaryDynamic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryDynamic

`func (o *Params) SetBinaryDynamic(v bool)`

SetBinaryDynamic sets BinaryDynamic field to given value.


### GetModelName

`func (o *Params) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *Params) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *Params) SetModelName(v string)`

SetModelName sets ModelName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


