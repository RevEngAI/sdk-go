# FunctionBoundary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndAddress** | **int64** |  | 
**IncludeInAnalysis** | Pointer to **NullableBool** |  | [optional] 
**MangledName** | **string** |  | 
**StartAddress** | **int64** |  | 

## Methods

### NewFunctionBoundary

`func NewFunctionBoundary(endAddress int64, mangledName string, startAddress int64, ) *FunctionBoundary`

NewFunctionBoundary instantiates a new FunctionBoundary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionBoundaryWithDefaults

`func NewFunctionBoundaryWithDefaults() *FunctionBoundary`

NewFunctionBoundaryWithDefaults instantiates a new FunctionBoundary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndAddress

`func (o *FunctionBoundary) GetEndAddress() int64`

GetEndAddress returns the EndAddress field if non-nil, zero value otherwise.

### GetEndAddressOk

`func (o *FunctionBoundary) GetEndAddressOk() (*int64, bool)`

GetEndAddressOk returns a tuple with the EndAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAddress

`func (o *FunctionBoundary) SetEndAddress(v int64)`

SetEndAddress sets EndAddress field to given value.


### GetIncludeInAnalysis

`func (o *FunctionBoundary) GetIncludeInAnalysis() bool`

GetIncludeInAnalysis returns the IncludeInAnalysis field if non-nil, zero value otherwise.

### GetIncludeInAnalysisOk

`func (o *FunctionBoundary) GetIncludeInAnalysisOk() (*bool, bool)`

GetIncludeInAnalysisOk returns a tuple with the IncludeInAnalysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeInAnalysis

`func (o *FunctionBoundary) SetIncludeInAnalysis(v bool)`

SetIncludeInAnalysis sets IncludeInAnalysis field to given value.

### HasIncludeInAnalysis

`func (o *FunctionBoundary) HasIncludeInAnalysis() bool`

HasIncludeInAnalysis returns a boolean if a field has been set.

### SetIncludeInAnalysisNil

`func (o *FunctionBoundary) SetIncludeInAnalysisNil(b bool)`

 SetIncludeInAnalysisNil sets the value for IncludeInAnalysis to be an explicit nil

### UnsetIncludeInAnalysis
`func (o *FunctionBoundary) UnsetIncludeInAnalysis()`

UnsetIncludeInAnalysis ensures that no value is present for IncludeInAnalysis, not even an explicit nil
### GetMangledName

`func (o *FunctionBoundary) GetMangledName() string`

GetMangledName returns the MangledName field if non-nil, zero value otherwise.

### GetMangledNameOk

`func (o *FunctionBoundary) GetMangledNameOk() (*string, bool)`

GetMangledNameOk returns a tuple with the MangledName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMangledName

`func (o *FunctionBoundary) SetMangledName(v string)`

SetMangledName sets MangledName field to given value.


### GetStartAddress

`func (o *FunctionBoundary) GetStartAddress() int64`

GetStartAddress returns the StartAddress field if non-nil, zero value otherwise.

### GetStartAddressOk

`func (o *FunctionBoundary) GetStartAddressOk() (*int64, bool)`

GetStartAddressOk returns a tuple with the StartAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAddress

`func (o *FunctionBoundary) SetStartAddress(v int64)`

SetStartAddress sets StartAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


