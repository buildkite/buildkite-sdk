# buildkite-sdk

[![Build status](https://badge.buildkite.com/a95a3beece2339d1783a0a819f4ceb323c1eb12fb9662be274.svg?branch=main)](https://buildkite.com/buildkite/pipeline-sdk)

A multi-language SDK for [Buildkite](https://buildkite.com)! 🪁 Consumes the [Buildkite pipeline schema](https://github.com/buildkite/pipeline-schema) and generates and publishes packages for TypeScript, Python, Go, and Ruby.

## Prerequisites

For development, you'll need current versions of the following tools:

-   [Node.js](https://nodejs.org/en/download), [Python](https://www.python.org/downloads/), [Go](https://go.dev/doc/install), [Ruby](https://www.ruby-lang.org/en/documentation/installation/)
-   For Python: [uv](https://docs.astral.sh/uv/), [Black](https://black.readthedocs.io/en/stable/)
-   For Ruby: [Bundler](https://bundler.io/)

See `mise.toml` for details. (We also recommend [Mise](https://mise.jdx.dev/) for tool-version management.) If you're on a Mac, and you use [Homebrew](https://brew.sh/), you can run `brew bundle` and `mise install` to get all you need:

```bash
brew bundle
mise install
```

If you hit any rough edges during development, please file an issue. Thanks!

### Useful commands

```bash
# Install all project dependencies.
npm install

# Test all SDKs and apps.
npm test

# Build all SDKs (and write them to ./dist/sdks).
npm run build

# Build all SDK docs (and write them to ./dist/docs).
npm run docs

# Serve the docs locally (which builds them implicitly).
npm run docs:serve

# Run all apps (which writes JSON and YAML pipelines to ./out).
npm run apps

# Watch all projects for changes (which rebuilds the docs and SDKs and re-runs all apps).
npm run watch

# Launch web servers for all docsets and watch all projects for changes. (Requires reload.)
npm run dev

# Format all SDK code.
npm run format

# Publish to npm, PyPi pkg.go.dev, and RubyGems.
npm run publish

# Clear away build and test artifacts.
npm run clean
```

## Installing and using the SDKs

The easiest way to use the SDK is to install the appropriate package for your language of choice, import the library into your program, assemble your pipeline steps programmatically, and serialize the pipeline to JSON or YAML, passing the output to `buildkite-agent pipeline upload`.

For example, if your language of choice were Ruby:

```bash
gem install buildkite-sdk
```

```ruby
# In ~/.buildkite/pipeline.rb:
require "buildkite"

pipeline = Buildkite::Pipeline.new

pipeline.add_step(
  label: "some-label",
  command: "echo 'Hello, World!'"
)

puts pipeline.to_json
```

```yaml
# In your pipeline's Settings > Steps:
steps:
    - label: ":pipeline: Generate pipeline"
      command: ruby .buildkite/pipeline.rb | buildkite-agent pipeline upload
```

This repository [uses this approach](./.buildkite/pipeline.rb) to ship the Buildkite SDK ... with the Buildkite SDK!

See below for more examples.

### Node.js

Install the package:

```bash
npm install @buildkite/buildkite-sdk
```

Use it in your program:

```javascript
const { Pipeline } = require("@buildkite/buildkite-sdk");

const pipeline = new Pipeline();

pipeline.addStep({
    command: "echo 'Hello, world!'",
});

console.log(pipeline.toJSON());
console.log(pipeline.toYAML());
```

### Python

Install the package:

```bash
uv add buildkite-sdk
```

Use it in your program:

```python
from buildkite_sdk import Pipeline

pipeline = Pipeline()
pipeline.add_step({"command": "echo 'Hello, world!'"})

print(pipeline.to_json())
print(pipeline.to_yaml())
```

### Go

Install the package:

```bash
go get github.com/buildkite/buildkite-sdk/sdk/go
```

Use it in your program:

```go
package main

import (
	"fmt"
	"github.com/buildkite/buildkite-sdk/sdk/go/sdk/buildkite"
)

func main() {
	pipeline := buildkite.Pipeline{}
	command := "echo 'Hello, world!"

	pipeline.AddCommandStep(buildkite.CommandStep{
		Command: &buildkite.CommandUnion{
			String: &command,
		},
	})

	fmt.Println(pipeline.ToJSON())
	fmt.Println(pipeline.ToYAML())
}
```

### Ruby

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

## Publishing new versions

All SDKs version on the same cadence. To publish a new version (of all SDKs), follow these steps:

1.  Commit all pending changes. We want the release commit to be "clean" (i.e., to consist only of changes related to the release itself.)

1.  Update the `VERSION_FROM` and `VERSION_TO` values in the `release:all` task in [`./project.json`](./project.json).

1.  Leaving that single change uncommitted, make sure you've exported a GitHub access token (as `GITHUB_TOKEN` -- see below) with push access to `main` branch of the repository, then run the release script:

    ```bash
    npm run release
    ```

    This script:

    -   Updates the version numbers in all affected files
    -   Rebuilds all SDKs
    -   Commits all changes (e.g., to version files, lockfiles, and anything else under `./sdk`)
    -   Adds two new tags to mark the release (`v0.0.0` and `sdk/go/v0.0.0`)
    -   Pushes the commit and tags to GitHub, triggering the `publish` task
    -   Creates a new GitHub release

    If for some reason the Buildkite publish job doesn't finish successfully, you can run some or all publish tasks from your local machine by exporting the applicable environment variables (again, see below), then running:

    ```bash
    npm run clean
    npm run build
    npm run publish                # To publish all packages
    npx nx publish sdk/typescript  # To publish only the Node.js package
    npx nx publish sdk/python      # To publish only the Python package
    npx nx publish sdk/go          # To publish only the Go package
    npx nx publish sdk/ruby        # To publish only the Ruby package
    ```

1.  Once the `publish` job completes, verify the releases at their respective URLs:

    -   https://github.com/buildkite/buildkite-sdk/releases
    -   https://www.npmjs.com/package/@buildkite/buildkite-sdk
    -   https://pypi.org/project/buildkite-sdk/
    -   https://pkg.go.dev/github.com/buildkite/buildkite-sdk/sdk/go (this usually takes a minute or two)
    -   https://rubygems.org/gems/buildkite-sdk

### Required environment variables

The following environment variables are required for releasing and publishing:

-   `GITHUB_TOKEN` for creating GitHub releases (with `@octokit/rest`)
-   `NPM_TOKEN` for publishing to npm (with `npm publish`)
-   `PYPI_TOKEN` fror publishing to PyPI (with `uv publish`)
-   `GEM_HOST_API_KEY` for publishing to RubyGems (with `gem push`)

See the `publish:all` and `release:all` tasks in `./project.json` for details.

## Upgrading nx

To upgrade the nx workspace to the latest version, use `nx migrate`. From the root of the project, run:

```bash
npx nx migrate latest
```

See the [nx guide](https://nx.dev/features/automate-updating-dependencies) for details.
