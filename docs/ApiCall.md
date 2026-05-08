# ApiCall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CalledFrom** | Pointer to **string** |  | [optional] 
**CalledFromRva** | Pointer to **string** |  | [optional] 
**FromModule** | Pointer to **string** |  | [optional] 
**Method** | **string** |  | 

## Methods

### NewApiCall

`func NewApiCall(method string, ) *ApiCall`

NewApiCall instantiates a new ApiCall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiCallWithDefaults

`func NewApiCallWithDefaults() *ApiCall`

NewApiCallWithDefaults instantiates a new ApiCall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCalledFrom

`func (o *ApiCall) GetCalledFrom() string`

GetCalledFrom returns the CalledFrom field if non-nil, zero value otherwise.

### GetCalledFromOk

`func (o *ApiCall) GetCalledFromOk() (*string, bool)`

GetCalledFromOk returns a tuple with the CalledFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalledFrom

`func (o *ApiCall) SetCalledFrom(v string)`

SetCalledFrom sets CalledFrom field to given value.

### HasCalledFrom

`func (o *ApiCall) HasCalledFrom() bool`

HasCalledFrom returns a boolean if a field has been set.

### GetCalledFromRva

`func (o *ApiCall) GetCalledFromRva() string`

GetCalledFromRva returns the CalledFromRva field if non-nil, zero value otherwise.

### GetCalledFromRvaOk

`func (o *ApiCall) GetCalledFromRvaOk() (*string, bool)`

GetCalledFromRvaOk returns a tuple with the CalledFromRva field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalledFromRva

`func (o *ApiCall) SetCalledFromRva(v string)`

SetCalledFromRva sets CalledFromRva field to given value.

### HasCalledFromRva

`func (o *ApiCall) HasCalledFromRva() bool`

HasCalledFromRva returns a boolean if a field has been set.

### GetFromModule

`func (o *ApiCall) GetFromModule() string`

GetFromModule returns the FromModule field if non-nil, zero value otherwise.

### GetFromModuleOk

`func (o *ApiCall) GetFromModuleOk() (*string, bool)`

GetFromModuleOk returns a tuple with the FromModule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromModule

`func (o *ApiCall) SetFromModule(v string)`

SetFromModule sets FromModule field to given value.

### HasFromModule

`func (o *ApiCall) HasFromModule() bool`

HasFromModule returns a boolean if a field has been set.

### GetMethod

`func (o *ApiCall) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *ApiCall) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *ApiCall) SetMethod(v string)`

SetMethod sets Method field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


