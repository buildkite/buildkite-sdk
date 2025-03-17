# buildkite-sdk

[![Build status](https://badge.buildkite.com/a95a3beece2339d1783a0a819f4ceb323c1eb12fb9662be274.svg?branch=main)](https://buildkite.com/buildkite/pipeline-sdk)

A Ruby SDK for [Buildkite](https://buildkite.com)! 🪁

## Usage

Install the package:

```bash
gem install buildkite-sdk
```

Use it in your program:

```ruby
require "buildkite"

pipeline = Buildkite::Pipeline.new

pipeline.add_step(
  label: "some-label",
  command: "echo 'Hello, World!'"
)

puts pipeline.to_json
puts pipeline.to_yaml
```
