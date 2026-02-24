# C# SDK Review — Linear Issues

---

## Issue 1: Add `[JsonDerivedType]` attributes to the `Field` abstract class

**Size:** 1
**Priority:** Urgent
**Blockers:** None

### Description

`Field` is an abstract base class for `TextField` and `SelectField`. It's used as `List<Field>` in `BlockStep.Fields` and `InputStep.Fields`.

Without `[JsonDerivedType]` attributes, `System.Text.Json` will only serialize the base-class properties (`Key`, `Required`, `Hint`) when it encounters a `TextField` or `SelectField` in the list. Derived properties like `TextField.Text`, `SelectField.Select`, and `SelectField.Options` will be silently dropped from JSON output. This means any block or input step with fields will produce broken pipelines.

**File:** `sdk/csharp/src/Buildkite.Sdk/Schema/BlockStep.cs`, line 60

The fix is to add two attributes to the `Field` class declaration:

```csharp
[JsonDerivedType(typeof(TextField))]
[JsonDerivedType(typeof(SelectField))]
public abstract class Field { ... }
```

This requires adding `using System.Text.Json.Serialization;` to the top of the file (it's not currently imported since no other type in that file uses serialization attributes).

### Acceptance Criteria

- [ ] `Field` class has `[JsonDerivedType(typeof(TextField))]` and `[JsonDerivedType(typeof(SelectField))]` attributes
- [ ] `using System.Text.Json.Serialization;` is added to `BlockStep.cs`
- [ ] Add a test that creates a `BlockStep` with a `TextField` and a `SelectField`, serializes to JSON via `Pipeline.ToJson()`, and asserts that `text`, `select`, and `options` keys are present in the output

---

## Issue 2: Make `WaitStep.Type` read-only to match other step types

**Size:** 1
**Priority:** Urgent
**Blockers:** None

### Description

`WaitStep.Type` is declared as a mutable property with a default value:

```csharp
public string? Type { get; set; } = "wait";
```

All other step types use read-only expression-bodied members:

```csharp
// BlockStep.cs
public string Type => "block";

// InputStep.cs
public string Type => "input";

// TriggerStep.cs
public string Type => "trigger";
```

A user could accidentally (or intentionally) set `WaitStep.Type` to any string, producing an invalid pipeline. This should be a read-only property for consistency and correctness.

**File:** `sdk/csharp/src/Buildkite.Sdk/Schema/WaitStep.cs`, line 42

### Acceptance Criteria

- [ ] `WaitStep.Type` is changed from `public string? Type { get; set; } = "wait";` to `public string Type => "wait";`
- [ ] Existing `WaitStep` tests still pass

---

## Issue 3: Add C# to top-level Nx aggregate targets (`build:all`, `test:all`, etc.)

**Size:** 2
**Priority:** Urgent
**Blockers:** None

### Description

The root `project.json` defines aggregate Nx targets that run across all SDK languages. C# is currently absent from all of them:

- `install:all` (line 3–14)
- `clean:all` (line 16–37)
- `format:all` (line 38–50) — needs a `format` target added to `sdk/csharp/project.json` first
- `build:all` (line 143–149)
- `test:all` (line 151–161)
- `run:all` (line 183–189)
- `publish:all` (line 191–197)
- `docs:all` (line 162–168)

Without this, running `nx build:all` or `nx test:all` skips C# entirely. This also means any CI steps that use aggregate targets won't include C#.

**File:** `project.json` (root)

Additionally, `sdk/csharp/project.json` is missing a `format` and `clean` target. Other SDKs have these. A `format` target should run `dotnet format` and a `clean` target should remove `bin`, `obj`, and `dist` directories.

### Acceptance Criteria

- [ ] `sdk-csharp:install` added to `install:all`
- [ ] `sdk-csharp:build` added to `build:all`
- [ ] `sdk-csharp:test` added to `test:all`
- [ ] `sdk-csharp:publish` added to `publish:all`
- [ ] `app-csharp:clean` added to `clean:all`
- [ ] `sdk-csharp:format` added to `format:all` (after adding the target)
- [ ] `app-csharp:run` added to `run:all`
- [ ] `sdk-csharp:docs:build` added to `docs:all`
- [ ] `sdk/csharp/project.json` has `format` and `clean` targets
- [ ] Running `nx build:all` includes C# without errors

---

## Issue 4: Replace custom `SnakeCaseNamingPolicy` with built-in `JsonNamingPolicy.SnakeCaseLower`

**Size:** 1
**Priority:** High
**Blockers:** None

### Description

`Pipeline.cs` contains a hand-rolled `SnakeCaseNamingPolicy` class (lines 9–33) that converts C# PascalCase property names to snake_case for JSON serialization. The implementation inserts an underscore before every uppercase character, which breaks on consecutive uppercase letters:

- `"HTMLParser"` becomes `"h_t_m_l_parser"` instead of `"html_parser"`
- `"IOStream"` becomes `"i_o_stream"` instead of `"io_stream"`

No current properties trigger this bug, but it's a latent defect.

Since the library already targets `net8.0` as its minimum, the built-in `JsonNamingPolicy.SnakeCaseLower` (added in .NET 8) can be used as a drop-in replacement.

**File:** `sdk/csharp/src/Buildkite.Sdk/Pipeline.cs`, lines 9–42

### Acceptance Criteria

- [ ] The `SnakeCaseNamingPolicy` class is removed
- [ ] `JsonOptions` uses `PropertyNamingPolicy = JsonNamingPolicy.SnakeCaseLower`
- [ ] All existing tests pass
- [ ] The `System.Text.Json` import remains (it's needed for `JsonSerializer`)

---

## Issue 5: Decide on wrapper types (`DependsOn`, `SoftFail`, `Skip`) — fix or remove

**Size:** 3
**Priority:** High
**Blockers:** None

### Description

The SDK defines three wrapper types in `CommonTypes.cs` that model Buildkite's union-type fields:

- `DependsOn` — wraps `string`, `List<string>`, or `List<Dependency>` (lines 9–23)
- `SoftFail` — wraps `bool` or `List<SoftFailCondition>` (lines 85–97)
- `Skip` — wraps `bool` or `string` (lines 111–124)

Each stores the inner value in a private `object _value` field, exposed via a `Value` property. They also define implicit conversion operators (e.g., `implicit operator DependsOn(string key)`).

**The problem:** The step properties that accept these values are typed as `object?`, not as the wrapper type:

```csharp
// CommandStep.cs
public object? DependsOn { get; set; }
public object? Skip { get; set; }
public object? SoftFail { get; set; }
```

If a user assigns a wrapper type instance to one of these `object?` properties, `System.Text.Json` will serialize the wrapper object itself (including its `Value` property), producing output like `{"value": "step-key"}` instead of just `"step-key"`. There are no custom `JsonConverter`s to handle this.

The tests only ever assign raw values (e.g., `DependsOn = "build"`), so this bug path is untested.

**Two valid approaches:**

**Option A — Remove wrapper types:** Delete `DependsOn`, `SoftFail`, and `Skip` classes. Keep the properties as `object?` and document the valid types in XML doc comments (already partially done). This is simpler but less type-safe.

**Option B — Make wrapper types work:** Change the step properties to use the wrapper types (e.g., `public DependsOn? DependsOn { get; set; }`), and implement `JsonConverter<DependsOn>` (and the others) that serializes the inner `_value` directly. Also implement YamlDotNet `IYamlTypeConverter` equivalents. This gives users a proper typed API.

**File:** `sdk/csharp/src/Buildkite.Sdk/Schema/CommonTypes.cs`

### Acceptance Criteria

- [ ] A decision is documented (option A or B)
- [ ] If Option A: wrapper types are removed, doc comments on `object?` properties clearly state valid types
- [ ] If Option B: step properties use wrapper types, `JsonConverter`s are implemented, serialization tests pass for all variants (e.g., `DependsOn` from string, from string array, from `Dependency` array)
- [ ] Tests cover whichever approach is chosen

---

## Issue 6: Add test coverage for `InputStep`

**Size:** 2
**Priority:** High
**Blockers:** None

### Description

`InputStep` is the only step type with zero test coverage. All other step types have dedicated test files.

`InputStep` is structurally similar to `BlockStep` (it has `Fields`, `Prompt`, `AllowedTeams`, `BlockedState`) but uses the `input` discriminator key instead of `block`.

**File to create:** `sdk/csharp/tests/Buildkite.Sdk.Tests/InputStepTests.cs`

Use the existing test files as a pattern reference (e.g., `BlockStepTests.cs`). Each test should create a `Pipeline`, add an `InputStep`, call `pipeline.ToYaml()`, and assert the expected keys/values are present.

### Acceptance Criteria

- [ ] New file `InputStepTests.cs` exists
- [ ] Tests cover:
  - Basic properties (`Input` label, `Prompt`)
  - `Key` configuration
  - `TextField` in `Fields`
  - `SelectField` with options in `Fields`
  - `BlockedState` configuration
  - Conditional (`If`)
  - `DependsOn`
- [ ] All tests pass

---

## Issue 7: Add test coverage for notification types

**Size:** 2
**Priority:** High
**Blockers:** None

### Description

All seven notification types (`EmailNotification`, `SlackNotification`, `WebhookNotification`, `PagerDutyNotification`, `BasecampNotification`, `GitHubCommitStatusNotification`, `GitHubCheckNotification`) have zero test coverage.

Notifications can appear in two places:

1. Pipeline-level: `Pipeline.AddNotify(notification)`
2. Step-level: `CommandStep.Notify` and `GroupStep.Notify`

Both contexts should be tested.

**File to create:** `sdk/csharp/tests/Buildkite.Sdk.Tests/NotificationTests.cs`

**File:** `sdk/csharp/src/Buildkite.Sdk/Schema/Notifications.cs` (reference for the types)

### Acceptance Criteria

- [ ] New file `NotificationTests.cs` exists
- [ ] Tests cover pipeline-level notification for each type:
  - `EmailNotification` with `Email` and `If`
  - `SlackNotification` with string channel and with `SlackConfig` object
  - `WebhookNotification` with `Webhook`
  - `PagerDutyNotification` with `PagerdutyChangeEvent`
  - `BasecampNotification` with `BasecampCampfire`
  - `GitHubCommitStatusNotification` with `GithubCommitStatus` config
  - `GitHubCheckNotification` with `GithubCheck` config
- [ ] At least one test covers step-level notification (`CommandStep.Notify`)
- [ ] All tests pass

---

## Issue 8: Add test coverage for `Retry`, `Matrix`, `Cache`, and `Signature` types

**Size:** 2
**Priority:** Medium
**Blockers:** None

### Description

The `CommonTypes.cs` file defines several complex types that are used by `CommandStep` but have no serialization tests:

- `Retry` (with `AutomaticRetry` and `ManualRetry`)
- `Matrix` (with `MatrixAdjustment`)
- `Cache`
- `Signature`

These types use nested objects and should be tested to verify they serialize correctly to both JSON and YAML with proper snake_case naming.

**Files:**
- `sdk/csharp/src/Buildkite.Sdk/Schema/CommonTypes.cs` (type definitions)
- `sdk/csharp/tests/Buildkite.Sdk.Tests/CommandStepTests.cs` (add tests here)

### Acceptance Criteria

- [ ] Tests added (to `CommandStepTests.cs` or a new file) covering:
  - `Retry` with `AutomaticRetry` (exit status, limit)
  - `Retry` with `ManualRetry` (allowed, permit on passed, reason)
  - `Cache` with name, paths, and size
  - `Signature` with algorithm, signed fields, and value
- [ ] Serialized YAML output uses snake_case (e.g., `exit_status`, `permit_on_passed`, `signed_fields`)
- [ ] All tests pass

---

## Issue 9: Reduce `object?` usage on `CommandStep` properties with typed alternatives

**Size:** 5
**Priority:** Medium
**Blockers:** Issue 5 (decides the wrapper type pattern)

### Description

`CommandStep` has 11 properties typed as `object?`, which means users get no IntelliSense, no compile-time checking, and must consult documentation to know what values are valid. This is the biggest ergonomic gap compared to the Go, TypeScript, and Python SDKs.

The properties and their actual valid types:

| Property | Valid types |
|---|---|
| `Command` | `string` or `List<string>` |
| `Commands` | `string` or `List<string>` |
| `Agents` | `AgentsObject` or `AgentsList` |
| `Branches` | `string` or `List<string>` |
| `IfChanged` | `string` or `List<string>` |
| `DependsOn` | `string`, `string[]`, or `Dependency[]` |
| `Skip` | `bool` or `string` |
| `SoftFail` | `bool` or `SoftFailCondition[]` |
| `ArtifactPaths` | `string` or `List<string>` |
| `Plugins` | `object[]` (strings or `Dictionary<string, object>`) |
| `Matrix` | `List<string>` or `Matrix` |
| `Cache` | `string`, `List<string>`, or `Cache` |
| `Secrets` | `string[]` or `Dictionary<string, string>` |

The same pattern applies to other step types (`BlockStep.Branches`, `BlockStep.DependsOn`, `WaitStep.DependsOn`, `GroupStep.Skip`, etc.).

**Approach:** For each union-type property, create a wrapper class (similar to the existing `DependsOn`/`SoftFail`/`Skip` pattern) with:
- Static factory methods for each variant
- Implicit conversion operators for common types
- A custom `JsonConverter` that serializes the inner value directly
- A YamlDotNet type converter equivalent

Start with the most common properties: `Command`, `DependsOn`, `Skip`, `SoftFail`, `Branches`.

**Files:**
- `sdk/csharp/src/Buildkite.Sdk/Schema/CommonTypes.cs` (add/modify wrapper types)
- `sdk/csharp/src/Buildkite.Sdk/Schema/CommandStep.cs` (change property types)
- All other step type files that share these properties

### Acceptance Criteria

- [ ] At minimum, `Command`, `DependsOn`, `Skip`, `SoftFail`, and `Branches` have typed wrapper classes
- [ ] Each wrapper has implicit operators so that simple assignment still works (e.g., `Command = "echo hello"`)
- [ ] Custom `JsonConverter`s serialize the inner value, not the wrapper object
- [ ] YAML serialization also works correctly
- [ ] All existing tests still pass (implicit operators preserve backwards compatibility)
- [ ] New tests cover each variant of each wrapper type

---

## Issue 10: Suppress redundant `type` field serialization on `BlockStep`, `InputStep`, `TriggerStep`

**Size:** 1
**Priority:** Medium
**Blockers:** None

### Description

`BlockStep`, `InputStep`, and `TriggerStep` each have a read-only `Type` property that always returns a fixed string (`"block"`, `"input"`, `"trigger"`). Because these properties are non-null, they are always included in serialized output:

```yaml
- block: "Deploy?"
  type: block      # <-- redundant
```

Buildkite identifies these step types by their discriminator key (`block:`, `input:`, `trigger:`), not by a `type` field. The other SDKs (Go, TypeScript, Python) do not emit this field. It adds noise to the output.

**Files:**
- `sdk/csharp/src/Buildkite.Sdk/Schema/BlockStep.cs`, line 54
- `sdk/csharp/src/Buildkite.Sdk/Schema/InputStep.cs`, line 51
- `sdk/csharp/src/Buildkite.Sdk/Schema/TriggerStep.cs`, line 54
- `sdk/csharp/src/Buildkite.Sdk/Schema/WaitStep.cs` (after Issue 2 is resolved)

### Acceptance Criteria

- [ ] `Type` properties are annotated with `[JsonIgnore]` (from `System.Text.Json.Serialization`) and `[YamlIgnore]` (from `YamlDotNet.Serialization`)
- [ ] Or alternatively, the `Type` properties are removed entirely if they serve no purpose in the public API
- [ ] Serialized JSON and YAML output no longer includes `type: block`, `type: input`, `type: trigger`, or `type: wait`
- [ ] All existing tests still pass (update any assertions that check for the `type` field)

---

## Issue 11: Change `Pipeline.Env` and `CommandStep.Env` value type from `object?` to `string`

**Size:** 1
**Priority:** Medium
**Blockers:** None

### Description

Environment variables are string key-value pairs, but the SDK types them as `Dictionary<string, object?>`:

- `Pipeline.Env` (`Pipeline.cs:50`)
- `CommandStep.Env` (`CommandStep.cs:33`)
- `BuildkitePipeline.Env` (`BuildkitePipeline.cs:12`)
- `TriggerBuild.Env` (`TriggerStep.cs:72`)

Using `object?` as the value type allows users to insert non-string values (integers, booleans, complex objects) that may serialize in unexpected ways. All Buildkite environment variables are strings.

**Files:**
- `sdk/csharp/src/Buildkite.Sdk/Pipeline.cs`
- `sdk/csharp/src/Buildkite.Sdk/Schema/CommandStep.cs`
- `sdk/csharp/src/Buildkite.Sdk/Schema/BuildkitePipeline.cs`
- `sdk/csharp/src/Buildkite.Sdk/Schema/TriggerStep.cs`

### Acceptance Criteria

- [ ] All `Env` properties use `Dictionary<string, string>` instead of `Dictionary<string, object?>`
- [ ] `AddEnvironmentVariable` on `Pipeline` accepts `string value` instead of `object? value`
- [ ] All existing tests updated to use string values (most already do)
- [ ] All tests pass

---

## Issue 12: Handle alias properties to prevent double serialization

**Size:** 2
**Priority:** Medium
**Blockers:** None

### Description

Every step type has alias properties that map to the same Buildkite field:

- `Label` / `Name` (on all steps)
- `Key` / `Id` / `Identifier` (on all steps)
- `Command` / `Commands` (on `CommandStep`)
- `Block` / `Label` / `Name` (on `BlockStep`)
- `Wait` / `Label` / `Name` (on `WaitStep`)

If a user sets both `Label` and `Name`, both properties will be serialized, producing invalid or confusing output:

```yaml
label: "Build"
name: "Also Build"   # <-- both emitted, confusing
```

**Recommended approach:** Make aliases delegate to the primary property and suppress them from serialization:

```csharp
[JsonIgnore]
[YamlIgnore]
public string? Name
{
    get => Label;
    set => Label = value;
}
```

This way the user can set either property, but only the primary is serialized.

**Files:** All step type files in `sdk/csharp/src/Buildkite.Sdk/Schema/`

### Acceptance Criteria

- [ ] Alias properties delegate to the primary property (getter and setter)
- [ ] Alias properties are annotated with `[JsonIgnore]` and `[YamlIgnore]`
- [ ] Setting `Name` and reading `Label` (or vice versa) returns the same value
- [ ] Serialized output only contains the primary key (e.g., `label`, not `name`)
- [ ] Tests verify alias delegation works
- [ ] Existing tests still pass

---

## Issue 13: Multi-target the test project or run tests against both TFMs in CI

**Size:** 2
**Priority:** Medium
**Blockers:** None

### Description

The library targets both `net8.0` and `net9.0`:

```xml
<TargetFrameworks>net8.0;net9.0</TargetFrameworks>
```

But the test project only targets `net9.0`:

```xml
<TargetFramework>net9.0</TargetFramework>
```

This means tests never run against the `net8.0` build. Any API behavioral differences or compatibility issues between the two frameworks won't be caught.

**Files:**
- `sdk/csharp/tests/Buildkite.Sdk.Tests/Buildkite.Sdk.Tests.csproj`
- CI pipeline configuration (if separate dotnet versions need installing)

### Acceptance Criteria

- [ ] Test project targets both frameworks: `<TargetFrameworks>net8.0;net9.0</TargetFrameworks>`
- [ ] `dotnet test` runs tests against both TFMs (this happens automatically with multi-targeting)
- [ ] CI has both .NET 8 and .NET 9 SDKs available
- [ ] All tests pass on both frameworks

---

## Issue 14: Add `dotnet format` target to SDK project and integrate with `format:all`

**Size:** 1
**Priority:** Low
**Blockers:** Issue 3 (adding C# to aggregate targets)

### Description

The C# app project (`apps/csharp/project.json`) has a `format` target that runs `dotnet format`, but the SDK project (`sdk/csharp/project.json`) does not. The root `format:all` target also doesn't include C#.

Consistent formatting enforcement ensures contributions stay clean.

**Files:**
- `sdk/csharp/project.json` (add `format` target)
- `project.json` (root — add to `format:all`, covered by Issue 3)

### Acceptance Criteria

- [ ] `sdk/csharp/project.json` has a `format` target: `dotnet format`
- [ ] Running `nx format sdk-csharp` works
- [ ] `format:all` includes `sdk-csharp:format` (may be handled by Issue 3)

---

## Issue 15: Rename `TriggerBuild.MetaData` to `Metadata` for .NET naming conventions

**Size:** 1
**Priority:** Low
**Blockers:** None

### Description

`TriggerBuild.MetaData` uses mid-word capitalization (`MetaData`) which doesn't follow .NET naming conventions. The standard .NET spelling is `Metadata` (single compound word). Since the JSON/YAML serializer converts to snake_case, both spellings produce the same output (`meta_data`), so this is purely a code style issue.

**File:** `sdk/csharp/src/Buildkite.Sdk/Schema/TriggerStep.cs`, line 75

### Acceptance Criteria

- [ ] Property renamed from `MetaData` to `Metadata`
- [ ] All usages updated (if any exist in tests or examples)
- [ ] Serialized output still produces `meta_data`

---

## Issue 16: Add `EnvironmentVariable` tests

**Size:** 2
**Priority:** Low
**Blockers:** None

### Description

The `EnvironmentVariable` static class (`Environment.cs`) provides strongly-typed access to 50+ Buildkite environment variables. It has zero test coverage.

While the implementation is straightforward (each property just calls `System.Environment.GetEnvironmentVariable`), tests are valuable to:
- Verify the correct environment variable names are used (typos would be silent bugs)
- Verify the boolean properties (`AgentDebug`, `GitSubmodules`, `IsBuildkite`, `IsCi`) parse correctly
- Serve as documentation for available environment variables

**File to create:** `sdk/csharp/tests/Buildkite.Sdk.Tests/EnvironmentVariableTests.cs`

Testing approach: Set environment variables in the test, read them via `EnvironmentVariable.*`, assert the values match. Clean up in a `finally` block or use `IDisposable` test fixture.

### Acceptance Criteria

- [ ] New file `EnvironmentVariableTests.cs` exists
- [ ] Tests cover a representative sample (at least 10 properties):
  - String properties: `BuildId`, `Branch`, `Commit`, `PipelineSlug`
  - Boolean properties: `IsBuildkite`, `IsCi`, `AgentDebug`
  - Null case: property returns `null` when env var is not set
- [ ] Environment variables are cleaned up after each test (no test pollution)
- [ ] All tests pass
