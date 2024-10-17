package typescript_code_gen

import (
	"testing"

	"github.com/buildkite/pipeline-sdk/pkg/schema"
	"github.com/stretchr/testify/assert"
)

func TestGenerateReturnStatement(t *testing.T) {
	t.Run("should handle a boolean", func(t *testing.T) {
		envVar := schema.NewEnvVar("TEST").Boolean()
		result := generateReturnStatement(envVar.GetDefinition())
		assert.Equal(t, "return Boolean(process.env.TEST!)", result)
	})

	t.Run("should handle a number", func(t *testing.T) {
		envVar := schema.NewEnvVar("TEST").Number()
		result := generateReturnStatement(envVar.GetDefinition())
		assert.Equal(t, "return Number(process.env.TEST!)", result)
	})

	t.Run("should handle a string", func(t *testing.T) {
		envVar := schema.NewEnvVar("TEST").String()
		result := generateReturnStatement(envVar.GetDefinition())
		assert.Equal(t, "return process.env.TEST!", result)
	})

	t.Run("should handle a string array", func(t *testing.T) {
		envVar := schema.NewEnvVar("TEST").StringArray(",")
		result := generateReturnStatement(envVar.GetDefinition())
		assert.Equal(t, "return process.env.TEST!.split(\",\")", result)
	})

	t.Run("should handle dynamic", func(t *testing.T) {
		envVar := schema.NewEnvVar("TEST").Dynamic().String()
		result := generateReturnStatement(envVar.GetDefinition())
		assert.Equal(t, "return process.env[strs.join(\"_\").toUpperCase()]!", result)
	})
}

func TestGenerateEnvironmentVariableMethod(t *testing.T) {
	t.Run("should generate an environment variable function", func(t *testing.T) {
		envVar := schema.NewEnvVar("TEST").Description("description").String()
		result := generateEnvironmentVariableMethod(envVar.GetDefinition())

		assert.Equal(t, 4, len(result))
		assert.Equal(t, "// description", result[0])
		assert.Equal(t, "public TEST(): string {", result[1])
		assert.Equal(t, "    return process.env.TEST!;", result[2])
		assert.Equal(t, "}", result[3])
	})

	t.Run("should generate a dynamic environment variable function", func(t *testing.T) {
		envVar := schema.NewEnvVar("TEST").Description("description").Dynamic().String()
		result := generateEnvironmentVariableMethod(envVar.GetDefinition())

		assert.Equal(t, 4, len(result))
		assert.Equal(t, "// description", result[0])
		assert.Equal(t, "public TEST(...strs: string[]): string {", result[1])
		assert.Equal(t, "    return process.env[strs.join(\"_\").toUpperCase()]!;", result[2])
		assert.Equal(t, "}", result[3])
	})
}

func TestNewEnvironmentFile(t *testing.T) {
	t.Run("should create an environment variable class", func(t *testing.T) {
		envVars := []schema.EnvironmentVariable{
			schema.NewEnvVar("TEST").Description("description").String(),
		}

		file := newEnvironmentFile(envVars)
		assert.Greater(t, len(file.Render()), 0)
	})
}
