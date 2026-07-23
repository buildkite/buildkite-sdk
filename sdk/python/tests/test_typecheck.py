import subprocess
import sys
from pathlib import Path


def test_checkout_usage_passes_pyright():
    sample = Path(__file__).with_name("typecheck_sample.py")
    result = subprocess.run(
        ["pyright", "--pythonpath", sys.executable, str(sample)],
        capture_output=True,
        text=True,
        cwd=Path(__file__).parent.parent,
    )
    assert result.returncode == 0, f"pyright failed:\n{result.stdout}\n{result.stderr}"
