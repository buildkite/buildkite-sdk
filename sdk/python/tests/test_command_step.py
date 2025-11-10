from buildkite_sdk import (
    Pipeline,
    CommandStep,
    CommandStepArgs,
    NestedCommandStep,
    AutomaticRetry,
    CommandStepRetry,
    SoftFailObject,
    DependsOnListObject,
    CommandStepSignature,
)
from .utils import TestRunner


class TestCommandStepNestingTypesClass(TestRunner):
    def test_type_script(self):
        pipeline = Pipeline(steps=[CommandStep(type="script")])
        self.validator.check_result(pipeline, {"steps": [{"type": "script"}]})

    def test_command_field(self):
        pipeline = Pipeline(steps=[CommandStep(command="bash.sh")])
        self.validator.check_result(pipeline, {"steps": [{"command": "bash.sh"}]})

    def test_type_script_command_field(self):
        pipeline = Pipeline(steps=[CommandStep(type="script", command="bash.sh")])
        self.validator.check_result(
            pipeline, {"steps": [{"type": "script", "command": "bash.sh"}]}
        )

    def test_type_command_command_field(self):
        pipeline = Pipeline(steps=[CommandStep(type="command", command="bash.sh")])
        self.validator.check_result(
            pipeline, {"steps": [{"type": "command", "command": "bash.sh"}]}
        )

    def test_nested_command(self):
        pipeline = Pipeline(
            steps=[NestedCommandStep(command=CommandStep(command="bash.sh"))]
        )
        self.validator.check_result(
            pipeline, {"steps": [{"command": {"command": "bash.sh"}}]}
        )

    def test_nested_commands(self):
        pipeline = Pipeline(
            steps=[NestedCommandStep(commands=CommandStep(command="bash.sh"))]
        )
        self.validator.check_result(
            pipeline, {"steps": [{"commands": {"command": "bash.sh"}}]}
        )

    def test_nested_script(self):
        pipeline = Pipeline(
            steps=[NestedCommandStep(script=CommandStep(command="bash.sh"))]
        )
        self.validator.check_result(
            pipeline, {"steps": [{"script": {"command": "bash.sh"}}]}
        )

    def test_only_plugin(self):
        pipeline = Pipeline(
            steps=[CommandStep(plugins=[{"a-plugin#v1.0.0": {"run": "app"}}])]
        )
        self.validator.check_result(
            pipeline, {"steps": [{"plugins": [{"a-plugin#v1.0.0": {"run": "app"}}]}]}
        )


class TestCommandStepNestingTypesDict(TestRunner):
    def test_type_script(self):
        pipeline = Pipeline(steps=[CommandStep.from_dict({"type": "script"})])
        self.validator.check_result(pipeline, {"steps": [{"type": "script"}]})

    def test_command_field(self):
        pipeline = Pipeline(steps=[CommandStep.from_dict({"command": "bash.sh"})])
        self.validator.check_result(pipeline, {"steps": [{"command": "bash.sh"}]})

    def test_type_script_command_field(self):
        pipeline = Pipeline(
            steps=[CommandStep.from_dict({"type": "script", "command": "bash.sh"})]
        )
        self.validator.check_result(
            pipeline, {"steps": [{"type": "script", "command": "bash.sh"}]}
        )

    def test_type_command_command_field(self):
        pipeline = Pipeline(
            steps=[CommandStep.from_dict({"type": "command", "command": "bash.sh"})]
        )
        self.validator.check_result(
            pipeline, {"steps": [{"type": "command", "command": "bash.sh"}]}
        )

    def test_nested_command(self):
        pipeline = Pipeline(
            steps=[NestedCommandStep.from_dict({"command": {"command": "bash.sh"}})]
        )
        self.validator.check_result(
            pipeline, {"steps": [{"command": {"command": "bash.sh"}}]}
        )

    def test_nested_commands(self):
        pipeline = Pipeline(
            steps=[NestedCommandStep.from_dict({"commands": {"command": "bash.sh"}})]
        )
        self.validator.check_result(
            pipeline, {"steps": [{"commands": {"command": "bash.sh"}}]}
        )

    def test_nested_script(self):
        pipeline = Pipeline(
            steps=[NestedCommandStep.from_dict({"script": {"command": "bash.sh"}})]
        )
        self.validator.check_result(
            pipeline, {"steps": [{"script": {"command": "bash.sh"}}]}
        )

    def test_only_plugin(self):
        pipeline = Pipeline(
            steps=[
                CommandStep.from_dict(
                    {"plugins": [{"a-plugin#v1.0.0": {"run": "app"}}]}
                )
            ]
        )
        self.validator.check_result(
            pipeline, {"steps": [{"plugins": [{"a-plugin#v1.0.0": {"run": "app"}}]}]}
        )


class TestCommandStepClass(TestRunner):
    def test_agents_dict(self):
        pipeline = Pipeline(steps=[CommandStep(command="test", agents={"os": "macOS"})])
        self.validator.check_result(
            pipeline, {"steps": [{"command": "test", "agents": {"os": "macOS"}}]}
        )

    def test_agents_list(self):
        pipeline = Pipeline(steps=[CommandStep(command="test", agents=["os=macOS"])])
        self.validator.check_result(
            pipeline, {"steps": [{"command": "test", "agents": ["os=macOS"]}]}
        )

    def test_artifact_paths_list(self):
        pipeline = Pipeline(
            steps=[CommandStep(command="test", artifact_paths=["one", "two"])]
        )
        self.validator.check_result(
            pipeline, {"steps": [{"command": "test", "artifact_paths": ["one", "two"]}]}
        )

    def test_artifact_paths_string(self):
        pipeline = Pipeline(steps=[CommandStep(command="test", artifact_paths="path")])
        self.validator.check_result(
            pipeline, {"steps": [{"command": "test", "artifact_paths": "path"}]}
        )

    def test_branches(self):
        pipeline = Pipeline(
            steps=[CommandStep(command="test", branches="master deploy-*")]
        )
        self.validator.check_result(
            pipeline, {"steps": [{"command": "test", "branches": "master deploy-*"}]}
        )

    def test_concurrency(self):
        pipeline = Pipeline(
            steps=[
                CommandStep(
                    command="test",
                    concurrency=1,
                    concurrency_group="my-group",
                    concurrency_method="eager",
                )
            ]
        )
        self.validator.check_result(
            pipeline,
            {
                "steps": [
                    {
                        "command": "test",
                        "concurrency": 1,
                        "concurrency_group": "my-group",
                        "concurrency_method": "eager",
                    }
                ]
            },
        )

    def test_env(self):
        pipeline = Pipeline(steps=[CommandStep(command="test", env={"ENV": "bar"})])
        self.validator.check_result(
            pipeline, {"steps": [{"command": "test", "env": {"ENV": "bar"}}]}
        )

    def test_id(self):
        pipeline = Pipeline(steps=[CommandStep(command="test", id="id")])
        self.validator.check_result(
            pipeline, {"steps": [{"command": "test", "id": "id"}]}
        )

    def test_identifier(self):
        pipeline = Pipeline(
            steps=[CommandStep(command="test", identifier="identifier")]
        )
        self.validator.check_result(
            pipeline, {"steps": [{"command": "test", "identifier": "identifier"}]}
        )

    def test_label(self):
        pipeline = Pipeline(steps=[CommandStep(command="test", label="label")])
        self.validator.check_result(
            pipeline, {"steps": [{"command": "test", "label": "label"}]}
        )

    def test_parallelism(self):
        pipeline = Pipeline(steps=[CommandStep(command="test", parallelism=42)])
        self.validator.check_result(
            pipeline, {"steps": [{"command": "test", "parallelism": 42}]}
        )

    def test_plugins_dict(self):
        pipeline = Pipeline(
            steps=[
                CommandStep(command="test", plugins={"a-plugin#v1.0.0": {"run": "app"}})
            ]
        )
        self.validator.check_result(
            pipeline,
            {
                "steps": [
                    {"command": "test", "plugins": {"a-plugin#v1.0.0": {"run": "app"}}}
                ]
            },
        )

    def test_plugins_dict_list(self):
        pipeline = Pipeline(
            steps=[
                CommandStep(
                    command="test", plugins=[{"a-plugin#v1.0.0": {"run": "app"}}]
                )
            ]
        )
        self.validator.check_result(
            pipeline,
            {
                "steps": [
                    {
                        "command": "test",
                        "plugins": [{"a-plugin#v1.0.0": {"run": "app"}}],
                    }
                ]
            },
        )

    def test_plugins_string_list(self):
        pipeline = Pipeline(
            steps=[CommandStep(command="test", plugins=["a-plugin#v1.0.0"])]
        )
        self.validator.check_result(
            pipeline, {"steps": [{"command": "test", "plugins": ["a-plugin#v1.0.0"]}]}
        )

    def test_retry_automatic_list(self):
        expected = {
            "command": "test",
            "retry": {
                "automatic": [
                    {"exit_status": -1, "signal_reason": "none"},
                    {"signal": "kill"},
                    {"exit_status": 255},
                    {"exit_status": 3, "limit": 3},
                ]
            },
        }
        pipeline = Pipeline(
            steps=[
                CommandStep(
                    command="test",
                    retry=CommandStepRetry(
                        automatic=[
                            AutomaticRetry(exit_status=-1, signal_reason="none"),
                            AutomaticRetry(signal="kill"),
                            AutomaticRetry(exit_status=255),
                            AutomaticRetry(exit_status=3, limit=3),
                        ]
                    ),
                )
            ]
        )
        self.validator.check_result(pipeline, {"steps": [expected]})

    def test_automatic_retry_exit_status(self):
        expected = {
            "command": "test",
            "retry": {"automatic": {"exit_status": -1}},
        }
        pipeline = Pipeline(
            steps=[
                CommandStep(
                    command="test",
                    retry=CommandStepRetry(automatic=AutomaticRetry(exit_status=-1)),
                )
            ]
        )
        self.validator.check_result(pipeline, {"steps": [expected]})

    def test_automatic_retry_bool(self):
        expected: CommandStepArgs = {
            "command": "test",
            "retry": {"automatic": True},
        }
        pipeline = Pipeline(
            steps=[CommandStep(command="test", retry=CommandStepRetry(automatic=True))]
        )
        self.validator.check_result(pipeline, {"steps": [expected]})

    def test_automatic_retry_exit_status_list(self):
        expected = {
            "command": "test",
            "retry": {"automatic": {"exit_status": [1, 2, 3]}},
        }
        pipeline = Pipeline(
            steps=[
                CommandStep(
                    command="test",
                    retry=CommandStepRetry(
                        automatic=AutomaticRetry(exit_status=[1, 2, 3])
                    ),
                )
            ]
        )
        self.validator.check_result(pipeline, {"steps": [expected]})

    def test_secrets_list(self):
        pipeline = Pipeline(
            steps=[
                CommandStep(command='test', secrets=['MY_SECRET'])
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{'command': 'test', 'secrets': ['MY_SECRET']}]})

    def test_secrets_object(self):
        pipeline = Pipeline(
            steps=[
                CommandStep(command='test', secrets={"MY_SECRET": "API_TOKEN"})
            ]
        )
        self.validator.check_result(pipeline, {'steps': [{'command': 'test', 'secrets': {"MY_SECRET": "API_TOKEN"}}]})

    def test_skip_bool(self):
        pipeline = Pipeline(steps=[CommandStep(command="test", skip=True)])
        self.validator.check_result(
            pipeline, {"steps": [{"command": "test", "skip": True}]}
        )

    def test_skip_string(self):
        pipeline = Pipeline(steps=[CommandStep(command="test", skip="reason")])
        self.validator.check_result(
            pipeline, {"steps": [{"command": "test", "skip": "reason"}]}
        )

    def test_timeout_in_minutes(self):
        pipeline = Pipeline(steps=[CommandStep(command="test", timeout_in_minutes=40)])
        self.validator.check_result(
            pipeline, {"steps": [{"command": "test", "timeout_in_minutes": 40}]}
        )

    def test_soft_fail_bool(self):
        pipeline = Pipeline(steps=[CommandStep(command="test", soft_fail=True)])
        self.validator.check_result(
            pipeline, {"steps": [{"command": "test", "soft_fail": True}]}
        )

    def test_soft_fail_exit_status_int(self):
        pipeline = Pipeline(
            steps=[
                CommandStep(command="test", soft_fail=[SoftFailObject(exit_status=-1)])
            ]
        )
        self.validator.check_result(
            pipeline,
            {"steps": [{"command": "test", "soft_fail": [{"exit_status": -1}]}]},
        )

    def test_soft_fail_exit_status_string(self):
        pipeline = Pipeline(
            steps=[
                CommandStep(command="test", soft_fail=[SoftFailObject(exit_status="*")])
            ]
        )
        self.validator.check_result(
            pipeline,
            {"steps": [{"command": "test", "soft_fail": [{"exit_status": "*"}]}]},
        )

    def test_if(self):
        pipeline = Pipeline(
            steps=[CommandStep(command="test", step_if="build.message !~ /skip tests/")]
        )
        self.validator.check_result(
            pipeline,
            {"steps": [{"command": "test", "if": "build.message !~ /skip tests/"}]},
        )

    def test_key(self):
        pipeline = Pipeline(steps=[CommandStep(command="test", key="key")])
        self.validator.check_result(
            pipeline, {"steps": [{"command": "test", "key": "key"}]}
        )

    def test_depends_on_string(self):
        pipeline = Pipeline(steps=[CommandStep(command="test", depends_on="step")])
        self.validator.check_result(
            pipeline, {"steps": [{"command": "test", "depends_on": "step"}]}
        )

    def test_depends_on_string_list(self):
        pipeline = Pipeline(
            steps=[CommandStep(command="test", depends_on=["one", "two"])]
        )
        self.validator.check_result(
            pipeline, {"steps": [{"command": "test", "depends_on": ["one", "two"]}]}
        )

    def test_depends_on_object_list(self):
        expected: CommandStepArgs = {
            "command": "test",
            "depends_on": [{"step": "one", "allow_failure": True}, {"step": "two"}],
        }
        pipeline = Pipeline(
            steps=[
                CommandStep(
                    command="test",
                    depends_on=[
                        DependsOnListObject(step="one", allow_failure=True),
                        DependsOnListObject(step="two"),
                    ],
                )
            ]
        )
        self.validator.check_result(pipeline, {"steps": [expected]})

    def test_depends_on_mixed_list(self):
        expected: CommandStepArgs = {
            "command": "test",
            "depends_on": [
                {"step": "one", "allow_failure": True},
                "two",
            ],
        }
        pipeline = Pipeline(
            steps=[
                CommandStep(
                    command="test",
                    depends_on=[
                        DependsOnListObject(step="one", allow_failure=True),
                        "two",
                    ],
                )
            ]
        )
        self.validator.check_result(pipeline, {"steps": [expected]})

    def test_allow_dependency_failure(self):
        pipeline = Pipeline(
            steps=[CommandStep(command="test", allow_dependency_failure=True)]
        )
        self.validator.check_result(
            pipeline, {"steps": [{"command": "test", "allow_dependency_failure": True}]}
        )

    def test_priority(self):
        pipeline = Pipeline(steps=[CommandStep(command="test", priority=1)])
        self.validator.check_result(
            pipeline, {"steps": [{"command": "test", "priority": 1}]}
        )

    def test_cancel_on_build_failing(self):
        pipeline = Pipeline(
            steps=[CommandStep(command="test", cancel_on_build_failing=True)]
        )
        self.validator.check_result(
            pipeline, {"steps": [{"command": "test", "cancel_on_build_failing": True}]}
        )

    def test_signature(self):
        expected: CommandStepArgs = {
            "command": "echo {{matrix}}",
            "env": {"FOO": "bar"},
            "plugins": [{"docker#v3.0.0": {"image": "alpine", "always-pull": True}}],
            "matrix": ["one", "two", "three"],
            "signature": {
                "value": "not a real signature value",
                "algorithm": "HS256",
                "signed_fields": ["command", "env::FOO", "plugins", "matrix"],
            },
        }
        pipeline = Pipeline(
            steps=[
                CommandStep(
                    command="echo {{matrix}}",
                    env={"FOO": "bar"},
                    plugins=[
                        {"docker#v3.0.0": {"image": "alpine", "always-pull": True}}
                    ],
                    matrix=["one", "two", "three"],
                    signature=CommandStepSignature(
                        value="not a real signature value",
                        algorithm="HS256",
                        signed_fields=["command", "env::FOO", "plugins", "matrix"],
                    ),
                )
            ]
        )
        self.validator.check_result(pipeline, {"steps": [expected]})

    def test_cache_string(self):
        pipeline = Pipeline(steps=[CommandStep(command="test", cache="dist/")])
        self.validator.check_result(
            pipeline, {"steps": [{"command": "test", "cache": "dist/"}]}
        )

    def test_cache_list(self):
        pipeline = Pipeline(
            steps=[CommandStep(command="test", cache=["dist/", "./src/target/"])]
        )
        self.validator.check_result(
            pipeline,
            {"steps": [{"command": "test", "cache": ["dist/", "./src/target/"]}]},
        )

    def test_if_changed(self):
        pipeline = Pipeline(steps=[CommandStep(command="test", if_changed="*.txt")])
        self.validator.check_result(
            pipeline, {"steps": [{"command": "test", "if_changed": "*.txt"}]}
        )


class TestCommandStepArgs(TestRunner):
    def test_agents_dict(self):
        pipeline = Pipeline(
            steps=[
                CommandStep.from_dict({"command": "test", "agents": {"os": "macOS"}})
            ]
        )
        self.validator.check_result(
            pipeline, {"steps": [{"command": "test", "agents": {"os": "macOS"}}]}
        )

    def test_agents_list(self):
        pipeline = Pipeline(
            steps=[CommandStep.from_dict({"command": "test", "agents": ["os=macOS"]})]
        )
        self.validator.check_result(
            pipeline, {"steps": [{"command": "test", "agents": ["os=macOS"]}]}
        )

    def test_artifact_paths_list(self):
        pipeline = Pipeline(
            steps=[
                CommandStep.from_dict(
                    {"command": "test", "artifact_paths": ["one", "two"]}
                )
            ]
        )
        self.validator.check_result(
            pipeline, {"steps": [{"command": "test", "artifact_paths": ["one", "two"]}]}
        )

    def test_artifact_paths_string(self):
        pipeline = Pipeline(
            steps=[CommandStep.from_dict({"command": "test", "artifact_paths": "path"})]
        )
        self.validator.check_result(
            pipeline, {"steps": [{"command": "test", "artifact_paths": "path"}]}
        )

    def test_branches(self):
        pipeline = Pipeline(
            steps=[
                CommandStep.from_dict(
                    {"command": "test", "branches": "master deploy-*"}
                )
            ]
        )
        self.validator.check_result(
            pipeline, {"steps": [{"command": "test", "branches": "master deploy-*"}]}
        )

    def test_concurrency(self):
        pipeline = Pipeline(
            steps=[
                CommandStep.from_dict(
                    {
                        "command": "test",
                        "concurrency": 1,
                        "concurrency_group": "my-group",
                        "concurrency_method": "eager",
                    }
                )
            ]
        )
        self.validator.check_result(
            pipeline,
            {
                "steps": [
                    {
                        "command": "test",
                        "concurrency": 1,
                        "concurrency_group": "my-group",
                        "concurrency_method": "eager",
                    }
                ]
            },
        )

    def test_env(self):
        pipeline = Pipeline(
            steps=[CommandStep.from_dict({"command": "test", "env": {"ENV": "bar"}})]
        )
        self.validator.check_result(
            pipeline, {"steps": [{"command": "test", "env": {"ENV": "bar"}}]}
        )

    def test_id(self):
        pipeline = Pipeline(
            steps=[CommandStep.from_dict({"command": "test", "id": "id"})]
        )
        self.validator.check_result(
            pipeline, {"steps": [{"command": "test", "id": "id"}]}
        )

    def test_identifier(self):
        pipeline = Pipeline(
            steps=[
                CommandStep.from_dict({"command": "test", "identifier": "identifier"})
            ]
        )
        self.validator.check_result(
            pipeline, {"steps": [{"command": "test", "identifier": "identifier"}]}
        )

    def test_label(self):
        pipeline = Pipeline(
            steps=[CommandStep.from_dict({"command": "test", "label": "label"})]
        )
        self.validator.check_result(
            pipeline, {"steps": [{"command": "test", "label": "label"}]}
        )

    def test_parallelism(self):
        pipeline = Pipeline(
            steps=[CommandStep.from_dict({"command": "test", "parallelism": 42})]
        )
        self.validator.check_result(
            pipeline, {"steps": [{"command": "test", "parallelism": 42}]}
        )

    def test_plugins_dict(self):
        pipeline = Pipeline(
            steps=[
                CommandStep.from_dict(
                    {"command": "test", "plugins": {"a-plugin#v1.0.0": {"run": "app"}}}
                )
            ]
        )
        self.validator.check_result(
            pipeline,
            {
                "steps": [
                    {"command": "test", "plugins": {"a-plugin#v1.0.0": {"run": "app"}}}
                ]
            },
        )

    def test_plugins_dict_list(self):
        pipeline = Pipeline(
            steps=[
                CommandStep.from_dict(
                    {
                        "command": "test",
                        "plugins": [{"a-plugin#v1.0.0": {"run": "app"}}],
                    }
                )
            ]
        )
        self.validator.check_result(
            pipeline,
            {
                "steps": [
                    {
                        "command": "test",
                        "plugins": [{"a-plugin#v1.0.0": {"run": "app"}}],
                    }
                ]
            },
        )

    def test_plugins_string_list(self):
        pipeline = Pipeline(
            steps=[
                CommandStep.from_dict(
                    {"command": "test", "plugins": ["a-plugin#v1.0.0"]}
                )
            ]
        )
        self.validator.check_result(
            pipeline, {"steps": [{"command": "test", "plugins": ["a-plugin#v1.0.0"]}]}
        )

    def test_retry_automatic_list(self):
        expected: CommandStepArgs = {
            "command": "test",
            "retry": {
                "automatic": [
                    {"exit_status": -1, "signal_reason": "none"},
                    {"signal": "kill"},
                    {"exit_status": 255},
                    {"exit_status": 3, "limit": 3},
                ]
            },
        }
        pipeline = Pipeline(steps=[CommandStep.from_dict(expected)])
        self.validator.check_result(pipeline, {"steps": [expected]})

    def test_automatic_retry_exit_status(self):
        expected: CommandStepArgs = {
            "command": "test",
            "retry": {"automatic": {"exit_status": -1}},
        }
        pipeline = Pipeline(steps=[CommandStep.from_dict(expected)])
        self.validator.check_result(pipeline, {"steps": [expected]})

    def test_automatic_retry_bool(self):
        expected: CommandStepArgs = {
            "command": "test",
            "retry": {"automatic": True},
        }
        pipeline = Pipeline(steps=[CommandStep.from_dict(expected)])
        self.validator.check_result(pipeline, {"steps": [expected]})

    def test_automatic_retry_exit_status_list(self):
        expected: CommandStepArgs = {
            "command": "test",
            "retry": {"automatic": {"exit_status": [1, 2, 3]}},
        }
        pipeline = Pipeline(steps=[CommandStep.from_dict(expected)])
        self.validator.check_result(pipeline, {"steps": [expected]})

    def test_skip_bool(self):
        pipeline = Pipeline(
            steps=[CommandStep.from_dict({"command": "test", "skip": True})]
        )
        self.validator.check_result(
            pipeline, {"steps": [{"command": "test", "skip": True}]}
        )

    def test_skip_string(self):
        pipeline = Pipeline(
            steps=[CommandStep.from_dict({"command": "test", "skip": "reason"})]
        )
        self.validator.check_result(
            pipeline, {"steps": [{"command": "test", "skip": "reason"}]}
        )

    def test_timeout_in_minutes(self):
        pipeline = Pipeline(
            steps=[CommandStep.from_dict({"command": "test", "timeout_in_minutes": 40})]
        )
        self.validator.check_result(
            pipeline, {"steps": [{"command": "test", "timeout_in_minutes": 40}]}
        )

    def test_soft_fail_bool(self):
        pipeline = Pipeline(
            steps=[CommandStep.from_dict({"command": "test", "soft_fail": True})]
        )
        self.validator.check_result(
            pipeline, {"steps": [{"command": "test", "soft_fail": True}]}
        )

    def test_soft_fail_exit_status_int(self):
        pipeline = Pipeline(
            steps=[
                CommandStep.from_dict(
                    {"command": "test", "soft_fail": [{"exit_status": -1}]}
                )
            ]
        )
        self.validator.check_result(
            pipeline,
            {"steps": [{"command": "test", "soft_fail": [{"exit_status": -1}]}]},
        )

    def test_soft_fail_exit_status_string(self):
        pipeline = Pipeline(
            steps=[
                CommandStep.from_dict(
                    {"command": "test", "soft_fail": [{"exit_status": "*"}]}
                )
            ]
        )
        self.validator.check_result(
            pipeline,
            {"steps": [{"command": "test", "soft_fail": [{"exit_status": "*"}]}]},
        )

    def test_if(self):
        pipeline = Pipeline(
            steps=[
                CommandStep.from_dict(
                    {"command": "test", "if": "build.message !~ /skip tests/"}
                )
            ]
        )
        self.validator.check_result(
            pipeline,
            {"steps": [{"command": "test", "if": "build.message !~ /skip tests/"}]},
        )

    def test_key(self):
        pipeline = Pipeline(
            steps=[CommandStep.from_dict({"command": "test", "key": "key"})]
        )
        self.validator.check_result(
            pipeline, {"steps": [{"command": "test", "key": "key"}]}
        )

    def test_depends_on_string(self):
        pipeline = Pipeline(
            steps=[CommandStep.from_dict({"command": "test", "depends_on": "step"})]
        )
        self.validator.check_result(
            pipeline, {"steps": [{"command": "test", "depends_on": "step"}]}
        )

    def test_depends_on_string_list(self):
        pipeline = Pipeline(
            steps=[
                CommandStep.from_dict({"command": "test", "depends_on": ["one", "two"]})
            ]
        )
        self.validator.check_result(
            pipeline, {"steps": [{"command": "test", "depends_on": ["one", "two"]}]}
        )

    def test_depends_on_object_list(self):
        expected: CommandStepArgs = {
            "command": "test",
            "depends_on": [{"step": "one", "allow_failure": True}, {"step": "two"}],
        }
        pipeline = Pipeline(steps=[CommandStep.from_dict(expected)])
        self.validator.check_result(pipeline, {"steps": [expected]})

    def test_depends_on_mixed_list(self):
        expected: CommandStepArgs = {
            "command": "test",
            "depends_on": [
                {"step": "one", "allow_failure": True},
                "two",
            ],
        }
        pipeline = Pipeline(steps=[CommandStep.from_dict(expected)])
        self.validator.check_result(pipeline, {"steps": [expected]})

    def test_allow_dependency_failure(self):
        pipeline = Pipeline(
            steps=[
                CommandStep.from_dict(
                    {"command": "test", "allow_dependency_failure": True}
                )
            ]
        )
        self.validator.check_result(
            pipeline, {"steps": [{"command": "test", "allow_dependency_failure": True}]}
        )

    def test_priority(self):
        pipeline = Pipeline(
            steps=[CommandStep.from_dict({"command": "test", "priority": 1})]
        )
        self.validator.check_result(
            pipeline, {"steps": [{"command": "test", "priority": 1}]}
        )

    def test_cancel_on_build_failing(self):
        pipeline = Pipeline(
            steps=[
                CommandStep.from_dict(
                    {"command": "test", "cancel_on_build_failing": True}
                )
            ]
        )
        self.validator.check_result(
            pipeline, {"steps": [{"command": "test", "cancel_on_build_failing": True}]}
        )

    def test_signature(self):
        expected: CommandStepArgs = {
            "command": "echo {{matrix}}",
            "env": {"FOO": "bar"},
            "plugins": [{"docker#v3.0.0": {"image": "alpine", "always-pull": True}}],
            "matrix": ["one", "two", "three"],
            "signature": {
                "value": "not a real signature value",
                "algorithm": "HS256",
                "signed_fields": ["command", "env::FOO", "plugins", "matrix"],
            },
        }
        pipeline = Pipeline(steps=[CommandStep.from_dict(expected)])
        self.validator.check_result(pipeline, {"steps": [expected]})

    def test_cache_string(self):
        pipeline = Pipeline(
            steps=[CommandStep.from_dict({"command": "test", "cache": "dist/"})]
        )
        self.validator.check_result(
            pipeline, {"steps": [{"command": "test", "cache": "dist/"}]}
        )

    def test_cache_list(self):
        pipeline = Pipeline(
            steps=[
                CommandStep.from_dict(
                    {"command": "test", "cache": ["dist/", "./src/target/"]}
                )
            ]
        )
        self.validator.check_result(
            pipeline,
            {"steps": [{"command": "test", "cache": ["dist/", "./src/target/"]}]},
        )

    def test_if_changed(self):
        pipeline = Pipeline(
            steps=[CommandStep.from_dict({"command": "test", "if_changed": "*.txt"})]
        )
        self.validator.check_result(
            pipeline, {"steps": [{"command": "test", "if_changed": "*.txt"}]}
        )
