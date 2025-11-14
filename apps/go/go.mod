module myapp

go 1.25.4

require github.com/buildkite/buildkite-sdk/sdk/go v0.0.1

require github.com/itchyny/json2yaml v0.1.4 // indirect

replace github.com/buildkite/buildkite-sdk/sdk/go => ../../sdk/go
