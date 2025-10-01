package utils

import "strings"

type codeBlock struct {
	lines []string
}

func (c *codeBlock) AddLines(lines ...string) {
	c.lines = append(c.lines, lines...)
}

func (c codeBlock) Length() int {
	return len(c.lines)
}

func (c codeBlock) String() string {
	return strings.Join(c.lines, "\n")
}

func NewCodeBlock(lines ...string) *codeBlock {
	return &codeBlock{lines: lines}
}
