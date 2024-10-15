package schema

import (
	"testing"

	"github.com/buildkite/pipeline-sdk/pkg/schema_types"
	"github.com/stretchr/testify/assert"
)

func TestEnvironmentVariable(t *testing.T) {
	t.Run("should create a new environment variable", func(t *testing.T) {
		envVar := NewEnvVar("ENV_VAR")
		assert.Equal(t, "ENV_VAR", envVar.name)
	})

	t.Run("should set the description", func(t *testing.T) {
		envVar := NewEnvVar("ENV_VAR").Description("env var description")
		assert.Equal(t, "env var description", envVar.description)
	})

	t.Run("should mark as dynamic", func(t *testing.T) {
		envVar := NewEnvVar("ENV_VAR").Dynamic()
		assert.Equal(t, true, envVar.dynamic)
	})

	t.Run("should create a string enviornment variable", func(t *testing.T) {
		envVar := NewEnvVar("ENV_VAR").String()
		assert.Equal(t, schema_types.SchemaString{}, envVar.typ)
	})

	t.Run("should create a string array environment variable", func(t *testing.T) {
		envVar := NewEnvVar("ENV_VAR").StringArray(",")
		assert.Equal(t, schema_types.SchemaArray{Items: schema_types.SchemaString{}}, envVar.typ)
	})

	t.Run("should create a number environment variable", func(t *testing.T) {
		envVar := NewEnvVar("ENV_VAR").Number()
		assert.Equal(t, schema_types.SchemaNumber{}, envVar.typ)
	})

	t.Run("should create a boolean enviornment variable", func(t *testing.T) {
		envVar := NewEnvVar("ENV_VAR").Boolean()
		assert.Equal(t, schema_types.SchemaBoolean{}, envVar.typ)
	})

	t.Run("should get an enviornment variable definition", func(t *testing.T) {
		envVar := NewEnvVar("ENV_VAR").Description("env var description").String()
		def := envVar.GetDefinition()

		assert.Equal(t, envVar.name, def.Name)
		assert.Equal(t, envVar.description, def.Description)
		assert.Equal(t, envVar.dynamic, def.Dynamic)
		assert.Equal(t, envVar.typ, def.Typ)
	})
}
