package gogen

import "fmt"

func NewGoComment(comment string) string {
	return fmt.Sprintf("// %s", comment)
}
