package typescript_code_gen

import (
	"fmt"

	"github.com/buildkite/pipeline-sdk/pkg/schema"
	"github.com/buildkite/pipeline-sdk/pkg/schema_types"
	"github.com/buildkite/pipeline-sdk/pkg/utils"
)

func newEnvironmentFile(envs []schema.EnvironmentVariable) string {
	file := NewFile()
	file.code = append(file.code, "class Environment {")

	for _, env := range envs {
		def := env.GetDefinition()
		envKey := fmt.Sprintf("process.env.%s", def.Name)
		if def.Dynamic {
			envKey = "process.env[strs.join(\"_\").toUpperCase()]"
		}

		var returnStatement string
		returnType := def.Typ.TypeScriptType()
		switch def.Typ.(type) {
		case schema_types.SchemaBoolean:
			returnStatement = fmt.Sprintf("return Boolean(%s)", envKey)
		case schema_types.SchemaNumber:
			returnStatement = fmt.Sprintf("return Number(%s)", envKey)
		case schema_types.SchemaArray:
			switch returnType {
			case "boolean[]":
				returnStatement = fmt.Sprintf("return %s.split(\"%s\").map(v => Boolean(v))", envKey, def.Metadata["delimiter"])
			case "number[]":
				returnStatement = fmt.Sprintf("return %s.split(\"%s\").map(v => Number(v))", envKey, def.Metadata["delimiter"])
			default:
				returnStatement = fmt.Sprintf("return %s.split(\"%s\")", envKey, def.Metadata["delimiter"])
			}
		default:
			returnStatement = fmt.Sprintf("return %s", envKey)
		}

		dynamicArgs := ""
		if def.Dynamic {
			dynamicArgs = "...strs: string[]"
		}

		block := utils.CodeBlock{
			utils.CodeGen.Comment.TypeScript(def.Description, 0),
			fmt.Sprintf("public %s(%s): %s {", def.Name, dynamicArgs, returnType),
			fmt.Sprintf("    %s;", returnStatement),
			"}",
		}

		file.code = append(file.code, block.DisplayIndent(4))
	}

	file.code = append(file.code, "}")
	file.code = append(file.code, "\n\nexport default Environment;")

	return file.String()
}
