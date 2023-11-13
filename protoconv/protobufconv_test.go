package protoconv_test

import (
	attrs "github.com/virtru/access-pdp/attributes"
	pbPDP "github.com/virtru/access-pdp/proto/accesspdp/v1"
	pbAttr "github.com/virtru/access-pdp/proto/attributes/v1"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/virtru/access-pdp/pdp"
	"github.com/virtru/access-pdp/protoconv"
)

func TestValueFailureToPbWithNilDataAttribute(t *testing.T) {
	failure := &pdp.ValueFailure{
		DataAttribute: nil,
		Message:       "test message",
	}

	pbFailure := protoconv.ValueFailureToPb(failure)

	require.NotNil(t, pbFailure, "The resulting pbFailure should not be nil")
	require.Nil(t, pbFailure.DataAttribute, "The DataAttribute in pbFailure should be nil")
	require.Equal(t, "test message", pbFailure.Message, "The Message in pbFailure should match the input")
}

func TestAttributeDefinitionToPb(t *testing.T) {
	type args struct {
		def *attrs.AttributeDefinition
	}
	testState := "published"
	tests := []struct {
		name string
		args args
		want *pbAttr.AttributeDefinition
	}{
		{
			name: "positive",
			args: args{
				def: &attrs.AttributeDefinition{
					Authority: "a",
					Name:      "b",
					Rule:      "c",
					State:     testState,
					Order:     nil,
					GroupBy:   nil,
				},
			},
			want: &pbAttr.AttributeDefinition{
				Authority: "a",
				Name:      "b",
				Rule:      "c",
				State:     &testState,
				GroupBy:   nil,
				Order:     nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := protoconv.AttributeDefinitionToPb(tt.args.def); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AttributeDefinitionToPb() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAttributeInstanceToPb(t *testing.T) {
	type args struct {
		def *attrs.AttributeInstance
	}
	tests := []struct {
		name string
		args args
		want *pbAttr.AttributeInstance
	}{
		{
			name: "positive",
			args: args{
				def: &attrs.AttributeInstance{
					Authority: "a",
					Name:      "b",
					Value:     "c",
				},
			},
			want: &pbAttr.AttributeInstance{
				Authority: "a",
				Name:      "b",
				Value:     "c",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := protoconv.AttributeInstanceToPb(tt.args.def); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AttributeInstanceToPb() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDataRuleResultsToPb(t *testing.T) {
	type args struct {
		results []pdp.DataRuleResult
	}
	tests := []struct {
		name string
		args args
		want []*pbPDP.DataRuleResult
	}{
		{
			name: "positive-nil",
			args: args{
				results: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := protoconv.DataRuleResultsToPb(tt.args.results); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DataRuleResultsToPb() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecisionToPbResponse(t *testing.T) {
	type args struct {
		entity   string
		decision *pdp.Decision
	}
	tests := []struct {
		name string
		args args
		want *pbPDP.DetermineAccessResponse
	}{
		{
			name: "positive-nil",
			args: args{
				entity: "a",
				decision: &pdp.Decision{
					Access:  false,
					Results: nil,
				},
			},
			want: &pbPDP.DetermineAccessResponse{
				Entity:  "a",
				Access:  false,
				Results: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := protoconv.DecisionToPbResponse(tt.args.entity, tt.args.decision); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecisionToPbResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPbToAttributeDefinition(t *testing.T) {
	type args struct {
		pbdef *pbAttr.AttributeDefinition
	}
	tests := []struct {
		name string
		args args
		want *attrs.AttributeDefinition
	}{
		{
			name: "positive-nil",
			args: args{
				pbdef: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := protoconv.PbToAttributeDefinition(tt.args.pbdef); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PbToAttributeDefinition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPbToAttributeDefinitions(t *testing.T) {
	type args struct {
		pbdefs []*pbAttr.AttributeDefinition
	}
	tests := []struct {
		name string
		args args
		want []attrs.AttributeDefinition
	}{
		{
			name: "positive-nil",
			args: args{
				pbdefs: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := protoconv.PbToAttributeDefinitions(tt.args.pbdefs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PbToAttributeDefinitions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPbToAttributeInstances(t *testing.T) {
	type args struct {
		pbinst []*pbAttr.AttributeInstance
	}
	tests := []struct {
		name string
		args args
		want []attrs.AttributeInstance
	}{
		{
			name: "positive-nil",
			args: args{
				pbinst: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := protoconv.PbToAttributeInstances(tt.args.pbinst); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PbToAttributeInstances() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPbToEntityAttrSets(t *testing.T) {
	type args struct {
		pbsets map[string]*pbPDP.ListOfAttributeInstances
	}
	tests := []struct {
		name string
		args args
		want map[string][]attrs.AttributeInstance
	}{
		{
			name: "positive-nil",
			args: args{
				pbsets: nil,
			},
			want: make(map[string][]attrs.AttributeInstance),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := protoconv.PbToEntityAttrSets(tt.args.pbsets); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PbToEntityAttrSets() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValueFailureToPb(t *testing.T) {
	type args struct {
		failure *pdp.ValueFailure
	}
	tests := []struct {
		name string
		args args
		want *pbPDP.ValueFailure
	}{
		{
			name: "negative-nil",
			args: args{
				failure: nil,
			},
			want: &pbPDP.ValueFailure{
				DataAttribute: nil,
				Message:       "failure is nil",
			},
		},
		{
			name: "negative-data-nil",
			args: args{
				failure: &pdp.ValueFailure{
					DataAttribute: nil,
					Message:       "data attribute is nil",
				},
			},
			want: &pbPDP.ValueFailure{
				DataAttribute: nil,
				Message:       "data attribute is nil",
			},
		},
		{
			name: "positive",
			args: args{
				failure: &pdp.ValueFailure{
					DataAttribute: &attrs.AttributeInstance{
						Authority: "a",
						Name:      "b",
						Value:     "c",
					},
					Message: "invalid authority",
				},
			},
			want: &pbPDP.ValueFailure{
				DataAttribute: &pbAttr.AttributeInstance{
					Authority: "a",
					Name:      "b",
					Value:     "c",
				},
				Message: "invalid authority",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := protoconv.ValueFailureToPb(tt.args.failure); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValueFailureToPb() = %v, want %v", got, tt.want)
			}
		})
	}
}
