package go_code_gen

import (
	"fmt"

	"github.com/buildkite/pipeline-sdk/pkg/schema"
	"github.com/buildkite/pipeline-sdk/pkg/schema_types"
	"github.com/buildkite/pipeline-sdk/pkg/utils"
)

var environmentFile = `// This file is auto-generated please do not edit

package buildkite

import (
	"os"
	"strings"
)

type environment struct{}`

func newEnvironmentFile(envs []schema.EnvironmentVariable) string {
	file := utils.CodeBlock{environmentFile}

	methods := utils.CodeBlock{}
	for _, env := range envs {
		def := env.GetDefinition()
		codeBlock := fmt.Sprintf(`str := os.Getenv("%s")`, def.Name)
		if def.Dynamic {
			codeBlock = fmt.Sprintf("%s\n%s\n",
				"envKey := strings.ToUpper(strings.Join(strs, \"_\"))",
				"str := os.Getenv(envKey)",
			)
		}

		methodBody := utils.CodeBlock{codeBlock}

		returnType := def.Typ.GoType()
		switch def.Typ.(type) {
		case schema_types.SchemaBoolean:
			methodBody = append(methodBody, "return ParseStringToBool(str)")
		case schema_types.SchemaNumber:
			methodBody = append(methodBody, "return ParseStringToInt(str)")
		case schema_types.SchemaArray:
			methodBody = append(methodBody, fmt.Sprintf("return strings.Split(str, \"%s\")", def.Metadata["delimiter"]))
		default:
			returnType = "string"
			methodBody = append(methodBody, "return str")
		}

		dynamicArgs := ""
		if def.Dynamic {
			dynamicArgs = "strs ...string"
		}

		methods = append(methods, fmt.Sprintf("%s\n%s\n%s\n%s\n",
			utils.NewCodeComment(def.Description, 0),
			fmt.Sprintf("func (e environment) %s(%s) %s {", def.Name, dynamicArgs, returnType),
			methodBody.DisplayIndent(4),
			"}",
		))
	}

	file = append(file, methods.Display())

	return file.Display()
}
