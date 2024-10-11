package typescript_code_gen

import "encoding/json"

type typescriptPackageJSON struct {
	Name            string            `json:"name"`
	Version         string            `json:"version"`
	Description     string            `json:"description"`
	Main            string            `json:"main"`
	Scripts         map[string]string `json:"scripts"`
	Keywords        []string          `json:"keywords"`
	Author          string            `json:"author"`
	License         string            `json:"license"`
	DevDependencies map[string]string `json:"devDependencies"`
}

func newPackageJSONFile() string {
	// Write package.json
	packageJSON := typescriptPackageJSON{
		Name:        "buildkite-pipeline-sdk",
		Version:     "0.0.1",
		Description: "",
		Main:        "index.ts",
		Scripts: map[string]string{
			"test": "echo \"Error: no test specified\" && exit 1",
		},
		Keywords: []string{},
		Author:   "",
		License:  "ISC",
		DevDependencies: map[string]string{
			"@types/node": "^20.11.30",
		},
	}

	packageJSONBytes, err := json.MarshalIndent(packageJSON, "", "\t")
	if err != nil {
		return "{}"
	}

	return string(packageJSONBytes)
}
