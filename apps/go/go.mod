module myapp

go 1.23.2

require github.com/buildkite/buildkite-sdk/sdk/go v0.0.1

require gopkg.in/yaml.v3 v3.0.1 // indirect

replace github.com/buildkite/buildkite-sdk/sdk/go => ../../sdk/go
