# OutcomeAnalysisSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComparisonCount** | **int32** |  | 
**RecommendationCounts** | [**[]OutcomeCountByRecommendationType**](OutcomeCountByRecommendationType.md) |  | 
**ReviewLabelCounts** | [**[]OutcomeCountByReviewLabel**](OutcomeCountByReviewLabel.md) |  | 
**ClassificationCounts** | [**[]OutcomeCountByClassification**](OutcomeCountByClassification.md) |  | 
**AlignedCount** | **int32** |  | 
**MismatchedCount** | **int32** |  | 
**UnresolvedCount** | **int32** |  | 

## Methods

### NewOutcomeAnalysisSummary

`func NewOutcomeAnalysisSummary(comparisonCount int32, recommendationCounts []OutcomeCountByRecommendationType, reviewLabelCounts []OutcomeCountByReviewLabel, classificationCounts []OutcomeCountByClassification, alignedCount int32, mismatchedCount int32, unresolvedCount int32, ) *OutcomeAnalysisSummary`

NewOutcomeAnalysisSummary instantiates a new OutcomeAnalysisSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutcomeAnalysisSummaryWithDefaults

`func NewOutcomeAnalysisSummaryWithDefaults() *OutcomeAnalysisSummary`

NewOutcomeAnalysisSummaryWithDefaults instantiates a new OutcomeAnalysisSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComparisonCount

`func (o *OutcomeAnalysisSummary) GetComparisonCount() int32`

GetComparisonCount returns the ComparisonCount field if non-nil, zero value otherwise.

### GetComparisonCountOk

`func (o *OutcomeAnalysisSummary) GetComparisonCountOk() (*int32, bool)`

GetComparisonCountOk returns a tuple with the ComparisonCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparisonCount

`func (o *OutcomeAnalysisSummary) SetComparisonCount(v int32)`

SetComparisonCount sets ComparisonCount field to given value.


### GetRecommendationCounts

`func (o *OutcomeAnalysisSummary) GetRecommendationCounts() []OutcomeCountByRecommendationType`

GetRecommendationCounts returns the RecommendationCounts field if non-nil, zero value otherwise.

### GetRecommendationCountsOk

`func (o *OutcomeAnalysisSummary) GetRecommendationCountsOk() (*[]OutcomeCountByRecommendationType, bool)`

GetRecommendationCountsOk returns a tuple with the RecommendationCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendationCounts

`func (o *OutcomeAnalysisSummary) SetRecommendationCounts(v []OutcomeCountByRecommendationType)`

SetRecommendationCounts sets RecommendationCounts field to given value.


### GetReviewLabelCounts

`func (o *OutcomeAnalysisSummary) GetReviewLabelCounts() []OutcomeCountByReviewLabel`

GetReviewLabelCounts returns the ReviewLabelCounts field if non-nil, zero value otherwise.

### GetReviewLabelCountsOk

`func (o *OutcomeAnalysisSummary) GetReviewLabelCountsOk() (*[]OutcomeCountByReviewLabel, bool)`

GetReviewLabelCountsOk returns a tuple with the ReviewLabelCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewLabelCounts

`func (o *OutcomeAnalysisSummary) SetReviewLabelCounts(v []OutcomeCountByReviewLabel)`

SetReviewLabelCounts sets ReviewLabelCounts field to given value.


### GetClassificationCounts

`func (o *OutcomeAnalysisSummary) GetClassificationCounts() []OutcomeCountByClassification`

GetClassificationCounts returns the ClassificationCounts field if non-nil, zero value otherwise.

### GetClassificationCountsOk

`func (o *OutcomeAnalysisSummary) GetClassificationCountsOk() (*[]OutcomeCountByClassification, bool)`

GetClassificationCountsOk returns a tuple with the ClassificationCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassificationCounts

`func (o *OutcomeAnalysisSummary) SetClassificationCounts(v []OutcomeCountByClassification)`

SetClassificationCounts sets ClassificationCounts field to given value.


### GetAlignedCount

`func (o *OutcomeAnalysisSummary) GetAlignedCount() int32`

GetAlignedCount returns the AlignedCount field if non-nil, zero value otherwise.

### GetAlignedCountOk

`func (o *OutcomeAnalysisSummary) GetAlignedCountOk() (*int32, bool)`

GetAlignedCountOk returns a tuple with the AlignedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlignedCount

`func (o *OutcomeAnalysisSummary) SetAlignedCount(v int32)`

SetAlignedCount sets AlignedCount field to given value.


### GetMismatchedCount

`func (o *OutcomeAnalysisSummary) GetMismatchedCount() int32`

GetMismatchedCount returns the MismatchedCount field if non-nil, zero value otherwise.

### GetMismatchedCountOk

`func (o *OutcomeAnalysisSummary) GetMismatchedCountOk() (*int32, bool)`

GetMismatchedCountOk returns a tuple with the MismatchedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMismatchedCount

`func (o *OutcomeAnalysisSummary) SetMismatchedCount(v int32)`

SetMismatchedCount sets MismatchedCount field to given value.


### GetUnresolvedCount

`func (o *OutcomeAnalysisSummary) GetUnresolvedCount() int32`

GetUnresolvedCount returns the UnresolvedCount field if non-nil, zero value otherwise.

### GetUnresolvedCountOk

`func (o *OutcomeAnalysisSummary) GetUnresolvedCountOk() (*int32, bool)`

GetUnresolvedCountOk returns a tuple with the UnresolvedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnresolvedCount

`func (o *OutcomeAnalysisSummary) SetUnresolvedCount(v int32)`

SetUnresolvedCount sets UnresolvedCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


