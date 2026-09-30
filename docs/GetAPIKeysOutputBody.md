# GetAPIKeysOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keys** | [**[]ApiKeyBody**](ApiKeyBody.md) | The caller&#39;s active API keys. Currently always exactly one, created on first request if the caller doesn&#39;t have one yet | 

## Methods

### NewGetAPIKeysOutputBody

`func NewGetAPIKeysOutputBody(keys []ApiKeyBody, ) *GetAPIKeysOutputBody`

NewGetAPIKeysOutputBody instantiates a new GetAPIKeysOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAPIKeysOutputBodyWithDefaults

`func NewGetAPIKeysOutputBodyWithDefaults() *GetAPIKeysOutputBody`

NewGetAPIKeysOutputBodyWithDefaults instantiates a new GetAPIKeysOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeys

`func (o *GetAPIKeysOutputBody) GetKeys() []ApiKeyBody`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *GetAPIKeysOutputBody) GetKeysOk() (*[]ApiKeyBody, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *GetAPIKeysOutputBody) SetKeys(v []ApiKeyBody)`

SetKeys sets Keys field to given value.


### SetKeysNil

`func (o *GetAPIKeysOutputBody) SetKeysNil(b bool)`

 SetKeysNil sets the value for Keys to be an explicit nil

### UnsetKeys
`func (o *GetAPIKeysOutputBody) UnsetKeys()`

UnsetKeys ensures that no value is present for Keys, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


