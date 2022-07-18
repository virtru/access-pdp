package attributes

// Clusterable is an interface that either AttributeInstances or AttributeDefinitions can implement,
// to support easily "clustering" or grouping a slice of either by their shared CanonicalName or Authority.
type Clusterable interface {
	// Type constraint (generics)
	// Both AttributeDefinitions and AttributeInstances are clusterable
	AttributeInstance | AttributeDefinition

	// Returns the canonical URI representation of this clusterable thing, in the format
	//  <scheme>://<hostname>/attr/<name>
	GetCanonicalName() string
	// Returns the authority of this clusterable thing, in the format
	//  <scheme>://<hostname>
	GetAuthority() string
}

// ClusterByAuthority takes a slice of Clusterables, and returns them as a map,
// where the map is keyed by each unique Authorities (e.g. 'https://myauthority.org') found in the slice of Clusterables
func ClusterByAuthority[attrCluster Clusterable](attrs []attrCluster) map[string][]attrCluster {
	clusters := make(map[string][]attrCluster)

	for _, instance := range attrs {
		clusters[instance.GetAuthority()] = append(clusters[instance.GetAuthority()], instance)
	}

	return clusters
}

// ClusterByCanonicalName takes a slice of Clusterables (AttributeInstance OR AttributeDefinition),
// and returns them as a map, where the map is keyed by each unique CanonicalName
// (e.g. Authority+Name, 'https://myauthority.org/attr/<name>') found in the slice of Clusterables
func ClusterByCanonicalName[attrCluster Clusterable](attrs []attrCluster) map[string][]attrCluster {
	clusters := make(map[string][]attrCluster)

	for _, instance := range attrs {
		clusters[instance.GetCanonicalName()] = append(clusters[instance.GetCanonicalName()], instance)
	}

	return clusters
}
