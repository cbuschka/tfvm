package version

type Collection []*TerraformVersion

func (v Collection) Len() int {
	return len(v)
}

func (v Collection) Less(i, j int) bool {
	return v[i].Version.LessThan(v[j].Version)
}

func (v Collection) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}
