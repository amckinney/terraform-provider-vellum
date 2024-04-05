// This file was auto-generated by Fern from our API Definition.

package vellum

type UpsertSandboxScenarioRequestRequest struct {
	Label *string `json:"label,omitempty" url:"label,omitempty"`
	// The inputs for the scenario
	Inputs []*ScenarioInputRequest `json:"inputs,omitempty" url:"inputs,omitempty"`
	// The id of the scenario to update. If none is provided, an id will be generated and a new scenario will be appended.
	ScenarioId *string `json:"scenario_id,omitempty" url:"scenario_id,omitempty"`
}
