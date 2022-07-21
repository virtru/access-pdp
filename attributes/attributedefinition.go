package attributes

import (
	"fmt"
)

// AttributeDefinition describes metadata about the attribute -
// it's name, it's authority, it's rule, it's valid values, etc.
//
// Instances, not Definitions, are compared for access decisions.
//
// An AttributeDefinition is not "an attribute" and cannot be used for access decisions,
// it simply described how a given AttributeInstance should be compared.
//
// Every Instance has a parent Definition, but not every Definition has an Instance.
type AttributeDefinition struct {
	Authority string `json:"authority"`
	Name      string `json:"name"`
	Rule      string `json:"rule"`
	State     string `json:"state,omitempty"`
	//'order' contains all the valid values an Instance of this Definition may
	//have. If the `rule` is == hierarchy, then the ordering of these values implies
	//their hierarchical position.
	Order   []string           `json:"order"`
	GroupBy *AttributeInstance `json:"group_by,omitempty"`
}

// Returns the canonical URI representation of this AttributeDefinition:
//  <scheme>://<hostname>/attr/<name>
func (attrdef AttributeDefinition) GetCanonicalName() string {
	return fmt.Sprintf("%s/attr/%s",
		attrdef.Authority,
		attrdef.Name,
	)
}

// Returns the authority of this AttributeDefinition:
//  <scheme>://<hostname>
func (attrdef AttributeDefinition) GetAuthority() string {
	return attrdef.Authority
}
