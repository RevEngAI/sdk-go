# OperationPrerequisitesBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operation** | **string** | The operation | 
**Prerequisites** | **[]string** | Operations that must complete before this operation can start | 

## Methods

### NewOperationPrerequisitesBody

`func NewOperationPrerequisitesBody(operation string, prerequisites []string, ) *OperationPrerequisitesBody`

NewOperationPrerequisitesBody instantiates a new OperationPrerequisitesBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationPrerequisitesBodyWithDefaults

`func NewOperationPrerequisitesBodyWithDefaults() *OperationPrerequisitesBody`

NewOperationPrerequisitesBodyWithDefaults instantiates a new OperationPrerequisitesBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperation

`func (o *OperationPrerequisitesBody) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *OperationPrerequisitesBody) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *OperationPrerequisitesBody) SetOperation(v string)`

SetOperation sets Operation field to given value.


### GetPrerequisites

`func (o *OperationPrerequisitesBody) GetPrerequisites() []string`

GetPrerequisites returns the Prerequisites field if non-nil, zero value otherwise.

### GetPrerequisitesOk

`func (o *OperationPrerequisitesBody) GetPrerequisitesOk() (*[]string, bool)`

GetPrerequisitesOk returns a tuple with the Prerequisites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrerequisites

`func (o *OperationPrerequisitesBody) SetPrerequisites(v []string)`

SetPrerequisites sets Prerequisites field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


