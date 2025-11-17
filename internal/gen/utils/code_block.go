package utils

import (
	"fmt"
	"strings"
)

type codeBlock struct {
	lines []string
}

func (c *codeBlock) AddLines(lines ...string) {
	c.lines = append(c.lines, lines...)
}

func (c *codeBlock) AddLinesWithIndent(indentSize int, lines ...string) {
	indent := strings.Repeat(" ", indentSize)
	for _, line := range lines {
		c.lines = append(c.lines, fmt.Sprintf("%s%s", indent, line))
	}
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
