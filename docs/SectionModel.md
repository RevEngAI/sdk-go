# SectionModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumberOfSections** | **int32** |  | 
**Sections** | [**[]SingleSectionModel**](SingleSectionModel.md) |  | 

## Methods

### NewSectionModel

`func NewSectionModel(numberOfSections int32, sections []SingleSectionModel, ) *SectionModel`

NewSectionModel instantiates a new SectionModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSectionModelWithDefaults

`func NewSectionModelWithDefaults() *SectionModel`

NewSectionModelWithDefaults instantiates a new SectionModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumberOfSections

`func (o *SectionModel) GetNumberOfSections() int32`

GetNumberOfSections returns the NumberOfSections field if non-nil, zero value otherwise.

### GetNumberOfSectionsOk

`func (o *SectionModel) GetNumberOfSectionsOk() (*int32, bool)`

GetNumberOfSectionsOk returns a tuple with the NumberOfSections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfSections

`func (o *SectionModel) SetNumberOfSections(v int32)`

SetNumberOfSections sets NumberOfSections field to given value.


### GetSections

`func (o *SectionModel) GetSections() []SingleSectionModel`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *SectionModel) GetSectionsOk() (*[]SingleSectionModel, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *SectionModel) SetSections(v []SingleSectionModel)`

SetSections sets Sections field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


