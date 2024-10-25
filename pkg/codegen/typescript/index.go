package typescript_code_gen

var indexFile = `// This file is auto-generated please do not edit

import Environment from "./environment";
import StepBuilder from "./stepBuilder";

export * as types from "./types";

export { Environment, StepBuilder };`

func newIndexFile() string {
	return indexFile
}
