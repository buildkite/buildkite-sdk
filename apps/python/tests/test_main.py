from src.main import generate_json, generate_yaml

def test_main():
    assert generate_json() == """{
    "steps": [
        {
            "label": "some-label",
            "command": "echo 'Hello, world!'"
        }
    ]
}"""
    assert generate_yaml() == """steps:
- command: echo 'Hello, world!'
  label: some-label
"""
