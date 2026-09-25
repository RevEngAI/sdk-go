# CreateURLRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnalysisScope** | Pointer to **string** |  | [optional] [default to "PRIVATE"]
**Url** | **string** |  | 

## Methods

### NewCreateURLRequest

`func NewCreateURLRequest(url string, ) *CreateURLRequest`

NewCreateURLRequest instantiates a new CreateURLRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateURLRequestWithDefaults

`func NewCreateURLRequestWithDefaults() *CreateURLRequest`

NewCreateURLRequestWithDefaults instantiates a new CreateURLRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalysisScope

`func (o *CreateURLRequest) GetAnalysisScope() string`

GetAnalysisScope returns the AnalysisScope field if non-nil, zero value otherwise.

### GetAnalysisScopeOk

`func (o *CreateURLRequest) GetAnalysisScopeOk() (*string, bool)`

GetAnalysisScopeOk returns a tuple with the AnalysisScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisScope

`func (o *CreateURLRequest) SetAnalysisScope(v string)`

SetAnalysisScope sets AnalysisScope field to given value.

### HasAnalysisScope

`func (o *CreateURLRequest) HasAnalysisScope() bool`

HasAnalysisScope returns a boolean if a field has been set.

### GetUrl

`func (o *CreateURLRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreateURLRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreateURLRequest) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


