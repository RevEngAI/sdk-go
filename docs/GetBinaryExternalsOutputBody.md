# GetBinaryExternalsOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BinaryId** | **int64** |  | 
**Externals** | [**BinaryExternalsBody**](BinaryExternalsBody.md) | Null until at least one external lookup has run for this binary&#39;s content hash | 

## Methods

### NewGetBinaryExternalsOutputBody

`func NewGetBinaryExternalsOutputBody(binaryId int64, externals BinaryExternalsBody, ) *GetBinaryExternalsOutputBody`

NewGetBinaryExternalsOutputBody instantiates a new GetBinaryExternalsOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBinaryExternalsOutputBodyWithDefaults

`func NewGetBinaryExternalsOutputBodyWithDefaults() *GetBinaryExternalsOutputBody`

NewGetBinaryExternalsOutputBodyWithDefaults instantiates a new GetBinaryExternalsOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBinaryId

`func (o *GetBinaryExternalsOutputBody) GetBinaryId() int64`

GetBinaryId returns the BinaryId field if non-nil, zero value otherwise.

### GetBinaryIdOk

`func (o *GetBinaryExternalsOutputBody) GetBinaryIdOk() (*int64, bool)`

GetBinaryIdOk returns a tuple with the BinaryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryId

`func (o *GetBinaryExternalsOutputBody) SetBinaryId(v int64)`

SetBinaryId sets BinaryId field to given value.


### GetExternals

`func (o *GetBinaryExternalsOutputBody) GetExternals() BinaryExternalsBody`

GetExternals returns the Externals field if non-nil, zero value otherwise.

### GetExternalsOk

`func (o *GetBinaryExternalsOutputBody) GetExternalsOk() (*BinaryExternalsBody, bool)`

GetExternalsOk returns a tuple with the Externals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternals

`func (o *GetBinaryExternalsOutputBody) SetExternals(v BinaryExternalsBody)`

SetExternals sets Externals field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


