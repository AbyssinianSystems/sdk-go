# ExplanationRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RuleId** | **string** |  | 
**Label** | **string** |  | 
**Feature** | **string** |  | 
**Contribution** | **float32** |  | 
**Category** | **string** |  | 
**Evidence** | Pointer to [**[]EvidenceRef**](EvidenceRef.md) |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 

## Methods

### NewExplanationRule

`func NewExplanationRule(ruleId string, label string, feature string, contribution float32, category string, ) *ExplanationRule`

NewExplanationRule instantiates a new ExplanationRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExplanationRuleWithDefaults

`func NewExplanationRuleWithDefaults() *ExplanationRule`

NewExplanationRuleWithDefaults instantiates a new ExplanationRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRuleId

`func (o *ExplanationRule) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *ExplanationRule) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *ExplanationRule) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.


### GetLabel

`func (o *ExplanationRule) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ExplanationRule) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ExplanationRule) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetFeature

`func (o *ExplanationRule) GetFeature() string`

GetFeature returns the Feature field if non-nil, zero value otherwise.

### GetFeatureOk

`func (o *ExplanationRule) GetFeatureOk() (*string, bool)`

GetFeatureOk returns a tuple with the Feature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeature

`func (o *ExplanationRule) SetFeature(v string)`

SetFeature sets Feature field to given value.


### GetContribution

`func (o *ExplanationRule) GetContribution() float32`

GetContribution returns the Contribution field if non-nil, zero value otherwise.

### GetContributionOk

`func (o *ExplanationRule) GetContributionOk() (*float32, bool)`

GetContributionOk returns a tuple with the Contribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContribution

`func (o *ExplanationRule) SetContribution(v float32)`

SetContribution sets Contribution field to given value.


### GetCategory

`func (o *ExplanationRule) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ExplanationRule) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ExplanationRule) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetEvidence

`func (o *ExplanationRule) GetEvidence() []EvidenceRef`

GetEvidence returns the Evidence field if non-nil, zero value otherwise.

### GetEvidenceOk

`func (o *ExplanationRule) GetEvidenceOk() (*[]EvidenceRef, bool)`

GetEvidenceOk returns a tuple with the Evidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidence

`func (o *ExplanationRule) SetEvidence(v []EvidenceRef)`

SetEvidence sets Evidence field to given value.

### HasEvidence

`func (o *ExplanationRule) HasEvidence() bool`

HasEvidence returns a boolean if a field has been set.

### GetNotes

`func (o *ExplanationRule) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ExplanationRule) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ExplanationRule) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *ExplanationRule) HasNotes() bool`

HasNotes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


