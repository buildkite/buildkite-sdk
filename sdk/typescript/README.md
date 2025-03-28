# buildkite-sdk

[![Build status](https://badge.buildkite.com/a95a3beece2339d1783a0a819f4ceb323c1eb12fb9662be274.svg?branch=main)](https://buildkite.com/buildkite/pipeline-sdk)

A TypeScript SDK for [Buildkite](https://buildkite.com)! 🪁

## Usage

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

change
