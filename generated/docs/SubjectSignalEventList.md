# SubjectSignalEventList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubjectId** | **string** |  | 
**Items** | [**[]SignalEvent**](SignalEvent.md) |  | 

## Methods

### NewSubjectSignalEventList

`func NewSubjectSignalEventList(subjectId string, items []SignalEvent, ) *SubjectSignalEventList`

NewSubjectSignalEventList instantiates a new SubjectSignalEventList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubjectSignalEventListWithDefaults

`func NewSubjectSignalEventListWithDefaults() *SubjectSignalEventList`

NewSubjectSignalEventListWithDefaults instantiates a new SubjectSignalEventList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubjectId

`func (o *SubjectSignalEventList) GetSubjectId() string`

GetSubjectId returns the SubjectId field if non-nil, zero value otherwise.

### GetSubjectIdOk

`func (o *SubjectSignalEventList) GetSubjectIdOk() (*string, bool)`

GetSubjectIdOk returns a tuple with the SubjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectId

`func (o *SubjectSignalEventList) SetSubjectId(v string)`

SetSubjectId sets SubjectId field to given value.


### GetItems

`func (o *SubjectSignalEventList) GetItems() []SignalEvent`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SubjectSignalEventList) GetItemsOk() (*[]SignalEvent, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *SubjectSignalEventList) SetItems(v []SignalEvent)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


