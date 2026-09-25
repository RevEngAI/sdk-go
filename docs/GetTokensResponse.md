# GetTokensResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AiDecomp** | **string** | Tokenised AI-decompilation. Includes generated comments. Empty until a run has succeeded. | 
**AnalysisId** | **int64** | Analysis the function belongs to. Scopes every data_type_id below. | 
**PlaceholderToRenderedToken** | [**map[string]RenderedToken**](RenderedToken.md) | Each placeholder token mapped to the value the server would render in its place, and the record it refers to. Null until a run has succeeded. | 
**PlaceholderToUserOverride** | [**map[string]Token**](Token.md) | The overrides on this function, keyed by token, each carrying who chose it. Null until a run has succeeded. | 

## Methods

### NewGetTokensResponse

`func NewGetTokensResponse(aiDecomp string, analysisId int64, placeholderToRenderedToken map[string]RenderedToken, placeholderToUserOverride map[string]Token, ) *GetTokensResponse`

NewGetTokensResponse instantiates a new GetTokensResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTokensResponseWithDefaults

`func NewGetTokensResponseWithDefaults() *GetTokensResponse`

NewGetTokensResponseWithDefaults instantiates a new GetTokensResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAiDecomp

`func (o *GetTokensResponse) GetAiDecomp() string`

GetAiDecomp returns the AiDecomp field if non-nil, zero value otherwise.

### GetAiDecompOk

`func (o *GetTokensResponse) GetAiDecompOk() (*string, bool)`

GetAiDecompOk returns a tuple with the AiDecomp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiDecomp

`func (o *GetTokensResponse) SetAiDecomp(v string)`

SetAiDecomp sets AiDecomp field to given value.


### GetAnalysisId

`func (o *GetTokensResponse) GetAnalysisId() int64`

GetAnalysisId returns the AnalysisId field if non-nil, zero value otherwise.

### GetAnalysisIdOk

`func (o *GetTokensResponse) GetAnalysisIdOk() (*int64, bool)`

GetAnalysisIdOk returns a tuple with the AnalysisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisId

`func (o *GetTokensResponse) SetAnalysisId(v int64)`

SetAnalysisId sets AnalysisId field to given value.


### GetPlaceholderToRenderedToken

`func (o *GetTokensResponse) GetPlaceholderToRenderedToken() map[string]RenderedToken`

GetPlaceholderToRenderedToken returns the PlaceholderToRenderedToken field if non-nil, zero value otherwise.

### GetPlaceholderToRenderedTokenOk

`func (o *GetTokensResponse) GetPlaceholderToRenderedTokenOk() (*map[string]RenderedToken, bool)`

GetPlaceholderToRenderedTokenOk returns a tuple with the PlaceholderToRenderedToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholderToRenderedToken

`func (o *GetTokensResponse) SetPlaceholderToRenderedToken(v map[string]RenderedToken)`

SetPlaceholderToRenderedToken sets PlaceholderToRenderedToken field to given value.


### GetPlaceholderToUserOverride

`func (o *GetTokensResponse) GetPlaceholderToUserOverride() map[string]Token`

GetPlaceholderToUserOverride returns the PlaceholderToUserOverride field if non-nil, zero value otherwise.

### GetPlaceholderToUserOverrideOk

`func (o *GetTokensResponse) GetPlaceholderToUserOverrideOk() (*map[string]Token, bool)`

GetPlaceholderToUserOverrideOk returns a tuple with the PlaceholderToUserOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholderToUserOverride

`func (o *GetTokensResponse) SetPlaceholderToUserOverride(v map[string]Token)`

SetPlaceholderToUserOverride sets PlaceholderToUserOverride field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


