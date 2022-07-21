package protoconv

import (
	pbPDP "github.com/virtru/access-pdp/v1/proto/accesspdp"
	pbAttr "github.com/virtru/access-pdp/v1/proto/attributes"

	attrs "github.com/virtru/access-pdp/v1/attributes"
	pdp "github.com/virtru/access-pdp/v1/pdp"
)

//Since codegen'd protobuf Go structs are (annoyingly, and unfixably, though arguably architecturally reasonably) not
//directly mappable to user-created Go structs, and are also not castable to those structs,
//conversion helper routines are required.

func PbToAttributeInstances(pbinst []*pbAttr.AttributeInstance) []attrs.AttributeInstance {
	var instances []attrs.AttributeInstance

	for _, v := range pbinst {
		instances = append(instances, attrs.AttributeInstance{Authority: v.Authority, Name: v.Name, Value: v.Value})
	}

	return instances
}

func AttributeInstanceToPb(def *attrs.AttributeInstance) *pbAttr.AttributeInstance {
	pbInst := pbAttr.AttributeInstance{
		Authority: def.Authority,
		Name:      def.Name,
		Value:     def.Value,
	}

	return &pbInst
}

func PbToAttributeDefinition(pbdef *pbAttr.AttributeDefinition) *attrs.AttributeDefinition {
	var def *attrs.AttributeDefinition
	if pbdef != nil {
		convAttr := &attrs.AttributeDefinition{
			Authority: pbdef.Authority,
			Name:      pbdef.Name,
			Rule:      pbdef.Rule,
			State:     *pbdef.State,
			Order:     pbdef.Order,
		}

		//GroupBy is optional - if it is present, it is just represented as another AttributeInstance
		if pbdef.GroupBy != nil {
			convAttr.GroupBy = &attrs.AttributeInstance{Authority: pbdef.GroupBy.Authority, Name: pbdef.GroupBy.Name, Value: pbdef.GroupBy.Value}
		}

		def = convAttr
	}
	return def
}

func PbToAttributeDefinitions(pbdefs []*pbAttr.AttributeDefinition) []attrs.AttributeDefinition {
	var defs []attrs.AttributeDefinition

	for _, v := range pbdefs {
		defs = append(defs, *PbToAttributeDefinition(v))
	}

	return defs
}

func AttributeDefinitionToPb(def *attrs.AttributeDefinition) *pbAttr.AttributeDefinition {
	pbDef := pbAttr.AttributeDefinition{
		Authority: def.Authority,
		Name:      def.Name,
		Rule:      def.Rule,
		State:     &def.State,
		Order:     def.Order,
	}

	//GroupBy is optional - if it is present, it is just represented as another AttributeInstance
	if def.GroupBy != nil {
		pbDef.GroupBy = &pbAttr.AttributeInstance{Authority: def.GroupBy.Authority, Name: def.GroupBy.Name, Value: def.GroupBy.Value}
	}

	return &pbDef
}

func PbToEntityAttrSets(pbsets map[string]*pbPDP.ListOfAttributeInstances) map[string][]attrs.AttributeInstance {
	entitySets := make(map[string][]attrs.AttributeInstance)

	for entity, instances := range pbsets {

		var convAttrs []attrs.AttributeInstance
		if instances != nil {
			convAttrs = PbToAttributeInstances(instances.AttributeInstances)
		}

		entitySets[entity] = convAttrs
	}
	return entitySets
}

func ValueFailureToPb(failure *pdp.ValueFailure) *pbPDP.ValueFailure {
	pbFail := pbPDP.ValueFailure{
		DataAttribute: AttributeInstanceToPb(failure.DataAttribute),
		Message:       failure.Message,
	}

	return &pbFail
}

func DataRuleResultsToPb(results []pdp.DataRuleResult) []*pbPDP.DataRuleResult {
	var pbresults []*pbPDP.DataRuleResult

	for _, v := range results {
		var convFails []*pbPDP.ValueFailure
		for _, fail := range v.ValueFailures {
			convFails = append(convFails, ValueFailureToPb(&fail))
		}
		pbresults = append(pbresults, &pbPDP.DataRuleResult{Passed: v.Passed, RuleDefinition: AttributeDefinitionToPb(v.RuleDefinition), ValueFailures: convFails})
	}

	return pbresults
}

func DecisionToPbResponse(entity string, decision *pdp.Decision) *pbPDP.DetermineAccessResponse {

	pbResults := DataRuleResultsToPb(decision.Results)

	return &pbPDP.DetermineAccessResponse{
		Entity:  entity,
		Access:  decision.Access,
		Results: pbResults,
	}
}
