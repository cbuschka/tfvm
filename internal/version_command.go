package tfvm

func printVersion() {
	buildInfo := GetBuildInfo()
	Print("%s", buildInfo.version)
}
