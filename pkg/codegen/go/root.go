package go_code_gen

var rootFile = `package buildkite

var Environment = environment{}
var StepBuilder = &stepBuilder{}`

func newRootFile() string {
	return rootFile
}
