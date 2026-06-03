module myapp

go 1.25.4

require github.com/buildkite/buildkite-sdk/sdk/go v0.10.0

require github.com/itchyny/json2yaml v0.1.5 // indirect

replace github.com/buildkite/buildkite-sdk/sdk/go => ../../sdk/go
