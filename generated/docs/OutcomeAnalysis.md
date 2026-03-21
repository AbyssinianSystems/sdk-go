# OutcomeAnalysis

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | [**OutcomeAnalysisFilter**](OutcomeAnalysisFilter.md) |  | 
**Summary** | [**OutcomeAnalysisSummary**](OutcomeAnalysisSummary.md) |  | 
**Rows** | [**[]OutcomeAnalysisRow**](OutcomeAnalysisRow.md) |  | 
**ReasonCodeStats** | [**[]OutcomeReasonCodeStat**](OutcomeReasonCodeStat.md) |  | 
**Comparisons** | [**[]OutcomeComparison**](OutcomeComparison.md) |  | 

## Methods

### NewOutcomeAnalysis

`func NewOutcomeAnalysis(filter OutcomeAnalysisFilter, summary OutcomeAnalysisSummary, rows []OutcomeAnalysisRow, reasonCodeStats []OutcomeReasonCodeStat, comparisons []OutcomeComparison, ) *OutcomeAnalysis`

NewOutcomeAnalysis instantiates a new OutcomeAnalysis object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutcomeAnalysisWithDefaults

`func NewOutcomeAnalysisWithDefaults() *OutcomeAnalysis`

NewOutcomeAnalysisWithDefaults instantiates a new OutcomeAnalysis object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *OutcomeAnalysis) GetFilter() OutcomeAnalysisFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *OutcomeAnalysis) GetFilterOk() (*OutcomeAnalysisFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *OutcomeAnalysis) SetFilter(v OutcomeAnalysisFilter)`

SetFilter sets Filter field to given value.


### GetSummary

`func (o *OutcomeAnalysis) GetSummary() OutcomeAnalysisSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *OutcomeAnalysis) GetSummaryOk() (*OutcomeAnalysisSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *OutcomeAnalysis) SetSummary(v OutcomeAnalysisSummary)`

SetSummary sets Summary field to given value.


### GetRows

`func (o *OutcomeAnalysis) GetRows() []OutcomeAnalysisRow`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *OutcomeAnalysis) GetRowsOk() (*[]OutcomeAnalysisRow, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *OutcomeAnalysis) SetRows(v []OutcomeAnalysisRow)`

SetRows sets Rows field to given value.


### GetReasonCodeStats

`func (o *OutcomeAnalysis) GetReasonCodeStats() []OutcomeReasonCodeStat`

GetReasonCodeStats returns the ReasonCodeStats field if non-nil, zero value otherwise.

### GetReasonCodeStatsOk

`func (o *OutcomeAnalysis) GetReasonCodeStatsOk() (*[]OutcomeReasonCodeStat, bool)`

GetReasonCodeStatsOk returns a tuple with the ReasonCodeStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCodeStats

`func (o *OutcomeAnalysis) SetReasonCodeStats(v []OutcomeReasonCodeStat)`

SetReasonCodeStats sets ReasonCodeStats field to given value.


### GetComparisons

`func (o *OutcomeAnalysis) GetComparisons() []OutcomeComparison`

GetComparisons returns the Comparisons field if non-nil, zero value otherwise.

### GetComparisonsOk

`func (o *OutcomeAnalysis) GetComparisonsOk() (*[]OutcomeComparison, bool)`

GetComparisonsOk returns a tuple with the Comparisons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparisons

`func (o *OutcomeAnalysis) SetComparisons(v []OutcomeComparison)`

SetComparisons sets Comparisons field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


