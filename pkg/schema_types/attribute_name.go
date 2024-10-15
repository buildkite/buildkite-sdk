package schema_types

import "github.com/buildkite/pipeline-sdk/pkg/utils"

type AttributeName string

func (p AttributeName) TitleCase() string {
	return utils.String.SnakeCaseToTitleCase(string(p))
}

func (p AttributeName) CamelCase() string {
	return utils.String.SnakeCaseToCamelCase(string(p))
}
