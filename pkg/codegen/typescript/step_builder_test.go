package typescript_code_gen

import (
	"testing"

	"github.com/buildkite/pipeline-sdk/pkg/schema"
	"github.com/stretchr/testify/assert"
)

func TestRenderStepFunction(t *testing.T) {
	t.Run("should render step builder function", func(t *testing.T) {
		step := schema.PipelinesSchema.Steps[0]
		result := renderStepFunction(step)
		assert.Equal(t, "    // A block step is used to pause the execution of a build and wait on a team member to unblock it using the web or the API.\n    public addBlockStep(args: types.Block): this {\n        this.steps.push({ ...args });\n        return this;\n    }", result)
	})
}

func TestNewStepBuilderFile(t *testing.T) {
	t.Run("should render a step build file", func(t *testing.T) {
		file := newStepBuilderFile(schema.PipelinesSchema)
		assert.Greater(t, len(file), 0)
	})
}
