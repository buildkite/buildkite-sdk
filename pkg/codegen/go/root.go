package go_code_gen

import "fmt"

func newVersionFile(version string) string {
	return fmt.Sprintf("package buildkite\n\nvar PkgVersion = \"%s\"", version)
}
