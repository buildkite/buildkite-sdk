module myapp

go 1.25.4

require github.com/buildkite/buildkite-sdk/sdk/go v0.11.0

require (
	github.com/buildkite/conditional v1.0.0 // indirect
	github.com/dlclark/regexp2 v1.12.0 // indirect
	github.com/itchyny/json2yaml v0.1.5 // indirect
)

replace github.com/buildkite/buildkite-sdk/sdk/go => ../../sdk/go
