package schema

import "github.com/buildkite/pipeline-sdk/pkg/schema_types"

type Step struct {
	Name        string
	Description string
	Fields      []schema_types.Field
}

func (s Step) ToObjectField() schema_types.Field {
	return schema_types.NewField().
		Name(s.Name).
		Description(s.Description).
		Object(s.Fields)
}
