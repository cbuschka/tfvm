package tfvm

func printVersion() {
	buildInfo := GetBuildInfo()
	Print("Tfvm version is %s, commitish=%s, built on %s.", buildInfo.version, buildInfo.commitish, buildInfo.buildTime)
}
