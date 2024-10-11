package schema_types

import "github.com/buildkite/pipeline-sdk/pkg/utils"

type AttributeName string

func (p AttributeName) TitleCase() string {
	return utils.SnakeCaseToTitleCase(string(p))
}

func (p AttributeName) CamelCase() string {
	return utils.SnakeCaseToCamelCase(string(p))
}
