package go_code_gen

import (
	"fmt"

	"github.com/buildkite/pipeline-sdk/pkg/schema"
	"github.com/buildkite/pipeline-sdk/pkg/schema_types"
	"github.com/buildkite/pipeline-sdk/pkg/utils"
)

var environmentFile = `type environment struct{}`

func renderGetenvBlock(name string, dynamic bool) string {
	if dynamic {
		return fmt.Sprintf("%s\n    %s\n",
			"envKey := strings.ToUpper(strings.Join(strs, \"_\"))",
			"str := os.Getenv(envKey)",
		)
	}

	return fmt.Sprintf(`str := os.Getenv("%s")`, name)
}

func renderReturnStatement(typ schema_types.SchemaType, metadata map[string]string) string {
	switch typ.(type) {
	case schema_types.SchemaBoolean:
		return "return ParseStringToBool(str)"
	case schema_types.SchemaNumber:
		return "return ParseStringToInt(str)"
	case schema_types.SchemaArray:
		return fmt.Sprintf("return strings.Split(str, \"%s\")", metadata["delimiter"])
	default:
		return "return str"
	}
}

func renderEnvVarArgs(dynamic bool) string {
	if dynamic {
		return "strs ...string"
	}
	return ""
}

func renderEnvironmentVaribleMethod(env schema.EnvironmentVariable) string {
	def := env.GetDefinition()

	methodBody := utils.CodeGen.NewCodeBlock(
		renderGetenvBlock(def.Name, def.Dynamic),
		renderReturnStatement(def.Typ, def.Metadata),
	)

	args := renderEnvVarArgs(def.Dynamic)

	return utils.CodeGen.NewCodeBlock(
		utils.CodeGen.Comment.Go(def.Description, 0),
		fmt.Sprintf("func (e environment) %s(%s) %s {", def.Name, args, def.Typ.GoType()),
		methodBody.DisplayIndent(4),
		"}",
	).Display()

}

func newEnvironmentFile(envs []schema.EnvironmentVariable) string {
	file := newFile()
	file.AddImport("os", "os")
	file.AddImport("strings", "strings")

	methods := utils.CodeGen.NewCodeBlock(environmentFile)

	for _, env := range envs {
		methods = append(methods, renderEnvironmentVaribleMethod(env))
	}

	file.AppendCode(methods.Display())

	return file.Render()
}
