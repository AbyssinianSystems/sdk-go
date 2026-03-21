# ReviewOutcome

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**SubjectId** | **string** |  | 
**ReviewedObject** | [**ArtifactRef**](ArtifactRef.md) |  | 
**Reviewer** | **string** |  | 
**Label** | **string** |  | 
**ReviewedAt** | **time.Time** |  | 
**Notes** | Pointer to **string** |  | [optional] 
**Disposition** | Pointer to **string** |  | [optional] 

## Methods

### NewReviewOutcome

`func NewReviewOutcome(id string, subjectId string, reviewedObject ArtifactRef, reviewer string, label string, reviewedAt time.Time, ) *ReviewOutcome`

NewReviewOutcome instantiates a new ReviewOutcome object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReviewOutcomeWithDefaults

`func NewReviewOutcomeWithDefaults() *ReviewOutcome`

NewReviewOutcomeWithDefaults instantiates a new ReviewOutcome object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReviewOutcome) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReviewOutcome) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReviewOutcome) SetId(v string)`

SetId sets Id field to given value.


### GetSubjectId

`func (o *ReviewOutcome) GetSubjectId() string`

GetSubjectId returns the SubjectId field if non-nil, zero value otherwise.

### GetSubjectIdOk

`func (o *ReviewOutcome) GetSubjectIdOk() (*string, bool)`

GetSubjectIdOk returns a tuple with the SubjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectId

`func (o *ReviewOutcome) SetSubjectId(v string)`

SetSubjectId sets SubjectId field to given value.


### GetReviewedObject

`func (o *ReviewOutcome) GetReviewedObject() ArtifactRef`

GetReviewedObject returns the ReviewedObject field if non-nil, zero value otherwise.

### GetReviewedObjectOk

`func (o *ReviewOutcome) GetReviewedObjectOk() (*ArtifactRef, bool)`

GetReviewedObjectOk returns a tuple with the ReviewedObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewedObject

`func (o *ReviewOutcome) SetReviewedObject(v ArtifactRef)`

SetReviewedObject sets ReviewedObject field to given value.


### GetReviewer

`func (o *ReviewOutcome) GetReviewer() string`

GetReviewer returns the Reviewer field if non-nil, zero value otherwise.

### GetReviewerOk

`func (o *ReviewOutcome) GetReviewerOk() (*string, bool)`

GetReviewerOk returns a tuple with the Reviewer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewer

`func (o *ReviewOutcome) SetReviewer(v string)`

SetReviewer sets Reviewer field to given value.


### GetLabel

`func (o *ReviewOutcome) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ReviewOutcome) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ReviewOutcome) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetReviewedAt

`func (o *ReviewOutcome) GetReviewedAt() time.Time`

GetReviewedAt returns the ReviewedAt field if non-nil, zero value otherwise.

### GetReviewedAtOk

`func (o *ReviewOutcome) GetReviewedAtOk() (*time.Time, bool)`

GetReviewedAtOk returns a tuple with the ReviewedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewedAt

`func (o *ReviewOutcome) SetReviewedAt(v time.Time)`

SetReviewedAt sets ReviewedAt field to given value.


### GetNotes

`func (o *ReviewOutcome) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ReviewOutcome) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ReviewOutcome) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *ReviewOutcome) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetDisposition

`func (o *ReviewOutcome) GetDisposition() string`

GetDisposition returns the Disposition field if non-nil, zero value otherwise.

### GetDispositionOk

`func (o *ReviewOutcome) GetDispositionOk() (*string, bool)`

GetDispositionOk returns a tuple with the Disposition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisposition

`func (o *ReviewOutcome) SetDisposition(v string)`

SetDisposition sets Disposition field to given value.

### HasDisposition

`func (o *ReviewOutcome) HasDisposition() bool`

HasDisposition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


