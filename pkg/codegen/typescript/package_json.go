package typescript_code_gen

import "encoding/json"

type typescriptPackageJSON struct {
	Name            string            `json:"name"`
	Version         string            `json:"version"`
	Description     string            `json:"description"`
	Main            string            `json:"main"`
	Types           string            `json:"types"`
	Scripts         map[string]string `json:"scripts"`
	Keywords        []string          `json:"keywords"`
	Author          string            `json:"author"`
	License         string            `json:"license"`
	DevDependencies map[string]string `json:"devDependencies"`
}

func newPackageJSONFile(version string) string {
	// Write package.json
	packageJSON := typescriptPackageJSON{
		Name:        "buildkite-pipeline-sdk",
		Version:     version,
		Description: "",
		Main:        "dist/index.js",
		Types:       "dist/index.d.ts",
		Scripts: map[string]string{
			"build": "tsc",
		},
		Keywords: []string{},
		Author:   "",
		License:  "MIT",
		DevDependencies: map[string]string{
			"@types/node": "^20.11.30",
			"typescript":  "^5.6.3",
		},
	}

	packageJSONBytes, err := json.MarshalIndent(packageJSON, "", "\t")
	if err != nil {
		return "{}"
	}

	return string(packageJSONBytes)
}
