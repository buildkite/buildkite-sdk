package go_code_gen

import (
	"testing"

	"github.com/buildkite/pipeline-sdk/pkg/schema"
	"github.com/buildkite/pipeline-sdk/pkg/schema_types"
	"github.com/stretchr/testify/assert"
)

func TestRenderGetenvBlock(t *testing.T) {
	t.Run("not dynamic", func(t *testing.T) {
		result := renderGetenvBlock("TEST", false)
		assert.Equal(t, "str := os.Getenv(\"TEST\")", result)
	})

	t.Run("dynamic", func(t *testing.T) {
		result := renderGetenvBlock("TEST", true)
		assert.Equal(t, "envKey := strings.ToUpper(strings.Join(strs, \"_\"))\n    str := os.Getenv(envKey)\n", result)
	})
}

func TestRenderReturnStatement(t *testing.T) {
	t.Run("should handle a boolean", func(t *testing.T) {
		val := schema_types.Simple.Boolean()
		result := renderReturnStatement(val, nil)
		assert.Equal(t, "return ParseStringToBool(str)", result)
	})

	t.Run("should handle a number", func(t *testing.T) {
		val := schema_types.Simple.Number()
		result := renderReturnStatement(val, nil)
		assert.Equal(t, "return ParseStringToInt(str)", result)
	})

	t.Run("should handle a string array", func(t *testing.T) {
		val := schema_types.Array.String()
		result := renderReturnStatement(val, map[string]string{"delimeter": ","})
		assert.Equal(t, "return strings.Split(str, \"\")", result)
	})

	t.Run("should handle a string", func(t *testing.T) {
		val := schema_types.Simple.String()
		result := renderReturnStatement(val, nil)
		assert.Equal(t, "return str", result)
	})
}

func TestRenderEnvVarArgs(t *testing.T) {
	t.Run("not dynamic", func(t *testing.T) {
		result := renderEnvVarArgs(false)
		assert.Equal(t, "", result)
	})

	t.Run("dynamic", func(t *testing.T) {
		result := renderEnvVarArgs(true)
		assert.Equal(t, "strs ...string", result)
	})
}

func TestRenderEnvironmentVaribleMethod(t *testing.T) {
	t.Run("should render an env var method", func(t *testing.T) {
		envVar := schema.NewEnvVar("TEST").Description("description").String()
		result := renderEnvironmentVaribleMethod(envVar)
		assert.Equal(t, "// description\nfunc (e environment) TEST() string {\n    str := os.Getenv(\"TEST\")\n    return str\n}", result)
	})
}

func TestNewEnvironmentFile(t *testing.T) {
	t.Run("should render an environment file", func(t *testing.T) {
		envVar := schema.NewEnvVar("TEST").Description("description").String()
		result := newEnvironmentFile([]schema.EnvironmentVariable{envVar})
		assert.Greater(t, len(result), 0)
	})
}
