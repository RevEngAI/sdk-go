# UpdateAnalysisInputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnalysisScope** | Pointer to **string** | The analysis&#39; visibility. Changing to a non-PUBLIC scope requires a subscription tier that supports private analyses | [optional] 
**BinaryName** | Pointer to **string** | Renames the analysis&#39; binary. Empty or whitespace-only is rejected | [optional] 

## Methods

### NewUpdateAnalysisInputBody

`func NewUpdateAnalysisInputBody() *UpdateAnalysisInputBody`

NewUpdateAnalysisInputBody instantiates a new UpdateAnalysisInputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAnalysisInputBodyWithDefaults

`func NewUpdateAnalysisInputBodyWithDefaults() *UpdateAnalysisInputBody`

NewUpdateAnalysisInputBodyWithDefaults instantiates a new UpdateAnalysisInputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalysisScope

`func (o *UpdateAnalysisInputBody) GetAnalysisScope() string`

GetAnalysisScope returns the AnalysisScope field if non-nil, zero value otherwise.

### GetAnalysisScopeOk

`func (o *UpdateAnalysisInputBody) GetAnalysisScopeOk() (*string, bool)`

GetAnalysisScopeOk returns a tuple with the AnalysisScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisScope

`func (o *UpdateAnalysisInputBody) SetAnalysisScope(v string)`

SetAnalysisScope sets AnalysisScope field to given value.

### HasAnalysisScope

`func (o *UpdateAnalysisInputBody) HasAnalysisScope() bool`

HasAnalysisScope returns a boolean if a field has been set.

### GetBinaryName

`func (o *UpdateAnalysisInputBody) GetBinaryName() string`

GetBinaryName returns the BinaryName field if non-nil, zero value otherwise.

### GetBinaryNameOk

`func (o *UpdateAnalysisInputBody) GetBinaryNameOk() (*string, bool)`

GetBinaryNameOk returns a tuple with the BinaryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryName

`func (o *UpdateAnalysisInputBody) SetBinaryName(v string)`

SetBinaryName sets BinaryName field to given value.

### HasBinaryName

`func (o *UpdateAnalysisInputBody) HasBinaryName() bool`

HasBinaryName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


