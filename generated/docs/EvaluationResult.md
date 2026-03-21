# EvaluationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**Reason** | Pointer to **string** |  | [optional] 
**SnapshotRef** | Pointer to [**ArtifactRef**](ArtifactRef.md) |  | [optional] 
**ScoreResultRef** | Pointer to [**ArtifactRef**](ArtifactRef.md) |  | [optional] 
**RecommendationRef** | Pointer to [**ArtifactRef**](ArtifactRef.md) |  | [optional] 

## Methods

### NewEvaluationResult

`func NewEvaluationResult(status string, ) *EvaluationResult`

NewEvaluationResult instantiates a new EvaluationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvaluationResultWithDefaults

`func NewEvaluationResultWithDefaults() *EvaluationResult`

NewEvaluationResultWithDefaults instantiates a new EvaluationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *EvaluationResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EvaluationResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EvaluationResult) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetReason

`func (o *EvaluationResult) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *EvaluationResult) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *EvaluationResult) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *EvaluationResult) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetSnapshotRef

`func (o *EvaluationResult) GetSnapshotRef() ArtifactRef`

GetSnapshotRef returns the SnapshotRef field if non-nil, zero value otherwise.

### GetSnapshotRefOk

`func (o *EvaluationResult) GetSnapshotRefOk() (*ArtifactRef, bool)`

GetSnapshotRefOk returns a tuple with the SnapshotRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotRef

`func (o *EvaluationResult) SetSnapshotRef(v ArtifactRef)`

SetSnapshotRef sets SnapshotRef field to given value.

### HasSnapshotRef

`func (o *EvaluationResult) HasSnapshotRef() bool`

HasSnapshotRef returns a boolean if a field has been set.

### GetScoreResultRef

`func (o *EvaluationResult) GetScoreResultRef() ArtifactRef`

GetScoreResultRef returns the ScoreResultRef field if non-nil, zero value otherwise.

### GetScoreResultRefOk

`func (o *EvaluationResult) GetScoreResultRefOk() (*ArtifactRef, bool)`

GetScoreResultRefOk returns a tuple with the ScoreResultRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreResultRef

`func (o *EvaluationResult) SetScoreResultRef(v ArtifactRef)`

SetScoreResultRef sets ScoreResultRef field to given value.

### HasScoreResultRef

`func (o *EvaluationResult) HasScoreResultRef() bool`

HasScoreResultRef returns a boolean if a field has been set.

### GetRecommendationRef

`func (o *EvaluationResult) GetRecommendationRef() ArtifactRef`

GetRecommendationRef returns the RecommendationRef field if non-nil, zero value otherwise.

### GetRecommendationRefOk

`func (o *EvaluationResult) GetRecommendationRefOk() (*ArtifactRef, bool)`

GetRecommendationRefOk returns a tuple with the RecommendationRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendationRef

`func (o *EvaluationResult) SetRecommendationRef(v ArtifactRef)`

SetRecommendationRef sets RecommendationRef field to given value.

### HasRecommendationRef

`func (o *EvaluationResult) HasRecommendationRef() bool`

HasRecommendationRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


