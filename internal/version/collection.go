package version

// Collection of terraform versions.
type Collection []*TerraformVersion

// Len gives length of collection.
func (v Collection) Len() int {
	return len(v)
}

// Less compares two terraform version items.
func (v Collection) Less(i, j int) bool {
	return v[i].Version.LessThan(v[j].Version)
}

// Swap swap two terraform version items.
func (v Collection) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}
