# CalleesCallerFunctionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseAddress** | **int32** | Base address of the binary | 
**Callees** | [**[]CalleeFunctionInfo**](CalleeFunctionInfo.md) | List of functions called by the target function | 
**Callers** | [**[]CallerFunctionInfo**](CallerFunctionInfo.md) | List of functions that call the target function | 

## Methods

### NewCalleesCallerFunctionsResponse

`func NewCalleesCallerFunctionsResponse(baseAddress int32, callees []CalleeFunctionInfo, callers []CallerFunctionInfo, ) *CalleesCallerFunctionsResponse`

NewCalleesCallerFunctionsResponse instantiates a new CalleesCallerFunctionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalleesCallerFunctionsResponseWithDefaults

`func NewCalleesCallerFunctionsResponseWithDefaults() *CalleesCallerFunctionsResponse`

NewCalleesCallerFunctionsResponseWithDefaults instantiates a new CalleesCallerFunctionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseAddress

`func (o *CalleesCallerFunctionsResponse) GetBaseAddress() int32`

GetBaseAddress returns the BaseAddress field if non-nil, zero value otherwise.

### GetBaseAddressOk

`func (o *CalleesCallerFunctionsResponse) GetBaseAddressOk() (*int32, bool)`

GetBaseAddressOk returns a tuple with the BaseAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAddress

`func (o *CalleesCallerFunctionsResponse) SetBaseAddress(v int32)`

SetBaseAddress sets BaseAddress field to given value.


### GetCallees

`func (o *CalleesCallerFunctionsResponse) GetCallees() []CalleeFunctionInfo`

GetCallees returns the Callees field if non-nil, zero value otherwise.

### GetCalleesOk

`func (o *CalleesCallerFunctionsResponse) GetCalleesOk() (*[]CalleeFunctionInfo, bool)`

GetCalleesOk returns a tuple with the Callees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallees

`func (o *CalleesCallerFunctionsResponse) SetCallees(v []CalleeFunctionInfo)`

SetCallees sets Callees field to given value.


### GetCallers

`func (o *CalleesCallerFunctionsResponse) GetCallers() []CallerFunctionInfo`

GetCallers returns the Callers field if non-nil, zero value otherwise.

### GetCallersOk

`func (o *CalleesCallerFunctionsResponse) GetCallersOk() (*[]CallerFunctionInfo, bool)`

GetCallersOk returns a tuple with the Callers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallers

`func (o *CalleesCallerFunctionsResponse) SetCallers(v []CallerFunctionInfo)`

SetCallers sets Callers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


