# GetAiDecompilationRatingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rating** | [**AiDecompilationRating**](AiDecompilationRating.md) | The rating the user has given to the AI decompilation response | 
**Reason** | **NullableString** |  | 

## Methods

### NewGetAiDecompilationRatingResponse

`func NewGetAiDecompilationRatingResponse(rating AiDecompilationRating, reason NullableString, ) *GetAiDecompilationRatingResponse`

NewGetAiDecompilationRatingResponse instantiates a new GetAiDecompilationRatingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAiDecompilationRatingResponseWithDefaults

`func NewGetAiDecompilationRatingResponseWithDefaults() *GetAiDecompilationRatingResponse`

NewGetAiDecompilationRatingResponseWithDefaults instantiates a new GetAiDecompilationRatingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRating

`func (o *GetAiDecompilationRatingResponse) GetRating() AiDecompilationRating`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *GetAiDecompilationRatingResponse) GetRatingOk() (*AiDecompilationRating, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *GetAiDecompilationRatingResponse) SetRating(v AiDecompilationRating)`

SetRating sets Rating field to given value.


### GetReason

`func (o *GetAiDecompilationRatingResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *GetAiDecompilationRatingResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *GetAiDecompilationRatingResponse) SetReason(v string)`

SetReason sets Reason field to given value.


### SetReasonNil

`func (o *GetAiDecompilationRatingResponse) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *GetAiDecompilationRatingResponse) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


