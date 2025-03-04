const fs = require("fs");
const yaml = require("yaml");
const mustache = require("mustache");

const {
    quicktype,
    InputData,
    JSONSchemaInput,
    TypeScriptTargetLanguage,
    PythonTargetLanguage,
    GoTargetLanguage,
    RubyTargetLanguage,
} = require("quicktype-core");

async function genTypes() {
    const inputData = new InputData();

    const result = await fetch(
        "https://raw.githubusercontent.com/buildkite/pipeline-schema/refs/heads/main/schema.json"
    );

    const schema = await result.json();
    const source = {
        name: "schema",
        schema: JSON.stringify(schema),
    };

    await inputData.addSource(
        "schema",
        source,
        () => new JSONSchemaInput(undefined)
    );

    const typescript = new TypeScriptTargetLanguage(
        "TypeScript",
        ["typescript"],
        "ts"
    );
    const python = new PythonTargetLanguage("Python", ["python"], "py");
    const go = new GoTargetLanguage("Go", ["go"], "go");
    const ruby = new RubyTargetLanguage("Ruby", ["ruby"], "rb");

    const opts = {
        typescript: {
            path: "./sdk/typescript/src/schema.ts",
            options: undefined,
        },
        python: {
            path: "./sdk/python/src/buildkite_sdk/schema.py",
            options: undefined,
        },
        go: {
            path: "./sdk/go/sdk/buildkite/schema.go",
            options: { package: "buildkite", omitEmpty: true },
        },
        ruby: {
            path: "./sdk/ruby/lib/schema.rb",
            options: undefined,
        },
    };

    for await (lang of [typescript, python, go, ruby]) {
        const langOpts = opts[lang.name];
        let { lines } = await quicktype({
            lang,
            inputData,
            rendererOptions: langOpts.options,
        });

        // Go formatter seems a bit buggy.
        if (lang.name === "go") {
            lines = lines.map((line) => {
                if (line.match(/Command[s]?\W+\*Branches/)) {
                    return line.replace("*Branches    ", "*CommandUnion");
                }

                if (line.match(/ArtifactPaths\W+\*Branches/)) {
                    return line.replace("*Branches", "[]string ");
                }

                return line;
            });
        }

        fs.writeFileSync(langOpts.path, lines.join("\n"), "utf-8");
    }
}

async function genEnvVars() {
    const result = await fetch(
        "https://raw.githubusercontent.com/buildkite/docs/a26485aba7049a7dedea8ad2988667a6b9245bf8/data/content/environment_variables.yaml"
    );

    const text = await result.text();
    let { variables } = yaml.parse(text);
    const langs = ["typescript", "python", "go", "ruby"];

    const opts = {
        typescript: {
            path: "./sdk/typescript/src/environment.ts",
            options: undefined,
        },
        python: {
            path: "./sdk/python/src/buildkite_sdk/environment.py",
            options: undefined,
        },
        go: {
            path: "./sdk/go/sdk/buildkite/environment.go",
            options: { package: "buildkite" },
        },
        ruby: {
            path: "./sdk/ruby/lib/environment.rb",
            options: undefined,
        },
    };

    // Make the appropriate tweaks.
    variables = variables.map((v) => {
        // Trim trailing newlines.
        v.desc = v.desc.trim();

        // Not valid, so provide the prefix, as it's still useful for interpolation.
        if (v.name === "BUILDKITE_AGENT_META_DATA_*") {
            return {
                ...v,
                name: "BUILDKITE_AGENT_META_DATA_",
            };
        }
        return v;
    });

    // TypeScript.
    let sdkPath = "./sdk/typescript";
    let template = fs
        .readFileSync(`${sdkPath}/env.mustache`, "utf-8")
        .toString();
    let rendered = mustache.render(template, {
        variables: variables.map((v) => {
            return {
                ...v,
                comment: [
                    "/**",
                    ...v.desc.split("\n").map((line) => `* ${line}`),
                    " */",
                ].join("\n"),
            };
        }),
    });
    fs.writeFileSync(`${sdkPath}/src/environment.ts`, rendered, "utf-8");

    // Python.
    sdkPath = "./sdk/python";
    template = fs.readFileSync(`${sdkPath}/env.mustache`, "utf-8").toString();
    rendered = mustache.render(template, {
        variables: variables.map((v) => {
            return {
                ...v,
                comment: [
                    ...v.desc.split("\n").map((line) => `# ${line}`),
                ].join("\n"),
            };
        }),
    });
    fs.writeFileSync(
        `${sdkPath}/src/buildkite_sdk/environment.py`,
        rendered,
        "utf-8"
    );

    // Go.
    sdkPath = "./sdk/go";
    template = fs.readFileSync(`${sdkPath}/env.mustache`, "utf-8").toString();
    rendered = mustache.render(template, {
        package: "buildkite",
        variables: variables.map((v) => {
            return {
                ...v,
                comment: [
                    ...v.desc.split("\n").map((line) => `// ${line}`),
                ].join("\n"),
            };
        }),
    });
    fs.writeFileSync(
        `${sdkPath}/sdk/buildkite/environment.go`,
        rendered,
        "utf-8"
    );

    // Ruby.
    sdkPath = "./sdk/ruby";
    template = fs.readFileSync(`${sdkPath}/env.mustache`, "utf-8").toString();
    rendered = mustache.render(template, {
        variables: variables.map((v) => {
            return {
                ...v,
                comment: [
                    ...v.desc
                        .split("\n")
                        .map((line, i) => `${i === 0 ? "#" : "  #"} ${line}`),
                ].join("\n"),
            };
        }),
    });
    fs.writeFileSync(`${sdkPath}/lib/environment.rb`, rendered, "utf-8");
}

(async () => {
    await genTypes();
    await genEnvVars();
})();
