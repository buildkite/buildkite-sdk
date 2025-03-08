from src.main import generate_json, generate_yaml

def test_main():
    assert generate_json() == """{
    "steps": [
        {
            "commands": "echo 'Hello, world!'",
            "label": "some-label"
        }
    ]
}"""
    assert generate_yaml() == """steps:
- commands: echo 'Hello, world!'
  label: some-label
"""
