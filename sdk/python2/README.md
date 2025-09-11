# buildkite-sdk

[![Build status](https://badge.buildkite.com/a95a3beece2339d1783a0a819f4ceb323c1eb12fb9662be274.svg?branch=main)](https://buildkite.com/buildkite/pipeline-sdk)

A Python SDK for [Buildkite](https://buildkite.com)! 🪁

## Usage

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
