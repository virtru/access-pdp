package attributes

import (
	"fmt"
)

type AttributeDefinition struct {
	Authority string `json:"authority"`
	Name      string `json:"name"`
	Rule      string `json:"rule"`
	State     string `json:"state,omitempty"`
	// 'order' contains all the valid values an Instance of this Definition may
	// have. If the `rule` is == hierarchy, then the ordering of these values implies
	// their hierarchical position.
	Order   []string           `json:"order"`
	GroupBy *AttributeInstance `json:"group_by,omitempty"`
}

// An AttributeDefinition describes metadata about the attribute -
// it's name, it's authority, it's rule, it's valid values, etc.
//
// An AttributeDefinition is not "an attribute" and cannot be used for access decisions.
//
// An attributes.AttributeInstance is a single, unique attribute, with a single value.
//
// Every Instance has a parent Definition, but not every Definition has an Instance.
//
// Instances, not Definitions, are compared for access decisions.

// Returns the canonical URI representation of this attribute definition.
// <scheme>://<hostname>/attr/<name>
func (attrdef AttributeDefinition) GetCanonicalName() string {
	return fmt.Sprintf("%s/attr/%s",
		attrdef.Authority,
		attrdef.Name,
	)
}

func (attrdef AttributeDefinition) GetAuthority() string {
	return attrdef.Authority
}
