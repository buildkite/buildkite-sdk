package schema_types

import (
	"fmt"
	"strings"
)

type SchemaUnion struct {
	Values []Field
}

func (s SchemaUnion) TypeScriptType() string {
	unionValues := make([]string, len(s.Values))
	for i, val := range s.Values {
		unionValues[i] = val.TypeScriptIdentifier()
	}

	return fmt.Sprintf("(%s)", strings.Join(unionValues, " | "))
}

func (s SchemaUnion) GoType() string {
	return ""
}
