# SecurityFinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckId** | Pointer to **string** | Semgrep rule ID that matched | [optional] 
**Confidence** | Pointer to **string** | Semgrep&#39;s confidence in the finding | [optional] 
**Cwe** | Pointer to **[]string** | CWE identifiers associated with the finding | [optional] 
**EndLine** | Pointer to **int64** | Line the finding ends on | [optional] 
**FunctionId** | Pointer to **int64** | ID of the function the finding was reported in | [optional] 
**FunctionName** | Pointer to **string** | Name of the function the finding was reported in | [optional] 
**Impact** | Pointer to **string** | Estimated impact of the finding | [optional] 
**Message** | Pointer to **string** | Human-readable description of the finding | [optional] 
**Severity** | Pointer to **string** | Severity of the finding | [optional] 
**SnippetLines** | Pointer to **[]string** | Source lines making up the reported snippet | [optional] 
**SnippetStartLine** | Pointer to **int64** | Line the reported snippet starts on | [optional] 
**StartLine** | Pointer to **int64** | Line the finding starts on | [optional] 

## Methods

### NewSecurityFinding

`func NewSecurityFinding() *SecurityFinding`

NewSecurityFinding instantiates a new SecurityFinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityFindingWithDefaults

`func NewSecurityFindingWithDefaults() *SecurityFinding`

NewSecurityFindingWithDefaults instantiates a new SecurityFinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckId

`func (o *SecurityFinding) GetCheckId() string`

GetCheckId returns the CheckId field if non-nil, zero value otherwise.

### GetCheckIdOk

`func (o *SecurityFinding) GetCheckIdOk() (*string, bool)`

GetCheckIdOk returns a tuple with the CheckId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckId

`func (o *SecurityFinding) SetCheckId(v string)`

SetCheckId sets CheckId field to given value.

### HasCheckId

`func (o *SecurityFinding) HasCheckId() bool`

HasCheckId returns a boolean if a field has been set.

### GetConfidence

`func (o *SecurityFinding) GetConfidence() string`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *SecurityFinding) GetConfidenceOk() (*string, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *SecurityFinding) SetConfidence(v string)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *SecurityFinding) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetCwe

`func (o *SecurityFinding) GetCwe() []string`

GetCwe returns the Cwe field if non-nil, zero value otherwise.

### GetCweOk

`func (o *SecurityFinding) GetCweOk() (*[]string, bool)`

GetCweOk returns a tuple with the Cwe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCwe

`func (o *SecurityFinding) SetCwe(v []string)`

SetCwe sets Cwe field to given value.

### HasCwe

`func (o *SecurityFinding) HasCwe() bool`

HasCwe returns a boolean if a field has been set.

### SetCweNil

`func (o *SecurityFinding) SetCweNil(b bool)`

 SetCweNil sets the value for Cwe to be an explicit nil

### UnsetCwe
`func (o *SecurityFinding) UnsetCwe()`

UnsetCwe ensures that no value is present for Cwe, not even an explicit nil
### GetEndLine

`func (o *SecurityFinding) GetEndLine() int64`

GetEndLine returns the EndLine field if non-nil, zero value otherwise.

### GetEndLineOk

`func (o *SecurityFinding) GetEndLineOk() (*int64, bool)`

GetEndLineOk returns a tuple with the EndLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndLine

`func (o *SecurityFinding) SetEndLine(v int64)`

SetEndLine sets EndLine field to given value.

### HasEndLine

`func (o *SecurityFinding) HasEndLine() bool`

HasEndLine returns a boolean if a field has been set.

### GetFunctionId

`func (o *SecurityFinding) GetFunctionId() int64`

GetFunctionId returns the FunctionId field if non-nil, zero value otherwise.

### GetFunctionIdOk

`func (o *SecurityFinding) GetFunctionIdOk() (*int64, bool)`

GetFunctionIdOk returns a tuple with the FunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionId

`func (o *SecurityFinding) SetFunctionId(v int64)`

SetFunctionId sets FunctionId field to given value.

### HasFunctionId

`func (o *SecurityFinding) HasFunctionId() bool`

HasFunctionId returns a boolean if a field has been set.

### GetFunctionName

`func (o *SecurityFinding) GetFunctionName() string`

GetFunctionName returns the FunctionName field if non-nil, zero value otherwise.

### GetFunctionNameOk

`func (o *SecurityFinding) GetFunctionNameOk() (*string, bool)`

GetFunctionNameOk returns a tuple with the FunctionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionName

`func (o *SecurityFinding) SetFunctionName(v string)`

SetFunctionName sets FunctionName field to given value.

### HasFunctionName

`func (o *SecurityFinding) HasFunctionName() bool`

HasFunctionName returns a boolean if a field has been set.

### GetImpact

`func (o *SecurityFinding) GetImpact() string`

GetImpact returns the Impact field if non-nil, zero value otherwise.

### GetImpactOk

`func (o *SecurityFinding) GetImpactOk() (*string, bool)`

GetImpactOk returns a tuple with the Impact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpact

`func (o *SecurityFinding) SetImpact(v string)`

SetImpact sets Impact field to given value.

### HasImpact

`func (o *SecurityFinding) HasImpact() bool`

HasImpact returns a boolean if a field has been set.

### GetMessage

`func (o *SecurityFinding) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SecurityFinding) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SecurityFinding) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SecurityFinding) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetSeverity

`func (o *SecurityFinding) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *SecurityFinding) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *SecurityFinding) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *SecurityFinding) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetSnippetLines

`func (o *SecurityFinding) GetSnippetLines() []string`

GetSnippetLines returns the SnippetLines field if non-nil, zero value otherwise.

### GetSnippetLinesOk

`func (o *SecurityFinding) GetSnippetLinesOk() (*[]string, bool)`

GetSnippetLinesOk returns a tuple with the SnippetLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnippetLines

`func (o *SecurityFinding) SetSnippetLines(v []string)`

SetSnippetLines sets SnippetLines field to given value.

### HasSnippetLines

`func (o *SecurityFinding) HasSnippetLines() bool`

HasSnippetLines returns a boolean if a field has been set.

### SetSnippetLinesNil

`func (o *SecurityFinding) SetSnippetLinesNil(b bool)`

 SetSnippetLinesNil sets the value for SnippetLines to be an explicit nil

### UnsetSnippetLines
`func (o *SecurityFinding) UnsetSnippetLines()`

UnsetSnippetLines ensures that no value is present for SnippetLines, not even an explicit nil
### GetSnippetStartLine

`func (o *SecurityFinding) GetSnippetStartLine() int64`

GetSnippetStartLine returns the SnippetStartLine field if non-nil, zero value otherwise.

### GetSnippetStartLineOk

`func (o *SecurityFinding) GetSnippetStartLineOk() (*int64, bool)`

GetSnippetStartLineOk returns a tuple with the SnippetStartLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnippetStartLine

`func (o *SecurityFinding) SetSnippetStartLine(v int64)`

SetSnippetStartLine sets SnippetStartLine field to given value.

### HasSnippetStartLine

`func (o *SecurityFinding) HasSnippetStartLine() bool`

HasSnippetStartLine returns a boolean if a field has been set.

### GetStartLine

`func (o *SecurityFinding) GetStartLine() int64`

GetStartLine returns the StartLine field if non-nil, zero value otherwise.

### GetStartLineOk

`func (o *SecurityFinding) GetStartLineOk() (*int64, bool)`

GetStartLineOk returns a tuple with the StartLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartLine

`func (o *SecurityFinding) SetStartLine(v int64)`

SetStartLine sets StartLine field to given value.

### HasStartLine

`func (o *SecurityFinding) HasStartLine() bool`

HasStartLine returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


