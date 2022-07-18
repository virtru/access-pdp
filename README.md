# ABAC Access Policy Decision Point

A reference implementation of an Attribute Based Access Control (ABAC) Access Policy Decision Point (PDP)

## What's an Access PDP, and how does it fit into an ABAC system?

![ABAC System](./resources/index.png)

- A Access PDP (Policy Decision Point) is a library or service that *makes Access decisions*. It is usually "wrapped" or used by an Access Policy Enforcement Point (PEP), which *enforces* whatever decision this Access PDP makes.



## Details
In this implementation, the Access PDP

1. Expects to be provided with:
  - The data attributes to make a decision against
  - Attribute definitions for every data attribute the decision is being made against
  - A list of entities the decision is being made against, and entitlements (entity attributes) for each entity
  
  as decision input.
  
  To the Access PDP, an "entity" is just a string identifier of any kind with entity attributes attached to it - this PDP
  doesn't care about entity subtypes (PE, NPE) or what kind of entity identifiers are being used, and "entities" have no meaning to the PDP except as a way to group decision results - they are simply there so the PEP invoking this library can correlate PDP requests with PDP results. Entity identifiers and "types" only have meaning to the caller of the PDP, not the PDP itself.
  
2. Returns:
   - For each entity identifer provided:
    1. A single, top-level boolean Access property indicating the overall access decision for that entity against the complete set of provided Data Attributes, according to the rules of the Data Attribute Definitions those map to (any-of, all-of, hierarchy).
    2. A set of DataRuleResults for each Data Attribute comparison that was done, which contributed to the top-level Access property of true or false.

## Interface

This library exposes gRPC endpoints, and so can be consumed by any code that understands the gRPC protocol. This library could be wrapped in a container and hosted out-of-process from an Access PEP, or it could be hosted in-process.
