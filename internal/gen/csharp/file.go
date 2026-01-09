package csharp

import (
	"fmt"
	"os"
	"strings"
)

// CSharpFile represents a C# source file
type CSharpFile struct {
	Namespace string
	FileName  string
	Usings    []string
	Content   string
}

// NewCSharpFile creates a new C# file
func NewCSharpFile(namespace, fileName string, usings []string, content string) *CSharpFile {
	return &CSharpFile{
		Namespace: namespace,
		FileName:  fileName,
		Usings:    usings,
		Content:   content,
	}
}

// Write writes the C# file to disk
func (f *CSharpFile) Write() error {
	var sb strings.Builder

	// Write usings
	for _, using := range f.Usings {
		sb.WriteString(fmt.Sprintf("using %s;\n", using))
	}
	if len(f.Usings) > 0 {
		sb.WriteString("\n")
	}

	// Write namespace and content
	sb.WriteString(fmt.Sprintf("namespace %s;\n\n", f.Namespace))
	sb.WriteString(f.Content)

	return os.WriteFile(f.FileName, []byte(sb.String()), 0644)
}

// CSharpClass represents a C# class definition
type CSharpClass struct {
	Name        string
	Description string
	Properties  []CSharpProperty
	Implements  []string
}

// CSharpProperty represents a property in a C# class
type CSharpProperty struct {
	Name         string
	Type         string
	JsonName     string
	Description  string
	IsNullable   bool
}

// NewCSharpClass creates a new C# class
func NewCSharpClass(name, description string) *CSharpClass {
	return &CSharpClass{
		Name:        name,
		Description: description,
		Properties:  []CSharpProperty{},
		Implements:  []string{},
	}
}

// AddProperty adds a property to the class
func (c *CSharpClass) AddProperty(prop CSharpProperty) {
	c.Properties = append(c.Properties, prop)
}

// AddImplements adds an interface implementation
func (c *CSharpClass) AddImplements(iface string) {
	c.Implements = append(c.Implements, iface)
}

// Generate generates the C# class code
func (c *CSharpClass) Generate() string {
	var sb strings.Builder

	// XML doc comment
	if c.Description != "" {
		sb.WriteString("/// <summary>\n")
		for _, line := range strings.Split(c.Description, "\n") {
			sb.WriteString(fmt.Sprintf("/// %s\n", strings.TrimSpace(line)))
		}
		sb.WriteString("/// </summary>\n")
	}

	// Class declaration
	sb.WriteString(fmt.Sprintf("public class %s", c.Name))
	if len(c.Implements) > 0 {
		sb.WriteString(" : ")
		sb.WriteString(strings.Join(c.Implements, ", "))
	}
	sb.WriteString("\n{\n")

	// Properties
	for _, prop := range c.Properties {
		if prop.Description != "" {
			sb.WriteString("    /// <summary>\n")
			sb.WriteString(fmt.Sprintf("    /// %s\n", prop.Description))
			sb.WriteString("    /// </summary>\n")
		}

		if prop.JsonName != "" && prop.JsonName != prop.Name {
			sb.WriteString(fmt.Sprintf("    [JsonPropertyName(\"%s\")]\n", prop.JsonName))
		}

		typeName := prop.Type
		if prop.IsNullable && !strings.HasSuffix(typeName, "?") {
			typeName += "?"
		}

		sb.WriteString(fmt.Sprintf("    public %s %s { get; set; }\n\n", typeName, prop.Name))
	}

	sb.WriteString("}\n")

	return sb.String()
}

// CSharpEnum represents a C# enum
type CSharpEnum struct {
	Name        string
	Description string
	Values      []CSharpEnumValue
}

// CSharpEnumValue represents an enum value
type CSharpEnumValue struct {
	Name  string
	Value string
}

// NewCSharpEnum creates a new C# enum
func NewCSharpEnum(name, description string) *CSharpEnum {
	return &CSharpEnum{
		Name:        name,
		Description: description,
		Values:      []CSharpEnumValue{},
	}
}

// AddValue adds a value to the enum
func (e *CSharpEnum) AddValue(name, value string) {
	e.Values = append(e.Values, CSharpEnumValue{Name: name, Value: value})
}

// Generate generates the C# enum code
func (e *CSharpEnum) Generate() string {
	var sb strings.Builder

	if e.Description != "" {
		sb.WriteString("/// <summary>\n")
		sb.WriteString(fmt.Sprintf("/// %s\n", e.Description))
		sb.WriteString("/// </summary>\n")
	}

	sb.WriteString(fmt.Sprintf("public enum %s\n{\n", e.Name))

	for i, val := range e.Values {
		sb.WriteString(fmt.Sprintf("    [JsonPropertyName(\"%s\")]\n", val.Value))
		sb.WriteString(fmt.Sprintf("    %s", val.Name))
		if i < len(e.Values)-1 {
			sb.WriteString(",")
		}
		sb.WriteString("\n")
	}

	sb.WriteString("}\n")

	return sb.String()
}

// GenerateInterfaces generates the marker interfaces for steps and notifications
func GenerateInterfaces() string {
	return `/// <summary>
/// Marker interface for all pipeline step types.
/// </summary>
[JsonDerivedType(typeof(CommandStep))]
[JsonDerivedType(typeof(BlockStep))]
[JsonDerivedType(typeof(InputStep))]
[JsonDerivedType(typeof(WaitStep))]
[JsonDerivedType(typeof(TriggerStep))]
[JsonDerivedType(typeof(GroupStep))]
public interface IStep { }

/// <summary>
/// Marker interface for all notification types.
/// </summary>
public interface INotification { }

/// <summary>
/// Marker interface for steps that can appear inside a group.
/// </summary>
[JsonDerivedType(typeof(CommandStep))]
[JsonDerivedType(typeof(BlockStep))]
[JsonDerivedType(typeof(InputStep))]
[JsonDerivedType(typeof(WaitStep))]
[JsonDerivedType(typeof(TriggerStep))]
public interface IGroupStep : IStep { }
`
}
