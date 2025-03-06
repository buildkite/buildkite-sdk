package buildkite

func Value[T any](val T) *T {
	return &val
}
