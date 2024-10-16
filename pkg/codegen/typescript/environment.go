package typescript_code_gen

import (
	"fmt"

	code_gen_utils "github.com/buildkite/pipeline-sdk/pkg/codegen/utils"
	"github.com/buildkite/pipeline-sdk/pkg/schema"
	"github.com/buildkite/pipeline-sdk/pkg/schema_types"
	"github.com/buildkite/pipeline-sdk/pkg/utils"
)

func generateReturnStatement(def schema.EnvironmentVariableDefinition) string {
	envKey := fmt.Sprintf("process.env.%s!", def.Name)
	if def.Dynamic {
		envKey = "process.env[strs.join(\"_\").toUpperCase()]!"
	}

	switch def.Typ.(type) {
	case schema_types.SchemaBoolean:
		return fmt.Sprintf("return Boolean(%s)", envKey)
	case schema_types.SchemaNumber:
		return fmt.Sprintf("return Number(%s)", envKey)
	case schema_types.SchemaArray:
		return fmt.Sprintf("return %s.split(\"%s\")", envKey, def.Metadata["delimiter"])
	default:
		return fmt.Sprintf("return %s", envKey)
	}
}

func generateEnvironmentVariableMethod(def schema.EnvironmentVariableDefinition) utils.CodeBlock {
	returnStatement := generateReturnStatement(def)

	dynamicArgs := ""
	if def.Dynamic {
		dynamicArgs = "...strs: string[]"
	}

	return utils.CodeGen.NewCodeBlock(
		utils.CodeGen.Comment.TypeScript(def.Description, 0),
		fmt.Sprintf("public %s(%s): %s {", def.Name, dynamicArgs, def.Typ.TypeScriptType()),
		fmt.Sprintf("    %s;", returnStatement),
		"}",
	)
}

func newEnvironmentFile(envs []schema.EnvironmentVariable) *code_gen_utils.CodeGenFile {
	file := newFile()

	file.AppendCode("class Environment {")

	for _, env := range envs {
		def := env.GetDefinition()
		block := generateEnvironmentVariableMethod(def)
		file.AppendCode(block.DisplayIndent(4))
	}

	file.AppendCode("}\n\nexport default Environment;")

	return file
}
