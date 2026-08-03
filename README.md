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

- [Node.js](https://nodejs.org/en/download), [Python](https://www.python.org/downloads/), [Go](https://go.dev/doc/install), [Ruby](https://www.ruby-lang.org/en/documentation/installation/)
- For Python: [uv](https://docs.astral.sh/uv/), [Black](https://black.readthedocs.io/en/stable/)
- For Ruby: [Bundler](https://bundler.io/)

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

- `sdk/typescript/src/types/`
- `sdk/python/src/buildkite_sdk/schema.py`
- `sdk/go/sdk/buildkite/`

Note that the type-generator binary (a Go program at `internal/gen/type-gen`) is automatically built when you run `npm run types`. If you need to rebuild that binary manually, run `npx nx gen:build`.

### Upgrading nx

We manage this repository with [Nx](https://nx.dev/). To upgrade the Nx workspace to the latest version, use `nx migrate`. From the root of the project, run:

```bash
npx nx migrate latest
```

See the [nx guide](https://nx.dev/features/automate-updating-dependencies) for details.

## Publishing new versions

Each SDK has its own version and `sdk/<language>/v<version>` tag. Always use
`release:create`; do not run `nx release` directly. The local command creates
the release commit and tags. Buildkite publishes the packages.

1.  Prepare the repository.

    Commit all intended changes, then confirm the working tree is clean:

    ```bash
    git status --short
    ```

    This command must produce no output. `release:create` also fetches tags and
    stops if a local tag conflicts with the remote.

1.  Preview the release and choose a bump.

    | Bump  | Preview                                        | Create                               |
    | ----- | ---------------------------------------------- | ------------------------------------ |
    | Patch | `npx nx release:create --dry-run`              | `npx nx release:create`              |
    | Minor | `npx nx release:create --dry-run --bump=minor` | `npx nx release:create --bump=minor` |
    | Major | `npx nx release:create --dry-run --bump=major` | `npx nx release:create --bump=major` |

    The helper selects every SDK changed since its own latest tag. One bump
    applies to every SDK selected by that run. Confirm the preview lists only
    the SDKs you intend to release before continuing.

    A change to `sdk/<language>/project.json` alone is ignored. To make a
    deliberate packaging release for such a change, add `--force` to both the
    preview and create commands.

1.  Create the release.

    Run exactly one command from the **Create** column above, using the same
    bump you previewed. For every selected SDK, the command:

    - Updates its version file (Go has no version file)
    - Updates its `CHANGELOG.md`
    - Creates one release commit
    - Creates an `sdk/<language>/v<version>` tag

    It does not push or publish anything.

1.  Review the release commit and tags.

    ```bash
    git show --stat
    git tag --points-at HEAD
    ```

    Confirm the commit contains only release changes and that there is one tag
    for each SDK shown in the preview. If anything is wrong, do not push it.

1.  Push the commit and tags.

    ```bash
    git push --follow-tags
    ```

1.  Publish from Buildkite.

    Get the release commit SHA:

    ```bash
    git rev-parse HEAD
    ```

    Manually trigger the **SDK Release Pipeline** against that exact commit.
    The pipeline checks each registry first, skips versions already published,
    and publishes only versions missing from their registry. A publish step
    turns red instead of publishing if its tag is missing or its SDK files do
    not match the tag. Other SDKs can continue independently.

1.  Create the GitHub Releases.

    After Buildkite succeeds, create one GitHub Release for each SDK published.
    Select its `sdk/<language>/v<version>` tag and use that SDK's new
    `CHANGELOG.md` entry as the release body.

### Version sources

| SDK        | Version lives in                                    |
| ---------- | --------------------------------------------------- |
| TypeScript | `sdk/typescript/package.json`                       |
| Python     | `sdk/python/pyproject.toml`                         |
| Ruby       | `sdk/ruby/lib/buildkite/version.rb`                 |
| C#         | `sdk/csharp/src/Buildkite.Sdk/Buildkite.Sdk.csproj` |
| Go         | the `sdk/go/v*` git tag, no file                    |

TypeScript uses Nx's built-in npm support. The other four are handled by [`tools/release/version-actions.ts`](./tools/release/version-actions.ts), wired up per SDK in the `release` block of [`nx.json`](./nx.json).

### Docs

The SDK language docs are managed by a Pulumi Program in `infra` and manually deployed after every release.

### Required environment variables

The local `release:create` command does not need registry credentials. The
Buildkite SDK Release Pipeline supplies these variables when publishing:

- `NPM_TOKEN` for publishing to npm (with `npm publish`)
- `PYPI_TOKEN` for publishing to PyPI (with `uv publish`)
- `GEM_HOST_API_KEY` for publishing to RubyGems (with `gem push`)
- `NUGET_API_KEY` for publishing to NuGet (with `dotnet nuget push`)

See the `publish:all` tasks in `./project.json` for details.
