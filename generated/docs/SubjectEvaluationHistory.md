# SubjectEvaluationHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubjectId** | **string** |  | 
**Items** | [**[]SubjectEvaluationBundle**](SubjectEvaluationBundle.md) |  | 

## Methods

### NewSubjectEvaluationHistory

`func NewSubjectEvaluationHistory(subjectId string, items []SubjectEvaluationBundle, ) *SubjectEvaluationHistory`

NewSubjectEvaluationHistory instantiates a new SubjectEvaluationHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubjectEvaluationHistoryWithDefaults

`func NewSubjectEvaluationHistoryWithDefaults() *SubjectEvaluationHistory`

NewSubjectEvaluationHistoryWithDefaults instantiates a new SubjectEvaluationHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubjectId

`func (o *SubjectEvaluationHistory) GetSubjectId() string`

GetSubjectId returns the SubjectId field if non-nil, zero value otherwise.

### GetSubjectIdOk

`func (o *SubjectEvaluationHistory) GetSubjectIdOk() (*string, bool)`

GetSubjectIdOk returns a tuple with the SubjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectId

`func (o *SubjectEvaluationHistory) SetSubjectId(v string)`

SetSubjectId sets SubjectId field to given value.


### GetItems

`func (o *SubjectEvaluationHistory) GetItems() []SubjectEvaluationBundle`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SubjectEvaluationHistory) GetItemsOk() (*[]SubjectEvaluationBundle, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *SubjectEvaluationHistory) SetItems(v []SubjectEvaluationBundle)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


