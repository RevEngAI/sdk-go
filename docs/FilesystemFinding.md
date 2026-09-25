# FilesystemFinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | Function&#39;s virtual address, hex-encoded | 
**Categories** | **[]string** | Distinct filesystem categories evidenced by this function | 
**Confidence** | **string** | High when a direct name match was found, medium when the function only calls into filesystem APIs | 
**DirectMatches** | Pointer to [**[]FilesystemDirectMatch**](FilesystemDirectMatch.md) | Matches against the function&#39;s own name | [optional] 
**EvidenceCount** | **int64** | Total number of direct matches and filesystem calls | 
**FilesystemCalls** | Pointer to [**[]FilesystemCall**](FilesystemCall.md) | Matches against names this function calls | [optional] 
**FunctionId** | **int64** | ID of the function the finding was reported in | 
**FunctionName** | **string** | Name of the function the finding was reported in | 
**FunctionSize** | **int64** | Size of the function in bytes | 
**Modifies** | **bool** | Whether this function evidences modifying the filesystem rather than only observing it | 
**Sources** | **[]string** | Distinct filesystem sources evidenced by this function | 
**Verification** | Pointer to [**FilesystemVerification**](FilesystemVerification.md) | LLM verdict checking this finding against its decompilation. Present only when the run verified this finding. | [optional] 

## Methods

### NewFilesystemFinding

`func NewFilesystemFinding(address string, categories []string, confidence string, evidenceCount int64, functionId int64, functionName string, functionSize int64, modifies bool, sources []string, ) *FilesystemFinding`

NewFilesystemFinding instantiates a new FilesystemFinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilesystemFindingWithDefaults

`func NewFilesystemFindingWithDefaults() *FilesystemFinding`

NewFilesystemFindingWithDefaults instantiates a new FilesystemFinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *FilesystemFinding) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *FilesystemFinding) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *FilesystemFinding) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetCategories

`func (o *FilesystemFinding) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *FilesystemFinding) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *FilesystemFinding) SetCategories(v []string)`

SetCategories sets Categories field to given value.


### SetCategoriesNil

`func (o *FilesystemFinding) SetCategoriesNil(b bool)`

 SetCategoriesNil sets the value for Categories to be an explicit nil

### UnsetCategories
`func (o *FilesystemFinding) UnsetCategories()`

UnsetCategories ensures that no value is present for Categories, not even an explicit nil
### GetConfidence

`func (o *FilesystemFinding) GetConfidence() string`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *FilesystemFinding) GetConfidenceOk() (*string, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *FilesystemFinding) SetConfidence(v string)`

SetConfidence sets Confidence field to given value.


### GetDirectMatches

`func (o *FilesystemFinding) GetDirectMatches() []FilesystemDirectMatch`

GetDirectMatches returns the DirectMatches field if non-nil, zero value otherwise.

### GetDirectMatchesOk

`func (o *FilesystemFinding) GetDirectMatchesOk() (*[]FilesystemDirectMatch, bool)`

GetDirectMatchesOk returns a tuple with the DirectMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectMatches

`func (o *FilesystemFinding) SetDirectMatches(v []FilesystemDirectMatch)`

SetDirectMatches sets DirectMatches field to given value.

### HasDirectMatches

`func (o *FilesystemFinding) HasDirectMatches() bool`

HasDirectMatches returns a boolean if a field has been set.

### SetDirectMatchesNil

`func (o *FilesystemFinding) SetDirectMatchesNil(b bool)`

 SetDirectMatchesNil sets the value for DirectMatches to be an explicit nil

### UnsetDirectMatches
`func (o *FilesystemFinding) UnsetDirectMatches()`

UnsetDirectMatches ensures that no value is present for DirectMatches, not even an explicit nil
### GetEvidenceCount

`func (o *FilesystemFinding) GetEvidenceCount() int64`

GetEvidenceCount returns the EvidenceCount field if non-nil, zero value otherwise.

### GetEvidenceCountOk

`func (o *FilesystemFinding) GetEvidenceCountOk() (*int64, bool)`

GetEvidenceCountOk returns a tuple with the EvidenceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceCount

`func (o *FilesystemFinding) SetEvidenceCount(v int64)`

SetEvidenceCount sets EvidenceCount field to given value.


### GetFilesystemCalls

`func (o *FilesystemFinding) GetFilesystemCalls() []FilesystemCall`

GetFilesystemCalls returns the FilesystemCalls field if non-nil, zero value otherwise.

### GetFilesystemCallsOk

`func (o *FilesystemFinding) GetFilesystemCallsOk() (*[]FilesystemCall, bool)`

GetFilesystemCallsOk returns a tuple with the FilesystemCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesystemCalls

`func (o *FilesystemFinding) SetFilesystemCalls(v []FilesystemCall)`

SetFilesystemCalls sets FilesystemCalls field to given value.

### HasFilesystemCalls

`func (o *FilesystemFinding) HasFilesystemCalls() bool`

HasFilesystemCalls returns a boolean if a field has been set.

### SetFilesystemCallsNil

`func (o *FilesystemFinding) SetFilesystemCallsNil(b bool)`

 SetFilesystemCallsNil sets the value for FilesystemCalls to be an explicit nil

### UnsetFilesystemCalls
`func (o *FilesystemFinding) UnsetFilesystemCalls()`

UnsetFilesystemCalls ensures that no value is present for FilesystemCalls, not even an explicit nil
### GetFunctionId

`func (o *FilesystemFinding) GetFunctionId() int64`

GetFunctionId returns the FunctionId field if non-nil, zero value otherwise.

### GetFunctionIdOk

`func (o *FilesystemFinding) GetFunctionIdOk() (*int64, bool)`

GetFunctionIdOk returns a tuple with the FunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionId

`func (o *FilesystemFinding) SetFunctionId(v int64)`

SetFunctionId sets FunctionId field to given value.


### GetFunctionName

`func (o *FilesystemFinding) GetFunctionName() string`

GetFunctionName returns the FunctionName field if non-nil, zero value otherwise.

### GetFunctionNameOk

`func (o *FilesystemFinding) GetFunctionNameOk() (*string, bool)`

GetFunctionNameOk returns a tuple with the FunctionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionName

`func (o *FilesystemFinding) SetFunctionName(v string)`

SetFunctionName sets FunctionName field to given value.


### GetFunctionSize

`func (o *FilesystemFinding) GetFunctionSize() int64`

GetFunctionSize returns the FunctionSize field if non-nil, zero value otherwise.

### GetFunctionSizeOk

`func (o *FilesystemFinding) GetFunctionSizeOk() (*int64, bool)`

GetFunctionSizeOk returns a tuple with the FunctionSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionSize

`func (o *FilesystemFinding) SetFunctionSize(v int64)`

SetFunctionSize sets FunctionSize field to given value.


### GetModifies

`func (o *FilesystemFinding) GetModifies() bool`

GetModifies returns the Modifies field if non-nil, zero value otherwise.

### GetModifiesOk

`func (o *FilesystemFinding) GetModifiesOk() (*bool, bool)`

GetModifiesOk returns a tuple with the Modifies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifies

`func (o *FilesystemFinding) SetModifies(v bool)`

SetModifies sets Modifies field to given value.


### GetSources

`func (o *FilesystemFinding) GetSources() []string`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *FilesystemFinding) GetSourcesOk() (*[]string, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *FilesystemFinding) SetSources(v []string)`

SetSources sets Sources field to given value.


### SetSourcesNil

`func (o *FilesystemFinding) SetSourcesNil(b bool)`

 SetSourcesNil sets the value for Sources to be an explicit nil

### UnsetSources
`func (o *FilesystemFinding) UnsetSources()`

UnsetSources ensures that no value is present for Sources, not even an explicit nil
### GetVerification

`func (o *FilesystemFinding) GetVerification() FilesystemVerification`

GetVerification returns the Verification field if non-nil, zero value otherwise.

### GetVerificationOk

`func (o *FilesystemFinding) GetVerificationOk() (*FilesystemVerification, bool)`

GetVerificationOk returns a tuple with the Verification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerification

`func (o *FilesystemFinding) SetVerification(v FilesystemVerification)`

SetVerification sets Verification field to given value.

### HasVerification

`func (o *FilesystemFinding) HasVerification() bool`

HasVerification returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


