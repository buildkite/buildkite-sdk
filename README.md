# buildkite-sdk

[![Build status](https://badge.buildkite.com/a95a3beece2339d1783a0a819f4ceb323c1eb12fb9662be274.svg?branch=main)](https://buildkite.com/buildkite/buildkite-sdk)

A multi-language SDK for [Buildkite](https://buildkite.com)! 🪁

Consumes the [Buildkite pipeline schema](https://github.com/buildkite/pipeline-schema) and generates and publishes packages for TypeScript or JavaScript, Python, Go, and Ruby.

## Installing and using the SDKs

Learn more about how to set up the Buildkite SDK for each langauge, and use it to work with your Buildkite pipelines, from the [Buildkite SDK](http://buildkite.com/docs/pipelines/configure/dynamic-pipelines/sdk) page of the Buildkite Docs.

## Upgrading to v0.4.0

In v0.4.0 we introduced type generation from Buildkite's [Pipeline Schema](https://github.com/buildkite/pipeline-schema). You can find a list of breaking changes [here](./docs/v0.0.4-breaking-changes.md).

## Development

### Prerequisites

To work on the SDK, you'll need current versions of the following tools:

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

# Publish the docs to AWS.
npm run docs:publish

# Clear away build and test artifacts.
npm run clean
```

### Regenerating types after a schema change

This SDK generates types from the [Buildkite pipeline schema](https://github.com/buildkite/pipeline-schema). When changes are made to the pipeline-schema repository, you can regenerate the types by running:

```bash
# Regenerate the types for all languages.
npm run types

# Regenerate the types for a specific language.
npm run types-ts
npm run types-py
npm run types-go
```

The type generator automatically fetches the latest schema from the `main` branch of the pipeline-schema repository. Generated types are then written to:

-   `sdk/typescript/src/types/`
-   `sdk/python/src/buildkite_sdk/schema.py`
-   `sdk/go/sdk/buildkite/`

Note that the type-generator binary (a Go program at `internal/gen/type-gen`) is automatically built when you run `npm run types`. If you need to rebuild that binary manually, run `npx nx gen:build`.

### Upgrading nx

We manage this repository with [Nx](https://nx.dev/). To upgrade the Nx workspace to the latest version, use `nx migrate`. From the root of the project, run:

```bash
npx nx migrate latest
```

See the [nx guide](https://nx.dev/features/automate-updating-dependencies) for details.

## Publishing new versions

Each SDK versions independently. Releasing one leaves the others untouched, and the pipeline publishes a version only if that SDK's registry does not already have it, so it is safe to re-run.

1.  Commit all pending changes. We want the release commit to be "clean" (i.e., to consist only of changes related to the release itself.)

1.  Release the SDKs you are shipping, with the bump each one needs:

    ```bash
    npx nx release patch --projects=sdk-python --skip-publish
    npx nx release minor --projects=sdk-typescript,sdk-go --skip-publish
    ```

    `--skip-publish` matters: publishing is the release pipeline's job, and
    without it `nx release` would push packages straight from your machine.

    For each project this:

    -   Bumps its version file (Go has none; its version is the tag)
    -   Writes its `CHANGELOG.md`
    -   Commits, and tags as `sdk/<language>/v<version>`

    Add `--dry-run` to preview without changing anything.

1. Push the commits and tags, then manually trigger the SDK Release Pipeline in Buildkite. For each SDK it takes the newest `sdk/<language>/v*` tag reachable from the build's commit and publishes it only if that version is missing from the registry, so SDKs you did not release are marked skipped rather than green, and re-running the pipeline is harmless. If no release tags are reachable at all it fails rather than quietly publishing nothing, which usually means tags were not fetched. After it has finished, create a release in GitHub ([example](https://github.com/buildkite/buildkite-sdk/releases/tag/v0.5.0)).

### Version sources

| SDK        | Version lives in                                     |
| ---------- | ---------------------------------------------------- |
| TypeScript | `sdk/typescript/package.json`                        |
| Python     | `sdk/python/pyproject.toml`                          |
| Ruby       | `sdk/ruby/lib/buildkite/version.rb`                  |
| C#         | `sdk/csharp/src/Buildkite.Sdk/Buildkite.Sdk.csproj`  |
| Go         | the `sdk/go/v*` git tag, no file                     |

TypeScript uses Nx's built-in npm support. The other four are handled by [`tools/release/version-actions.ts`](./tools/release/version-actions.ts), wired up per SDK in the `release` block of [`nx.json`](./nx.json).

### Docs

The SDK language docs are managed by a Pulumi Program in `infra` and manually deployed after every release.

### Required environment variables

The following environment variables are required for releasing and publishing:

-   `NPM_TOKEN` for publishing to npm (with `npm publish`)
-   `PYPI_TOKEN` for publishing to PyPI (with `uv publish`)
-   `GEM_HOST_API_KEY` for publishing to RubyGems (with `gem push`)
-   `NUGET_API_KEY` for publishing to NuGet (with `dotnet nuget push`)

See the `publish:all` tasks in `./project.json` for details.
