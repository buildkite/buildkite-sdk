package typescript

import "fmt"

func NewTypeDocComment(comment string) string {
	if comment == "" {
		return ""
	}

	return fmt.Sprintf("/**\n     * %s\n    */", comment)
}
