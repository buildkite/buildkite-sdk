package go_code_gen

var rootFile = `package buildkite

var Environment = environment{}
var PipelineBuilder = &StepBuilder{}`

func newRootFile() string {
	return rootFile
}
