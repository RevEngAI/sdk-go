# ListArchiveContentsOutputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entries** | [**[]ArchiveContentEntry**](ArchiveContentEntry.md) | Files inside the archive, with paths relative to the archive root | 
**HasNext** | **bool** | Whether a further page of entries follows this one. | 
**Page** | **int64** | Page number of this response (1-indexed). | 
**PageSize** | **int64** | Number of entries per page. | 
**TotalCount** | **int64** | Total number of file entries in the archive, ignoring pagination. | 

## Methods

### NewListArchiveContentsOutputBody

`func NewListArchiveContentsOutputBody(entries []ArchiveContentEntry, hasNext bool, page int64, pageSize int64, totalCount int64, ) *ListArchiveContentsOutputBody`

NewListArchiveContentsOutputBody instantiates a new ListArchiveContentsOutputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListArchiveContentsOutputBodyWithDefaults

`func NewListArchiveContentsOutputBodyWithDefaults() *ListArchiveContentsOutputBody`

NewListArchiveContentsOutputBodyWithDefaults instantiates a new ListArchiveContentsOutputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntries

`func (o *ListArchiveContentsOutputBody) GetEntries() []ArchiveContentEntry`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *ListArchiveContentsOutputBody) GetEntriesOk() (*[]ArchiveContentEntry, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *ListArchiveContentsOutputBody) SetEntries(v []ArchiveContentEntry)`

SetEntries sets Entries field to given value.


### SetEntriesNil

`func (o *ListArchiveContentsOutputBody) SetEntriesNil(b bool)`

 SetEntriesNil sets the value for Entries to be an explicit nil

### UnsetEntries
`func (o *ListArchiveContentsOutputBody) UnsetEntries()`

UnsetEntries ensures that no value is present for Entries, not even an explicit nil
### GetHasNext

`func (o *ListArchiveContentsOutputBody) GetHasNext() bool`

GetHasNext returns the HasNext field if non-nil, zero value otherwise.

### GetHasNextOk

`func (o *ListArchiveContentsOutputBody) GetHasNextOk() (*bool, bool)`

GetHasNextOk returns a tuple with the HasNext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNext

`func (o *ListArchiveContentsOutputBody) SetHasNext(v bool)`

SetHasNext sets HasNext field to given value.


### GetPage

`func (o *ListArchiveContentsOutputBody) GetPage() int64`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ListArchiveContentsOutputBody) GetPageOk() (*int64, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ListArchiveContentsOutputBody) SetPage(v int64)`

SetPage sets Page field to given value.


### GetPageSize

`func (o *ListArchiveContentsOutputBody) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ListArchiveContentsOutputBody) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ListArchiveContentsOutputBody) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.


### GetTotalCount

`func (o *ListArchiveContentsOutputBody) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ListArchiveContentsOutputBody) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ListArchiveContentsOutputBody) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


