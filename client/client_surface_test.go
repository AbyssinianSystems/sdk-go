package client

import (
	"reflect"
	"testing"
)

func TestClientPublicMethodSurface(t *testing.T) {
	type methodExpectation struct {
		name string
	}

	expected := []methodExpectation{
		{name: "ClearBearerToken"},
		{name: "CompareSubjectRuleset"},
		{name: "GetLatestSubjectEvaluation"},
		{name: "GetOutcomeAnalysis"},
		{name: "GetSubjectInvestigation"},
		{name: "Healthz"},
		{name: "IngestSignal"},
		{name: "IngestSignalAndGetLatestEvaluation"},
		{name: "ListSubjectEvaluations"},
		{name: "ListSubjectSignalEvents"},
		{name: "Livez"},
		{name: "Readyz"},
		{name: "RecordReviewOutcome"},
		{name: "RecomputeSubject"},
		{name: "SetBearerToken"},
	}

	clientType := reflect.TypeOf(&Client{})
	for _, method := range expected {
		if _, ok := clientType.MethodByName(method.name); !ok {
			t.Fatalf("Client method %q missing; public SDK surface changed", method.name)
		}
	}
}
