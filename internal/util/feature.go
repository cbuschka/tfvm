package util

type Features struct {
	AlphaEnabled bool
}

var DefaultFeatures = newDefaultFeatures()

func newDefaultFeatures() Features {
	return Features{AlphaEnabled: IsEnvVarTrue("TFVM_ALPHA_ENABLED")}
}

func PrintAlphaFeatureWaningIfEnabled(features Features) {
	if features.AlphaEnabled {
		Print("Alpha features ENABLED!")
	}
}
