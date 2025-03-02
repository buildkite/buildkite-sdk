from buildkite_sdk import Pipeline
from os import makedirs

def generate_json():
    pipeline = Pipeline()
    pipeline.add_command_step({"label": "some-label", "command": "echo 'Hello, world!'"})
    return pipeline.to_json()

def generate_yaml():
    pipeline = Pipeline()
    pipeline.add_command_step({"label": "some-label", "command": "echo 'Hello, world!'"})
    return pipeline.to_yaml()

makedirs("../../out/apps/python", exist_ok=True)

with open("../../out/apps/python/pipeline.json", "w") as file:
    file.write(generate_json())

with open("../../out/apps/python/pipeline.yaml", "w") as file:
    file.write(generate_yaml())
