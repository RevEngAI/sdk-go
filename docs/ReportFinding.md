# ReportFinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Evidence** | [**[]ReportFindingEvidenceInner**](ReportFindingEvidenceInner.md) | The evidence for or against the finding. | 
**PropertyId** | **string** | Identifier of the property the finding asserts. | 
**Reachability** | **string** | Whether the code behind the finding is reachable from an entry point. | 

## Methods

### NewReportFinding

`func NewReportFinding(evidence []ReportFindingEvidenceInner, propertyId string, reachability string, ) *ReportFinding`

NewReportFinding instantiates a new ReportFinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportFindingWithDefaults

`func NewReportFindingWithDefaults() *ReportFinding`

NewReportFindingWithDefaults instantiates a new ReportFinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvidence

`func (o *ReportFinding) GetEvidence() []ReportFindingEvidenceInner`

GetEvidence returns the Evidence field if non-nil, zero value otherwise.

### GetEvidenceOk

`func (o *ReportFinding) GetEvidenceOk() (*[]ReportFindingEvidenceInner, bool)`

GetEvidenceOk returns a tuple with the Evidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidence

`func (o *ReportFinding) SetEvidence(v []ReportFindingEvidenceInner)`

SetEvidence sets Evidence field to given value.


### GetPropertyId

`func (o *ReportFinding) GetPropertyId() string`

GetPropertyId returns the PropertyId field if non-nil, zero value otherwise.

### GetPropertyIdOk

`func (o *ReportFinding) GetPropertyIdOk() (*string, bool)`

GetPropertyIdOk returns a tuple with the PropertyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyId

`func (o *ReportFinding) SetPropertyId(v string)`

SetPropertyId sets PropertyId field to given value.


### GetReachability

`func (o *ReportFinding) GetReachability() string`

GetReachability returns the Reachability field if non-nil, zero value otherwise.

### GetReachabilityOk

`func (o *ReportFinding) GetReachabilityOk() (*string, bool)`

GetReachabilityOk returns a tuple with the Reachability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachability

`func (o *ReportFinding) SetReachability(v string)`

SetReachability sets Reachability field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


