# OutcomeAnalysisRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecommendationType** | Pointer to **string** |  | [optional] 
**PrimaryReason** | Pointer to [**RecommendationReasonCode**](RecommendationReasonCode.md) |  | [optional] 
**ReviewLabel** | **string** |  | 
**OutcomeClassification** | [**OutcomeClassification**](OutcomeClassification.md) |  | 
**Count** | **int32** |  | 

## Methods

### NewOutcomeAnalysisRow

`func NewOutcomeAnalysisRow(reviewLabel string, outcomeClassification OutcomeClassification, count int32, ) *OutcomeAnalysisRow`

NewOutcomeAnalysisRow instantiates a new OutcomeAnalysisRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutcomeAnalysisRowWithDefaults

`func NewOutcomeAnalysisRowWithDefaults() *OutcomeAnalysisRow`

NewOutcomeAnalysisRowWithDefaults instantiates a new OutcomeAnalysisRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecommendationType

`func (o *OutcomeAnalysisRow) GetRecommendationType() string`

GetRecommendationType returns the RecommendationType field if non-nil, zero value otherwise.

### GetRecommendationTypeOk

`func (o *OutcomeAnalysisRow) GetRecommendationTypeOk() (*string, bool)`

GetRecommendationTypeOk returns a tuple with the RecommendationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendationType

`func (o *OutcomeAnalysisRow) SetRecommendationType(v string)`

SetRecommendationType sets RecommendationType field to given value.

### HasRecommendationType

`func (o *OutcomeAnalysisRow) HasRecommendationType() bool`

HasRecommendationType returns a boolean if a field has been set.

### GetPrimaryReason

`func (o *OutcomeAnalysisRow) GetPrimaryReason() RecommendationReasonCode`

GetPrimaryReason returns the PrimaryReason field if non-nil, zero value otherwise.

### GetPrimaryReasonOk

`func (o *OutcomeAnalysisRow) GetPrimaryReasonOk() (*RecommendationReasonCode, bool)`

GetPrimaryReasonOk returns a tuple with the PrimaryReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryReason

`func (o *OutcomeAnalysisRow) SetPrimaryReason(v RecommendationReasonCode)`

SetPrimaryReason sets PrimaryReason field to given value.

### HasPrimaryReason

`func (o *OutcomeAnalysisRow) HasPrimaryReason() bool`

HasPrimaryReason returns a boolean if a field has been set.

### GetReviewLabel

`func (o *OutcomeAnalysisRow) GetReviewLabel() string`

GetReviewLabel returns the ReviewLabel field if non-nil, zero value otherwise.

### GetReviewLabelOk

`func (o *OutcomeAnalysisRow) GetReviewLabelOk() (*string, bool)`

GetReviewLabelOk returns a tuple with the ReviewLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewLabel

`func (o *OutcomeAnalysisRow) SetReviewLabel(v string)`

SetReviewLabel sets ReviewLabel field to given value.


### GetOutcomeClassification

`func (o *OutcomeAnalysisRow) GetOutcomeClassification() OutcomeClassification`

GetOutcomeClassification returns the OutcomeClassification field if non-nil, zero value otherwise.

### GetOutcomeClassificationOk

`func (o *OutcomeAnalysisRow) GetOutcomeClassificationOk() (*OutcomeClassification, bool)`

GetOutcomeClassificationOk returns a tuple with the OutcomeClassification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcomeClassification

`func (o *OutcomeAnalysisRow) SetOutcomeClassification(v OutcomeClassification)`

SetOutcomeClassification sets OutcomeClassification field to given value.


### GetCount

`func (o *OutcomeAnalysisRow) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *OutcomeAnalysisRow) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *OutcomeAnalysisRow) SetCount(v int32)`

SetCount sets Count field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


