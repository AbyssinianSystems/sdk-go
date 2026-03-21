# ReviewOutcomeWriteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReviewedObject** | [**ReviewTargetRef**](ReviewTargetRef.md) |  | 
**Reviewer** | **string** |  | 
**Label** | **string** |  | 
**Notes** | Pointer to **string** |  | [optional] 
**Disposition** | Pointer to **string** |  | [optional] 

## Methods

### NewReviewOutcomeWriteRequest

`func NewReviewOutcomeWriteRequest(reviewedObject ReviewTargetRef, reviewer string, label string, ) *ReviewOutcomeWriteRequest`

NewReviewOutcomeWriteRequest instantiates a new ReviewOutcomeWriteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReviewOutcomeWriteRequestWithDefaults

`func NewReviewOutcomeWriteRequestWithDefaults() *ReviewOutcomeWriteRequest`

NewReviewOutcomeWriteRequestWithDefaults instantiates a new ReviewOutcomeWriteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReviewedObject

`func (o *ReviewOutcomeWriteRequest) GetReviewedObject() ReviewTargetRef`

GetReviewedObject returns the ReviewedObject field if non-nil, zero value otherwise.

### GetReviewedObjectOk

`func (o *ReviewOutcomeWriteRequest) GetReviewedObjectOk() (*ReviewTargetRef, bool)`

GetReviewedObjectOk returns a tuple with the ReviewedObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewedObject

`func (o *ReviewOutcomeWriteRequest) SetReviewedObject(v ReviewTargetRef)`

SetReviewedObject sets ReviewedObject field to given value.


### GetReviewer

`func (o *ReviewOutcomeWriteRequest) GetReviewer() string`

GetReviewer returns the Reviewer field if non-nil, zero value otherwise.

### GetReviewerOk

`func (o *ReviewOutcomeWriteRequest) GetReviewerOk() (*string, bool)`

GetReviewerOk returns a tuple with the Reviewer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewer

`func (o *ReviewOutcomeWriteRequest) SetReviewer(v string)`

SetReviewer sets Reviewer field to given value.


### GetLabel

`func (o *ReviewOutcomeWriteRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ReviewOutcomeWriteRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ReviewOutcomeWriteRequest) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetNotes

`func (o *ReviewOutcomeWriteRequest) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ReviewOutcomeWriteRequest) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ReviewOutcomeWriteRequest) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *ReviewOutcomeWriteRequest) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetDisposition

`func (o *ReviewOutcomeWriteRequest) GetDisposition() string`

GetDisposition returns the Disposition field if non-nil, zero value otherwise.

### GetDispositionOk

`func (o *ReviewOutcomeWriteRequest) GetDispositionOk() (*string, bool)`

GetDispositionOk returns a tuple with the Disposition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisposition

`func (o *ReviewOutcomeWriteRequest) SetDisposition(v string)`

SetDisposition sets Disposition field to given value.

### HasDisposition

`func (o *ReviewOutcomeWriteRequest) HasDisposition() bool`

HasDisposition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


