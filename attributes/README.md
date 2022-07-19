# Attributes

ABAC systems rely on "attributes", and Virtru itself chooses to represent attributes as canonical URIs for convenience's sake, but the word "attribute" by itself is too generic for implementation use.

To avoid confusion this codebase uses the following terms for the following "parts" of an attribute:

Assuming an example attribute in the following URI form: https://derp.com/attr/Blob/value/Green ->

 - **Attribute Namespace** = `https://derp.com`
 - **Attribute Name** = `Blob`
 - **Attribute Canonical Name** = Attribute Namespace + Attribute Name = `https://derp.com/attr/Blob`
 - **Attribute Value** = `Green`
 - **Attribute Instance** = Attribute Namespace + Attribute Name + Attribute Specific value = `https://derp.com/attr/Blob/value/Green`
 - **Attribute Definition** = Metadata (rule type: allof/anyof/hierarchy, etc, allowed values) associated with a specific Attribute Canonical Name
 
**Attribute Instances** are basically just **Attribute Definitions**, but taken with a specific **Attribute Value**.
 
Every **Attribute Instance** necessarily has a corresponding **Attribute Definition**. Multiple **Attribute Instances** may share the same **Attribute Definition**.

## Example

If an **Attribute Definition** for the Canonical Name `https://derp.com/attr/Blob` defined the following allowed Attribute Values: `[Green, Red, Blue, Purple]` for that Canonical Name:


then there would be 4 possible unique **Attribute Instances** for that single **Attribute Definition**:

`https://derp.com/attr/Blob/value/Green`
`https://derp.com/attr/Blob/value/Red`
`https://derp.com/attr/Blob/value/Blue`
`https://derp.com/attr/Blob/value/Purple`


1. These **Attribute Instances** all share the same Canonical Name
1. These **Attribute Instances** all share the same **Attribute Definition**
1. These **Attribute Instances** may be mapped to either data or entities in an ABAC system

## Struct representations

- [AttributeDefinition](./attributedefinition.go)
- [AttributeInstance](./attributeinstance.go)
