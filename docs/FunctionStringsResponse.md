# FunctionStringsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Strings** | [**[]FunctionString**](FunctionString.md) | The strings associated with this function | 
**TotalStrings** | **int32** | The total number of strings associated with this function | 

## Methods

### NewFunctionStringsResponse

`func NewFunctionStringsResponse(strings []FunctionString, totalStrings int32, ) *FunctionStringsResponse`

NewFunctionStringsResponse instantiates a new FunctionStringsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionStringsResponseWithDefaults

`func NewFunctionStringsResponseWithDefaults() *FunctionStringsResponse`

NewFunctionStringsResponseWithDefaults instantiates a new FunctionStringsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStrings

`func (o *FunctionStringsResponse) GetStrings() []FunctionString`

GetStrings returns the Strings field if non-nil, zero value otherwise.

### GetStringsOk

`func (o *FunctionStringsResponse) GetStringsOk() (*[]FunctionString, bool)`

GetStringsOk returns a tuple with the Strings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrings

`func (o *FunctionStringsResponse) SetStrings(v []FunctionString)`

SetStrings sets Strings field to given value.


### GetTotalStrings

`func (o *FunctionStringsResponse) GetTotalStrings() int32`

GetTotalStrings returns the TotalStrings field if non-nil, zero value otherwise.

### GetTotalStringsOk

`func (o *FunctionStringsResponse) GetTotalStringsOk() (*int32, bool)`

GetTotalStringsOk returns a tuple with the TotalStrings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalStrings

`func (o *FunctionStringsResponse) SetTotalStrings(v int32)`

SetTotalStrings sets TotalStrings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


