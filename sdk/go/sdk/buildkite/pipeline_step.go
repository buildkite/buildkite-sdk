package buildkite

import (
	"github.com/buildkite/buildkite-sdk/sdk/go/sdk/schema"
)

// Concurrency Methods
type conncurrencyMethod struct {
	Eager   schema.ConcurrencyMethod
	Ordered schema.ConcurrencyMethod
}

var ConcurrencyMethod = conncurrencyMethod{
	Eager:   schema.Eager,
	Ordered: schema.Ordered,
}

// Cache
type Cache interface {
	toSchemaCache() *schema.Cache
}

type CacheString string

func (c CacheString) toSchemaCache() *schema.Cache {
	key := string(c)
	return &schema.Cache{
		String: &key,
	}
}

type CacheStringArray []string

func (c CacheStringArray) toSchemaCache() *schema.Cache {
	return &schema.Cache{
		StringArray: c,
	}
}

type CacheObject schema.CacheClass

func (c CacheObject) toSchemaCache() *schema.Cache {
	return &schema.Cache{
		CacheClass: &schema.CacheClass{
			Name:  c.Name,
			Paths: c.Paths,
			Size:  c.Size,
		},
	}
}

// Depends On
type DependsOn interface {
	toSchema() *schema.DependsOn
}

type DependsOnString string

func (d DependsOnString) toSchema() *schema.DependsOn {
	key := string(d)
	return &schema.DependsOn{
		String: &key,
	}
}

type DependsOnStringArray []string

func (d DependsOnStringArray) toSchema() *schema.DependsOn {
	items := make([]schema.DependsOnElement, len(d))
	for i, item := range d {
		items[i] = schema.DependsOnElement{
			String: &item,
		}
	}

	return &schema.DependsOn{
		UnionArray: items,
	}
}

type dependsOnItem struct {
	AllowFailure bool
	Step         string
}

type DependsOnItemArray []dependsOnItem

func (d DependsOnItemArray) toSchema() *schema.DependsOn {
	items := make([]schema.DependsOnElement, len(d))
	for i, item := range d {
		items[i] = schema.DependsOnElement{
			DependsOnClass: &schema.DependsOnClass{
				Step: &item.Step,
				AllowFailure: &schema.AllowDependencyFailureUnion{
					Bool: &item.AllowFailure,
				},
			},
		}
	}

	return &schema.DependsOn{
		UnionArray: items,
	}
}

// Matrix
type Matrix interface {
	toSchema() *schema.MatrixUnion
}

type MatrixSimple []string

func (s MatrixSimple) toSchema() *schema.MatrixUnion {
	items := make([]schema.MatrixElement, len(s))
	for i, item := range s {
		items[i] = schema.MatrixElement{
			String: &item,
		}
	}

	return &schema.MatrixUnion{
		UnionArray: items,
	}
}

type MatrixAdvancedSetup map[string][]string

type MatrixAdvancedAdjustment struct {
	With     map[string]string
	Skip     bool
	SoftFail bool
}

type MatrixAdvanced struct {
	Setup       MatrixAdvancedSetup
	Adjustments []MatrixAdvancedAdjustment
}

func (m MatrixAdvanced) toSchema() *schema.MatrixUnion {
	setup := make(map[string][]schema.MatrixElement)
	for key, vals := range m.Setup {
		setupVals := make([]schema.MatrixElement, len(vals))
		for i, item := range vals {
			setupVals[i] = schema.MatrixElement{
				String: &item,
			}
		}

		setup[key] = setupVals
	}

	adjustments := make([]schema.Adjustment, len(m.Adjustments))
	for i, item := range m.Adjustments {
		adjustments[i] = schema.Adjustment{
			Skip: &schema.Skip{
				Bool: &item.Skip,
			},
			SoftFail: &schema.SoftFail{
				Bool: &item.SoftFail,
			},
			With: &schema.With{
				StringMap: item.With,
			},
		}
	}

	return &schema.MatrixUnion{
		MatrixClass: &schema.MatrixClass{
			Setup: &schema.Setup{
				UnionArrayMap: setup,
			},
			Adjustments: adjustments,
		},
	}
}

// Retry
type Retry interface {
	toSchema() *schema.Retry
}

// Signature
type Signature struct {
	Algorithm    string
	SignedFields []string
	Value        string
}

func (s Signature) toSchema() *schema.Signature {
	return &schema.Signature{
		Algorithm:    &s.Algorithm,
		SignedFields: s.SignedFields,
		Value:        &s.Value,
	}
}

// Soft Fail
type SoftFail interface {
	toSchema() *schema.SoftFail
}

type SoftFailSimple bool

func (s SoftFailSimple) toSchema() *schema.SoftFail {
	val := bool(s)
	return &schema.SoftFail{
		Bool: &val,
	}
}

type SoftFailAdvancedItem struct {
	ExitStatus int64
}

type SoftFailAdvanced []SoftFailAdvancedItem

func (s SoftFailAdvanced) toSchema() *schema.SoftFail {
	items := make([]schema.SoftFailElement, len(s))
	for i, val := range s {
		items[i] = schema.SoftFailElement{
			ExitStatus: &schema.SoftFailExitStatus{
				Integer: &val.ExitStatus,
			},
		}
	}

	return &schema.SoftFail{
		SoftFailElementArray: items,
	}
}

// TODO: make these more user friendly
type CommandNotify schema.NotifyElement

type PipelineStep struct {
	AllowDependencyFailure *schema.AllowDependencyFailureUnion `json:"allow_dependency_failure,omitempty"`
	Block                  *schema.Block                       `json:"block,omitempty"`
	BlockedState           *schema.BlockedState                `json:"blocked_state,omitempty"`
	Branches               *schema.Branches                    `json:"branches,omitempty"`
	DependsOn              *schema.DependsOn                   `json:"depends_on,omitempty"`
	Fields                 []schema.Field                      `json:"fields,omitempty"`
	ID                     *string                             `json:"id,omitempty"`
	Identifier             *string                             `json:"identifier,omitempty"`
	If                     *string                             `json:"if,omitempty"`
	Key                    *string                             `json:"key,omitempty"`
	Label                  *string                             `json:"label,omitempty"`
	Name                   *string                             `json:"name,omitempty"`
	Prompt                 *string                             `json:"prompt,omitempty"`
	Type                   *schema.BlockStepType               `json:"type,omitempty"`
	Input                  *schema.Input                       `json:"input,omitempty"`
	Agents                 *schema.Agents                      `json:"agents,omitempty"`
	ArtifactPaths          []string                            `json:"artifact_paths,omitempty"`
	Cache                  *schema.Cache                       `json:"cache,omitempty"`
	CancelOnBuildFailing   *schema.AllowDependencyFailureUnion `json:"cancel_on_build_failing,omitempty"`
	Command                *schema.CommandUnion                `json:"command,omitempty"`
	Commands               *schema.CommandUnion                `json:"commands,omitempty"`
	Concurrency            *int64                              `json:"concurrency,omitempty"`
	ConcurrencyGroup       *string                             `json:"concurrency_group,omitempty"`
	ConcurrencyMethod      *schema.ConcurrencyMethod           `json:"concurrency_method,omitempty"`
	Env                    map[string]interface{}              `json:"env,omitempty"`
	Matrix                 *schema.MatrixUnion                 `json:"matrix,omitempty"`
	Notify                 []schema.BlockStepNotify            `json:"notify,omitempty"`
	Parallelism            *int64                              `json:"parallelism,omitempty"`
	Plugins                *schema.Plugins                     `json:"plugins,omitempty"`
	Priority               *int64                              `json:"priority,omitempty"`
	Retry                  *schema.Retry                       `json:"retry,omitempty"`
	Signature              *schema.Signature                   `json:"signature,omitempty"`
	Skip                   *schema.Skip                        `json:"skip,omitempty"`
	SoftFail               *schema.SoftFail                    `json:"soft_fail,omitempty"`
	TimeoutInMinutes       *int64                              `json:"timeout_in_minutes,omitempty"`
	Script                 *schema.CommandStep                 `json:"script,omitempty"`
	ContinueOnFailure      *schema.AllowDependencyFailureUnion `json:"continue_on_failure,omitempty"`
	Wait                   *schema.Label                       `json:"wait,omitempty"`
	Waiter                 *schema.WaitStep                    `json:"waiter,omitempty"`
	Async                  *schema.AllowDependencyFailureUnion `json:"async,omitempty"`
	Build                  *schema.Build                       `json:"build,omitempty"`
	Trigger                *schema.Trigger                     `json:"trigger,omitempty"`
	Group                  *string                             `json:"group,omitempty"`
	Steps                  []PipelineStep                      `json:"steps,omitempty"`
}
