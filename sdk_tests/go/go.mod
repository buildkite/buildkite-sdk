module github.com/buildkite/pipeline-sdk/sdk_tests/go

go 1.23.2

require (
	github.com/buildkite/pipeline-sdk/sdk/go v0.0.0-20241023211635-76b9e09faa8e
	github.com/santhosh-tekuri/jsonschema/v6 v6.0.1
	github.com/stretchr/testify v1.9.0
)

replace github.com/buildkite/pipeline-sdk/sdk/go => ../../sdk/go

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
