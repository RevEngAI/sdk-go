# ListFunctionStringsOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Strings** | [**[]FunctionStringItem**](FunctionStringItem.md) |  | 
**TotalStrings** | **int64** |  | 

## Methods

### NewListFunctionStringsOutputBody

`func NewListFunctionStringsOutputBody(strings []FunctionStringItem, totalStrings int64, ) *ListFunctionStringsOutputBody`

NewListFunctionStringsOutputBody instantiates a new ListFunctionStringsOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListFunctionStringsOutputBodyWithDefaults

`func NewListFunctionStringsOutputBodyWithDefaults() *ListFunctionStringsOutputBody`

NewListFunctionStringsOutputBodyWithDefaults instantiates a new ListFunctionStringsOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStrings

`func (o *ListFunctionStringsOutputBody) GetStrings() []FunctionStringItem`

GetStrings returns the Strings field if non-nil, zero value otherwise.

### GetStringsOk

`func (o *ListFunctionStringsOutputBody) GetStringsOk() (*[]FunctionStringItem, bool)`

GetStringsOk returns a tuple with the Strings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrings

`func (o *ListFunctionStringsOutputBody) SetStrings(v []FunctionStringItem)`

SetStrings sets Strings field to given value.


### SetStringsNil

`func (o *ListFunctionStringsOutputBody) SetStringsNil(b bool)`

 SetStringsNil sets the value for Strings to be an explicit nil

### UnsetStrings
`func (o *ListFunctionStringsOutputBody) UnsetStrings()`

UnsetStrings ensures that no value is present for Strings, not even an explicit nil
### GetTotalStrings

`func (o *ListFunctionStringsOutputBody) GetTotalStrings() int64`

GetTotalStrings returns the TotalStrings field if non-nil, zero value otherwise.

### GetTotalStringsOk

`func (o *ListFunctionStringsOutputBody) GetTotalStringsOk() (*int64, bool)`

GetTotalStringsOk returns a tuple with the TotalStrings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalStrings

`func (o *ListFunctionStringsOutputBody) SetTotalStrings(v int64)`

SetTotalStrings sets TotalStrings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


