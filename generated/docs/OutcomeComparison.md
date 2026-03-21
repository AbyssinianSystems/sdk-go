# OutcomeComparison

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubjectId** | **string** |  | 
**Recommendation** | Pointer to [**Recommendation**](Recommendation.md) |  | [optional] 
**ReviewOutcome** | [**ReviewOutcome**](ReviewOutcome.md) |  | 
**RecommendationType** | Pointer to **string** |  | [optional] 
**PrimaryReason** | Pointer to [**RecommendationReasonCode**](RecommendationReasonCode.md) |  | [optional] 
**AdditionalReasons** | Pointer to [**[]RecommendationReasonCode**](RecommendationReasonCode.md) |  | [optional] 
**ReviewLabel** | **string** |  | 
**OutcomeClassification** | [**OutcomeClassification**](OutcomeClassification.md) |  | 

## Methods

### NewOutcomeComparison

`func NewOutcomeComparison(subjectId string, reviewOutcome ReviewOutcome, reviewLabel string, outcomeClassification OutcomeClassification, ) *OutcomeComparison`

NewOutcomeComparison instantiates a new OutcomeComparison object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutcomeComparisonWithDefaults

`func NewOutcomeComparisonWithDefaults() *OutcomeComparison`

NewOutcomeComparisonWithDefaults instantiates a new OutcomeComparison object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubjectId

`func (o *OutcomeComparison) GetSubjectId() string`

GetSubjectId returns the SubjectId field if non-nil, zero value otherwise.

### GetSubjectIdOk

`func (o *OutcomeComparison) GetSubjectIdOk() (*string, bool)`

GetSubjectIdOk returns a tuple with the SubjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectId

`func (o *OutcomeComparison) SetSubjectId(v string)`

SetSubjectId sets SubjectId field to given value.


### GetRecommendation

`func (o *OutcomeComparison) GetRecommendation() Recommendation`

GetRecommendation returns the Recommendation field if non-nil, zero value otherwise.

### GetRecommendationOk

`func (o *OutcomeComparison) GetRecommendationOk() (*Recommendation, bool)`

GetRecommendationOk returns a tuple with the Recommendation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendation

`func (o *OutcomeComparison) SetRecommendation(v Recommendation)`

SetRecommendation sets Recommendation field to given value.

### HasRecommendation

`func (o *OutcomeComparison) HasRecommendation() bool`

HasRecommendation returns a boolean if a field has been set.

### GetReviewOutcome

`func (o *OutcomeComparison) GetReviewOutcome() ReviewOutcome`

GetReviewOutcome returns the ReviewOutcome field if non-nil, zero value otherwise.

### GetReviewOutcomeOk

`func (o *OutcomeComparison) GetReviewOutcomeOk() (*ReviewOutcome, bool)`

GetReviewOutcomeOk returns a tuple with the ReviewOutcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewOutcome

`func (o *OutcomeComparison) SetReviewOutcome(v ReviewOutcome)`

SetReviewOutcome sets ReviewOutcome field to given value.


### GetRecommendationType

`func (o *OutcomeComparison) GetRecommendationType() string`

GetRecommendationType returns the RecommendationType field if non-nil, zero value otherwise.

### GetRecommendationTypeOk

`func (o *OutcomeComparison) GetRecommendationTypeOk() (*string, bool)`

GetRecommendationTypeOk returns a tuple with the RecommendationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendationType

`func (o *OutcomeComparison) SetRecommendationType(v string)`

SetRecommendationType sets RecommendationType field to given value.

### HasRecommendationType

`func (o *OutcomeComparison) HasRecommendationType() bool`

HasRecommendationType returns a boolean if a field has been set.

### GetPrimaryReason

`func (o *OutcomeComparison) GetPrimaryReason() RecommendationReasonCode`

GetPrimaryReason returns the PrimaryReason field if non-nil, zero value otherwise.

### GetPrimaryReasonOk

`func (o *OutcomeComparison) GetPrimaryReasonOk() (*RecommendationReasonCode, bool)`

GetPrimaryReasonOk returns a tuple with the PrimaryReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryReason

`func (o *OutcomeComparison) SetPrimaryReason(v RecommendationReasonCode)`

SetPrimaryReason sets PrimaryReason field to given value.

### HasPrimaryReason

`func (o *OutcomeComparison) HasPrimaryReason() bool`

HasPrimaryReason returns a boolean if a field has been set.

### GetAdditionalReasons

`func (o *OutcomeComparison) GetAdditionalReasons() []RecommendationReasonCode`

GetAdditionalReasons returns the AdditionalReasons field if non-nil, zero value otherwise.

### GetAdditionalReasonsOk

`func (o *OutcomeComparison) GetAdditionalReasonsOk() (*[]RecommendationReasonCode, bool)`

GetAdditionalReasonsOk returns a tuple with the AdditionalReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalReasons

`func (o *OutcomeComparison) SetAdditionalReasons(v []RecommendationReasonCode)`

SetAdditionalReasons sets AdditionalReasons field to given value.

### HasAdditionalReasons

`func (o *OutcomeComparison) HasAdditionalReasons() bool`

HasAdditionalReasons returns a boolean if a field has been set.

### GetReviewLabel

`func (o *OutcomeComparison) GetReviewLabel() string`

GetReviewLabel returns the ReviewLabel field if non-nil, zero value otherwise.

### GetReviewLabelOk

`func (o *OutcomeComparison) GetReviewLabelOk() (*string, bool)`

GetReviewLabelOk returns a tuple with the ReviewLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewLabel

`func (o *OutcomeComparison) SetReviewLabel(v string)`

SetReviewLabel sets ReviewLabel field to given value.


### GetOutcomeClassification

`func (o *OutcomeComparison) GetOutcomeClassification() OutcomeClassification`

GetOutcomeClassification returns the OutcomeClassification field if non-nil, zero value otherwise.

### GetOutcomeClassificationOk

`func (o *OutcomeComparison) GetOutcomeClassificationOk() (*OutcomeClassification, bool)`

GetOutcomeClassificationOk returns a tuple with the OutcomeClassification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcomeClassification

`func (o *OutcomeComparison) SetOutcomeClassification(v OutcomeClassification)`

SetOutcomeClassification sets OutcomeClassification field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


