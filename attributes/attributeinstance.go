package attributes

import (
	"fmt"
	"net/url"
	"strings"
)

// AttributeInstance is created by selecting the Authority, Name and a specific Value from
// an AttributeDefinition.
//
// An AttributeInstance is a single, unique attribute, with a single value.
//
// Applied to an entity, the AttributeInstance becomes an entity attribute.
// Applied to data, the AttributeInstance becomes a data attribute.
//
// When making an access decisions, these two kinds of AttributeInstances are compared with each other.
//
// Example AttributeInstance:
// https://derp.com/attr/Blob/value/Green ->
//  Authority = https://derp.com
//  Name = Blob
//  CanonicalName = Authority + Name https://derp.com/attr/Blob
//  Value = Green
type AttributeInstance struct {
	Authority string `json:"authority"`
	Name      string `json:"name"`
	Value     string `json:"value"`
}

// Implement the standard "stringify" interface
// and return a string in the canonical AttributeInstance format of
//  <authority>/attr/<name>/value/<value>
func (attr AttributeInstance) String() string {
	return fmt.Sprintf("%s/attr/%s/value/%s",
		attr.Authority,
		attr.Name,
		attr.Value,
	)
}

// For cases where just the canonical name of this AttributeInstance is required
// (e.g. <authority>/attr/<name> - the authority and name, but not the value):
//  <authority>/attr/<name>
func (attr AttributeInstance) GetCanonicalName() string {
	return fmt.Sprintf("%s/attr/%s",
		attr.Authority,
		attr.Name,
	)
}

func (attrdef AttributeInstance) GetAuthority() string {
	return attrdef.Authority
}

// Accepts a valid attribute instance URI (authority + name + value in the canonical
// format 'https://example.org/attr/MyAttrName/value/MyAttrValue') and returns an
// AttributeInstance.
//
// Strings that are not valid URLs will result in a parsing failure, and return an error.
func ParseInstanceFromURI(attributeURI string) (AttributeInstance, error) {

	parsedAttr, err := url.Parse(attributeURI)
	if err != nil {
		return AttributeInstance{}, err
	}

	// Needs to be absolute - that is, rooted with a scheme, and not relative.
	if !parsedAttr.IsAbs() {
		return AttributeInstance{}, fmt.Errorf("Could not parse attributeURI %s - is not an absolute URI", attributeURI)
	}

	pathParts := strings.Split(strings.Trim(parsedAttr.Path, "/"), "/")
	// If we don't end up with exactly 4 segments, e.g. `attr/MyAttrName/value/MyAttrValue` ->
	// then something is wrong, this is not a canonical attr representation and we need to return an error
	if len(pathParts) != 4 {
		return AttributeInstance{}, fmt.Errorf("Could not parse attributeURI %s - path %s is not in canonical format, parts were %s", attributeURI, parsedAttr.Path, pathParts)
	}

	authority := fmt.Sprintf("%s://%s", parsedAttr.Scheme, parsedAttr.Hostname()) // == https://example.org
	name := pathParts[1]                                                          // == MyAttrName
	value := pathParts[3]                                                         // == MyAttrValue

	return AttributeInstance{
		Authority: authority, //Just scheme://host of the attribute - that is, the authority
		Name:      name,
		Value:     value,
	}, nil
}

// Accepts attribute namespace, name and value strings, and returns an AttributeInstance
func ParseInstanceFromParts(namespace, name, value string) (AttributeInstance, error) {
	fmtAttr := fmt.Sprintf("%s/attr/%s/value/%s", namespace, name, value)
	return ParseInstanceFromURI(fmtAttr)
}
