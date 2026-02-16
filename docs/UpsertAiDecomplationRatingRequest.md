# UpsertAiDecomplationRatingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rating** | [**AiDecompilationRating**](AiDecompilationRating.md) | The rating for the AI decompilation response | 
**Reason** | **NullableString** |  | 

## Methods

### NewUpsertAiDecomplationRatingRequest

`func NewUpsertAiDecomplationRatingRequest(rating AiDecompilationRating, reason NullableString, ) *UpsertAiDecomplationRatingRequest`

NewUpsertAiDecomplationRatingRequest instantiates a new UpsertAiDecomplationRatingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpsertAiDecomplationRatingRequestWithDefaults

`func NewUpsertAiDecomplationRatingRequestWithDefaults() *UpsertAiDecomplationRatingRequest`

NewUpsertAiDecomplationRatingRequestWithDefaults instantiates a new UpsertAiDecomplationRatingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRating

`func (o *UpsertAiDecomplationRatingRequest) GetRating() AiDecompilationRating`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *UpsertAiDecomplationRatingRequest) GetRatingOk() (*AiDecompilationRating, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *UpsertAiDecomplationRatingRequest) SetRating(v AiDecompilationRating)`

SetRating sets Rating field to given value.


### GetReason

`func (o *UpsertAiDecomplationRatingRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *UpsertAiDecomplationRatingRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *UpsertAiDecomplationRatingRequest) SetReason(v string)`

SetReason sets Reason field to given value.


### SetReasonNil

`func (o *UpsertAiDecomplationRatingRequest) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *UpsertAiDecomplationRatingRequest) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


