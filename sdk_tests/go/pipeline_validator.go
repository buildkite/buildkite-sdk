package sdk_tests_go

import (
	"encoding/json"
	"strings"
	"testing"

	bk "github.com/buildkite/pipeline-sdk/sdk/go"
	"github.com/stretchr/testify/assert"

	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	"github.com/santhosh-tekuri/jsonschema/v6"
)

type HTTPURLLoader http.Client

func (l *HTTPURLLoader) Load(url string) (any, error) {
	client := (*http.Client)(l)
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		_ = resp.Body.Close()
		return nil, fmt.Errorf("%s returned status code %d", url, resp.StatusCode)
	}
	defer resp.Body.Close()

	return jsonschema.UnmarshalJSON(resp.Body)
}

func newHTTPURLLoader(insecure bool) *HTTPURLLoader {
	httpLoader := HTTPURLLoader(http.Client{
		Timeout: 15 * time.Second,
	})
	if insecure {
		httpLoader.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}
	return &httpLoader
}

type pipelineSchemaValidator struct {
	schema *jsonschema.Schema
}

func (p pipelineSchemaValidator) ValidatePipeline(t *testing.T, pipeline *bk.StepBuilder) {
	pipelineJSON, err := json.Marshal(pipeline)
	assert.NoError(t, err)

	j, err := jsonschema.UnmarshalJSON(strings.NewReader(string(pipelineJSON)))
	assert.NoError(t, err)

	err = p.schema.Validate(j)
	if err != nil {
		fmt.Println(string(pipelineJSON))
	}
	assert.NoError(t, err)
}

func newPipelineSchemaValidator(t *testing.T) pipelineSchemaValidator {
	schemaURL := "https://raw.githubusercontent.com/buildkite/pipeline-schema/refs/heads/main/schema.json"

	loader := jsonschema.SchemeURLLoader{
		"file":  jsonschema.FileLoader{},
		"http":  newHTTPURLLoader(false),
		"https": newHTTPURLLoader(false),
	}

	c := jsonschema.NewCompiler()
	c.UseLoader(loader)
	sch, err := c.Compile(schemaURL)
	assert.NoError(t, err)

	return pipelineSchemaValidator{schema: sch}
}
