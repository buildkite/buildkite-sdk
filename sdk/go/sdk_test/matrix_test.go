package sdk_test

import (
	"testing"

	buildkite "github.com/buildkite/buildkite-sdk/sdk/go/sdk/buildkite"
)

type testMatrixSetup struct {
	Setup buildkite.MatrixSetup `json:"setup"`
}

func TestMatrixSetup(t *testing.T) {
	t.Run("MatrixElementList", func(t *testing.T) {
		value := "value"
		list := []buildkite.MatrixElement{
			{
				String: &value,
			},
		}
		val := testMatrixSetup{
			Setup: buildkite.MatrixSetup{
				MatrixElementList: &list,
			},
		}
		CheckResult(t, val, `{"setup":["value"]}`)
	})

	t.Run("MatrixSetup", func(t *testing.T) {
		elementValue := "bar"
		value := map[string][]buildkite.MatrixElement{
			"foo": {
				{
					String: &elementValue,
				},
			},
		}
		val := testMatrixSetup{
			Setup: buildkite.MatrixSetup{
				MatrixSetup: &value,
			},
		}
		CheckResult(t, val, `{"setup":{"foo":["bar"]}}`)
	})
}

type testMatrixAdjustments struct {
	Adjustments buildkite.MatrixAdjustments `json:"adjustments"`
}

func TestMatrixAdjustments(t *testing.T) {
	t.Run("With", func(t *testing.T) {
		t.Run("MatrixElementList", func(t *testing.T) {
			value := "value"
			list := []buildkite.MatrixElement{
				{
					String: &value,
				},
			}
			val := testMatrixAdjustments{
				Adjustments: buildkite.MatrixAdjustments{
					With: &buildkite.MatrixAdjustmentsWith{
						MatrixElementList: &list,
					},
				},
			}
			CheckResult(t, val, `{"adjustments":{"with":["value"]}}`)
		})

		t.Run("MatrixAdjustmentsObject", func(t *testing.T) {
			value := buildkite.MatrixAdjustmentsWithObject{
				"foo": "bar",
			}
			val := testMatrixAdjustments{
				Adjustments: buildkite.MatrixAdjustments{
					With: &buildkite.MatrixAdjustmentsWith{
						MatrixAdjustmentsWithObject: &value,
					},
				},
			}
			CheckResult(t, val, `{"adjustments":{"with":{"foo":"bar"}}}`)
		})
	})

	t.Run("Skip", func(t *testing.T) {
		skip := "skip"
		val := testMatrixAdjustments{
			Adjustments: buildkite.MatrixAdjustments{
				Skip: &buildkite.Skip{
					String: &skip,
				},
			},
		}
		CheckResult(t, val, `{"adjustments":{"skip":"skip"}}`)
	})

	t.Run("Softfail", func(t *testing.T) {
		softFail := true
		val := testMatrixAdjustments{
			Adjustments: buildkite.MatrixAdjustments{
				SoftFail: &buildkite.SoftFail{
					SoftFailEnum: &buildkite.SoftFailEnum{
						Bool: &softFail,
					},
				},
			},
		}
		CheckResult(t, val, `{"adjustments":{"soft_fail":true}}`)
	})

	t.Run("All", func(t *testing.T) {
		skip := "skip"
		softFail := true
		with := buildkite.MatrixAdjustmentsWithObject{
			"foo": "bar",
		}
		val := testMatrixAdjustments{
			Adjustments: buildkite.MatrixAdjustments{
				Skip: &buildkite.Skip{
					String: &skip,
				},
				SoftFail: &buildkite.SoftFail{
					SoftFailEnum: &buildkite.SoftFailEnum{
						Bool: &softFail,
					},
				},
				With: &buildkite.MatrixAdjustmentsWith{
					MatrixAdjustmentsWithObject: &with,
				},
			},
		}
		CheckResult(t, val, `{"adjustments":{"skip":"skip","soft_fail":true,"with":{"foo":"bar"}}}`)
	})
}

func TestMatrixObject(t *testing.T) {
	t.Run("Setup", func(t *testing.T) {
		elementValue := "bar"
		value := map[string][]buildkite.MatrixElement{
			"foo": {
				{
					String: &elementValue,
				},
			},
		}
		val := buildkite.MatrixObject{
			Setup: &buildkite.MatrixSetup{
				MatrixSetup: &value,
			},
		}
		CheckResult(t, val, `{"setup":{"foo":["bar"]}}`)
	})

	t.Run("Adjustments", func(t *testing.T) {
		value := buildkite.MatrixAdjustmentsWithObject{
			"foo": "bar",
		}
		val := buildkite.MatrixObject{
			Adjustments: []buildkite.MatrixAdjustments{
				{
					With: &buildkite.MatrixAdjustmentsWith{
						MatrixAdjustmentsWithObject: &value,
					},
				},
			},
		}
		CheckResult(t, val, `{"adjustments":[{"with":{"foo":"bar"}}]}`)
	})

	t.Run("Both", func(t *testing.T) {
		elementValue := "bar"
		setup := map[string][]buildkite.MatrixElement{
			"foo": {
				{
					String: &elementValue,
				},
			},
		}
		adjustments := buildkite.MatrixAdjustmentsWithObject{
			"foo": "bar",
		}
		val := buildkite.MatrixObject{
			Setup: &buildkite.MatrixSetup{
				MatrixSetup: &setup,
			},
			Adjustments: []buildkite.MatrixAdjustments{
				{
					With: &buildkite.MatrixAdjustmentsWith{
						MatrixAdjustmentsWithObject: &adjustments,
					},
				},
			},
		}
		CheckResult(t, val, `{"adjustments":[{"with":{"foo":"bar"}}],"setup":{"foo":["bar"]}}`)
	})
}

type testMatrix struct {
	Matrix buildkite.Matrix `json:"matrix"`
}

func TestMatrix(t *testing.T) {
	t.Run("MatrixElementList", func(t *testing.T) {
		element := "value"
		list := buildkite.MatrixElementList{
			{
				String: &element,
			},
		}
		val := testMatrix{
			Matrix: buildkite.Matrix{
				MatrixElementList: &list,
			},
		}
		CheckResult(t, val, `{"matrix":["value"]}`)
	})

	t.Run("MatrixObject", func(t *testing.T) {
		elementValue := "bar"
		value := map[string][]buildkite.MatrixElement{
			"foo": {
				{
					String: &elementValue,
				},
			},
		}
		val := testMatrix{
			Matrix: buildkite.Matrix{
				MatrixObject: &buildkite.MatrixObject{
					Setup: &buildkite.MatrixSetup{
						MatrixSetup: &value,
					},
				},
			},
		}
		CheckResult(t, val, `{"matrix":{"setup":{"foo":["bar"]}}}`)
	})
}
